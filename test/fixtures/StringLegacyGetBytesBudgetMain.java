// Owned 128-byte copy for separate VM scan/copy budget refusals.
public final class StringLegacyGetBytesBudgetMain {
    public static void main(String[] args) {
        String s = "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef" +
                   "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef";
        byte[] bytes = new byte[128];
        s.getBytes(0, 128, bytes, 0);
        System.out.println(bytes[0]);
        System.out.println(bytes[127]);
    }
}
