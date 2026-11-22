// JDK accepts negative dimensions; bounded Magnolia geometry refuses them
// rather than interpreting an unmapped peer/layout state as a real rectangle.
import java.awt.Canvas;
import java.awt.Component;
public final class AwtCanvasNegativeBoundsMain {
    public static void main(String[] args) {
        Component canvas = (Component) new Canvas();
        canvas.setBounds(0, 0, -1, 3);
        System.out.println(canvas.bounds().width);
    }
}
