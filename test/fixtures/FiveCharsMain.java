// Owned Java 8 String(char[]) exact-ASCII subset; the original SParser is
// only a separate static inventory, never executed by this reference.
public final class FiveCharsMain {
    public static void main(String[] ignored) {
        char[] chars = {'A', '\0', 'z', '!'};
        String text = new String(chars);
        chars[0] = 'X'; // The String must be a snapshot, not the source array.
        System.out.println(text.length());
        System.out.println((int) text.charAt(0));
        System.out.println((int) text.charAt(1));
        System.out.println((int) text.charAt(2));
        System.out.println((int) text.charAt(3));
        try {
            new String((char[]) null);
            System.out.println("BAD-null");
        } catch (NullPointerException expected) {
            System.out.println("null-caught");
        }
    }
    // JDK-valid non-ASCII UTF-16; deliberately refused by the ASCII-only VM.
    public static void wideChar() {
        System.out.println(new String(new char[]{'\u00e9'}).length());
    }
    public static void oversized() {
        System.out.println(new String(new char[4097]).length());
    }
    // This distinct JDK constructor must fail selected preflight, not run.
    public static void rangeConstructor() {
        System.out.println(new String(new char[]{'A'}, 0, 1));
    }
}
