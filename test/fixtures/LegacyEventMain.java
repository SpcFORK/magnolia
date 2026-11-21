import java.applet.Applet;
import java.awt.Event;

/** Small JDK-verifiable oracle for Component.handleEvent's legacy callbacks. */
@SuppressWarnings("deprecation")
public final class LegacyEventMain extends Applet {
    public int hits;
    public int checksum;

    @Override public boolean handleEvent(Event event) {
        if (event.id == 777) { checksum += 99; return true; }
        return super.handleEvent(event);
    }
    @Override public boolean keyDown(Event event, int key) {
        hits++;
        checksum += key;
        return true;
    }
    @Override public boolean mouseDown(Event event, int x, int y) {
        hits++;
        checksum += x * 100 + y;
        return true;
    }
    @Override public boolean mouseDrag(Event event, int x, int y) {
        hits++;
        checksum += x * 100 + y;
        return false;
    }
}
