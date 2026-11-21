package javax.microedition.midlet;

// Descriptor-compatible owned model; AMS exception transitions are not yet
// supported. An uncaught exception stops Magnolia's bounded probe.
public class MIDletStateChangeException extends Exception {
    public MIDletStateChangeException() { super(); }
    public MIDletStateChangeException(String message) { super(message); }
}
