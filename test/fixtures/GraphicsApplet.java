// Real AWT draw calls used to test Magnolia's first graphical applet bridge.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Graphics;

public class GraphicsApplet extends Applet {
    @Override public void paint(Graphics g) {
        g.setColor(Color.red);
        g.fillRect(4, 5, 30, 12);
        g.setColor(new Color(0, 128, 255));
        g.drawString("Oak canvas", 2, 38);
        g.drawLine(2, 41, 62, 41);
    }
}
