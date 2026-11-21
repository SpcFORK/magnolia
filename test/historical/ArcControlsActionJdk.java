// JDK 11 -Xverify:all/Xvfb reference for the UNCHANGED historical
// ArcControls.action(Event,Object) and ArcCanvas.redraw(boolean,int,int).
// It calls the original methods directly with owner-created AWT objects;
// there is no Magnolia action admission, event queue, layout or arc pixels.
// Compile with --release 8 against the externally SHA-pinned three classes.
import java.awt.Button;
import java.awt.Component;
import java.awt.Container;
import java.awt.Event;
import java.awt.TextField;

public final class ArcControlsActionJdk {
    private static void state(String label, ArcCanvas canvas) {
        System.out.println(label + "|" + canvas.filled + "|" +
            canvas.startAngle + "|" + canvas.endAngle);
    }
    private static void result(String label, boolean handled, ArcCanvas canvas) {
        System.out.println(label + "|" + handled + "|" + canvas.filled +
            "|" + canvas.startAngle + "|" + canvas.endAngle);
    }
    public static void main(String[] args) {
        if (args.length != 0) return;
        ArcCanvas canvas = new ArcCanvas();
        ArcControls controls = new ArcControls(canvas);
        Component[] children = ((Container) controls).getComponents();
        Button fill = (Button) children[2];
        Button draw = (Button) children[3];
        System.out.println("buttons|" + children.length + "|" +
            fill.getLabel() + "|" + draw.getLabel() + "|" +
            (controls.s == children[0]) + "|" + (controls.e == children[1]));
        state("initial", canvas);

        // The branch checks Event.target's TYPE, not id/target membership or
        // argument type. A wrong target must never cast/trim the argument.
        result("non-button", controls.action(
            new Event((TextField) children[0], Event.ACTION_EVENT, "Fill"),
            Integer.valueOf(7)), canvas);
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

        // Any Button target is accepted by this *direct* action call, even
        // one outside controls. The Event id is not consulted by the method.
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
            controls.action(new Event(fill, Event.ACTION_EVENT, null),
                Integer.valueOf(1));
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
