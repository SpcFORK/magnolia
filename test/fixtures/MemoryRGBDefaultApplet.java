// Owned Java 8 bytecode: the five-argument MemoryImageSource int[] overload.
// The runtime only claims finite, detached, synchronous ARGB snapshots, not
// a display peer, consumer callbacks, animation, or a DS worker frame.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Component;
import java.awt.Graphics;
import java.awt.Image;
import java.awt.image.ColorModel;
import java.awt.image.MemoryImageSource;

public class MemoryRGBDefaultApplet extends Applet {
    MemoryImageSource producer;
    MemoryImageSource explicitProducer;
    Image source;
    Image explicitSource;
    Image canvas;

    @Override public void init() {
        ColorModel model = ColorModel.getRGBdefault();
        int[] pixels = new int[16];
        pixels[5] = 0xff102030;
        pixels[6] = 0x00112233;
        pixels[9] = 0xff00ff80;
        pixels[10] = 0xff801020;
        producer = new MemoryImageSource(2, 2, pixels, 5, 4);
        pixels[5] = 0xff405060; // Live array BEFORE createImage.
        source = ((Component) this).createImage(producer);
        pixels[5] = 0xff708090; // After createImage: the first Image is fixed.
        explicitProducer = new MemoryImageSource(2, 2, model, pixels, 5, 4);
        explicitSource = ((Component) this).createImage(explicitProducer);
        pixels[5] = 0xffa0b0c0; // Neither already-created Image changes.
        canvas = ((Component) this).createImage(4, 3);
        Graphics graphics = canvas.getGraphics();
        graphics.setColor(Color.blue);
        graphics.fillRect(0, 0, 4, 3);
        System.out.println(graphics.drawImage(source, 1, 0, this));
        System.out.println("default-argb-ready");
    }

    // Another selected method: Image width/height queries on the same
    // bounded offscreen producer, both with and without an ImageObserver.
    // BufferedImage.flush() must not destroy the already-materialized pixels.
    public void boundsOnly() {
        int[] pixels = {0xff112233, 0x00112233, 0xff405060,
                        0xff708090, 0xffa0b0c0, 0xffe0d0c0};
        source = ((Component) this).createImage(
            new MemoryImageSource(2, 3, pixels, 0, 2));
        System.out.println(source.getWidth(null));
        System.out.println(source.getHeight(this));
        source.flush();
        System.out.println(source.getWidth(this));
        System.out.println(source.getHeight(null));
    }

    // JDK constructors accept these windows but a consumer may fail later.
    // They must be rejected before any guest output by the bounded VM.
    public void badWindow() {
        System.out.println("guest-prefix");
        new MemoryImageSource(2, 2, new int[4], 1, 2);
        System.out.println("invalid-window-returned");
    }
    public void badStride() {
        System.out.println("guest-prefix");
        new MemoryImageSource(2, 2, new int[4], 0, 1);
        System.out.println("small-stride-returned");
    }
    public void nullPixels() {
        System.out.println("guest-prefix");
        new MemoryImageSource(1, 1, (int[]) null, 0, 1);
        System.out.println("null-pixels-returned");
    }
}
