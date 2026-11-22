import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// notifyPaused in pauseApp must not remain latched after the AMS resumes.
public final class NotifiedPauseMIDlet extends MIDlet {
    private int starts;
    public NotifiedPauseMIDlet() { System.out.println("TRACE:notify-construct"); }
    protected void startApp() throws MIDletStateChangeException {
        starts++;
        System.out.println(starts == 1
            ? "TRACE:notify-start" : "TRACE:notify-resume");
    }
    protected void pauseApp() {
        System.out.println("TRACE:notify-pause");
        notifyPaused();
    }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException {
        System.out.println("TRACE:notify-destroy");
        notifyDestroyed();
    }
}
