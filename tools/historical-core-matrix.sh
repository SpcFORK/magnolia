#!/usr/bin/env bash
# Differential, genuine-bytecode oracle for two historical JVM 45.3 class sets.
# No third-party binaries are committed; all build products stay in mktemp.
# Offline: FLASH_ZIP=/path/to/FlashApi.ZIP FLIPLAY_CLASS=/path/to/fliplay.class \
#          FLICKFRAME_CLASS=/path/to/flickframe.class \
#          MAGNOLIA=/path/to/magnolia bash tools/historical-core-matrix.sh
# Online:  MAGNOLIA=/path/to/magnolia bash tools/historical-core-matrix.sh --download
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
JAVAC="${JAVAC:-javac}"
JAVA="${JAVA:-java}"
for tool in "$MAGNOLIA" "$JAVA" "$JAVAC" node python3 sha256sum cmp diff grep; do
    if ! command -v "$tool" >/dev/null 2>&1; then
        echo "missing historical-core-matrix dependency: $tool" >&2
        exit 2
    fi
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/fliplay" "$scratch/flash/com/macromedia/flash" "$scratch/converted"

if [[ "${1:-}" == '--download' && "$#" -eq 1 ]]; then
    command -v curl >/dev/null || { echo 'missing curl' >&2; exit 2; }
    # Download only for an explicitly requested network run. Hashes below
    # identify exactly which class bytes are tested; do not redistribute them.
    curl --fail --location --silent --show-error --max-time 45 \
        --user-agent 'Mozilla/5.0' -H 'Accept: */*' \
        --output "$scratch/fliplay/fliplay.class" \
        'https://www.oortman3d.com/fliplay/fliplay.class'
    curl --fail --location --silent --show-error --max-time 45 \
        --user-agent 'Mozilla/5.0' -H 'Accept: */*' \
        --output "$scratch/fliplay/flickframe.class" \
        'https://www.oortman3d.com/fliplay/flickframe.class'
    curl --fail --location --silent --show-error --max-time 90 \
        --output "$scratch/FlashApi.ZIP" \
        'https://raw.githubusercontent.com/macromedia-os/Macromedia-Flash-SDK/main/JavaFlash/FlashApi.ZIP'
elif [[ "$#" -eq 0 ]]; then
    : "${FLIPLAY_CLASS:?set FLIPLAY_CLASS to the original fliplay.class or use --download}"
    : "${FLICKFRAME_CLASS:?set FLICKFRAME_CLASS to the original flickframe.class or use --download}"
    : "${FLASH_ZIP:?set FLASH_ZIP to the original FlashApi.ZIP or use --download}"
    cp "$FLIPLAY_CLASS" "$scratch/fliplay/fliplay.class"
    cp "$FLICKFRAME_CLASS" "$scratch/fliplay/flickframe.class"
    cp "$FLASH_ZIP" "$scratch/FlashApi.ZIP"
else
    echo 'usage: bash tools/historical-core-matrix.sh [--download]' >&2
    exit 2
fi

check_hash() {
    local actual
    actual="$(sha256sum "$1")"
    actual="${actual%% *}"
    if [[ "$actual" != "$2" ]]; then
        echo "unexpected SHA-256 for $1: $actual (expected $2)" >&2
        exit 1
    fi
}
check_hash "$scratch/fliplay/fliplay.class" \
    bc315872dafebbf3c634e64c0fc251b38031ed49b7c4698d39c94990de9fe20b
check_hash "$scratch/fliplay/flickframe.class" \
    12feb5321cd8c9aeb0bf285f055596d981e3d45ccfb8c62d029f2f6d95ccdf9f
check_hash "$scratch/FlashApi.ZIP" \
    fbb2149179c1c5a32bd9aa7f479a1e760e36e18dbd8e48147729ca67822932c2

python3 - "$scratch/FlashApi.ZIP" "$scratch/flash/com/macromedia/flash" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[1]) as archive:
    for cls in ('Matrix', 'SParser', 'Rect'):
        member = archive.getinfo('com/macromedia/flash/' + cls + '.class')
        if member.file_size > 1048576:
            raise SystemExit('unexpected ' + cls + '.class size')
        with open(sys.argv[2] + '/' + cls + '.class', 'wb') as output:
            output.write(archive.read(member))
PY
check_hash "$scratch/flash/com/macromedia/flash/Matrix.class" \
    df934e7dce8756e79abe3360d1cea5a29818f4ed25f813ad0d1739abe8fca6c3
