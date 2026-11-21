// A detached tree has no BorderLayout geometry: Magnolia must not fabricate
// Component.bounds() for a child once an AWT parent can lay it out.
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;
public final class AwtCanvasAttachedBoundsMain {
    public static void main(String[] args) {
        Panel parent = new Panel();
        Component child = (Component) new Canvas();
        ((Container) parent).add(child);
        System.out.println(child.bounds().width);
    }
}
