// Owned JDK 11 AppletStub reference. Constructs URLs but never opens them.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.net.URL;

public final class AppletBaseHostJdk {
    private static URL base(String host, String path) throws Exception {
        return host.isEmpty() ? new URL("file:/" + path) :
            new URL("http://" + host + "/" + path);
    }
    private static final class Stub implements AppletStub {
        final URL document, code;
        Stub(URL document, URL code) { this.document = document; this.code = code; }
        public boolean isActive() { return false; }
        public URL getDocumentBase() { return document; }
        public URL getCodeBase() { return code; }
        public String getParameter(String key) { return null; }
        public AppletContext getAppletContext() { return null; }
        public void appletResize(int w, int h) { throw new AssertionError("unexpected resize"); }
    }
    public static void main(String[] args) throws Exception {
        if (args.length != 2) throw new IllegalArgumentException("DOCUMENT_HOST CODE_HOST");
        AppletBaseHostProbe applet = new AppletBaseHostProbe();
        applet.setStub(new Stub(base(args[0], "page.html"), base(args[1], "assets/")));
        applet.probe();
    }
}
