// Explicitly constructed peerless Applet, one independently measured JDK 11
// logical font's printable-ASCII stringWidth; NO glyph/pixel painting.
import java.applet.Applet;
import java.awt.Component;
import java.awt.Font;
import java.awt.FontMetrics;
public final class FontMetricsHelveticaMain extends Applet {
    public void init() {
        Font font = new Font("Helvetica", Font.BOLD, 12);
        FontMetrics metrics = ((Component) this).getFontMetrics(font);
        System.out.println("source|" + metrics.stringWidth("www.dseffects.com"));
        System.out.println("empty|" + metrics.stringWidth(""));
        System.out.println("sample|" + metrics.stringWidth("HELLO WORLD!"));
        try { metrics.stringWidth(null); }
        catch (NullPointerException expected) { System.out.println("null|NPE"); }
    }
}
