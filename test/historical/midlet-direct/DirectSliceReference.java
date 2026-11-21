import java.lang.reflect.Field;
import java.lang.reflect.Method;
import javax.microedition.lcdui.Display;
import javax.microedition.midlet.MIDlet;
import org.microemu.app.Headless;

/**
 * Real, separately pinned MicroEmulator observation of the same TWO owned
 * classfiles that Magnolia interprets. This observer makes direct calls on
 * the still-UNSELECTED Canvas, not hardware/device-delivered key events.
 */
public final class DirectSliceReference extends MIDlet {
    protected void startApp() {
        try {
            DirectSliceMIDlet game = new DirectSliceMIDlet();
            Field viewField = DirectSliceMIDlet.class.getDeclaredField("view");
            Field displayField = DirectSliceMIDlet.class.getDeclaredField("display");
            viewField.setAccessible(true);
            displayField.setAccessible(true);
            DirectSliceMIDlet.View view = (DirectSliceMIDlet.View)viewField.get(game);
            Display display = (Display)displayField.get(game);
            if (display != Display.getDisplay(game) ||
                    display.getCurrent() != null ||
                    view.action != 0 || view.presses != 0 || view.down)
                throw new AssertionError("constructed Canvas already selected/modified");
            int[] keys = {-5, -5, 49};
            String[] methods = {"keyPressed", "keyReleased", "keyRepeated"};
            StringBuilder trace = new StringBuilder();
            for (int index = 0; index < keys.length; ++index) {
                Method method = DirectSliceMIDlet.View.class.getMethod(
                        methods[index], int.class);
                method.invoke(view, keys[index]);
                if (index != 0) trace.append(',');
                trace.append(methods[index]).append('(').append(keys[index])
                        .append(")=").append(view.action).append(':')
                        .append(view.down).append(':').append(view.presses);
                if (display.getCurrent() != null)
                    throw new AssertionError("direct key selected a device Canvas");
            }
            System.out.println("REFERENCE:game=" + game.getClass().getName()
                    + ";view=" + view.getClass().getName()
                    + ";before=0:false:0;trace=" + trace.toString()
                    + ";selected=false;deviceInput=false;startApp=false;paint=false");
            System.out.flush();
            System.exit(0);
        } catch (Throwable failure) {
            System.err.println("REFERENCE:failed " + failure);
            failure.printStackTrace();
            System.exit(2);
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }

    public static void main(String[] args) throws Exception {
        if (args.length != 0) throw new IllegalArgumentException("no URLs/arguments");
        Headless.main(new String[]{"--rms", "memory", "DirectSliceReference"});
        Thread.sleep(3500L);
        System.err.println("REFERENCE:proxy-not-called");
        System.exit(2);
    }
}
