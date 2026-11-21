import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Java SE AWT is present in this JDK, not in the narrow CLDC/MIDP host.
public final class DesktopAPIMIDlet extends MIDlet {
    public DesktopAPIMIDlet() { System.out.println("BEGIN"); }
    protected void startApp() throws MIDletStateChangeException { }
    protected void pauseApp() {
        java.awt.Point p = new java.awt.Point(3, 4);
        System.out.println(p.x);
    }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { }
}
