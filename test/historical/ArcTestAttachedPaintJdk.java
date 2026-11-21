// Independent reference for the original ArcTest.init() -> attached child
// ArcCanvas.paint path. This is NOT part of the detached draw-call oracle.
import java.awt.Graphics2D;
import java.awt.image.BufferedImage;
public final class ArcTestAttachedPaintJdk {
    public static void main(String[] args) {
        if (args.length != 0) throw new IllegalArgumentException("no arguments expected");
        Graphics2D graphics = new BufferedImage(1, 1,
            BufferedImage.TYPE_INT_ARGB).createGraphics();
        try { new ArcTestAttachedPaintApplet().paint(graphics); }
        finally { graphics.dispose(); }
        System.out.println("JDK original ArcTest attached child paint returned");
    }
}
