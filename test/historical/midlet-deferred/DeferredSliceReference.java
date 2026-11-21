import java.lang.reflect.Field;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Independent MicroEmulator observer for the exact same owned classfiles
// used by the native generic start-return profile. The real worker must run;
// no actual device pixel, scheduler ordering or keypad parity is asserted.
public final class DeferredSliceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs or arguments");
            Headless.main(new String[]{"--rms", "memory", "DeferredSliceMIDlet"});
            for (int tick = 0; tick < 50; tick++) {
                MIDlet game = MIDletBridge.getCurrentMIDlet();
                if (game != null && "DeferredSliceMIDlet".equals(game.getClass().getName())) {
                    Field viewField = game.getClass().getDeclaredField("view");
                    viewField.setAccessible(true);
                    Object view = viewField.get(game);
                    if (view != null && "DeferredSliceMIDlet$View".equals(
                            view.getClass().getName())) {
                        Field state = view.getClass().getDeclaredField("workerState");
                        state.setAccessible(true);
                        Displayable selected = Display.getDisplay(game).getCurrent();
                        if (state.getInt(view) == 7 && selected == view) {
                            System.out.println("REFERENCE:midlet=" + game.getClass().getName()
                                    + ";view=" + view.getClass().getName()
                                    + ";workerState=" + state.getInt(view)
                                    + ";selected=true");
                            System.exit(0);
                        }
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real worker or Canvas selection never completed");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
