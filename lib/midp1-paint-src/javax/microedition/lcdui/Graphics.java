package javax.microedition.lcdui;

// Opt-in Canvas paint surface plus the existing bounded offscreen methods.
// The native host owns the pixels and a one-frame 176x176 clip; this is not
// a general LCDUI renderer, font rasterizer, affine transform or AWT Graphics.
public final class Graphics {
    private Graphics() { }

    public void setColor(int rgb) { ImageBridge.setColor(this, rgb); }
    public int getColor() { return ImageBridge.getColor(this); }
    public int getClipWidth() { return ImageBridge.clipWidth(this); }
    public void fillRect(int x, int y, int width, int height) {
        ImageBridge.fillRect(this, x, y, width, height);
    }
    public void drawImage(Image image, int x, int y, int anchor) {
        ImageBridge.drawImage(this, image, x, y, anchor);
    }
}
