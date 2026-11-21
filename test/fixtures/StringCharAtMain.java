// Bounded ASCII String.charAt and catchable index errors, independently
// compared across the JDK and Magnolia's admitted target runtimes.
public final class StringCharAtMain {
    public static void main(String[] args) {
        String text = "AZaz 09!";
        System.out.println((int) text.charAt(0));
        System.out.println((int) text.charAt(3));
        System.out.println((int) text.charAt(text.length() - 1));
        try {
            System.out.println((int) text.charAt(-1));
            System.out.println("BAD:negative");
        } catch (StringIndexOutOfBoundsException expected) {
            System.out.println("caught-negative");
        }
        try {
            System.out.println((int) text.charAt(text.length()));
            System.out.println("BAD:length");
        } catch (StringIndexOutOfBoundsException expected) {
            System.out.println("caught-length");
        }
    }
}
