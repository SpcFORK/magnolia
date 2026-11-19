#!/usr/bin/env bash
# Independent Java 8 -Xverify:all oracle for the narrow putfield rule on
# uninitialized constructor `this`. This does NOT verify all constructor,
# access-control, exception-handler or StackMapTable rules.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" javac java python3 node timeout cmp diff; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing preinit-field oracle dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/variants"
javac --release 8 -g:none -d "$scratch/original" \
    test/fixtures/ConstructorFrameMain.java \
    test/fixtures/PreSuperFieldMain.java
for name in ConstructorFrameMain PreSuperFieldMain PreSuperFieldBase; do
    cmp "test/fixtures/$name.class" "$scratch/original/$name.class"
done
for name in ConstructorFrameMain PreSuperFieldMain; do
    timeout 10s java -Xverify:all -cp "$scratch/original" "$name" \
        > "$scratch/$name.jdk"
    diff -u "test/fixtures/$name.expected" "$scratch/$name.jdk"
    inputs=("$scratch/original/$name.class")
    if [[ "$name" == PreSuperFieldMain ]]; then
        inputs+=("$scratch/original/PreSuperFieldBase.class")
    fi
    timeout 15s "$MAGNOLIA" build run \
        --entry "${inputs[0]}" "${inputs[@]:1}" \
        --class "$name" > "$scratch/$name.oak"
    diff -u "$scratch/$name.jdk" "$scratch/$name.oak"
done

# Parse the constant pool and *named* Code methods, asserting the exact
# constructor/field refs before any same-length mutation. An inherited
# fieldref may symbolically name the subclass yet resolve in its superclass.
python3 - "$scratch/original" "$scratch/variants" <<'PY'
from pathlib import Path
import sys, shutil
source, dest = map(Path, sys.argv[1:])

def decode(raw):
    pos = 0
    def take(n):
        nonlocal pos
        data = raw[pos:pos+n]
        assert len(data) == n, 'truncated fixture'
        pos += n
        return data
    def u1(): return take(1)[0]
    def u2(): return int.from_bytes(take(2), 'big')
    def u4(): return int.from_bytes(take(4), 'big')
    assert take(4) == b'\xca\xfe\xba\xbe'
    assert take(4) == b'\0\0\0\x34', 'expected Java 8'
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
        else: raise ValueError(f'unexpected fixture pool tag {tag}')
        i += 1
    def utf(i):
        assert pool[i][0] == 1
        return pool[i][1]
    def clazz(i):
        assert pool[i][0] == 7
        return utf(pool[i][1])
    def member(tag, owner, name, desc):
        found = [i for i,e in enumerate(pool) if e and e[0] == tag and
                 clazz(e[1]) == owner and
                 (utf(pool[e[2]][1]),utf(pool[e[2]][2])) == (name,desc)]
        assert len(found) == 1, (tag,owner,name,desc,found)
        return found[0]
    def attributes():
        result = []
        for _ in range(u2()):
            name,size = utf(u2()),u4()
            at = pos
            take(size)
            result.append((name,at,size))
        return result
    take(6)
    take(2*u2())
    fields = []
    for _ in range(u2()):
        flags,name,desc = u2(),utf(u2()),utf(u2())
        fields.append((flags,name,desc))
        attributes()
    methods = {}
    for _ in range(u2()):
        take(2)
        name,desc = utf(u2()),utf(u2())
        for kind,at,size in attributes():
            if kind == 'Code':
                length = int.from_bytes(raw[at+4:at+8], 'big')
                assert 8+length <= size
                methods[(name,desc)] = (at+8,raw[at+8:at+8+length])
    return clazz,member,fields,methods

def save(label,owner,raw,start,old,new):
    assert len(old) == len(new), label
    assert raw[start:start+len(old)] == old, label
    changed = bytearray(raw)
    changed[start:start+len(old)] = new
    folder = dest/label
    folder.mkdir()
    (folder/(owner+'.class')).write_bytes(changed)
    if owner == 'PreSuperFieldMain':
        shutil.copyfile(source/'PreSuperFieldBase.class',
                        folder/'PreSuperFieldBase.class')
    print(f'[preinit-field] {label}: {owner} byte edit @ file offset {start}',
          file=sys.stderr)

