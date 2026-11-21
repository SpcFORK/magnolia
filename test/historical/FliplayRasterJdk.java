import java.awt.Image;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;
import java.awt.image.ImageConsumer;
import java.awt.image.ImageProducer;
import java.io.DataInputStream;
import java.io.FileInputStream;
import java.lang.reflect.Field;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.MessageDigest;
import java.util.Hashtable;

/**
 * JDK 11 -Xverify:all reference for original Fliplay image/patch methods.
 * In a displayless test process, an Applet subclass replaces ONLY the host
 * createImage calls with synchronous BufferedImages. Its ImageConsumer uses
 * JDK ColorModel.getRGB, not an FLC decoder. Unsafe.allocateInstance bypasses
 * Applet's GUI constructor under -Djava.awt.headless=true; this does NOT
 * establish actual Applet lifecycle or UI behavior. The original Fliplay
 * header, frame loader, decoder, process_first_frame/process_loop_frame and
 * flickframe ctor/add_patch/delete_patches/draw_patches ARE invoked.
 */
public final class FliplayRasterJdk {
  public static final class HeadlessApplet extends fliplay {
    @Override public Image createImage(int width, int height) {
      return new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
    }
    @Override public Image createImage(ImageProducer source) {
      Consumer consumer = new Consumer();
      source.startProduction(consumer);
      if (!consumer.complete || consumer.pixels == null ||
          consumer.width < 1 || consumer.height < 1) {
        throw new IllegalStateException("synchronous ImageProducer required");
      }
      BufferedImage image = new BufferedImage(consumer.width, consumer.height,
                                              BufferedImage.TYPE_INT_ARGB);
      image.setRGB(0, 0, consumer.width, consumer.height, consumer.pixels,
                   0, consumer.width);
      return image;
    }
  }
  private static final class Consumer implements ImageConsumer {
    int width, height;
    int[] pixels;
    boolean complete;
    @Override public void setDimensions(int w, int h) {
      width = w; height = h; pixels = new int[w * h];
    }
    @Override public void setColorModel(ColorModel model) {}
    @Override public void setProperties(Hashtable<?, ?> properties) {}
    @Override public void setHints(int hints) {}
    @Override public void setPixels(int x, int y, int w, int h,
        ColorModel model, byte[] source, int offset, int scan) {
      for (int row = 0; row < h; row++) for (int col = 0; col < w; col++) {
        pixels[(y + row) * width + x + col] =
            model.getRGB(source[offset + row * scan + col] & 255);
      }
    }
    @Override public void setPixels(int x, int y, int w, int h,
        ColorModel model, int[] source, int offset, int scan) {
      for (int row = 0; row < h; row++) for (int col = 0; col < w; col++) {
        pixels[(y + row) * width + x + col] =
            model.getRGB(source[offset + row * scan + col]);
      }
    }
    @Override public void imageComplete(int status) {
      if (status == ImageConsumer.STATICIMAGEDONE) complete = true;
      // Some JDK 11 MemoryImageSource paths send IMAGEERROR after completion.
      // Do not discard an already delivered, complete synchronous frame.
    }
  }
  private static HeadlessApplet newAppletWithoutDisplay() throws Exception {
    Class<?> unsafeClass = Class.forName("sun.misc.Unsafe");
    Field field = unsafeClass.getDeclaredField("theUnsafe");
    field.setAccessible(true);
    Object unsafe = field.get(null);
    HeadlessApplet applet = (HeadlessApplet) unsafeClass
        .getMethod("allocateInstance", Class.class)
        .invoke(unsafe, HeadlessApplet.class);
    applet.frame_data_len = -1;
    return applet;
  }
  private static String hex(byte[] bytes) {
    StringBuilder text = new StringBuilder();
    for (byte b : bytes) text.append(String.format("%02x", b & 255));
    return text.toString();
  }
  private static byte[] rgb(BufferedImage image) {
    byte[] bytes = new byte[image.getWidth() * image.getHeight() * 3];
    int at = 0;
    for (int y = 0; y < image.getHeight(); y++)
      for (int x = 0; x < image.getWidth(); x++) {
        int argb = image.getRGB(x, y);
        bytes[at++] = (byte)(argb >>> 16);
        bytes[at++] = (byte)(argb >>> 8);
        bytes[at++] = (byte)argb;
      }
    return bytes;
  }
  public static void main(String[] args) throws Exception {
    if (args.length != 3) throw new IllegalArgumentException("FLC, output directory, limit required");
    int limit = Integer.parseInt(args[2]);
    if (limit < 1 || limit > 61) throw new IllegalArgumentException("limit 1..61");
    HeadlessApplet player = newAppletWithoutDisplay();
    player.red = new byte[256]; player.green = new byte[256];
    player.blue = new byte[256];
    try (DataInputStream input = new DataInputStream(new FileInputStream(args[0]))) {
      player.load_header(input);
      if (player.fli_frames != 60 || player.image_width != 320 ||
          player.image_height != 240) throw new AssertionError("unexpected original header");
      player.pixels = new byte[320 * 240];
      player.row_start = new int[240]; player.row_end = new int[240];
      player.block_height = 64; player.patch_blocks = 4;
      player.frame = new flickframe[61];
      player.bimage = player.createImage(320, 240);
      player.bgraphics = player.bimage.getGraphics();
      MessageDigest aggregate = MessageDigest.getInstance("SHA-256");
      for (int i = 0; i < limit; ++i) {
        player.frame[i] = new flickframe(4);
        player.load_frame(input, i);
        player.extract_frame(i);
        if (i == 0) player.process_first_frame();
        else {
          player.process_loop_frame(i);
          player.frame[i].draw_patches(player.bgraphics);
        }
        byte[] pixels = rgb((BufferedImage)player.bimage);
        String frame = String.format("%03d", i);
        Files.write(Paths.get(args[1], frame + ".rgb"), pixels);
        if (i < 60) aggregate.update(pixels);
        StringBuilder boxes = new StringBuilder();
        for (int k = 0; k < player.frame[i].count; ++k) {
          if (k != 0) boxes.append(';');
          Image patch = player.frame[i].patch[k];
          boxes.append(player.frame[i].xpos[k]).append(',')
              .append(player.frame[i].ypos[k]).append(',')
              .append(patch.getWidth(null)).append(',').append(patch.getHeight(null));
        }
        System.out.println("JDK RENDER " + frame + " patches=" + player.frame[i].count +
            " patchPixels=" + player.pixel_sum + " boxes=" +
            (boxes.length() == 0 ? "-" : boxes.toString()) + " rgbSHA256=" +
            hex(MessageDigest.getInstance("SHA-256").digest(pixels)));
      }
      if (limit == 61) {
        if (input.read() != -1) throw new AssertionError("trailing movie data");
        System.out.println("JDK ALL_SHA256 " + hex(aggregate.digest()));
      }
    }
  }
}
