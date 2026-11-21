// Compile ONLY alongside separately obtained, hash-pinned MyApplet.class.
// This checks the original applet's lifecycle/color state, not its Font paint.
import java.awt.Color;
public final class MyAppletInitJdk {
    public static void main(String[] argv) {
        MyApplet applet = new MyApplet();
        applet.init();
        applet.start();
        System.out.println(applet.getBackground() == Color.PINK);
        System.out.println(applet.getForeground() == Color.BLUE);
    }
}
