// Harness-owned Applet entrypoint; never substitutes or recompiles ArcCanvas.
// The SAME compiled class runs under JDK and Magnolia. Detached Canvas bounds
// are set explicitly, without BorderLayout, a peer or a repaint scheduler.
import java.applet.Applet;
import java.awt.Component;
import java.awt.Font;
import java.awt.Graphics;

public final class ArcCanvasPaintApplet extends Applet {
    @Override public void paint(Graphics graphics) {
        ArcCanvas canvas = new ArcCanvas();
        ((Component) canvas).setBounds(7, 11, 31, 21);
        canvas.startAngle = -45;
        canvas.endAngle = 225;
        canvas.paint(graphics);

        ((Component) canvas).setBounds(-4, 5, 20, 20);
        canvas.filled = true;
        canvas.startAngle = 350;
        canvas.endAngle = -450;
        canvas.font = new Font("Dialog", Font.BOLD, 16);
        canvas.paint(graphics);

        ((Component) canvas).setBounds(0, 0, 1, 1);
        canvas.filled = false;
        canvas.startAngle = 0;
        canvas.endAngle = 45;
        canvas.font = null; // Graphics.setFont(null) retains the prior Font.
        canvas.paint(graphics);

        ((Component) canvas).setBounds(0, 0, 0, 0);
        canvas.filled = true;
        canvas.startAngle = 18;
        canvas.endAngle = 0;
        canvas.paint(graphics); // Keep even no-op/negative-size draw CALLS.
    }
}
