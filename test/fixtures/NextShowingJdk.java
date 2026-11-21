// JDK 11 -Xverify:all reference: detached Applets/Canvas and true real peer.
import java.awt.Canvas;
import java.awt.Frame;

public final class NextShowingJdk {
    public static void main(String[] args) {
        if (args.length != 1) throw new IllegalArgumentException("mode");
        if (args[0].equals("peer")) {
            Frame frame = new Frame("JDK peer counterexample");
            Canvas canvas = new Canvas();
            frame.add(canvas);
            frame.setSize(40, 40);
            frame.setVisible(true);
            try {
                System.out.println(frame.isShowing());
                System.out.println(canvas.isShowing());
            } finally { frame.dispose(); }
            return;
        }
        NextShowingApplet applet = args[0].equals("override")
            ? new NextShowingOverrideApplet() : new NextShowingApplet();
        if (args[0].equals("base") || args[0].equals("override"))
            applet.probe();
        else if (args[0].equals("attached")) applet.attached();
        else if (args[0].equals("other")) applet.otherDisplayMember();
        else throw new IllegalArgumentException(args[0]);
    }
}
