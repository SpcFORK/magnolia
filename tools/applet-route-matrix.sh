#!/usr/bin/env bash
# Cross-route differential oracle for the bounded JVM fixture subset.
# Run from anywhere: MAGNOLIA=/path/to/magnolia bash tools/applet-route-matrix.sh
# Requires a JDK, Node.js, Python 3 and the Magnolia interpreter; no network.
# --tail reruns HTML/edit/coverage checks independently after a timed full run.
# External callers may wrap each Magnolia invocation with timeout and tracing.
set -euo pipefail
cd "$(dirname "$0")/.."
if [[ $# -gt 1 || ($# -eq 1 && "$1" != '--tail') ]]; then
    echo 'usage: applet-route-matrix.sh [--tail]' >&2
    exit 2
fi
tail_only=false
if [[ ${1:-} == '--tail' ]]; then tail_only=true; fi
MAGNOLIA="${MAGNOLIA:-magnolia}"
for command in "$MAGNOLIA" java javac node python3; do
    if ! command -v "$command" >/dev/null 2>&1; then
        echo "missing route-matrix dependency: $command" >&2
        exit 2
    fi
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/classes" "$scratch/edited"

if [[ $tail_only == false ]]; then
for name in NarrowIntMain PrimitiveArrayMain StackShuffleMain SwitchMain WideLocalMain ReferenceIdentityMain PointGeometryMain ArrayCopyMain StackHeightMain LongWidthMain; do
    printf '[route] checking %s: JDK, native, Oak, JS, Java, restore\n' "$name" >&2
    source="test/fixtures/$name.class"
    arguments=()
    jdk_arguments=()
    if [[ "$name" == ReferenceIdentityMain ]]; then
        arguments=(--args '["same","same"]')
        jdk_arguments=(same same)
    fi
    java -Xverify:all -cp test/fixtures "$name" "${jdk_arguments[@]}" > "$scratch/$name.reference"

    "$MAGNOLIA" build run --entry "$source" "${arguments[@]}" > "$scratch/$name.native"
    "$MAGNOLIA" build --entry "$source" "${arguments[@]}" --no-optimize \
        --output "$scratch/$name.oak" >/dev/null
    "$MAGNOLIA" "$scratch/$name.oak" > "$scratch/$name.oak-output"
    "$MAGNOLIA" build --entry "$source" "${arguments[@]}" --output "$scratch/$name.js" >/dev/null
    node "$scratch/$name.js" > "$scratch/$name.js-output"
    "$MAGNOLIA" build --entry "$source" "${arguments[@]}" \
        --output "$scratch/${name}Converted.java" >/dev/null
    javac -d "$scratch/classes" "$scratch/${name}Converted.java"
    java -Xverify:all -cp "$scratch/classes" "${name}Converted" \
        > "$scratch/$name.java-output"
    for route in native oak-output js-output java-output; do
        diff -u "$scratch/$name.reference" "$scratch/$name.$route"
    done
    "$MAGNOLIA" build --entry "$scratch/$name.oak" \
        --output "$scratch/$name.class" >/dev/null
    cmp "$source" "$scratch/$name.class"
    printf 'PASS %-19s JVM = native Oak = standalone Oak = JS = Java; exact class restore\n' "$name"
done

# A two-slot ladd -> lsub edit changes one observed result, produces a real
# new JDK-verified class, and never confuses restoration with recompilation.
python3 - "$scratch/LongWidthMain.oak" "$scratch/LongWidthMainEdited.oak" \
    "$scratch/LongWidthMainBadWidth.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
head, body = source.split('// LongWidthMain.bump(J)J\n', 1)
method, rest = body.split('// LongWidthMain.main(', 1)
assert method.count('[97, 0, ?],') == 1
with open(sys.argv[2], 'w', encoding='utf-8') as out:
    out.write(head + '// LongWidthMain.bump(J)J\n' +
              method.replace('[97, 0, ?],', '[101, 0, ?],') +
              '// LongWidthMain.main(' + rest)
old = "'bump(J)J': {locals: 2, stack: 4, argWidths: [2]"
assert source.count(old) == 1
with open(sys.argv[3], 'w', encoding='utf-8') as out:
    out.write(source.replace(old, old.replace('[2]', '[1]')))
PY
"$MAGNOLIA" build --entry "$scratch/LongWidthMainEdited.oak" \
    --output "$scratch/edited/LongWidthMain.class" >/dev/null
if cmp -s test/fixtures/LongWidthMain.class "$scratch/edited/LongWidthMain.class"; then
    echo 'FAIL: edited long class is byte-identical to original' >&2
    exit 1
fi
java -Xverify:all -cp "$scratch/edited" LongWidthMain \
    > "$scratch/LongWidthMainEdited.reference"
"$MAGNOLIA" "$scratch/LongWidthMainEdited.oak" > "$scratch/LongWidthMainEdited.oak-output"
diff -u "$scratch/LongWidthMainEdited.reference" \
    "$scratch/LongWidthMainEdited.oak-output"
python3 - "$scratch/LongWidthMain.reference" "$scratch/LongWidthMainEdited.reference" <<'PY'
import sys
original, edited = [open(name, encoding='utf-8').read().splitlines()
                    for name in sys.argv[1:]]
assert len(original) == len(edited) == 13
assert original[2] == '1' and edited[2] == '-1'
assert original[:2] + original[3:] == edited[:2] + edited[3:]
PY
"$MAGNOLIA" build restore --entry "$scratch/LongWidthMainEdited.oak" \
    --output "$scratch/LongWidthMain.original.class" >/dev/null
cmp test/fixtures/LongWidthMain.class "$scratch/LongWidthMain.original.class"
if "$MAGNOLIA" build --entry "$scratch/LongWidthMainBadWidth.oak" \
    --output "$scratch/LongWidthMainBadWidth.class" > "$scratch/bad-long-width.log" 2>&1; then
    echo 'FAIL: altered two-slot method signature metadata was accepted' >&2
    exit 1
fi
grep -q 'argument slots changed in Oak' "$scratch/bad-long-width.log"
printf 'PASS checked long Oak edit = new JDK class; byte-exact original restore; forged argument slots refused\n'

# The same category-two method also edits inside an actual JAR. Preserve its
# manifest and nonclass resource while producing a different verified class.
python3 - test/fixtures/LongWidthMain.class "$scratch/LongWidthMain.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[2], 'w', compression=zipfile.ZIP_DEFLATED) as jar:
    jar.writestr('META-INF/MANIFEST.MF',
                 'Manifest-Version: 1.0\r\nMain-Class: LongWidthMain\r\n\r\n')
    jar.write(sys.argv[1], 'LongWidthMain.class')
    jar.writestr('assets/long.txt', 'untouched resource\n')
PY
java -Xverify:all -jar "$scratch/LongWidthMain.jar" > "$scratch/LongWidthMain.jar-reference"
"$MAGNOLIA" build run --entry "$scratch/LongWidthMain.jar" \
    > "$scratch/LongWidthMain.jar-oak"
diff -u "$scratch/LongWidthMain.reference" "$scratch/LongWidthMain.jar-reference"
diff -u "$scratch/LongWidthMain.reference" "$scratch/LongWidthMain.jar-oak"
"$MAGNOLIA" build --entry "$scratch/LongWidthMain.jar" \
    --output "$scratch/LongWidthJar.oak" >/dev/null
python3 - "$scratch/LongWidthJar.oak" "$scratch/LongWidthJarEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
head, body = source.split('// LongWidthMain.bump(J)J\n', 1)
method, rest = body.split('// LongWidthMain.main(', 1)
assert method.count('[97, 0, ?],') == 1
with open(sys.argv[2], 'w', encoding='utf-8') as out:
    out.write(head + '// LongWidthMain.bump(J)J\n' +
              method.replace('[97, 0, ?],', '[101, 0, ?],') +
              '// LongWidthMain.main(' + rest)
PY
"$MAGNOLIA" build --entry "$scratch/LongWidthJarEdited.oak" \
    --output "$scratch/LongWidthMainEdited.jar" >/dev/null
java -Xverify:all -jar "$scratch/LongWidthMainEdited.jar" \
    > "$scratch/LongWidthMainEdited.jar-reference"
diff -u "$scratch/LongWidthMainEdited.reference" \
    "$scratch/LongWidthMainEdited.jar-reference"
"$MAGNOLIA" build restore --entry "$scratch/LongWidthJarEdited.oak" \
    --output "$scratch/LongWidthMainRestored.jar" >/dev/null
cmp "$scratch/LongWidthMain.jar" "$scratch/LongWidthMainRestored.jar"
python3 - "$scratch/LongWidthMain.jar" "$scratch/LongWidthMainEdited.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[1]) as original, zipfile.ZipFile(sys.argv[2]) as edited:
    assert original.namelist() == edited.namelist()
    for name in ('META-INF/MANIFEST.MF', 'assets/long.txt'):
        assert original.read(name) == edited.read(name)
    assert original.read('LongWidthMain.class') != edited.read('LongWidthMain.class')
PY
printf 'PASS exact long JAR round trip and edited long JAR verify on JDK; resources preserved\n'

# Both Java-input backends share the linked VM; the Oak-only frontend must
# refuse unsupported reference comparison instead of emitting integer opcodes.
"$MAGNOLIA" build run --entry test/fixtures/ReferenceIdentityMain.java \
    --java-backend javac --args '["same","same"]' > "$scratch/ReferenceIdentityMain.javac-input"
diff -u "$scratch/ReferenceIdentityMain.reference" \
    "$scratch/ReferenceIdentityMain.javac-input"
"$MAGNOLIA" build run --entry test/fixtures/PointGeometryMain.java \
    --java-backend javac > "$scratch/PointGeometryMain.javac-input"
diff -u "$scratch/PointGeometryMain.reference" \
    "$scratch/PointGeometryMain.javac-input"
"$MAGNOLIA" build run --entry test/fixtures/ArrayCopyMain.java \
    --java-backend javac > "$scratch/ArrayCopyMain.javac-input"
diff -u "$scratch/ArrayCopyMain.reference" \
    "$scratch/ArrayCopyMain.javac-input"
"$MAGNOLIA" build run --entry test/fixtures/LongWidthMain.java \
    --java-backend javac > "$scratch/LongWidthMain.javac-input"
diff -u "$scratch/LongWidthMain.reference" \
    "$scratch/LongWidthMain.javac-input"
cat > "$scratch/IdentitySubsetReject.java" <<'JAVA'
public class IdentitySubsetReject {
    public static void main(String[] args) {
        String left = args[0];
        String right = args[1];
        if (left == right) { System.out.println(1); }
    }
}
JAVA
if "$MAGNOLIA" build run --entry "$scratch/IdentitySubsetReject.java" \
    --java-backend oak --args '["same","same"]' > "$scratch/IdentitySubsetReject.log" 2>&1; then
    echo 'FAIL: Oak-only Java frontend accepted reference comparisons without modeling them' >&2
    exit 1
fi
grep -q 'comparison requires operands of the same int/boolean type' \
    "$scratch/IdentitySubsetReject.log"
printf 'PASS isolated javac inputs execute identity, Point, arraycopy and exact long; Oak-only Java rejects out-of-subset comparisons\n'

# JDK exceptions on invalid arraycopy inputs are currently reported as
# explicit bounded VM errors, not catchable exceptions. Never treat a failed
# copy, incompatible primitive array or partial covariant copy as success.
for cause in range negative null type store; do
    if "$MAGNOLIA" build run --entry test/fixtures/ArrayCopyMain.class \
        --args "[\"$cause\"]" > "$scratch/arraycopy-$cause.log" 2>&1; then
        echo "FAIL: invalid System.arraycopy ($cause) unexpectedly succeeded" >&2
        exit 1
    fi
    case "$cause" in
        range|negative) expected='System.arraycopy range out of bounds' ;;
        null) expected='null System.arraycopy array' ;;
        type) expected='System.arraycopy primitive array type mismatch' ;;
        store) expected='System.arraycopy array store type mismatch' ;;
    esac
    grep -q "$expected" "$scratch/arraycopy-$cause.log"
