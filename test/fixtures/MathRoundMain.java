// Independently compiled Java 8 raw-bit Math.round(float/double) oracle.
// The double immediately BELOW +0.5 catches a naive host (x + 0.5) error.
// The 2^52/2^63 boundaries check exact two-limb long output and saturation.
public final class MathRoundMain {
    private static void floats() {
        int[] bits = {
            0x00000000, 0x80000000, 0x00000001, 0x80000001,
            0x3effffff, 0x3f000000, 0x3f000001,
            0xbeffffff, 0xbf000000, 0xbf000001,
            0x3f7fffff, 0xbf7fffff, 0x3fc00000, 0xbfc00000,
            0x4affffff, 0x4b000000, 0x4b000001,
            0x4effffff, 0x4f000000, 0xcf000000,
            0x7f7fffff, 0xff7fffff, 0x7f800000, 0xff800000,
            0x7fc00001, 0xffc12345
        };
        for (int value : bits)
            System.out.println(Math.round(Float.intBitsToFloat(value)));
    }

    private static void doubles() {
        long[] bits = {
            0x0000000000000000L, 0x8000000000000000L,
            0x0000000000000001L, 0x8000000000000001L,
            0x3fdfffffffffffffL, 0x3fe0000000000000L,
            0x3fe0000000000001L, 0xbfdfffffffffffffL,
            0xbfe0000000000000L, 0xbfe0000000000001L,
            0x3ff0000000000000L, 0xbff0000000000000L,
            0x3ff8000000000000L, 0xbff8000000000000L,
            0x432fffffffffffffL, 0x4330000000000000L,
            0x4330000000000001L, 0xc32fffffffffffffL,
            0xc330000000000000L, 0xc330000000000001L,
            0x43dfffffffffffffL, 0x43e0000000000000L,
            0xc3dfffffffffffffL, 0xc3e0000000000000L,
            0x7fefffffffffffffL, 0xffefffffffffffffL,
            0x7ff0000000000000L, 0xfff0000000000000L,
            0x7ff8000000000001L, 0xfff8000000000001L
        };
        for (long value : bits)
            System.out.println(Math.round(Double.longBitsToDouble(value)));
    }

    public static void main(String[] args) {
        floats();
        doubles();
    }
}
