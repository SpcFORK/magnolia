// JVM int narrowing conversions: i2b, i2c, i2s retain signed/unsigned rules.
public class NarrowIntMain {
    public static void main(String[] args) {
        int raw = 255;
        int negative = -1;
        int wide = 65535;
        int narrowedByte = (byte) raw;
        int unsignedChar = (char) negative;
        int narrowedShort = (short) wide;
        System.out.println(narrowedByte);
        System.out.println(unsignedChar);
        System.out.println(narrowedShort);
    }
}
