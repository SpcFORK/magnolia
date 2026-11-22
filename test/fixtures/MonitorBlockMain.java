/** JDK 8-verified explicit monitorenter/exit with compiler-generated finally. */
public final class MonitorBlockMain {
  private static int total;
  private static int step(Object lock, int n) {
    synchronized (lock) {
      total += n;
      synchronized (lock) { total++; } // explicit reentrance, same identity
    }
    return total;
  }
  private static void throwWhileHeld(Object lock) {
    synchronized (lock) { throw new IllegalArgumentException("expected"); }
  }
  public static void main(String[] ignored) {
    Object first = new Object();
    Object second = new Object();
    System.out.println(step(first, 3));
    System.out.println(step(second, 2));
    try { throwWhileHeld(first); }
    catch (IllegalArgumentException expected) { System.out.println("caught"); }
    System.out.println(step(first, 1));
    try { synchronized ((Object)null) { System.out.println("unreachable"); } }
    catch (NullPointerException expected) { System.out.println("null"); }
    int[] array = new int[2];
    synchronized (array) { total++; }
    System.out.println(total);
  }
}
