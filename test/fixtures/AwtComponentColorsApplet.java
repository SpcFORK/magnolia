// JDK 11 reference fixture: inherited Component colors on a detached Applet.
// JVM class is compiled with `javac --release 8 -g:none`. The Oak run invokes
// the same init/start path, not Java's static main or an AWT paint callback.
import java.applet.Applet;
import java.awt.Color;
import java.awt.Component;

public final class AwtComponentColorsApplet extends Applet {
    public void init() {
        System.out.println(getBackground() == null);
        System.out.println(getForeground() == null);
        System.out.println(!isBackgroundSet());
        System.out.println(!isForegroundSet());
        Color background = Color.PINK;
        Color foreground = Color.BLUE;
        setBackground(background);
        setForeground(foreground);
        System.out.println(getBackground() == background);
        System.out.println(getForeground() == foreground);
        System.out.println(isBackgroundSet());
        System.out.println(isForegroundSet());
        System.out.println(getBackground() == Color.pink);
        System.out.println(getForeground() == Color.blue);
        setBackground(null);
        System.out.println(getBackground() == null);
        System.out.println(!isBackgroundSet());
        Component asComponent = this;
        asComponent.setForeground(Color.RED);
        System.out.println(asComponent.getForeground() == Color.RED);
        System.out.println(asComponent.isForegroundSet());
        System.out.println(getForeground() == foreground);
    }
    public void start() { System.out.println("started"); }
    public static void main(String[] argv) {
        AwtComponentColorsApplet player = new AwtComponentColorsApplet();
        player.init();
        player.start();
    }
}
