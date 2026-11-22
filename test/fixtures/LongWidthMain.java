// End-to-end two-slot JVM long oracle. No 64-bit integer is routed through a
// JavaScript Number: 2^53+1 must remain distinct from 2^53.
public final class LongWidthMain {
    public static final long CONSTANT = 9007199254740993L;
    private static long current;
    private long instance;

    public static long bump(long input) { return input + 1L; }
    public long remember(long input) { instance = input; return instance; }

    public static void main(String[] args) {
        long exact = 9007199254740993L;
        System.out.println(exact);
        System.out.println(exact + 2L);
        System.out.println(bump(exact) - exact);
        System.out.println((int) exact);
        System.out.println((long) -1);
        System.out.println(exact > 9007199254740992L ? 1 : 0);
        long min = Long.MIN_VALUE;
        System.out.println(-min); // signed 64-bit overflow wraps to MIN_VALUE
        System.out.println(min < 0L ? 1 : 0);
        System.out.println(current); // default zero field
        System.out.println(current = exact); // category-two dup2
        System.out.println(current);
        LongWidthMain box = new LongWidthMain();
        System.out.println(box.instance); // default zero instance field
        System.out.println(box.remember(exact));
        bump(1L); // discarded category-two result uses pop2
    }
}
