import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Owned same-bytecode exercise for one 176x176 Canvas selection/paint.
// MicroEmulator is the independent reference; this is not an Acotel class.
public final class PaintSliceMIDlet extends MIDlet {
    public static volatile View view;
    public static volatile int detachedPaints;
    public static volatile int selectionRequestedPaints;
    public static final class View extends Canvas {
        public volatile int shown;
        public volatile int paints;
        public volatile int clipWidth;

        protected void showNotify() {
            shown++;
            System.out.println("TRACE:shown");
        }
        protected void paint(Graphics g) {
            clipWidth = g.getClipWidth();
            g.setColor(0);
            g.fillRect(0, 0, 176, 176);
            g.setColor(0xFF0000);
            g.fillRect(10, 12, 16, 9);
            paints++;
            System.out.println("TRACE:paint");
        }
    }
    protected void startApp() {
        view = new View();
        System.out.println("TRACE:start");
        view.repaint(); // detached Canvas must NOT paint immediately
        detachedPaints = view.paints;
        System.out.println("TRACE:detached");
        Display.getDisplay(this).setCurrent(view);
        selectionRequestedPaints = view.paints;
        System.out.println("TRACE:selected");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
