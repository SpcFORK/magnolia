// Java 8 class: exact, bounded Font constructor/property shape used by
// original MyApplet paint. The JDK is the behavioral oracle, not this text.
import java.awt.Font;
import java.io.Serializable;

public final class AwtFontStateMain {
    public static void main(String[] args) {
        Font requested = new Font("Tahoma", Font.BOLD | Font.ITALIC, 20);
        System.out.println(requested.getName());
        System.out.println(requested.getStyle());
        System.out.println(requested.getSize());
        System.out.println(requested instanceof Serializable);
        Font normalized = new Font(null, 7, 12);
        System.out.println(normalized.getName());
        System.out.println(normalized.getStyle());
        System.out.println(normalized.getSize());
        System.out.println(normalized != requested);
    }
}
