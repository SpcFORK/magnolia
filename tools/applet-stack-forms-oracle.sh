#!/usr/bin/env bash
# Independent JDK -Xverify:all oracle for every legal JVMS dup_x2/dup2_x1/
# dup2_x2 category form and swap. A tiny one-off class writer emits *actual*
# Code bytes; JDK verification, not a Magnolia-runner expectation, is the oracle.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for cmd in "$MAGNOLIA" java javac node python3 timeout diff cmp; do
    command -v "$cmd" >/dev/null 2>&1 || { echo "missing: $cmd" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/generated"
run() {
    local label="$1" seconds="$2" start=$SECONDS rc
    shift 2
    printf '[stack-forms] START %s (max %ss)\n' "$label" "$seconds" >&2
    if timeout --signal=TERM --kill-after=3s "$seconds" "$@"; then
        printf '[stack-forms] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        rc=$?
        printf '[stack-forms] FAIL %s exit=%s (%ss)\n' "$label" "$rc" "$((SECONDS-start))" >&2
        return "$rc"
    fi
}
python3 - "$scratch/StackForms.class" "$scratch/invalid/StackForms.class" "$scratch/expected.out" <<'PY'
from pathlib import Path
import struct,sys
u1=lambda n:struct.pack('>B',n)
u2=lambda n:struct.pack('>H',n)
u4=lambda n:struct.pack('>I',n)
cp=[]
def add(tag,payload):
    cp.append(u1(tag)+payload)
    return len(cp)
def utf(text):return add(1,u2(len(text.encode()))+text.encode())
def clazz(name):return add(7,u2(utf(name)))
def nat(name,desc):return add(12,u2(utf(name))+u2(utf(desc)))
def methodref(owner,name,desc):return add(10,u2(owner)+u2(nat(name,desc)))
def fieldref(owner,name,desc):return add(9,u2(owner)+u2(nat(name,desc)))
this=clazz('StackForms');superclass=clazz('java/lang/Object')
code_attr=utf('Code')
system_class=clazz('java/lang/System');out=fieldref(system_class,'out','Ljava/io/PrintStream;')
printstream=clazz('java/io/PrintStream')
println=methodref(printstream,'println','(I)V')
init_ref=methodref(superclass,'<init>','()V')
methods=[]
def m(access,name,desc,stack,locals,code):
    methods.append((access,utf(name),utf(desc),stack,locals,bytes(code)))
m(1,'<init>','()V',1,1,[42,183,*u2(init_ref),177])
# Explicit categories from JVMS 6.5. Each form stores the *entire* resulting
# stack into typed locals before computing a positional checksum. An incorrect
# shuffle is either rejected by -Xverify:all or produces a different checksum.
# (stack input categories, opcode, stack output positions into the input)
forms=[
 ('swap',                 'II',   95, [1,0]),
 ('dupX2Narrow',          'III',  91, [2,0,1,2]),
 ('dupX2Wide',            'JI',   91, [1,0,1]),
 ('dup2X1Narrow',         'III',  93, [1,2,0,1,2]),
 ('dup2X1Wide',           'IJ',   93, [1,0,1]),
 ('dup2X2Narrow',         'IIII', 94, [2,3,0,1,2,3]),
 ('dup2X2BottomWide',     'JII',  94, [1,2,0,1,2]),
 ('dup2X2TopWide',        'IIJ',  94, [2,0,1,2]),
 ('dup2X2TwoWide',        'JJ',   94, [1,0,1]),
]
weights=[2,3,5,7,11,13]
main=[];expected=[]
for name,kinds,op,order in forms:
    desc='('+kinds+')I'
    # Explicit local loads, rather than pre-pushed constants, let the same
    # bytecode be JDK-verified and executed through all output routes.
    body=[];base=0
    values=[]
    for index,kind in enumerate(kinds):
        body.append((26 if kind=='I' else 30)+base)
        values.append((13+19*index) if kind=='I' else index%2)
        base+=1 if kind=='I' else 2
    body.append(op)
    cursor=base;stored=[]
    for index in reversed(order):
        kind=kinds[index]
        body += [54 if kind=='I' else 55,cursor]
        stored.append((kind,cursor,index))
        cursor += 1 if kind=='I' else 2
    body.append(3)  # iconst_0 accumulator
    for ordinal,(kind,local,index) in enumerate(stored):
        body += [21 if kind=='I' else 22,local]
        if kind=='J':body.append(136)  # l2i
        body += [16,weights[ordinal],104,96]  # bipush, imul, iadd
    body.append(172)
    max_stack=max(3,sum(1 if kinds[index]=='I' else 2 for index in order))
    m(9,name,desc,max_stack,cursor,body)
    expected.append(sum(values[index]*weights[j] for j,index in enumerate(reversed(order))))
    ref=methodref(this,name,desc)
    main += [178,*u2(out)]
    for kind,value in zip(kinds,values):
        if kind=='J':main.append(9+value)  # lconst_0 or lconst_1
        else:main += [16,value]  # bipush 13..70
    main += [184,*u2(ref),182,*u2(println)]
main.append(177)
m(9,'main','([Ljava/lang/String;)V',7,1,main)
Path(sys.argv[3]).write_text(''.join(str(value)+'\n' for value in expected))
method_bytes=b''
for access,name,desc,stack,locals,body in methods:
    payload=u2(stack)+u2(locals)+u4(len(body))+body+u2(0)+u2(0)
    method_bytes+=u2(access)+u2(name)+u2(desc)+u2(1)+u2(code_attr)+u4(len(payload))+payload
header=b'\xca\xfe\xba\xbe'+u2(0)+u2(52)+u2(len(cp)+1)+b''.join(cp)
image=header+u2(0x21)+u2(this)+u2(superclass)+u2(0)+u2(0)+u2(len(methods))+method_bytes+u2(0)
Path(sys.argv[1]).write_bytes(image)
# Change actual Code operands in independently JDK-verifiable classes, keeping
# constant pools and code lengths fixed. Each mutation violates the opcode's
# category form at the shuffle, not merely a later result comparison.
variants=[
 ('swap-wide', [30,28,91], [30,28,95]),
 ('dup2x1-wide-below-wide', [30,32,94], [30,32,93]),
 ('dup2x1-narrow-over-wide', [30,28,29,94], [30,28,29,93]),
 ('dup2x2-second-wide', [30,28,29,94], [28,30,29,94]),
 ('dup2x2-wide-over-wide-narrow', [26,27,32,94], [32,26,32,94]),
 ('dupx2-wide-top', [26,27,32,94], [26,27,32,91]),
 ('dupx2-split-wide', [30,28,29,94], [30,28,29,91]),
]
root=Path(sys.argv[2]).parent.parent
for label,old,new in variants:
    old,new=bytes(old),bytes(new)
    assert len(old)==len(new) and image.count(old)==1,(label,image.count(old))
    target=root/('invalid-'+label)/'StackForms.class'
    target.parent.mkdir(parents=True,exist_ok=True)
    target.write_bytes(image.replace(old,new))
print('[stack-forms] emitted',len(methods),'real JVM methods,',len(variants),'negative category cases; class bytes',len(image),file=sys.stderr)
PY
run 'JDK -Xverify:all nine legal forms' 12 java -Xverify:all -cp "$scratch" StackForms >"$scratch/jdk.out"
diff -u "$scratch/expected.out" "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 9 ]]
run 'native Oak actual class' 25 "$MAGNOLIA" build run --entry "$scratch/StackForms.class" >"$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run 'build standalone Oak' 40 "$MAGNOLIA" build --entry "$scratch/StackForms.class" --no-optimize --output "$scratch/StackForms.oak" >/dev/null
run 'execute standalone Oak' 18 "$MAGNOLIA" "$scratch/StackForms.oak" >"$scratch/oak.out"
diff -u "$scratch/jdk.out" "$scratch/oak.out"
run 'build bundled JS' 95 "$MAGNOLIA" build --entry "$scratch/StackForms.class" --output "$scratch/StackForms.js" >/dev/null
run 'execute bundled JS' 25 node "$scratch/StackForms.js" >"$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run 'generate Java' 45 "$MAGNOLIA" build --entry "$scratch/StackForms.class" --output "$scratch/StackFormsConverted.java" >/dev/null
run 'compile generated Java' 20 javac -d "$scratch/generated" "$scratch/StackFormsConverted.java"
run 'execute generated Java' 12 java -Xverify:all -cp "$scratch/generated" StackFormsConverted >"$scratch/java.out"
diff -u "$scratch/jdk.out" "$scratch/java.out"
run 'restore original class' 15 "$MAGNOLIA" build restore --entry "$scratch/StackForms.oak" --output "$scratch/restored.class" >/dev/null
cmp "$scratch/StackForms.class" "$scratch/restored.class"
for label in swap-wide dup2x1-wide-below-wide dup2x1-narrow-over-wide \
        dup2x2-second-wide dup2x2-wide-over-wide-narrow dupx2-wide-top dupx2-split-wide; do
    if timeout --signal=TERM --kill-after=2s 8s java -Xverify:all \
            -cp "$scratch/invalid-$label" StackForms >"$scratch/invalid.jdk" 2>&1; then
        echo "FAIL: JDK accepted illegal $label" >&2; exit 1
    fi
    grep -Eq 'VerifyError|ClassFormatError' "$scratch/invalid.jdk" || {
        echo "FAIL: unexpected JDK result on $label" >&2; cat "$scratch/invalid.jdk"; exit 1;
    }
    if timeout --signal=TERM --kill-after=2s 8s "$MAGNOLIA" build run \
            --entry "$scratch/invalid-$label/StackForms.class" >"$scratch/invalid.oak" 2>&1; then
        echo "FAIL: Magnolia accepted illegal $label" >&2; exit 1
    fi
    grep -Eq 'wrong value category|operand stack|narrow value' "$scratch/invalid.oak" || {
        echo "FAIL: unexpected Magnolia diagnostic on $label" >&2; cat "$scratch/invalid.oak"; exit 1;
    }
    printf '[stack-forms] PASS invalid %s rejected by JDK and Magnolia\n' "$label" >&2
done
printf 'PASS 9 legal JVM shuffle forms across JDK/Oak/JS/Java; 7 invalid category forms rejected by both verifiers; original exact\n'
