// Additional headless Java 8 fixture: supported host-call and null semantics.
import java.applet.Applet;

public class ValuesApplet extends Applet {
    @Override public void init() {
        System.out.println(true);
        System.out.println(false);
        System.out.println(new StringBuilder().append((String) null).toString());
        System.out.println(Integer.parseInt(getParameter("n")));
        System.out.println(Integer.toString(-42));
        if ("a".equals("a")) System.out.print("ok\n");
    }
}
