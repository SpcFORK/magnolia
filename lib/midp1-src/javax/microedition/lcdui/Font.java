package javax.microedition.lcdui;

// Owned finite default-device metrics, NOT a substitute for a platform font.
// These are measured on pinned MicroEmulator 2.0.4/JDK 11 headless with the
// default J2SE device, and are checked again against the real implementation
// by the byte-executed fixture. Device/font/JDK changes must fail the oracle.
// The original Acotel game requests only FACE_SYSTEM, STYLE_BOLD and these
// three sizes. The widths are ASCII 32..126, encoded as one hex digit each.
public final class Font {
    public static final int FACE_SYSTEM = 0;
    public static final int STYLE_BOLD = 1;
    public static final int SIZE_MEDIUM = 0;
    public static final int SIZE_SMALL = 8;
    public static final int SIZE_LARGE = 16;

    private static final String SMALL =
        "3458698344583433666666666644888597777667833769887876677A777434855665664663363966664546686656368";
    private static final String MEDIUM =
        "45698BA3556945448888888888449996B988988994497B998988899C888545966787875884474B888857587A7768489";
    private static final String LARGE =
        "567B9DB4667B5555999999999955BBB8DAAAB99BB55A8DBBABA99BAEA99656B77998996994494E999968698C888959B";

    private static Font small;
    private static Font medium;
    private static Font large;
    private final String widths;
    private final int height;

    private Font(String table, int pixels) { widths = table; height = pixels; }

    public static Font getFont(int face, int style, int size) {
        if (face != FACE_SYSTEM || style != STYLE_BOLD)
            throw new IllegalArgumentException();
        if (size == SIZE_SMALL) {
            if (small == null) small = new Font(SMALL, 12);
            return small;
        }
        if (size == SIZE_MEDIUM) {
            if (medium == null) medium = new Font(MEDIUM, 14);
            return medium;
        }
        if (size == SIZE_LARGE) {
            if (large == null) large = new Font(LARGE, 17);
            return large;
        }
        throw new IllegalArgumentException();
    }

    public int getHeight() { return height; }

    public int charWidth(char character) {
        if (character < 32 || character > 126)
            throw new IllegalArgumentException();
        char digit = widths.charAt(character - 32);
        return digit <= '9' ? digit - '0' : digit - 'A' + 10;
    }

    public int stringWidth(String text) {
        if (text == null) throw new NullPointerException();
        if (text.length() > 256) throw new IllegalArgumentException();
        int width = 0;
        for (int i = 0; i < text.length(); ++i)
            width += charWidth(text.charAt(i));
        return width;
    }
}
