// Owned Java 8 bytecode: Component.getGraphics on a constructed but detached
// Applet/Canvas. A null result is NOT a graphics peer or a drawing surface.
import java.applet.Applet;
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.Graphics;
import java.awt.Panel;
import java.awt.image.BufferedImage;
import java.awt.image.ColorModel;

public class PeerlessGraphicsApplet extends Applet {
    @Override public void init() {
        Component component = this;
        Component canvas = new Canvas();
        System.out.println(component.getParent() == null);
        System.out.println(component.getGraphics() == null);
        System.out.println(component.getGraphics() == null);
        // javac names this user Applet subclass as symbolic owner; the linker
        // must resolve the exact inherited Component method or its override.
        System.out.println(getGraphics() == null);
        System.out.println(canvas.getGraphics() == null);
    }

    // Isolated selection for generated-Java refusal on getGraphics itself,
    // rather than on unrelated Canvas construction in init().
    public void graphicsOnly() {
        System.out.println(((Component) this).getGraphics() == null);
    }

    // An existing headless ColorModel contract also needs virtual override
    // linking when its call-site owner is the built-in Component.
    public void modelOverrideProbe() {
        ColorModel model = ((Component) this).getColorModel();
        System.out.println(model == ColorModel.getRGBdefault());
    }

    // A JDK-valid non-displayable attached Applet still returns null. Magnolia
    // deliberately refuses attached components: it has no parent/peer host.
    public void attachedProbe() {
        Panel parent = new Panel();
        ((Container) parent).add(this);
        System.out.println(((Component) this).getGraphics() == null);
    }
}

// A symbolic invokevirtual on Component MUST preserve virtual overrides.
class PeerlessGraphicsOverrideApplet extends PeerlessGraphicsApplet {
    @Override public Graphics getGraphics() {
        System.out.println("override-graphics");
        return null;
    }
    @Override public ColorModel getColorModel() {
        System.out.println("override-model");
        return ColorModel.getRGBdefault();
    }
}

// Legal Java: a user override may supply its *own* offscreen surface even
// without a peer. Magnolia must preflight/refuse this unsupported body, not
// erase the override by returning its built-in detached null.
class PeerlessGraphicsSurfaceApplet extends PeerlessGraphicsApplet {
    @Override public Graphics getGraphics() {
        return new BufferedImage(1, 1, BufferedImage.TYPE_INT_ARGB).getGraphics();
    }
}
