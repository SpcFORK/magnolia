// Owned Java 8 thread/time bytecode. The real JDK clock has OS scheduling
// drift; the bounded VM is tested separately with an opt-in logical epoch.
// The oracle compares only observable deadline order and minimum elapsed time,
// never claims a fixed wall-clock timestamp or identical JDK wake-up time.
public final class FiveVirtualClockMain implements Runnable {
    public static long first, second, third;
    public static Thread worker;

    public static void main(String[] ignored) {
        first = System.currentTimeMillis();
        worker = new Thread(new FiveVirtualClockMain());
        worker.start();
    }
    // JDK-valid but a different, unmodeled clock/precision/scheduler API.
    public static void otherClock() {
        long sampled = System.nanoTime();
        System.out.println(sampled == sampled);
    }
    @Override public void run() {
        try {
            Thread.sleep(5);
            second = System.currentTimeMillis();
            Thread.sleep(7);
            third = System.currentTimeMillis();
            System.out.println(second - first >= 5);
            System.out.println(third - second >= 7);
            System.out.println(third - first >= 12);
            System.out.println(third >= second && second >= first);
        } catch (InterruptedException unexpected) {
            System.out.println("BAD-interrupted");
        }
    }
}
