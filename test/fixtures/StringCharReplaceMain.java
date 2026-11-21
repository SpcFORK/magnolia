// Owned Java 8 JDK differential for the ASCII String.replace(char,char)
// slice used by the historical Java Flash Applet's bgcolor parameter.
public final class StringCharReplaceMain {
    public static void main(String[] args) {
        if (args.length > 0) {
            System.out.println(args[0].replace('a', 'b'));
            return;
        }
        String color = "#a#";
        String changed = color.replace('#', '0');
        System.out.println(changed);
        System.out.println(color == changed);
        System.out.println(color.equals(changed));
        System.out.println(color == color.replace('#', '#'));
        System.out.println(color == color.replace('z', '0'));
        System.out.println("ABA".replace('A', 'b'));
        System.out.println("".replace('#', '0') == "");
        String emptyReceiver = null;
        try {
            emptyReceiver.replace('x', 'y');
            System.out.println("null accepted");
        } catch (NullPointerException expected) {
            System.out.println("null caught");
        }
    }
}
