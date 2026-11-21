// JDK-accepted Unicode digit is deliberately outside Magnolia's ASCII
// radix profile. The test must refuse, never return an invented number.
public final class IntegerRadixUnicodeMain {
    public static void main(String[] args) {
        System.out.println(Integer.parseInt("\u0661", 16));
    }
}
