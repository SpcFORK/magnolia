// Independent JDK 11 -Xverify:all oracle for detached headless color models.
// Bypass only Applet's display-dependent constructor. No Java plugin/window.
import java.awt.Image;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;
import java.awt.image.ImageConsumer;
import java.awt.image.ImageProducer;
import java.lang.reflect.Field;
import java.util.Hashtable;

public final class DefaultRGBModelJdk {
    public static final class Host extends DefaultRGBModelApplet {
        @Override public Image createImage(int width, int height) {
            return new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
        }
        @Override public Image createImage(ImageProducer producer) {
            final class Consumer implements ImageConsumer {
                int width, height;
                int[] argb;
                boolean complete;
                public void setDimensions(int w, int h) {
                    width = w;
                    height = h;
                    argb = new int[w * h];
                }
                public void setColorModel(ColorModel ignored) {}
                public void setProperties(Hashtable<?, ?> ignored) {}
                public void setHints(int ignored) {}
                public void setPixels(int x, int y, int w, int h,
                        ColorModel model, byte[] values, int offset, int scan) {
                    throw new AssertionError("expected int[] RGB pixels");
                }
                public void setPixels(int x, int y, int w, int h,
                        ColorModel model, int[] values, int offset, int scan) {
                    for (int row = 0; row < h; row++) {
                        for (int col = 0; col < w; col++) {
                            argb[(y + row) * width + x + col] =
                                model.getRGB(values[offset + row * scan + col]);
                        }
                    }
                }
                public void imageComplete(int status) {
                    if (status == STATICIMAGEDONE) complete = true;
                }
            }
            Consumer output = new Consumer();
            producer.startProduction(output);
            if (!output.complete || output.width > 16 || output.height > 16)
                throw new AssertionError("bounded synchronous RGB source required");
            BufferedImage image = new BufferedImage(output.width, output.height,
                                                    BufferedImage.TYPE_INT_ARGB);
            image.setRGB(0, 0, output.width, output.height,
                         output.argb, 0, output.width);
            return image;
        }
    }

    public static void main(String[] args) throws Exception {
        Class<?> type = Class.forName("sun.misc.Unsafe");
        Field field = type.getDeclaredField("theUnsafe");
        field.setAccessible(true);
        Host applet = (Host) type.getMethod("allocateInstance", Class.class)
            .invoke(field.get(null), Host.class);
        if (!Boolean.getBoolean("java.awt.headless"))
            throw new AssertionError("headless reference only");
        if (args.length == 1) {
            if ("model-only".equals(args[0])) applet.modelOnly();
            else if ("indexed".equals(args[0])) applet.indexedIntSource();
            else if ("indexed-read".equals(args[0])) applet.indexedRead();
            else if ("partial".equals(args[0])) applet.partialAlpha();
            else throw new IllegalArgumentException(args[0]);
            return;
        }
        applet.init();
        show((BufferedImage) applet.source);
        show((BufferedImage) applet.canvas);
    }

    private static void show(BufferedImage image) {
        for (int y = 0; y < image.getHeight(); y++) {
            for (int x = 0; x < image.getWidth(); x++) {
                System.out.print(String.format("%08x", image.getRGB(x, y)) +
                                 (x + 1 == image.getWidth() ? "\n" : " "));
            }
        }
    }
}
