// Owned Java 8 bytecode, independently exercised on a JDK BufferedImage
// host. Each selected method isolates one original Flash AWT member.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Component;
import java.awt.Graphics;
import java.awt.Image;

public class FiveRasterApplet extends Applet {
    public Image canvas;

    private Image image(int width, int height) {
        return ((Component) this).createImage(width, height);
    }

    public void opaqueColor() {
        canvas = image(4, 1);
        Graphics graphics = canvas.getGraphics();
        int[] packed = {0x11223344, 0x80aabbcc, -1, 0};
        for (int x = 0; x < packed.length; x++) {
            graphics.setColor(new Color(packed[x]));
            graphics.fillRect(x, 0, 1, 1);
        }
        System.out.println("opaque-color-ready");
    }

    public void intersectClip() {
        canvas = image(5, 4);
        Graphics graphics = canvas.getGraphics();
        graphics.setColor(Color.red);
        graphics.fillRect(0, 0, 5, 4);
        graphics.clipRect(1, 1, 4, 3);
        graphics.setColor(Color.blue);
        graphics.fillRect(0, 0, 5, 4);
        graphics.clipRect(2, 0, 2, 3);
        graphics.setColor(Color.green);
        graphics.fillRect(-3, -3, 11, 11);
        graphics.setColor(Color.magenta);
        graphics.drawRect(1, 1, 3, 2);
        Image yellow = image(2, 2);
        Graphics source = yellow.getGraphics();
        source.setColor(Color.yellow);
        source.fillRect(0, 0, 2, 2);
        System.out.println(graphics.drawImage(yellow, 2, 0, null));
        // No later positive clip may reopen an empty intersection.
        graphics.clipRect(50, 50, 1, 1);
        graphics.setColor(Color.white);
        graphics.fillRect(0, 0, 5, 4);
        graphics.clipRect(0, 0, 5, 4);
        graphics.drawRect(0, 0, 4, 3);
        System.out.println("clip-ready");
    }

    // The existing, separately JDK-checked centered-arrow raster also has
    // to respect the newly intersected clip rather than draw outside it.
    public void intersectArrowClip() {
        canvas = image(7, 6);
        Graphics graphics = canvas.getGraphics();
        graphics.setColor(Color.blue);
        graphics.fillRect(0, 0, 7, 6);
        graphics.clipRect(2, 1, 3, 3);
        graphics.setColor(Color.yellow);
        graphics.fillPolygon(new int[]{1, 5, 1}, new int[]{0, 2, 4}, 3);
        System.out.println("clipped-arrow-ready");
    }

    // Isolate the two newly admitted symbols from Component.createImage for
    // generated-Java's exact pre-emission refusal checks.
    public void colorOnly() {
        Color made = new Color(0x00123456);
        System.out.println(made != null);
    }
    public void clipOnly(Graphics graphics) {
        graphics.clipRect(0, 0, 1, 1);
    }

    // Distinct JDK Color(int,boolean) alpha constructor: never silently
    // treat a semitransparent value as the admitted opaque Color(int).
    public void colorWithAlpha() {
        canvas = image(1, 1);
        Graphics graphics = canvas.getGraphics();
        graphics.setColor(new Color(0x80112233, true));
        graphics.fillRect(0, 0, 1, 1);
        System.out.println("alpha-constructor-accepted");
    }

    // JDK-valid clip replacement is NOT sequential clip intersection.
    public void replaceClip() {
        canvas = image(2, 2);
        Graphics graphics = canvas.getGraphics();
        graphics.setClip(0, 0, 1, 1);
        graphics.setColor(Color.blue);
        graphics.fillRect(0, 0, 2, 2);
        System.out.println("setClip-accepted");
    }

    // JDK-valid but outside the bounded positive-size, raster-only clip.
    public void negativeClipSize() {
        canvas = image(2, 2);
        Graphics graphics = canvas.getGraphics();
        graphics.clipRect(1, 1, -2, 1);
        System.out.println("JDK-negative-size-accepted");
    }

    // A detached Component returns null, not an invented writable Graphics.
    public void noPeerClip() {
        Graphics graphics = ((Component) this).getGraphics();
        if (graphics == null) System.out.println("no-peer");
    }
}
