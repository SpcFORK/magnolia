// Bounded, detached AWT tree constructor/accessor/exception differential.
// No peer, layout geometry, paint, event dispatch or implicit network access.
import java.awt.BorderLayout;
import java.awt.Button;
import java.awt.Canvas;
import java.awt.Component;
import java.awt.Container;
import java.awt.FlowLayout;
import java.awt.Panel;
import java.awt.TextField;

public final class AwtTreeStateMain {
    private static void part(int n) { System.out.print(n); System.out.print("|"); }
    private static void part(boolean b) { System.out.print(b); System.out.print("|"); }
    private static void part(String s) { System.out.print(s); System.out.print("|"); }
    public static void main(String[] args) {
        Panel panel = new Panel();
        System.out.print("flow|");
        part(((Container) panel).getLayout() instanceof FlowLayout);
        System.out.println(((Component) panel).getParent() == null);

        BorderLayout layout = new BorderLayout();
        ((Container) panel).setLayout(layout);
        Canvas center = new Canvas();
        Canvas south = new Canvas();
        ((Container) panel).add("Center", center);
        ((Container) panel).add("South", south);
        Component[] children = ((Container) panel).getComponents();
        System.out.print("border|");
        part(children.length);
        part((String) layout.getConstraints(children[0]));
        part((String) layout.getConstraints(children[1]));
        part(((Component) center).getParent() == panel);
        System.out.println(((Component) south).getParent() == panel);
        children[0] = null;
        System.out.print("snapshot|");
        System.out.println(((Container) panel).getComponents()[0] == center);

        boolean invalid = false;
        try { ((Container) panel).add("Diagonal", new Canvas()); }
        catch (IllegalArgumentException exception) { invalid = true; }
        System.out.print("invalid|");
        part(invalid);
        System.out.println(((Container) panel).getComponents().length);
        boolean missing = false;
        try { ((Container) panel).add("North", (Component) null); }
        catch (NullPointerException exception) { missing = true; }
        System.out.print("null|");
        part(missing);
        System.out.println(((Container) panel).getComponents().length);

        Canvas replacement = new Canvas();
        ((Container) panel).add("Center", replacement);
        System.out.print("duplicate|");
        part(layout.getConstraints(center) == null);
        part((String) layout.getConstraints(replacement));
        System.out.println(((Container) panel).getComponents().length);

        Panel parent = new Panel();
        Panel child = new Panel();
        ((Container) parent).add(child);
        boolean cycle = false;
        try { ((Container) child).add(parent); }
        catch (IllegalArgumentException exception) { cycle = true; }
        System.out.print("cycle|");
        System.out.println(cycle);

        TextField text = new TextField(null, -3);
        System.out.print("text|");
        part(text.getText());
        System.out.println(text.getColumns());
        Button button = new Button(null);
        System.out.print("button|");
        System.out.println(button.getLabel());
        System.out.print("enabled|");
        part(((Component) center).isEnabled());
        ((Component) center).disable();
        part(((Component) center).isEnabled());
        ((Component) center).enable();
        System.out.println(((Component) center).isEnabled());
    }
}
