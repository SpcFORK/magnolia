// JDK-valid detached geometry intentionally outside Magnolia's 4096 bound.
import java.awt.Canvas;
import java.awt.Component;
public final class AwtCanvasBoundsOverBudgetMain {
    public static void main(String[] args) {
        Component canvas = (Component) new Canvas();
        canvas.setBounds(0, 0, 4097, 1);
        System.out.println(canvas.bounds().width);
    }
}
