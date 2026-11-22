// Independent JDK 11 -Xverify:all reference for MemoryRGBDefaultApplet's
// unchanged compiled Java 8 bytecode. Unsafe bypasses Applet's display-only
// constructor; the host consumes original JDK MemoryImageSource synchronously.
// This is NOT an AWT peer, DS thread, or a browser playback oracle.
import java.awt.Image;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;
import java.awt.image.ImageConsumer;
import java.awt.image.ImageProducer;
import java.lang.reflect.Field;
import java.util.Hashtable;

public final class MemoryRGBDefaultJdk {
    public static final class Host extends MemoryRGBDefaultApplet {
        int produced;
        boolean onlyDefaultModel = true;
        @Override public Image createImage(int width, int height) {
            return new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
        }
        @Override public Image createImage(ImageProducer producer) {
            final class Consumer implements ImageConsumer {
                int width, height;
                int[] argb;
                boolean complete, correctModel = true;
                public void setDimensions(int w, int h) {
                    width = w; height = h;
                    if (w != 2 || (h != 2 && h != 3))
                        throw new AssertionError("unexpected dimensions");
                    argb = new int[w * h];
                }
                public void setColorModel(ColorModel model) {
                    correctModel &= model == ColorModel.getRGBdefault();
                }
                public void setProperties(Hashtable<?, ?> ignored) {}
                public void setHints(int ignored) {}
                public void setPixels(int x, int y, int w, int h,
                        ColorModel model, byte[] data, int offset, int scan) {
                    throw new AssertionError("expected JDK int[] producer");
                }
                public void setPixels(int x, int y, int w, int h,
                        ColorModel model, int[] data, int offset, int scan) {
                    correctModel &= model == ColorModel.getRGBdefault();
                    for (int row = 0; row < h; ++row) {
                        for (int col = 0; col < w; ++col) {
                            argb[(y + row) * width + x + col] =
                                model.getRGB(data[offset + row * scan + col]);
                        }
                    }
                }
                public void imageComplete(int status) {
                    if (status == STATICIMAGEDONE) complete = true;
                }
            }
            Consumer output = new Consumer();
            producer.startProduction(output);
            if (!output.complete || !output.correctModel || output.argb == null)
                throw new AssertionError("JDK producer not static default ARGB");
            onlyDefaultModel &= output.correctModel;
            ++produced;
            BufferedImage image = new BufferedImage(output.width, output.height,
                                                    BufferedImage.TYPE_INT_ARGB);
            image.setRGB(0, 0, output.width, output.height,
                         output.argb, 0, output.width);
            return image;
        }
    }
    private static void show(BufferedImage image) {
        for (int y = 0; y < image.getHeight(); y++) {
            for (int x = 0; x < image.getWidth(); x++) {
                System.out.print(String.format("%08x", image.getRGB(x, y)) +
                                 (x + 1 == image.getWidth() ? "\n" : " "));
            }
        }
    }
    public static void main(String[] args) throws Exception {
        Class<?> unsafe = Class.forName("sun.misc.Unsafe");
        Field field = unsafe.getDeclaredField("theUnsafe");
        field.setAccessible(true);
        Host applet = (Host) unsafe.getMethod("allocateInstance", Class.class)
            .invoke(field.get(null), Host.class);
        if (!Boolean.getBoolean("java.awt.headless"))
            throw new AssertionError("headless reference only");
        if (args.length == 1) {
            if ("badWindow".equals(args[0])) applet.badWindow();
            else if ("badStride".equals(args[0])) applet.badStride();
            else if ("nullPixels".equals(args[0])) applet.nullPixels();
            else if ("boundsOnly".equals(args[0])) {
                applet.boundsOnly();
                show((BufferedImage) applet.source); // Also checks post-flush pixels.
            }
            else throw new IllegalArgumentException(args[0]);
            return;
        }
        if (args.length != 0) throw new IllegalArgumentException("argument count");
        applet.onlyDefaultModel = true; // allocateInstance skips field initializers.
        applet.init();
        System.out.println("default-model=" + applet.onlyDefaultModel +
                           "|produced=" + applet.produced);
        show((BufferedImage) applet.source);
        show((BufferedImage) applet.explicitSource);
        show((BufferedImage) applet.canvas);
    }
}
