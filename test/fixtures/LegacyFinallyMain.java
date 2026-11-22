// This readable source specifies the owned legacy-finally behavior. The
// reproducible test tool replaces checked(int)'s javac inline-finally bytecode
// with a Java 5-era jsr/ret Code attribute (class version 49.0). The SAME
// patched class, never the javac intermediate, is run by JDK 11 and Magnolia.
public final class LegacyFinallyMain {
    static int tally;

    static int checked(int value) {
        tally = value;
        try {
            if (value < 0) throw new IllegalArgumentException();
            return tally;
        } finally {
            tally += 10;
        }
    }

    public static void main(String[] ignored) {
        System.out.println(checked(1));
        System.out.println(tally);
        try {
            checked(-1);
            System.out.println("BAD-no-exception");
        } catch (IllegalArgumentException expected) {
            System.out.println(tally);
        }
        System.out.println(checked(0));
        System.out.println(tally);
    }
}
