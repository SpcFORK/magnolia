import java.applet.Applet;
import java.awt.Component;
import java.awt.Graphics;
import java.awt.Image;
import java.awt.image.IndexColorModel;
import java.awt.image.MemoryImageSource;

/** Small, non-historical Java 8 image and patch-island regression. */
public class RasterProbeApplet extends Applet {
  Image canvas;
  Image source;
  // Force the inherited Component descriptor: the bounded VM deliberately
  // does not intercept arbitrary user-class createImage overloads.
  private Image image(int width, int height) {
    return ((Component)this).createImage(width, height);
  }
  private Image image(MemoryImageSource producer) {
    return ((Component)this).createImage(producer);
  }
  @Override public void init() {
    byte[] red = new byte[256], green = new byte[256], blue = new byte[256];
    red[0] = 10; green[0] = 20; blue[0] = 30;
    red[1] = 40; green[1] = 50; blue[1] = 60;
    red[2] = 70; green[2] = 80; blue[2] = 90;
    red[255] = (byte)250; green[255] = (byte)251; blue[255] = (byte)252;
    IndexColorModel palette = new IndexColorModel(8, 256, red, green, blue);
    red[1] = 99; // Model must retain the ORIGINAL palette snapshot.
    byte[] indices = new byte[16];
    indices[5] = 0; indices[6] = 1;
    indices[9] = (byte)255; indices[10] = 2;
    MemoryImageSource producer = new MemoryImageSource(2, 2, palette,
                                                      indices, 5, 4);
    source = image(producer);
    indices[5] = 2; // Already-created Image retains its copied pixels.
    canvas = image(4, 3);
    Graphics g = canvas.getGraphics();
    g.setColor(java.awt.Color.red);
    g.fillRect(0, 0, 4, 3);
    System.out.println(g.drawImage(source, 1, 0, this));
    Image patch = image(2, 2);
    Graphics pg = patch.getGraphics();
    System.out.println(pg.drawImage(source, 0, 0, null));
    patch.flush();
    System.out.println(g.drawImage(patch, 2, 1, null));
    g.setColor(java.awt.Color.blue);
    g.drawRect(0, 0, 3, 2);
    System.out.println("raster-ready");
  }
}
