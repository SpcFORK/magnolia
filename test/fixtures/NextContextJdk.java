// JDK 11 -Xverify:all AppletStub/AppletContext oracle. The stub RECORDS
// showDocument requests; it never opens a browser, network or local file.
import java.applet.Applet;
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.applet.AudioClip;
import java.awt.Image;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Collections;
import java.util.Enumeration;
import java.util.Iterator;
import java.util.List;

public final class NextContextJdk {
    static final class Recorder implements AppletContext {
        final List<String> records = new ArrayList<String>();
        final String documentPrefix, codePrefix;
        Recorder(URL document, URL code) {
            String doc = document.toExternalForm();
            documentPrefix = doc.substring(0, doc.lastIndexOf('/') + 1);
            codePrefix = code.toExternalForm();
            if (!codePrefix.endsWith("/")) throw new AssertionError("code base");
        }
        public AudioClip getAudioClip(URL url) { throw new AssertionError("audio"); }
        public Image getImage(URL url) { throw new AssertionError("image"); }
        public Applet getApplet(String name) { return null; }
        public Enumeration<Applet> getApplets() {
            return Collections.enumeration(Collections.<Applet>emptyList());
        }
        public void showDocument(URL url) { throw new AssertionError("one-arg navigation"); }
        public void showDocument(URL url, String target) {
            String text = url.toExternalForm();
            String base = text.startsWith(documentPrefix) ? "documentbase:" :
                text.startsWith(codePrefix) ? "codebase:" : "external:";
            String name = base.equals("documentbase:")
                ? text.substring(documentPrefix.length()) :
                base.equals("codebase:") ? text.substring(codePrefix.length()) : text;
            records.add(base + name + ":" + target);
        }
        public void showStatus(String status) { throw new AssertionError("status"); }
        public void setStream(String key, InputStream value) throws IOException {
            throw new AssertionError("stream");
        }
        public InputStream getStream(String key) { throw new AssertionError("stream"); }
        public Iterator<String> getStreamKeys() { return Collections.<String>emptyList().iterator(); }
    }
    static final class Stub implements AppletStub {
        final URL document, code;
        final Recorder recorder;
        Stub(URL document, URL code) {
            this.document = document;
            this.code = code;
            recorder = new Recorder(document, code);
        }
        public boolean isActive() { return true; }
        public URL getDocumentBase() { return document; }
        public URL getCodeBase() { return code; }
        public String getParameter(String key) { return null; }
        public AppletContext getAppletContext() { return recorder; }
        public void appletResize(int width, int height) { throw new AssertionError("resize"); }
    }
    public static void main(String[] args) throws Exception {
        if (args.length != 2) throw new IllegalArgumentException("mode base-directory");
        Path root = Paths.get(args[1]).toAbsolutePath();
        Files.createDirectories(root.resolve("document"));
        Files.createDirectories(root.resolve("assets"));
        URL document = root.resolve("document/page.html").toUri().toURL();
        URL code = root.resolve("assets/").toUri().toURL();
        Stub stub = new Stub(document, code);
        boolean override = args[0].equals("override");
        NextContextApplet applet = override
            ? new NextContextOverrideApplet() : new NextContextApplet();
        applet.setStub(stub);
        if (!override && (applet.getAppletContext() != stub.recorder ||
            applet.getDocumentBase() != stub.document ||
            applet.getCodeBase() != stub.code))
            throw new AssertionError("the JDK must delegate to the exact stub");
        if (args[0].equals("probe")) applet.probe();
        else if (args[0].equals("lookup")) applet.appletLookup();
        else if (args[0].equals("target")) applet.unknownTarget();
        else if (args[0].equals("remote")) applet.remoteURL();
        else if (args[0].equals("bound")) applet.tooManyDocuments();
        else if (override) applet.contextOnly();
        else throw new IllegalArgumentException(args[0]);
        for (String record : stub.recorder.records) System.out.println(record);
    }
}
