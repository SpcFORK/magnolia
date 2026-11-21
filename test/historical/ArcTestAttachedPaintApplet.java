// JDK-valid but OUTSIDE Magnolia's measured draw-call gate: the actual
// ArcTest.init() attaches its Canvas to a BorderLayout. Without layout/peer
// geometry Magnolia must refuse when its original paint asks for bounds().
import java.applet.Applet;
import java.awt.Container;
import java.awt.Graphics;
public final class ArcTestAttachedPaintApplet extends Applet {
    @Override public void paint(Graphics graphics) {
        ArcTest original = new ArcTest();
        original.init();
        ArcCanvas child = (ArcCanvas) ((Container) original).getComponents()[0];
        child.paint(graphics);
    }
}