done
printf 'PASS five invalid arraycopy paths fail with explicit diagnostics\n'
"$MAGNOLIA" build --entry test/fixtures/ArrayCopyMain.class \
    --args '["store"]' --output "$scratch/ArrayCopyInvalid.js" >/dev/null
if node "$scratch/ArrayCopyInvalid.js" > "$scratch/ArrayCopyInvalid.js.log" 2>&1; then
    echo 'FAIL: JS arraycopy ignored an incompatible covariant element' >&2
    exit 1
fi
grep -q 'System.arraycopy array store type mismatch' "$scratch/ArrayCopyInvalid.js.log"
"$MAGNOLIA" build --entry test/fixtures/ArrayCopyMain.class \
    --args '["store"]' --output "$scratch/ArrayCopyInvalidConverted.java" >/dev/null
javac -d "$scratch/classes" "$scratch/ArrayCopyInvalidConverted.java"
if java -Xverify:all -cp "$scratch/classes" ArrayCopyInvalidConverted \
    > "$scratch/ArrayCopyInvalid.java.log" 2>&1; then
    echo 'FAIL: generated Java arraycopy ignored an incompatible element' >&2
    exit 1
fi
grep -q 'System.arraycopy array store type mismatch' "$scratch/ArrayCopyInvalid.java.log"
printf 'PASS JS and generated Java also reject incompatible reference stores\n'
else
    # The full run has already exercised these class routes. Recreate only
    # the JDK outputs and editable Oak listings needed by the remaining checks.
    printf '[route] --tail: regenerating four JDK oracles and three Oak listings\n' >&2
    for name in ReferenceIdentityMain PointGeometryMain ArrayCopyMain LongWidthMain; do
        jdk_arguments=()
        if [[ "$name" == ReferenceIdentityMain ]]; then
            jdk_arguments=(same same)
        fi
        java -Xverify:all -cp test/fixtures "$name" "${jdk_arguments[@]}" \
            > "$scratch/$name.reference"
    done
    for name in ReferenceIdentityMain SwitchMain WideLocalMain; do
        arguments=()
        if [[ "$name" == ReferenceIdentityMain ]]; then
            arguments=(--args '["same","same"]')
        fi
        "$MAGNOLIA" build --entry "test/fixtures/$name.class" \
            "${arguments[@]}" --no-optimize --output "$scratch/$name.oak" >/dev/null
    done
