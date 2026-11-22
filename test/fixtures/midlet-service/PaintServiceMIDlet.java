import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Owned, separately checked two-frame fixture. Original Acotel bytes are NOT
// edited or embedded. One selected repaint request plus serviceRepaints
// should produce one further paint with a different full-frame hash.
public final class PaintServiceMIDlet extends MIDlet {
    public static volatile View view;
    public static volatile int detachedPaints;
    public static final class View extends Canvas {
        public volatile int shown;
        public volatile int paints;
        public volatile int clipWidth;
        protected void showNotify() {
            shown++;
            System.out.println("TRACE:shown");
        }
        protected void paint(Graphics graphics) {
            clipWidth = graphics.getClipWidth();
            graphics.setColor(0);
            graphics.fillRect(0, 0, 176, 176);
            if (paints == 0) {
                graphics.setColor(0xFF0000);
                graphics.fillRect(10, 12, 16, 9);
                System.out.println("TRACE:paint-one");
            } else {
                graphics.setColor(0x00FF00);
                graphics.fillRect(32, 41, 13, 7);
                System.out.println("TRACE:paint-two");
            }
            paints++;
        }
    }
    protected void startApp() {
        view = new View();
        System.out.println("TRACE:start");
        view.repaint();
        detachedPaints = view.paints;
        System.out.println("TRACE:detached");
        Display.getDisplay(this).setCurrent(view);
        System.out.println("TRACE:selected");
    }
    public void triggerService() {
        view.repaint();
        System.out.println("TRACE:request-returned");
        view.serviceRepaints();
        System.out.println("TRACE:service-returned");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
