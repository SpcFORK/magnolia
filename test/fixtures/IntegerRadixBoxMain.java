// Owned Java 8 comparison of the exact Integer.valueOf(String,int) call
// reached by the 1998 Java Flash Applet's bgcolor parameter. Not a player.
public final class IntegerRadixBoxMain {
    private static void invalid(String text, int radix) {
        try {
            Integer.valueOf(text, radix);
            System.out.println("BAD accepted");
        } catch (NumberFormatException expected) {
            System.out.println("NumberFormatException");
        }
    }

    public static void main(String[] args) {
        if (args.length > 0) {
            System.out.println(Integer.valueOf(args[0], 16).intValue());
            return;
        }
        String color = "#00Ff".replace('#', '0');
        System.out.println(color);
        System.out.println(Integer.valueOf(color, 16).intValue());
        Integer a = Integer.valueOf("+007f", 16);
        System.out.println(a.intValue());
        System.out.println(a == Integer.valueOf("127", 10));
        System.out.println(Integer.valueOf("-80", 16) == Integer.valueOf("-128", 10));
        System.out.println(Integer.valueOf("80", 16) == Integer.valueOf("128", 16));
        System.out.println(Integer.valueOf("-80000000", 16).intValue());
        System.out.println(Integer.valueOf("7fffffff", 16).intValue());
        System.out.println(Integer.valueOf("101010", 2).intValue());
        System.out.println(Integer.valueOf("z", 36).intValue());
        invalid(null, 16);
        invalid("", 16);
        invalid("+", 16);
        invalid("0x1f", 16);
        invalid("100000000", 16);
        invalid("-80000001", 16);
        invalid(" 1", 16);
        invalid("1", 1);
        invalid("1", 37);
        invalid("G", 16);
    }
}
