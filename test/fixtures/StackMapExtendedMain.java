// Java 8 fixture whose exception target is far enough from the start of the
// protected range to emit frame_type 247 (same_locals_1_stack_item_extended).
// Keep the sequential increments: javac otherwise emits a shorter frame.
public final class StackMapExtendedMain {
    static int read(int n) {
        try {
            int value = n;
            value += 1;
            value += 2;
            value += 3;
            value += 4;
            value += 5;
            value += 6;
            value += 7;
            value += 8;
            value += 9;
            value += 10;
            value += 11;
            value += 12;
            value += 13;
            value += 14;
            value += 15;
            value += 16;
            value += 17;
            value += 18;
            value += 19;
            value += 20;
            value += 21;
            value += 22;
            value += 23;
            value += 24;
            value += 25;
            value += 26;
            value += 27;
            value += 28;
            return value;
        } catch (ArithmeticException ex) {
            return -1;
        }
    }

    public static void main(String[] args) {
        System.out.println(read(3));
    }
}
