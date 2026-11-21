// Independent JDK 11 -Xverify:all oracle for the externally pinned original
// Cyberster MyApplet.class. This records Graphics state at original bytecode
// draw calls; the companion PNG is a JDK-only raster reference, NOT a claim
// that Magnolia's SVG/Canvas text raster has identical font metrics or pixels.
// Compile against the hash-checked class, with an EMPTY -sourcepath directory.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.awt.Color;
import java.awt.Font;
import java.awt.FontMetrics;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Rectangle;
import java.awt.Shape;
import java.awt.image.BufferedImage;
import java.awt.image.ImageObserver;
import java.io.File;
import java.net.URL;
import java.text.AttributedCharacterIterator;
import javax.imageio.ImageIO;

public final class MyAppletPaintJdk {
    private static String color(Color c) {
        return "rgb(" + c.getRed() + "," + c.getGreen() + "," + c.getBlue() + ")";
    }
    private static String font(Font f) {
        return f.getName() + "|" + f.getStyle() + "|" + f.getSize();
    }
    private static final class Stub implements AppletStub {
        private boolean active;
        public boolean isActive() { return active; }
        void setActive(boolean value) { active = value; }
        public URL getDocumentBase() { return getCodeBase(); }
        public URL getCodeBase() {
            try { return new URL("file:/"); }
            catch (Exception e) { throw new AssertionError(e); }
        }
        public String getParameter(String name) {
            if (name.equals("var1")) return "first";
            if (name.equals("var2")) return "second";
            return null;
        }
        public AppletContext getAppletContext() { return null; }
        public void appletResize(int width, int height) { }
    }
    private static final class RecordingGraphics extends Graphics {
        private final Graphics2D g;
        RecordingGraphics(Graphics2D graphics) { g = graphics; }
        public Graphics create() { return new RecordingGraphics((Graphics2D) g.create()); }
        public void translate(int x, int y) { g.translate(x, y); }
        public Color getColor() { return g.getColor(); }
        public void setColor(Color c) {
            g.setColor(c);
            System.out.println("setColor|" + color(g.getColor()));
        }
        public void setPaintMode() { g.setPaintMode(); }
        public void setXORMode(Color c) { g.setXORMode(c); }
        public Font getFont() { return g.getFont(); }
        public void setFont(Font f) {
            g.setFont(f);
            System.out.println("setFont|" + font(g.getFont()));
        }
        public FontMetrics getFontMetrics(Font f) { return g.getFontMetrics(f); }
        public Rectangle getClipBounds() { return g.getClipBounds(); }
        public void clipRect(int x, int y, int w, int h) { g.clipRect(x, y, w, h); }
        public void setClip(int x, int y, int w, int h) { g.setClip(x, y, w, h); }
        public Shape getClip() { return g.getClip(); }
        public void setClip(Shape s) { g.setClip(s); }
        public void copyArea(int x, int y, int w, int h, int dx, int dy) {
            g.copyArea(x, y, w, h, dx, dy);
        }
        public void drawLine(int a, int b, int c, int d) { g.drawLine(a, b, c, d); }
        public void fillRect(int a, int b, int c, int d) { g.fillRect(a, b, c, d); }
        public void clearRect(int a, int b, int c, int d) { g.clearRect(a, b, c, d); }
        public void drawRoundRect(int a, int b, int c, int d, int e, int f) {
            g.drawRoundRect(a, b, c, d, e, f);
        }
        public void fillRoundRect(int a, int b, int c, int d, int e, int f) {
            g.fillRoundRect(a, b, c, d, e, f);
        }
        public void drawOval(int a, int b, int c, int d) { g.drawOval(a, b, c, d); }
        public void fillOval(int a, int b, int c, int d) { g.fillOval(a, b, c, d); }
        public void drawArc(int a, int b, int c, int d, int e, int f) {
            g.drawArc(a, b, c, d, e, f);
        }
        public void fillArc(int a, int b, int c, int d, int e, int f) {
            g.fillArc(a, b, c, d, e, f);
        }
        public void drawPolyline(int[] x, int[] y, int n) { g.drawPolyline(x, y, n); }
        public void drawPolygon(int[] x, int[] y, int n) { g.drawPolygon(x, y, n); }
        public void fillPolygon(int[] x, int[] y, int n) { g.fillPolygon(x, y, n); }
        public void drawString(String s, int x, int y) {
            System.out.println("draw|" + s + "|" + x + "|" + y + "|" +
                color(g.getColor()) + "|" + font(g.getFont()));
            g.drawString(s, x, y);
        }
        public void drawString(AttributedCharacterIterator it, int x, int y) {
            g.drawString(it, x, y);
        }
        public boolean drawImage(java.awt.Image img, int x, int y, ImageObserver o) {
            return g.drawImage(img, x, y, o);
        }
        public boolean drawImage(java.awt.Image img, int x, int y, int w, int h, ImageObserver o) {
            return g.drawImage(img, x, y, w, h, o);
        }
        public boolean drawImage(java.awt.Image img, int x, int y, Color bg, ImageObserver o) {
            return g.drawImage(img, x, y, bg, o);
        }
        public boolean drawImage(java.awt.Image img, int x, int y, int w, int h, Color bg,
                ImageObserver o) { return g.drawImage(img, x, y, w, h, bg, o); }
        public boolean drawImage(java.awt.Image img, int dx1, int dy1, int dx2, int dy2,
                int sx1, int sy1, int sx2, int sy2, ImageObserver o) {
            return g.drawImage(img, dx1, dy1, dx2, dy2, sx1, sy1, sx2, sy2, o);
        }
        public boolean drawImage(java.awt.Image img, int dx1, int dy1, int dx2, int dy2,
                int sx1, int sy1, int sx2, int sy2, Color bg, ImageObserver o) {
            return g.drawImage(img, dx1, dy1, dx2, dy2, sx1, sy1, sx2, sy2, bg, o);
        }
        public void dispose() { g.dispose(); }
    }
    public static void main(String[] args) throws Exception {
        if (args.length != 1) throw new IllegalArgumentException("PNG output path required");
        MyApplet applet = new MyApplet();
        Stub stub = new Stub();
        applet.setStub(stub);
        applet.init();
        stub.setActive(true); // AppletStub says active just before start.
        applet.start();
        if (!Color.PINK.equals(applet.getBackground()) ||
                !Color.BLUE.equals(applet.getForeground()))
            throw new AssertionError("original init colors changed");
        BufferedImage image = new BufferedImage(320, 200, BufferedImage.TYPE_INT_ARGB);
        Graphics2D graphics = image.createGraphics();
        graphics.setColor(applet.getForeground());
        graphics.setFont(new Font("Dialog", Font.PLAIN, 12));
        RecordingGraphics recorded = new RecordingGraphics(graphics);
        try { applet.paint(recorded); }
        finally { recorded.dispose(); }
        if (!ImageIO.write(image, "PNG", new File(args[0])))
            throw new AssertionError("PNG writer unavailable");
        stub.setActive(false); // Inactive just before stop.
        applet.stop();
        applet.destroy();
    }
}
