package javax.microedition.lcdui;

// Bounded MIDP-1 command identity/data. There are NO softkeys, menus or
// device-delivered command callbacks in this model. Only exact bytecode APIs
// exercised by the independent fixture and the original game's constructor.
public final class Command {
    public static final int SCREEN = 1;
    public static final int BACK = 2;
    public static final int CANCEL = 3;
    public static final int OK = 4;
    public static final int HELP = 5;
    public static final int STOP = 6;
    public static final int EXIT = 7;
    public static final int ITEM = 8;

    private final String label;
    private final int type;
    private final int priority;

    public Command(String label, int type, int priority) {
        // The pinned MicroEmulator retains even null label / out-of-range
        // type in this three-argument constructor. Do not invent an exception
        // from a different handset or MIDP revision.
        this.label = label;
        this.type = type;
        this.priority = priority;
    }
    public String getLabel() { return label; }
    public int getCommandType() { return type; }
    public int getPriority() { return priority; }
}
