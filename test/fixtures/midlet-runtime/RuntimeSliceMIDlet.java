import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// One actual MicroEmulator/JDK freeMemory observation is replayed as an
// explicit Magnolia snapshot. The second, low snapshot is branch-logic only:
// no genuine device was forced to have exactly 4096 free bytes.
public final class RuntimeSliceMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }

    public static boolean low(long free) { return free < 5000L; }

    protected void startApp() throws MIDletStateChangeException {
        Runtime runtime = Runtime.getRuntime();
        long snapshot = runtime.freeMemory();
        System.out.println("TRACE:runtime-begin");
        System.out.println(runtime == Runtime.getRuntime());
        System.out.println(snapshot);
        if (low(snapshot)) {
            System.out.println("TRACE:below-five-thousand");
            runtime.gc();
        } else {
            System.out.println("TRACE:at-least-five-thousand");
        }
        runtime.gc();
        System.out.println("TRACE:runtime-end");
        Display.getDisplay(this).setCurrent(new View());
    }

    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
