import java.lang.reflect.Field;
import java.security.MessageDigest;
import java.util.HashSet;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;
import org.microemu.device.Device;
import org.microemu.device.DeviceDisplay;
import org.microemu.device.DeviceFactory;
import org.microemu.device.MutableImage;

/**
 * Reference-only original Acotel JAR probe, NOT a Magnolia game implementation.
 * Run in a short-lived JVM with a separately hash-checked MicroEmulator and
 * original JAR on the classpath, private home/in-memory RMS, and an outer wall
 * timeout. No downloaded/JAD URL is used. Extract MIDlet/Canvas and device
 * backing-buffer observations; never mistake those for verified gameplay.
 */
public final class AcotelMicroemuReference {
    private static String hash(int[] raster) throws Exception {
        MessageDigest md = MessageDigest.getInstance("SHA-256");
        for (int value : raster) {
            md.update((byte)(value >>> 24));
            md.update((byte)(value >>> 16));
            md.update((byte)(value >>> 8));
            md.update((byte)value);
        }
        byte[] digest = md.digest();
        StringBuilder result = new StringBuilder();
        for (byte value : digest) {
            int v = value & 255;
            result.append(Character.forDigit(v >>> 4, 16));
            result.append(Character.forDigit(v & 15, 16));
        }
        return result.toString();
    }

    public static void main(String[] args) {
        try {
            if (args.length != 0)
                throw new IllegalArgumentException("no untrusted arguments/URLs");
            Headless.main(new String[] {"--rms", "memory", "ACO_XmasGifts_s"});
            MIDlet game = MIDletBridge.getCurrentMIDlet();
            if (game == null || !"ACO_XmasGifts_s".equals(game.getClass().getName()))
                throw new IllegalStateException("original MIDlet not constructed");
            Device device = DeviceFactory.getDevice();
            if (device == null || device.getDeviceDisplay() == null)
                throw new IllegalStateException("emulator device absent");
            DeviceDisplay display = device.getDeviceDisplay();
            for (int i = 0; i < 7; i++) {
                if (i > 0) Thread.sleep(250L);
                Displayable current = Display.getDisplay(game).getCurrent();
                MutableImage image = display.getDisplayImage();
                int[] pixels = image == null ? null : image.getData();
                String currentName = current == null ? "null" : current.getClass().getName();
                String offscreen = "null";
                if (current != null && "DS_Effects".equals(currentName)) {
                    Field field = current.getClass().getDeclaredField("_fld9");
                    field.setAccessible(true);
                    Image frame = (Image) field.get(current);
                    if (frame != null) {
                        int w = frame.getWidth(), h = frame.getHeight();
                        if (w < 1 || h < 1 || (long) w * h > 100000L)
                            throw new IllegalStateException("offscreen image exceeds bound");
                        int[] argb = new int[w * h];
                        frame.getRGB(argb, 0, w, 0, 0, w, h);
                        HashSet<Integer> colors = new HashSet<Integer>();
                        for (int color : argb) colors.add(color);
                        offscreen = w + "x" + h + ",colors=" + colors.size() +
                            ",sha256=" + hash(argb);
                    }
                }
                System.out.println("REF:tick=" + i + ",midlet=" +
                    game.getClass().getName() + ",current=" + currentName +
                    ",device=" + display.getWidth() + "x" + display.getHeight() +
                    ",pixels=" + (pixels == null ? "null" : pixels.length) +
                    ",argbSHA256=" + (pixels == null ? "null" : hash(pixels)) +
                    ",gameOffscreen=" + offscreen);
            }
            System.out.println("REF:observed-original-startup");
            System.exit(0); // emulator/game threads must not outlive this probe
        } catch (Throwable error) {
            System.err.println("REF:FAILED " + error.toString());
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
