import java.applet.Applet;
import java.awt.Graphics;
import java.awt.Color;

class BaseApplet extends Applet {
    public void init() { System.out.println("inherited init"); }
    public void paint(Graphics g) { g.setColor(Color.green); g.fillRect(1, 2, 8, 6); }
}

public class InheritedApplet extends BaseApplet {
    public void start() { System.out.println("child start"); }
}
