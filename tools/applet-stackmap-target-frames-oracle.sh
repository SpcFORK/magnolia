#!/usr/bin/env bash
# Selected-entry control-flow target frames, distinct from StackMapTable
# structure and from full declared/inferred type verification. Version 51+
# uses the required-frame gate; JDK 11's version-50 inference fallback (and
# legacy version 49) must remain accepted. Class-file edits are located by
# named Code methods and correct all enclosing attribute lengths/counts.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" java javac python3 timeout cmp diff; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing target-frame oracle dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/cases"
javac --release 8 -g:none -d "$scratch/original" \
    test/fixtures/StackMapFrameMain.java \
    test/fixtures/StackMapExtendedMain.java \
    test/fixtures/StackMapEntryLoopMain.java \
    test/fixtures/SwitchMain.java
for name in StackMapFrameMain FrameCtor StackMapExtendedMain \
            StackMapEntryLoopMain SwitchMain; do
    cmp "test/fixtures/$name.class" "$scratch/original/$name.class"
done
for name in StackMapFrameMain StackMapExtendedMain \
            StackMapEntryLoopMain SwitchMain; do
    timeout 12s java -Xverify:all -cp "$scratch/original" "$name" \
        > "$scratch/$name.jdk"
    inputs=("$scratch/original/$name.class")
    [[ "$name" != StackMapFrameMain ]] || \
        inputs+=("$scratch/original/FrameCtor.class")
    timeout 15s "$MAGNOLIA" build run \
        --entry "${inputs[0]}" "${inputs[@]:1}" --class "$name" \
        > "$scratch/$name.oak"
    diff -u "$scratch/$name.jdk" "$scratch/$name.oak"
    [[ "$name" == SwitchMain ]] || \
        diff -u "test/fixtures/$name.expected" "$scratch/$name.jdk"
    printf '[target-frames] PASS original JDK/Oak: %s\n' "$name"
done

python3 - "$scratch/original" "$scratch/cases" <<'PY'
from pathlib import Path
import shutil, sys
source, cases = map(Path, sys.argv[1:])
def be2(n): return n.to_bytes(2, 'big')
def be4(n): return n.to_bytes(4, 'big')

def decode(name):
    raw = (source/(name+'.class')).read_bytes()
    pos = 0
    def take(n):
        nonlocal pos
        chunk = raw[pos:pos+n]
        assert len(chunk) == n, (name, pos, n)
        pos += n
        return chunk
    def u1(): return take(1)[0]
    def u2(): return int.from_bytes(take(2), 'big')
    def u4(): return int.from_bytes(take(4), 'big')
    assert take(4) == b'\xca\xfe\xba\xbe'
    assert take(4) == b'\x00\x00\x00\x34'
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
        else: raise AssertionError((name, 'pool tag', tag))
        i += 1
    def utf(index):
        assert pool[index] and pool[index][0] == 1
        return pool[index][1]
    def attributes():
        at_count = pos
        result = []
        for _ in range(u2()):
            start = pos
            kind = utf(u2())
            size_at = pos
            length = u4()
            payload_at = pos
            take(length)
            result.append(dict(name=kind, start=start, size_at=size_at,
                               length=length, payload_at=payload_at, end=pos))
        return at_count, result
    take(6)
    take(2*u2())
    for _ in range(u2()): take(6); attributes()
    methods = {}
    for _ in range(u2()):
        take(2)
        method_name, desc = utf(u2()), utf(u2())
        _, outer = attributes()
        for a in outer:
            if a['name'] != 'Code': continue
            prior = pos
            pos = a['payload_at']
            max_stack, max_locals, length = u2(), u2(), u4()
            code_at = pos
            code = take(length)
            handlers = [(u2(),u2(),u2(),u2()) for _ in range(u2())]
            count_at, nested = attributes()
            assert pos == a['end'], (name, method_name)
            pos = prior
            methods[(method_name,desc)] = dict(outer=a, code_at=code_at,
                code=code, max_stack=max_stack, max_locals=max_locals,
                handlers=handlers, count_at=count_at, nested=nested)
    attributes()
    assert pos == len(raw), (name, pos, len(raw))
    return raw, methods

def save(label, name, raw, version=52):
    buf = bytearray(raw)
    assert buf[6:8] == b'\x00\x34'
    buf[6:8] = be2(version)
    folder = cases/label
    folder.mkdir()
    (folder/(name+'.class')).write_bytes(buf)
    if name == 'StackMapFrameMain':
        shutil.copyfile(source/'FrameCtor.class', folder/'FrameCtor.class')

def without_table(raw, method):
    attrs = [a for a in method['nested'] if a['name'] == 'StackMapTable']
    assert len(attrs) == 1
    a = attrs[0]
    assert int.from_bytes(raw[method['count_at']:method['count_at']+2],'big') == 1
    original_len = a['end'] - a['start']
    buf = bytearray(raw)
    outer = method['outer']
    buf[outer['size_at']:outer['size_at']+4] = be4(outer['length']-original_len)
    buf[method['count_at']:method['count_at']+2] = be2(0)
    del buf[a['start']:a['end']]
    return buf, a

