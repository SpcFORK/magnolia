// Java 8 StackMapTable fixture: the ternary branches while an object from
// `new` is still uninitialized. javac emits full_frame entries containing
// ITEM_Uninitialized offsets pointing back to that allocation site. This
// exercises structural frame decoding, not full frame type verification.
final class FrameCtor {
    int value;
    FrameCtor(int value) { this.value = value; }
}

public final class StackMapFrameMain {
    static int make(int choice) {
        FrameCtor item = new FrameCtor(choice != 0 ? 7 : 9);
        return item.value;
    }

    public static void main(String[] args) {
        // Must not print this marker if a reachable make() frame fails
        // selected-entry preflight, even though main calls make later.
        System.out.println(101);
        System.out.println(make(1));
        System.out.println(make(0));
    }
}
