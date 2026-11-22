// Java 8/JDK 11 reference for exact built-in Panel parent-color inheritance.
// No peer, AppletStub, layout geometry, paint or event scheduler is involved.
import java.awt.Canvas;
import java.awt.Color;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;

public final class AwtTreeExplicitInheritedColorsMain {
    private static void state(String label, Component child, Color background,
                              Color foreground) {
        System.out.print(label);
        System.out.print("|");
        System.out.print(child.getBackground() == background);
        System.out.print("|");
        System.out.print(child.getForeground() == foreground);
        System.out.print("|");
        System.out.print(child.isBackgroundSet());
        System.out.print("|");
        System.out.println(child.isForegroundSet());
    }
    public static void main(String[] args) {
        if (args.length != 0) return;
        Panel root = new Panel();
        Panel middle = new Panel();
        Canvas leaf = new Canvas();
        ((Container) root).add(middle);
        ((Container) middle).add(leaf);
        Component top = (Component) root;
        Component mid = (Component) middle;
        Component child = (Component) leaf;
        state("default", child, null, null);
        top.setBackground(Color.RED);
        top.setForeground(Color.BLUE);
        state("root", child, Color.RED, Color.BLUE);
        mid.setBackground(Color.GREEN);
        mid.setForeground(Color.CYAN);
        state("middle", child, Color.GREEN, Color.CYAN);
        child.setBackground(Color.YELLOW);
        child.setForeground(Color.MAGENTA);
        state("child", child, Color.YELLOW, Color.MAGENTA);
        mid.setBackground(null);
        mid.setForeground(null);
        state("middle-reset", child, Color.YELLOW, Color.MAGENTA);
        child.setBackground(null);
        child.setForeground(null);
        state("child-reset", child, Color.RED, Color.BLUE);
        top.setBackground(null);
        top.setForeground(null);
        state("root-reset", child, null, null);
    }
}
