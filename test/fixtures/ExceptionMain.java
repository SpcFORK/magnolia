// Java 8 fixture for ordered exception tables, explicit and implicit
// throw sites, catch-all finally, cross-frame propagation, and host types.
public final class ExceptionMain {
    private static int finalized;
    static int divide(int left, int right) { return left / right; }
    static void bounce(RuntimeException value) { throw value; }
    static void failInFinally() {
        try { bounce(new IllegalStateException("leave")); }
        finally { finalized = 17; }
    }
    public static void main(String[] args) {
        try { System.out.println(Integer.parseInt("bad")); }
        catch (NumberFormatException ex) { System.out.println("parse"); }
        try { divide(7, 0); }
        catch (ArithmeticException ex) {
            System.out.println(ex.getMessage().equals("/ by zero"));
        }
        try { long value = 3L / 0L; System.out.println(value); }
        catch (ArithmeticException ex) { System.out.println("long"); }
        RuntimeException original = new IllegalArgumentException("signal");
        try { bounce(original); }
        catch (IllegalArgumentException ex) { System.out.println(ex == original); }
        catch (Exception ex) { System.out.println("wrong catch"); }
        try { throw null; }
        catch (NullPointerException ex) { System.out.println("null throw"); }
        try { String missing = null; System.out.println(missing.length()); }
        catch (NullPointerException ex) { System.out.println("null call"); }
        try { int[] values = new int[1]; System.out.println(values[2]); }
        catch (ArrayIndexOutOfBoundsException ex) { System.out.println("bounds"); }
        try { int[] values = new int[-1]; System.out.println(values.length); }
        catch (NegativeArraySizeException ex) { System.out.println("negative"); }
        try { String wrongly = (String) new Object(); System.out.println(wrongly); }
        catch (ClassCastException ex) { System.out.println("cast"); }
        Object[] source = {"copied", new Object()};
        String[] dest = new String[2];
        try { System.arraycopy(source, 0, dest, 0, 2); }
        catch (ArrayStoreException ex) { System.out.println(dest[0]); }
        try { failInFinally(); }
        catch (RuntimeException ex) {
            System.out.println(ex.getMessage());
            System.out.println(finalized);
        }
        try { throw new java.io.IOException("io"); }
        catch (java.io.IOException ex) { System.out.println(ex.getMessage()); }
    }
}