name = 'StackMapFrameMain'
raw, methods = decode(name)
make = methods[('make','(I)I')]
assert make['code'] == bytes.fromhex('bb 00 02 59 1a 99 00 08 10 07 '
                                      'a7 00 05 10 09 b7 00 03 4c 2b b4 00 04 ac')
empty, a = without_table(raw, make)
assert raw[a['payload_at']:a['end']] == bytes.fromhex(
    '00 02 ff 00 0d 00 01 01 00 02 08 00 00 08 00 00 '
    'ff 00 01 00 01 01 00 03 08 00 00 08 00 00 01')
for version in (49, 50, 51, 52):
    save(f'branch-empty-v{version}', name, empty, version)
# Keep the first well-formed full_frame at PC 13. The missing second frame
# at PC 15 must fail; this is not an attribute-length corruption.
p = a['payload_at']
removed = raw[p+16:a['end']]
assert len(removed) == 15 and raw[p+16] == 255
second = bytearray(raw)
outer = make['outer']
second[outer['size_at']:outer['size_at']+4] = be4(outer['length']-len(removed))
second[a['size_at']:a['size_at']+4] = be4(a['length']-len(removed))
second[p:p+2] = be2(1)
del second[p+16:a['end']]
save('branch-second-v52', name, second)
# A valid version-51 class with its two original frames stays accepted.
save('branch-intact-v51', name, raw, 51)

name = 'StackMapExtendedMain'
raw, methods = decode(name)
read = methods[('read','(I)I')]
assert len(read['code']) == 91 and read['handlers'] == [(0,87,88,2)]
empty, a = without_table(raw, read)
assert raw[a['payload_at']:a['end']] == b'\x00\x01\xf7\x00\x58\x07\x00\x02'
for version in (49, 50, 51, 52):
    save(f'handler-empty-v{version}', name, empty, version)

name = 'SwitchMain'
raw, methods = decode(name)
dense = methods[('dense','(I)I')]
assert 170 in dense['code'] and len(dense['code']) == 53
empty, a = without_table(raw, dense)
assert raw[a['payload_at']:a['end']] == b'\x00\x06\x24\x02\x02\x01\x02\x02'
for version in (49, 50, 51, 52):
    save(f'switch-empty-v{version}', name, empty, version)

name = 'StackMapEntryLoopMain'
raw, methods = decode(name)
spin = methods[('spin','()V')]
assert spin['code'] == b'\xa7\x00\x00' and spin['handlers'] == []
empty, a = without_table(raw, spin)
assert raw[a['payload_at']:a['end']] == b'\x00\x01\x00'
for version in (49, 50, 51, 52):
    save(f'entry-empty-v{version}', name, empty, version)
print('[target-frames] PASS named-method / corrected-length mutations')
PY

# JDK errors concern *coverage*, not malformed table lengths. Entry-loop
# verifies even the uncalled spin() on JDK class load; Magnolia only verifies
# selected methods, so its negative is run with spin selected for preflight.
negative_cases=(
    'branch-empty-v52|StackMapFrameMain|13|make(I)I'
    'branch-empty-v51|StackMapFrameMain|13|make(I)I'
    'branch-second-v52|StackMapFrameMain|15|make(I)I'
    'handler-empty-v52|StackMapExtendedMain|88|read(I)I'
    'handler-empty-v51|StackMapExtendedMain|88|read(I)I'
    'switch-empty-v52|SwitchMain|50|dense(I)I'
    'switch-empty-v51|SwitchMain|50|dense(I)I'
    'entry-empty-v52|StackMapEntryLoopMain|0|spin()V'
    'entry-empty-v51|StackMapEntryLoopMain|0|spin()V'
)
for spec in "${negative_cases[@]}"; do
    IFS='|' read -r label name target method <<< "$spec"
    folder="$scratch/cases/$label"
    if timeout 12s java -Xverify:all -cp "$folder" "$name" \
        > "$scratch/$label.jdk.log" 2>&1; then
        echo "FAIL JDK accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq "java.lang.VerifyError: Expecting a stackmap frame at branch target $target" \
        "$scratch/$label.jdk.log" || {
        echo "FAIL wrong JDK error for $label" >&2
        head -n 20 "$scratch/$label.jdk.log" >&2; exit 1;
    }
    inputs=("$folder/$name.class")
    [[ "$name" != StackMapFrameMain ]] || inputs+=("$folder/FrameCtor.class")
    options=(--class "$name")
    [[ "$name" != StackMapEntryLoopMain ]] || options+=(--method spin --desc '()V')
    if timeout 8s "$MAGNOLIA" build run \
        --entry "${inputs[0]}" "${inputs[@]:1}" "${options[@]}" \
        > "$scratch/$label.oak.log" 2>&1; then
        echo "FAIL Magnolia accepted $label" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || {
            echo "FAIL Magnolia hung on $label" >&2; exit 1;
        }
    fi
    grep -Fq "$name.$method at byte $target: missing StackMapTable frame at reachable control-flow target" \
        "$scratch/$label.oak.log" || {
        echo "FAIL wrong Magnolia diagnostic for $label" >&2
        cat "$scratch/$label.oak.log" >&2; exit 1;
    }
    if grep -Exq '(101|7|9|409|11|31)' "$scratch/$label.oak.log"; then
        echo "FAIL output before selected-entry preflight: $label" >&2
        cat "$scratch/$label.oak.log" >&2; exit 1
    fi
    printf '[target-frames] PASS JDK VerifyError target %s = pre-output refusal: %s\n' \
        "$target" "$label"
