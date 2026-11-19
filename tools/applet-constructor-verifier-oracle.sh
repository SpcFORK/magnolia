#!/usr/bin/env bash
# Bounded constructor/allocation-site verifier regression, independently
# checked against Java 8 class files under JDK -Xverify:all. Mutants alter
# only named Code payloads, preserving lengths and constant-pool indexes.
# This does NOT claim complete JVM verification or StackMapTable enforcement.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" javac java python3 node timeout diff cmp; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing constructor-verifier dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/cases"

javac --release 8 -g:none -d "$scratch/original" \
    test/fixtures/ConstructorFrameMain.java
cmp test/fixtures/ConstructorFrameMain.class \
    "$scratch/original/ConstructorFrameMain.class"
timeout 10s java -Xverify:all -cp "$scratch/original" \
    ConstructorFrameMain > "$scratch/jdk.out"
diff -u test/fixtures/ConstructorFrameMain.expected "$scratch/jdk.out"
timeout 15s "$MAGNOLIA" build run \
    --entry "$scratch/original/ConstructorFrameMain.class" > "$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"

# A valid super() and this() delegation, plus fresh-object aliases through
# dup/astore/aload, must survive both standalone backends and exact restore.
timeout 60s "$MAGNOLIA" build --entry \
    "$scratch/original/ConstructorFrameMain.class" \
    --output "$scratch/ConstructorFrameMain.oak" > "$scratch/oak-build.log"
timeout 20s "$MAGNOLIA" "$scratch/ConstructorFrameMain.oak" > "$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
timeout 110s "$MAGNOLIA" build --entry \
    "$scratch/original/ConstructorFrameMain.class" \
    --output "$scratch/ConstructorFrameMain.js" > "$scratch/js-build.log"
timeout 20s node "$scratch/ConstructorFrameMain.js" > "$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
timeout 15s "$MAGNOLIA" build restore \
    --entry "$scratch/ConstructorFrameMain.oak" \
    --output "$scratch/restored.class" > "$scratch/restore.log"
cmp "$scratch/original/ConstructorFrameMain.class" "$scratch/restored.class"
printf '[constructor-verifier] PASS Java 8 source = pinned class; JDK = native/Oak/JS; exact restore\n'

# Decode just enough class-file metadata to assert the *named* target and
# its constant-pool reference. No JDK-produced mutant is treated as valid
# merely because it can be parsed: every negative must cause VerifyError.
python3 - "$scratch/original/ConstructorFrameMain.class" "$scratch/cases" <<'PY'
from pathlib import Path
import sys
raw = Path(sys.argv[1]).read_bytes()
out = Path(sys.argv[2])
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
assert take(4) == b'\x00\x00\x00\x34', 'expected Java 8 class'
pool = [None] * u2()
i = 1
while i < len(pool):
    tag = u1()
    if tag == 1: pool[i] = (tag, take(u2()).decode('utf-8'))
    elif tag in (3,4): take(4)
    elif tag in (5,6): take(8); i += 1
    elif tag in (7,8,16,19,20): pool[i] = (tag,u2())
    elif tag in (9,10,11,12,17,18): pool[i] = (tag,u2(),u2())
    elif tag == 15: take(3)
    else: raise ValueError(f'unsupported fixture pool tag {tag}')
    i += 1
def utf(i):
    assert pool[i][0] == 1
    return pool[i][1]
def classname(i):
    assert pool[i][0] == 7
    return utf(pool[i][1])
def member(owner, name, desc):
    result = [i for i,e in enumerate(pool) if e and e[0] == 10 and
              classname(e[1]) == owner and
              (utf(pool[e[2]][1]),utf(pool[e[2]][2])) == (name,desc)]
    assert len(result) == 1, (owner,name,desc,result)
    return result[0]
def clazz(name):
    result = [i for i,e in enumerate(pool) if e and e[0] == 7 and
              classname(i) == name]
    assert len(result) == 1, (name,result)
    return result[0]
def attributes():
    result = []
    for _ in range(u2()):
        name,size = utf(u2()),u4()
        start = pos
        take(size)
        result.append((name,start,size))
    return result

