import java.awt.Image;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;
import java.awt.image.ImageConsumer;
import java.awt.image.ImageProducer;
import java.lang.reflect.Field;
import java.net.URL;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.security.MessageDigest;
import java.util.Hashtable;

/**
 * JDK 11 -Xverify:all oracle for ORIGINAL fliplay.<init>/init/load_fli on an
 * actual local FLC or ZIP stream. GUI Applet construction must be bypassed in
 * headless mode; only parameters/codebase, synchronous image construction,
 * status and repaint scheduling are adapted. No original method is rewritten.
 * Repaints are counted, not treated as a completed animation/event thread.
 */
public final class FliplayLoadJdk {
  private static String mode;
  private static URL base;
  private static Path output;
  public static final class HeadlessApplet extends fliplay {
    int repaints;
    int statuses;
    @Override public String getParameter(String name) {
      if (name.equals("File")) return mode.equals("zip") ? "moonclouds.zip" : "moonclouds.flc";
      if (name.equals("ZipFile") && mode.equals("zip")) return "yes";
      return null;
    }
    @Override public URL getCodeBase() { return base; }
    @Override public void showStatus(String status) { statuses++; }
    @Override public void repaint() { repaints++; }
    @Override public Image createImage(int w, int h) {
      return new BufferedImage(w, h, BufferedImage.TYPE_INT_ARGB);
    }
    @Override public Image createImage(ImageProducer source) {
      Consumer c = new Consumer();
      source.startProduction(c);
      if (!c.complete || c.pixels == null || c.width < 1 || c.height < 1)
        throw new IllegalStateException("synchronous ImageProducer required");
      BufferedImage image = new BufferedImage(c.width, c.height,
                                              BufferedImage.TYPE_INT_ARGB);
      image.setRGB(0, 0, c.width, c.height, c.pixels, 0, c.width);
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
    @Override public void setProperties(Hashtable<?, ?> props) {}
    @Override public void setHints(int hints) {}
    @Override public void setPixels(int x, int y, int w, int h, ColorModel model,
        byte[] source, int offset, int scan) {
      for (int row = 0; row < h; row++) for (int col = 0; col < w; col++)
        pixels[(y + row) * width + x + col] =
            model.getRGB(source[offset + row * scan + col] & 255);
    }
    @Override public void setPixels(int x, int y, int w, int h, ColorModel model,
        int[] source, int offset, int scan) {
      for (int row = 0; row < h; row++) for (int col = 0; col < w; col++)
        pixels[(y + row) * width + x + col] =
            model.getRGB(source[offset + row * scan + col]);
    }
    @Override public void imageComplete(int status) {
      if (status == ImageConsumer.STATICIMAGEDONE) complete = true;
    }
  }
  private static HeadlessApplet newWithoutDisplay() throws Exception {
    Class<?> clazz = Class.forName("sun.misc.Unsafe");
    Field field = clazz.getDeclaredField("theUnsafe");
    field.setAccessible(true);
    Object unsafe = field.get(null);
    HeadlessApplet player = (HeadlessApplet) clazz
        .getMethod("allocateInstance", Class.class)
        .invoke(unsafe, HeadlessApplet.class);
    // Exact defaults assigned by the original fliplay.<init>()V.
    player.frame_data_len = -1; player.klick_dim = 24;
    player.loop_flag = true; player.klick_panel_mode = true;
    player.keyboard_control = true;
    return player;
  }
  private static byte[] argb(Image image) {
    if (!(image instanceof BufferedImage)) throw new AssertionError("not buffered");
    BufferedImage raster = (BufferedImage) image;
    int w = raster.getWidth(), h = raster.getHeight();
    if (w < 1 || h < 1 || w * h > 81920) throw new AssertionError("raster bounds");
    byte[] raw = new byte[w * h * 4];
    int at = 0;
    for (int y = 0; y < h; y++) for (int x = 0; x < w; x++) {
      int p = raster.getRGB(x, y);
      raw[at++] = (byte)(p >>> 24); raw[at++] = (byte)(p >>> 16);
      raw[at++] = (byte)(p >>> 8); raw[at++] = (byte)p;
    }
    return raw;
  }
  private static String hex(byte[] data) {
    StringBuilder s = new StringBuilder();
    for (byte b : data) s.append(String.format("%02x", b & 255));
    return s.toString();
  }
  public static void main(String[] args) throws Exception {
    if (args.length != 3 || (!args[0].equals("raw") && !args[0].equals("zip")))
      throw new IllegalArgumentException("raw|zip INPUT DIRECTORY required");
    mode = args[0];
    Path supplied = Paths.get(args[1]).toAbsolutePath();
    base = supplied.getParent().toUri().toURL();
    output = Paths.get(args[2]);
    if (!Files.isDirectory(output)) throw new AssertionError("missing output directory");
    HeadlessApplet player = newWithoutDisplay();
    player.init();
    player.load_fli(); // all original control flow, URL/ZIP and frame methods
    if (!player.loaded || !player.extracted || player.load_failed ||
        player.fli_frames != 60 || player.frame_number != 60 ||
        player.frame == null || player.frame.length != 61 ||
        player.repaints != 61 || player.statuses != 61)
      throw new AssertionError("original loader did not finish");
    Files.write(output.resolve("icon.argb"), argb(player.klick_image));
    Files.write(output.resolve("frame0.argb"), argb(player.bimage));
    int count = 0;
    long cells = 0;
    for (int i = 0; i < 61; ++i) {
      flickframe f = player.frame[i];
      if (f == null || f.count < 0 || f.count > 4)
        throw new AssertionError("bad frame/patch count " + i);
      for (int p = 0; p < f.count; p++) {
        Image patch = f.patch[p];
        int w = patch.getWidth(null), h = patch.getHeight(null);
        String name = String.format("patch-%03d-%d-%d-%d-%d-%d.argb",
            i, p, f.xpos[p], f.ypos[p], w, h);
        byte[] raw = argb(patch);
        Files.write(output.resolve(name), raw);
        cells += w * h;
        count++;
      }
    }
    System.out.println("JDK LIVE LOAD " + mode + " frames=" + player.fli_frames +
        " ring=" + player.frame_number + " patches=" + count +
        " cells=" + cells + " patchPixels=" + player.pixel_sum +
        " repaints=" + player.repaints + " statuses=" + player.statuses +
        " frame0SHA256=" +
        hex(MessageDigest.getInstance("SHA-256").digest(argb(player.bimage))));
  }
}
