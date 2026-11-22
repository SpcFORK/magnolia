import java.io.IOException;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Byte-identical fixture against real MicroEmulator, with UNCHANGED original
// a and s0 PNGs in its local JAR. Only top-left binary-alpha offscreen blits.
public final class BlitSliceMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }
    protected void startApp() throws MIDletStateChangeException {
        try {
            System.out.println("TRACE:blit-begin");
            Image a = Image.createImage("/a");
            Image sprite = Image.createImage("/s0");
            Image target = Image.createImage(20, 12);
            Graphics graphics = target.getGraphics();
            graphics.setColor(0xff123456);
            graphics.fillRect(0, 0, 20, 12);
            graphics.drawImage(a, -5, 2, Graphics.LEFT | Graphics.TOP);
            graphics.drawImage(sprite, 12, -8, Graphics.LEFT | Graphics.TOP);
            int[] pixels = new int[20 * 12];
            target.getRGB(pixels, 0, 20, 0, 0, 20, 12);
            int checksum = 0x811c9dc5;
            int background = 0;
            for (int i = 0; i < pixels.length; i++) {
                checksum = (checksum ^ pixels[i]) * 16777619;
                if (pixels[i] == 0xff123456) background++;
            }
            System.out.println(checksum);
            System.out.println(background);
            System.out.println(pixels[0]);
            System.out.println(pixels[2 * 20 + 1]);
            System.out.println(pixels[11 * 20 + 19]);
            System.out.println("TRACE:blit-end");
            Display.getDisplay(this).setCurrent(new View());
        } catch (IOException unexpected) {
            System.out.println("BAD:blit-image-resource");
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
