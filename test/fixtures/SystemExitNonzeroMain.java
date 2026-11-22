// JDK legitimately exits with code 7. Finite Magnolia deliberately refuses
// nonzero codes (even with the zero-exit capability), before guest output.
public final class SystemExitNonzeroMain {
    public static void main(String[] args) {
        System.exit(7);
        System.out.println("UNREACHABLE after nonzero exit");
    }
}
