import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.security.MessageDigest;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import org.microemu.app.Headless;

/**
 * Separately pinned real-MicroEmulator observer, not a replacement game.
 * Its OWN startApp constructs the unchanged external original Acotel MIDlet,
 * then invokes exactly three original DS_Effects key callbacks by reflection
 * on the unselected Canvas. This is NOT a device/AMS keypad event. It never
 * invokes the original game's startApp, paint, run, or commandAction.
 */
public final class AcotelDirectKeyReference extends MIDlet {
    private static Field field(Class<?> owner, String name) throws Exception {
        Field f = owner.getDeclaredField(name);
        f.setAccessible(true);
        return f;
    }
    private static String pixelHash(int[] pixels) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        for (int pixel : pixels) {
            digest.update((byte)(pixel >>> 24));
            digest.update((byte)(pixel >>> 16));
            digest.update((byte)(pixel >>> 8));
            digest.update((byte)pixel);
        }
        StringBuilder text = new StringBuilder();
        for (byte item : digest.digest()) {
            text.append(Character.forDigit((item & 255) >>> 4, 16));
            text.append(Character.forDigit(item & 15, 16));
        }
        return text.toString();
    }
    protected void startApp() {
        try {
            Class<?> original = Class.forName("ACO_XmasGifts_s");
            Object game = original.getConstructor().newInstance();
            Object canvas = field(original, "_fld2").get(game);
            Class<?> view = canvas.getClass();
            Display display = (Display)field(original, "_fld1").get(game);
            if (!"DS_Effects".equals(view.getName()) ||
                display != Display.getDisplay((MIDlet)game) ||
                display.getCurrent() != null ||
                field(view, "_fld3b").get(canvas) != null)
                throw new AssertionError("original game not constructed and unselected");
            Field action = field(view, "_fld2h");
            Field down = field(view, "_fld3h");
            int initial = action.getInt(canvas);
            boolean initiallyDown = down.getBoolean(canvas);
            String[] methods = {"keyPressed", "keyReleased", "keyRepeated"};
            int[] keys = {50, 50, -4};
            StringBuilder trace = new StringBuilder();
            for (int i = 0; i < methods.length; ++i) {
                Method callback = view.getMethod(methods[i], int.class);
                callback.invoke(canvas, keys[i]);
                if (i != 0) trace.append(',');
                trace.append(methods[i]).append('(').append(keys[i]).append(")=")
                     .append(action.getInt(canvas)).append(':')
                     .append(down.getBoolean(canvas));
                if (display.getCurrent() != null ||
                    field(view, "_fld3b").get(canvas) != null)
                    throw new AssertionError("direct callback selected display or started worker");
            }
            Image image = (Image)field(view, "_fld9").get(canvas);
            if (image == null || image.getWidth() != 176 || image.getHeight() != 176)
                throw new AssertionError("original constructor image extent changed");
            int[] pixels = new int[176 * 176];
            image.getRGB(pixels, 0, 176, 0, 0, 176, 176);
            for (int pixel : pixels) {
                if (pixel != 0xFF000000)
                    throw new AssertionError("unstarted game unexpectedly drew pixels");
            }
            System.out.println("DIRECT:constructor=" + original.getName() +
                ";canvas=" + view.getName() + ";before=" + initial + ":" +
                initiallyDown + ";trace=" + trace + ";offscreen=176x176" +
                ";argbSHA256=" + pixelHash(pixels) +
                ";displaySelected=false;workerStarted=false;" +
                "gameStartAppCalled=false;gamePaintCalled=false;" +
                "deviceInputDelivered=false");
            System.out.flush();
            System.exit(0);
        } catch (Throwable failure) {
            System.err.println("DIRECT:failed " + failure);
            failure.printStackTrace();
            System.exit(2);
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
    public static void main(String[] args) throws Exception {
        if (args.length != 0) throw new IllegalArgumentException("no arguments or URLs");
        Headless.main(new String[]{"--rms", "memory", "AcotelDirectKeyReference"});
        Thread.sleep(3500L);
        System.err.println("DIRECT:proxy-not-called");
        System.exit(2);
    }
}
