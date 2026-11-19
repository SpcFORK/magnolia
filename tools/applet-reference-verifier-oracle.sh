#!/usr/bin/env bash
# Java 8 -Xverify:all reference-type oracle. Each mutant is one same-width
# bytecode/constant-pool-index edit in a real javac class. In particular,
# aastore accepts any reference during verification and checks the element
# dynamically: rejecting that instruction at preflight would be incorrect.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for cmd in "$MAGNOLIA" java javac node python3 timeout diff cmp; do
    command -v "$cmd" >/dev/null 2>&1 || { echo "missing: $cmd" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/mutants"
run() {
    local label="$1" seconds="$2" start=$SECONDS rc
    shift 2
    printf '[ref-verifier] START %s (max %ss)\n' "$label" "$seconds" >&2
    if timeout --signal=TERM --kill-after=2s "$seconds" "$@"; then
        printf '[ref-verifier] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[ref-verifier] FAIL %s exit=%s (%ss)\n' "$label" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
classes=(RefSubtypeMain RefParent RefChild RefSibling RefOther RefScore)
run 'compile Java 8 sources' 20 javac --release 8 -g:none -d "$scratch/original" test/fixtures/RefSubtypeMain.java
for name in "${classes[@]}"; do
    cmp "test/fixtures/$name.class" "$scratch/original/$name.class"
done
run 'independent verified JDK output' 10 java -Xverify:all -cp "$scratch/original" RefSubtypeMain > "$scratch/jdk.out"
diff -u test/fixtures/RefSubtypeMain.expected "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 13 ]]
inputs=()
for name in "${classes[@]}"; do inputs+=("$scratch/original/$name.class"); done
run 'verified native Oak output' 15 "$MAGNOLIA" build run \
    --entry "${inputs[0]}" "${inputs[@]:1}" --class RefSubtypeMain > "$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"

# One local multi-class JAR exercises Builder's shared linker and keeps all
# historical/third-party binaries out of the test and patch.
python3 - "$scratch/original" "$scratch/RefSubtypeMain.jar" <<'PY'
from pathlib import Path
from zipfile import ZipFile, ZIP_DEFLATED
import sys
src, target = Path(sys.argv[1]), sys.argv[2]
with ZipFile(target, 'w', ZIP_DEFLATED) as jar:
    jar.writestr('META-INF/MANIFEST.MF',
                 'Manifest-Version: 1.0\r\nMain-Class: RefSubtypeMain\r\n\r\n')
    for name in ('RefSubtypeMain','RefParent','RefChild','RefSibling','RefOther','RefScore'):
        jar.write(src/(name+'.class'), name+'.class')
PY
run 'multi-class JAR native route' 20 "$MAGNOLIA" build run \
    --entry "$scratch/RefSubtypeMain.jar" > "$scratch/jar.out"
diff -u "$scratch/jdk.out" "$scratch/jar.out"
run 'Builder standalone Oak' 65 "$MAGNOLIA" build \
    --entry "$scratch/RefSubtypeMain.jar" --output "$scratch/RefSubtypeMain.oak"
run 'execute standalone Oak' 20 "$MAGNOLIA" "$scratch/RefSubtypeMain.oak" > "$scratch/standalone.out"
diff -u "$scratch/jdk.out" "$scratch/standalone.out"
run 'restore original JAR' 25 "$MAGNOLIA" build restore \
    --entry "$scratch/RefSubtypeMain.oak" --output "$scratch/restored.jar"
cmp "$scratch/RefSubtypeMain.jar" "$scratch/restored.jar"
run 'Builder bundled JS' 110 "$MAGNOLIA" build \
    --entry "$scratch/RefSubtypeMain.jar" --output "$scratch/RefSubtypeMain.js"
run 'execute bundled JS' 25 node "$scratch/RefSubtypeMain.js" > "$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run 'Builder offline HTML' 110 "$MAGNOLIA" build \
    --entry "$scratch/RefSubtypeMain.jar" --output "$scratch/RefSubtypeMain.html"
run 'execute offline HTML' 20 node - "$scratch/RefSubtypeMain.html" "$scratch/jdk.out" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2);
const output = {textContent: ''}, canvas = {style: {}};
const ctx = {document: {getElementById(id) {return id === 'output' ? output :
    id === 'drawing' ? canvas : null;}},
    console: {log() {}, error(...args) {throw Error(args.join(' '));}}};
vm.createContext(ctx);
for (const script of scripts) vm.runInContext(script, ctx, {timeout: 10000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(ctx.MagnoliaApplet.inspect().type, 'ok');
JS
run 'Builder generated Java' 60 "$MAGNOLIA" build \
    --entry "$scratch/RefSubtypeMain.jar" --output "$scratch/RefSubtypeConverted.java"
mkdir -p "$scratch/generated"
run 'compile generated Java' 25 javac -d "$scratch/generated" "$scratch/RefSubtypeConverted.java"
run 'execute generated Java' 12 java -Xverify:all -cp "$scratch/generated" \
    RefSubtypeConverted > "$scratch/generated.out"
diff -u "$scratch/jdk.out" "$scratch/generated.out"

# The parser walks only the fixture class's pool and Code attributes; every
# edit is asserted against a named method, not an arbitrary occurrence in the
# class file. Keep the original and all dependency classes in each variant.
python3 - "$scratch/original" "$scratch/mutants" <<'PY'
from pathlib import Path
import sys, shutil
src, dest = map(Path, sys.argv[1:])
raw = (src/'RefSubtypeMain.class').read_bytes()
pos = 0
def take(n):
    global pos
    chunk = raw[pos:pos+n]
    assert len(chunk) == n, 'truncated fixture'
    pos += n
    return chunk
def u1(): return take(1)[0]
def u2(): return int.from_bytes(take(2), 'big')
def u4(): return int.from_bytes(take(4), 'big')
def be2(n): return n.to_bytes(2, 'big')
assert take(4) == b'\xca\xfe\xba\xbe'
take(4)
pool = [None]*u2()
i = 1
while i < len(pool):
    tag = u1()
    if tag == 1: pool[i] = (tag, take(u2()).decode('utf-8'))
    elif tag in (3,4): take(4)
    elif tag in (5,6): take(8); i += 1
    elif tag in (7,8,16,19,20): pool[i] = (tag, u2())
    elif tag in (9,10,11,12,17,18): pool[i] = (tag, u2(),u2())
    elif tag == 15: take(3)
    else: raise ValueError(f'unknown fixture tag {tag}')
    i += 1

def utf(i):
    assert pool[i][0] == 1
    return pool[i][1]
def clazz(i):
    assert pool[i][0] == 7
    return utf(pool[i][1])
def member_index(tag, owner, name, desc):
    matches = [i for i, e in enumerate(pool) if e and e[0] == tag and
               clazz(e[1]) == owner and
               (utf(pool[e[2]][1]), utf(pool[e[2]][2])) == (name, desc)]
    assert len(matches) == 1, (tag, owner, name, desc, matches)
    return matches[0]
def class_index(name):
    matches = [i for i, e in enumerate(pool) if e and e[0] == 7 and clazz(i) == name]
    assert len(matches) == 1, (name, matches)
    return matches[0]
def attributes():
    found = []
    for _ in range(u2()):
        name, length = utf(u2()), u4()
        start = pos
        take(length)
        found.append((name, start, length))
    return found

take(6)
take(2*u2())
for _ in range(u2()):
    take(6)
    attributes()
methods = {}
for _ in range(u2()):
    take(2)
    name, desc = utf(u2()), utf(u2())
    for attr, start, size in attributes():
        if attr == 'Code':
            length = int.from_bytes(raw[start+4:start+8], 'big')
            assert 8+length <= size
            methods[(name,desc)] = (start+8, raw[start+8:start+8+length])

files = ['RefSubtypeMain','RefParent','RefChild','RefSibling','RefOther','RefScore']
def patch(label, name, desc, pattern, offset, before, after):
    assert len(before) == len(after)
    start, code = methods[(name,desc)]
    assert code.count(pattern) == 1, (label, code.count(pattern), code.hex())
    at = start + code.index(pattern) + offset
    assert raw[at:at+len(before)] == before, label
    output = bytearray(raw)
    output[at:at+len(before)] = after
    folder = dest/label
    folder.mkdir()
    for file in files[1:]: shutil.copyfile(src/(file+'.class'), folder/(file+'.class'))
    (folder/'RefSubtypeMain.class').write_bytes(output)

main = '([Ljava/lang/String;)V'
parent = member_index(10,'RefSubtypeMain','fromParent','(LRefParent;)I')
score = member_index(10,'RefSubtypeMain','fromScore','(LRefScore;)I')
array = member_index(10,'RefSubtypeMain','fromArray','([LRefParent;)I')
from_object = member_index(10,'RefSubtypeMain','fromObject','(Ljava/lang/Object;)I')
field = member_index(9,'RefSubtypeMain','saved','LRefParent;')
patch('sibling-call','main',main,b'\x2b\xb8'+be2(parent),0,b'\x2b',b'\x2d')
patch('interface-call','main',main,b'\x2c\xb8'+be2(score),0,b'\x2c',b'\x2d')
patch('array-call','main',main,b'\x19\x05\xb8'+be2(array),1,b'\x05',b'\x03')
patch('field-store','stash','(LRefParent;LRefOther;)V',
      b'\x2a\xb3'+be2(field),0,b'\x2a',b'\x2b')
patch('reference-return','choose','(LRefParent;LRefOther;)LRefParent;',
      b'\x2a\xb0',0,b'\x2a',b'\x2b')
patch('merged-object-call','join','(ZLRefChild;LRefOther;)I',
      b'\xb8'+be2(from_object),1,be2(from_object),be2(parent))
# A reference-array store with a statically incompatible value MUST verify:
# the actual array's component is checked at execution, not at verification.
patch('runtime-array-store','main',main,
      b'\x19\x05\x03\x2c\x53',3,b'\x2c',b'\x2d')
PY

# The interface-call edit is NOT a verifier error on the reference JDK:
# bytecode may pass a class-typed reference to an interface parameter and
# receive IncompatibleClassChangeError at dynamic invocation instead.
negatives=(sibling-call array-call field-store reference-return merged-object-call)
for label in "${negatives[@]}"; do
    folder="$scratch/mutants/$label"
    printf '[ref-verifier] START negative %s, two preflight oracles (caps 8/8s)\n' "$label" >&2
    if timeout --signal=TERM --kill-after=2s 8s java -Xverify:all -cp "$folder" RefSubtypeMain \
        > "$scratch/$label.jdk" 2>&1; then
        echo "FAIL JDK accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || { echo "JDK timeout: $label" >&2; exit 1; }
    fi
    grep -q 'VerifyError' "$scratch/$label.jdk" || {
        echo "JDK failure was not verification: $label" >&2
        cat "$scratch/$label.jdk" >&2; exit 1;
    }
    if timeout --signal=TERM --kill-after=2s 8s "$MAGNOLIA" build run \
        --entry "$folder/RefSubtypeMain.class" \
        "$folder/RefParent.class" "$folder/RefChild.class" \
        "$folder/RefSibling.class" "$folder/RefOther.class" \
        "$folder/RefScore.class" --class RefSubtypeMain \
        > "$scratch/$label.oak" 2>&1; then
        echo "FAIL Magnolia accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || { echo "Oak timeout: $label" >&2; exit 1; }
    fi
    grep -q 'reference subtype' "$scratch/$label.oak" || {
        echo "unexpected Magnolia rejection: $label" >&2
        cat "$scratch/$label.oak" >&2; exit 1;
    }
    if grep -E '^[0-9]+$' "$scratch/$label.oak"; then
        echo "FAIL: $label printed output before preflight" >&2; exit 1
    fi
    printf '[ref-verifier] PASS %s JDK VerifyError = Oak preflight rejection\n' "$label" >&2
done

runtime_case() {
    local label="$1" jdk_error="$2" oak_error="$3" lines="$4" folder="$scratch/mutants/$1" rc
    printf '[ref-verifier] START JDK-valid dynamic %s (caps 8/8s)\n' "$label" >&2
    if timeout --signal=TERM --kill-after=2s 8s java -Xverify:all -cp "$folder" RefSubtypeMain \
        > "$scratch/$label.jdk" 2>&1; then
        echo "FAIL JDK did not throw at $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -q "$jdk_error" "$scratch/$label.jdk"
    ! grep -q 'VerifyError' "$scratch/$label.jdk"
    if timeout --signal=TERM --kill-after=2s 8s "$MAGNOLIA" build run \
        --entry "$folder/RefSubtypeMain.class" "$folder/RefParent.class" \
        "$folder/RefChild.class" "$folder/RefSibling.class" \
        "$folder/RefOther.class" "$folder/RefScore.class" \
        --class RefSubtypeMain > "$scratch/$label.oak" 2>&1; then
        echo "FAIL Magnolia did not throw at $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -q "$oak_error" "$scratch/$label.oak"
    head -n "$lines" "$scratch/$label.jdk" > "$scratch/prefix.jdk"
    head -n "$lines" "$scratch/$label.oak" > "$scratch/prefix.oak"
    diff -u "$scratch/prefix.jdk" "$scratch/prefix.oak"
    printf '[ref-verifier] PASS %s runtime error after %s matching lines\n' "$label" "$lines" >&2
}
runtime_case runtime-array-store java.lang.ArrayStoreException java/lang/ArrayStoreException 5
runtime_case interface-call java.lang.IncompatibleClassChangeError java/lang/IncompatibleClassChangeError 4
printf '[ref-verifier] PASS 13 JDK lines = native/JAR/Oak/JS/offline HTML/generated Java; 5 JDK-rejected preflight cases, 2 JDK-accepted runtime errors, exact JAR restore\n'
