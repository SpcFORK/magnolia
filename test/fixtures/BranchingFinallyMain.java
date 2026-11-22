// Owned Java 8 behavior for a bounded *branching* legacy jsr/ret finally.
// The independently pinned fixture builder replaces checked(int)'s javac
// inline-finally Code with Java 5-era jsr/ret bytecode at versions 49 and 45.
// This is not Macromedia bytecode, an audio backend, or a Flash movie.
public final class BranchingFinallyMain {
    static int tally;

    static int checked(int value) {
        tally = value;
        try {
            if (value < 0) throw new IllegalArgumentException();
            return tally;
        } finally {
            if (tally > 0) {
                int remaining = tally;
                while (remaining > 0) {
                    tally++;
                    remaining--;
                }
            }
            // A direct conditional edge to the legacy ret when zero tests
            // the same JDK-valid shape as the original Bitmap cleanup.
            if (tally != 0) tally += 10;
        }
    }

    public static void main(String[] ignored) {
        System.out.println(checked(3));
        System.out.println(tally);
        try {
            checked(-2);
            System.out.println("BAD-no-exception");
        } catch (IllegalArgumentException expected) {
            System.out.println(tally);
        }
        System.out.println(checked(0));
        System.out.println(tally);
        System.out.println(checked(1));
        System.out.println(tally);
    }
}
