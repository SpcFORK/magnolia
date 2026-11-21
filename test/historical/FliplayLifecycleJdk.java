import java.awt.Event;
import java.awt.Graphics;
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
 * JDK 11 reference for finite ORIGINAL fliplay.init/start/Runnable.run,
 * load_fli/update/synchronized next_frame/paint. Headless AWT setup differs
 * from a real browser: only synchronous createImage and repaint->update are
 * adapted, with an actual JDK Thread and the actual local FLC/ZIP stream.
 */
public final class FliplayLifecycleJdk {
  private static String mode;
  private static URL base;
  private static boolean full;
  private static boolean capture;
  private static Path frames;
  private static MessageDigest aggregate;
  private static void snapshot(BufferedImage image, int number) {
    byte[] rgba = argb(image);
    byte[] rgb = new byte[230400];
    for (int i = 0, at = 0; i < rgba.length; i += 4) {
      rgb[at++] = rgba[i + 1]; rgb[at++] = rgba[i + 2]; rgb[at++] = rgba[i + 3];
    }
    if (number < 60) aggregate.update(rgb);
    try { Files.write(frames.resolve(String.format("%03d.rgb", number)), rgb); }
    catch (java.io.IOException e) { throw new IllegalStateException("frame export failed", e); }
  }
  public static final class HeadlessApplet extends fliplay {
    int repaints, statuses, captured, lastCaptured;
    BufferedImage screen;
    @Override public String getParameter(String name) {
      if (name.equals("File")) return mode.equals("zip") ? "moonclouds.zip" : "moonclouds.flc";
      if (name.equals("ZipFile") && mode.equals("zip")) return "yes";
      if (name.equals("Loops")) return "1";
      if (name.equals("End")) return full ? "60" : "5";
      if (name.equals("Speed")) return "1";
      return null;
    }
    @Override public URL getCodeBase() { return base; }
    @Override public void showStatus(String text) { statuses++; }
    @Override public void repaint() {
      repaints++;
      // Loading repaint requests coalesce; the original first frame has
      // already been drawn to bimage when frame_number is zero.
      if (capture && captured == 0 && display_frame && frame_number == 0) {
        snapshot((BufferedImage)bimage, 0);
        lastCaptured = 0;
        captured++;
      }
      // The original run() scheduling request synchronously invokes original
      // update(Graphics), which invokes original synchronized next_frame().
      if (loaded && engine != null) {
        if (screen == null) screen = new BufferedImage(320, 240, BufferedImage.TYPE_INT_ARGB);
        Graphics g = screen.getGraphics();
        try { update(g); } finally { g.dispose(); }
        if (capture && frame_number != lastCaptured && captured < 61) {
          snapshot(screen, frame_number);
          lastCaptured = frame_number;
          captured++;
        }
      }
    }
    @Override public Image createImage(int w, int h) {
      return new BufferedImage(w, h, BufferedImage.TYPE_INT_ARGB);
    }
    @Override public Image createImage(ImageProducer source) {
      Consumer c = new Consumer();
      source.startProduction(c);
      if (!c.complete || c.pixels == null || c.width < 1 || c.height < 1)
        throw new IllegalStateException("synchronous ImageProducer required");
      BufferedImage image = new BufferedImage(c.width, c.height, BufferedImage.TYPE_INT_ARGB);
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
    // Exact original fliplay.<init>()V instance defaults.
    player.frame_data_len = -1; player.klick_dim = 24;
    player.loop_flag = true; player.klick_panel_mode = true;
    player.keyboard_control = true;
    return player;
  }
  private static byte[] argb(BufferedImage image) {
    if (image == null || image.getWidth() != 320 || image.getHeight() != 240)
      throw new AssertionError("missing final 320x240 screen");
    byte[] raw = new byte[320 * 240 * 4];
    int at = 0;
    for (int y = 0; y < 240; y++) for (int x = 0; x < 320; x++) {
      int p = image.getRGB(x, y);
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
    if ((args.length != 3 && args.length != 4 && args.length != 5) ||
        (!args[0].equals("raw") && !args[0].equals("zip")) ||
        (args.length == 5 && !args[3].equals("--full")) ||
        (args.length == 4 && !args[3].equals("--events") &&
            !args[3].equals("--events-full")))
      throw new IllegalArgumentException("raw|zip INPUT SCREEN.argb [--full FRAME_DIR | --events | --events-full] required");
    mode = args[0];
    boolean events = args.length == 4;
    capture = args.length == 5;
    full = capture || (events && args[3].equals("--events-full"));
    if (capture) {
      frames = Paths.get(args[4]);
      if (!Files.isDirectory(frames)) throw new AssertionError("missing frame output directory");
      aggregate = MessageDigest.getInstance("SHA-256");
    }
    Path input = Paths.get(args[1]).toAbsolutePath();
    base = input.getParent().toUri().toURL();
    HeadlessApplet player = newWithoutDisplay();
    player.init();
    player.start(); // JDK executes original run() on a real Thread
    if (player.engine == null) throw new AssertionError("no original worker thread");
    player.engine.join(30_000L);
    if (player.engine.isAlive()) throw new AssertionError("original worker did not terminate");
    if (!player.loaded || player.load_failed ||
        player.frame_number != (full ? 60 : 5) ||
        player.work_loops != 0 || player.loop_flag || !player.userpause ||
        player.repaints != (full ? 122 : 67) || player.statuses != 63 ||
        (capture && player.captured != 61))
      throw new AssertionError("original finite lifecycle disagrees: frame="+
          player.frame_number + " loops=" + player.work_loops +
          " repaint=" + player.repaints + " status=" + player.statuses +
          " captured=" + player.captured);
    if (events) {
      // A second, explicitly bounded acceptance slice: dispatch real legacy
      // java.awt.Event objects to original fliplay.handleEvent/mouseDown/
      // mouseDrag bytecode *after* the finite original run has paused.
      int[][] records = {
          {2032,0,0,0}, {401,43,0,0}, {401,45,0,0},
          {401,114,0,0}, {501,0,36,12}, {506,0,55,20},
          {501,0,31,20}, {401,46,0,0}, {401,48,0,0}
      };
      StringBuilder handled = new StringBuilder();
      for (int i = 0; i < records.length; i++) {
        int[] record = records[i];
        Event event = new Event(player, record[0], null);
        event.key = record[1]; event.x = record[2]; event.y = record[3];
        handled.append(player.handleEvent(event) ? '1' : '0');
        if (full) System.out.println("JDK GUI EVENT " + i +
            " handled=" + handled.charAt(i) +
            " frame=" + player.frame_number +
            " repaints=" + player.repaints +
            " statuses=" + player.statuses +
            " screenSHA256=" + hex(MessageDigest.getInstance("SHA-256")
                .digest(argb(player.screen))));
      }
      byte[] eventScreen = argb(player.screen);
      System.out.println("JDK EVENTS " + mode + " handled=" + handled +
          " frame=" + player.frame_number + " repaints=" + player.repaints +
          " statuses=" + player.statuses + " reverse=" + player.reverse_flag +
          " drag=" + player.drag_flag + " x=" + player.klick_panel_xpos +
          " y=" + player.klick_panel_ypos + " delay=" + player.frame_delay +
          " screenSHA256=" +
          hex(MessageDigest.getInstance("SHA-256").digest(eventScreen)));
    }
    byte[] screen = argb(player.screen);
    Files.write(Paths.get(args[2]), screen);
    System.out.println("JDK LIFECYCLE " + mode + " loaded=" + player.loaded +
        " frame=" + player.frame_number + " loops=" + player.work_loops +
        " repaints=" + player.repaints + " statuses=" + player.statuses +
        " screenSHA256=" +
        hex(MessageDigest.getInstance("SHA-256").digest(screen)) +
        (capture ? " first60RGBSHA256=" + hex(aggregate.digest()) : ""));
  }
}
