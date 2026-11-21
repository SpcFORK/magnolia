import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

public final class CanvasSizeReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "CanvasSizeMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "CanvasSizeMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "CanvasSizeMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:canvas-geometry-selected");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real Canvas geometry fixture not selected");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
