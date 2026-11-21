import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.HashSet;
import java.util.Vector;
import javax.microedition.lcdui.Command;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Image;
import javax.microedition.midlet.MIDlet;
import org.microemu.app.Headless;

/**
 * Independent, SHA-gated real-MicroEmulator observation of the UNCHANGED
 * game's actual startApp() method. This proxy's startApp constructs the
 * original game, invokes its startApp by reflection, then observes that its
 * REAL worker runs and draws. Magnolia's separate checkpoint deliberately
 * stops before Runnable.run, so these are not a frame/playback differential.
 * Never load a JAD URL, substitute original bytecode, or deliver an input.
 */
public final class AcotelScheduledStartReference extends MIDlet {
    private static Field field(Class<?> owner, String name) throws Exception {
        Field result = owner.getDeclaredField(name);
        result.setAccessible(true);
        return result;
    }
    private static int[] pixels(Image image) {
        if (image == null || image.getWidth() != 176 || image.getHeight() != 176)
            throw new AssertionError("original image size changed");
        int[] result = new int[176 * 176];
        image.getRGB(result, 0, 176, 0, 0, 176, 176);
        return result;
    }
    protected void startApp() {
        try {
            Class<?> original = Class.forName("ACO_XmasGifts_s");
            Object game = original.getConstructor().newInstance();
            Object canvas = field(original, "_fld2").get(game);
            Class<?> view = canvas.getClass();
            Display display = (Display)field(original, "_fld1").get(game);
            Command command = (Command)field(original, "_fld3").get(game);
            if (!"DS_Effects".equals(view.getName()) ||
                display != Display.getDisplay((MIDlet)game) ||
                display.getCurrent() != null ||
                field(view, "_fld3b").get(canvas) != null ||
                field(view, "_fld2b").getBoolean(canvas) ||
                field(view, "_fld3c").getInt(canvas) != 0 ||
                field(view, "_fld1c").getInt(canvas) != 20)
                throw new AssertionError("original pre-start state changed");
            Class<?> displayable = view.getSuperclass().getSuperclass();
            Method commandsMethod = displayable.getDeclaredMethod("getCommands");
            commandsMethod.setAccessible(true);
            Vector<?> commands = (Vector<?>)commandsMethod.invoke(canvas);
            if (commands.size() != 1 || commands.get(0) != command ||
                !"Fine".equals(command.getLabel()) ||
                command.getCommandType() != 1 || command.getPriority() != 0 ||
                field(displayable, "listener").get(canvas) != game)
                throw new AssertionError("original command/listener identity changed");
            Image image = (Image)field(view, "_fld9").get(canvas);
            int[] initial = pixels(image);
            for (int pixel : initial)
                if (pixel != 0xFF000000)
                    throw new AssertionError("original pre-start offscreen not black");
            // The unchanged original game's method, NOT a proxy replacement.
            original.getMethod("startApp").invoke(game);
            Object worker = field(view, "_fld3b").get(canvas);
            if (!(worker instanceof Thread) || !((Thread)worker).isAlive() ||
                field(view, "_fld2b").getBoolean(canvas))
                throw new AssertionError("original startApp did not start its worker");
            // MicroEmulator's display selection is asynchronous: reading
            // getCurrent immediately after return is a race. Wait for the
            // selected Canvas with a finite deadline, not an assumed tick.
            boolean selected = display.getCurrent() == canvas;
            for (int i = 0; i < 25 && !selected; ++i) {
                Thread.sleep(20L);
                selected = display.getCurrent() == canvas;
            }
            if (!selected || !((Thread)worker).isAlive() ||
                field(view, "_fld2b").getBoolean(canvas))
                throw new AssertionError("original worker/display did not stay active");
            Thread.sleep(250L);
            HashSet<Integer> colors = new HashSet<Integer>();
            for (int pixel : pixels(image)) colors.add(pixel);
            if (colors.size() <= 1 || !((Thread)worker).isAlive())
                throw new AssertionError("original worker failed to draw offscreen");
            System.out.println("START:original=ACO_XmasGifts_s;canvas=DS_Effects;" +
                "preBlack=true;preSelected=false;preWorker=false;" +
                "registered=Fine:1:0;listener=same-object;" +
                "originalStartAppReturned=true;workerStarted=true;" +
                "eventuallySelected=true;workerDrewColors=" + colors.size() +
                ";deviceInputDelivered=false;noPixelParityWithMagnolia=true");
            System.out.flush();
            System.exit(0);
        } catch (Throwable failure) {
            System.err.println("START:failed " + failure);
            failure.printStackTrace();
            System.exit(2);
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
    public static void main(String[] args) throws Exception {
        if (args.length != 0) throw new IllegalArgumentException("no JAD, URL or arguments");
        Headless.main(new String[]{"--rms", "memory", "AcotelScheduledStartReference"});
        Thread.sleep(4000L);
        System.err.println("START:proxy-not-called");
        System.exit(2);
    }
}