be2 = lambda value: value.to_bytes(2,'big')
raw = (source/'ConstructorFrameMain.class').read_bytes()
clazz,member,fields,methods = decode(raw)
assert (0,'value','I') in fields and (8,'made','I') in fields
own = member(9,'ConstructorFrameMain','value','I')
static = member(9,'ConstructorFrameMain','made','I')
init = member(10,'ConstructorFrameMain','<init>','()V')
init_int = member(10,'ConstructorFrameMain','<init>','(I)V')
start,code = methods[('<init>','(I)V')]
original = b'\x2a\xb7'+be2(init)+b'\x2a\x1b\xb5'+be2(own)
assert code.startswith(original), code.hex()
save('own-before-this','ConstructorFrameMain',raw,start,original,
     b'\x2a\x1b\xb5'+be2(own)+b'\x2a\xb7'+be2(init))
save('get-before-this','ConstructorFrameMain',raw,start,original,
     b'\x2a\xb4'+be2(own)+b'\x57\x2a\xb7'+be2(init))
save('static-before-this','ConstructorFrameMain',raw,start,original,
     b'\x2a\x1b\xb5'+be2(static)+b'\x2a\xb7'+be2(init))
main_at,main = methods[('main','([Ljava/lang/String;)V')]
# Use the real new-class operand from main, not a guessed pool index.
assert main[0] == 0xbb
assert clazz(int.from_bytes(main[1:3],'big')) == 'ConstructorFrameMain'
new_call = main[:6] + b'\xb7'+be2(init_int)
assert main.startswith(new_call)
save('new-object-field-before-init','ConstructorFrameMain',raw,
     main_at+6,b'\xb7'+be2(init_int),b'\xb5'+be2(own))

raw = (source/'PreSuperFieldMain.class').read_bytes()
clazz,member,fields,methods = decode(raw)
assert (1,'own','I') in fields and not any(f[1] == 'inherited' for f in fields)
base_init = member(10,'PreSuperFieldBase','<init>','()V')
first = member(9,'PreSuperFieldMain','own','I')
inherited = member(9,'PreSuperFieldMain','inherited','I')
start,code = methods[('<init>','(I)V')]
original = b'\x2a\xb7'+be2(base_init)+b'\x2a\x1b\xb5'+be2(first)
assert code.startswith(original), code.hex()
assert code.count(b'\xb5'+be2(inherited)) == 1
save('own-before-super','PreSuperFieldMain',raw,start,original,
     b'\x2a\x1b\xb5'+be2(first)+b'\x2a\xb7'+be2(base_init))
save('inherited-before-super','PreSuperFieldMain',raw,start,original,
     b'\x2a\x1b\xb5'+be2(inherited)+b'\x2a\xb7'+be2(base_init))
PY

for variant in own-before-this own-before-super; do
    if [[ "$variant" == own-before-this ]]; then
        name=ConstructorFrameMain
    else
        name=PreSuperFieldMain
    fi
    folder="$scratch/variants/$variant"
    inputs=("$folder/$name.class")
    [[ "$name" == ConstructorFrameMain ]] || inputs+=("$folder/PreSuperFieldBase.class")
    timeout 10s java -Xverify:all -cp "$folder" "$name" > "$scratch/$variant.jdk"
    diff -u "test/fixtures/$name.expected" "$scratch/$variant.jdk"
    timeout 15s "$MAGNOLIA" build run --entry "${inputs[0]}" \
        "${inputs[@]:1}" --class "$name" > "$scratch/$variant.oak"
    diff -u "$scratch/$variant.jdk" "$scratch/$variant.oak"
    printf '[preinit-field] PASS JDK-accepted %s = Magnolia native\n' "$variant"
done
# The valid same-class case also survives Builder's standalone JS route.
timeout 105s "$MAGNOLIA" build \
    --entry "$scratch/variants/own-before-this/ConstructorFrameMain.class" \
    --output "$scratch/OwnBeforeThis.js" > "$scratch/js-build.log"
timeout 25s node "$scratch/OwnBeforeThis.js" > "$scratch/OwnBeforeThis.out"
diff -u test/fixtures/ConstructorFrameMain.expected "$scratch/OwnBeforeThis.out"
printf '[preinit-field] PASS JDK-accepted own-field store = bundled JS\n'

