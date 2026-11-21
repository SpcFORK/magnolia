import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// A real MicroEmulator 2.0.4 reference for the SAME Java 8 fixture bytes
// passed separately to Magnolia. No JAD URL/network and no third-party game.
public final class KeyActionReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "KeyActionMIDlet"});
            for (int attempt = 0; attempt < 40; ++attempt) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null &&
                        "KeyActionMIDlet".equals(current.getClass().getName())) {
                    Displayable screen = Display.getDisplay(current).getCurrent();
                    if (screen instanceof Canvas &&
                            "KeyActionMIDlet$View".equals(screen.getClass().getName())) {
                        System.out.println("REFERENCE:original-MIDP-started");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("original MIDP callback/display never reached");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
