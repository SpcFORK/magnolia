#!/usr/bin/env bash
# Selected Java 8 long[] oracle: exact limbs, nested arrays, arraycopy,
# self-contained Oak/JS/HTML/Java and checked edited class/JAR output.
# No network; external JDK is an oracle/build-time tool, not a delivered runtime.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for tool in "$MAGNOLIA" java javac node python3 timeout cmp diff; do
    command -v "$tool" >/dev/null 2>&1 || {
        echo "missing long-array matrix dependency: $tool" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/translated" "$scratch/edited"
log_run() {
    local label="$1" seconds="$2" result started=$SECONDS
    shift 2
    printf '[%s] START %s (max %ss):' "$(date '+%H:%M:%S')" "$label" "$seconds" >&2
    printf ' %q' "$@" >&2
    printf '\n' >&2
    if timeout --signal=TERM --kill-after=3s "$seconds" "$@"; then
        printf '[%s] PASS %s (%ss)\n' "$(date '+%H:%M:%S')" "$label" "$((SECONDS-started))" >&2
    else
        result=$?
        printf '[%s] FAIL %s exit=%s (%ss)\n' "$(date '+%H:%M:%S')" "$label" "$result" "$((SECONDS-started))" >&2
        return "$result"
    fi
}
fixture=test/fixtures/LongArrayMain.class
log_run 'javac --release 8' 15 javac --release 8 -g:none \
    -d "$scratch/original" test/fixtures/LongArrayMain.java
cmp "$fixture" "$scratch/original/LongArrayMain.class"
log_run 'independent JDK verification/oracle' 10 java -Xverify:all \
    -cp "$scratch/original" LongArrayMain > "$scratch/jdk.output"
diff -u test/fixtures/LongArrayMain.expected "$scratch/jdk.output"
printf 'PASS fixture bytes reproducible; JDK oracle has %s exact values\n' \
    "$(wc -l < "$scratch/jdk.output")"

log_run 'native Oak .class' 20 "$MAGNOLIA" build run \
    --entry "$fixture" > "$scratch/native.output"
log_run 'isolated javac .java input' 20 "$MAGNOLIA" build run \
    --entry test/fixtures/LongArrayMain.java --java-backend javac \
    > "$scratch/javac-input.output"
for route in native javac-input; do
    diff -u "$scratch/jdk.output" "$scratch/$route.output"
done
printf 'PASS .class and isolated .java input use the same long-array VM\n'

log_run 'standalone Oak build' 30 "$MAGNOLIA" build --entry "$fixture" \
    --no-optimize --output "$scratch/LongArrayMain.oak" >/dev/null
log_run 'standalone Oak run' 12 "$MAGNOLIA" "$scratch/LongArrayMain.oak" \
    > "$scratch/oak.output"
diff -u "$scratch/jdk.output" "$scratch/oak.output"
log_run 'Builder JS build' 75 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/LongArrayMain.js" >/dev/null
log_run 'self-contained JS run' 15 node "$scratch/LongArrayMain.js" \
    > "$scratch/js.output"
diff -u "$scratch/jdk.output" "$scratch/js.output"
log_run 'Builder HTML build' 110 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/LongArrayMain.html" >/dev/null
log_run 'offline inline-HTML DOM oracle' 20 node - \
    "$scratch/LongArrayMain.html" "$scratch/jdk.output" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2, 'browser output must contain exactly two inline scripts');
