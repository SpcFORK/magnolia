// JDK-valid 17-node Panel chain. Magnolia's bounded detached tree refuses
// an attachment deeper than 16 before any output (fatal, not catchable IAE).
import java.awt.Container;
import java.awt.Panel;
public final class AwtTreeDepthBudgetMain {
    public static void main(String[] args) {
        Panel current = new Panel();
        for (int i = 0; i < 16; i++) {
            Panel next = new Panel();
            ((Container) current).add(next);
            current = next;
        }
        System.out.println("JDK accepted 17 Panels");
    }
}
