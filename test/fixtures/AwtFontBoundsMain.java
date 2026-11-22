// This is JDK-valid. Magnolia deliberately refuses this size *as a resource
// policy*, not as a JVM VerifyError or a claim that Java disallows the Font.
import java.awt.Font;
public final class AwtFontBoundsMain {
    public static void main(String[] args) {
        System.out.println(new Font("Tahoma", Font.PLAIN, 257).getSize());
    }
}
