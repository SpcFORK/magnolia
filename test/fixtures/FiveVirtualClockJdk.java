// Real JDK 11 time/thread oracle for the SAME Java 8 fixture bytecode.
// Check the JDK anchor is near the wall clock; do not assert exact millis,
// thread interleaving or that an opt-in simulated epoch equals JDK time.
public final class FiveVirtualClockJdk {
    public static void main(String[] args) throws Exception {
        long before = System.currentTimeMillis();
        FiveVirtualClockMain.main(new String[0]);
        FiveVirtualClockMain.worker.join(5000);
        long after = System.currentTimeMillis();
        if (FiveVirtualClockMain.worker.isAlive() ||
            FiveVirtualClockMain.first < before ||
            FiveVirtualClockMain.first > after ||
            FiveVirtualClockMain.second - FiveVirtualClockMain.first < 5 ||
            FiveVirtualClockMain.third - FiveVirtualClockMain.second < 7)
            throw new AssertionError("JDK wall clock or Runnable deadline differs");
    }
}
