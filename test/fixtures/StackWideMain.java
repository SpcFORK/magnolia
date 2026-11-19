// Java 8 compiler exercises category-1/2 stack shuffles for compound
// field/array writes without replacing Java bytecode with a fake sequence.
public final class StackWideMain {
    static long[] longs = {7L, 11L};
    static int[] ints = {2, 3};
    long count;
    double rate;
    static int incInt(int[] a, int i) { return a[i] += 4; }
    static long incLong(long[] a, int i) { return a[i] += 3L; }
    static long incField(StackWideMain v) { return v.count += 5L; }
    static double incDouble(StackWideMain v) { return v.rate += 0.5; }
    public static void main(String[] args) {
        StackWideMain v = new StackWideMain();
        System.out.println(incInt(ints, 1));
        System.out.println(ints[1]);
        System.out.println(incLong(longs, 0));
        System.out.println(longs[0]);
        System.out.println(incField(v));
        System.out.println(v.count);
        System.out.println(Double.doubleToRawLongBits(incDouble(v)));
        System.out.println(Double.doubleToRawLongBits(v.rate));
    }
}
