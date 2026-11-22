// Owned Java 8 bytecode: one monitor, one waiter, one notifier, no AWT peer.
// The independent JDK runner joins the outer worker; Magnolia's bounded
// scheduler executes that worker synchronously, then runs a queued notifier
// ONLY during Object.wait() after releasing the whole reentrant monitor.
public final class FlashWaitHandoffMain implements Runnable {
    static final Object LOCK = new Object();
    static final Object OTHER = new Object();
    static Thread worker;
    static int phase;
    private final int role;

    public FlashWaitHandoffMain(int role) { this.role = role; }

    public static void positive() {
        phase = 0;
        worker = new Thread(new FlashWaitHandoffMain(0));
        worker.start();
    }

    public static void wrongNotifier() {
        phase = 0;
        worker = new Thread(new FlashWaitHandoffMain(3));
        worker.start();
    }

    public static void noNotifier() {
        phase = 0;
        worker = new Thread(new FlashWaitHandoffMain(4));
        worker.start();
    }

    public static void twoNotifiers() {
        phase = 0;
        worker = new Thread(new FlashWaitHandoffMain(5));
        worker.start();
    }

    public static void noMonitorStartInTask() {
        worker = new Thread(new FlashWaitHandoffMain(6));
        worker.start();
    }

    public static void priorNotification() {
        phase = 0;
        worker = new Thread(new FlashWaitHandoffMain(7));
        worker.start();
    }

    public static void notOwned() {
        try {
            LOCK.wait();
            System.out.println("BUG-unowned-returned");
        } catch (IllegalMonitorStateException expected) {
            System.out.println("unowned");
        } catch (InterruptedException unexpected) {
            System.out.println("BUG-unowned-interrupted");
        }
    }

    public static void nullWait() {
        Object none = null;
        try {
            none.wait();
            System.out.println("BUG-null-returned");
        } catch (NullPointerException expected) {
            System.out.println("null");
        } catch (InterruptedException unexpected) {
            System.out.println("BUG-null-interrupted");
        }
    }

    public static void timedWait() {
        synchronized (LOCK) {
            try { LOCK.wait(1L); } catch (InterruptedException impossible) {
                System.out.println("BUG-timed-interrupted");
            }
        }
        System.out.println("jdk-timeout");
    }

    public static void noMonitorStart() {
        new Thread(new FlashWaitHandoffMain(1)).start();
    }

    public void run() {
        if (role == 6) {
            new Thread(new FlashWaitHandoffMain(1)).start();
        } else if (role == 1) {
            synchronized (LOCK) {
                phase = 7;
                System.out.println("notifier");
                LOCK.notifyAll();
            }
        } else if (role == 2) {
            synchronized (OTHER) { OTHER.notifyAll(); }
        } else {
            synchronized (LOCK) {
                synchronized (LOCK) {
                    if (role == 5) {
                        new Thread(new FlashWaitHandoffMain(1)).start();
                        new Thread(new FlashWaitHandoffMain(1)).start();
                    } else if (role != 4) {
                        new Thread(new FlashWaitHandoffMain(
                                role == 3 || role == 7 ? 2 : 1)).start();
                    }
                    if (role == 7) LOCK.notifyAll(); // before wait: lost
                    if (role == 0) System.out.println("before");
                    while (phase == 0) {
                        try {
                            LOCK.wait();
                        } catch (InterruptedException interrupted) {
                            if (role == 0) System.out.println("BUG-positive-interrupted");
                            else System.out.println("jdk-interrupted");
                            return;
                        }
                    }
                    if (role == 0) {
                        LOCK.notifyAll(); // verifies re-acquired ownership
                        System.out.println("after");
                        System.out.println(phase);
                    }
                }
            }
        }
    }
}
