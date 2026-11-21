import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Owned fixture against REAL MicroEmulator 2.0.4. Not an Acotel class and
// not the stubbed Magnolia LCDUI implementation. No physical device or game.
public final class RepaintProbeMIDlet extends MIDlet {
    public static volatile View view;
    public static volatile int detachedPaints;
    public static volatile int selectionRequestedPaints;
    public static final class View extends Canvas {
        public volatile int paints;
        public volatile int shown;
        public volatile int clipWidth;
        public volatile String paintThread;
        protected void showNotify() {
            System.out.println("TRACE:shown");
            shown++;
        }
        protected void paint(Graphics g) {
            int next = paints + 1;
            System.out.println("TRACE:paint");
            clipWidth = g.getClipWidth();
            paintThread = Thread.currentThread().getName();
            // Publish the counter after the entire trace line. A reference
            // reader observing the volatile count has seen the print too.
            paints = next;
        }
    }
    protected void startApp() {
        view = new View();
        System.out.println("TRACE:start");
        view.repaint();
        detachedPaints = view.paints;
        System.out.println("TRACE:detached-repaint");
        Display.getDisplay(this).setCurrent(view);
        selectionRequestedPaints = view.paints;
        System.out.println("TRACE:selection-requested");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
