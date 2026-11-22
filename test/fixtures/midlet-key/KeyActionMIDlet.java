import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// The original Acotel DS_Effects.keyPressed/keyReleased calls precisely this
// Canvas.getGameAction member. Test the model as executed Java bytecode, not a
// guessed label or a static descriptor. The reference runs on MicroEmulator.
public final class KeyActionMIDlet extends MIDlet {
    private final Display display;
    private final View view;
    public KeyActionMIDlet() {
        display = Display.getDisplay(this);
        view = new View();
    }
    public static final class View extends Canvas {
        public View() { super(); }
        protected void paint(Graphics graphics) { }
    }
    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:key-begin");
        System.out.println(view.getGameAction(-5));
        System.out.println(view.getGameAction(-4));
        System.out.println(view.getGameAction(-3));
        System.out.println(view.getGameAction(-2));
        System.out.println(view.getGameAction(-1));
        System.out.println(view.getGameAction(0));
        System.out.println(view.getGameAction(8));
        System.out.println(view.getGameAction(48));
        System.out.println(view.getGameAction(49));
        System.out.println(view.getGameAction(50));
        System.out.println(view.getGameAction(51));
        System.out.println(view.getGameAction(52));
        System.out.println(view.getGameAction(53));
        System.out.println(view.getGameAction(54));
        System.out.println(view.getGameAction(55));
        System.out.println(view.getGameAction(56));
        System.out.println(view.getGameAction(57));
        System.out.println(view.getGameAction(42));
        System.out.println(view.getGameAction(35));
        System.out.println(view.getGameAction(999));
        System.out.println("TRACE:key-end");
        display.setCurrent(view);
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
            throws MIDletStateChangeException { notifyDestroyed(); }
}
