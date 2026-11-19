// The Java frontend compiles against a verified, explicitly supplied helper
// class without relying on javac's ambient classpath or source discovery.
public class UsesHelper {
    public static void main(String[] args) {
        System.out.println(Helper.ignored());
    }
}
