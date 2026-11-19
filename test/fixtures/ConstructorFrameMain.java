// Java 8 fixture: a constructor must initialize `this`, and a fresh `new`
// reference is unusable until the exact object's <init> returns. This class
// includes aliases through dup and local variables for verifier mutations.
public final class ConstructorFrameMain {
    static int made;
    int value;

    public ConstructorFrameMain(int value) {
        this();
        this.value = value;
        made++;
    }

    public ConstructorFrameMain() { }

    static Object fresh() {
        Object copy = new Object();
        return copy;
    }

    static Object two() {
        Object first = new Object();
        Object second = new Object();
        return first;
    }

    public static void main(String[] args) {
        ConstructorFrameMain item = new ConstructorFrameMain(7);
        System.out.println(item.value);
        System.out.println(made);
        Object first = fresh();
        System.out.println(first != new Object());
        System.out.println(two() != null);
        ConstructorFrameMain empty = new ConstructorFrameMain();
        System.out.println(empty.value);
    }
}
