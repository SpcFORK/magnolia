import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// This finite fixture waits until the *real* device worker has exited before
// comparing lines. Magnolia's bounded task executes synchronously; timing,
// interleavings, a live-game loop and repaint are explicitly NOT compared.
public final class ThreadSliceMIDlet extends MIDlet implements Runnable {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }
    private volatile boolean done;

    public void run() {
        System.out.println("TRACE:task-begin");
        try {
            Thread.sleep(7L);
        } catch (InterruptedException unexpected) {
            System.out.println("BAD:interrupted");
        }
        System.out.println("TRACE:task-end");
        done = true;
    }

    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:thread-begin");
        Thread worker = new Thread(this);
        System.out.println(worker.isAlive() ? 1 : 0);
        worker.start();
        int polls = 0;
        while (worker.isAlive() && polls < 10000000) polls++;
        System.out.println(worker.isAlive() ? "BAD:worker-not-finished" : "TRACE:worker-finished");
        System.out.println(done ? 1 : 0);
        System.out.println("TRACE:thread-end");
        Display.getDisplay(this).setCurrent(new View());
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
