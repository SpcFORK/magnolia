import java.lang.reflect.Field;
import java.security.MessageDigest;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;
import org.microemu.device.DeviceDisplay;
import org.microemu.device.DeviceFactory;
import org.microemu.device.MutableImage;

// Real MicroEmulator/JDK observer; does not invoke Magnolia or the original
// game. Uses the emulator's own MIDlet classloader for the selected fixture.
public final class PaintSliceReference {
    private static String argbSha256(int[] data, int width, int originY) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-256");
        for (int y = 0; y < 176; y++) for (int x = 0; x < 176; x++) {
            int value = data[(y + originY) * width + x];
            int expected = x >= 10 && x < 26 && y >= 12 && y < 21 ?
                0xFFFF0000 : 0xFF000000;
            if (value != expected) throw new AssertionError("unexpected Canvas pixel at " + x + "," + y);
            md.update((byte)(value >>> 24));
            md.update((byte)(value >>> 16));
            md.update((byte)(value >>> 8));
            md.update((byte)value);
        }
        StringBuilder hex = new StringBuilder();
        for (byte value : md.digest()) {
            int b = value & 255;
            hex.append(Character.forDigit(b >>> 4, 16));
            hex.append(Character.forDigit(b & 15, 16));
        }
        return hex.toString();
    }
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "PaintSliceMIDlet"});
            long deadline = System.currentTimeMillis() + 2000;
            MIDlet owner = null;
            Canvas selected = null;
            while (System.currentTimeMillis() < deadline) {
                owner = MIDletBridge.getCurrentMIDlet();
                if (owner != null && "PaintSliceMIDlet".equals(owner.getClass().getName())) {
                    Displayable candidate = Display.getDisplay(owner).getCurrent();
                    if (candidate instanceof Canvas &&
                        "PaintSliceMIDlet$View".equals(candidate.getClass().getName())) {
                        selected = (Canvas) candidate;
                        break;
                    }
                }
                Thread.sleep(5L);
            }
            if (selected == null) throw new AssertionError("fixture Canvas not selected");
            Field shown = selected.getClass().getField("shown");
            Field paints = selected.getClass().getField("paints");
            Field clip = selected.getClass().getField("clipWidth");
            deadline = System.currentTimeMillis() + 1000;
            while ((shown.getInt(selected) != 1 || paints.getInt(selected) != 1) &&
                   System.currentTimeMillis() < deadline) Thread.sleep(5L);
            Field detached = owner.getClass().getField("detachedPaints");
            Field requested = owner.getClass().getField("selectionRequestedPaints");
            DeviceDisplay display = DeviceFactory.getDevice().getDeviceDisplay();
            MutableImage image = display.getDisplayImage();
            if (image == null) throw new AssertionError("no device image");
            int[] pixels = image.getData();
            int width = display.getWidth();
            if (width != 176 || pixels.length < 176 * 176 || pixels.length % width != 0)
                throw new AssertionError("device geometry outside bounded profile");
            int reds = 0, minX = width, minY = pixels.length / width,
                maxX = -1, maxY = -1, blacks = 0;
            for (int i = 0; i < pixels.length; i++) {
                if (pixels[i] == 0xFFFF0000) {
                    int x = i % width, y = i / width;
                    reds++;
                    if (x < minX) minX = x;
                    if (x > maxX) maxX = x;
                    if (y < minY) minY = y;
                    if (y > maxY) maxY = y;
                }
                if (pixels[i] == 0xFF000000) blacks++;
            }
            if (shown.getInt(selected) != 1 || paints.getInt(selected) != 1 ||
                detached.getInt(null) != 0 || requested.getInt(null) != 0 ||
                clip.getInt(selected) != 176 || reds != 16 * 9 ||
                blacks != 176 * 176 - reds ||
                minX != 10 || maxX != 25 || minY != 34 || maxY != 42)
                throw new AssertionError("reference paint state or pixels differ");
            // MicroEmulator's 38,720-pixel backing includes a 22-row top UI
            // and a 22-row bottom UI. Compare the actual Canvas rectangle,
            // not the entire device backing image, to the native 176x176 host.
            String frame = argbSha256(pixels, width, 22);
            System.out.println("REFERENCE:shown=" + shown.getInt(selected) +
                ";paints=" + paints.getInt(selected) +
                ";detached=" + detached.getInt(null) +
                ";selectedCount=" + requested.getInt(null) +
                ";clip=" + clip.getInt(selected) +
                ";device=" + display.getWidth() + "x" + display.getHeight() +
                ";backing=" + pixels.length +
                ";reds=" + reds + ";black=" + blacks +
                ";box=" + minX + "," + minY + "," + maxX + "," + maxY +
                ";canvasArgbSha256=" + frame);
            if (frame.length() != 64)
                throw new AssertionError("reference paint state or pixels differ");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
