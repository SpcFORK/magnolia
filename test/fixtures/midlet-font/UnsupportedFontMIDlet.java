import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Font;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// A valid MIDP Font request outside the measured Acotel default-device
// profile; Magnolia refuses before output instead of fabricating metrics.
public final class UnsupportedFontMIDlet extends MIDlet {
    private final Display display = Display.getDisplay(this);
    private final View view = new View();
    public static final class View extends Canvas {
        public View() { super(); }
        protected void paint(Graphics graphics) { }
    }
    protected void startApp() throws MIDletStateChangeException {
        Font other = Font.getFont(32, 1, 8); // FACE_MONOSPACE
        System.out.println("BAD:unmeasured-font-accepted");
        System.out.println(other.getHeight());
        display.setCurrent(view);
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
