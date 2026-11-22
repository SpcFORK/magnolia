// Independent Java 8/JDK 11 -Xverify:all detached-Canvas geometry oracle.
// No peer, parent layout, rendering, images or event delivery is involved.
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Rectangle;

public final class AwtCanvasBoundsMain {
    private static void snapshot(String label, Component component) {
        Rectangle first = component.bounds();
        Rectangle second = component.bounds();
        System.out.print(label + "|" + first.x + "|" + first.y + "|" +
            first.width + "|" + first.height + "|");
        System.out.println(first != second);
    }
    public static void main(String[] args) {
        if (args.length != 0) return;
        Component canvas = (Component) new Canvas();
        snapshot("default", canvas);
        canvas.setBounds(7, 11, 31, 21);
        snapshot("odd", canvas);
        canvas.setBounds(-4, 5, 20, 20);
        snapshot("shifted", canvas);
        canvas.setBounds(0, 0, 1, 1);
        snapshot("unit", canvas);
        canvas.setBounds(0, 0, 0, 0);
        snapshot("empty", canvas);
    }
}