const output = {textContent: ''};
const canvas = {style: {}, getContext() { throw Error('unexpected drawing'); }};
const context = {document: {getElementById(id) {
    return id === 'output' ? output : id === 'drawing' ? canvas : null;
}}, console: {log() {}, error(...args) {throw Error(args.join(' '));}}};
vm.createContext(context);
for (const script of scripts) vm.runInContext(script, context, {timeout: 10000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(context.MagnoliaApplet.inspect().type, 'ok');
JS
printf 'PASS JDK = native Oak = standalone Oak = JS = offline HTML\n'

log_run 'generated Java source' 30 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/LongArrayConverted.java" >/dev/null
log_run 'compile generated Java' 20 javac -d "$scratch/translated" \
    "$scratch/LongArrayConverted.java"
log_run 'execute generated Java' 10 java -Xverify:all \
    -cp "$scratch/translated" LongArrayConverted > "$scratch/java.output"
diff -u "$scratch/jdk.output" "$scratch/java.output"
log_run 'restore original class bytes' 12 "$MAGNOLIA" build restore \
    --entry "$scratch/LongArrayMain.oak" \
    --output "$scratch/original-restored.class" >/dev/null
cmp "$fixture" "$scratch/original-restored.class"
printf 'PASS generated Java matches JDK; original class restored byte-exactly\n'

# The one same-width edit is inside combine([J)J: lxor -> lor. It changes
# exactly one printed value while retaining valid two-slot array loads.
python3 - "$scratch/LongArrayMain.oak" "$scratch/LongArrayEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
old, new = '[131, 0, ?],', '[129, 0, ?],'
assert source.count(old) == 1 and source.count('// LongArrayMain.combine([J)J') == 1
open(sys.argv[2], 'w', encoding='utf-8').write(source.replace(old, new))
PY
log_run 'checked Oak edit -> NEW class' 25 "$MAGNOLIA" build \
    --entry "$scratch/LongArrayEdited.oak" \
    --output "$scratch/edited/LongArrayMain.class" >/dev/null
if cmp -s "$fixture" "$scratch/edited/LongArrayMain.class"; then
    echo 'FAIL edited class equals original bytes' >&2; exit 1
fi
log_run 'JDK verify edited class' 10 java -Xverify:all \
    -cp "$scratch/edited" LongArrayMain > "$scratch/edited.jdk"
log_run 'edited Oak execution' 12 "$MAGNOLIA" \
    "$scratch/LongArrayEdited.oak" > "$scratch/edited.oak"
diff -u "$scratch/edited.jdk" "$scratch/edited.oak"
python3 - "$scratch/jdk.output" "$scratch/edited.jdk" <<'PY'
import sys
before, after = [open(name, encoding='utf-8').readlines() for name in sys.argv[1:]]
assert len(before) == len(after) == 14
assert before[5] == '-9007199254740994\n' and after[5] == '-1\n'
assert before[:5] + before[6:] == after[:5] + after[6:]
PY
log_run 'restore original despite edited Oak' 12 "$MAGNOLIA" build restore \
    --entry "$scratch/LongArrayEdited.oak" \
    --output "$scratch/edited-original.class" >/dev/null
cmp "$fixture" "$scratch/edited-original.class"
printf 'PASS lxor -> lor changes exactly one value in a NEW JDK-verified class; restore still exact\n'

python3 - "$fixture" "$scratch/LongArrayMain.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[2], 'w', compression=zipfile.ZIP_DEFLATED) as jar:
    jar.writestr('META-INF/MANIFEST.MF',
                 'Manifest-Version: 1.0\r\nMain-Class: LongArrayMain\r\n\r\n')
    jar.write(sys.argv[1], 'LongArrayMain.class')
    jar.writestr('assets/array.txt', 'keep this resource\n')
PY
log_run 'original JDK -jar' 10 java -Xverify:all \
    -jar "$scratch/LongArrayMain.jar" > "$scratch/original.jar-output"
log_run 'original JAR on Oak' 20 "$MAGNOLIA" build run \
    --entry "$scratch/LongArrayMain.jar" > "$scratch/original.jar-oak"
for route in original.jar-output original.jar-oak; do
    diff -u "$scratch/jdk.output" "$scratch/$route"
done
log_run 'JAR -> standalone Oak' 30 "$MAGNOLIA" build \
    --entry "$scratch/LongArrayMain.jar" --no-optimize \
    --output "$scratch/LongArrayJar.oak" >/dev/null
python3 - "$scratch/LongArrayJar.oak" "$scratch/LongArrayJarEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
assert source.count('[131, 0, ?],') == 1
open(sys.argv[2], 'w', encoding='utf-8').write(
    source.replace('[131, 0, ?],', '[129, 0, ?],'))
PY
log_run 'checked Oak edit -> NEW JAR' 30 "$MAGNOLIA" build \
    --entry "$scratch/LongArrayJarEdited.oak" \
    --output "$scratch/LongArrayEdited.jar" >/dev/null
log_run 'JDK verify edited JAR' 10 java -Xverify:all \
    -jar "$scratch/LongArrayEdited.jar" > "$scratch/edited.jar-jdk"
diff -u "$scratch/edited.jdk" "$scratch/edited.jar-jdk"
log_run 'restore original JAR' 12 "$MAGNOLIA" build restore \
    --entry "$scratch/LongArrayJarEdited.oak" \
    --output "$scratch/LongArrayOriginal.jar" >/dev/null
cmp "$scratch/LongArrayMain.jar" "$scratch/LongArrayOriginal.jar"
python3 - "$scratch/LongArrayMain.jar" "$scratch/LongArrayEdited.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[1]) as original, zipfile.ZipFile(sys.argv[2]) as edited:
    assert original.namelist() == edited.namelist()
    for name in ('META-INF/MANIFEST.MF', 'assets/array.txt'):
        assert original.read(name) == edited.read(name)
    assert original.read('LongArrayMain.class') != edited.read('LongArrayMain.class')
PY
printf 'PASS edited JAR verifies on JDK, manifest/resource preserved, original JAR exact\n'

# Java exceptions cannot be caught by this subset yet. Verify that bounded
# Oak fails explicitly for the same three input categories, never succeeds.
for cause in bounds mismatch null; do
    case "$cause" in
        bounds) jdk_args=(a); oak_args='["a"]'; jdk_error='ArrayIndexOutOfBoundsException'; oak_error='array index out of bounds' ;;
        mismatch) jdk_args=(a b); oak_args='["a","b"]'; jdk_error='ArrayStoreException'; oak_error='System.arraycopy primitive array type mismatch' ;;
        null) jdk_args=(a b c); oak_args='["a","b","c"]'; jdk_error='NullPointerException'; oak_error='null array reference' ;;
    esac
    if timeout 10s java -Xverify:all -cp "$scratch/original" \
        LongArrayMain "${jdk_args[@]}" > "$scratch/$cause.jdk" 2>&1; then
        echo "FAIL JDK $cause unexpectedly succeeded" >&2; exit 1
    fi
    if timeout 10s "$MAGNOLIA" build run --entry "$fixture" \
        --args "$oak_args" > "$scratch/$cause.oak" 2>&1; then
        echo "FAIL Oak $cause unexpectedly succeeded" >&2; exit 1
    fi
    grep -q "$jdk_error" "$scratch/$cause.jdk"
    grep -q "$oak_error" "$scratch/$cause.oak"
    printf 'PASS %-8s JDK raises %s; Oak rejects explicitly (not catchable yet)\n' \
        "$cause" "$jdk_error"
done
log_run 'whole-input coverage' 12 "$MAGNOLIA" build coverage \
    --entry "$fixture" --json > "$scratch/coverage.json"
python3 - "$scratch/coverage.json" <<'PY'
import json, sys
coverage = json.load(open(sys.argv[1], encoding='utf-8'))
assert coverage['decodedBodies'] == coverage['methods'] == 4
assert coverage['allBodiesPreflighted?'] and not coverage['blockers']
PY
printf 'PASS all 4 fixture methods decoded; long[] route matrix complete\n'
