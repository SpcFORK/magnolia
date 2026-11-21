// JDK-valid parent-color inheritance. Magnolia deliberately refuses instead
// of returning a false null/color when a new detached child has a parent.
import java.awt.Canvas;
import java.awt.Color;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;
public final class AwtTreeInheritedColorMain {
    public static void main(String[] args) {
        Panel parent = new Panel();
        Canvas child = new Canvas();
        ((Component) parent).setBackground(Color.RED);
        ((Container) parent).add(child);
        System.out.println(((Component) child).getBackground() == Color.RED);
    }
}
