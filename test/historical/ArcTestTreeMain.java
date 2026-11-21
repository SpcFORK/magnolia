// Independent Java 8-compiled, JDK 11 -Xverify:all/Xvfb state oracle.
// Calls the THREE unchanged original Sun 1.0.2 classes. The same four class
// bytes (including THIS harness) run under Magnolia's bounded, detached AWT
// tree. This checks init/start/stop/restart state; no painting, events,
// window, AWT peer, layout geometry or applet context is exercised.
// Compile against external pinned classes with an EMPTY -sourcepath.
import java.awt.BorderLayout;
import java.awt.Button;
import java.awt.Component;
import java.awt.Container;

public final class ArcTestTreeMain {
    private static void part(int value) {
        System.out.print(value);
        System.out.print("|");
    }
    private static void part(boolean value) {
        System.out.print(value);
        System.out.print("|");
    }
    private static void part(String value) {
        System.out.print(value);
        System.out.print("|");
    }
    public static void main(String[] args) {
        if (args.length != 0) return;
        ArcTest applet = new ArcTest();
        applet.init();
        // Exact AWT declaring owners keep the bounded Magnolia host contract
        // explicit (no broad java.awt subclass-method inference).
        Component[] children = ((Container) applet).getComponents();
        ArcControls controls = applet.controls;
        Component[] parts = ((Container) controls).getComponents();
        BorderLayout layout = (BorderLayout) ((Container) applet).getLayout();
        System.out.print("init|");
        part(children.length);
        part(children[0] instanceof ArcCanvas);
        part(children[1] instanceof ArcControls);
        part(children[0].getParent() == applet);
        part(children[1].getParent() == applet);
        part(controls == children[1]);
        part(controls.canvas == children[0]);
        part((String) layout.getConstraints(children[0]));
        part((String) layout.getConstraints(children[1]));
        part(((ArcCanvas) children[0]).startAngle);
        part(((ArcCanvas) children[0]).endAngle);
        part(((ArcCanvas) children[0]).filled);
        part(parts.length);
        part(parts[0] == controls.s);
        part(parts[1] == controls.e);
        part(parts[2] instanceof Button);
        part(parts[3] instanceof Button);
        part(parts[0].getParent() == controls);
        part(parts[3].getParent() == controls);
        part(controls.s.getText());
        part(controls.s.getColumns());
        part(controls.e.getText());
        part(controls.e.getColumns());
        part(((Button) parts[2]).getLabel());
        part(((Button) parts[3]).getLabel());
        System.out.println(((Component) controls).isEnabled());
        applet.start();
        System.out.print("start|");
        System.out.println(((Component) controls).isEnabled());
        applet.stop();
        System.out.print("stop|");
        System.out.println(((Component) controls).isEnabled());
        applet.start();
        System.out.print("restart|");
        System.out.println(((Component) controls).isEnabled());
        applet.stop();
    }
}
