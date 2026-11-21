// Independent JDK 11 -Xverify:all/Xvfb oracle for the UNCHANGED original
// ArcCanvas.paint(Graphics) and the separately compiled harness Applet.
// Records calls before rasterization; neither this trace nor Magnolia's
// SVG/HTML output establishes arc pixels, clipping or font/glyph parity.
// Compile against the pinned originals with an EMPTY -sourcepath directory.
import java.awt.Color;
import java.awt.Font;
import java.awt.FontMetrics;
import java.awt.Graphics;
import java.awt.Graphics2D;
import java.awt.Rectangle;
import java.awt.Shape;
import java.awt.image.BufferedImage;
import java.awt.image.ImageObserver;
import java.text.AttributedCharacterIterator;

public final class ArcCanvasPaintJdk {
    private static String color(Color value) {
        return "rgb(" + value.getRed() + "," + value.getGreen() + "," +
            value.getBlue() + ")";
    }
    private static String font(Font value) {
        return value.getName() + "|" + value.getStyle() + "|" + value.getSize();
    }
    private static final class RecordingGraphics extends Graphics {
        private final Graphics2D delegate;
        private final StringBuilder calls = new StringBuilder();
        private final StringBuilder stateChanges = new StringBuilder();
        RecordingGraphics(Graphics2D delegate) { this.delegate = delegate; }
        String trace() { return calls.toString(); }
        String mutations() { return stateChanges.toString(); }
        private AssertionError unexpected() { return new AssertionError("unexpected Graphics API"); }
        public Graphics create() { throw unexpected(); }
        public void translate(int x, int y) { throw unexpected(); }
        public Color getColor() { return delegate.getColor(); }
        public void setColor(Color value) {
            delegate.setColor(value);
            stateChanges.append("setColor|").append(color(getColor())).append('\n');
        }
        public void setPaintMode() { throw unexpected(); }
        public void setXORMode(Color value) { throw unexpected(); }
        public Font getFont() { return delegate.getFont(); }
        public void setFont(Font value) {
            delegate.setFont(value);
            stateChanges.append("setFont|").append(value == null ? "null" : font(value))
                .append("|effective|").append(font(getFont())).append('\n');
        }
        public FontMetrics getFontMetrics(Font value) { throw unexpected(); }
        public Rectangle getClipBounds() { throw unexpected(); }
        public void clipRect(int x, int y, int width, int height) { throw unexpected(); }
        public void setClip(int x, int y, int width, int height) { throw unexpected(); }
        public Shape getClip() { throw unexpected(); }
        public void setClip(Shape clip) { throw unexpected(); }
        public void copyArea(int x, int y, int width, int height, int dx, int dy) {
            throw unexpected();
        }
        public void drawLine(int x1, int y1, int x2, int y2) {
            calls.append("line|").append(x1).append('|').append(y1).append('|')
                .append(x2).append('|').append(y2).append('|')
                .append(color(getColor())).append('\n');
        }
        public void fillRect(int x, int y, int width, int height) { throw unexpected(); }
        public void clearRect(int x, int y, int width, int height) { throw unexpected(); }
        public void drawRoundRect(int x, int y, int width, int height,
                                  int arcWidth, int arcHeight) { throw unexpected(); }
        public void fillRoundRect(int x, int y, int width, int height,
                                  int arcWidth, int arcHeight) { throw unexpected(); }
        public void drawOval(int x, int y, int width, int height) { throw unexpected(); }
        public void fillOval(int x, int y, int width, int height) { throw unexpected(); }
        private void arc(String type, int x, int y, int width, int height,
                         int start, int extent) {
            calls.append(type).append('|').append(x).append('|').append(y)
                .append('|').append(width).append('|').append(height).append('|')
                .append(start).append('|').append(extent).append('|')
                .append(color(getColor())).append('\n');
        }
        public void drawArc(int x, int y, int width, int height, int start, int extent) {
            arc("drawArc", x, y, width, height, start, extent);
        }
        public void fillArc(int x, int y, int width, int height, int start, int extent) {
            arc("fillArc", x, y, width, height, start, extent);
        }
        public void drawPolyline(int[] x, int[] y, int count) { throw unexpected(); }
        public void drawPolygon(int[] x, int[] y, int count) { throw unexpected(); }
        public void fillPolygon(int[] x, int[] y, int count) { throw unexpected(); }
        public void drawString(String text, int x, int y) {
            calls.append("text|").append(text).append('|').append(x).append('|')
                .append(y).append('|').append(color(getColor())).append('|')
                .append(font(getFont())).append('\n');
        }
        public void drawString(AttributedCharacterIterator text, int x, int y) {
            throw unexpected();
        }
        public boolean drawImage(java.awt.Image image, int x, int y,
                                 ImageObserver observer) { throw unexpected(); }
        public boolean drawImage(java.awt.Image image, int x, int y,
                                 int width, int height, ImageObserver observer) {
            throw unexpected();
        }
        public boolean drawImage(java.awt.Image image, int x, int y, Color background,
                                 ImageObserver observer) { throw unexpected(); }
        public boolean drawImage(java.awt.Image image, int x, int y,
                                 int width, int height, Color background,
                                 ImageObserver observer) { throw unexpected(); }
        public boolean drawImage(java.awt.Image image, int dx1, int dy1,
                                 int dx2, int dy2, int sx1, int sy1,
                                 int sx2, int sy2, ImageObserver observer) {
            throw unexpected();
        }
        public boolean drawImage(java.awt.Image image, int dx1, int dy1,
                                 int dx2, int dy2, int sx1, int sy1,
                                 int sx2, int sy2, Color background,
                                 ImageObserver observer) { throw unexpected(); }
        public void dispose() { delegate.dispose(); }
    }
    public static void main(String[] args) {
        if (args.length != 0) throw new IllegalArgumentException("no arguments expected");
        // The tiny image supplies REAL JDK Graphics color/font semantics.
        // The recorder does not draw into it; it records the original calls.
        Graphics2D graphics = new BufferedImage(1, 1,
            BufferedImage.TYPE_INT_ARGB).createGraphics();
        graphics.setColor(Color.black);
        graphics.setFont(new Font("Dialog", Font.PLAIN, 12));
        RecordingGraphics recorded = new RecordingGraphics(graphics);
        try { new ArcCanvasPaintApplet().paint(recorded); }
        finally { recorded.dispose(); }
        System.out.print(recorded.trace());
        System.err.print(recorded.mutations()); // separately pinned JDK-only behavior
    }
}
