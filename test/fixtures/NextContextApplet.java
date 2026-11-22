// Owned Java 8 Applet exercising ONLY a record-only AppletContext and safe
// base-relative local URL requests. This is not browser navigation or I/O.
import java.applet.Applet;
import java.applet.AppletContext;
import java.net.URL;

public class NextContextApplet extends Applet {
    public void probe() throws Exception {
        AppletContext context = getAppletContext();
        System.out.println(context != null && context == getAppletContext());
        context.showDocument(new URL(getDocumentBase(), "page2.html"), "_self");
        context.showDocument(new URL(getCodeBase(), "movie.bin"), "_blank");
        System.out.println("recorded");
    }

    // Selected independently for generated-Java pre-emission refusal.
    public void contextOnly() {
        System.out.println(getAppletContext() != null);
    }
    public void showOnly(AppletContext context, URL url, String target) {
        context.showDocument(url, target);
    }

    // Both members run on the JDK but are intentionally outside the record
    // profile: no Applet enumeration or arbitrary named browser target.
    public void appletLookup() {
        System.out.println(getAppletContext().getApplet("other") == null);
    }
    public void unknownTarget() throws Exception {
        getAppletContext().showDocument(
            new URL(getDocumentBase(), "page2.html"), "_new");
        System.out.println("JDK-accepted-target");
    }
    public void remoteURL() throws Exception {
        getAppletContext().showDocument(new URL("https://example.test/"), "_self");
        System.out.println("JDK-accepted-URL");
    }
    public void tooManyDocuments() throws Exception {
        AppletContext context = getAppletContext();
        URL url = new URL(getDocumentBase(), "page2.html");
        for (int i = 0; i < 17; i++) context.showDocument(url, "_self");
        System.out.println("JDK-17-records");
    }
}

class NextContextOverrideApplet extends NextContextApplet {
    @Override public AppletContext getAppletContext() {
        System.out.println("override-context");
        return null;
    }
}
