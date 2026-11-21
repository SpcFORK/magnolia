package javax.microedition.lcdui;

// One pinned default-device key-to-game-action map. This executes as real
// Java bytecode in Magnolia's finite MIDlet host; it is not an event queue.
// In this DEFAULT model other Canvas/device APIs (including repaint and
// getKeyCode) still refuse; a separate opt-in one-frame model is kept under
// lib/midp1-paint-src. The original Acotel game uses this exact
// method before dispatching its numeric-key fallbacks in keyPressed/keyReleased.
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

    protected Canvas() { super(); }
    protected abstract void paint(Graphics graphics);

    // Pinned MicroEmulator 2.0.4 default J2SE device: physical arrows/fire
    // are -1..-5. The four extra game keys use odd numbered numeric codes,
    // not 2/4/5/6/8. Zero is checked for the other listed input codes only;
    // a new/unoracled key refuses rather than pretending all devices agree.
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
