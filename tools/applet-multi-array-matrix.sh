#!/usr/bin/env bash
# Bounded JVM multianewarray cross-route oracle. The independent JDK verifies
# the checked-in class, edited NEW class/JAR and failure paths; delivered
# standalone Oak, JS and HTML do not depend on a browser-side Java runtime.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" java javac node python3 timeout cmp diff; do
    command -v "$executable" >/dev/null 2>&1 || {
        echo "missing multianewarray matrix dependency: $executable" >&2; exit 2;
    }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/generated" "$scratch/edited"
run_bounded() {
    local title="$1" limit="$2" start=$SECONDS status
    shift 2
    printf '[%s] START %s (max %ss):' "$(date '+%H:%M:%S')" "$title" "$limit" >&2
    printf ' %q' "$@" >&2
    printf '\n' >&2
    if timeout --signal=TERM --kill-after=3s "$limit" "$@"; then
        printf '[%s] PASS %s (%ss)\n' "$(date '+%H:%M:%S')" "$title" "$((SECONDS-start))" >&2
    else
        status=$?
        printf '[%s] FAIL %s exit=%s (%ss)\n' "$(date '+%H:%M:%S')" "$title" "$status" "$((SECONDS-start))" >&2
        return "$status"
    fi
}
fixture=test/fixtures/MultiArrayMain.class
run_bounded 'javac Java 8 fixture' 15 javac --release 8 -g:none \
    -d "$scratch/original" test/fixtures/MultiArrayMain.java
cmp "$fixture" "$scratch/original/MultiArrayMain.class"
run_bounded 'JDK verifier/oracle' 10 java -Xverify:all \
    -cp "$scratch/original" MultiArrayMain > "$scratch/jdk.out"
diff -u test/fixtures/MultiArrayMain.expected "$scratch/jdk.out"
printf 'PASS reproducible fixture; independent JDK oracle has %s lines\n' \
    "$(wc -l < "$scratch/jdk.out")"
run_bounded 'native Oak class' 16 "$MAGNOLIA" build run \
    --entry "$fixture" > "$scratch/native.out"
run_bounded 'isolated javac Java-input' 20 "$MAGNOLIA" build run \
    --entry test/fixtures/MultiArrayMain.java --java-backend javac \
    > "$scratch/javac-input.out"
run_bounded 'standalone Oak build' 35 "$MAGNOLIA" build --entry "$fixture" \
    --no-optimize --output "$scratch/MultiArrayMain.oak" >/dev/null
run_bounded 'standalone Oak execution' 15 "$MAGNOLIA" \
    "$scratch/MultiArrayMain.oak" > "$scratch/oak.out"
run_bounded 'Builder JS build' 75 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/MultiArrayMain.js" >/dev/null
run_bounded 'self-contained JS execution' 20 node "$scratch/MultiArrayMain.js" \
    > "$scratch/js.out"
for route in native javac-input oak js; do
    diff -u "$scratch/jdk.out" "$scratch/$route.out"
done
printf 'PASS original JDK = native/standalone Oak = isolated Java-input = JS\n'
run_bounded 'Builder HTML build' 110 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/MultiArrayMain.html" >/dev/null
run_bounded 'offline HTML inline JS oracle' 20 node - \
    "$scratch/MultiArrayMain.html" "$scratch/jdk.out" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2, 'two self-contained inline scripts required');
