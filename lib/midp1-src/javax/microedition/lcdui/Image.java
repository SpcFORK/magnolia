package javax.microedition.lcdui;

import java.io.IOException;

// Owned MIDP-1 bytecode for the deliberately limited native Image island.
// ImageBridge is an internal, descriptor-checked VM service, not part of the
// Java ME API and not a class taken from MicroEmulator. Its build-time stub is
// NOT shipped in midp1-classes. The native host alone owns the bounded raster.
public final class Image {
    private int width;
    private int height;
    private boolean mutable;

    private Image() { }

    public static Image createImage(int width, int height) {
        return ImageBridge.blank(width, height);
    }

    public static Image createImage(String resourceName) throws IOException {
        if (resourceName == null) throw new NullPointerException();
        return ImageBridge.load(resourceName);
    }

    public int getWidth() { return width; }
    public int getHeight() { return height; }

    public Graphics getGraphics() {
        if (!mutable) throw new IllegalStateException();
        return ImageBridge.graphics(this);
    }

    public void getRGB(int[] rgb, int offset, int scanlength,
                       int x, int y, int width, int height) {
        ImageBridge.getRGB(this, rgb, offset, scanlength,
                           x, y, width, height);
    }
}
