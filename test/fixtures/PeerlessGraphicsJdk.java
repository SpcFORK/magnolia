// Independent JDK 11 reference. The fixture uses real Java Component calls;
// this harness checks headless Canvas, Xvfb detached/attached Applets and a
// displayable peer. No historical Flash class, installer or SWF is executed.
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Frame;
import java.awt.Graphics;
import java.awt.GraphicsEnvironment;

public final class PeerlessGraphicsJdk {
    public static void main(String[] args) {
        if (args.length != 1) throw new IllegalArgumentException("one mode");
        switch (args[0]) {
            case "headless-canvas":
                if (!GraphicsEnvironment.isHeadless()) throw new AssertionError("headless only");
                System.out.println(((Component) new Canvas()).getGraphics() == null);
                break;
            case "plain":
                requireDisplay();
                PeerlessGraphicsApplet plain = new PeerlessGraphicsApplet();
                if (plain.isDisplayable()) throw new AssertionError("unexpected peer");
                plain.init();
                break;
            case "override":
                requireDisplay();
                PeerlessGraphicsApplet overridden = new PeerlessGraphicsOverrideApplet();
                if (overridden.isDisplayable()) throw new AssertionError("unexpected peer");
                overridden.init();
                break;
            case "custom-surface":
                requireDisplay();
                new PeerlessGraphicsSurfaceApplet().init();
                break;
            case "only":
                requireDisplay();
                new PeerlessGraphicsApplet().graphicsOnly();
                break;
            case "model-override":
                requireDisplay();
                new PeerlessGraphicsOverrideApplet().modelOverrideProbe();
                break;
            case "attached":
                requireDisplay();
                new PeerlessGraphicsApplet().attachedProbe();
                break;
            case "peer":
                requireDisplay();
                Frame frame = new Frame("JDK peer probe");
                try {
                    PeerlessGraphicsApplet child = new PeerlessGraphicsApplet();
                    frame.add(child);
                    frame.setSize(80, 80);
                    frame.setVisible(true);
                    Graphics graphics = ((Component) child).getGraphics();
                    try {
                        System.out.println(child.getParent() != null);
                        System.out.println(child.isDisplayable());
                        System.out.println(graphics != null);
                    } finally {
                        if (graphics != null) graphics.dispose();
                    }
                } finally {
                    frame.dispose();
                }
                break;
            default:
                throw new IllegalArgumentException(args[0]);
        }
    }

    private static void requireDisplay() {
        if (GraphicsEnvironment.isHeadless()) throw new AssertionError("display required");
    }
}
