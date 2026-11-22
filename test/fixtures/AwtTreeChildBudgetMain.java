// JDK-valid 17th child. Magnolia deliberately refuses before any output:
// its detached-tree policy caps each Container at 16 children.
import java.awt.Canvas;
import java.awt.Container;
import java.awt.Panel;
public final class AwtTreeChildBudgetMain {
    public static void main(String[] args) {
        Panel panel = new Panel();
        for (int i = 0; i < 17; i++)
            ((Container) panel).add(new Canvas());
        System.out.println(((Container) panel).getComponents().length);
    }
}
