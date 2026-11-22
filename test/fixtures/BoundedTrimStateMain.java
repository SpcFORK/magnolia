// Independent JDK 11 and Magnolia native/JS/HTML probe for ASCII String.trim
// and TextComponent.setText. Not a general Unicode/JVM String or GUI peer.
import java.awt.Container;
import java.awt.Panel;
import java.awt.TextField;
public final class BoundedTrimStateMain {
    public static void main(String[] args) {
        String unchanged = "A\t B";
        System.out.println(unchanged.trim() == unchanged);
        System.out.println(" \t A\t B \n".trim().equals(unchanged));
        System.out.println(" \t\n ".trim().length());
        System.out.println(" \u001fA! \n".trim().equals("A!"));
        TextField field = new TextField("old", 4);
        Panel owner = new Panel();
        ((Container) owner).add(field);
        field.setText("  -15\t");
        System.out.println(field.getText().trim().equals("-15"));
        System.out.println(field.getColumns());
        field.setText(null);
        System.out.println(field.getText().equals(""));
    }
}
