import java.applet.Applet;
import java.awt.Color;
import java.awt.Graphics;
import java.awt.Image;

/** Java source executed as bytecode in Magnolia (NOT the JDK oracle). */
public final class PolygonArrowFixture extends Applet {
  public Image[] images;
  public void generate(int first, int last) {
    if (first < 0 || last <= first || last > 3000 || last - first > 200)
      return;
    images = new Image[last - first];
    int count = 0;
    for (int top = -3; top <= 9; top += 3)
      for (int height = 2; height <= 60; height += 2)
        for (int base = 3; base <= 48; base += 5)
          for (int side = -1; side <= 1; side += 2) {
            if (count < first || count >= last) { count++; continue; }
            Image icon = ((java.awt.Component)this).createImage(16, 16);
            Graphics g = icon.getGraphics();
            int[] xs = {base, base + side * height, base};
            int[] ys = {top, top + height / 2, top + height};
            g.setColor(Color.RED);
            g.fillPolygon(xs, ys, 3);
            images[count++ - first] = icon;
          }
  }
  public void invalid() {
    Image icon = ((java.awt.Component)this).createImage(16, 16);
    icon.getGraphics().fillPolygon(new int[]{2, 7, 15}, new int[]{2, 5, 8}, 3);
  }
  public void extreme() {
    Image icon = ((java.awt.Component)this).createImage(16, 16);
    icon.getGraphics().fillPolygon(new int[]{2147483647, 1, 2147483647},
        new int[]{-2147483648, 3, 100}, 3);
  }
  public void shortArrays() {
    Image icon = ((java.awt.Component)this).createImage(16, 16);
    icon.getGraphics().fillPolygon(new int[]{1, 2}, new int[]{1, 2}, 3);
  }
}
