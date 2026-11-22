// JDK-valid detached Canvas repaint() returns; Magnolia's bounded
// constructed peerless Canvas path also returns without a peer, queue or pixels.
import java.awt.Canvas;
import java.awt.Component;
public final class AwtTreeRepaintMain {
    public static void main(String[] args) {
        Canvas canvas = new Canvas();
        ((Component) canvas).repaint();
        System.out.println("JDK Canvas repaint returned");
    }
}
