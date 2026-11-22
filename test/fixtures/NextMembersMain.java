// Owned Java 8: three separate bounded member behaviors from original Flash.
// Each selected method has its own JDK 11 -Xverify:all expected output.
import java.awt.Rectangle;

public final class NextMembersMain {
    public static void rectangleOnly() {
        Rectangle first = new Rectangle(-10, Integer.MAX_VALUE, -2, 0);
        Rectangle second = new Rectangle(7, -5, 0, 4);
        System.out.println(first != second);
        System.out.println(first.x);
        System.out.println(first.y);
        System.out.println(first.width);
        System.out.println(first.height);
        System.out.println(second.x);
        System.out.println(second.y);
        System.out.println(second.width);
        System.out.println(second.height);
    }

    public synchronized void implicitNotify() {
        notifyAll(); // ACC_SYNCHRONIZED owns this monitor without monitorenter.
        System.out.println("implicit-owner");
    }

    public static void notifyOnly() {
        Object lock = new Object();
        synchronized (lock) {
            lock.notifyAll();
            synchronized (lock) {
                lock.notifyAll();
                System.out.println("nested-owner");
            }
        }
        new NextMembersMain().implicitNotify();
        try {
            lock.notifyAll();
            System.out.println("BAD-unowned");
        } catch (IllegalMonitorStateException expected) {
            System.out.println("not-owner");
        }
        try {
            ((Object) null).notifyAll();
            System.out.println("BAD-null");
        } catch (NullPointerException expected) {
            System.out.println("null-caught");
        }
    }

    public static void yieldOnly() {
        int[] state = new int[1];
        state[0] = 17;
        Thread.yield();
        System.out.println(state[0]);
        state[0] = 29;
        Thread.yield();
        System.out.println(state[0]);
    }

    public static void yieldOverBound() {
        for (int i = 0; i < 129; i++) Thread.yield();
        System.out.println("JDK-129-yields");
    }
    public static void notifyOverBound() {
        Object lock = new Object();
        synchronized (lock) {
            for (int i = 0; i < 257; i++) lock.notifyAll();
        }
        System.out.println("JDK-257-notifications");
    }

    // Isolated symbols for exact generated-Java refusal checks.
    public static void notifyParameter(Object lock) { lock.notifyAll(); }
    public static void rectangleConstructor() {
        System.out.println(new Rectangle(1, 2, 3, 4) != null);
    }

    // Each of these is valid JDK bytecode but is NOT within the new contracts.
    public static void replaceRectangle() {
        Rectangle rectangle = new Rectangle();
        System.out.println(rectangle.width);
    }
    public static void mutateRectangle() {
        Rectangle rectangle = new Rectangle(1, 2, 3, 4);
        rectangle.x = 99;
        System.out.println(rectangle.x);
    }
    public static void notifyOne() {
        Object lock = new Object();
        synchronized (lock) { lock.notify(); }
    }
    public static void realWait() throws InterruptedException {
        Object lock = new Object();
        synchronized (lock) { lock.wait(1L); }
    }
    public static void sleepWithoutTask() throws InterruptedException {
        Thread.sleep(1L);
        System.out.println("JDK-slept");
    }
}