done

# Under this JDK, 50.0 may fall back to inference; 49.0 also ignores the
# absent table. No blanket 'table required' rule may reject those positives.
for version in 49 50; do
    for name in StackMapFrameMain StackMapExtendedMain SwitchMain \
                StackMapEntryLoopMain; do
        case "$name" in
            StackMapFrameMain) prefix=branch;;
            StackMapExtendedMain) prefix=handler;;
            SwitchMain) prefix=switch;;
            *) prefix=entry;;
        esac
        folder="$scratch/cases/$prefix-empty-v$version"
        timeout 12s java -Xverify:all -cp "$folder" "$name" \
            > "$scratch/$prefix-v$version.jdk"
        inputs=("$folder/$name.class")
        [[ "$name" != StackMapFrameMain ]] || \
            inputs+=("$folder/FrameCtor.class")
        timeout 15s "$MAGNOLIA" build run \
            --entry "${inputs[0]}" "${inputs[@]:1}" --class "$name" \
            > "$scratch/$prefix-v$version.oak"
        diff -u "$scratch/$prefix-v$version.jdk" \
            "$scratch/$prefix-v$version.oak"
    done
    printf '[target-frames] PASS all four version %s inference positives\n' "$version"
done
folder="$scratch/cases/branch-intact-v51"
timeout 12s java -Xverify:all -cp "$folder" StackMapFrameMain \
    > "$scratch/v51-intact.jdk"
timeout 15s "$MAGNOLIA" build run \
    --entry "$folder/StackMapFrameMain.class" "$folder/FrameCtor.class" \
    --class StackMapFrameMain > "$scratch/v51-intact.oak"
diff -u "$scratch/v51-intact.jdk" "$scratch/v51-intact.oak"
diff -u test/fixtures/StackMapFrameMain.expected "$scratch/v51-intact.jdk"
printf '[target-frames] PASS complete version-51 declared frames\n'

# This is a SELECTED-entry gate, not all-method Java class verification:
# JDK refuses entry-empty-v52 even when main would not call spin, whereas
# Magnolia can run that selected main. Pin the limit rather than hide it.
folder="$scratch/cases/entry-empty-v52"
timeout 15s "$MAGNOLIA" build run \
    --entry "$folder/StackMapEntryLoopMain.class" \
    --class StackMapEntryLoopMain > "$scratch/unselected-limit.oak"
diff -u test/fixtures/StackMapEntryLoopMain.expected \
    "$scratch/unselected-limit.oak"
printf '[target-frames] LIMIT JDK verifies unselected spin; Magnolia main selection does not\n'

# A later make() error must be caught before main's earlier 101 print and
# before the Builder writes even a partial artifact for any target.
for suffix in oak js html java; do
    target="$scratch/refused.$suffix"
    if timeout 115s "$MAGNOLIA" build \
        --entry "$scratch/cases/branch-second-v52/StackMapFrameMain.class" \
        --class-files "$scratch/cases/branch-second-v52/FrameCtor.class" \
        --output "$target" > "$scratch/refused-$suffix.log" 2>&1; then
        echo "FAIL Builder emitted $suffix for missing frame" >&2; exit 1
    else
        rc=$?; [[ "$rc" != 124 && "$rc" != 137 ]] || exit 1
    fi
    grep -Fq 'missing StackMapTable frame at reachable control-flow target' \
        "$scratch/refused-$suffix.log" || {
        echo "FAIL Builder rejected for the wrong reason ($suffix)" >&2
        cat "$scratch/refused-$suffix.log" >&2; exit 1;
    }
    [[ ! -e "$target" ]] || {
        echo "FAIL Builder left a partial $suffix file" >&2; exit 1;
    }
    printf '[target-frames] PASS Builder %s refuses before emission\n' "$suffix"
done
printf 'PASS required-target StackMapTable oracle: 9 JDK VerifyError/pre-output agreements, 9 JDK-valid versioned positives, 4 Builder no-emission checks, 1 explicit unselected-method non-parity\n'
