// Owned JDK 11 -Xverify:all/Xvfb *and* Magnolia harness for UNCHANGED
// external ArcControls.action(Event,Object) and ArcCanvas.redraw. Never
// embeds/recompiles the original classes, attaches a GUI peer, or dispatches
// an actual AWT event: these are deliberate direct method calls.
import java.awt.Button;
import java.awt.Component;
import java.awt.Container;
import java.awt.Event;
import java.awt.TextField;

public final class ArcControlsActionMain {
    private static void part(String value) { System.out.print(value); System.out.print("|"); }
    private static void part(boolean value) { System.out.print(value); System.out.print("|"); }
    private static void part(int value) { System.out.print(value); System.out.print("|"); }
    private static void state(String label, ArcCanvas canvas) {
        part(label);
        part(canvas.filled);
        part(canvas.startAngle);
        System.out.println(canvas.endAngle);
    }
    private static void result(String label, boolean handled, ArcCanvas canvas) {
        part(label);
        part(handled);
        part(canvas.filled);
        part(canvas.startAngle);
        System.out.println(canvas.endAngle);
    }
    public static void main(String[] args) {
        if (args.length != 0) return;
        ArcCanvas canvas = new ArcCanvas();
        ArcControls controls = new ArcControls(canvas);
        Component[] children = ((Container) controls).getComponents();
        Button fill = (Button) children[2];
        Button draw = (Button) children[3];
        part("buttons");
        part(children.length);
        part(fill.getLabel());
        part(draw.getLabel());
        part(controls.s == children[0]);
        System.out.println(controls.e == children[1]);
        state("initial", canvas);

        // A non-Button or null target returns before casting the wrong arg.
        result("non-button", controls.action(
            new Event((TextField) children[0], Event.ACTION_EVENT, "Fill"),
            new Object()), canvas);
        result("null-target", controls.action(
            new Event(null, Event.ACTION_EVENT, "Fill"), "Fill"), canvas);

        controls.s.setText("  -15  ");
        controls.e.setText("+120");
        result("fill", controls.action(
            new Event(fill, Event.ACTION_EVENT, "unused"), "Fill"), canvas);
        controls.s.setText("18");
        controls.e.setText(" -90 \t");
        result("draw", controls.action(
            new Event(draw, Event.ACTION_EVENT, "Fill"), "Draw"), canvas);

        // Direct call checks only the Button TYPE, not membership or id.
        controls.s.setText("-4");
        controls.e.setText("270");
        result("outsider-key", controls.action(
            new Event(new Button("outsider"), Event.KEY_PRESS, null),
            "Fill"), canvas);

        try {
            controls.action(new Event(fill, Event.ACTION_EVENT, null), null);
            System.out.println("null-arg|unexpected");
        } catch (NullPointerException expected) {
            state("null-arg|NPE", canvas);
        }
        try {
            controls.action(new Event(fill, Event.ACTION_EVENT, null), new Object());
            System.out.println("non-string-arg|unexpected");
        } catch (ClassCastException expected) {
            state("non-string-arg|CCE", canvas);
        }
        controls.s.setText("bad");
        controls.e.setText("51");
        try {
            controls.action(new Event(fill, Event.ACTION_EVENT, null), "Draw");
            System.out.println("bad-first|unexpected");
        } catch (NumberFormatException expected) {
            state("bad-first|NFE", canvas);
        }
        controls.s.setText("32");
        controls.e.setText("2147483648");
        try {
            controls.action(new Event(fill, Event.ACTION_EVENT, null), "Draw");
            System.out.println("bad-second|unexpected");
        } catch (NumberFormatException expected) {
            state("bad-second|NFE", canvas);
        }
        try {
            controls.action(null, "Fill");
            System.out.println("null-event|unexpected");
        } catch (NullPointerException expected) {
            state("null-event|NPE", canvas);
        }
    }
}
