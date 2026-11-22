import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Negative bound: the service profile may paint ONCE after the initial
// callback. A second service call must fail rather than returning success.
public final class PaintServiceTwiceMIDlet extends MIDlet {
    private View view;
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) {
            graphics.setColor(0xFF0000);
            graphics.fillRect(10, 12, 16, 9);
        }
    }
    protected void startApp() {
        view = new View();
        Display.getDisplay(this).setCurrent(view);
    }
    public void triggerService() {
        view.repaint();
        view.serviceRepaints();
        view.repaint();
        view.serviceRepaints();
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
