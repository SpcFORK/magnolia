// JDK permits an arbitrary Event.arg object; Magnolia bounds the directly
// constructed guest Event to null or short ASCII strings, not a general AWT.
import java.awt.Event;
public final class AwtEventNonstringArgMain {
    public static void main(String[] args) {
        new Event(null, Event.ACTION_EVENT, new Object());
        System.out.println("JDK accepted non-String Event argument");
    }
}