fi

# Each browser target must contain only inline scripts, and its simulated
# DOM output must equal an independent JDK oracle without a Java plugin.
for name in ReferenceIdentityMain PointGeometryMain ArrayCopyMain LongWidthMain; do
    printf '[route] checking %s HTML against JDK\n' "$name" >&2
    html_args=()
    if [[ "$name" == ReferenceIdentityMain ]]; then
        html_args=(--args '["same","same"]')
    fi
    "$MAGNOLIA" build --entry "test/fixtures/$name.class" \
        "${html_args[@]}" --output "$scratch/$name.html" >/dev/null
    node - "$scratch/$name.html" "$scratch/$name.reference" <<'JS'
const fs = require('fs'), vm = require('vm'), assert = require('assert');
const html = fs.readFileSync(process.argv[2], 'utf8');
const expected = fs.readFileSync(process.argv[3], 'utf8');
const scripts = [...html.matchAll(/<script>([\s\S]*?)<\/script>/g)].map(m => m[1]);
assert.strictEqual(scripts.length, 2, 'HTML must contain only two inline scripts');
const output = {textContent: ''};
const canvas = {style: {}, getContext() { throw Error('unexpected graphics usage'); }};
const context = {
    document: {getElementById(id) {
        return id === 'output' ? output : id === 'drawing' ? canvas : null;
    }},
    console: {log() {}, error(...messages) {throw Error(messages.join(' '));}},
};
vm.createContext(context);
for (const code of scripts) vm.runInContext(code, context, {timeout: 10000});
assert.strictEqual(output.textContent, expected);
assert.strictEqual(context.MagnoliaApplet.inspect().type, 'ok');
JS
    printf 'PASS %s HTML (two inline scripts) = independent JDK oracle\n' "$name"
