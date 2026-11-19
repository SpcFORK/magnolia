#!/usr/bin/env bash
# Version-aware, bounded StackMapTable *structural* gate. The selected-entry
# verifier additionally checks decoded instruction boundaries, but does not
# compare declared stack-map types with inferred frames or require frames at
# branch targets. Independently compare every named mutant with JDK 11
# -Xverify:all; never count a JDK-valid dynamic failure as verifier parity.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" javac java python3 node timeout cmp diff; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing StackMapTable oracle dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/cases"

javac --release 8 -g:none -d "$scratch/original" \
    test/fixtures/StackMapFrameMain.java \
    test/fixtures/StackMapExtendedMain.java \
    test/fixtures/ConstructorFrameMain.java
for name in StackMapFrameMain FrameCtor StackMapExtendedMain ConstructorFrameMain; do
    cmp "test/fixtures/$name.class" "$scratch/original/$name.class"
done
for name in StackMapFrameMain StackMapExtendedMain ConstructorFrameMain; do
    timeout 12s java -Xverify:all -cp "$scratch/original" "$name" \
        > "$scratch/$name.jdk"
    diff -u "test/fixtures/$name.expected" "$scratch/$name.jdk"
    inputs=("$scratch/original/$name.class")
    [[ "$name" != StackMapFrameMain ]] || \
        inputs+=("$scratch/original/FrameCtor.class")
    timeout 20s "$MAGNOLIA" build run \
        --entry "${inputs[0]}" "${inputs[@]:1}" --class "$name" \
        > "$scratch/$name.oak"
    diff -u "$scratch/$name.jdk" "$scratch/$name.oak"
    printf '[stackmap] PASS pinned Java 8 / JDK / Oak: %s\n' "$name"
done
# Positive standalone Oak and bundled JS exercise the shared Builder
# selected-entry path (the latter must not depend on a browser/network).
timeout 75s "$MAGNOLIA" build \
    --entry "$scratch/original/StackMapFrameMain.class" \
    --class-files "$scratch/original/FrameCtor.class" \
    --output "$scratch/StackMapFrameMain.oak" > "$scratch/positive-oak-build.log"
timeout 25s "$MAGNOLIA" "$scratch/StackMapFrameMain.oak" \
    > "$scratch/standalone.oak.out"
diff -u "$scratch/StackMapFrameMain.jdk" "$scratch/standalone.oak.out"
timeout 115s "$MAGNOLIA" build \
    --entry "$scratch/original/StackMapFrameMain.class" \
    --class-files "$scratch/original/FrameCtor.class" \
    --output "$scratch/StackMapFrameMain.js" > "$scratch/positive-js-build.log"
timeout 25s node "$scratch/StackMapFrameMain.js" > "$scratch/bundled.js.out"
diff -u "$scratch/StackMapFrameMain.jdk" "$scratch/bundled.js.out"
printf '[stackmap] PASS original standalone Oak / bundled JS = JDK\n'

# Parse named methods, their *nested* Code attributes, and the class version.
# Assert real frame encodings and operand offsets before each same-size edit.
# Only duplicate/removal deliberately relocate attribute bytes; both fix
# Code.attribute_length and Code.attributes_count without changing bytecode.
python3 - "$scratch/original" "$scratch/cases" <<'PY'
from pathlib import Path
import shutil, sys
source, dest = map(Path, sys.argv[1:])
def be2(n): return n.to_bytes(2, 'big')
def be4(n): return n.to_bytes(4, 'big')

