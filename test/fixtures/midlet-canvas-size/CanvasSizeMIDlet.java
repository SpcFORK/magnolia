import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Pinned *default MicroEmulator device*, not all MIDP displays. Observe
// geometry both before and after an AMS display selection on identical bytes.
public final class CanvasSizeMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }
    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:geometry-begin");
        View view = new View();
        System.out.println(view.getWidth());
        System.out.println(view.getHeight());
        Display.getDisplay(this).setCurrent(view);
        System.out.println(view.getWidth());
        System.out.println(view.getHeight());
        System.out.println("TRACE:geometry-end");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
