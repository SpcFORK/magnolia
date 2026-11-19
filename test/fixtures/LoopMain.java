// Oak-only Java subset and javac reference: locals, loop and conditional.
public class LoopMain {
    public static void main(String[] args) {
        int i = 0;
        int sum = 0;
        while (i < 5) {
            sum = sum + i;
            i = i + 1;
        }
        if (sum == 10) {
            System.out.print("sum=");
            System.out.println(sum);
        } else {
            System.out.println(0);
        }
    }
}
