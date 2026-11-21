import java.lang.reflect.Field;
import javax.microedition.lcdui.Display;
import javax.microedition.midlet.MIDlet;
import org.microemu.app.Headless;

/**
 * Independent constructor-only observation of the exact same two owned
 * DirectSliceMIDlet classfiles used by the native generic constructor host.
 * No startApp, direct callback, device key, worker or paint is invoked.
 */
public final class DirectConstructorReference extends MIDlet {
    protected void startApp() {
        try {
            DirectSliceMIDlet target = new DirectSliceMIDlet();
            Field viewField = DirectSliceMIDlet.class.getDeclaredField("view");
            Field displayField = DirectSliceMIDlet.class.getDeclaredField("display");
            viewField.setAccessible(true);
            displayField.setAccessible(true);
            DirectSliceMIDlet.View view = (DirectSliceMIDlet.View)viewField.get(target);
            Display display = (Display)displayField.get(target);
            if (display != Display.getDisplay(target) ||
                    display.getCurrent() != null ||
                    view.action != 0 || view.presses != 0 || view.down)
                throw new AssertionError("constructor changed unselected Canvas state");
            System.out.println("CONSTRUCTOR:game=" + target.getClass().getName()
                    + ";view=" + view.getClass().getName()
                    + ";displaySame=true;selected=false;state=0:false:0"
                    + ";targetStartApp=false;callbacks=false;paint=false;deviceInput=false");
            System.out.flush();
            System.exit(0);
        } catch (Throwable failure) {
            System.err.println("CONSTRUCTOR:failed " + failure);
            failure.printStackTrace();
            System.exit(2);
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }

    public static void main(String[] args) throws Exception {
        if (args.length != 0) throw new IllegalArgumentException("no URLs/arguments");
        Headless.main(new String[]{"--rms", "memory", "DirectConstructorReference"});
        Thread.sleep(3500L);
        System.err.println("CONSTRUCTOR:proxy-not-called");
        System.exit(2);
    }
}
