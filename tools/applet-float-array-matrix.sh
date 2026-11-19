#!/usr/bin/env bash
# Java 8/JDK verifier oracle for typed float/double arrays, exact bits,
# conversions, independent backends and checked same-width bytecode edits.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for cmd in "$MAGNOLIA" java javac node python3 timeout diff cmp; do
    command -v "$cmd" >/dev/null 2>&1 || { echo "missing: $cmd" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/generated" "$scratch/edited"
run() {
    local label="$1" seconds="$2" start=$SECONDS rc
    shift 2
    printf '[fp-arrays] START %s (max %ss)\n' "$label" "$seconds" >&2
    if timeout --signal=TERM --kill-after=3s "$seconds" "$@"; then
        printf '[fp-arrays] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[fp-arrays] FAIL %s exit=%s (%ss)\n' "$label" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
fixture=test/fixtures/FloatArrayMain.class
run 'reproduce Java 8 fixture' 20 javac --release 8 -g:none -d "$scratch/original" test/fixtures/FloatArrayMain.java
cmp "$fixture" "$scratch/original/FloatArrayMain.class"
run 'independent JDK oracle' 12 java -Xverify:all -cp "$scratch/original" FloatArrayMain >"$scratch/jdk.out"
diff -u test/fixtures/FloatArrayMain.expected "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 24 ]]
run 'native Oak class' 20 "$MAGNOLIA" build run --entry "$fixture" >"$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run 'build standalone Oak' 35 "$MAGNOLIA" build --entry "$fixture" --no-optimize --output "$scratch/FloatArrayMain.oak" >/dev/null
run 'execute standalone Oak' 18 "$MAGNOLIA" "$scratch/FloatArrayMain.oak" >"$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
run 'build bundled JS' 95 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/FloatArrayMain.js" >/dev/null
run 'execute bundled JS' 25 node "$scratch/FloatArrayMain.js" >"$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run 'build offline HTML' 110 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/FloatArrayMain.html" >/dev/null
run 'execute offline inline HTML' 20 node - "$scratch/FloatArrayMain.html" "$scratch/jdk.out" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2);
const output = {textContent: ''}, canvas = {style: {}};
const ctx = {document: {getElementById(id) {return id === 'output' ? output : id === 'drawing' ? canvas : null;}},
             console: {log() {}, error(...args) {throw Error(args.join(' '));}}};
vm.createContext(ctx);
for (const script of scripts) vm.runInContext(script, ctx, {timeout: 10000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(ctx.MagnoliaApplet.inspect().type, 'ok');
JS
run 'generate Java' 45 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/FloatArrayConverted.java" >/dev/null
run 'compile generated Java' 20 javac -d "$scratch/generated" "$scratch/FloatArrayConverted.java"
run 'execute generated Java' 12 java -Xverify:all -cp "$scratch/generated" FloatArrayConverted >"$scratch/java.out"
diff -u "$scratch/jdk.out" "$scratch/java.out"
run 'restore original class' 20 "$MAGNOLIA" build restore --entry "$scratch/FloatArrayMain.oak" --output "$scratch/restore.class" >/dev/null
cmp "$fixture" "$scratch/restore.class"
python3 - "$scratch/FloatArrayMain.oak" "$scratch/FloatArrayEdited.oak" <<'PY'
from pathlib import Path
import sys
src = Path(sys.argv[1]).read_text()
old = '[111, 0, ?],'
new = '[107, 0, ?],'
assert src.count(old) == 1 and 'FloatArrayMain.returnDouble([D)D' in src
Path(sys.argv[2]).write_text(src.replace(old, new))
PY
run 'checked edit ddiv -> dmul' 45 "$MAGNOLIA" build --entry "$scratch/FloatArrayEdited.oak" --output "$scratch/edited/FloatArrayMain.class" >/dev/null
! cmp -s "$fixture" "$scratch/edited/FloatArrayMain.class"
run 'independent JDK-verified edited bytecode' 12 java -Xverify:all -cp "$scratch/edited" FloatArrayMain >"$scratch/edited.jdk"
run 'execute edited standalone Oak' 18 "$MAGNOLIA" "$scratch/FloatArrayEdited.oak" >"$scratch/edited.oak"
diff -u "$scratch/edited.jdk" "$scratch/edited.oak"
! cmp -s "$scratch/jdk.out" "$scratch/edited.jdk"
run 'restore pre-edit original' 20 "$MAGNOLIA" build restore --entry "$scratch/FloatArrayEdited.oak" --output "$scratch/edited-original.class" >/dev/null
cmp "$fixture" "$scratch/edited-original.class"
printf 'PASS 24 JDK lines = native/standalone/JS/HTML/Java, edited JDK class matches edited Oak, exact original restored\n'
