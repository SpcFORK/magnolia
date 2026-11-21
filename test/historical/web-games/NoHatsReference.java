import java.applet.Applet;
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.applet.AudioClip;
import java.awt.Frame;
import java.awt.Graphics2D;
import java.awt.Image;
import java.awt.Toolkit;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.InputStream;
import java.net.URL;
import java.security.MessageDigest;
import java.util.Collections;
import java.util.Enumeration;
import java.util.HashSet;
import java.util.Set;

// JDK 11/Xvfb reference for the independently pinned original noHats.java
// recompiled unchanged to Java 8 bytecode, with its original, local GIF.
// Only an explicitly local file URL is granted; no browser, audio, HTTP or
// scoreboard service. Does NOT certify a Magnolia game run or input delivery.
public final class NoHatsReference {
    private static String hex(byte[] bytes) {
        StringBuilder b = new StringBuilder(bytes.length * 2);
        for (byte value : bytes) b.append(String.format("%02x", value & 255));
        return b.toString();
    }
    private static final class LocalHost implements AppletStub, AppletContext {
        final File imageFile;
        final URL base;
        LocalHost(File imageFile) throws Exception {
            this.imageFile = imageFile.getCanonicalFile();
            if (!this.imageFile.isFile() || this.imageFile.length() != 39546 ||
                !this.imageFile.getName().equals("noHats.gif"))
                throw new IllegalArgumentException("missing pinned local GIF");
            this.base = this.imageFile.getParentFile().toURI().toURL();
            MessageDigest md = MessageDigest.getInstance("SHA-256");
            try (java.io.FileInputStream in = new java.io.FileInputStream(this.imageFile)) {
                byte[] buffer = new byte[4096];
                for (int n; (n = in.read(buffer)) >= 0;) md.update(buffer, 0, n);
            }
            if (!hex(md.digest()).equals("9de9f3fb87d9219a9ca2a038d5e2d59cf4871e954cc15db29a1ca11144bfc0b9"))
                throw new IllegalArgumentException("local GIF SHA-256 differs");
        }
        public boolean isActive() { return true; }
        public URL getDocumentBase() { return base; }
        public URL getCodeBase() { return base; }
        public String getParameter(String name) { return null; }
        public AppletContext getAppletContext() { return this; }
        public void appletResize(int width, int height) {
            if (width != 512 || height != 240) throw new IllegalArgumentException("unexpected resize");
        }
        public AudioClip getAudioClip(URL url) { throw new SecurityException("no audio"); }
        public Image getImage(URL url) {
            try {
                if (!url.getProtocol().equals("file") ||
                    !new File(url.toURI()).getCanonicalFile().equals(imageFile))
                    throw new SecurityException("nonlocal or wrong image URL");
            } catch (java.net.URISyntaxException | java.io.IOException e) {
                throw new IllegalArgumentException("bad image URL", e);
            }
            return Toolkit.getDefaultToolkit().createImage(imageFile.getPath());
        }
        public Applet getApplet(String name) { return null; }
        public Enumeration<Applet> getApplets() { return Collections.enumeration(Collections.<Applet>emptyList()); }
        public void showDocument(URL url) { throw new SecurityException("no document navigation"); }
        public void showDocument(URL url, String target) { throw new SecurityException("no document navigation"); }
        public void showStatus(String status) { System.out.println("REF_STATUS:" + status); }
        public void setStream(String key, InputStream stream) { throw new SecurityException("no streams"); }
        public InputStream getStream(String key) { return null; }
        public java.util.Iterator<String> getStreamKeys() { return Collections.<String>emptyList().iterator(); }
    }
    public static void main(String[] args) throws Exception {
        if (args.length != 1) throw new IllegalArgumentException("one local GIF path required");
        LocalHost stub = new LocalHost(new File(args[0]));
        noHats game = new noHats();
        game.setStub(stub);
        Frame window = new Frame("pinned No Hats splash reference");
        window.setLayout(null);
        window.add(game);
        game.setBounds(0, 0, 512, 240);
        window.setSize(512, 240);
        window.setVisible(true);
        game.init();
        if (game.gameStatus != 0 || game.score != 0 || game.lives != 0)
            throw new AssertionError("original splash state differs");
        BufferedImage screen = new BufferedImage(512, 240, BufferedImage.TYPE_INT_ARGB);
        Graphics2D g = screen.createGraphics();
        game.update(g); // one synchronous original update/paint, no game thread
        g.dispose();
        int[] pixels = screen.getRGB(0, 0, 512, 240, null, 0, 512);
        byte[] argb = new byte[pixels.length * 4];
        Set<Integer> distinct = new HashSet<Integer>();
        for (int i = 0; i < pixels.length; i++) {
            int px = pixels[i];
            distinct.add(px);
            argb[i * 4] = (byte) (px >>> 24);
            argb[i * 4 + 1] = (byte) (px >>> 16);
            argb[i * 4 + 2] = (byte) (px >>> 8);
            argb[i * 4 + 3] = (byte) px;
        }
        System.out.println("REF_SPLASH:status=" + game.gameStatus +
                ";score=" + game.score + ";lives=" + game.lives +
                ";updates=" + game.animationCount + ";pixels=" + pixels.length +
                ";colors=" + distinct.size() +
                ";argbSha256=" + hex(MessageDigest.getInstance("SHA-256").digest(argb)));
        window.dispose();
        System.exit(0); // only this bounded harness, not the original class, exits
    }
}
