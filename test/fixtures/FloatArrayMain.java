// Reproducible --release 8 fixture for F/D array, conversion, and bit paths.
// Values are printed as raw int/long bits to keep the JDK as the decimal oracle.
public final class FloatArrayMain {
    static float held;
    static double remembered;
    static float returnFloat(float[] values) { return values[1] - values[0]; }
    static double returnDouble(double[] values) { return values[1] / values[0]; }

    public static void main(String[] args) {
        float[] floats = new float[3];
        double[] doubles = new double[3];
        System.out.println(Float.floatToRawIntBits(floats[0]));
        System.out.println(Double.doubleToRawLongBits(doubles[0]));
        floats[0] = Float.intBitsToFloat(0x80000000);
        floats[1] = 0.75f;
        doubles[0] = Double.longBitsToDouble(0x8000000000000000L);
        doubles[1] = 0.75;
        System.out.println(Float.floatToRawIntBits(floats[0]));
        System.out.println(Double.doubleToRawLongBits(doubles[0]));
        held = returnFloat(floats);
        remembered = returnDouble(doubles);
        System.out.println(Float.floatToRawIntBits(held));
        System.out.println(Double.doubleToRawLongBits(remembered));
        float[] copy = new float[3];
        double[] copied = new double[3];
        System.arraycopy(floats, 0, copy, 0, 3);
        System.arraycopy(doubles, 0, copied, 0, 3);
        System.out.println(Float.floatToRawIntBits(copy[1]));
        System.out.println(Double.doubleToRawLongBits(copied[1]));
        float[][] floatGrid = new float[2][3];
        double[][] doubleGrid = new double[2][3];
        System.out.println(Float.floatToRawIntBits(floatGrid[1][1]));
        System.out.println(Double.doubleToRawLongBits(doubleGrid[0][2]));
        floatGrid[1][1] = copy[1];
        doubleGrid[0][2] = copied[1];
        System.out.println(Float.floatToRawIntBits(floatGrid[1][1]));
        System.out.println(Double.doubleToRawLongBits(doubleGrid[0][2]));
        System.out.println((int) Float.intBitsToFloat(0x7f800000));
        System.out.println((int) Float.intBitsToFloat(0xff800000));
        System.out.println((int) Float.intBitsToFloat(0x7fc12345));
        System.out.println((long) Double.longBitsToDouble(0x7ff0000000000000L));
        System.out.println((long) Double.longBitsToDouble(0xfff0000000000000L));
        System.out.println((long) Double.longBitsToDouble(0x7ff8123456789abcL));
        System.out.println(Float.floatToRawIntBits((float) Long.MAX_VALUE));
        System.out.println(Double.doubleToRawLongBits((double) Long.MIN_VALUE));
        System.out.println(Float.floatToRawIntBits((float) 16777217));
        System.out.println(Double.doubleToRawLongBits((double) 2147483647));
        System.out.println(Double.doubleToRawLongBits((double) held));
        System.out.println(Float.floatToRawIntBits((float) remembered));
    }
}
