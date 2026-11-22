// Independently checked exact Integer.parseInt(String,int) semantics needed
// by the unchanged DS_BlobsMenu._mth4b body; selected ASCII/int32 only.
public final class IntegerRadixParseMain {
    private static void invalid(String text, int radix) {
        try {
            Integer.parseInt(text, radix);
            System.out.println("BAD accepted");
        } catch (NumberFormatException expected) {
            System.out.println("NumberFormatException");
        }
    }
    public static void main(String[] args) {
        if (args.length > 0) {
            System.out.println(Integer.parseInt(args[0], 16));
            return;
        }
        System.out.println(Integer.parseInt("00ff80", 16));
        System.out.println(Integer.parseInt("+0ABcd", 16));
        System.out.println(Integer.parseInt("-80000000", 16));
        System.out.println(Integer.parseInt("7fffffff", 16));
        System.out.println(Integer.parseInt("101010", 2));
        System.out.println(Integer.parseInt("z", 36));
        invalid(null, 16);
        invalid("", 16);
        invalid("+", 16);
        invalid("0x1f", 16);
        invalid("100000000", 16);
        invalid("-80000001", 16);
        invalid(" 1", 16);
        invalid("1", 1);
        invalid("1", 37);
        invalid("G", 16);
    }
}
