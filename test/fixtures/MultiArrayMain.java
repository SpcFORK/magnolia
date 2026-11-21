// JVM multianewarray: fully/partially allocated arrays, nested exact longs,
// reference/primitive leaves, zero-length levels and a static initializer.
// Compile reproducibly: javac --release 8 -g:none MultiArrayMain.java
public final class MultiArrayMain {
    private static final long[][] INITIAL = new long[1][1];

    public static long combine(long[][] values) {
        return values[0][1] ^ values[1][2];
    }

    public static void main(String[] args) {
        long[][] values = new long[2][3];
        System.out.println(values.length);
        System.out.println(values[0].length);
        System.out.println(values[1][2]);
        System.out.println(INITIAL[0][0]);
        values[0][1] = 9007199254740993L;
        values[1][2] = Long.MIN_VALUE | 1L; // share one bit with values[0][1]
        System.out.println(combine(values));
        System.out.println(values[0][1]);
        System.out.println(values[1][2]);
        System.out.println(values instanceof Object[]);
        System.out.println(((Object) values) instanceof int[][]);
        System.out.println(values[0] instanceof Cloneable);

        long[][] copied = new long[2][];
        System.arraycopy(values, 0, copied, 0, 2);
        System.out.println(copied[1] == values[1]);
        System.arraycopy(values[0], 1, values[1], 0, 1);
        System.out.println(values[1][0]);

        long[][][] partial = new long[2][1][];
        System.out.println(partial[0][0] == null);
        partial[0][0] = new long[1];
        System.out.println(partial[0][0][0]);
        System.out.println((new long[0][1][1]).length);

        String[][] strings = new String[1][2];
        System.out.println(strings[0][0] == null);
        strings[0][1] = "hello";
        System.out.println(strings[0][1]);
        int[][][] ints = new int[2][1][1];
        ints[1][0][0] = -17;
        System.out.println(ints[1][0][0]);
        System.out.println((new boolean[1][1][0]).length);
        System.out.println((new byte[1][1][0]).length);
        System.out.println((new char[1][1][0]).length);
        System.out.println((new short[1][1][0]).length);

        // Failures are explicit bounded VM errors, not catchable exceptions.
        if (args.length == 1) System.out.println((new long[1][-1]).length);
        if (args.length == 2) System.out.println((new long[0][-1]).length);
        if (args.length == 3) System.out.println((new long[300][300]).length);
    }
}
