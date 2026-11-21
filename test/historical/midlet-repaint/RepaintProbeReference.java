import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Uses the emulator's own MIDlet classloader; reflection reads only counters
// on the selected original fixture. The reference does not expose a general
// Java SE Applet or run the unchanged Acotel game inside Magnolia.
public final class RepaintProbeReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "RepaintProbeMIDlet"});
            long deadline = System.currentTimeMillis() + 2000;
            Canvas selected = null;
            MIDlet owner = null;
            while (System.currentTimeMillis() < deadline) {
                owner = MIDletBridge.getCurrentMIDlet();
                if (owner != null && "RepaintProbeMIDlet".equals(owner.getClass().getName())) {
                    Displayable candidate = Display.getDisplay(owner).getCurrent();
                    if (candidate instanceof Canvas &&
                        "RepaintProbeMIDlet$View".equals(candidate.getClass().getName())) {
                        selected = (Canvas) candidate;
                        break;
                    }
                }
                Thread.sleep(10L);
            }
            if (selected == null) throw new AssertionError("Canvas was never selected");
            java.lang.reflect.Field paints = selected.getClass().getField("paints");
            java.lang.reflect.Field shown = selected.getClass().getField("shown");
            java.lang.reflect.Field clip = selected.getClass().getField("clipWidth");
            java.lang.reflect.Field thread = selected.getClass().getField("paintThread");
            java.lang.reflect.Field detached = owner.getClass().getField("detachedPaints");
            java.lang.reflect.Field requested = owner.getClass().getField("selectionRequestedPaints");
            deadline = System.currentTimeMillis() + 1000;
            while (paints.getInt(selected) < 1 && System.currentTimeMillis() < deadline)
                Thread.sleep(10L);
            System.out.println("REFERENCE:before-repaint:shown=" + shown.getInt(selected) +
                ";paints=" + paints.getInt(selected) +
                ";detached=" + detached.getInt(null) +
                ";selectionRequested=" + requested.getInt(null) +
                ";clip=" + clip.getInt(selected) +
                ";thread=" + thread.get(selected) +
                ";selected=" + (Display.getDisplay(owner).getCurrent() == selected));
            selected.repaint();
            selected.repaint();
            System.out.println("REFERENCE:requests-returned:paints=" + paints.getInt(selected));
            selected.serviceRepaints();
            System.out.println("REFERENCE:after-service:paints=" + paints.getInt(selected));
            if (shown.getInt(selected) != 1 || paints.getInt(selected) != 2 ||
                detached.getInt(null) != 0 || requested.getInt(null) != 0 ||
                clip.getInt(selected) != 176 || !"event-thread".equals(thread.get(selected)))
                throw new AssertionError("bounded repaint observation differs");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
