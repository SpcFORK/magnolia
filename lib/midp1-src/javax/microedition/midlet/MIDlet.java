package javax.microedition.midlet;

// Owned, intentionally incomplete MIDP-1 model BYTECODE for Magnolia's
// separate headless MIDlet probe. NEVER a Java SE applet superclass or a
// drop-in MIDP implementation. The real MicroEmulator API is the oracle.
public abstract class MIDlet {
    // 0: newly constructed/paused; 2: self-notified paused; 3: destroyed.
    // AMS transitions are tracked by the separate Oak controller, not by
    // inventing an observable MIDP getter for this private implementation.
    private int modelSignal;

    protected MIDlet() { }
    protected abstract void startApp() throws MIDletStateChangeException;
    protected abstract void pauseApp();
    protected abstract void destroyApp(boolean unconditional)
        throws MIDletStateChangeException;

    public final void notifyPaused() { modelSignal = 2; }
    public final void notifyDestroyed() { modelSignal = 3; }
}
