// JVM fixture for invokevirtual receiver dispatch and inherited invokestatic.
class DispatchBase {
    public int value() { return 1; }
    static int shared() { return 7; }
}
class DispatchChild extends DispatchBase {
    @Override public int value() { return 2; }
}
public class DispatchMain {
    public static void main(String[] args) {
        DispatchBase receiver = new DispatchChild();
        System.out.println(receiver.value());
        System.out.println(DispatchChild.shared());
    }
}