take(6)
take(2*u2())  # interfaces
for _ in range(u2()): take(6); attributes()
methods = {}
for _ in range(u2()):
    take(2)
    name,desc = utf(u2()),utf(u2())
    for kind,start,size in attributes():
        if kind == 'Code':
            code_length = int.from_bytes(raw[start+4:start+8], 'big')
            assert 8 + code_length <= size
            methods[(name,desc)] = (start+8, raw[start+8:start+8+code_length])

object_init = member('java/lang/Object','<init>','()V')
this_init = member('ConstructorFrameMain','<init>','()V')
call_two = member('ConstructorFrameMain','two','()Ljava/lang/Object;')
object_class = clazz('java/lang/Object')
this_class = clazz('ConstructorFrameMain')

def patch(label, name, desc, original, edited, nth=0, count=1):
    assert len(original) == len(edited), label
    start,code = methods[(name,desc)]
    assert code.count(original) == count, (label,code.count(original),code.hex())
    at = -1
    for _ in range(nth+1): at = code.index(original,at+1)
    result = bytearray(raw)
    assert result[start+at:start+at+len(original)] == original
    result[start+at:start+at+len(original)] = edited
    target = out/label
    target.mkdir()
    (target/'ConstructorFrameMain.class').write_bytes(result)
    print(f'[constructor-verifier] {label}: {name}{desc} @ byte {at}',file=sys.stderr)

# The no-arg constructor is called *after* prior prints by main, but must
# fail preflight before even the first line is emitted.
patch('missing-super','<init>','()V',
      b'\xb7'+be2(object_init),b'\x57\x00\x00')
patch('discarded-this','<init>','()V',
      b'\x2a\xb7'+be2(object_init),b'\x01\x4b\x00\x00')
patch('missing-this','<init>','(I)V',
      b'\xb7'+be2(this_init),b'\x57\x00\x00')
# fresh() is called only after two successful print operations. The copy in
# its local and the original dup alias both retain the uninitialized token.
patch('uninitialized-return','fresh','()Ljava/lang/Object;',
      b'\xb7'+be2(object_init),b'\x57\x00\x00')
# Two distinct `new Object` sites must retain distinct U tokens: initializing
# the second one cannot initialize the first object's local alias.
patch('other-object-initialized','two','()Ljava/lang/Object;',
      b'\xb7'+be2(object_init),b'\x57\x00\x00',count=2)
# Replace the second allocation with a load of the initialized first
# object, then try to invoke Object.<init> on that initialized alias.
patch('constructor-twice','two','()Ljava/lang/Object;',
      b'\xbb'+be2(object_class),b'\x2a\x00\x00',nth=1,count=2)
# A freshly allocated ConstructorFrameMain must invoke ITS constructor,
# not Object.<init> directly; the class is linked but JDK verification fails.
main = '([Ljava/lang/String;)V'
late = b'\xbb'+be2(this_class)+b'\x59\xb7'+be2(this_init)
patch('wrong-owner','main',main,late,
      b'\xbb'+be2(this_class)+b'\x59\xb7'+be2(object_init))
patch('initializer-virtual','main',main,late,
      b'\xbb'+be2(this_class)+b'\x59\xb6'+be2(this_init))
# The JDK DOES allow uninitialized references in identity and null tests:
# these are positive counterexamples to an overzealous category check.
comparison = b'\xbb'+be2(object_class)+b'\x59\xb7'+be2(object_init)+b'\xa5'
patch('compare-uninitialized','main',main,comparison,
      b'\xbb'+be2(object_class)+b'\x59\x57\x00\x00\xa5')
call_and_null = b'\xb8'+be2(call_two)+b'\xc6'
patch('ifnull-uninitialized','main',main,call_and_null,
      b'\xbb'+be2(object_class)+b'\xc6')
patch('ifnonnull-uninitialized','main',main,call_and_null,
      b'\xbb'+be2(object_class)+b'\xc7')
PY

