// javac emits an explicit StackMapTable frame at bytecode PC 0 for the
// backwards goto. JDK 11 requires that frame even though the method also
// has its implicit initial frame; main does not call the infinite loop.
public final class StackMapEntryLoopMain {
    static void spin() { for (;;) { } }

    public static void main(String[] args) {
        System.out.println(11);
    }
}
