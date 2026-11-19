// Java 8 class-initialization regression: preparation, parent-before-child,
// active-use triggers, recursion through getstatic, and exactly-once behavior.
class InitBase {
    static int base = 4;
    static {
        System.out.println("base");
        base += 2;
    }
}

class InitHelper {
    static int runs = 0;
    static {
        runs++;
        System.out.println("helper");
    }
    static int twice() { return runs * 2; }
}

// Merely supplying a class must not force preflight of its unused initializer.
class UnreachableInit {
    static { new java.awt.Font("Dialog", java.awt.Font.PLAIN, 12); }
}

public class InitMain extends InitBase {
    public static final int MAGIC = 42;
    public static final String LABEL = "prepared";
    static int count = 1;
    static {
        count += base;
        System.out.println("entry");
    }
    public static void main(String[] args) {
        System.out.println(MAGIC);
        System.out.println(LABEL);
        System.out.println(count);
        System.out.println(InitHelper.twice());
        System.out.println(InitHelper.twice());
        System.out.println(count);
    }
}
