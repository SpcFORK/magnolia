// JDK 11 -Xverify:all/Xvfb oracle for THREE original Sun class files:
// ArcTest, ArcCanvas, ArcControls. Init/start/stop only, with a synthetic
// AppletStub. No canvas painting, interaction, event queue, applet context,
// image/font pixel parity or Magnolia execution is claimed by this harness.
// Compile against hash-checked external bytes and an EMPTY -sourcepath.
import java.applet.AppletContext;
import java.applet.AppletStub;
import java.awt.BorderLayout;
import java.awt.Button;
import java.awt.Component;
import java.net.URL;

public final class ArcTestInitJdk {
    private static final class Stub implements AppletStub {
        private boolean active;
        public boolean isActive() { return active; }
        void setActive(boolean value) { active = value; }
        public URL getDocumentBase() { return getCodeBase(); }
        public URL getCodeBase() {
            try { return new URL("file:/"); }
            catch (Exception e) { throw new AssertionError(e); }
        }
        public String getParameter(String name) { return null; }
        public AppletContext getAppletContext() { return null; }
        public void appletResize(int width, int height) {
            throw new AssertionError("unexpected resize during tested lifecycle");
        }
    }
    public static void main(String[] args) {
        if (args.length != 0) throw new IllegalArgumentException("no external inputs");
        ArcTest applet = new ArcTest();
        Stub stub = new Stub();
        applet.setStub(stub);
        applet.init();
        Component[] children = applet.getComponents();
        if (children.length != 2 || !(children[0] instanceof ArcCanvas) ||
                !(children[1] instanceof ArcControls) ||
                children[0].getParent() != applet || children[1].getParent() != applet)
            throw new AssertionError("unexpected original applet child graph");
        ArcCanvas canvas = (ArcCanvas) children[0];
        ArcControls controls = (ArcControls) children[1];
        if (applet.controls != controls || controls.canvas != canvas)
            throw new AssertionError("unexpected original control/canvas references");
        Component[] buttons = controls.getComponents();
        if (buttons.length != 4 || buttons[0] != controls.s ||
                buttons[1] != controls.e || !(buttons[2] instanceof Button) ||
                !(buttons[3] instanceof Button))
            throw new AssertionError("unexpected original controls graph");
        BorderLayout layout = (BorderLayout) applet.getLayout();
        System.out.println("init|" + layout.getClass().getName() +
                "|" + children.length +
                "|" + children[0].getClass().getName() +
                "|" + children[1].getClass().getName() +
                "|" + layout.getConstraints(children[0]) +
                "|" + layout.getConstraints(children[1]) +
                "|" + canvas.startAngle +
                "|" + canvas.endAngle +
                "|" + canvas.filled +
                "|" + controls.s.getText() +
                "|" + controls.s.getColumns() +
                "|" + controls.e.getText() +
                "|" + controls.e.getColumns() +
                "|" + ((Button) buttons[2]).getLabel() +
                "|" + ((Button) buttons[3]).getLabel() +
                "|" + controls.isEnabled());
        stub.setActive(true);
        applet.start();
        System.out.println("start|" + controls.isEnabled());
        stub.setActive(false);
        applet.stop();
        System.out.println("stop|" + controls.isEnabled());
        stub.setActive(true);
        applet.start();
        System.out.println("restart|" + controls.isEnabled());
        stub.setActive(false);
        applet.stop();
        applet.destroy();
    }
}
