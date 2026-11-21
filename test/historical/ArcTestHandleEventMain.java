// Call the UNCHANGED original ArcTest.handleEvent(Event) directly, without
// an event queue or Frame. On "destroy", JDK System.exit(0) stops the JVM;
// the Magnolia opt-in VM may terminate its OWN bounded guest run, not its
// embedding browser/host process. Compile against SHA-pinned original class.
import java.awt.Event;
public final class ArcTestHandleEventMain {
    public static void main(String[] args) {
        ArcTest applet = new ArcTest();
        Event ordinary = new Event(null, Event.ACTION_EVENT, null);
        System.out.print("ordinary|");
        System.out.println(applet.handleEvent(ordinary));
        if (args.length == 0) {
            System.out.println("no-destroy");
            return;
        }
        System.out.println("before-destroy");
        applet.handleEvent(new Event(null, Event.WINDOW_DESTROY, null));
        System.out.println("UNREACHABLE after destroy");
    }
}