const output = {textContent: ''};
const canvas = {style: {}, getContext() { throw Error('unexpected draw'); }};
const context = {document: {getElementById(id) {
    return id === 'output' ? output : id === 'drawing' ? canvas : null;
}}, console: {log() {}, error(...args) {throw Error(args.join(' '));}}};
vm.createContext(context);
for (const script of scripts) vm.runInContext(script, context, {timeout: 12000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(context.MagnoliaApplet.inspect().type, 'ok');
JS
run_bounded 'generated Java source' 35 "$MAGNOLIA" build --entry "$fixture" \
    --output "$scratch/MultiArrayConverted.java" >/dev/null
run_bounded 'compile generated Java' 20 javac -d "$scratch/generated" \
    "$scratch/MultiArrayConverted.java"
run_bounded 'generated Java execution' 10 java -Xverify:all \
    -cp "$scratch/generated" MultiArrayConverted > "$scratch/generated.out"
diff -u "$scratch/jdk.out" "$scratch/generated.out"
run_bounded 'restore exact original class' 15 "$MAGNOLIA" build restore \
    --entry "$scratch/MultiArrayMain.oak" \
    --output "$scratch/original-restored.class" >/dev/null
cmp "$fixture" "$scratch/original-restored.class"
printf 'PASS inline HTML and generated Java agree with JDK; original bytes exact\n'

# A same-shape edit in combine([[J)J retains every 4-byte multianewarray
# operand, stack map and code offset. Recompile must scan past opcode 197.
python3 - "$scratch/MultiArrayMain.oak" "$scratch/MultiArrayEdited.oak" <<'PY'
from pathlib import Path
import sys
source = Path(sys.argv[1]).read_text(encoding='utf-8')
assert source.count('// MultiArrayMain.combine([[J)J') == 1
assert source.count('[131, 0, ?],') == 1
Path(sys.argv[2]).write_text(source.replace('[131, 0, ?],', '[129, 0, ?],'), encoding='utf-8')
PY
run_bounded 'checked Oak edit -> NEW class' 30 "$MAGNOLIA" build \
    --entry "$scratch/MultiArrayEdited.oak" \
    --output "$scratch/edited/MultiArrayMain.class" >/dev/null
if cmp -s "$fixture" "$scratch/edited/MultiArrayMain.class"; then
    echo 'FAIL edited class equals original' >&2; exit 1
fi
run_bounded 'JDK verifies edited class' 10 java -Xverify:all \
    -cp "$scratch/edited" MultiArrayMain > "$scratch/edited.jdk"
run_bounded 'edited Oak execution' 15 "$MAGNOLIA" \
    "$scratch/MultiArrayEdited.oak" > "$scratch/edited.oak"
diff -u "$scratch/edited.jdk" "$scratch/edited.oak"
python3 - "$scratch/jdk.out" "$scratch/edited.jdk" <<'PY'
from pathlib import Path
import sys
original, edited = [Path(name).read_text(encoding='utf-8').splitlines() for name in sys.argv[1:]]
assert len(original) == len(edited) == 22
assert sum(a != b for a,b in zip(original, edited)) == 1
assert original[4] != edited[4], 'only the expected long XOR result must change'
PY
run_bounded 'restore original despite edited Oak' 15 "$MAGNOLIA" build restore \
    --entry "$scratch/MultiArrayEdited.oak" \
    --output "$scratch/edited-original.class" >/dev/null
cmp "$fixture" "$scratch/edited-original.class"
printf 'PASS edited NEW class JDK verifies and differs in one value; original restores exactly\n'

python3 - "$fixture" "$scratch/MultiArrayMain.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[2], 'w', compression=zipfile.ZIP_DEFLATED) as jar:
    jar.writestr('META-INF/MANIFEST.MF',
                 'Manifest-Version: 1.0\r\nMain-Class: MultiArrayMain\r\n\r\n')
    jar.write(sys.argv[1], 'MultiArrayMain.class')
    jar.writestr('assets/multi.txt', 'preserve independent resource\n')
PY
run_bounded 'JDK original JAR' 10 java -Xverify:all \
    -jar "$scratch/MultiArrayMain.jar" > "$scratch/original.jar-jdk"
run_bounded 'native Oak original JAR' 20 "$MAGNOLIA" build run \
    --entry "$scratch/MultiArrayMain.jar" > "$scratch/original.jar-oak"
for route in original.jar-jdk original.jar-oak; do
    diff -u "$scratch/jdk.out" "$scratch/$route"
done
run_bounded 'JAR -> standalone Oak' 35 "$MAGNOLIA" build \
    --entry "$scratch/MultiArrayMain.jar" --no-optimize \
    --output "$scratch/MultiArrayJar.oak" >/dev/null
python3 - "$scratch/MultiArrayJar.oak" "$scratch/MultiArrayJarEdited.oak" <<'PY'
from pathlib import Path
import sys
source=Path(sys.argv[1]).read_text(encoding='utf-8')
assert source.count('[131, 0, ?],') == 1
Path(sys.argv[2]).write_text(source.replace('[131, 0, ?],', '[129, 0, ?],'), encoding='utf-8')
PY
run_bounded 'checked Oak edit -> NEW JAR' 35 "$MAGNOLIA" build \
    --entry "$scratch/MultiArrayJarEdited.oak" \
    --output "$scratch/MultiArrayEdited.jar" >/dev/null
run_bounded 'JDK verifies edited JAR' 10 java -Xverify:all \
    -jar "$scratch/MultiArrayEdited.jar" > "$scratch/edited.jar-jdk"
diff -u "$scratch/edited.jdk" "$scratch/edited.jar-jdk"
run_bounded 'restore exact original JAR' 15 "$MAGNOLIA" build restore \
    --entry "$scratch/MultiArrayJarEdited.oak" \
    --output "$scratch/MultiArrayRestored.jar" >/dev/null
cmp "$scratch/MultiArrayMain.jar" "$scratch/MultiArrayRestored.jar"
python3 - "$scratch/MultiArrayMain.jar" "$scratch/MultiArrayEdited.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[1]) as original, zipfile.ZipFile(sys.argv[2]) as edited:
    assert original.namelist() == edited.namelist()
    assert original.read('META-INF/MANIFEST.MF') == edited.read('META-INF/MANIFEST.MF')
    assert original.read('assets/multi.txt') == edited.read('assets/multi.txt')
    assert original.read('MultiArrayMain.class') != edited.read('MultiArrayMain.class')
PY
printf 'PASS edited multi-array JAR JDK verifies; manifest/resource kept; original byte-exact\n'

for kind in first-negative zero-then-negative oversized; do
    case "$kind" in
        first-negative) jdk_args=(one); jdk_error=NegativeArraySizeException; oak_error='invalid multianewarray descriptor or dimension';;
        zero-then-negative) jdk_args=(one two); jdk_error=NegativeArraySizeException; oak_error='invalid multianewarray descriptor or dimension';;
        oversized) jdk_args=(one two three); jdk_error=''; oak_error='multianewarray allocation budget exceeded';;
    esac
    args="$(python3 - "${jdk_args[@]}" <<'PY'
