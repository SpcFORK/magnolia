package javax.microedition.lcdui;

// Logical displayable identity and at most four registered commands. There
// is no hardware softkey, show callback, painting or command event dispatch.
public abstract class Displayable {
    private final Command[] registered = new Command[4];
    private int commandCount;
    private CommandListener commandListener;

    protected Displayable() { }

    // The independently observed default MicroEmulator 2.0.4 device is
    // 176x176 both before and after this Canvas is selected. These are NOT
    // handset-independent dimensions or a visible framebuffer.
    public int getWidth() { return 176; }
    public int getHeight() { return 176; }

    public void addCommand(Command command) {
        if (command == null) throw new NullPointerException();
        for (int i = 0; i < commandCount; i++) {
            if (registered[i] == command) return;
        }
        if (commandCount == registered.length) throw new IllegalStateException();
        registered[commandCount] = command;
        commandCount++;
    }
    public void setCommandListener(CommandListener listener) {
        commandListener = listener;
    }
}
