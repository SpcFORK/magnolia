// ASCII String VM-library calls also found in historical Java Flash classes.
public class StringCaseMain {
    public static void main(String[] args) {
        System.out.println("Low".equalsIgnoreCase("lOW"));
        System.out.println("HIGH".equalsIgnoreCase(args[0]));
        System.out.println("Hello".indexOf('l'));
        System.out.println("Hello".indexOf('z'));
        System.out.println(String.valueOf(-13));
    }
}
