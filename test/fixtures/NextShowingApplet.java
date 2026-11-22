// Owned Java 8 Applet: isShowing() on detached, peerless Components.
// A real display peer is a separate JDK counterexample, not a Magnolia host.
import java.applet.Applet;
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.Panel;

public class NextShowingApplet extends Applet {
    public void probe() {
        Component applet = this;
        Component canvas = new Canvas();
        System.out.println(applet.isShowing());
        // javac may use this user-class owner for an inherited Component call.
        System.out.println(isShowing());
        System.out.println(canvas.isShowing());
    }

    // JDK-valid peerless attached Applets also return false; this bounded
    // host refuses attached geometry rather than claiming a full AWT tree.
    public void attached() {
        Panel parent = new Panel();
        ((Container) parent).add(this);
        System.out.println(((Component) this).isShowing());
    }

    public void otherDisplayMember() {
        System.out.println(((Component) this).isDisplayable());
    }

    public void showingOnly() {
        System.out.println(((Component) this).isShowing());
    }
}

class NextShowingOverrideApplet extends NextShowingApplet {
    @Override public boolean isShowing() {
        System.out.println("override-showing");
        return true;
    }
}