done

# A same-width if_acmpeq -> if_acmpne edit produces a DIFFERENT verified
# class and changes exactly one boolean. Restoring that Oak still returns
# the ORIGINAL bytes; an integer-branch opcode in its place is rejected.
python3 - "$scratch/ReferenceIdentityMain.oak" \
    "$scratch/ReferenceIdentityMainEdited.oak" \
    "$scratch/ReferenceIdentityMainInvalid.oak" \
    "$scratch/ReferenceIdentityMain.reference" \
    "$scratch/ReferenceIdentityMainEdited.expected" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
old = '[165, 32, ?],'
if source.count(old) != 1:
    raise SystemExit('reference-identity fixture must have one selected branch')
with open(sys.argv[2], 'w', encoding='utf-8') as f:
    f.write(source.replace(old, '[166, 32, ?],'))
with open(sys.argv[3], 'w', encoding='utf-8') as f:
    f.write(source.replace(old, '[159, 32, ?],'))
original = open(sys.argv[4], encoding='utf-8').readlines()
if original[2] != 'true\n':
    raise SystemExit('expected third reference-identity result must be true')
original[2] = 'false\n'
with open(sys.argv[5], 'w', encoding='utf-8') as f:
    f.writelines(original)
PY
"$MAGNOLIA" build --entry "$scratch/ReferenceIdentityMainEdited.oak" \
    --output "$scratch/ReferenceIdentityMainEdited.class" >/dev/null
