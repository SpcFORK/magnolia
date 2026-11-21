/** For a checked JDK-verified wrong-receiver monitorexit mutation. */
public final class MonitorMismatchMain {
  static void gate(Object held, Object other) {
    synchronized (held) { System.out.println("inside"); }
  }
  public static void main(String[] args) {
    try { gate(new Object(), new Object()); System.out.println("after"); }
    catch (IllegalMonitorStateException expected) { System.out.println("wrong"); }
  }
}
