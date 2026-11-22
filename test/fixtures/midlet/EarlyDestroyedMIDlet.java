import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Self-notified destruction during startup must suppress later AMS callbacks.
public final class EarlyDestroyedMIDlet extends MIDlet {
    public EarlyDestroyedMIDlet() { System.out.println("TRACE:early-construct"); }
    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:early-stop");
        notifyDestroyed();
    }
    protected void pauseApp() { System.out.println("BAD:paused-after-destroy"); }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException {
        System.out.println("BAD:destroyed-twice");
    }
}