cp "$scratch/ReferenceIdentityMainEdited.class" "$scratch/edited/ReferenceIdentityMain.class"
if cmp -s test/fixtures/ReferenceIdentityMain.class "$scratch/edited/ReferenceIdentityMain.class"; then
    echo 'FAIL: edited identity branch produced original class' >&2
    exit 1
fi
java -Xverify:all -cp "$scratch/edited" ReferenceIdentityMain same same \
    > "$scratch/ReferenceIdentityMainEdited.reference"
"$MAGNOLIA" "$scratch/ReferenceIdentityMainEdited.oak" \
    > "$scratch/ReferenceIdentityMainEdited.oak-output"
diff -u "$scratch/ReferenceIdentityMainEdited.expected" \
    "$scratch/ReferenceIdentityMainEdited.reference"
diff -u "$scratch/ReferenceIdentityMainEdited.reference" \
    "$scratch/ReferenceIdentityMainEdited.oak-output"
"$MAGNOLIA" build --entry "$scratch/edited/ReferenceIdentityMain.class" \
    --args '["same","same"]' --output "$scratch/ReferenceIdentityMainEdited.js" >/dev/null
node "$scratch/ReferenceIdentityMainEdited.js" > "$scratch/ReferenceIdentityMainEdited.js-output"
diff -u "$scratch/ReferenceIdentityMainEdited.reference" \
    "$scratch/ReferenceIdentityMainEdited.js-output"
