import java.awt.Color;
import java.awt.Graphics;
import java.awt.Polygon;
import java.awt.image.BufferedImage;
import java.io.ByteArrayOutputStream;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.security.MessageDigest;

/** Independently generated JDK 11 arrow-icon oracle for the supported fillPolygon slice. */
public final class PolygonArrowJdk {
  private static String hex(byte[] data) {
    StringBuilder s = new StringBuilder();
    for (byte b : data) s.append(String.format("%02x", b & 255));
    return s.toString();
  }
  public static void main(String[] args) throws Exception {
    if (args.length != 1) throw new IllegalArgumentException("OUTPUT.argb required");
    int count = 0;
    ByteArrayOutputStream raw = new ByteArrayOutputStream();
    for (int top = -3; top <= 9; top += 3)
      for (int height = 2; height <= 60; height += 2)
        for (int base = 3; base <= 48; base += 5)
          for (int side = -1; side <= 1; side += 2) {
            BufferedImage icon = new BufferedImage(16, 16, BufferedImage.TYPE_INT_ARGB);
            int[] xs = {base, base + side * height, base};
            int[] ys = {top, top + height / 2, top + height};
            Graphics g = icon.getGraphics();
            g.setColor(Color.RED);
            g.fillPolygon(xs, ys, 3);
            g.dispose();
            for (int y = 0; y < 16; ++y)
              for (int x = 0; x < 16; ++x) {
                int pixel = icon.getRGB(x, y);
                raw.write(pixel >>> 24); raw.write(pixel >>> 16);
                raw.write(pixel >>> 8); raw.write(pixel);
              }
            count++;
          }
    if (count != 3000) throw new AssertionError("case count: " + count);
    byte[] data = raw.toByteArray();
    Files.write(Paths.get(args[0]), data);
    System.out.println("JDK ARROWS cases=" + count + " bytes=" + data.length +
        " SHA256=" + hex(MessageDigest.getInstance("SHA-256").digest(data)));
  }
}
