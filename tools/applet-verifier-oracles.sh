#!/usr/bin/env bash
# Reproducible negative JVM-oracle matrix for height and I/F/A/J/D frames.
# Mutate checked-in javac --release 8 fixtures at Code payload offsets found
# by a tiny class-file reader; compare Magnolia rejection with -Xverify:all.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" java python3 timeout; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing verifier-oracle dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
python3 - test/fixtures "$scratch" <<'PY'
from pathlib import Path
import sys
source, out = map(Path, sys.argv[1:])

# Only the metadata needed to locate named method Code payloads is decoded;
# malformed class parsing itself is exercised separately by Magnolia tests.
def methods(raw):
    pos = 0
    def take(n):
        nonlocal pos
        result = raw[pos:pos+n]
        assert len(result) == n, 'truncated fixture'
        pos += n
        return result
    def u1(): return take(1)[0]
    def u2(): return int.from_bytes(take(2), 'big')
    def u4(): return int.from_bytes(take(4), 'big')
    assert take(4) == b'\xca\xfe\xba\xbe'
    take(4)  # minor/major version
    count = u2()
    pool = [None] * count
    i = 1
    while i < count:
        tag = u1()
        if tag == 1: pool[i] = take(u2()).decode('utf-8')
        elif tag in (3, 4): take(4)
        elif tag in (5, 6): take(8); i += 1
        elif tag in (7, 8, 16, 19, 20): take(2)
        elif tag in (9, 10, 11, 12, 17, 18): take(4)
        elif tag == 15: take(3)
        else: raise ValueError(f'unexpected fixture constant-pool tag {tag}')
        i += 1
    take(6)  # access flags, this_class, super_class
    take(2 * u2())  # interfaces
    def attributes():
        found = []
        for _ in range(u2()):
            name = pool[u2()]
            size = u4()
            offset = pos
            take(size)
            found.append((name, offset, size))
        return found
    for _ in range(u2()):
        take(6)  # field access/name/descriptor
        attributes()
    result = {}
    for _ in range(u2()):
        take(2)  # access flags
        name, desc = pool[u2()], pool[u2()]
        for attr, offset, size in attributes():
            if attr == 'Code':
                length = int.from_bytes(raw[offset+4:offset+8], 'big')
                assert 8 + length <= size
                result[name + desc] = (offset, offset + 8, raw[offset+8:offset+8+length])
    return result

def patch(name, raw, offset, before, after):
    assert raw[offset:offset+len(before)] == before, f'fixture drift: {name}'
    folder = out / name
    folder.mkdir()
    data = bytearray(raw)
    data[offset:offset+len(before)] = after
    (folder / class_name).write_bytes(data)

class_name = 'StackHeightMain.class'
raw = (source / class_name).read_bytes()
_, code, body = methods(raw)['main([Ljava/lang/String;)V']
patch('height-underflow', raw, code, body[0:1], b'\x57')
patch('height-overflow', raw, code - 8, raw[code-8:code-6], b'\0\0')
patch('height-merge', raw, code + 9, body[9:11], b'\0\0')
patch('height-local', raw, code + 11, body[11:12], b'\x3e')
patch('height-end', raw, code + 19, body[19:20], b'\0')
class_name = 'LongWidthMain.class'
raw = (source / class_name).read_bytes()
_, main, body = methods(raw)['main([Ljava/lang/String;)V']
_, bump, bump_body = methods(raw)['bump(J)J']
patch('long-split-pop', raw, main + 164, b'\x58', b'\x57')
patch('long-uninit-local', raw, main + 3, b'\x40', b'\x41')
patch('long-wrong-ldc', raw, main + 2, body[2:3], b'\x05')
patch('long-wrong-return', raw, bump + 3, b'\xad', b'\xac')
patch('long-overflow', raw, main - 7, raw[main-7:main-6], b'\x01')
class_name = 'HelloMain.class'
raw = (source / class_name).read_bytes()
_, twice, twice_code = methods(raw)['twice(I)I']
patch('typed-ref-as-int', raw, twice + 1, b'\x1a', b'\x01')  # iload_0 -> null
class_name = 'ReferenceIdentityMain.class'
raw = (source / class_name).read_bytes()
_, identity, identity_code = methods(raw)['main([Ljava/lang/String;)V']
patch('typed-int-as-ref', raw, identity + 21, b'\x2b', b'\x03')  # aload_1 -> iconst_0
class_name = 'StackHeightMain.class'
raw = (source / class_name).read_bytes()
_, shape, shape_code = methods(raw)['main([Ljava/lang/String;)V']
patch('typed-ref-arraylength', raw, shape, b'\x2a', b'\x03')  # aload_0 -> iconst_0
patch('typed-merge', raw, shape + 5, b'\x07', b'\x01')  # iconst_4 -> null at merge
class_name = 'LongArrayMain.class'
raw = (source / class_name).read_bytes()
_, combine, combine_code = methods(raw)['combine([J)J']
_, array_main, array_main_code = methods(raw)['main([Ljava/lang/String;)V']
patch('long-array-index-reference', raw, combine + 1, b'\x03', b'\x01')
patch('long-array-store-category', raw, array_main + 18, b'\x50', b'\x4f')
newlong = array_main_code.index(b'\xbc\x0b')
patch('known-array-component-mismatch', raw, array_main + newlong + 1,
      b'\x0b', b'\x0a')  # newarray long -> int; laload must fail verifier
