// JDK-valid counterexample: outside Magnolia's bounded ASCII String model.
public final class StringLegacyGetBytesUnicodeMain {
    public static void main(String[] args) {
        byte[] result = new byte[1];
        "\u00e9".getBytes(0, 1, result, 0);
        System.out.println(result[0]);
    }
}