def decode(filename):
    raw = filename.read_bytes()
    pos = 0
    def take(n):
        nonlocal pos
        data = raw[pos:pos+n]
        assert len(data) == n, f'truncated fixture: {filename}'
        pos += n
        return data
    def u1(): return take(1)[0]
    def u2(): return int.from_bytes(take(2), 'big')
    def u4(): return int.from_bytes(take(4), 'big')
    assert take(4) == b'\xca\xfe\xba\xbe'
    assert take(4) == b'\x00\x00\x00\x34', 'expected Java 8 class'
    pool = [None] * u2()
    i = 1
    while i < len(pool):
        tag = u1()
        if tag == 1: pool[i] = (tag, take(u2()).decode('utf-8'))
        elif tag in (3, 4): take(4)
        elif tag in (5, 6): take(8); i += 1
        elif tag in (7, 8, 16, 19, 20): pool[i] = (tag, u2())
        elif tag in (9, 10, 11, 12, 17, 18): pool[i] = (tag, u2(), u2())
        elif tag == 15: take(3)
        else: raise AssertionError((filename, 'pool tag', tag))
        i += 1
    def utf(index):
        assert pool[index] is not None and pool[index][0] == 1
        return pool[index][1]
    def attrs():
        count_at = pos
        result = []
        for _ in range(u2()):
            at = pos
            name = utf(u2())
            size_at = pos
            length = u4()
            payload_at = pos
            take(length)
            result.append(dict(name=name, at=at, size_at=size_at,
                               length=length, payload_at=payload_at, end=pos))
        return count_at, result
    take(6)
    take(2*u2())
    for _ in range(u2()): take(6); attrs()
    methods = {}
    for _ in range(u2()):
        take(2)
        name, desc = utf(u2()), utf(u2())
        _, outer = attrs()
        for attr in outer:
            if attr['name'] != 'Code': continue
            previous = pos
            pos = attr['payload_at']
            maximum_stack, maximum_locals, code_len = u2(), u2(), u4()
            code_at = pos
            code = take(code_len)
            take(8*u2())
            attr_count_at, nested = attrs()
            assert pos == attr['end'], (filename, name)
            pos = previous
            methods[(name,desc)] = dict(outer=attr, code=code,
                code_at=code_at, max_stack=maximum_stack,
                max_locals=maximum_locals, count_at=attr_count_at,
                nested=nested)
    attrs()
    assert pos == len(raw)
    return raw, pool, methods

def stackmap(method):
    entries = [x for x in method['nested'] if x['name'] == 'StackMapTable']
    assert len(entries) == 1, entries
    return entries[0]

def save(label, owner, data):
    folder = dest/label
    folder.mkdir()
    (folder/(owner+'.class')).write_bytes(data)
    if owner == 'StackMapFrameMain':
        shutil.copyfile(source/'FrameCtor.class', folder/'FrameCtor.class')

def patch(label, owner, raw, *edits, version=None):
    buf = bytearray(raw)
    for offset, before, after in edits:
        assert len(before) == len(after), label
        assert buf[offset:offset+len(before)] == before, (label, offset)
        buf[offset:offset+len(before)] = after
    if version is not None:
        assert buf[6:8] == b'\x00\x34'
        buf[6:8] = be2(version)
    save(label, owner, buf)

owner = 'StackMapFrameMain'
raw, pool, methods = decode(source/(owner+'.class'))
make = methods[('make','(I)I')]
attr = stackmap(make)
p = attr['payload_at']
expected = bytes.fromhex('00 02 ff 00 0d 00 01 01 00 02 08 00 00 08 00 00 '
                         'ff 00 01 00 01 01 00 03 08 00 00 08 00 00 01')
assert raw[p:attr['end']] == expected
assert make['code'].startswith(b'\xbb\x00\x02\x59\x1a\x99')
assert make['code'][8:10] == b'\x10\x07'
assert len(make['code']) == 24 and make['max_locals'] == 2
patch('reserved', owner, raw, (p+2, b'\xff', b'\xc8'))
patch('invalid-type', owner, raw, (p+10, b'\x08', b'\x09'))
patch('invalid-new', owner, raw, (p+11, b'\x00\x00', b'\x00\x01'))
patch('invalid-new-outside', owner, raw,
      (p+11, b'\x00\x00', b'\x00\xff'))
patch('frame-outside-code', owner, raw,
      (p+3, b'\x00\x0d', b'\x00\x18'))
patch('frame-in-operand', owner, raw,
      (p+3, b'\x00\x0d', b'\x00\x0c'))
patch('too-many-locals', owner, raw,
      (p+5, b'\x00\x01', b'\x00\x03'))
patch('fewer-entries', owner, raw,
      (p, b'\x00\x02', b'\x00\x01'))
patch('extra-entries', owner, raw,
      (p, b'\x00\x02', b'\x00\x03'))
# The raw operand byte equals 0xbb but is not an instruction boundary.
patch('operand-fake-new', owner, raw,
      (make['code_at']+9, b'\x07', b'\xbb'),
      (p+11, b'\x00\x00', b'\x00\x09'))

