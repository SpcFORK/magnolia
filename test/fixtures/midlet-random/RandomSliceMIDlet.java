import java.util.Random;
import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Compile ONCE against the independent MicroEmulator classpath. Execute the
// identical class bytes in real MicroEmulator and the narrow Magnolia host.
// All samples have explicit long seeds: NO default-device clock is oracled.
public final class RandomSliceMIDlet extends MIDlet {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }

    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:random-begin");
        Random zero = new Random(0L);
        System.out.println(zero.nextInt());
        System.out.println(zero.nextInt());
        System.out.println(zero.nextInt(1));
        System.out.println(zero.nextInt(2));
        System.out.println(zero.nextInt(1024));
        System.out.println(zero.nextInt(2147483647));
        try {
            zero.nextInt(0);
            System.out.println("BAD:zero-bound");
        } catch (IllegalArgumentException expected) {
            System.out.println("TRACE:zero-bound-caught");
        }
        try {
            zero.nextInt(-1);
            System.out.println("BAD:negative-bound");
        } catch (IllegalArgumentException expected) {
            System.out.println("TRACE:negative-bound-caught");
        }
        System.out.println(zero.nextInt());
        Random negative = new Random(-1L);
        System.out.println(negative.nextInt());
        System.out.println(negative.nextInt(1073741825));
        System.out.println(negative.nextInt(1073741825));
        System.out.println(negative.nextInt(1073741824));
        Random maximum = new Random(0x7fffffffffffffffL);
        System.out.println(maximum.nextInt());
        maximum.setSeed(-0x8000000000000000L);
        System.out.println(maximum.nextInt(2147483647));
        maximum.setSeed(0L);
        System.out.println(maximum.nextInt());
        // Allocation with Random() is allowed, but its unpinned clock may
        // never be sampled; setSeed first, then compare with explicit seed.
        Random reset = new Random();
        reset.setSeed(0L);
        System.out.println(reset.nextInt());
        System.out.println("TRACE:random-end");
        Display.getDisplay(this).setCurrent(new View());
    }

    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
