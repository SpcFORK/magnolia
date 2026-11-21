import java.security.MessageDigest;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

/** Real MicroEmulator Image decoder and ORIGINAL Acotel JAR resources only.
 * The Magnolia runtime never uses this class or copies emulator code/assets.
 * Reading all PNGs after selecting the original Canvas is a reference-side
 * resource oracle, not a claim that Magnolia started or painted the game.
 */
public final class AcotelPngReference {
    private static String hex(byte[] raw) {
        StringBuilder result = new StringBuilder();
        for (byte b : raw) {
            int value = b & 255;
            result.append(Character.forDigit(value >>> 4, 16));
            result.append(Character.forDigit(value & 15, 16));
        }
        return result.toString();
    }

    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "ACO_XmasGifts_s"});
            Displayable screen = null;
            for (int attempt = 0; attempt < 40; ++attempt) {
                MIDlet midlet = MIDletBridge.getCurrentMIDlet();
                if (midlet != null && "ACO_XmasGifts_s".equals(
                        midlet.getClass().getName())) {
                    screen = Display.getDisplay(midlet).getCurrent();
                    if (screen != null && "DS_Effects".equals(
                            screen.getClass().getName())) break;
                }
                Thread.sleep(50L);
            }
            if (screen == null || !"DS_Effects".equals(screen.getClass().getName()))
                throw new AssertionError("original game did not select DS_Effects");
            String[] names = new String[34];
            names[0] = "a"; names[1] = "b1";
            for (int i = 0; i < 32; ++i) names[i + 2] = "s" + i;
            int totalPixels = 0;
            for (String name : names) {
                Image image = Image.createImage("/" + name);
                int width = image.getWidth(), height = image.getHeight();
                long area = (long)width * height;
                if (width < 1 || width > 512 || height < 1 ||
                        height > 256 || area > 65536 ||
                        totalPixels + area > 100000)
                    throw new AssertionError("resource geometry outside bound: " + name);
                int[] argb = new int[(int)area];
                image.getRGB(argb, 0, width, 0, 0, width, height);
                MessageDigest digest = MessageDigest.getInstance("SHA-256");
                int transparent = 0;
                for (int pixel : argb) {
                    digest.update((byte)(pixel >>> 24));
                    digest.update((byte)(pixel >>> 16));
                    digest.update((byte)(pixel >>> 8));
                    digest.update((byte)pixel);
                    if ((pixel >>> 24) == 0) ++transparent;
                }
                totalPixels += (int)area;
                System.out.println("REFPNG:" + name + ":" + width + "x" +
                    height + ":sha=" + hex(digest.digest()) +
                    ":transparent=" + transparent);
            }
            System.out.println("REFPNG:finished:images=34:pixels=" + totalPixels);
            System.exit(0); // original game's device threads must not persist
        } catch (Throwable error) {
            System.err.println("REFPNG:FAILED " + error.toString());
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
