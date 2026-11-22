// IEEE-754 class pool, two-slot values, methods, statics/instances, primitive
// arrays, multianewarray, arithmetic, comparisons, conversion and zero signs.
// Compile reproducibly: javac --release 8 -g:none FloatingMain.java
public final class FloatingMain {
    private static final double NEGATIVE_ZERO = -0.0;
    private static final double LEAST_DOUBLE = Double.MIN_VALUE;
    private static final float LEAST_FLOAT = Float.MIN_VALUE;
    private static double saved;
    private static float narrowed;
    private double field;

    private FloatingMain(double initial) { this.field = initial; }
    private double apply(double value) { return field + value; }
    private static float smaller(float value) { return value / 2.0f; }

    public static void main(String[] args) {
        FloatingMain object = new FloatingMain(0.75);
        System.out.println(object.apply(1.25));
        saved = object.apply(2.25);
        System.out.println(saved);
        narrowed = smaller(0.5f);
        System.out.println(narrowed);
        System.out.println(LEAST_DOUBLE > 0.0);
        System.out.println(LEAST_DOUBLE / 2.0 == 0.0);
        System.out.println(LEAST_FLOAT > 0.0f);
        System.out.println(LEAST_FLOAT / 2.0f == 0.0f);
        System.out.println(NEGATIVE_ZERO == 0.0);
        System.out.println(Double.doubleToRawLongBits(NEGATIVE_ZERO));
        System.out.println(Double.doubleToRawLongBits(LEAST_DOUBLE));
        System.out.println(Float.floatToRawIntBits(LEAST_FLOAT));
        // Bit intrinsics must not canonicalize a class-file NaN payload.
        double payload = Double.longBitsToDouble(0x7ff8123456789abcL);
        System.out.println(Double.doubleToRawLongBits(payload));
        System.out.println(Float.floatToRawIntBits(Float.intBitsToFloat(0x7fc12345)));
        System.out.println(1.0 / NEGATIVE_ZERO);
        System.out.println(0.0 / 0.0);
        System.out.println(Double.POSITIVE_INFINITY - Double.POSITIVE_INFINITY);
        System.out.println(Double.POSITIVE_INFINITY + Double.NEGATIVE_INFINITY);
        double nan = 0.0 / 0.0;
        System.out.println(nan != nan);
        System.out.println(nan < 1.0);
        System.out.println(nan > 1.0);
        System.out.println((int) Double.POSITIVE_INFINITY);
        System.out.println((int) Double.NEGATIVE_INFINITY);
        System.out.println((int) nan);
        System.out.println((long) Double.NEGATIVE_INFINITY);
        System.out.println((long) Double.POSITIVE_INFINITY);
        System.out.println((long) nan);
        System.out.println((double) 9007199254740993L == 9007199254740992.0);
        System.out.println((float) 16777217 == 16777216.0f);
        System.out.println((int) (-1.75));
        System.out.println((long) (-1.75));
        System.out.println((double) ((float) 0.5));
        System.out.println((float) ((double) 0.5));
        System.out.println(-0.0);
        System.out.println(-0.0f);
        System.out.println(-0.5 % 0.5);
        System.out.println(5.5 % 2.0);
        double[] doubles = new double[2];
        doubles[0] = object.apply(0.25);
        System.arraycopy(doubles, 0, doubles, 1, 1);
        System.out.println(doubles[1]);
        System.out.println(doubles[0] == doubles[1]);
        float[] floats = new float[1];
        floats[0] = narrowed;
        System.out.println(floats[0]);
        double[][] grid = new double[1][2];
        grid[0][1] = 2.5;
        System.out.println(grid[0][1]);
        System.out.println(grid[0][0]);
    }
}
