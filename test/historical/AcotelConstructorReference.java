import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.security.MessageDigest;
import java.util.HashSet;
import java.util.Vector;
import javax.microedition.lcdui.Command;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import org.microemu.app.Headless;

/**
 * Own proxy MIDlet used ONLY to initialize a real MicroEmulator 2.0.4 device.
 * Its startApp() constructs the untouched, external original Acotel MIDlet by
 * reflection; it NEVER invokes the original game's startApp(), run(), paint(),
 * or a key/softkey callback. No original JAR/class/resource is copied here.
 * The owning process is bounded externally and explicitly exits after the
 * constructor returns so emulator timers cannot keep this reference alive.
 */
public final class AcotelConstructorReference extends MIDlet {
    private static Field field(Class<?> owner, String name) throws Exception {
        Field result = owner.getDeclaredField(name);
        result.setAccessible(true);
        return result;
    }
    private static String argbSha256(int[] pixels) throws Exception {
        MessageDigest digest = MessageDigest.getInstance("SHA-256");
        for (int pixel : pixels) {
            digest.update((byte)(pixel >>> 24));
            digest.update((byte)(pixel >>> 16));
            digest.update((byte)(pixel >>> 8));
            digest.update((byte)pixel);
        }
        StringBuilder result = new StringBuilder();
        for (byte item : digest.digest()) {
            int value = item & 255;
            result.append(Character.forDigit(value >>> 4, 16));
            result.append(Character.forDigit(value & 15, 16));
        }
        return result.toString();
    }
    protected void startApp() {
        try {
            // The proxy's startApp(), NOT original Acotel's. The unchanged
            // external original class is loaded by the real reference JVM.
            Class<?> gameClass = Class.forName("ACO_XmasGifts_s");
            Object game = gameClass.getConstructor().newInstance();
            Object canvas = field(gameClass, "_fld2").get(game);
            Display display = (Display)field(gameClass, "_fld1").get(game);
            Command command = (Command)field(gameClass, "_fld3").get(game);
            if (canvas == null || !"DS_Effects".equals(canvas.getClass().getName()) ||
                display == null || display != Display.getDisplay((MIDlet)game) ||
                display.getCurrent() != null ||
                field(canvas.getClass(), "_fld3b").get(canvas) != null ||
                field(canvas.getClass(), "_fld2b").getBoolean(canvas) ||
                field(canvas.getClass(), "_fld3c").getInt(canvas) != 0 ||
                field(canvas.getClass(), "_fld1c").getInt(canvas) != 20)
                throw new AssertionError("original constructor/Canvas/device state changed");

            Class<?> displayable = canvas.getClass().getSuperclass().getSuperclass();
            Method commandsMethod = displayable.getDeclaredMethod("getCommands");
            commandsMethod.setAccessible(true);
            Vector<?> commands = (Vector<?>)commandsMethod.invoke(canvas);
            if (commands.size() != 1 || commands.get(0) != command ||
                !"Fine".equals(command.getLabel()) ||
                command.getCommandType() != 1 || command.getPriority() != 0 ||
                field(displayable, "listener").get(canvas) != game)
                throw new AssertionError("original command registration changed");

            Image image = (Image)field(canvas.getClass(), "_fld9").get(canvas);
            if (image == null || image.getWidth() != 176 || image.getHeight() != 176)
                throw new AssertionError("original offscreen image missing/changed");
            int[] pixels = new int[176 * 176];
            image.getRGB(pixels, 0, 176, 0, 0, 176, 176);
            HashSet<Integer> colors = new HashSet<Integer>();
            for (int pixel : pixels) colors.add(pixel);
            if (colors.size() != 1 || !colors.contains(0xFF000000))
                throw new AssertionError("constructor image unexpectedly drawn");
            System.out.println("REFERENCE:constructor=ACO_XmasGifts_s;canvas=DS_Effects;" +
                "offscreen=176x176;colors=" + colors.size() + ";argbSHA256=" +
                argbSha256(pixels) + ";registered=Fine:1:0;listener=same-object;" +
                "selected=false;worker=false;state=0;sleep=20;stop=false;" +
                "gameStartAppCalled=false;gamePaintCalled=false");
            System.out.flush();
            System.exit(0);
        } catch (Throwable failure) {
            System.err.println("REFERENCE:constructor-failed " + failure);
            failure.printStackTrace();
            System.exit(2);
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
    public static void main(String[] args) throws Exception {
        if (args.length != 0) throw new IllegalArgumentException("no arguments or URLs");
        Headless.main(new String[]{"--rms", "memory", "AcotelConstructorReference"});
        // MicroEmulator loads its own classloader copy of this proxy MIDlet;
        // do NOT wait on a static field in the main loader. Its startApp calls
        // System.exit itself after flushing the observed constructor record.
        Thread.sleep(3500L);
        System.err.println("REFERENCE:constructor-proxy-never-called");
        System.exit(2);
    }
}