"$MAGNOLIA" build --entry "$scratch/edited/ReferenceIdentityMain.class" \
    --args '["same","same"]' --output "$scratch/ReferenceIdentityEditedConverted.java" >/dev/null
javac -d "$scratch/classes" "$scratch/ReferenceIdentityEditedConverted.java"
java -Xverify:all -cp "$scratch/classes" ReferenceIdentityEditedConverted \
    > "$scratch/ReferenceIdentityMainEdited.java-output"
diff -u "$scratch/ReferenceIdentityMainEdited.reference" \
    "$scratch/ReferenceIdentityMainEdited.java-output"
"$MAGNOLIA" build restore --entry "$scratch/ReferenceIdentityMainEdited.oak" \
    --output "$scratch/ReferenceIdentityMain.original.class" >/dev/null
cmp test/fixtures/ReferenceIdentityMain.class "$scratch/ReferenceIdentityMain.original.class"
if "$MAGNOLIA" build --entry "$scratch/ReferenceIdentityMainInvalid.oak" \
    --output "$scratch/ReferenceIdentityMainInvalid.class" > "$scratch/invalid.log" 2>&1; then
    echo 'FAIL: integer branch substituted for reference identity branch' >&2
    exit 1
fi
grep -q 'not a safe same-width edit' "$scratch/invalid.log"
printf 'PASS edited if_acmpeq -> if_acmpne: JDK/Oak/JS/Java match, restore original, unsafe edit rejected\n'

# A JAR with the same class follows the same decoder/VM/checked assembler;
# the recompiled JAR keeps its manifest/resource and verifies with -jar.
python3 - test/fixtures/ReferenceIdentityMain.class "$scratch/ReferenceIdentityMain.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[2], 'w', compression=zipfile.ZIP_DEFLATED) as jar:
    jar.writestr('META-INF/MANIFEST.MF',
                 'Manifest-Version: 1.0\r\nMain-Class: ReferenceIdentityMain\r\n\r\n')
    jar.write(sys.argv[1], 'ReferenceIdentityMain.class')
    jar.writestr('assets/identity.txt', 'original resource\n')
PY
"$MAGNOLIA" build run --entry "$scratch/ReferenceIdentityMain.jar" \
    --args '["same","same"]' > "$scratch/ReferenceIdentityMain.jar-output"
diff -u "$scratch/ReferenceIdentityMain.reference" "$scratch/ReferenceIdentityMain.jar-output"
"$MAGNOLIA" build --entry "$scratch/ReferenceIdentityMain.jar" --no-optimize \
    --args '["same","same"]' --output "$scratch/ReferenceIdentityJar.oak" >/dev/null
python3 - "$scratch/ReferenceIdentityJar.oak" "$scratch/ReferenceIdentityJarEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
old = '[165, 32, ?],'
if source.count(old) != 1:
    raise SystemExit('JAR identity branch not present once')
with open(sys.argv[2], 'w', encoding='utf-8') as f:
    f.write(source.replace(old, '[166, 32, ?],'))
PY
"$MAGNOLIA" build --entry "$scratch/ReferenceIdentityJarEdited.oak" \
    --output "$scratch/ReferenceIdentityEdited.jar" >/dev/null
java -Xverify:all -jar "$scratch/ReferenceIdentityEdited.jar" same same \
    > "$scratch/ReferenceIdentityJarEdited.reference"
diff -u "$scratch/ReferenceIdentityMainEdited.reference" \
    "$scratch/ReferenceIdentityJarEdited.reference"
"$MAGNOLIA" build restore --entry "$scratch/ReferenceIdentityJarEdited.oak" \
    --output "$scratch/ReferenceIdentityJar.original.jar" >/dev/null
