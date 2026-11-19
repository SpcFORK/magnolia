import java.applet.Applet;
import java.awt.Graphics;
import java.awt.Color;

public class PublicBaseApplet extends Applet {
    public void init() { System.out.println("public base init"); }
    public void paint(Graphics g) { g.setColor(Color.red); g.fillRect(2, 3, 6, 4); }
}
