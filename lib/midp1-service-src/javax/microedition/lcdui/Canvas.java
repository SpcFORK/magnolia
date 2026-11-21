package javax.microedition.lcdui;

// Separate opt-in bounded SERVICE profile. The default Canvas still has no
// repaint method; the one-frame paint Canvas still has no serviceRepaints.
// The trusted host attaches the first selected frame ONLY AFTER the first
// showNotify/paint callback. Exactly one subsequent selected repaint may be
// synchronously serviced in the verified MIDlet triggerService() method.
// This is not a concurrent event thread, a timer or a general MIDP queue.
public abstract class Canvas extends Displayable {
    public static final int UP = 1;
    public static final int DOWN = 6;
    public static final int LEFT = 2;
    public static final int RIGHT = 5;
    public static final int FIRE = 8;
    public static final int GAME_A = 9;
    public static final int GAME_B = 10;
    public static final int GAME_C = 11;
    public static final int GAME_D = 12;

    private boolean repaintPending;
    private Image serviceFrame;
    private int serviceCount;

    protected Canvas() { super(); }
    protected abstract void paint(Graphics graphics);
    protected void showNotify() { }

    public final void repaint() { repaintPending = true; }

    // The frame is never guest-assigned. Invocations before the trusted host
    // has selected and attached its first frame, or after one service, fail.
    // A pending request is consumed before calling user paint bytecode, so a
    // nested/reentrant service cannot generate unbounded paint callbacks.
    public final void serviceRepaints() {
        if (serviceFrame == null || serviceCount != 0)
            throw new IllegalStateException();
        if (repaintPending) {
            repaintPending = false;
            serviceCount = 1;
            paint(ImageBridge.graphics(serviceFrame));
        }
    }

    public int getGameAction(int keyCode) {
        if (keyCode == -1) return UP;
        if (keyCode == -2) return DOWN;
        if (keyCode == -3) return LEFT;
        if (keyCode == -4) return RIGHT;
        if (keyCode == -5) return FIRE;
        if (keyCode == 49) return GAME_A;
        if (keyCode == 51) return GAME_B;
        if (keyCode == 55) return GAME_C;
        if (keyCode == 57) return GAME_D;
        if (keyCode == 0 || keyCode == 8 || keyCode == 48 ||
                keyCode == 50 || keyCode == 52 || keyCode == 53 ||
                keyCode == 54 || keyCode == 56 || keyCode == 42 ||
                keyCode == 35 || keyCode == 999) return 0;
        throw new IllegalArgumentException();
    }
}
