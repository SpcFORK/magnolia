#!/usr/bin/env bash
# Exercise the Oak selective ZIP reader on a JAR too large for eager entries().
# Python generates disposable adversarial ZIP32 fixtures; it does not parse or
# execute JARs in Magnolia, and no third-party binaries enter the repository.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
command -v "$MAGNOLIA" >/dev/null || { echo "missing Magnolia executable" >&2; exit 2; }
command -v python3 >/dev/null || { echo "missing Python fixture generator" >&2; exit 2; }
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
python3 - "$scratch" test/fixtures/CounterApplet.class <<'PY'
from pathlib import Path
import io
import struct
import sys
import warnings
import zipfile

out, source = Path(sys.argv[1]), Path(sys.argv[2])
cls = source.read_bytes()

def jar(entries):
    stream = io.BytesIO()
    with warnings.catch_warnings():
        warnings.simplefilter('ignore', UserWarning)
        with zipfile.ZipFile(stream, 'w', compression=zipfile.ZIP_DEFLATED) as archive:
            for name, value in entries:
                archive.writestr(name, value)
    return stream.getvalue()

def save(name, raw): (out / (name + '.jar')).write_bytes(raw)

# 4,101 entries but only one class; ordinary entries() must still refuse it.
save('large', jar([('CounterApplet.class', cls)] +
    [(f'assets/{i:05d}.dat', b'') for i in range(4100)]))
normal = jar([('CounterApplet.class', cls), ('assets/ok.txt', b'hello')])
save('normal', normal)
save('duplicate', jar([('CounterApplet.class', cls), ('CounterApplet.class', cls)]))
save('unsafe', jar([('CounterApplet.class', cls), ('../escape.txt', b'bad')]))

def mutate(name, callback):
    data = bytearray(normal)
    end = len(data) - 22
    assert data[end:end+4] == b'PK\x05\x06'
    central = struct.unpack_from('<I', data, end+16)[0]
    assert data[central:central+4] == b'PK\x01\x02'
    callback(data, central, end)
    save(name, data)

def wrong_crc(b, p, e):
    # Keep central and local headers internally consistent; the payload must
    # still fail the independent CRC of the *actual uncompressed bytes*.
    checksum = struct.unpack_from('<I', b, p+16)[0] ^ 1
    struct.pack_into('<I', b, p+16, checksum)
    struct.pack_into('<I', b, 14, checksum)
mutate('crc', wrong_crc)
mutate('local-name', lambda b, p, e: b.__setitem__(30, b[30] ^ 1))
mutate('local-crc', lambda b, p, e: b.__setitem__(14, b[14] ^ 1))
mutate('local-size', lambda b, p, e: b.__setitem__(22, b[22] ^ 1))
mutate('flag-mismatch', lambda b, p, e: b.__setitem__(p+8, b[p+8] ^ 8))
mutate('ratio', lambda b, p, e: struct.pack_into('<I', b, p+20, 1))
mutate('zip64-size', lambda b, p, e: struct.pack_into('<I', b, p+24, 0xffffffff))
mutate('encrypted', lambda b, p, e: b.__setitem__(p+8, b[p+8] | 1))
mutate('unsupported-method', lambda b, p, e: struct.pack_into('<H', b, p+10, 12))
mutate('entry-count', lambda b, p, e: (struct.pack_into('<H', b, e+8, 50001),
                                     struct.pack_into('<H', b, e+10, 50001)))
def corrupt_later(b, p, e):
    name_len, extra_len, comment_len = struct.unpack_from('<HHH', b, p+28)
    second = p + 46 + name_len + extra_len + comment_len
    assert b[second:second+4] == b'PK\x01\x02'
    b[second] ^= 1
mutate('later-central', corrupt_later)
PY
cat > "$scratch/verify.oak" <<OAK
std := import('std')
str := import('str')
fs := import('fs')
zip := import('$PWD/lib/applet-zip.oak')
fn require(name, condition) {
    if !condition {
        true -> {
            std.println('FAIL ' + name)
            exit(1)
        }
        _ -> std.println('PASS ' + name)
    }
}
original := fs.readFile('test/fixtures/CounterApplet.class')
large := fs.readFile('$scratch/large.jar')
chosen := zip.selectClass(large, 'CounterApplet')
require('select one class from 4101-entry ZIP32', chosen.type = :ok &
    chosen.value = original & chosen.archiveEntries = 4101)
require('ordinary whole-JAR bridge still enforces entry cap',
    zip.entries(large).type = :error)
normal := zip.selectClass(fs.readFile('$scratch/normal.jar'), 'CounterApplet.class')
require('select preserves original class bytes', normal.type = :ok & normal.value = original)
[
    ['duplicate', 'duplicate selected ZIP class'],
    ['unsafe', 'unsafe ZIP entry name'],
    ['crc', 'CRC32 mismatch'],
    ['local-name', 'inconsistent ZIP local header'],
    ['local-crc', 'inconsistent ZIP local header'],
    ['local-size', 'inconsistent ZIP local header'],
    ['flag-mismatch', 'inconsistent ZIP local header'],
    ['ratio', 'compression-ratio limit'],
    ['zip64-size', 'ZIP64 extension'],
    ['encrypted', 'encrypted JAR member'],
    ['unsupported-method', 'unsupported ZIP compression method'],
    ['entry-count', 'oversized selective JAR index'],
    ['later-central', 'invalid selective ZIP central-directory entry']
] |> std.each(fn(case) {
    selected := zip.selectClass(fs.readFile('$scratch/' + case.(0) + '.jar'), 'CounterApplet')
    require(case.(0) + ' rejected', selected.type = :error &
        str.contains?(selected.message, case.(1)))
})
OAK
"$MAGNOLIA" "$scratch/verify.oak"
# Cover the user-facing Builder flag, not only the library function.
"$MAGNOLIA" build --entry "$scratch/large.jar" \
    --select-class CounterApplet --param name=Ada \
    --output "$scratch/Selected.oak" > "$scratch/build.log" 2>&1 || {
    cat "$scratch/build.log" >&2; exit 1;
}
"$MAGNOLIA" "$scratch/Selected.oak" > "$scratch/run.log" 2>&1
printf 'PASS Builder selective-JAR standalone Oak output: '
tr '\n' '/' < "$scratch/run.log"
printf '\n'
if ! cmp -s "$scratch/run.log" <(printf 'Hello from Magnolia\n8\n14\nAda\n'); then
    echo 'FAIL: unexpected selected applet behavior' >&2; exit 1
fi
if "$MAGNOLIA" build --entry "$scratch/large.jar" --select-class CounterApplet \
    --entry-class DoesNotMatch --output "$scratch/Bad.oak" > "$scratch/bad.log" 2>&1; then
    echo 'FAIL: Builder accepted mismatched entry class' >&2; exit 1
fi
grep -q 'entryClass must match selectClass' "$scratch/bad.log"
printf 'PASS Builder refuses inconsistent selected entry class\n'
