import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Negative: the selected extra triggerService() body must be linked before
// this constructor or startApp can print. Real MicroEmulator has getCurrent;
// Magnolia's narrow Display model intentionally does not.
public final class PaintServiceUnsupportedMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) {
            graphics.setColor(0xFF0000);
            graphics.fillRect(10, 12, 16, 9);
        }
    }
    private final View view = new View();
    protected void startApp() {
        System.out.println("BAD:start");
        Display.getDisplay(this).setCurrent(view);
    }
    public void triggerService() {
        System.out.println("BAD:service");
        Display.getDisplay(this).getCurrent();
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
