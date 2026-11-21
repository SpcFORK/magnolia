import java.lang.reflect.Field;
import java.lang.reflect.Method;
import java.util.Vector;
import javax.microedition.lcdui.Command;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// Observe the real registered command/listener on a genuine MicroEmulator
// Displayable, with bounded AMS wait. Reflection is reference OBSERVATION;
// no fixture bytecode is changed or substituted in the production VM.
public final class CommandSliceReference {
    public static void main(String[] args) {
        try {
            if (args.length != 0) throw new IllegalArgumentException("no URLs");
            Headless.main(new String[]{"--rms", "memory", "CommandSliceMIDlet"});
            for (int attempt = 0; attempt < 40; attempt++) {
                MIDlet current = MIDletBridge.getCurrentMIDlet();
                if (current != null && "CommandSliceMIDlet".equals(
                        current.getClass().getName())) {
                    Displayable view = Display.getDisplay(current).getCurrent();
                    if (view != null && "CommandSliceMIDlet$View".equals(
                            view.getClass().getName())) {
                        Method method = Displayable.class.getDeclaredMethod("getCommands");
                        method.setAccessible(true);
                        Vector commands = (Vector) method.invoke(view);
                        Field field = Displayable.class.getDeclaredField("listener");
                        field.setAccessible(true);
                        Object listener = field.get(view);
                        if (commands == null || commands.size() != 1 ||
                                !"Fine".equals(((Command) commands.get(0)).getLabel()) ||
                                listener != current) {
                            throw new AssertionError("real command/listener registration differs");
                        }
                        System.out.println("REFERENCE:command-count=1;label=Fine;listener=entry");
                        System.exit(0);
                    }
                }
                Thread.sleep(50L);
            }
            throw new AssertionError("real command fixture never selected Canvas");
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
