// JVM subtype checks over the bounded host objects and array forms.
// Their result must match Java rather than silently returning false.
public class HostSubtypeMain {
    public static void main(String[] args) {
        Object builder = new StringBuilder("ok");
        System.out.println(builder instanceof Object);
        System.out.println(builder instanceof StringBuilder);
        System.out.println(args instanceof Object);
        System.out.println(args instanceof Cloneable);
        System.out.println(args instanceof java.io.Serializable);
        System.out.println(args instanceof Object[]);
        System.out.println(args instanceof String[]);
        Object[] refs = new String[2];
        System.out.println(refs instanceof Object);
        System.out.println(refs instanceof Object[]);
        System.out.println(refs instanceof String[]);
        Object numberArray = new int[1];
        System.out.println(numberArray instanceof Object);
        System.out.println(numberArray instanceof Cloneable);
        System.out.println(numberArray instanceof Object[]);
        Object color = new java.awt.Color(1, 2, 3);
        System.out.println(color instanceof Object);
    }
}
