// Owned, tiny Graphics-only fixture for exact arc descriptor and renderer
// refusal tests; original ArcCanvas.paint has a separate pinned JDK oracle.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Graphics;
public final class AwtArcPaintApplet extends Applet {
    @Override public void paint(Graphics graphics) {
        graphics.setColor(Color.red);
        graphics.drawArc(2, 3, 30, 25, -45, 225);
        graphics.fillArc(4, 5, 0, -1, 350, -450);
    }
}
