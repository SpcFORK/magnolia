import java.io.IOException;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Compile ONCE against real MicroEmulator, then execute identical bytes in
// MicroEmulator and Magnolia. The JAR embeds three UNCHANGED original Acotel
// PNG members selected in Magnolia by --resources; this is an owned fixture,
// NOT the original game's classfiles, painting or delivered input.
public final class ImageResourceMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }

    protected void startApp() throws MIDletStateChangeException {
        try {
            System.out.println("TRACE:image-begin");
            Image a = Image.createImage("/a");
            Image b = Image.createImage("/b1");
            Image sprite = Image.createImage("/s0");
            System.out.println(a.getWidth());
            System.out.println(a.getHeight());
            System.out.println(b.getWidth());
            System.out.println(b.getHeight());
            System.out.println(sprite.getWidth());
            System.out.println(sprite.getHeight());
            int[] sample = new int[11];
            a.getRGB(sample, 1, 5, 0, 0, 3, 2);
            System.out.println(sample[1]);
            System.out.println(sample[2]);
            System.out.println(sample[6]);
            sprite.getRGB(sample, 2, 4, 4, 4, 2, 2);
            System.out.println(sample[2]);
            System.out.println(sample[3]);
            System.out.println(sample[6]);
            try {
                a.getGraphics();
                System.out.println("BAD:immutable-graphics");
            } catch (IllegalStateException expected) {
                System.out.println("TRACE:immutable-graphics-caught");
            }
            Image mutable = Image.createImage(4, 3);
            int[] pixels = new int[12];
            mutable.getRGB(pixels, 0, 4, 0, 0, 4, 3);
            System.out.println(pixels[0]);
            Graphics graphics = mutable.getGraphics();
            System.out.println(graphics.getColor());
            graphics.setColor(0xff123456);
            System.out.println(graphics.getColor());
            graphics.fillRect(-1, -1, 3, 3);
            mutable.getRGB(pixels, 0, 4, 0, 0, 4, 3);
            System.out.println(pixels[0]);
            System.out.println(pixels[1]);
            System.out.println(pixels[5]);
            System.out.println(pixels[10]);
            try {
                Image.createImage(0, 1);
                System.out.println("BAD:zero-width");
            } catch (IllegalArgumentException expected) {
                System.out.println("TRACE:zero-width-caught");
            }
            try {
                Image.createImage("/absent");
                System.out.println("BAD:missing-resource");
            } catch (IOException expected) {
                System.out.println("TRACE:missing-resource-caught");
            }
            System.out.println("TRACE:image-end");
            Display.getDisplay(this).setCurrent(new View());
        } catch (IOException unexpected) {
            System.out.println("BAD:image-io");
        }
    }

    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
