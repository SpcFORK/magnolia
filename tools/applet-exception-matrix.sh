#!/usr/bin/env bash
# Independent Java 8 oracle for ordered exception tables, athrow, catch-all
# finally, implicit Java exceptions, cross-call propagation, and checked edits.
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
    printf '[exceptions] START %s (max %ss)\n' "$label" "$seconds" >&2
    if timeout --signal=TERM --kill-after=3s "$seconds" "$@"; then
        printf '[exceptions] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[exceptions] FAIL %s exit=%s (%ss)\n' "$label" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
fixture=test/fixtures/ExceptionMain.class
run 'reproduce javac --release 8 fixture' 20 javac --release 8 -g:none -d "$scratch/original" test/fixtures/ExceptionMain.java
cmp "$fixture" "$scratch/original/ExceptionMain.class"
run 'JDK verifier and 13-line oracle' 12 java -Xverify:all -cp "$scratch/original" ExceptionMain >"$scratch/jdk.out"
diff -u test/fixtures/ExceptionMain.expected "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 13 ]]
run 'native Oak original class' 20 "$MAGNOLIA" build run --entry "$fixture" >"$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run 'build standalone Oak with bounded signed handler IR' 40 "$MAGNOLIA" build --entry "$fixture" --no-optimize --output "$scratch/ExceptionMain.oak" >/dev/null
run 'execute standalone Oak' 20 "$MAGNOLIA" "$scratch/ExceptionMain.oak" >"$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
run 'build self-contained JS' 95 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/ExceptionMain.js" >/dev/null
run 'execute self-contained JS' 25 node "$scratch/ExceptionMain.js" >"$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run 'build offline HTML' 110 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/ExceptionMain.html" >/dev/null
run 'execute offline inline HTML' 20 node - "$scratch/ExceptionMain.html" "$scratch/jdk.out" <<'JS'
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
run 'generate Java' 45 "$MAGNOLIA" build --entry "$fixture" --output "$scratch/ExceptionConverted.java" >/dev/null
run 'compile generated Java' 20 javac -d "$scratch/generated" "$scratch/ExceptionConverted.java"
run 'execute generated Java' 12 java -Xverify:all -cp "$scratch/generated" ExceptionConverted >"$scratch/java.out"
diff -u "$scratch/jdk.out" "$scratch/java.out"
run 'restore original byte-exactly' 15 "$MAGNOLIA" build restore --entry "$scratch/ExceptionMain.oak" --output "$scratch/restored.class" >/dev/null
cmp "$fixture" "$scratch/restored.class"
run 'zero-edit recompile keeps handler bytes' 40 "$MAGNOLIA" build --entry "$scratch/ExceptionMain.oak" --output "$scratch/original-recompiled.class" >/dev/null
cmp "$fixture" "$scratch/original-recompiled.class"
python3 - "$scratch/ExceptionMain.oak" "$scratch/ExceptionEdited.oak" "$scratch/Tampered.oak" <<'PY'
from pathlib import Path
import sys
source=Path(sys.argv[1]).read_text()
start=source.index('// ExceptionMain.failInFinally()V\n')
end=source.index('// ExceptionMain.main(', start)
block=source[start:end]
old='    // 9: bipush\n    [16, 17, ?],'
assert block.count(old)==1
Path(sys.argv[2]).write_text(source[:start]+block.replace(old,old.replace('17','18'))+source[end:])
# Handler metadata is not an instruction edit. Reverse compilation MUST refuse
# to discard this change and instead preserve the original exception table.
marker="'failInFinally()V': {locals: 1, stack: 3, argWidths: [], handlers: [{'catch': ?\n'end': 5"
assert source.count(marker)==1
Path(sys.argv[3]).write_text(source.replace(marker, marker.replace("'end': 5", "'end': 4")))
PY
run 'checked finally edit -> NEW .class' 45 "$MAGNOLIA" build --entry "$scratch/ExceptionEdited.oak" --output "$scratch/edited/ExceptionMain.class" >/dev/null
! cmp -s "$fixture" "$scratch/edited/ExceptionMain.class"
run 'JDK verifier on edited .class' 12 java -Xverify:all -cp "$scratch/edited" ExceptionMain >"$scratch/edited.jdk"
run 'execute edited standalone Oak' 20 "$MAGNOLIA" "$scratch/ExceptionEdited.oak" >"$scratch/edited.oak"
diff -u "$scratch/edited.jdk" "$scratch/edited.oak"
python3 - "$scratch/jdk.out" "$scratch/edited.jdk" <<'PY'
from pathlib import Path
import sys
original,edited=[Path(p).read_text().splitlines() for p in sys.argv[1:]]
assert len(original)==len(edited)==13
assert [(i+1,a,b) for i,(a,b) in enumerate(zip(original,edited)) if a!=b]==[(12,'17','18')]
PY
run 'restore pre-edit original byte-exactly' 15 "$MAGNOLIA" build restore --entry "$scratch/ExceptionEdited.oak" --output "$scratch/edited-original.class" >/dev/null
cmp "$fixture" "$scratch/edited-original.class"
if timeout --signal=TERM --kill-after=3s 40s "$MAGNOLIA" build --entry "$scratch/Tampered.oak" --output "$scratch/tampered.class" >"$scratch/rejected.log" 2>&1; then
    echo 'FAIL: changed handler table was silently accepted' >&2; exit 1
fi
grep -q 'method stack/locals/flags/argument slots changed' "$scratch/rejected.log" || {
    echo 'FAIL: unexpected handler rejection'; cat "$scratch/rejected.log"; exit 1;
}
printf 'PASS 13 JDK exception lines = native/Oak/JS/HTML/Java; checked edited class; handler tamper refused; original restored\n'
