// JDK-valid Applet. Even with --allow-system-exit-zero, the bounded
// application-only termination capability must reject an Applet before init.
public final class SystemExitApplet extends java.applet.Applet {
    public void init() {
        System.out.println("UNREACHABLE applet init");
        System.exit(0);
    }
}
