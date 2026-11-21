import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.security.MessageDigest;
import java.awt.image.BufferedImage;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;
import org.microemu.device.DeviceDisplay;
import org.microemu.device.DeviceFactory;
import org.microemu.device.MutableImage;
import org.microemu.device.j2se.J2SEMutableImage;

// JDK 11 observer on REAL MicroEmulator 2.0.4; fixture class is loaded by
// emulator's own classloader. Never launches the unchanged original game.
public final class PaintServiceReference {
    // J2SEMutableImage.getData() caches its PixelGrabber after the first read;
    // a second call can return the stale FIRST frame. Read the same real
    // device BufferedImage into a FRESH array after each service boundary.
    private static int[] pixels(MutableImage image) {
        BufferedImage backing = (BufferedImage)((J2SEMutableImage)image).getImage();
        return backing.getRGB(0, 0, 176, 220, null, 0, 176);
    }
    private static String frameHash(int[] backing, int width, int expectedColor,
                                    int x0, int y0, int w, int h) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        if (width != 176 || backing.length != 176 * 220)
            throw new AssertionError("unexpected device framebuffer extent");
        int colors = 0;
        for (int y = 0; y < 176; y++) for (int x = 0; x < 176; x++) {
            int pixel = backing[(22 + y) * width + x];
            int expected = x >= x0 && x < x0 + w && y >= y0 && y < y0 + h
                ? expectedColor : 0xFF000000;
            if (pixel != expected) throw new AssertionError("Canvas pixel mismatch at " + x + "," + y);
            if (pixel == expectedColor) colors++;
            digest.update((byte)(pixel >>> 24));
            digest.update((byte)(pixel >>> 16));
            digest.update((byte)(pixel >>> 8));
            digest.update((byte)pixel);
        }
        if (colors != w * h) throw new AssertionError("rectangle missing");
        StringBuilder hex = new StringBuilder();
        for (byte value : digest.digest()) {
            int b = value & 255;
            hex.append(Character.forDigit(b >>> 4, 16));
            hex.append(Character.forDigit(b & 15, 16));
        }
        return hex.toString();
    }
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no untrusted inputs");
            Headless.main(new String[]{"--rms", "memory", "PaintServiceMIDlet"});
            long deadline = System.currentTimeMillis() + 2000;
            MIDlet owner = null;
            Canvas selected = null;
            while (System.currentTimeMillis() < deadline) {
                owner = MIDletBridge.getCurrentMIDlet();
                if (owner != null && "PaintServiceMIDlet".equals(owner.getClass().getName())) {
                    Displayable candidate = Display.getDisplay(owner).getCurrent();
                    if (candidate instanceof Canvas &&
                        "PaintServiceMIDlet$View".equals(candidate.getClass().getName())) {
                        selected = (Canvas) candidate;
                        break;
                    }
                }
                Thread.sleep(5L);
            }
            if (selected == null) throw new AssertionError("owned Canvas not selected");
            Field paints = selected.getClass().getField("paints");
            Field shown = selected.getClass().getField("shown");
            Field clip = selected.getClass().getField("clipWidth");
            Field detached = owner.getClass().getField("detachedPaints");
            deadline = System.currentTimeMillis() + 1000;
            while (paints.getInt(selected) < 1 && System.currentTimeMillis() < deadline)
                Thread.sleep(5L);
            DeviceDisplay display = DeviceFactory.getDevice().getDeviceDisplay();
            if (display == null || display.getWidth() != 176 || display.getHeight() != 176)
                throw new AssertionError("default device geometry differs");
            MutableImage image = display.getDisplayImage();
            if (image == null) throw new AssertionError("device has no backing image");
            String first = frameHash(pixels(image), 176, 0xFFFF0000, 10, 12, 16, 9);
            System.out.println("REFERENCE:first=" + first + ";shown=" + shown.getInt(selected) +
                ";paints=" + paints.getInt(selected) + ";clip=" + clip.getInt(selected) +
                ";detached=" + detached.getInt(null));
            if (paints.getInt(selected) != 1 || shown.getInt(selected) != 1 ||
                clip.getInt(selected) != 176 || detached.getInt(null) != 0)
                throw new AssertionError("first frame scheduling differs");
            Method trigger = owner.getClass().getMethod("triggerService");
            trigger.invoke(owner);
            // A serviceRepaints return should leave this particular selected
            // finite request sequence complete. The event-thread paint may
            // race the caller's request-returned line, not the final result.
            deadline = System.currentTimeMillis() + 1000;
            while (paints.getInt(selected) < 2 && System.currentTimeMillis() < deadline)
                Thread.sleep(5L);
            String second = frameHash(pixels(display.getDisplayImage()), 176,
                                      0xFF00FF00, 32, 41, 13, 7);
            System.out.println("REFERENCE:second=" + second + ";shown=" + shown.getInt(selected) +
                ";paints=" + paints.getInt(selected) + ";clip=" + clip.getInt(selected));
            if (paints.getInt(selected) != 2 || shown.getInt(selected) != 1 ||
                clip.getInt(selected) != 176 || first.equals(second))
                throw new AssertionError("bounded serviceRepaints observation differs");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
