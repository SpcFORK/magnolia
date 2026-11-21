package javax.microedition.lcdui;

import java.io.IOException;

// javac-only type stub for owned Image/Graphics bytecode. NEVER ship this
// classfile or execute its native declarations. The Oak VM admits exactly the
// bridge descriptors and only under the explicitly selected MIDlet profile.
final class ImageBridge {
    private ImageBridge() { }
    static native Image blank(int width, int height);
    static native Image load(String resourceName) throws IOException;
    static native Graphics graphics(Image image);
    static native void getRGB(Image image, int[] rgb, int offset,
                              int scanlength, int x, int y,
                              int width, int height);
    static native void setColor(Graphics graphics, int rgb);
    static native int getColor(Graphics graphics);
    // Only the opt-in one-frame paint model calls this geometry bridge.
    static native int clipWidth(Graphics graphics);
    static native void fillRect(Graphics graphics, int x, int y,
                                int width, int height);
    static native void drawImage(Graphics graphics, Image source,
                                 int x, int y, int anchor);
}
