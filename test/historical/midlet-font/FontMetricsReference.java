import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.lcdui.Font;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// This harness uses actual pinned MicroEmulator classes only. The fixture
// FontMetricsMIDlet.class is the SAME byte array subsequently run in Magnolia.
// Tables are extracted from the real device in the isolated reference JVM;
// no MicroEmulator implementation is copied into Magnolia or this project.
public final class FontMetricsReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "FontMetricsMIDlet"});
            boolean selected = false;
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "FontMetricsMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable screen = Display.getDisplay(current).getCurrent();
                    if (screen instanceof Canvas && "FontMetricsMIDlet$View".equals(
                            screen.getClass().getName())) {
                        selected = true;
                        break;
                    }
                }
                Thread.sleep(50L);
            }
            if (!selected) throw new AssertionError("fixture never selected Canvas");
            for (int size : new int[]{8, 0, 16}) {
                Font real = Font.getFont(0, 1, size);
                StringBuilder widths = new StringBuilder();
                for (char c = 32; c <= 126; c++) {
                    int w = real.charWidth(c);
                    if (w < 0 || w > 15) throw new AssertionError("unbounded font table");
                    widths.append(Character.toUpperCase(Character.forDigit(w, 16)));
                }
                System.out.println("REFERENCE:ascii-size=" + size + ":" + widths);
            }
            System.out.println("REFERENCE:font-finished");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
