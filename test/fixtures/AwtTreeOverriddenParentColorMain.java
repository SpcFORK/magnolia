// JDK-valid override: reading a child's color must dispatch to its parent's
// virtual getBackground, not just read the parent's raw Component field.
// Magnolia's finite built-in-Panel profile must refuse this child lookup.
import java.awt.Canvas;
import java.awt.Color;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;

public final class AwtTreeOverriddenParentColorMain {
    static final class TintPanel extends Panel {
        @Override public Color getBackground() { return Color.ORANGE; }
    }
    public static void main(String[] args) {
        Panel parent = new TintPanel();
        ((Component) parent).setBackground(Color.RED);
        Component child = (Component) new Canvas();
        ((Container) parent).add(child);
        System.out.println(child.getBackground() == Color.ORANGE);
    }
}
