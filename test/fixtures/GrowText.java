// Exercises the bounded Oak VM's text-builder growth guard.
public class GrowText {
    public static void main(String[] args) {
        StringBuilder builder = new StringBuilder(args[0]);
        builder.append("x");
        System.out.println("unreachable for a full-size input");
    }
}
