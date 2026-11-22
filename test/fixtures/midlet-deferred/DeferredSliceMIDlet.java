import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Same bytecode on MicroEmulator and Magnolia. The real emulator executes
// View.run(); Magnolia's opt-in start-return profile deliberately does not.
// There is no paint callback, visible device or input claim in Magnolia.
public final class DeferredSliceMIDlet extends MIDlet {
    private View view;

    public DeferredSliceMIDlet() {
        System.out.println("TRACE:constructed");
    }

    public static final class View extends Canvas implements Runnable {
        public volatile int workerState;
        public void run() {
            workerState = 7;
            System.out.println("TRACE:worker-ran");
        }
        protected void paint(Graphics graphics) { }
    }

    protected void startApp() {
        System.out.println("TRACE:start-entered");
        view = new View();
        Display.getDisplay(this).setCurrent(view);
        Thread worker = new Thread(view);
        System.out.println(worker.isAlive() ? "BAD:early-alive" : "TRACE:not-alive-before-start");
        worker.start();
        System.out.println("TRACE:start-returning");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
