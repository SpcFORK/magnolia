// Java 1.1-style StringBuffer calls still used in historical applets.
import java.applet.Applet;

public class BufferApplet extends Applet {
    @Override public void init() {
        StringBuffer message = new StringBuffer("Ready: ");
        message.append(7);
        System.out.println(message.toString());
    }
}
