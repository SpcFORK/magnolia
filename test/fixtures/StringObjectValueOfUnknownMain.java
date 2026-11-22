// JDK-valid negative control: the runtime value is a String, but its verifier
// type is Object. Magnolia refuses this broader toString path before output.
public final class StringObjectValueOfUnknownMain {
    private static Object widen(String text) { return text; }
    public static void main(String[] args) {
        System.out.println("BEGIN");
        Object unknown = widen("Mint");
        System.out.println(String.valueOf(unknown));
    }
}
