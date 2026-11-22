import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;
import javax.microedition.lcdui.Display;

// Real MicroEmulator accepts getCurrent(); the finite Magnolia model omits
// it. Even though only pauseApp uses it, Magnolia must refuse BEFORE BEGIN.
public final class UnsupportedPauseMIDlet extends MIDlet {
    private final Display display;
    public UnsupportedPauseMIDlet() {
        System.out.println("BEGIN");
        display = Display.getDisplay(this);
    }
    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:start");
    }
    protected void pauseApp() {
        System.out.println(display.getCurrent() == null);
    }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
