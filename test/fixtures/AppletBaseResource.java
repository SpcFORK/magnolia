// Owned Java 8 fixture. The JDK AppletStub oracle is in
// test/historical/AppletBasesJdk.java. This is NOT the Flash player.
import java.applet.Applet;
import java.io.DataInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;

public final class AppletBaseResource extends Applet {
    private void readBase(String label, URL base, String relative)
            throws IOException {
        InputStream stream = new URL(base, relative).openStream();
        DataInputStream data = new DataInputStream(stream);
        byte[] bytes = new byte[2];
        int count = data.read(bytes, 0, 2);
        stream.close();
        System.out.println(label);
        System.out.println(count);
        System.out.println(bytes[0]);
        System.out.println(bytes[1]);
    }

    public void probe() throws IOException {
        readBase("document", ((Applet) this).getDocumentBase(), "movie.bin");
        readBase("code", ((Applet) this).getCodeBase(), "movie.bin");
    }

    // Selected individually to check missing/unsafe input without opening any
    // file, contacting a host, or accidentally using the other base's map.
    public void document(String relative) throws IOException {
        readBase("document", ((Applet) this).getDocumentBase(), relative);
    }
    public void code(String relative) throws IOException {
        readBase("code", ((Applet) this).getCodeBase(), relative);
    }
}
