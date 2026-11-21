package javax.microedition.lcdui;

// DEFAULT native offscreen raster surface, NOT a Canvas paint callback or
// device renderer. The separate opt-in one-frame variant is in
// lib/midp1-paint-src. Only top-left binary-alpha Image blits join the measured
// color/fill island; other anchors and drawing/text APIs remain unmodeled.
public final class Graphics {
    private Graphics() { }

    public void setColor(int rgb) { ImageBridge.setColor(this, rgb); }
    public int getColor() { return ImageBridge.getColor(this); }
    public void fillRect(int x, int y, int width, int height) {
        ImageBridge.fillRect(this, x, y, width, height);
    }
    public void drawImage(Image image, int x, int y, int anchor) {
        ImageBridge.drawImage(this, image, x, y, anchor);
    }
}
