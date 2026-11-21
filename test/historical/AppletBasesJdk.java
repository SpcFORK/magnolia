// JDK 11 -Xverify:all AppletStub reference for the owned Java 8 fixture.
// Both bases are file: URLs in separate temporary directories; no network.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.net.URL;
import java.nio.file.Path;
import java.nio.file.Paths;

public final class AppletBasesJdk {
    private static final class Stub implements AppletStub {
        final URL document, code;
        Stub(URL document, URL code) {
            this.document = document;
            this.code = code;
        }
        public boolean isActive() { return false; }
        public URL getDocumentBase() { return document; }
        public URL getCodeBase() { return code; }
        public String getParameter(String key) { return null; }
        public AppletContext getAppletContext() { return null; }
        public void appletResize(int width, int height) {
            throw new AssertionError("unexpected resize");
        }
    }

    public static void main(String[] args) throws Exception {
        if (args.length != 2) throw new IllegalArgumentException("DOCUMENT_PAGE CODE_DIRECTORY");
        Path page = Paths.get(args[0]).toAbsolutePath();
        Path assets = Paths.get(args[1]).toAbsolutePath();
        Stub stub = new Stub(page.toUri().toURL(), assets.toUri().toURL());
        if (stub.document.equals(stub.code) ||
                !stub.document.toExternalForm().endsWith("/page.html") ||
                !stub.code.toExternalForm().endsWith("/assets/")) {
            throw new AssertionError("bases must be a distinct page and code directory");
        }
        AppletBaseResource applet = new AppletBaseResource();
        applet.setStub(stub);
        if (!applet.getDocumentBase().equals(stub.document) ||
                !applet.getCodeBase().equals(stub.code)) {
            throw new AssertionError("Applet did not delegate to the AppletStub");
        }
        applet.probe();
    }
}
