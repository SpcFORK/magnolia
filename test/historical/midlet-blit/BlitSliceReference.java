import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

public final class BlitSliceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "BlitSliceMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "BlitSliceMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "BlitSliceMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:blit-fixture-selected");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real blit fixture never selected Canvas");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
