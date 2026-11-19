// VM-library boundary parity: Integer and ASCII String host intrinsics.
public class LibraryMain {
    public static void main(String[] args) {
        System.out.println(Integer.sum(Integer.MAX_VALUE, 1));
        System.out.println(Integer.compare(Integer.MIN_VALUE, 0));
        System.out.println(Integer.max(-4, 6));
        System.out.println(Integer.min(-4, 6));
        System.out.println("Oak".concat("VM").length());
        System.out.println("Oak".equals(args[0]));
        System.out.println(Integer.parseInt("-2147483648"));
        System.out.println(Integer.toString(-13));
    }
}
