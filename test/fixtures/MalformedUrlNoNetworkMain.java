// JDK accepts URL parsing without performing network I/O; Magnolia must refuse
// its unadmitted single-string constructor before publishing guest output.
import java.net.URL;
public final class MalformedUrlNoNetworkMain {
    public static void main(String[] args) throws Exception {
        System.out.println("guest-prefix");
        System.out.println(new URL("http://example.invalid/path").getHost());
    }
}
