// Reference invocation of the exact owned .class under a real JDK 11.
// It must exit 1, retain guest-prefix on stdout and diagnose the real Java
// IllegalArgumentException on stderr.
public final class AppletOutputBoundaryJdk {
    public static void main(String[] args) {
        new AppletOutputBoundaryMain().init();
    }
}
