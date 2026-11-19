#!/usr/bin/env bash
# Cross-backend IEEE-754 bit roundtrips, with Python's independent binary
# struct packing as the finite/signed-zero input oracle. NaN payloads enter
# via .class pool tests; arithmetic NaNs may legally canonicalize.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" node python3 timeout; do
    command -v "$executable" >/dev/null 2>&1 || { echo "missing: $executable" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
printf '[fp-bits] generate 512 deterministic finite/signed-zero bit cases\n' >&2
python3 - "$scratch/probe.oak" "$(pwd)/lib/ieee754.oak" <<'PY'
import random, struct, sys
from pathlib import Path
rng = random.Random(0x1eeef754)
base64 = [0, 1, 2, (1<<52)-1, 1<<52, (1<<52)+1,
          0x3fe0000000000000, 0x3ff0000000000000,
          0x7fefffffffffffff, 0x7fe0000000000000,
          0x8000000000000000, 0x8000000000000001]
base32 = [0, 1, 2, (1<<23)-1, 1<<23, (1<<23)+1,
          0x3f000000, 0x3f800000, 0x7f7fffff, 0x7f000000,
          0x80000000, 0x80000001]
# Cover all exponent decades as well as arbitrary finite mantissas/signs.
for exponent in range(0, 2047, 31):
    base64.append((exponent << 52) | rng.getrandbits(52))
for exponent in range(0, 255, 4):
    base32.append((exponent << 23) | rng.getrandbits(23))
while len(base64) < 256:
    bits = rng.getrandbits(64)
    if (bits >> 52) & 2047 != 2047: base64.append(bits)
while len(base32) < 256:
    bits = rng.getrandbits(32)
    if (bits >> 23) & 255 != 255: base32.append(bits)
base64, base32 = base64[:256], base32[:256]
for bits in base64:
    assert struct.unpack('>Q', struct.pack('>d', struct.unpack('>d',
        bits.to_bytes(8, 'big'))[0]))[0] == bits
for bits in base32:
    assert struct.unpack('>I', struct.pack('>f', struct.unpack('>f',
        bits.to_bytes(4, 'big'))[0]))[0] == bits
lines=["std := import('std')", f"fp := import('{sys.argv[2]}')",
       "doubles := [" + ','.join(f'[{n >> 32}, {n & 4294967295}]' for n in base64) + '] ',
       "floats := [" + ','.join(str(n) for n in base32) + '] ',
       'mismatches := 0',
       "doubles |> std.each(fn(pair) {",
       "    raw := fp.f64(pair.(0), pair.(1))",
       "    rebuilt := fp.fromNumber(fp.numeric(raw), true, fp.sign(raw) = 1)",
       "    if rebuilt.hi != raw.hi | rebuilt.lo != raw.lo -> {",
       "        std.println('double mismatch: ' + string(pair) + ' -> ' + string(rebuilt))",
       "        mismatches <- mismatches + 1",
       "    }",
       "})",
       "floats |> std.each(fn(bits) {",
       "    raw := fp.f32(bits)",
       "    rebuilt := fp.fromNumber(fp.numeric(raw), false, fp.sign(raw) = 1)",
       "    if rebuilt.bits != bits -> {",
       "        std.println('float mismatch: ' + string(bits) + ' -> ' + string(rebuilt))",
       "        mismatches <- mismatches + 1",
       "    }",
       "})",
       "if mismatches != 0 -> exit(1)",
       "std.println('PASS IEEE finite bit roundtrip 512/512')"]
Path(sys.argv[1]).write_text('\n'.join(lines) + '\n', encoding='utf-8')
PY
printf '[fp-bits] native Oak (35s cap)\n' >&2
timeout --signal=TERM --kill-after=2s 35s "$MAGNOLIA" "$scratch/probe.oak" >"$scratch/native.log"
cat "$scratch/native.log"
printf '[fp-bits] self-contained JS compile (40s cap)\n' >&2
timeout --signal=TERM --kill-after=2s 40s "$MAGNOLIA" build --entry \
    "$scratch/probe.oak" --output "$scratch/probe.js" --web >/dev/null
printf '[fp-bits] self-contained JS execution (18s cap)\n' >&2
timeout --signal=TERM --kill-after=2s 18s node "$scratch/probe.js" >"$scratch/js.log"
diff -u "$scratch/native.log" "$scratch/js.log"
cat "$scratch/js.log"
