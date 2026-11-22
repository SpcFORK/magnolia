// Owned Java 8 fixture: URL host metadata only. No URL is opened.
import java.applet.Applet;
import java.net.URL;

public final class AppletBaseHostProbe extends Applet {
    public void probe() throws Exception {
        URL document = ((Applet) this).getDocumentBase();
        URL code = ((Applet) this).getCodeBase();
        System.out.println("document");
        System.out.println(document.getHost());
        System.out.println("document-relative");
        System.out.println(new URL(document, "movie.bin").getHost());
        System.out.println("code");
        System.out.println(code.getHost());
        System.out.println("code-relative");
        System.out.println(new URL(code, "movie.bin").getHost());
        if (document.getHost() == document.getHost()) {
            System.out.println("same-host-reference");
        } else {
            System.out.println("different-host-reference");
        }
    }
}
