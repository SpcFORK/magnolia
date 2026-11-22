// JDK-valid reparenting. Magnolia deliberately refuses before any output:
// a nonempty child parent is outside this finite detached-tree island.
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;
public final class AwtTreeReparentMain {
    public static void main(String[] args) {
        Panel first = new Panel(), second = new Panel();
        Canvas child = new Canvas();
        ((Container) first).add(child);
        ((Container) second).add(child);
        System.out.println(((Component) child).getParent() == second);
    }
}
