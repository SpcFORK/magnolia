// Ordinary Java program, not an Applet. Exercises static main(String[]).
public class HelloMain {
    public static void main(String[] args) {
        System.out.println(args.length);
        System.out.println(args[0]);
        System.out.println(twice(7));
    }
    static int twice(int n) { return n + n; }
}
