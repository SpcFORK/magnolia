// JDK oracle: String.valueOf(Object) returns its String argument by identity,
// and returns the interned literal "null" for a null Object reference.
public final class StringObjectValueOfMain {
    public static void main(String[] args) {
        String word = args.length == 0 ? "Mint" : args[0];
        String same = String.valueOf((Object) word);
        System.out.println(same == word);
        System.out.println(same);
        String absent = String.valueOf((Object) null);
        System.out.println(absent == "null");
        System.out.println(absent);
        System.out.println(String.valueOf((Object) word) == same);
    }
}
