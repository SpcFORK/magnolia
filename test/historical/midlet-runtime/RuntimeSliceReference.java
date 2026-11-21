import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Run the unchanged fixture on the pinned real MIDP API and JDK 11. Do not
// substitute a Java SE Applet for the MIDlet or require a fabricated heap.
public final class RuntimeSliceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            if (!RuntimeSliceMIDlet.low(4096L) ||
                RuntimeSliceMIDlet.low(5000L) ||
                RuntimeSliceMIDlet.low(16777216L))
                throw new AssertionError("JDK long-branch reference changed");
            Headless.main(new String[]{"--rms", "memory", "RuntimeSliceMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "RuntimeSliceMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "RuntimeSliceMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:runtime-fixture-selected");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real Runtime fixture never selected Canvas");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
