// Display-dependent *counterexample* for Component.getColorModel(): the
// detached JDK model is not universally the default RGB model. Run under a
// specific Xvfb 24-bit screen and compare with headless=true, no Applet.
import java.awt.Canvas;
import java.awt.image.ColorModel;

public final class DefaultRGBDisplayJdk {
    public static void main(String[] args) {
        ColorModel component = new Canvas().getColorModel();
        System.out.println(component == ColorModel.getRGBdefault());
        System.out.println(component.getPixelSize());
        System.out.println(component.hasAlpha());
    }
}
