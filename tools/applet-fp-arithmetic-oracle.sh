#!/usr/bin/env bash
# Independent, seeded JDK -Xverify:all oracle for IEEE primitive operations.
# Bit outputs canonicalize arithmetic NaNs but preserve finite and signed-zero
# results. The class itself is generated/compiled only in a temporary folder.
set -euo pipefail
cd "$(dirname "$0")/.."
MAGNOLIA="${MAGNOLIA:-magnolia}"
for executable in "$MAGNOLIA" java javac node python3 timeout diff; do
    command -v "$executable" >/dev/null 2>&1 || { echo "missing: $executable" >&2; exit 2; }
done
scratch="$(mktemp -d)"
trap 'rm -rf "$scratch"' EXIT
mkdir -p "$scratch/jdk" "$scratch/generated"
run_bounded() {
    local label="$1" limit="$2" start=$SECONDS status
    shift 2
    printf '[fp-arithmetic] START %s (max %ss)\n' "$label" "$limit" >&2
    if timeout --signal=TERM --kill-after=3s "$limit" "$@"; then
        printf '[fp-arithmetic] PASS %s (%ss)\n' "$label" "$((SECONDS-start))" >&2
    else
        status=$?
        printf '[fp-arithmetic] FAIL %s exit=%s (%ss)\n' "$label" "$status" "$((SECONDS-start))" >&2
        return "$status"
    fi
}
python3 - "$scratch/FPSeedMain.java" <<'PY'
from pathlib import Path
import random, sys
r = random.Random(0x7f32d64)
base64 = [0, 0x8000000000000000, 1, 2, 0x000fffffffffffff,
          0x0010000000000000, 0x3ff0000000000000,
          0x4008000000000000, 0x7fefffffffffffff,
          0xffefffffffffffff, 0x7ff0000000000000, 0xfff0000000000000]
base32 = [0, 0x80000000, 1, 2, 0x007fffff,
          0x00800000, 0x3f800000, 0x40400000,
          0x7f7fffff, 0xff7fffff, 0x7f800000, 0xff800000]
def gen64():
    while True:
        n=r.getrandbits(64)
        if n>>52 & 2047 != 2047:return n

def gen32():
    while True:
        n=r.getrandbits(32)
        if n>>23 & 255 != 255:return n
A, B, FA, FB = [], [], [], []
for i in range(64):
    A.append(base64[i%len(base64)] if i < len(base64) else gen64())
    B.append(base64[(i+2)%len(base64)] if i < len(base64) else gen64())
    FA.append(base32[i%len(base32)] if i < len(base32) else gen32())
    FB.append(base32[(i+2)%len(base32)] if i < len(base32) else gen32())
arrays = '\n'.join('  static final %s[] %s = {%s};' % (
    typ, name, ','.join('0x%0*x%s' % (digits, n, suffix) for n in nums))
    for typ, name, nums, digits, suffix in [
        ('long', 'A', A, 16, 'L'), ('long', 'B', B, 16, 'L'),
        ('int', 'FA', FA, 8, ''), ('int', 'FB', FB, 8, '')])
Path(sys.argv[1]).write_text('''public final class FPSeedMain {
%s
  public static void main(String[] args) {
    for (int i = 0; i < A.length; i++) {
      double a = Double.longBitsToDouble(A[i]);
      double b = Double.longBitsToDouble(B[i]);
      System.out.println(Double.doubleToLongBits(a + b));
      System.out.println(Double.doubleToLongBits(a - b));
      System.out.println(Double.doubleToLongBits(a * b));
      System.out.println(Double.doubleToLongBits(a / b));
      System.out.println(Double.doubleToLongBits(a %% b));
      float x = Float.intBitsToFloat(FA[i]);
      float y = Float.intBitsToFloat(FB[i]);
      System.out.println(Float.floatToIntBits(x + y));
      System.out.println(Float.floatToIntBits(x - y));
      System.out.println(Float.floatToIntBits(x * y));
      System.out.println(Float.floatToIntBits(x / y));
      System.out.println(Float.floatToIntBits(x %% y));
    }
  }
}
''' % arrays)
PY
run_bounded 'javac Java 8 seeded program' 20 javac --release 8 -g:none \
    -d "$scratch/jdk" "$scratch/FPSeedMain.java"
run_bounded 'independent JDK verifier and 640 bit results' 12 java -Xverify:all \
    -cp "$scratch/jdk" FPSeedMain > "$scratch/jdk.out"
[[ $(wc -l < "$scratch/jdk.out") == 640 ]]
run_bounded 'native Oak seeded program' 40 "$MAGNOLIA" build run \
    --entry "$scratch/jdk/FPSeedMain.class" > "$scratch/native.out"
diff -u "$scratch/jdk.out" "$scratch/native.out"
run_bounded 'self-contained JS build' 95 "$MAGNOLIA" build \
    --entry "$scratch/jdk/FPSeedMain.class" --output "$scratch/FPSeedMain.js" \
    >/dev/null
run_bounded 'self-contained JS execution' 30 node "$scratch/FPSeedMain.js" \
    > "$scratch/js.out"
diff -u "$scratch/jdk.out" "$scratch/js.out"
run_bounded 'generated Java source' 40 "$MAGNOLIA" build \
    --entry "$scratch/jdk/FPSeedMain.class" \
    --output "$scratch/FPSeedConverted.java" >/dev/null
run_bounded 'javac generated Java' 20 javac -d "$scratch/generated" \
    "$scratch/FPSeedConverted.java"
run_bounded 'generated Java execution' 12 java -Xverify:all \
    -cp "$scratch/generated" FPSeedConverted > "$scratch/generated.out"
diff -u "$scratch/jdk.out" "$scratch/generated.out"
printf 'PASS 64 seeded cases × 10 primitive operations = 640 JDK bit results on Oak, JS and Java\n'
