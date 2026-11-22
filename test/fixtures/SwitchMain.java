// Differential test of JVM tableswitch/lookupswitch, signed keys and defaults.
public final class SwitchMain {
    static int dense(int n) {
        switch (n) {
            case -2: return 11;
            case -1: return -7;
            case 0: return 0;
            case 1: return 17;
            case 2: return 20;
            default: return 31;
        }
    }

    static int sparse(int n) {
        switch (n) {
            case -2147483648: return 5;
            case -100: return 7;
            case 37: return 9;
            case 1000000: return 13;
            default: return 17;
        }
    }

    public static void main(String[] args) {
        for (int n = -3; n <= 3; n++) System.out.println(dense(n));
        System.out.println(sparse(-2147483648));
        System.out.println(sparse(-100));
        System.out.println(sparse(37));
        System.out.println(sparse(1000000));
        System.out.println(sparse(0));
    }
}
