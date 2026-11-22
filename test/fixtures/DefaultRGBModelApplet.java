// Owned Java 8 bytecode for a *detached headless* AWT color-model island.
// The JDK reference materializes int[] MemoryImageSource pixels; this is not
// the original Flash Applet and does not use a window or an SWF.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Component;
import java.awt.Graphics;
import java.awt.Image;
import java.awt.image.ColorModel;
import java.awt.image.IndexColorModel;
import java.awt.image.MemoryImageSource;

public class DefaultRGBModelApplet extends Applet {
    Image source;
    Image canvas;

    private Image image(int width, int height) {
        return ((Component) this).createImage(width, height);
    }
    private Image image(MemoryImageSource producer) {
        return ((Component) this).createImage(producer);
    }

    @Override public void init() {
        ColorModel first = ColorModel.getRGBdefault();
        ColorModel second = ColorModel.getRGBdefault();
        ColorModel detached = ((Component) this).getColorModel();
        System.out.println(first == second);
        System.out.println(first == detached);
        System.out.println(first instanceof IndexColorModel);
        System.out.println(first.getRGB(0x80402010));
        System.out.println(detached.getRGB(0xff123456));

        int[] pixels = new int[16];
        pixels[5] = 0xff112233;
        pixels[6] = 0x00112233; // Fully transparent: destination stays blue.
        pixels[9] = 0xff00ff80;
        pixels[10] = 0xff801020;
        MemoryImageSource producer = new MemoryImageSource(
            2, 2, first, pixels, 5, 4);
        pixels[5] = 0xff445566; // The producer retains the live source array.
        source = image(producer);
        pixels[5] = 0xff667788; // The already-created Image is a snapshot.
        canvas = image(4, 3);
        Graphics g = canvas.getGraphics();
        g.setColor(Color.blue);
        g.fillRect(0, 0, 4, 3);
        System.out.println(g.drawImage(source, 1, 0, this));
        System.out.println("rgb-ready");
    }

    // Select independently to prove generated Java refuses this model API
    // before emission (rather than failing later at createImage).
    public void modelOnly() {
        ColorModel first = ColorModel.getRGBdefault();
        System.out.println(first == ((Component) this).getColorModel());
        System.out.println(first.getRGB(0xff123456));
    }

    // JDK-valid, but this *different* model/int[] combination is not the
    // bounded default-RGB profile. Never treat every ColorModel as identity.
    public void indexedIntSource() {
        byte[] red = new byte[256], green = new byte[256], blue = new byte[256];
        red[1] = (byte) 255;
        IndexColorModel indexed = new IndexColorModel(8, 256, red, green, blue);
        int[] pixels = {1};
        Image image = image(new MemoryImageSource(1, 1, indexed, pixels, 0, 1));
        System.out.println(image != null);
    }
    public void indexedRead() {
        byte[] red = new byte[256], green = new byte[256], blue = new byte[256];
        red[1] = (byte) 255;
        ColorModel indexed = new IndexColorModel(8, 256, red, green, blue);
        System.out.println(indexed.getRGB(1));
    }

    // Valid JDK behavior, deliberately outside Magnolia's opaque-alpha-only
    // blit policy. Select this method separately for an explicit refusal.
    public void partialAlpha() {
        int[] pixels = {0x80112233};
        Image part = image(new MemoryImageSource(
            1, 1, ColorModel.getRGBdefault(), pixels, 0, 1));
        Graphics g = image(1, 1).getGraphics();
        g.setColor(Color.blue);
        g.fillRect(0, 0, 1, 1);
        g.drawImage(part, 0, 0, null);
        System.out.println("partial-drawn");
    }
}
