// Java 8-compiled JDK oracle for the exact Integer.valueOf(String) and
// Integer.intValue() pair reached by original Sun Blink.init. This does not
// grant arbitrary Number/Integer constructors or printing boxed objects.
import java.io.Serializable;

public final class IntegerBoxStateMain {
    private static void rejected(String text) {
        try {
            Integer.valueOf(text);
            System.out.println("BAD accepted invalid decimal");
        } catch (NumberFormatException expected) {
            System.out.println("NumberFormatException");
        }
    }

    public static void main(String[] args) {
        Integer low = Integer.valueOf("-128");
        Integer a = Integer.valueOf("+00000127");
        System.out.println(low.intValue());
        System.out.println(a.intValue());
        System.out.println(a == Integer.valueOf("127"));
        System.out.println(low == Integer.valueOf("-128"));
        System.out.println(Integer.valueOf("-0") == Integer.valueOf("0"));
        System.out.println(Integer.valueOf("128") == Integer.valueOf("128"));
        System.out.println(Integer.valueOf("-129") == Integer.valueOf("-129"));
        System.out.println(a instanceof Number);
        System.out.println(a instanceof Serializable);
        System.out.println(Integer.valueOf("2147483647").intValue());
        System.out.println(Integer.valueOf("-2147483648").intValue());
        rejected(null);
        rejected("");
        rejected(" 1");
        rejected("12x");
        rejected("2147483648");
        rejected("-2147483649");
    }
}
