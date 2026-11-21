package javax.microedition.lcdui;

// Opt-in, pinned-default-device Canvas model. The normal MIDlet profile still
// uses the separate no-paint Canvas classfile. A repaint only records a pending
// request: the trusted finite host dispatches one showNotify + paint on the
// selected Canvas after startApp returns. No timer, serviceRepaints, keypad,
// second frame or generic MIDP event queue is claimed.
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

    protected Canvas() { super(); }
    protected abstract void paint(Graphics graphics);
    protected void showNotify() { }

    // A detached request never invokes paint. Selecting a Canvas schedules
    // its initial paint even if no repaint was requested; the finite host
    // consumes at most one pending request at the post-start dispatch point.
    public void repaint() { repaintPending = true; }

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