# Duplicate/removal preserves a structurally valid enclosing Code length.
outer = make['outer']
assert int.from_bytes(raw[make['count_at']:make['count_at']+2], 'big') == 1
encoded = raw[attr['at']:attr['end']]
assert len(encoded) == attr['length']+6
buf = bytearray(raw)
buf[outer['size_at']:outer['size_at']+4] = be4(outer['length']+len(encoded))
buf[make['count_at']:make['count_at']+2] = be2(2)
buf[attr['end']:attr['end']] = encoded
save('duplicate', owner, buf)
buf = bytearray(raw)
buf[outer['size_at']:outer['size_at']+4] = be4(outer['length']-len(encoded))
buf[make['count_at']:make['count_at']+2] = be2(0)
del buf[attr['at']:attr['end']]
save('missing-frame', owner, buf)
patch('wrong-declared-type', owner, raw,
      (p+7, b'\x01', b'\x02'))

owner = 'StackMapExtendedMain'
raw, pool, methods = decode(source/(owner+'.class'))
read = methods[('read','(I)I')]
attr = stackmap(read)
p = attr['payload_at']
assert raw[p:attr['end']] == b'\x00\x01\xf7\x00\x58\x07\x00\x02'
assert len(read['code']) == 91
patch('extended-outside-code', owner, raw,
      (p+3, b'\x00\x58', b'\x00\x5b'))

owner = 'ConstructorFrameMain'
raw, pool, methods = decode(source/(owner+'.class'))
main = methods[('main','([Ljava/lang/String;)V')]
attr = stackmap(main)
p = attr['payload_at']
assert raw[p:p+10] == b'\x00\x04\xff\x00\x33\x00\x03\x07\x00\x1a'
assert pool[0x1a][0] == 7 and pool[1][0] == 10
patch('reserved-v52', owner, raw, (p+2, b'\xff', b'\xc8'))
patch('reserved-v50', owner, raw,
      (p+2, b'\xff', b'\xc8'), version=50)
patch('reserved-v49', owner, raw,
      (p+2, b'\xff', b'\xc8'), version=49)
patch('bad-class-index', owner, raw,
      (p+8, b'\x00\x1a', b'\x00\x00'))
patch('bad-class-tag', owner, raw,
      (p+8, b'\x00\x1a', b'\x00\x01'))
print('[stackmap] PASS named-method structural mutation assertions')
PY

# JDK 11 can classify an invalid frame as either ClassFormatError or
# VerifyError; pin the actual distinction for each tested mutation.
negative_cases=(
    'reserved|StackMapFrameMain|ClassFormatError|reserved StackMapTable frame type'
    'invalid-type|StackMapFrameMain|ClassFormatError|invalid StackMapTable verification type'
    'invalid-new|StackMapFrameMain|ClassFormatError|invalid StackMapTable uninitialized new offset'
    'invalid-new-outside|StackMapFrameMain|ClassFormatError|invalid StackMapTable uninitialized new offset'
    'frame-outside-code|StackMapFrameMain|VerifyError|StackMapTable frame beyond Code'
    'frame-in-operand|StackMapFrameMain|VerifyError|StackMapTable frame is not on an instruction boundary'
    'too-many-locals|StackMapFrameMain|ClassFormatError|too many StackMapTable locals'
    'fewer-entries|StackMapFrameMain|ClassFormatError|invalid StackMapTable length'
    'extra-entries|StackMapFrameMain|ClassFormatError|truncated class file'
    'operand-fake-new|StackMapFrameMain|ClassFormatError|StackMapTable uninitialized offset is not a new instruction'
    'duplicate|StackMapFrameMain|ClassFormatError|duplicate StackMapTable attribute'
    'extended-outside-code|StackMapExtendedMain|VerifyError|StackMapTable frame beyond Code'
    'reserved-v52|ConstructorFrameMain|ClassFormatError|reserved StackMapTable frame type'
    'reserved-v50|ConstructorFrameMain|ClassFormatError|reserved StackMapTable frame type'
    'bad-class-index|ConstructorFrameMain|ClassFormatError|invalid StackMapTable Object constant-pool index'
    'bad-class-tag|ConstructorFrameMain|ClassFormatError|invalid StackMapTable Object constant-pool index'
    'missing-frame|StackMapFrameMain|VerifyError|missing StackMapTable frame at reachable control-flow target'
)
for spec in "${negative_cases[@]}"; do
    IFS='|' read -r label name jdkError oakDiagnostic <<< "$spec"
    folder="$scratch/cases/$label"
    inputs=("$folder/$name.class")
    [[ "$name" != StackMapFrameMain ]] || inputs+=("$folder/FrameCtor.class")
    if timeout 12s java -Xverify:all -cp "$folder" "$name" \
        > "$scratch/$label.jdk.log" 2>&1; then
        echo "FAIL JDK accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq "java.lang.$jdkError" "$scratch/$label.jdk.log" || {
        echo "FAIL JDK did not raise $jdkError for $label" >&2
        cat "$scratch/$label.jdk.log" >&2; exit 1;
    }
    if grep -Exq '(101|7|9|409)' "$scratch/$label.jdk.log"; then
        echo "FAIL JDK produced program output for $label" >&2; exit 1
    fi
    if timeout 20s "$MAGNOLIA" build run \
        --entry "${inputs[0]}" "${inputs[@]:1}" --class "$name" \
        > "$scratch/$label.oak.log" 2>&1; then
        echo "FAIL Magnolia accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq "$oakDiagnostic" "$scratch/$label.oak.log" || {
        echo "FAIL Magnolia rejected $label for the wrong reason" >&2
        cat "$scratch/$label.oak.log" >&2; exit 1;
    }
    if grep -Exq '(101|7|9|409)' "$scratch/$label.oak.log"; then
        echo "FAIL program output before selected-entry preflight: $label" >&2
        cat "$scratch/$label.oak.log" >&2; exit 1
    fi
    printf '[stackmap] PASS JDK %s / pre-output refusal: %s\n' "$jdkError" "$label"
