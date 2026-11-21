package javax.microedition.lcdui;

import javax.microedition.midlet.MIDlet;

// Single-MIDlet, headless *requested* display state. There is no device
// display, event queue, paint callback, getCurrent contract or rasterization.
public final class Display {
    private static Display instance;
    private static MIDlet owner;
    private Displayable requested;

    private Display() { }

    public static Display getDisplay(MIDlet midlet) {
        if (midlet == null) throw new NullPointerException();
        if (instance == null) {
            instance = new Display();
            owner = midlet;
        }
        if (owner != midlet) throw new IllegalArgumentException();
        return instance;
    }

    public void setCurrent(Displayable screen) { requested = screen; }
}
