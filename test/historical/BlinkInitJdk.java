// Independent JDK 11 -Xverify:all oracle for *original* pinned Sun Blink.class.
// Requires a real display (Xvfb); the local AppletStub supplies only two HTML
// parameters. Invoke init ONLY, not start/run/stop/paint or an AWT event loop.
// Compile against the externally hash-checked .class with empty -sourcepath.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.net.URL;

public final class BlinkInitJdk {
    private static final class Stub implements AppletStub {
        private final String speed, label;
        Stub(String speed, String label) { this.speed = speed; this.label = label; }
        public boolean isActive() { return false; } // init precedes start.
        public URL getDocumentBase() { return getCodeBase(); }
        public URL getCodeBase() {
            try { return new URL("file:/"); }
            catch (Exception e) { throw new AssertionError(e); }
        }
        public String getParameter(String name) {
            if ("speed".equals(name)) return speed;
            if ("lbl".equals(name)) return label;
            return null;
        }
        public AppletContext getAppletContext() { return null; }
        public void appletResize(int width, int height) {
            throw new AssertionError("unexpected resize during init");
        }
    }
    public static void main(String[] args) {
        if (args.length != 2) throw new IllegalArgumentException("speed/label required");
        String speed = "~".equals(args[0]) ? null : args[0];
        String label = "~".equals(args[1]) ? null : args[1];
        Blink applet = new Blink();
        applet.setStub(new Stub(speed, label));
        try {
            applet.init();
            System.out.println("init|" + applet.font.getName() + "|" +
                applet.font.getStyle() + "|" + applet.font.getSize() + "|" +
                applet.speed + "|" + applet.lbl);
        } catch (NumberFormatException | ArithmeticException error) {
            System.out.println("error|" + error.getClass().getName());
        }
    }
}
