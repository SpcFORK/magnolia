// Independent Java 8 conversion oracle. The 2^55+2^31 +/- 1 values expose
// double-rounding if a 64-bit JVM long is first converted through binary64.
// Reproduce: javac --release 8 -g:none LongFloatRoundingMain.java
public final class LongFloatRoundingMain {
    static float narrow(long value) { return (float) value; }
    static double widen(long value) { return (double) value; }
    public static void main(String[] args) {
        long halfway = (1L << 55) + (1L << 31);
        long[] values = {
            0L, 1L, -1L, 16777215L, 16777216L, 16777217L,
            16777218L, 16777219L, 9007199254740993L,
            halfway - 1, halfway, halfway + 1,
            -halfway - 1, -halfway, -halfway + 1,
            Long.MAX_VALUE, Long.MIN_VALUE, (1L << 62) + (1L << 38) + 1
        };
        for (long value : values) {
            System.out.println(Float.floatToRawIntBits(narrow(value)));
            System.out.println(Double.doubleToRawLongBits(widen(value)));
        }
    }
}
