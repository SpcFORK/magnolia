import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;

// Owned, unrelated MIDP bytecode. Neither original game names nor its
// key script or field layout appear here. Only the constructor executes in
// the opt-in Magnolia direct-callback image: NOT startApp or device input.
public final class DirectSliceMIDlet extends MIDlet {
    private final Display display;
    private final View view;

    public DirectSliceMIDlet() {
        display = Display.getDisplay(this);
        view = new View();
    }

    public static final class View extends Canvas {
        public int action;
        public int presses;
        public boolean down;

        public View() { super(); }
        public void keyPressed(int key) {
            action = getGameAction(key);
            presses++;
            down = true;
        }
        public void keyReleased(int key) {
            action = getGameAction(key);
            down = false;
        }
        public void keyRepeated(int key) {
            action = getGameAction(key);
            presses += 2;
            down = true;
        }
        protected void paint(Graphics graphics) {
            System.out.println("BAD:unexpected-paint");
        }
    }

    protected void startApp() {
        System.out.println("BAD:unexpected-startApp");
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean force) { }
}
