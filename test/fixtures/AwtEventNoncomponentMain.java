// JDK-valid Event(Object,int,Object) with a plain Object target. The
// Magnolia direct-Event profile admits only null/constructed AWT Components.
import java.awt.Event;
public final class AwtEventNoncomponentMain {
    public static void main(String[] args) {
        Event event = new Event(new Object(), Event.ACTION_EVENT, null);
        System.out.println("JDK accepted non-Component target");
    }
}