check_hash "$scratch/flash/com/macromedia/flash/SParser.class" \
    9fc21c5ec64e9daa32f163588f95d5ad18aae6cf1f73288c2932e11d882338db
check_hash "$scratch/flash/com/macromedia/flash/Rect.class" \
    36e7b476cec00972cc31929a7dfb32d7ba5abb729e4356288dec8e7487110f74
printf 'PASS original FlashApi.ZIP and Fliplay JVM 45.3 class hashes verified\n'

"$JAVAC" --release 8 -g:none -cp "$scratch/fliplay" -d "$scratch/fliplay" \
    test/historical/FliplayCoreMain.java
"$JAVAC" --release 8 -g:none -cp "$scratch/flash" -d "$scratch/flash" \
    test/historical/com/macromedia/flash/FlashMatrixCoreMain.java \
    test/historical/com/macromedia/flash/FlashParserCoreMain.java \
    test/historical/com/macromedia/flash/FlashPointRectCoreMain.java
printf '255\n128\n33023\n13440\n305430783\n-1\n' > "$scratch/FliplayCoreMain.expected"
printf '23\n-2147483648\n2147483647\n11\n63\n0\n' > "$scratch/FlashMatrixCoreMain.expected"
printf '33023\n-1\n-2147478988\n8\n' > "$scratch/FlashParserCoreMain.expected"
printf 'true\nfalse\ntrue\nfalse\ntrue\ntrue\nfalse\ntrue\nfalse\ntrue\nfalse\n' \
    > "$scratch/FlashPointRectCoreMain.expected"

# A reference-JVM run actually loads the original classes; it does not read
# Magnolia output. A successful checksum alone never counts as execution.
"$JAVA" -Xverify:all -Djava.awt.headless=false -Dfile.encoding=UTF-8 \
    -cp "$scratch/fliplay" FliplayCoreMain > "$scratch/FliplayCoreMain.reference"
"$JAVA" -Xverify:all -Dfile.encoding=UTF-8 \
    -cp "$scratch/flash" com.macromedia.flash.FlashMatrixCoreMain \
    > "$scratch/FlashMatrixCoreMain.reference"
"$JAVA" -Xverify:all -Dfile.encoding=UTF-8 \
    -cp "$scratch/flash" com.macromedia.flash.FlashParserCoreMain \
    > "$scratch/FlashParserCoreMain.reference"
"$JAVA" -Xverify:all -Dfile.encoding=UTF-8 \
    -cp "$scratch/flash" com.macromedia.flash.FlashPointRectCoreMain \
    > "$scratch/FlashPointRectCoreMain.reference"
for name in FliplayCoreMain FlashMatrixCoreMain FlashParserCoreMain FlashPointRectCoreMain; do
    diff -u "$scratch/$name.expected" "$scratch/$name.reference"
done
printf 'PASS independent JDK -Xverify:all executed original constructors, parser, Matrix, and Point/Rect geometry\n'

run_routes() {
    local name="$1" entry="$2" owner="$3" helpers="$4"
    "$MAGNOLIA" build run --entry "$entry" --class "$owner" \
        --include "$helpers" > "$scratch/$name.native"
    diff -u "$scratch/$name.reference" "$scratch/$name.native"
    printf 'PASS %-20s native Magnolia executes original class methods\n' "$name"

    "$MAGNOLIA" build --entry "$entry" --entry-class "$owner" \
        --class-files "$helpers" --output "$scratch/$name.oak" >/dev/null
    "$MAGNOLIA" "$scratch/$name.oak" > "$scratch/$name.oak-output"
    diff -u "$scratch/$name.reference" "$scratch/$name.oak-output"

    "$MAGNOLIA" build --entry "$entry" --entry-class "$owner" \
        --class-files "$helpers" --output "$scratch/$name.js" >/dev/null
    node "$scratch/$name.js" > "$scratch/$name.js-output"
    diff -u "$scratch/$name.reference" "$scratch/$name.js-output"

    "$MAGNOLIA" build --entry "$entry" --entry-class "$owner" \
        --class-files "$helpers" --output "$scratch/${name}Converted.java" >/dev/null
    "$JAVAC" -d "$scratch/converted" "$scratch/${name}Converted.java"
    "$JAVA" -Xverify:all -Dfile.encoding=UTF-8 -cp "$scratch/converted" \
        "${name}Converted" > "$scratch/$name.java-output"
    diff -u "$scratch/$name.reference" "$scratch/$name.java-output"

    "$MAGNOLIA" build restore --entry "$scratch/$name.oak" \
        --output "$scratch/$name.restored" >/dev/null
    cmp "$entry" "$scratch/$name.restored/$owner.class"
    printf 'PASS %-20s reference JVM = standalone Oak = Node JS = generated Java; exact harness recovery\n' "$name"
}
run_routes FliplayCoreMain \
    "$scratch/fliplay/FliplayCoreMain.class" FliplayCoreMain \
    "$scratch/fliplay/fliplay.class,$scratch/fliplay/flickframe.class"
