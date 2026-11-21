package javax.microedition.lcdui;

// Type-only callback contract; the native MIDlet probe does not deliver
// device softkeys or invent commandAction() events.
public interface CommandListener {
    void commandAction(Command command, Displayable displayable);
}