done

# A version-49 class file uses legacy inference and can ignore a malformed
# StackMapTable payload. A version-50/52 class above must not ignore it.
folder="$scratch/cases/reserved-v49"
timeout 12s java -Xverify:all -cp "$folder" ConstructorFrameMain \
    > "$scratch/v49.jdk"
diff -u test/fixtures/ConstructorFrameMain.expected "$scratch/v49.jdk"
timeout 20s "$MAGNOLIA" build run \
    --entry "$folder/ConstructorFrameMain.class" \
    > "$scratch/v49.oak"
diff -u "$scratch/v49.jdk" "$scratch/v49.oak"
printf '[stackmap] PASS JDK 49.0 inference ignores a reserved frame type\n'

# Explicit, pinned NON-parity: declared/inferred frame type comparison remains
# open even after the later required-control-flow-target frame slice.
for label in wrong-declared-type; do
    folder="$scratch/cases/$label"
    if timeout 12s java -Xverify:all -cp "$folder" StackMapFrameMain \
        > "$scratch/$label.jdk.limit" 2>&1; then
        echo "FAIL expected JDK VerifyError for known limitation $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'java.lang.VerifyError' "$scratch/$label.jdk.limit"
    timeout 20s "$MAGNOLIA" build run \
        --entry "$folder/StackMapFrameMain.class" "$folder/FrameCtor.class" \
        --class StackMapFrameMain > "$scratch/$label.oak.limit"
    diff -u test/fixtures/StackMapFrameMain.expected \
        "$scratch/$label.oak.limit"
    printf '[stackmap] LIMIT JDK VerifyError but bounded Magnolia accepts: %s\n' "$label"
done

# An invalid reachable *later* method must not produce even an incomplete
# Builder artifact, including the JS path whose generation is slower.
for suffix in oak js html java; do
    target="$scratch/refused.$suffix"
    if timeout 115s "$MAGNOLIA" build \
        --entry "$scratch/cases/frame-in-operand/StackMapFrameMain.class" \
        --class-files "$scratch/cases/frame-in-operand/FrameCtor.class" \
        --output "$target" > "$scratch/refused-$suffix.log" 2>&1; then
        echo "FAIL Builder accepted frame-in-operand ($suffix)" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'StackMapTable frame is not on an instruction boundary' \
        "$scratch/refused-$suffix.log" || {
        echo "FAIL Builder rejected for the wrong reason: $suffix" >&2
        cat "$scratch/refused-$suffix.log" >&2; exit 1;
    }
    [[ ! -e "$target" ]] || {
        echo "FAIL Builder left a partial $suffix file" >&2; exit 1;
    }
    printf '[stackmap] PASS Builder %s refuses before emission\n' "$suffix"
done
printf 'PASS StackMapTable structural oracle: 3 pinned JDK/Oak positives, standalone Oak/JS, 17 JDK rejections at preflight (including the later frame-coverage gate), 4 Builder no-emission targets, legacy v49 acceptance, 1 explicit declared-type non-parity probe\n'
