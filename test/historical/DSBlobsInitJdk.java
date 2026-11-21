// Independent JDK 11 -Xverify:all/Xvfb reference for the ORIGINAL
// DS_BlobsMenu.init(), not its run/start/paint. Credits are supplied exactly
// as requested by the historic bytecode; absent/mismatched credits loop.
// The one logical Helvetica BOLD 12 width table is also measured here.
import java.applet.Applet;
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.applet.AudioClip;
import java.awt.Font;
import java.awt.FontMetrics;
import java.awt.Image;
import java.net.URL;
import java.util.Enumeration;
import java.util.Collections;

public final class DSBlobsInitJdk {
    private static final class Context implements AppletContext {
        String status;
        public AudioClip getAudioClip(URL url) { throw new AssertionError("audio host"); }
        public Image getImage(URL url) { throw new AssertionError("image host"); }
        public Applet getApplet(String name) { throw new AssertionError("applet registry"); }
        public Enumeration<Applet> getApplets() {
            return Collections.enumeration(Collections.<Applet>emptyList());
        }
        public void showDocument(URL url) { throw new AssertionError("navigation host"); }
        public void showDocument(URL url, String target) { throw new AssertionError("navigation host"); }
        public void showStatus(String text) { status = text; }
        public void setStream(String key, java.io.InputStream stream) {
            throw new AssertionError("stream host");
        }
        public java.io.InputStream getStream(String key) { throw new AssertionError("stream host"); }
        public java.util.Iterator<String> getStreamKeys() { throw new AssertionError("stream host"); }
    }
    private static final class Stub implements AppletStub {
        final Context context = new Context();
        public boolean isActive() { return false; }
        public URL getDocumentBase() { return getCodeBase(); }
        public URL getCodeBase() {
            try { return new URL("file:/"); }
            catch (Exception e) { throw new AssertionError(e); }
        }
        public String getParameter(String name) {
            return "credits".equals(name) ?
                "Applet by Dario Sciacca (www.dseffects.com)" : null;
        }
        public AppletContext getAppletContext() { return context; }
        public void appletResize(int width, int height) {
            throw new AssertionError("unexpected resize");
        }
    }
    public static void main(String[] args) {
        if (args.length != 1) throw new IllegalArgumentException("state or widths");
        DS_BlobsMenu applet = new DS_BlobsMenu();
        Stub stub = new Stub();
        applet.setSize(320, 200);
        applet.setStub(stub);
        if ("state".equals(args[0])) {
            applet.init(); // unchanged original bytecode
            Font f = applet._fld7;
            FontMetrics metrics = applet.getFontMetrics(f);
            System.out.println("state|" + applet._fld8c + "|" +
                applet._fld8 + "|" + applet._fld9 + "|" +
                applet._fld9b + "|" + stub.context.status + "|" +
                f.getName() + "|" + f.getStyle() + "|" + f.getSize());
            System.out.println("sourceWidth|" + metrics.stringWidth("www.dseffects.com"));
        } else if ("widths".equals(args[0])) {
            Font f = applet._fld7;
            FontMetrics metrics = applet.getFontMetrics(f);
            System.out.println("font|" + f.getName() + "|" + f.getStyle() + "|" + f.getSize());
            System.out.print("ascii32-126|");
            for (int code = 32; code <= 126; code++) {
                if (code > 32) System.out.print(",");
                System.out.print(metrics.stringWidth(Character.toString((char) code)));
            }
            System.out.println();
            for (String text : new String[]{"", "www.dseffects.com", "HELLO WORLD!",
                    "AV VA", "Java 8: 0123456789", "  '[]~{}"}) {
                int sum = 0;
                for (int i = 0; i < text.length(); i++)
                    sum += metrics.stringWidth(text.substring(i, i + 1));
                System.out.println("sample|" + text + "|" + metrics.stringWidth(text) + "|" + sum);
                if (sum != metrics.stringWidth(text))
                    throw new AssertionError("non-additive JDK glyph widths");
            }
        } else throw new IllegalArgumentException(args[0]);
    }
}
