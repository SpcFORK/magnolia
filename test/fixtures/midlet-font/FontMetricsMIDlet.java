import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Font;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Compile once against real MicroEmulator. Execute these identical bytes in
// that emulator and Magnolia. The original Acotel requests exactly these
// three Font variants and uses their stringWidth/getHeight in game UI.
public final class FontMetricsMIDlet extends MIDlet {
    private final Display display = Display.getDisplay(this);
    private final View view = new View();
    public static final class View extends Canvas {
        public View() { super(); }
        protected void paint(Graphics graphics) { }
    }
    private void metrics(int size) {
        Font font = Font.getFont(0, 1, size);
        System.out.println(size);
        System.out.println(font == Font.getFont(0, 1, size) ? 1 : 0);
        System.out.println(font.getHeight());
        System.out.println(font.charWidth('W'));
        System.out.println(font.charWidth('i'));
        System.out.println(font.charWidth(' '));
        System.out.println(font.stringWidth(""));
        System.out.println(font.stringWidth("xxChristmas"));
        System.out.println(font.stringWidth("? www.dseffects.com"));
        System.out.println(font.stringWidth("Tu hai completato il gioco"));
        System.out.println(font.stringWidth("Xmas Gifts"));
    }
    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:font-begin");
        metrics(8);
        metrics(0);
        metrics(16);
        System.out.println("TRACE:charat-valid");
        System.out.println((int)"abc".charAt(1));
        try {
            "abc".charAt(-1);
            System.out.println("BAD:charat-negative");
        } catch (StringIndexOutOfBoundsException expected) {
            System.out.println("TRACE:charat-negative-caught");
        }
        try {
            "abc".charAt(3);
            System.out.println("BAD:charat-length");
        } catch (StringIndexOutOfBoundsException expected) {
            System.out.println("TRACE:charat-length-caught");
        }
        System.out.println("TRACE:font-end");
        display.setCurrent(view);
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
