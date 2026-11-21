import java.applet.Applet;
import java.awt.Component;
import java.awt.Font;
import java.awt.FontMetrics;

// Independent all-glyph differential to the unchanged original DS JDK
// harness. Does not reproduce any JDK widths: it queries admitted bytecode.
public final class FontMetricsWidthTableMain extends Applet {
    public void init() {
        Font font = new Font("Helvetica", Font.BOLD, 12);
        FontMetrics metrics = ((Component) this).getFontMetrics(font);
        System.out.println("font|" + font.getName() + "|" +
                font.getStyle() + "|" + font.getSize());
        String ascii = " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~";
        // The independent byte-for-byte JDK oracle checks the full literal.
        System.out.print("ascii32-126|");
        for (int i = 0; i < 95; i++) {
            if (i != 0) System.out.print(",");
            System.out.print(metrics.stringWidth(ascii.substring(i, i + 1)));
        }
        System.out.println();
        String[] samples = {"", "www.dseffects.com", "HELLO WORLD!",
                "AV VA", "Java 8: 0123456789", "  '[]~{}"};
        for (int j = 0; j < samples.length; j++) {
            String text = samples[j];
            int sum = 0;
            for (int i = 0; i < text.length(); i++)
                sum += metrics.stringWidth(text.substring(i, i + 1));
            System.out.println("sample|" + text + "|" +
                    metrics.stringWidth(text) + "|" + sum);
        }
    }
}
