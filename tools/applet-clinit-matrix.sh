#!/usr/bin/env bash
# Abrupt <clinit> JVM oracle over a three-class app: one-time EIIE wrapping,
# Error propagation, subsequent NCDFE, causes, and all offline output routes.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for cmd in "$MAGNOLIA" java javac node python3 timeout diff cmp; do
    command -v "$cmd" >/dev/null 2>&1 || { echo "missing: $cmd" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/original" "$scratch/generated"
run() {
    local name="$1" limit="$2" start=$SECONDS rc
    shift 2
    printf '[clinit] START %s (max %ss)\n' "$name" "$limit" >&2
    if timeout --signal=TERM --kill-after=3s "$limit" "$@"; then
        printf '[clinit] PASS %s (%ss)\n' "$name" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[clinit] FAIL %s exit=%s (%ss)\n' "$name" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
run 'reproduce 3 Java 8 class images' 20 javac --release 8 -g:none -d "$scratch/original" test/fixtures/ExceptionInitMain.java
for name in ExceptionInitMain 'ExceptionInitMain$BrokenInit' 'ExceptionInitMain$BrokenError'; do
    cmp "test/fixtures/$name.class" "$scratch/original/$name.class"
done
run 'JDK verifier and 5-line class-initialization oracle' 12 java -Xverify:all -cp "$scratch/original" ExceptionInitMain >"$scratch/jdk.out"
diff -u test/fixtures/ExceptionInitMain.expected "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 5 ]]
run 'native Oak linked 3-class program' 20 "$MAGNOLIA" build run \
    --entry test/fixtures/ExceptionInitMain.class \
    'test/fixtures/ExceptionInitMain$BrokenInit.class' \
    'test/fixtures/ExceptionInitMain$BrokenError.class' \
    --class ExceptionInitMain >"$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run 'build portable 3-class archive' 25 "$MAGNOLIA" build build \
    --entry test/fixtures/ExceptionInitMain.class \
    'test/fixtures/ExceptionInitMain$BrokenInit.class' \
    'test/fixtures/ExceptionInitMain$BrokenError.class' \
    --class ExceptionInitMain --output "$scratch/ExceptionInitMain.applet.json" >/dev/null
archive="$scratch/ExceptionInitMain.applet.json"
run 'build standalone multi-class Oak' 45 "$MAGNOLIA" build --entry "$archive" \
    --no-optimize --output "$scratch/ExceptionInitMain.oak" >/dev/null
run 'execute standalone multi-class Oak' 18 "$MAGNOLIA" "$scratch/ExceptionInitMain.oak" >"$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
run 'build self-contained JS' 95 "$MAGNOLIA" build --entry "$archive" \
    --output "$scratch/ExceptionInitMain.js" >/dev/null
run 'execute JS' 25 node "$scratch/ExceptionInitMain.js" >"$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run 'build offline HTML' 110 "$MAGNOLIA" build --entry "$archive" \
    --output "$scratch/ExceptionInitMain.html" >/dev/null
run 'execute offline inline HTML' 20 node - "$scratch/ExceptionInitMain.html" "$scratch/jdk.out" <<'JS'
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
run 'build generated Java' 45 "$MAGNOLIA" build --entry "$archive" \
    --output "$scratch/ExceptionInitConverted.java" >/dev/null
run 'compile generated Java' 20 javac -d "$scratch/generated" "$scratch/ExceptionInitConverted.java"
run 'execute generated Java' 12 java -Xverify:all -cp "$scratch/generated" ExceptionInitConverted >"$scratch/java.out"
diff -u "$scratch/jdk.out" "$scratch/java.out"
run 'restore original 3 classes from Oak' 20 "$MAGNOLIA" build restore \
    --entry "$scratch/ExceptionInitMain.oak" --output "$scratch/restored" >/dev/null
for name in ExceptionInitMain 'ExceptionInitMain$BrokenInit' 'ExceptionInitMain$BrokenError'; do
    cmp "test/fixtures/$name.class" "$scratch/restored/$name.class"
done
printf 'PASS 5 class-init JDK lines = native/Oak/JS/HTML/Java; 3 original class images restored byte-exactly\n'