import json, sys
print(json.dumps(sys.argv[1:]))
PY
)"
    if [[ "$kind" == oversized ]]; then
        run_bounded 'JDK permits large array (Oak stays bounded)' 10 java \
            -Xverify:all -cp "$scratch/original" MultiArrayMain \
            "${jdk_args[@]}" > "$scratch/$kind.jdk"
        [[ $(tail -n 1 "$scratch/$kind.jdk") == 300 ]]
    else
        if timeout 10s java -Xverify:all -cp "$scratch/original" \
            MultiArrayMain "${jdk_args[@]}" > "$scratch/$kind.jdk" 2>&1; then
            echo "FAIL JDK $kind should throw" >&2; exit 1
        fi
        grep -q "$jdk_error" "$scratch/$kind.jdk"
    fi
    if timeout 15s "$MAGNOLIA" build run --entry "$fixture" \
        --args "$args" > "$scratch/$kind.oak" 2>&1; then
        echo "FAIL Oak should refuse $kind" >&2; exit 1
    fi
    grep -q "$oak_error" "$scratch/$kind.oak"
    printf 'PASS %-18s JDK result/exception checked; Oak rejects explicitly\n' "$kind"
done
run_bounded 'whole-input coverage' 15 "$MAGNOLIA" build coverage \
    --entry "$fixture" --json > "$scratch/coverage.json"
python3 - "$scratch/coverage.json" <<'PY'
import json, sys
coverage=json.load(open(sys.argv[1],encoding='utf-8'))
assert coverage['allBodiesPreflighted?'] and coverage['decodedBodies'] == coverage['methods'] == 4
assert not coverage['blockers']
PY
printf 'PASS all four fixture methods decode; multianewarray matrix complete\n'
