// JDK System.exit(0) stops at this call. Magnolia's opt-in, application-only
// bounded VM must stop the guest and preserve only the preceding output.
public final class SystemExitZeroMain {
    public static void main(String[] args) {
        System.out.println("before-zero");
        System.exit(0);
        System.out.println("UNREACHABLE after exit");
    }
}
