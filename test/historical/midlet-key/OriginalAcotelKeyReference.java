import java.lang.reflect.Field;
import java.lang.reflect.Method;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Optional, separate *original game* MicroEmulator evidence. Does not run the
// game in Magnolia or deliver a physical/AMS keypad event.
public final class OriginalAcotelKeyReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "ACO_XmasGifts_s"});
            MIDlet game = MIDletBridge.getCurrentMIDlet();
            if (game == null || !"ACO_XmasGifts_s".equals(game.getClass().getName()))
                throw new AssertionError("original game not constructed");
            Displayable displayable = null;
            for (int attempt = 0; attempt < 40; ++attempt) {
                displayable = Display.getDisplay(game).getCurrent();
                if (displayable != null) break;
                Thread.sleep(50L);
            }
            if (!(displayable instanceof Canvas) ||
                    !"DS_Effects".equals(displayable.getClass().getName()))
                throw new AssertionError("original Canvas not current");
            Canvas canvas = (Canvas)displayable;
            int[] keys = {-5, -4, -3, -2, -1, 0, 8, 48, 49, 50,
                          51, 52, 53, 54, 55, 56, 57, 42, 35, 999};
            System.out.println("ORIGINAL:keys-begin");
            for (int key : keys) System.out.println(canvas.getGameAction(key));
            System.out.println("ORIGINAL:keys-end");
            // The original callback is invoked on its live object, rather
            // than inferred from static disassembly. The running game may
            // consume its flags asynchronously, so the snapshots are notes,
            // NOT timing-stable assertions or emulator event delivery.
            Method press = displayable.getClass().getMethod("keyPressed", int.class);
            Method release = displayable.getClass().getMethod("keyReleased", int.class);
            Field action = displayable.getClass().getDeclaredField("_fld2h");
            Field down = displayable.getClass().getDeclaredField("_fld3h");
            action.setAccessible(true);
            down.setAccessible(true);
            press.invoke(displayable, 50);
            System.out.println("ORIGINAL:pressed50=" + action.getInt(displayable) +
                               ",down=" + down.getBoolean(displayable));
            release.invoke(displayable, 50);
            System.out.println("ORIGINAL:released50=" + action.getInt(displayable) +
                               ",down=" + down.getBoolean(displayable));
            System.out.println("ORIGINAL:reference-finished");
            System.exit(0);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
