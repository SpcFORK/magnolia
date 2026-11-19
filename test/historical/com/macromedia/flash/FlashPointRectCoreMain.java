package com.macromedia.flash;

import java.awt.Point;

// Independent JDK oracle invoking the ORIGINAL Flash 2 Rect class and AWT
// Point fields/constructors. The Unicode escapes are original JVM member
// identifiers, not substituted geometry algorithms.
public final class FlashPointRectCoreMain {
    public static void main(String[] args) {
        Rect box = new Rect(2, 3, 10, 12);
        Point sample = new Point(5, 6);
        System.out.println(box.\u0181(sample));
        sample.x = 11;
        System.out.println(box.\u0181(sample));
        sample.x = 2;
        sample.y = 12;
        System.out.println(box.\u0181(sample));
        sample.y = 13;
        System.out.println(box.\u0181(sample));

        Rect accumulated = new Rect();
        accumulated.\u017f(new Point(4, 5));
        accumulated.\u017f(new Point(-2, 8));
        System.out.println(accumulated.\u0181(new Point(-2, 8)));
        System.out.println(accumulated.\u0181(new Point(3, 5)));
        System.out.println(accumulated.\u0181(new Point(5, 8)));
        System.out.println(accumulated.\u0180(box));
        System.out.println(accumulated.\u0180(new Rect(50, 50, 60, 60)));
        accumulated.\u017f(new Point(20, 30));
        System.out.println(accumulated.\u0181(new Point(20, 30)));
        System.out.println(accumulated.\u0181(new Point(21, 30)));
    }
}
