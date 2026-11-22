import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Preflight must inspect the reachable host-invoked paint callback BEFORE
// this constructor/startApp's BAD line can reach the CLI. drawLine remains
// unmodeled on the opt-in one-frame MIDP Graphics, despite AWT having one.
public final class PaintSliceUnsupportedMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics g) { g.drawLine(0, 0, 1, 1); }
    }
    protected void startApp() {
        System.out.println("BAD:started-before-unsupported-paint");
        Display.getDisplay(this).setCurrent(new View());
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
