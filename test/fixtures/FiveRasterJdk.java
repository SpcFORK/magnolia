// Independent JDK 11 -Xverify:all, headless BufferedImage oracle.
// Only this harness bypasses Applet's display-bound constructor; the exact
// fixture's method bytecode still calls its virtual Component.createImage.
import java.awt.Image;
import java.awt.image.BufferedImage;
import java.lang.reflect.Field;

public final class FiveRasterJdk {
    public static final class Host extends FiveRasterApplet {
        @Override public Image createImage(int width, int height) {
            return new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
        }
    }
    public static void main(String[] args) throws Exception {
        if (!Boolean.getBoolean("java.awt.headless"))
            throw new AssertionError("headless reference required");
        Class<?> type = Class.forName("sun.misc.Unsafe");
        Field field = type.getDeclaredField("theUnsafe");
        field.setAccessible(true);
        Host applet = (Host) type.getMethod("allocateInstance", Class.class)
            .invoke(field.get(null), Host.class);
        if (args.length != 1) throw new IllegalArgumentException("mode");
        if ("color".equals(args[0])) applet.opaqueColor();
        else if ("clip".equals(args[0])) applet.intersectClip();
        else if ("arrow-clip".equals(args[0])) applet.intersectArrowClip();
        else if ("color-only".equals(args[0])) applet.colorOnly();
        else if ("clip-only".equals(args[0]))
            applet.clipOnly(new BufferedImage(1, 1,
                BufferedImage.TYPE_INT_ARGB).getGraphics());
        else if ("alpha".equals(args[0])) applet.colorWithAlpha();
        else if ("replace".equals(args[0])) applet.replaceClip();
        else if ("negative".equals(args[0])) applet.negativeClipSize();
        else if ("peerless".equals(args[0])) applet.noPeerClip();
        else throw new IllegalArgumentException(args[0]);
        if (applet.canvas != null) {
            BufferedImage image = (BufferedImage) applet.canvas;
            for (int y = 0; y < image.getHeight(); y++) {
                for (int x = 0; x < image.getWidth(); x++)
                    System.out.print(String.format("%08x", image.getRGB(x, y)) +
                        (x + 1 == image.getWidth() ? "\n" : " "));
            }
        }
    }
}
