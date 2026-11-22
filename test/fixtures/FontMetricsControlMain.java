// JDK accepts control characters in stringWidth; Magnolia's measured
// table is printable ASCII U+0020..U+007E only and must refuse pre-output.
import java.applet.Applet;
import java.awt.Component;
import java.awt.Font;
public final class FontMetricsControlMain extends Applet {
    public void init() {
        System.out.println(((Component) this).getFontMetrics(new Font("Helvetica", Font.BOLD, 12))
            .stringWidth("ABC\nDEF"));
    }
}