class_name = 'MultiArrayMain.class'
raw = (source / class_name).read_bytes()
_, clinit, clinit_body = methods(raw)['<clinit>()V']
assert clinit_body.count(b'\xc5') == 1, 'expected one static multi-array allocation'
array_at = clinit_body.index(b'\xc5')
assert clinit_body[array_at+3] == 2
patch('multi-array-zero-dimensions', raw, clinit + array_at + 3, b'\x02', b'\x00')
patch('multi-array-over-rank', raw, clinit + array_at + 3, b'\x02', b'\x03')
assert clinit_body[array_at-1:array_at] == b'\x04'
patch('multi-array-reference-count', raw, clinit + array_at - 1, b'\x04', b'\x01')
class_name = 'FloatingMain.class'
raw = (source / class_name).read_bytes()
_, apply, apply_body = methods(raw)['apply(D)D']
_, smaller, smaller_body = methods(raw)['smaller(F)F']
_, floating_main, floating_body = methods(raw)['main([Ljava/lang/String;)V']
assert apply_body == b'\x2a\xb4\x00\x04\x27\x63\xaf'
assert smaller_body == b'\x22\x0d\x6e\xae'
patch('double-argument-as-long', raw, apply + 4, b'\x27', b'\x1f')
patch('double-arithmetic-as-long', raw, apply + 5, b'\x63', b'\x61')
patch('float-arithmetic-as-double', raw, smaller + 2, b'\x6e', b'\x6f')
assert floating_body.count(b'\xbc\x07') == 1, 'expected one double newarray'
array_at = floating_body.index(b'\xbc\x07')
patch('double-array-as-float', raw, floating_main + array_at + 1, b'\x07', b'\x06')
PY
while read -r cause fixture diagnostic; do
    file="$scratch/$cause/$fixture.class"
    if timeout 8s "$MAGNOLIA" build run --entry "$file" \
        > "$scratch/$cause.oak.log" 2>&1; then
        echo "FAIL: Magnolia accepted malformed $cause" >&2; exit 1
    fi
    grep -q "$diagnostic" "$scratch/$cause.oak.log" || {
        echo "FAIL: unexpected Magnolia rejection for $cause" >&2
        cat "$scratch/$cause.oak.log" >&2; exit 1
    }
    # This mutation lives in a later method: rejecting it at runtime after
    # the first result has printed is not equivalent to JDK preflight.
    if [[ "$cause" == long-shift-reference ]] &&
        grep -qx '27021597764222979' "$scratch/$cause.oak.log"; then
        echo 'FAIL: malformed long shift produced output before preflight' >&2
        exit 1
    fi
    if [[ "$cause" == long-array-store-category ]] &&
        grep -qx '0' "$scratch/$cause.oak.log"; then
        echo 'FAIL: malformed long array store produced output before preflight' >&2
        exit 1
    fi
    if timeout 8s java -Xverify:all -cp "$scratch/$cause" "$fixture" \
        > "$scratch/$cause.jdk.log" 2>&1; then
        echo "FAIL: JDK accepted malformed $cause" >&2; exit 1
    fi
    grep -Eq 'VerifyError|ClassFormatError' "$scratch/$cause.jdk.log" || {
        echo "FAIL: JDK failed for a reason other than verification: $cause" >&2
        cat "$scratch/$cause.jdk.log" >&2; exit 1
    }
    printf 'PASS %-19s rejected by both Magnolia preflight and JDK -Xverify:all\n' "$cause"
done <<'CASES'
height-underflow StackHeightMain operand.stack.underflow
height-overflow StackHeightMain operand.stack.exceeds.max_stack
height-merge StackHeightMain inconsistent.operand.stack.height
height-local StackHeightMain local.variable.index.exceeds.max_locals
height-end StackHeightMain control.flow.leaves.the.method
long-split-pop LongWidthMain wrong.value.category
long-uninit-local LongWidthMain invalid.or.uninitialized.long.frame.local
long-wrong-ldc LongWidthMain ldc2_w.requires.a.long.constant
long-wrong-return LongWidthMain ireturn.does.not.match.method.descriptor
long-overflow LongWidthMain operand.stack.exceeds.max_stack
typed-ref-as-int HelloMain expected.I,.found.A
typed-int-as-ref ReferenceIdentityMain expected.A,.found.I
typed-ref-arraylength StackHeightMain expected.A,.found.I
typed-merge StackHeightMain inconsistent.value.categories.at.merge
long-array-index-reference LongArrayMain expected.I,.found.A
long-array-store-category LongArrayMain operand.stack.exceeds.max_stack
known-array-component-mismatch LongArrayMain array.opcode.incompatible.with.verifier-known.component.type
multi-array-zero-dimensions MultiArrayMain invalid.or.unsupported.multianewarray
multi-array-over-rank MultiArrayMain invalid.or.unsupported.multianewarray
multi-array-reference-count MultiArrayMain expected.I,.found.A
double-argument-as-long FloatingMain invalid.or.uninitialized.long.frame.local
double-arithmetic-as-long FloatingMain expected.J,.found.D
float-arithmetic-as-double FloatingMain operand.stack.underflow
double-array-as-float FloatingMain array.opcode.incompatible.with.verifier-known.component.type
CASES
