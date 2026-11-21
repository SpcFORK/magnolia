import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Proves FACE_MONOSPACE is a valid real MIDP request even though the narrow
// Magnolia font profile correctly refuses it instead of guessing metrics.
public final class UnsupportedFontReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "UnsupportedFontMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "UnsupportedFontMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "UnsupportedFontMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:alternate-font-accepted");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real MIDP alternative font did not finish");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