negative_cases=(
    'get-before-this|ConstructorFrameMain|uninitialized reference used before constructor invocation'
    'new-object-field-before-init|ConstructorFrameMain|uninitialized reference used before constructor invocation'
    'inherited-before-super|PreSuperFieldMain|uninitialized reference used before constructor invocation'
)
for spec in "${negative_cases[@]}"; do
    IFS='|' read -r variant name diagnostic <<< "$spec"
    folder="$scratch/variants/$variant"
    inputs=("$folder/$name.class")
    [[ "$name" == ConstructorFrameMain ]] || inputs+=("$folder/PreSuperFieldBase.class")
    if timeout 10s java -Xverify:all -cp "$folder" "$name" \
        > "$scratch/$variant.jdk" 2>&1; then
        echo "FAIL JDK accepted $variant" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'java.lang.VerifyError' "$scratch/$variant.jdk" || {
        echo "FAIL JDK non-verification error: $variant" >&2
        cat "$scratch/$variant.jdk" >&2; exit 1;
    }
    if timeout 15s "$MAGNOLIA" build run --entry "${inputs[0]}" \
        "${inputs[@]:1}" --class "$name" \
        > "$scratch/$variant.oak" 2>&1; then
        echo "FAIL Magnolia accepted $variant" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq "$diagnostic" "$scratch/$variant.oak" || {
        echo "FAIL unexpected Magnolia diagnostic: $variant" >&2
        cat "$scratch/$variant.oak" >&2; exit 1;
    }
    if grep -Exq '(before|7|8|1|true|false|0)' "$scratch/$variant.oak"; then
        echo "FAIL output before preflight: $variant" >&2
        cat "$scratch/$variant.oak" >&2; exit 1
    fi
    printf '[preinit-field] PASS JDK VerifyError = pre-output refusal: %s\n' "$variant"
done
# A putfield referencing a *static* field of this class is NOT a JDK
# VerifyError: the JDK accepts verification and raises a dynamic
# IncompatibleClassChangeError. Magnolia's bounded linker refuses the
# unsupported opcode/field-kind combination at preflight instead. Do not
# misreport this as agreement between verifiers.
folder="$scratch/variants/static-before-this"
if timeout 10s java -Xverify:all -cp "$folder" ConstructorFrameMain \
        > "$scratch/static.jdk" 2>&1; then
    echo 'FAIL JDK did not dynamically reject static-before-this' >&2; exit 1
else
    rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
fi
grep -Fq 'java.lang.IncompatibleClassChangeError' "$scratch/static.jdk"
! grep -Fq 'VerifyError' "$scratch/static.jdk"
if timeout 15s "$MAGNOLIA" build run \
        --entry "$folder/ConstructorFrameMain.class" \
        > "$scratch/static.oak" 2>&1; then
    echo 'FAIL Magnolia accepted unsupported static putfield' >&2; exit 1
else
    rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
fi
grep -Fq 'uninitialized reference used before constructor invocation' \
    "$scratch/static.oak"
! grep -Exq '(7|8|1|true|false|0)' "$scratch/static.oak"
printf '[preinit-field] PASS JDK-valid static putfield is a DYNAMIC error; Magnolia refuses the unsupported form at preflight\n'
# One inherited-field rejection is in a constructor invoked *after* a print;
# Builder must refuse before emitting a JS artifact (not a late runtime trap).
if timeout 105s "$MAGNOLIA" build \
    --entry "$scratch/variants/inherited-before-super/PreSuperFieldMain.class" \
    --class-files "$scratch/variants/inherited-before-super/PreSuperFieldBase.class" \
    --output "$scratch/refused.js" > "$scratch/refused.log" 2>&1; then
    echo 'FAIL Builder accepted inherited-before-super' >&2; exit 1
else
    rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
fi
grep -Fq 'uninitialized reference used before constructor invocation' \
    "$scratch/refused.log" || {
    echo 'FAIL Builder rejected for the wrong reason' >&2
    cat "$scratch/refused.log" >&2; exit 1;
}
[[ ! -e "$scratch/refused.js" ]] || {
    echo 'FAIL Builder left a partial JS file' >&2; exit 1;
}
printf 'PASS preinit field oracle: two JDK-accepted pre-super stores, three JDK VerifyError/preflight refusals, one documented dynamic-error distinction, no late output or JS emission\n'
