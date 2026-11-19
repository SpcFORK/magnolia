#!/usr/bin/env bash
# Independent negative JDK -Xverify:all comparisons for exception tables and
# athrow. Reject before any fixture output, not merely after a partial run.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" java python3 timeout grep; do
    command -v "$executable" >/dev/null 2>&1 || { echo "missing: $executable" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
python3 - test/fixtures/ExceptionMain.class "$scratch" <<'PY'
from pathlib import Path
import sys
raw=Path(sys.argv[1]).read_bytes();out=Path(sys.argv[2]);p=0

def take(n):
    global p
    b=raw[p:p+n];assert len(b)==n;p+=n;return b

def u1():return take(1)[0]
def u2():return int.from_bytes(take(2),'big')
def u4():return int.from_bytes(take(4),'big')
assert take(4)==b'\xca\xfe\xba\xbe';take(4)
cp=[None]*u2();i=1
while i<len(cp):
    tag=u1()
    if tag==1:cp[i]=(tag,take(u2()).decode('utf-8'))
    elif tag in (3,4):take(4)
    elif tag in (5,6):take(8);i+=1
    elif tag in (7,8,16,19,20):cp[i]=(tag,u2())
    elif tag in (9,10,11,12,17,18):take(4)
    elif tag==15:take(3)
    else:raise ValueError(tag)
    i+=1
take(2);this=u2();take(2);assert cp[cp[this][1]][1]=='ExceptionMain'
take(2*u2())
def attrs():
    found=[]
    for _ in range(u2()):
        name=cp[u2()][1];size=u4();offset=p;take(size)
        found.append((name,offset,size))
    return found
for _ in range(u2()):take(6);attrs()
main=None
for _ in range(u2()):
    take(2);name=cp[u2()][1];desc=cp[u2()][1]
    for kind,offset,size in attrs():
        if kind=='Code' and name+desc=='main([Ljava/lang/String;)V':
            n=int.from_bytes(raw[offset+4:offset+8],'big')
            code=offset+8;table=code+n
            assert n>40 and int.from_bytes(raw[table:table+2],'big')>5
            main=(code,raw[code:code+n],table+2,n)
assert main is not None
code,body,table,n=main

def patch(label,place,before,after):
    assert len(before)==len(after) and raw[place:place+len(before)]==before, label
    folder=out/label;folder.mkdir();altered=bytearray(raw)
    altered[place:place+len(before)]=after
    (folder/'ExceptionMain.class').write_bytes(altered)
first=raw[table:table+8];assert len(first)==8
start,end,target,catch=(int.from_bytes(first[i:i+2],'big') for i in (0,2,4,6))
assert start < end <= n and target < n and catch != 0
patch('empty-range',table,first[:2],first[2:4])
patch('end-past-code',table+2,first[2:4],b'\xff\xff')
patch('target-inside-instruction',table+4,first[4:6],b'\x00\x01')
patch('non-throwable-catch',table+6,first[6:8],this.to_bytes(2,'big'))
assert body.count(b'\x01\xbf')==1
throw=body.index(b'\x01\xbf')
patch('athrow-int',code+throw,b'\x01',b'\x03')
PY
while read -r label diagnostic; do
    printf '[exception-oracles] START %s (8s per tool)\n' "$label" >&2
    folder="$scratch/$label"
    if timeout --signal=TERM --kill-after=2s 8s "$MAGNOLIA" build run \
        --entry "$folder/ExceptionMain.class" >"$folder/native.log" 2>&1; then
        echo "FAIL: Magnolia accepted $label" >&2; exit 1
    fi
    grep -q "$diagnostic" "$folder/native.log" || {
        echo "FAIL: unexpected Magnolia diagnostic for $label" >&2
        tail -28 "$folder/native.log" >&2; exit 1;
    }
    if timeout --signal=TERM --kill-after=2s 8s java -Xverify:all \
        -cp "$folder" ExceptionMain >"$folder/jdk.log" 2>&1; then
        echo "FAIL: JDK accepted $label" >&2; exit 1
    fi
    if ! grep -Eq 'VerifyError|ClassFormatError' "$folder/jdk.log"; then
        echo "FAIL: JDK did not reject $label as invalid bytecode" >&2
        tail -28 "$folder/jdk.log" >&2; exit 1
    fi
    if grep -Eq '^parse$|^true$|^wrapped$' "$folder/native.log"; then
        echo "FAIL: malformed $label ran before Magnolia preflight" >&2; exit 1
    fi
    printf 'PASS %-27s rejected by Magnolia preflight and JDK -Xverify:all\n' "$label"
done <<'CASES'
empty-range invalid.exception.table
end-past-code invalid.exception.table
target-inside-instruction exception.handler.is.not.on.instruction.boundaries
non-throwable-catch unsupported.or.non-Throwable.catch.class
athrow-int expected.A,.found.I
CASES
