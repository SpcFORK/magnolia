/** JVM ACC_SYNCHRONIZED reentrance/release fixture; explicit monitorenter is separate. */
public final class SyncMonitorMain {
  private static int total;
  private int value = 5;
  private static synchronized int recurse(int n) {
    if (n == 0) return total;
    total++;
    return recurse(n - 1);
  }
  private synchronized int outer() { return inner() + 1; }
  private synchronized int inner() { return value + 1; }
  private synchronized void fail() { throw new IllegalStateException("expected"); }
  public static void main(String[] argv) {
    System.out.println(recurse(4));
    SyncMonitorMain first = new SyncMonitorMain();
    SyncMonitorMain other = new SyncMonitorMain();
    System.out.println(first.outer());
    try { first.fail(); }
    catch (IllegalStateException expected) { System.out.println("caught"); }
    System.out.println(first.outer());
    System.out.println(other.outer());
    System.out.println(recurse(1));
  }
}
