import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Observe the exact same compiled finite Runnable class on real MicroEmulator.
// No timing or scheduling equivalence is inferred from completion ordering.
public final class ThreadSliceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "ThreadSliceMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "ThreadSliceMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "ThreadSliceMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:finite-thread-selected");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real finite task never selected Canvas");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
