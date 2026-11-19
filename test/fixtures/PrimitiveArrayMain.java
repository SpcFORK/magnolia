// Signed short and unsigned char array reads/writes through real JVM opcodes.
public final class PrimitiveArrayMain {
    public static void main(String[] args) {
        int high = 65535;
        int wrap = 65536;
        int sign = 32768;
        int negative = -32769;
        char[] chars = new char[2];
        short[] shorts = new short[2];
        chars[0] = (char) high;
        chars[1] = (char) wrap;
        shorts[0] = (short) sign;
        shorts[1] = (short) negative;
        System.out.println((int) chars[0]);
        System.out.println((int) chars[1]);
        System.out.println((int) shorts[0]);
        System.out.println((int) shorts[1]);
        System.out.println((int) chars.length);
        System.out.println((int) shorts.length);
    }
}
