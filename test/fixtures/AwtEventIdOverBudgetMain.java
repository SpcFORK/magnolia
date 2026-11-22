// JDK permits an arbitrary Event id; Magnolia's direct-call envelope
// deliberately refuses values outside 0..65535 before application output.
import java.awt.Event;
public final class AwtEventIdOverBudgetMain {
    public static void main(String[] args) {
        new Event(null, 65536, null);
        System.out.println("JDK accepted wide Event id");
    }
}
