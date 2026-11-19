import java.awt.Point;

// Independent Java 8/JDK oracle for the bounded, mutable AWT Point host.
// This does not exercise an AWT window or a Flash player lifecycle.
public final class PointGeometryMain {
    public static void main(String[] args) {
        Point origin = new Point();
        Point a = new Point(7, -9);
        Point b = new Point(a);
        System.out.println(origin.x);
        System.out.println(origin.y);
        System.out.println(a.x);
        System.out.println(a.y);
        System.out.println(a == b);
        System.out.println(a == a);
        System.out.println(a.equals(b));

        a.x = Integer.MIN_VALUE;
        a.y = Integer.MAX_VALUE;
        System.out.println(a.x);
        System.out.println(a.y);
        System.out.println(b.x);
        System.out.println(b.y);
        System.out.println(a.equals(b));

        a.move(13, -14);
        b.setLocation(-33, 44);
        System.out.println(a.x);
        System.out.println(a.y);
        System.out.println(b.x);
        System.out.println(b.y);
        Point c = b.getLocation();
        System.out.println(b == c);
        System.out.println(b.equals(c));
        System.out.println(b.equals(null));
        c.y = 55;
        System.out.println(b.y);
        System.out.println(c.y);
        System.out.println(c instanceof java.awt.geom.Point2D);
        System.out.println(c instanceof java.io.Serializable);
        System.out.println(c instanceof Cloneable);
        Point[] copies = new Point[2];
        copies[0] = c;
        System.out.println(copies[0] == c);
        System.out.println(copies[0].x);
    }
}
