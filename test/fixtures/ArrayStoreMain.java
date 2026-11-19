// Legal Java source that throws ArrayStoreException at the assignment.
// The bounded Oak VM must fail, not continue as though String[] were Object[].
public class ArrayStoreMain {
    public static void main(String[] args) {
        Object[] values = new String[1];
        System.out.println(values instanceof Object[]);
        values[0] = new StringBuilder("bad");
        System.out.println("unreachable");
    }
}
