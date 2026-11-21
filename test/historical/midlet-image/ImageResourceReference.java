import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Executes the SAME owned ImageResourceMIDlet.class in the real pinned Java
// ME implementation with its three original resource bytes on the classpath.
// NO Magnolia/ImageBridge code, downloaded JAD or surrogate Acotel MIDlet.
public final class ImageResourceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "ImageResourceMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "ImageResourceMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "ImageResourceMIDlet$View".equals(
                            view.getClass().getName())) {
                        System.out.println("REFERENCE:image-fixture-selected");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real Image fixture never selected Canvas");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
