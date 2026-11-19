// Exact two-slot long[] values across defaults, loads/stores, method calls,
// overlapping System.arraycopy, a static field and a nested reference array.
// Compile reproducibly with: javac --release 8 -g:none LongArrayMain.java
public final class LongArrayMain {
    private static long[] retained;

    public static long combine(long[] values) {
        return values[0] ^ values[2];
    }

    public static long select(long[][] rows, int row, int column) {
        return rows[row][column];
    }

    public static void main(String[] args) {
        long[] values = new long[4];
        System.out.println(values[0]);
        values[0] = 9007199254740993L;
        values[1] = Long.MIN_VALUE;
        values[2] = -1L;
        values[3] = values[0] + values[2];
        System.out.println(values[0]);
        System.out.println(values[1]);
        System.out.println(values[2]);
        System.out.println(values[3]);
        System.out.println(combine(values));

        System.arraycopy(values, 0, values, 1, 3);
        System.out.println(values[0]);
        System.out.println(values[1]);
        System.out.println(values[2]);
        System.out.println(values[3]);

        long[] copy = new long[4];
        System.arraycopy(values, 0, copy, 0, copy.length);
        retained = copy;
        System.out.println(retained[2]);
        System.out.println(copy instanceof Cloneable);

        long[][] rows = new long[2][];
        rows[0] = values;
        rows[1] = retained;
        System.out.println(rows[0] == values);
        System.out.println(select(rows, 1, 3));

        // Java errors are currently explicit bounded VM failures, not
        // catchable JVM exceptions. Keep these paths out of normal output.
        if (args.length == 1) System.out.println(values[4]);
        if (args.length == 2) System.arraycopy(values, 0, new int[4], 0, 0);
        if (args.length == 3) {
            long[] missing = null;
            System.out.println(missing[0]);
        }
    }
}
