// JDK-valid public Event.target mutation, but deliberately outside the
// read-only target profile: refuse at preflight before printing anything.
import java.awt.Button;
import java.awt.Event;
public final class AwtEventTargetWriteMain {
    public static void main(String[] args) {
        Event event = new Event(new Button("button"), Event.ACTION_EVENT, null);
        event.target = null;
        System.out.println("JDK accepted Event.target write");
    }
}
