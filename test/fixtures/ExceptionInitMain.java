// Class-initialization failures are abrupt: non-Error is wrapped once, an
// Error is not, and later active uses throw NoClassDefFoundError.
public final class ExceptionInitMain {
    static final class BrokenInit {
        static int value = Integer.parseInt("not-an-int");
    }
    static final class BrokenError {
        static int value = fail();
        static int fail() { throw new LinkageError("linkage"); }
    }
    public static void main(String[] args) {
        try { System.out.println(BrokenInit.value); }
        catch (ExceptionInInitializerError ex) {
            System.out.println("wrapped");
            System.out.println(ex.getCause() instanceof NumberFormatException);
        }
        try { System.out.println(BrokenInit.value); }
        catch (NoClassDefFoundError ex) { System.out.println("erroneous"); }
        try { System.out.println(BrokenError.value); }
        catch (LinkageError ex) { System.out.println(ex.getMessage()); }
        try { System.out.println(BrokenError.value); }
        catch (NoClassDefFoundError ex) { System.out.println("failed error class"); }
    }
}