cmp "$scratch/fliplay/fliplay.class" "$scratch/FliplayCoreMain.restored/fliplay.class"
cmp "$scratch/fliplay/flickframe.class" "$scratch/FliplayCoreMain.restored/flickframe.class"
run_routes FlashMatrixCoreMain \
    "$scratch/flash/com/macromedia/flash/FlashMatrixCoreMain.class" \
    com/macromedia/flash/FlashMatrixCoreMain \
    "$scratch/flash/com/macromedia/flash/Matrix.class"
cmp "$scratch/flash/com/macromedia/flash/Matrix.class" \
    "$scratch/FlashMatrixCoreMain.restored/com/macromedia/flash/Matrix.class"
run_routes FlashParserCoreMain \
    "$scratch/flash/com/macromedia/flash/FlashParserCoreMain.class" \
    com/macromedia/flash/FlashParserCoreMain \
    "$scratch/flash/com/macromedia/flash/SParser.class"
cmp "$scratch/flash/com/macromedia/flash/SParser.class" \
    "$scratch/FlashParserCoreMain.restored/com/macromedia/flash/SParser.class"
run_routes FlashPointRectCoreMain \
    "$scratch/flash/com/macromedia/flash/FlashPointRectCoreMain.class" \
    com/macromedia/flash/FlashPointRectCoreMain \
    "$scratch/flash/com/macromedia/flash/Rect.class"
cmp "$scratch/flash/com/macromedia/flash/Rect.class" \
    "$scratch/FlashPointRectCoreMain.restored/com/macromedia/flash/Rect.class"
for name in FlashParserCoreMain FlashPointRectCoreMain; do
    helper=Rect
    if [[ "$name" == FlashParserCoreMain ]]; then helper=SParser; fi
    "$MAGNOLIA" build --entry "$scratch/flash/com/macromedia/flash/$name.class" \
        --entry-class "com/macromedia/flash/$name" \
        --class-files "$scratch/flash/com/macromedia/flash/$helper.class" \
        --output "$scratch/$name.html" >/dev/null
    node - "$scratch/$name.html" "$scratch/$name.reference" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2, 'HTML must contain two inline scripts, no external scripts');
