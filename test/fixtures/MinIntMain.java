// Java's minimum int is legal only as the direct operand of unary minus.
public class MinIntMain {
    public static void main(String[] args) {
        System.out.println(-2147483648);
        if (true == true) { System.out.println(1); }
    }
}
