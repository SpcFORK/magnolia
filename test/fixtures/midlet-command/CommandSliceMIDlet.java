import javax.microedition.lcdui.Canvas;
import javax.microedition.lcdui.Command;
import javax.microedition.lcdui.CommandListener;
import javax.microedition.lcdui.Display;
import javax.microedition.lcdui.Displayable;
import javax.microedition.lcdui.Graphics;
import javax.microedition.midlet.MIDlet;
import javax.microedition.midlet.MIDletStateChangeException;

// Identical original fixture bytes on real MicroEmulator and native Magnolia.
// commandAction is explicitly CALLED here by the fixture, not delivered by
// a softkey/device event. Registration is inspected by each bounded harness.
public final class CommandSliceMIDlet extends MIDlet implements CommandListener {
    public static final class View extends Canvas {
        protected void paint(Graphics graphics) { }
    }
    private Command exit;
    private View view;

    protected void startApp() throws MIDletStateChangeException {
        System.out.println("TRACE:command-begin");
        exit = new Command("Fine", Command.SCREEN, 0);
        System.out.println(exit.getLabel());
        System.out.println(exit.getCommandType());
        System.out.println(exit.getPriority());
        Command emptyLabel = new Command(null, Command.EXIT, 7);
        System.out.println(emptyLabel.getLabel() == null
            ? "TRACE:null-label-retained" : "BAD:null-label-not-retained");
        Command oddType = new Command("invalid", 0, -5);
        System.out.println(oddType.getCommandType());
        System.out.println(oddType.getPriority());
        view = new View();
        try {
            view.addCommand(null);
            System.out.println("BAD:null-command-was-accepted");
        } catch (NullPointerException expected) {
            System.out.println("TRACE:null-command-caught");
        }
        view.addCommand(exit);
        view.addCommand(exit); // duplicate must not create a second command
        view.setCommandListener(null);
        view.setCommandListener(this);
        commandAction(exit, view); // explicitly invoked, NOT a softkey
        System.out.println("TRACE:command-end");
        Display.getDisplay(this).setCurrent(view);
    }

    public void commandAction(Command selected, Displayable source) {
        if (selected == exit && source == view) {
            System.out.println("TRACE:direct-callback-identity");
        } else {
            System.out.println("BAD:command-identity");
        }
    }
    protected void pauseApp() { }
    protected void destroyApp(boolean unconditional)
        throws MIDletStateChangeException { notifyDestroyed(); }
}
