import java.awt.Image;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;
import java.awt.image.ImageConsumer;
import java.awt.image.ImageProducer;
import java.lang.reflect.Field;
import java.util.Hashtable;

/** JDK -Xverify:all independent 4x3 reference, no X11 or Applet lifecycle. */
public final class RasterProbeJdk {
  public static final class Host extends RasterProbeApplet {
    @Override public Image createImage(int w, int h) {
      return new BufferedImage(w, h, BufferedImage.TYPE_INT_ARGB);
    }
    @Override public Image createImage(ImageProducer producer) {
      final class Consumer implements ImageConsumer {
        int w, h; int[] pixels; boolean done;
        public void setDimensions(int width, int height) {
          w = width; h = height; pixels = new int[w * h];
        }
        public void setColorModel(ColorModel model) {}
        public void setProperties(Hashtable<?, ?> properties) {}
        public void setHints(int flags) {}
        public void setPixels(int x, int y, int width, int height,
            ColorModel model, byte[] data, int offset, int scan) {
          for (int row = 0; row < height; row++)
            for (int col = 0; col < width; col++)
              pixels[(y + row) * w + x + col] =
                  model.getRGB(data[offset + row * scan + col] & 255);
        }
        public void setPixels(int x, int y, int width, int height,
            ColorModel model, int[] data, int offset, int scan) {
          for (int row = 0; row < height; row++)
            for (int col = 0; col < width; col++)
              pixels[(y + row) * w + x + col] =
                  model.getRGB(data[offset + row * scan + col]);
        }
        public void imageComplete(int status) {
          if (status == STATICIMAGEDONE) done = true;
          // A later IMAGEERROR must not erase completed synchronous pixels.
        }
      }
      Consumer consumer = new Consumer();
      producer.startProduction(consumer);
      if (!consumer.done) throw new IllegalStateException("synchronous indexed source required");
      BufferedImage image = new BufferedImage(consumer.w, consumer.h,
                                              BufferedImage.TYPE_INT_ARGB);
      image.setRGB(0, 0, consumer.w, consumer.h, consumer.pixels, 0, consumer.w);
      return image;
    }
  }
  public static void main(String[] args) throws Exception {
    // Only the JDK reference bypasses the display-bound Applet constructor.
    Class<?> type = Class.forName("sun.misc.Unsafe");
    Field field = type.getDeclaredField("theUnsafe"); field.setAccessible(true);
    Object unsafe = field.get(null);
    Host applet = (Host) type.getMethod("allocateInstance", Class.class)
        .invoke(unsafe, Host.class);
    applet.init();
    BufferedImage original = (BufferedImage) applet.source;
    for (int y = 0; y < 2; y++) {
      for (int x = 0; x < 2; x++)
        System.out.print(String.format("%08x", original.getRGB(x, y)) +
                         (x == 1 ? "\n" : " "));
    }
    BufferedImage image = (BufferedImage) applet.canvas;
    for (int y = 0; y < 3; y++) {
      for (int x = 0; x < 4; x++)
        System.out.print(String.format("%08x", image.getRGB(x, y)) +
                         (x == 3 ? "\n" : " "));
    }
  }
}