const output = {textContent: ''};
const canvas = {style: {}, getContext() {throw Error('unexpected graphics usage')}};
const context = {
    document: {getElementById(id) {
        return id === 'output' ? output : id === 'drawing' ? canvas : null;
    }},
    console: {log() {}, error(...messages) {throw Error(messages.join(' '))}},
};
vm.createContext(context);
for (const code of scripts) vm.runInContext(code, context, {timeout: 10000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(context.MagnoliaApplet.inspect().type, 'ok');
JS
    printf 'PASS original Flash %s HTML has two inline scripts and DOM output = JDK oracle\n' "$helper"
done
python3 - "$scratch/FlashMatrixCoreMainConverted.java" \
    "$scratch/FlashParserCoreMainConverted.java" <<'PY'
import sys
matrix = open(sys.argv[1], encoding='ascii').read()
parser = open(sys.argv[2], encoding='ascii').read()
for source, methods in ((matrix, (('\\u0177', '(I)I'), ('\\u0178', '(II)I'))),
                        (parser, (('\\u00c1', '()I'), ('\\u00c2', '()I'), ('\\u00c3', '()I')))):
    for name, signature in methods:
        if source.count(name) < 2 or name + signature not in source:
            raise SystemExit('Java target collapsed an original Flash method name: ' + name)
    if '\\u0000(I)I' in source or '\\u0000()I' in source:
        raise SystemExit('Java target silently replaced an original Flash method name')
PY
printf 'PASS Java target preserved original Flash method identifiers as distinct Unicode escapes\n'

# This is the DIFFERENT route: statically assemble one permitted same-width
# instruction edit in an authentic parser method. JDK verifier + changed
# behavior check the NEW class; restore must still return the ORIGINAL bytes.
"$MAGNOLIA" build --entry "$scratch/flash/com/macromedia/flash/FlashParserCoreMain.class" \
    --entry-class com/macromedia/flash/FlashParserCoreMain \
    --class-files "$scratch/flash/com/macromedia/flash/SParser.class" \
    --no-optimize --output "$scratch/FlashParserEditable.oak" >/dev/null
python3 - "$scratch/FlashParserEditable.oak" "$scratch/FlashParserEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
start = source.index('// com/macromedia/flash/SParser.Á()I\n')
end = source.index('\n// com/macromedia/flash/SParser.', start + 1)
method = source[start:end]
old, new = '    [17, 255, ?],', '    [17, 127, ?],'
if method.count(old) != 2:
    raise SystemExit('authentic parser method no longer has two unsigned-byte masks')
with open(sys.argv[2], 'w', encoding='utf-8') as output:
    output.write(source[:start] + method.replace(old, new, 1) + source[end:])
PY
"$MAGNOLIA" build recompile --entry "$scratch/FlashParserEdited.oak" \
    --output "$scratch/FlashParserEdited.classes" > "$scratch/recompiled.log"
grep -q '(1 changed instructions)' "$scratch/recompiled.log"
if cmp -s "$scratch/flash/com/macromedia/flash/SParser.class" \
    "$scratch/FlashParserEdited.classes/com/macromedia/flash/SParser.class"; then
    echo 'FAIL: new edited parser class is byte-identical to the original' >&2
    exit 1
fi
"$JAVA" -Xverify:all -Dfile.encoding=UTF-8 \
    -cp "$scratch/FlashParserEdited.classes" com.macromedia.flash.FlashParserCoreMain \
    > "$scratch/FlashParserEdited.reference"
"$MAGNOLIA" "$scratch/FlashParserEdited.oak" > "$scratch/FlashParserEdited.oak-output"
printf '32895\n-129\n-2147478988\n8\n' > "$scratch/FlashParserEdited.expected"
diff -u "$scratch/FlashParserEdited.expected" "$scratch/FlashParserEdited.reference"
diff -u "$scratch/FlashParserEdited.reference" "$scratch/FlashParserEdited.oak-output"
"$MAGNOLIA" build restore --entry "$scratch/FlashParserEdited.oak" \
    --output "$scratch/FlashParserOriginal.classes" >/dev/null
cmp "$scratch/flash/com/macromedia/flash/SParser.class" \
    "$scratch/FlashParserOriginal.classes/com/macromedia/flash/SParser.class"
printf 'PASS edited original Flash parser Oak -> NEW JDK-verified class matches edited Oak; original stays byte-exact\n'

# Change one SAME-WIDTH conditional in the original Rect.Ɓ(Point) method:
# reject a point on the minimum-x boundary. The new class must verify and
# behave differently on a real JVM, while restore returns original bytes.
"$MAGNOLIA" build --entry "$scratch/flash/com/macromedia/flash/FlashPointRectCoreMain.class" \
    --entry-class com/macromedia/flash/FlashPointRectCoreMain \
    --class-files "$scratch/flash/com/macromedia/flash/Rect.class" \
    --no-optimize --output "$scratch/FlashPointRectEditable.oak" >/dev/null
python3 - "$scratch/FlashPointRectEditable.oak" "$scratch/FlashPointRectEdited.oak" \
    "$scratch/FlashPointRectCoreMain.reference" "$scratch/FlashPointRectEdited.expected" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
start = source.index('// com/macromedia/flash/Rect.Ɓ(Ljava/awt/Point;)Z\n')
end = source.index('\n]\n', start) + 3
method = source[start:end]
old, new = '    [163, 20, ?],', '    [162, 20, ?],'
if method.count(old) != 3:
    raise SystemExit('original Rect must have three guarded outside branches')
with open(sys.argv[2], 'w', encoding='utf-8') as edited:
    edited.write(source[:start] + method.replace(old, new, 1) + source[end:])
lines = open(sys.argv[3], encoding='utf-8').readlines()
if len(lines) != 11 or lines[2] != 'true\n' or lines[4] != 'true\n':
    raise SystemExit('unexpected JDK Rect baseline; review selected boundary edit')
lines[2] = lines[4] = 'false\n'
with open(sys.argv[4], 'w', encoding='utf-8') as expected:
    expected.writelines(lines)
PY
"$MAGNOLIA" build recompile --entry "$scratch/FlashPointRectEdited.oak" \
    --output "$scratch/FlashPointRectEdited.classes" > "$scratch/point-recompiled.log"
grep -q '(1 changed instructions)' "$scratch/point-recompiled.log"
if cmp -s "$scratch/flash/com/macromedia/flash/Rect.class" \
    "$scratch/FlashPointRectEdited.classes/com/macromedia/flash/Rect.class"; then
    echo 'FAIL: edited original Rect class is byte-identical to original' >&2
    exit 1
fi
"$JAVA" -Xverify:all -Dfile.encoding=UTF-8 \
    -cp "$scratch/FlashPointRectEdited.classes" com.macromedia.flash.FlashPointRectCoreMain \
    > "$scratch/FlashPointRectEdited.reference"
"$MAGNOLIA" "$scratch/FlashPointRectEdited.oak" > "$scratch/FlashPointRectEdited.oak-output"
diff -u "$scratch/FlashPointRectEdited.expected" "$scratch/FlashPointRectEdited.reference"
diff -u "$scratch/FlashPointRectEdited.reference" "$scratch/FlashPointRectEdited.oak-output"
"$MAGNOLIA" build --entry "$scratch/FlashPointRectEdited.classes/com/macromedia/flash/FlashPointRectCoreMain.class" \
    --entry-class com/macromedia/flash/FlashPointRectCoreMain \
    --class-files "$scratch/FlashPointRectEdited.classes/com/macromedia/flash/Rect.class" \
    --output "$scratch/FlashPointRectEdited.js" >/dev/null
node "$scratch/FlashPointRectEdited.js" > "$scratch/FlashPointRectEdited.js-output"
diff -u "$scratch/FlashPointRectEdited.reference" "$scratch/FlashPointRectEdited.js-output"
"$MAGNOLIA" build --entry "$scratch/FlashPointRectEdited.classes/com/macromedia/flash/FlashPointRectCoreMain.class" \
    --entry-class com/macromedia/flash/FlashPointRectCoreMain \
    --class-files "$scratch/FlashPointRectEdited.classes/com/macromedia/flash/Rect.class" \
    --output "$scratch/FlashPointRectEditedConverted.java" >/dev/null
"$JAVAC" -d "$scratch/converted" "$scratch/FlashPointRectEditedConverted.java"
"$JAVA" -Xverify:all -cp "$scratch/converted" FlashPointRectEditedConverted \
    > "$scratch/FlashPointRectEdited.java-output"
diff -u "$scratch/FlashPointRectEdited.reference" "$scratch/FlashPointRectEdited.java-output"
"$MAGNOLIA" build restore --entry "$scratch/FlashPointRectEdited.oak" \
    --output "$scratch/FlashPointRectOriginal.classes" >/dev/null
cmp "$scratch/flash/com/macromedia/flash/Rect.class" \
    "$scratch/FlashPointRectOriginal.classes/com/macromedia/flash/Rect.class"
cmp "$scratch/flash/com/macromedia/flash/FlashPointRectCoreMain.class" \
    "$scratch/FlashPointRectOriginal.classes/com/macromedia/flash/FlashPointRectCoreMain.class"
printf 'PASS edited original Rect Oak -> new JDK-verified class; Oak/JS/Java match; original bytes restored\n'

# Full-player coverage is deliberately separate: four *selected* slices
# pass, whereas the player class sets still contain unsupported methods/APIs.
if "$MAGNOLIA" build coverage \
    --entry "$scratch/fliplay/fliplay.class" \
    --include "$scratch/fliplay/flickframe.class" --json > "$scratch/fliplay.coverage"; then
    echo 'unexpectedly unblocked full Fliplay coverage; re-audit player before claiming it runs' >&2
    exit 1
fi
if "$MAGNOLIA" build coverage --entry "$scratch/FlashApi.ZIP" \
    --json > "$scratch/flash.coverage"; then
    echo 'unexpectedly unblocked full Flash coverage; re-audit player before claiming it runs' >&2
    exit 1
fi
python3 - "$scratch/fliplay.coverage" "$scratch/flash.coverage" <<'PY'
import json, sys
for label, filename in zip(('Fliplay', 'Flash'), sys.argv[1:]):
    report = json.load(open(filename, encoding='utf-8'))
    assert report['type'] == 'ok' and not report['allBodiesPreflighted?']
    assert report['decodedBodies'] < report['methods'] and report['blockers']
    print('LIMIT %s: %d/%d decoded method bodies; %d blockers; full player NOT RUNNABLE' %
          (label, report['decodedBodies'], report['methods'], len(report['blockers'])))
PY