cmp "$scratch/ReferenceIdentityMain.jar" "$scratch/ReferenceIdentityJar.original.jar"
python3 - "$scratch/ReferenceIdentityMain.jar" "$scratch/ReferenceIdentityEdited.jar" <<'PY'
import sys, zipfile
with zipfile.ZipFile(sys.argv[1]) as original, zipfile.ZipFile(sys.argv[2]) as edited:
    assert original.namelist() == edited.namelist(), 'JAR entries changed'
    for name in ('META-INF/MANIFEST.MF', 'assets/identity.txt'):
        assert original.read(name) == edited.read(name), name
    assert original.read('ReferenceIdentityMain.class') != edited.read('ReferenceIdentityMain.class')
PY
printf 'PASS reference identity JAR: new JDK-verified class, manifest/resource intact, original JAR restored\n'

python3 - "$scratch/SwitchMain.oak" "$scratch/SwitchMainEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
old, new = '[16, 20, ?],', '[16, 21, ?],'
if source.count(old) != 1:
    raise SystemExit('switch fixture must have exactly one editable 20 literal')
with open(sys.argv[2], 'w', encoding='utf-8') as result:
    result.write(source.replace(old, new))
PY
"$MAGNOLIA" build --entry "$scratch/SwitchMainEdited.oak" \
    --output "$scratch/SwitchMainEdited.class" >/dev/null
cp "$scratch/SwitchMainEdited.class" "$scratch/edited/SwitchMain.class"
java -Xverify:all -cp "$scratch/edited" SwitchMain > "$scratch/SwitchMainEdited.reference"
"$MAGNOLIA" "$scratch/SwitchMainEdited.oak" > "$scratch/SwitchMainEdited.oak-output"
diff -u "$scratch/SwitchMainEdited.reference" "$scratch/SwitchMainEdited.oak-output"
if cmp -s test/fixtures/SwitchMain.class "$scratch/edited/SwitchMain.class"; then
    echo 'FAIL: edited switch class is identical to the original' >&2
    exit 1
fi
printf 'PASS edited switch-method literal -> new JDK-verified class and matching Oak\n'

python3 - "$scratch/WideLocalMain.oak" "$scratch/WideLocalMainEdited.oak" <<'PY'
import sys
source = open(sys.argv[1], encoding='utf-8').read()
old, new = '[196, 260, [132, 200]],', '[196, 260, [132, 201]],'
if source.count(old) != 1:
    raise SystemExit('wide fixture must have exactly one editable iinc 260, 200')
with open(sys.argv[2], 'w', encoding='utf-8') as result:
    result.write(source.replace(old, new))
PY
"$MAGNOLIA" build --entry "$scratch/WideLocalMainEdited.oak" \
    --output "$scratch/WideLocalMainEdited.class" >/dev/null
cp "$scratch/WideLocalMainEdited.class" "$scratch/edited/WideLocalMain.class"
java -Xverify:all -cp "$scratch/edited" WideLocalMain > "$scratch/WideLocalMainEdited.reference"
"$MAGNOLIA" "$scratch/WideLocalMainEdited.oak" > "$scratch/WideLocalMainEdited.oak-output"
diff -u "$scratch/WideLocalMainEdited.reference" "$scratch/WideLocalMainEdited.oak-output"
if cmp -s test/fixtures/WideLocalMain.class "$scratch/edited/WideLocalMain.class"; then
    echo 'FAIL: edited wide class is identical to the original' >&2
    exit 1
fi
printf 'PASS edited wide iinc -> new JDK-verified class and matching Oak\n'

"$MAGNOLIA" build coverage --entry test/fixtures/SwitchMain.class \
    --json > "$scratch/coverage-success"
if "$MAGNOLIA" build coverage --entry test/fixtures/DefaultSpeaker.class \
    > "$scratch/coverage-failure"; then
    echo 'FAIL: unsupported interface default method passed whole-input coverage' >&2
    exit 1
fi
grep -q 'interface default/instance method bodies are not linked' "$scratch/coverage-failure"
printf 'PASS whole-input coverage accepts supported class and rejects unreachable blocker\n'
