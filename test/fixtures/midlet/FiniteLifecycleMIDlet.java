import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Graphics;

// Compile against the REAL MicroEmulator MIDP API, not Magnolia's model.
// The independent AMS harness calls start/pause/start/destroy on one MIDlet.
public final class FiniteLifecycleMIDlet extends MIDlet {
    private final Display display;
    private final Canvas view;
    private int starts;

    public FiniteLifecycleMIDlet() {
        System.out.println("TRACE:construct");
        display = Display.getDisplay(this);
        view = new View();
    }

    static final class View extends Canvas {
        public View() { super(); }
        protected void paint(Graphics graphics) { }
    }

    protected void startApp() throws MIDletStateChangeException {
        starts++;
        System.out.println(starts == 1 ? "TRACE:start-first" : "TRACE:start-resume");
        System.out.println(display == Display.getDisplay(this)
            ? "TRACE:display-same" : "TRACE:display-changed");
        display.setCurrent(view);
    }

    protected void pauseApp() { System.out.println("TRACE:pause"); }

    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException {
        System.out.println(unconditional
            ? "TRACE:destroy-true" : "TRACE:destroy-false");
        notifyDestroyed();
    }
}
