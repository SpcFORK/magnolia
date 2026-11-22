// Component invokevirtual must dispatch to the preflighted concrete user
// override, not the built-in Color field accessor. The parent-color profile
// still refuses virtual getter dispatch on a parent subclass.
import java.awt.Canvas;
import java.awt.Color;
import java.awt.Component;

public final class AwtTreeOverriddenReceiverColorMain {
    static final class TintCanvas extends Canvas {
        @Override public Color getBackground() { return Color.RED; }
    }
    public static void main(String[] args) {
        Component child = new TintCanvas();
        System.out.println(child.getBackground() == Color.RED);
    }
}
