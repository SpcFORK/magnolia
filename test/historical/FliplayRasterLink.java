import java.awt.Component;
import java.awt.Graphics;
import java.awt.Image;
import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.IOException;

/**
 * Selection-only fixture: makes the original Fliplay frame loader and AWT
 * patch-renderer methods reachable without embedding the historical classes
 * or the movie. The link method itself is NEVER run. The bounded Oak oracle
 * supplies exact movie bytes and calls selected original methods directly.
 */
public final class FliplayRasterLink {
  private FliplayRasterLink() {}

  public static fliplay player() { return new fliplay(); }
  public static DataInputStream stream(byte[] bytes) {
    return new DataInputStream(new ByteArrayInputStream(bytes));
  }
  public static byte[] byteArray(int length) { return new byte[length]; }
  public static int[] intArray(int length) { return new int[length]; }
  public static flickframe[] frames(int length) { return new flickframe[length]; }
  public static Image image(Component component, int width, int height) {
    return component.createImage(width, height);
  }
  public static Graphics graphics(Image image) { return image.getGraphics(); }

  public static void link(byte[] header) throws IOException {
    fliplay p = player();
    DataInputStream in = stream(header);
    p.load_header(in);
    p.pixels = byteArray(320 * 240);
    p.red = byteArray(256);
    p.green = byteArray(256);
    p.blue = byteArray(256);
    p.row_start = intArray(240);
    p.row_end = intArray(240);
    p.frame = frames(61);
    p.frame[0] = new flickframe(4);
    p.bimage = image(p, 320, 240);
    p.bgraphics = graphics(p.bimage);
    p.load_frame(in, 0);
    p.extract_frame(0);
    p.process_first_frame();
    p.process_loop_frame(1);
    p.frame[0].draw_patches(p.bgraphics);
  }
}
