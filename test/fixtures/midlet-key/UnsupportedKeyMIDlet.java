import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Real MicroEmulator returns 0 for this out-of-fixture key. Magnolia's finite
// checked-device profile deliberately refuses instead of inventing support.
public final class UnsupportedKeyMIDlet extends MIDlet {
    private final View view = new View();
    public static final class View extends Canvas {
        public View() { super(); }
        protected void paint(Graphics graphics) { }
    }
    protected void startApp() throws MIDletStateChangeException {
        int action = view.getGameAction(1000);
        System.out.println("BAD:unexpected-key");
        System.out.println(action);
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
