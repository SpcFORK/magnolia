import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.IOException;

/**
 * Java 8 reachability anchor for the separately supplied, unmodified Fliplay
 * class files. This class contains no animation data or decoder algorithm.
 * Magnolia selects link(byte[]) for preflight, then the Oak oracle invokes the
 * compiled helper methods and original fliplay methods on bounded input.
 * link itself is never run: frame objects are prepared by the oracle before
 * calling the original load_frame/extract_frame methods.
 */
public final class FliplayFrameLink {
  private FliplayFrameLink() {}

  public static fliplay player() { return new fliplay(); }

  public static DataInputStream stream(byte[] bytes) {
    return new DataInputStream(new ByteArrayInputStream(bytes));
  }

  public static byte[] pixelArray(int length) { return new byte[length]; }
  public static int[] rowArray(int length) { return new int[length]; }
  public static flickframe[] frames(int length) { return new flickframe[length]; }

  public static void link(byte[] bytes) throws IOException {
    fliplay p = player();
    DataInputStream input = stream(bytes);
    p.load_header(input);
    p.pixels = pixelArray(320 * 240);
    p.red = pixelArray(256);
    p.green = pixelArray(256);
    p.blue = pixelArray(256);
    p.row_start = rowArray(240);
    p.row_end = rowArray(240);
    p.frame = frames(61);
    p.load_frame(input, 0);
    p.extract_frame(0);
  }
}
