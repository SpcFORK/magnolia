import java.awt.Color;

// Owned probe: calls the unchanged package-private DS_BlobsMenu._mth4b
// bytecode. JDK resolves these synthetic parameters via this override;
// Magnolia injects the SAME values via the explicit Applet parameter map.
// This NEVER calls DS_BlobsMenu.run/start or draws its blobs.
public final class DSBlobsColorMain extends DS_BlobsMenu {
    public String getParameter(String key) {
        if ("good".equals(key)) return "00ff80";
        if ("plus".equals(key)) return "+0ABcd";
        if ("badhex".equals(key)) return "ffx000";
        if ("wronglen".equals(key)) return "ffffff0";
        return null;
    }
    public void init() {
        Color fallback = new Color(0x010203);
        Color good = _mth4b("good", fallback);
        Color plus = _mth4b("plus", fallback);
        Color badhex = _mth4b("badhex", fallback);
        Color wronglen = _mth4b("wronglen", fallback);
        Color absent = _mth4b("absent", fallback);
        System.out.println("color|" + (good.getRGB() & 0xffffff) + "|" +
                (plus.getRGB() & 0xffffff));
        System.out.println("fallback|" + (badhex == fallback ? "true" : "false") + "|" +
                (wronglen == fallback ? "true" : "false") + "|" +
                (absent == fallback ? "true" : "false"));
    }
    public static void main(String[] args) { new DSBlobsColorMain().init(); }
}
