#!/usr/bin/env bash
# Independent JDK conversion oracle: signed 64-bit long -> binary32 must
# round ONCE, not through binary64. Also checks long -> binary64 and targets.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for cmd in "$MAGNOLIA" java javac python3 node timeout diff cmp; do
    command -v "$cmd" >/dev/null 2>&1 || { echo "missing: $cmd" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/reproduced" "$scratch/generated" "$scratch/invalid"
run() {
    local label="$1" seconds="$2" start=$SECONDS rc
    shift 2
    printf '[long-float] START %s (max %ss)\n' "$label" "$seconds" >&2
    if timeout --signal=TERM --kill-after=2s "$seconds" "$@"; then
        printf '[long-float] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[long-float] FAIL %s exit=%s (%ss)\n' "$label" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
fixture=test/fixtures/LongFloatRoundingMain.class
run 'reproduce Java 8 class' 12 javac --release 8 -g:none -d "$scratch/reproduced" test/fixtures/LongFloatRoundingMain.java
cmp "$fixture" "$scratch/reproduced/LongFloatRoundingMain.class"
run 'JDK -Xverify:all 18 boundary longs' 8 java -Xverify:all -cp "$scratch/reproduced" LongFloatRoundingMain >"$scratch/jdk.out"
diff -u test/fixtures/LongFloatRoundingMain.expected "$scratch/jdk.out"
[[ $(wc -l <"$scratch/jdk.out") == 36 ]]
run 'native Oak' 12 "$MAGNOLIA" build run --entry "$fixture" >"$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run 'build standalone Oak' 23 "$MAGNOLIA" build --entry "$fixture" --no-optimize --output "$scratch/LongFloatRoundingMain.oak" >/dev/null
run 'execute standalone Oak' 10 "$MAGNOLIA" "$scratch/LongFloatRoundingMain.oak" >"$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
run 'build bundled JS' 25 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/LongFloatRoundingMain.js" >/dev/null
run 'execute bundled JS' 12 node "$scratch/LongFloatRoundingMain.js" >"$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
if [[ "${1:-}" == '--skip-html' ]]; then
    printf '[long-float] SKIP offline HTML (optional quick subset only)\n' >&2
else
    run 'build offline HTML' 35 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/LongFloatRoundingMain.html" >/dev/null
    run 'execute offline HTML' 10 node - "$scratch/LongFloatRoundingMain.html" "$scratch/jdk.out" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.equal(scripts.length, 2);
const output = {textContent: ''}, canvas = {style: {}};
const context = {document: {getElementById(id) {return id === 'output' ? output : id === 'drawing' ? canvas : null;}},
    console: {log() {}, error(...args) {throw Error(args.join(' '));}}};
vm.createContext(context);
for (const script of scripts) vm.runInContext(script, context, {timeout: 6000});
assert.equal(output.textContent, expected);
JS
fi
run 'generate Java from VM image' 25 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/LongFloatConverted.java" >/dev/null
run 'compile generated Java' 12 javac -d "$scratch/generated" "$scratch/LongFloatConverted.java"
run 'execute generated Java' 8 java -Xverify:all -cp "$scratch/generated" LongFloatConverted >"$scratch/generated.out"
diff -u "$scratch/jdk.out" "$scratch/generated.out"
run 'restore byte-exact original' 10 "$MAGNOLIA" build restore --entry "$scratch/LongFloatRoundingMain.oak" --output "$scratch/restored.class" >/dev/null
cmp "$fixture" "$scratch/restored.class"
python3 - "$fixture" "$scratch/invalid/LongFloatRoundingMain.class" <<'PY'
from pathlib import Path
import sys
image=Path(sys.argv[1]).read_bytes()
pattern=bytes([30,137,174])  # lload_0; l2f; freturn
assert image.count(pattern)==1
Path(sys.argv[2]).write_bytes(image.replace(pattern, bytes([30,134,174]))) # i2f on a long
PY
if timeout --signal=TERM --kill-after=2s 8s java -Xverify:all -cp "$scratch/invalid" LongFloatRoundingMain >"$scratch/bad.jdk" 2>&1; then
    echo 'FAIL: JDK accepted wrong-category conversion' >&2; exit 1
fi
grep -Eq 'VerifyError|ClassFormatError' "$scratch/bad.jdk" || {
    echo 'FAIL: unexpected JDK diagnostic' >&2; cat "$scratch/bad.jdk"; exit 1;
}
if timeout --signal=TERM --kill-after=2s 8s "$MAGNOLIA" build run \
        --entry "$scratch/invalid/LongFloatRoundingMain.class" >"$scratch/bad.oak" 2>&1; then
    echo 'FAIL: Magnolia accepted wrong-category conversion' >&2; exit 1
fi
grep -Eq 'wrong value category|operand stack|requires' "$scratch/bad.oak" || {
    echo 'FAIL: unexpected Magnolia diagnostic' >&2; cat "$scratch/bad.oak"; exit 1;
}
# Seeded stress fixture is generated only in a temporary directory, keeping
# independent JDK compilation and Oracle checks for adjacent f32 midpoints.
python3 - "$scratch/LongFloatSeeds.java" <<'PY'
from pathlib import Path
import random,sys
values={0,1,-1,(1<<63)-1,-(1<<63)}
for exp in range(24,63):
    center=(1<<exp)+(1<<(exp-24))
    for delta in (-1,0,1):
        for sign in (1,-1):values.add(sign*(center+delta))
rng=random.Random(632421)
for _ in range(148):
    v=rng.getrandbits(64)
    values.add(v if v<(1<<63) else v-(1<<64))
values=sorted(values)
assert len(values)>370
source='public class LongFloatSeeds { public static void main(String[] args) { long[] samples = {\n'
source+=',\n'.join(str(v)+'L' if v!=-(1<<63) else 'Long.MIN_VALUE' for v in values)
source+='}; for (long v : samples) System.out.println(Float.floatToRawIntBits((float)v)); } }\n'
Path(sys.argv[1]).write_text(source)
print('[long-float] seeded oracle values:',len(values),file=sys.stderr)
PY
run 'javac seeded midpoint corpus' 12 javac --release 8 -g:none -d "$scratch" "$scratch/LongFloatSeeds.java"
run 'JDK -Xverify:all seeded oracle' 8 java -Xverify:all -cp "$scratch" LongFloatSeeds >"$scratch/seeds.jdk"
run 'native Oak seeded oracle' 12 "$MAGNOLIA" build run --entry "$scratch/LongFloatSeeds.class" >"$scratch/seeds.oak"
diff -u "$scratch/seeds.jdk" "$scratch/seeds.oak"
run 'build seeded JS' 25 "$MAGNOLIA" build --entry "$scratch/LongFloatSeeds.class" --output "$scratch/LongFloatSeeds.js" >/dev/null
run 'execute seeded JS' 12 node "$scratch/LongFloatSeeds.js" >"$scratch/seeds.js"
diff -u "$scratch/seeds.jdk" "$scratch/seeds.js"
printf 'PASS JDK bits for 18 boundary longs on Oak/JS/Java%s; seeded JDK/native/JS midpoint corpus; invalid conversion rejected; exact restoration\n' "$(if [[ "${1:-}" == '--skip-html' ]]; then printf ' (HTML skipped)'; else printf '/HTML'; fi)"
