// JDK oracle for Magnolia's category-one control-flow stack-height gate.
// The ternary has two incoming branches that must present one int to istore.
public final class StackHeightMain {
    public static void main(String[] args) {
        int chosen = args.length == 0 ? 4 : 7;
        System.out.println(chosen);
    }
}
