// JDK 11 -Xverify:all reference for BulkStreamApplet's *actual* file reads.
// No network or Flash bytecode: a real AppletStub uses separate local bases.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.net.URL;
import java.nio.file.Path;
import java.nio.file.Paths;

public final class BulkStreamJdk {
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
            throw new AssertionError("unexpected Applet resize");
        }
    }

    public static void main(String[] args) throws Exception {
        if (args.length < 2 || args.length > 3) {
            throw new IllegalArgumentException("PAGE CODE_DIR [MODE]");
        }
        Path page = Paths.get(args[0]).toAbsolutePath();
        Path assets = Paths.get(args[1]).toAbsolutePath();
        Stub stub = new Stub(page.toUri().toURL(), assets.toUri().toURL());
        if (stub.document.equals(stub.code) ||
            !stub.document.toExternalForm().endsWith("/page.html") ||
            !stub.code.toExternalForm().endsWith("/assets/")) {
            throw new AssertionError("separate local page/code directories required");
        }
        BulkStreamApplet applet = new BulkStreamApplet();
        applet.setStub(stub);
        String mode = args.length == 3 ? args[2] : "probe";
        switch (mode) {
            case "probe": applet.probe(); break;
            case "wrappedZero": applet.wrappedZero(); break;
            case "overlap": applet.overlap(); break;
            case "budgetRead": applet.budgetRead(); break;
            case "closedResource": applet.closedResource(); break;
            case "unmodeledSubtype": applet.unmodeledSubtype(); break;
            case "unsupportedOverload": applet.unsupportedOverload(); break;
            case "unsupportedOwner": applet.unsupportedOwner(); break;
            default: throw new IllegalArgumentException(mode);
        }
    }
}