for label in compare-uninitialized ifnull-uninitialized ifnonnull-uninitialized; do
    folder="$scratch/cases/$label"
    timeout 10s java -Xverify:all -cp "$folder" ConstructorFrameMain \
        > "$scratch/$label.jdk"
    timeout 15s "$MAGNOLIA" build run \
        --entry "$folder/ConstructorFrameMain.class" > "$scratch/$label.oak"
    diff -u "$scratch/$label.jdk" "$scratch/$label.oak"
    if [[ "$label" == ifnonnull-uninitialized ]]; then
        printf '7\n1\ntrue\nfalse\n0\n' > "$scratch/positive.expected"
    else
        cp test/fixtures/ConstructorFrameMain.expected "$scratch/positive.expected"
    fi
    diff -u "$scratch/positive.expected" "$scratch/$label.jdk"
    printf '[constructor-verifier] PASS JDK-accepted %s = Magnolia native\n' "$label"
done

negative_cases=(
    'missing-super|constructor returns before this()/super() initialization|<init>()V'
    'discarded-this|constructor returns before this()/super() initialization|<init>()V'
    'missing-this|constructor returns before this()/super() initialization|<init>(I)V'
    'uninitialized-return|uninitialized reference used before constructor invocation|fresh()Ljava/lang/Object;'
    'other-object-initialized|uninitialized reference used before constructor invocation|two()Ljava/lang/Object;'
    'constructor-twice|constructor requires its uninitialized receiver|two()Ljava/lang/Object;'
    'wrong-owner|constructor requires its uninitialized receiver|main([Ljava/lang/String;)V'
    'initializer-virtual|invalid initializer invocation form|main([Ljava/lang/String;)V'
)
for spec in "${negative_cases[@]}"; do
    IFS='|' read -r label diagnostic method <<< "$spec"
    folder="$scratch/cases/$label"
    if timeout 10s java -Xverify:all -cp "$folder" ConstructorFrameMain \
            > "$scratch/$label.jdk" 2>&1; then
        echo "FAIL JDK accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'java.lang.VerifyError' "$scratch/$label.jdk" || {
        echo "FAIL JDK rejected $label for a non-verification reason" >&2
        cat "$scratch/$label.jdk" >&2; exit 1;
    }
    grep -Fq "ConstructorFrameMain.$method" "$scratch/$label.jdk" || {
        echo "FAIL JDK failure in the wrong method: $label" >&2
        cat "$scratch/$label.jdk" >&2; exit 1;
    }
    if timeout 15s "$MAGNOLIA" build run \
            --entry "$folder/ConstructorFrameMain.class" \
            > "$scratch/$label.oak" 2>&1; then
        echo "FAIL Magnolia accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq "$diagnostic" "$scratch/$label.oak" || {
        echo "FAIL unexpected Magnolia diagnostic: $label" >&2
        cat "$scratch/$label.oak" >&2; exit 1;
    }
    grep -Fq "ConstructorFrameMain.$method" "$scratch/$label.oak" || {
        echo "FAIL Magnolia rejected in the wrong method: $label" >&2
        cat "$scratch/$label.oak" >&2; exit 1;
    }
    if grep -Exq '(7|1|true|false|0)' "$scratch/$label.oak"; then
        echo "FAIL $label produced output before selected-entry preflight" >&2
        cat "$scratch/$label.oak" >&2; exit 1
    fi
    printf '[constructor-verifier] PASS JDK VerifyError = pre-output refusal: %s\n' "$label"
done

# A reachable *later* method must not produce even a partial .oak/.js/.html/
# .java artifact; this checks preflight before Builder emits each route.
for suffix in oak js html java; do
    target="$scratch/refused.$suffix"
    if timeout 110s "$MAGNOLIA" build \
            --entry "$scratch/cases/uninitialized-return/ConstructorFrameMain.class" \
            --output "$target" > "$scratch/refused-$suffix.log" 2>&1; then
        echo "FAIL Builder emitted $suffix despite uninitialized return" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'uninitialized reference used before constructor invocation' \
        "$scratch/refused-$suffix.log" || {
        echo "FAIL Builder rejected for the wrong reason: $suffix" >&2
        cat "$scratch/refused-$suffix.log" >&2; exit 1;
    }
    if [[ -e "$target" ]]; then
        echo "FAIL Builder left a partial $suffix file" >&2; exit 1
    fi
    printf '[constructor-verifier] PASS Builder %s refuses before emission\n' "$suffix"
done
printf 'PASS constructor verifier: 4 JDK-accepted executions (JDK/native/Oak/JS for original), 8 JDK VerifyError preflight refusals, 4 no-emission targets\n'
