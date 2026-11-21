import java.awt.Image;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.MessageDigest;
import java.lang.reflect.Field;

/**
 * JDK 11 reference for the ORIGINAL Fliplay init() bytecode, not Applet
 * lifecycle execution. A displayless Applet cannot be constructed normally:
 * bypass only its host constructor, seed the original constructor defaults,
 * supply explicit parameters, and override createImage(II) with a synchronous
 * JDK BufferedImage. All init() logic and Graphics calls remain original.
 */
public final class FliplayInitJdk {
  private static String mode;
  public static final class HeadlessApplet extends fliplay {
    @Override public String getParameter(String name) {
      if (mode.equals("themed")) {
        if (name.equals("PanelColor")) return "2";
        if (name.equals("TextColor")) return "red";
        if (name.equals("Block")) return "32";
        if (name.equals("Speed")) return "71";
        if (name.equals("Box")) return "yes";
        if (name.equals("Reverse")) return "yes";
        if (name.equals("Keyboard")) return "no";
        if (name.equals("File")) return "moonclouds.flc";
      }
      return null;
    }
    @Override public Image createImage(int width, int height) {
      return new BufferedImage(width, height, BufferedImage.TYPE_INT_ARGB);
    }
  }
  private static HeadlessApplet newWithoutDisplay() throws Exception {
    Class<?> unsafeClass = Class.forName("sun.misc.Unsafe");
    Field singleton = unsafeClass.getDeclaredField("theUnsafe");
    singleton.setAccessible(true);
    Object unsafe = singleton.get(null);
    HeadlessApplet player = (HeadlessApplet)unsafeClass
        .getMethod("allocateInstance", Class.class)
        .invoke(unsafe, HeadlessApplet.class);
    // Values actually written by the unmodified fliplay.<init>()V. Bypassing
    // the Applet host constructor is unavoidable in a headless JDK process.
    player.frame_data_len = -1;
    player.klick_dim = 24;
    player.loop_flag = true;
    player.klick_panel_mode = true;
    player.keyboard_control = true;
    return player;
  }
  private static String hex(byte[] data) {
    StringBuilder output = new StringBuilder();
    for (byte b : data) output.append(String.format("%02x", b & 255));
    return output.toString();
  }
  public static void main(String[] args) throws Exception {
    if (args.length != 2 || (!args[0].equals("default") && !args[0].equals("themed")))
      throw new IllegalArgumentException("default|themed OUTPUT.argb required");
    mode = args[0];
    HeadlessApplet player = newWithoutDisplay();
    player.init(); // invokes the original method, including two fillPolygon calls
    if (!(player.klick_image instanceof BufferedImage) ||
        player.klick_image.getWidth(null) != 72 ||
        player.klick_image.getHeight(null) != 24)
      throw new AssertionError("original init did not create a 72x24 icon");
    BufferedImage icon = (BufferedImage)player.klick_image;
    ByteArrayOutputStream raw = new ByteArrayOutputStream(72 * 24 * 4);
    for (int y = 0; y < 24; ++y) for (int x = 0; x < 72; ++x) {
      int pixel = icon.getRGB(x, y);
      raw.write(pixel >>> 24); raw.write(pixel >>> 16);
      raw.write(pixel >>> 8); raw.write(pixel);
    }
    byte[] bytes = raw.toByteArray();
    Files.write(Paths.get(args[1]), bytes);
    System.out.println("JDK INIT " + mode + " iconSHA256=" +
        hex(MessageDigest.getInstance("SHA-256").digest(bytes)) +
        " block=" + player.block_height + " delay=" + player.frame_delay +
        " box=" + player.box_flag + " reverse=" + player.reverse_flag +
        " keyboard=" + player.keyboard_control +
        " file=" + player.file_name);
  }
}
