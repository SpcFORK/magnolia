// Synthetic fixture for the Magnolia applet subsystem (no network or AWT UI).
// Regenerate only if needed: javac --release 8 test/fixtures/CounterApplet.java
import java.applet.Applet;

public class CounterApplet extends Applet {
    int count;

    public void init() {
        count = 2 + 3;
        int i = 0;
        while (i < 3) {
            count += i;
            i++;
        }
        System.out.println("Hello from Magnolia");
        System.out.println(count);
        System.out.println(twice(7));
        String name = getParameter("name");
        if (name != null) System.out.println(name);
    }

    public int twice(int n) { return n * 2; }
}
