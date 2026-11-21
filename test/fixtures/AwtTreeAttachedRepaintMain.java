// JDK-valid attached-but-peerless repaint. Bounded Magnolia deliberately
// refuses: attached child geometry, repaint scheduling and pixels are open.
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;
public final class AwtTreeAttachedRepaintMain {
    public static void main(String[] args) {
        Panel parent = new Panel();
        Canvas child = new Canvas();
        ((Container) parent).add(child);
        ((Component) child).repaint();
        System.out.println("JDK attached Canvas repaint returned");
    }
}
