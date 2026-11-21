import java.lang.reflect.Field;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.lcdui.Font;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Reference ONLY: the unchanged game's real Canvas Font fields on pinned
// MicroEmulator. This class does NOT run the original game inside Magnolia.
public final class OriginalAcotelFontReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "ACO_XmasGifts_s"});
            Displayable screen = null;
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "ACO_XmasGifts_s".equals(
                        current.getClass().getName())) {
                    screen = Display.getDisplay(current).getCurrent();
                    if (screen != null && "DS_Effects".equals(
                            screen.getClass().getName())) break;
                }
                Thread.sleep(50L);
            }
            if (screen == null || !"DS_Effects".equals(screen.getClass().getName()))
                throw new AssertionError("original Canvas not selected");
            String[] names = {"_fld6", "_fld7", "_fld8"};
            int[] sizes = {8, 0, 16};
            for (int i = 0; i < names.length; i++) {
                Field field = screen.getClass().getDeclaredField(names[i]);
                field.setAccessible(true);
                Font font = (Font)field.get(screen);
                if (font == null || font.getFace() != 0 ||
                        font.getStyle() != 1 || font.getSize() != sizes[i])
                    throw new AssertionError("original font field not initialized");
                System.out.println("ORIGINAL:font-size=" + sizes[i] +
                    ",height=" + font.getHeight() +
                    ",xmas-width=" + font.stringWidth("Xmas Gifts"));
            }
            System.out.println("ORIGINAL:font-reference-finished");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
