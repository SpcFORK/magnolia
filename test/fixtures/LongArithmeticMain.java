// Exact 64-bit arithmetic, bitwise and masked-shift differential fixture.
// The named methods prevent javac from folding the tested JVM instructions.
public final class LongArithmeticMain {
    public static long mul(long a, long b) { return a * b; }
    public static long div(long a, long b) { return a / b; }
    public static long rem(long a, long b) { return a % b; }
    public static long and(long a, long b) { return a & b; }
    public static long or(long a, long b) { return a | b; }
    public static long xor(long a, long b) { return a ^ b; }
    public static long shl(long a, int n) { return a << n; }
    public static long shr(long a, int n) { return a >> n; }
    public static long ushr(long a, int n) { return a >>> n; }

    public static void zero() { System.out.println(div(1L, 0L)); }

    public static void main(String[] args) {
        long exact = 9007199254740993L; // 2^53 + 1 cannot be a JS Number
        long minus = -exact;
        System.out.println(mul(exact, 3L));
        System.out.println(mul(Long.MAX_VALUE, Long.MAX_VALUE));
        System.out.println(mul(Long.MIN_VALUE, -1L));
        System.out.println(div(exact, 3L));
        System.out.println(div(minus, 7L));
        System.out.println(div(Long.MIN_VALUE, -1L));
        System.out.println(rem(exact, 7L));
        System.out.println(rem(minus, 7L));
        System.out.println(rem(Long.MIN_VALUE, -1L));
        System.out.println(and(exact, -1L));
        System.out.println(or(exact, 0xf000000000000000L));
        System.out.println(xor(exact, -1L));
        System.out.println(shl(exact, 31));
        System.out.println(shl(exact, 32));
        System.out.println(shl(exact, 65));
        System.out.println(shl(exact, -65));
        System.out.println(shr(-1L, 63));
        System.out.println(shr(Long.MIN_VALUE, 31));
        System.out.println(shr(Long.MIN_VALUE, 33));
        System.out.println(ushr(-1L, 1));
        System.out.println(ushr(-1L, 32));
        System.out.println(ushr(Long.MIN_VALUE, 63));
        System.out.println(ushr(exact, 64));
        System.out.println(ushr(exact, -1));
    }
}
