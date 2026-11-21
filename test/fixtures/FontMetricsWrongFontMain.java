// JDK-valid, but Font.PLAIN has different metrics. The finite Magnolia
// Helvetica BOLD 12 table must reject this before producing guest output.
import java.applet.Applet;
import java.awt.Component;
import java.awt.Font;
public final class FontMetricsWrongFontMain extends Applet {
    public void init() {
        System.out.println(((Component) this).getFontMetrics(new Font("Helvetica", Font.PLAIN, 12))
            .stringWidth("www.dseffects.com"));
    }
}
