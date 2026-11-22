// JDK-valid, resource-policy-negative case: Integer.valueOf accepts
// thousands of leading zeroes, but Magnolia caps this parse at 4096 ASCII
// characters. No output is allowed before the Magnolia policy refusal.
public final class IntegerBoxBoundsMain {
    public static void main(String[] args) {
        Integer.valueOf(args[0]);
        System.out.println("accepted");
    }
}
