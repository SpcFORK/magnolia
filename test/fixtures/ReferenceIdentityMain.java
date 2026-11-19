// Reference comparisons are JVM identity, not Oak/JavaScript structural equality.
// Exercises aliases, equal-but-distinct objects/arrays/strings, null, interned
// literals, host singletons and independent (but text-equal) main arguments.
public final class ReferenceIdentityMain {
    int value = 7;

    static String literal() {
        return "equal";
    }

    public static void main(String[] args) {
        ReferenceIdentityMain first = new ReferenceIdentityMain();
        ReferenceIdentityMain second = new ReferenceIdentityMain();
        ReferenceIdentityMain alias = first;
        System.out.println(first == alias);
        System.out.println(first == second);
        System.out.println(first != second);
        System.out.println(first.value == second.value);
        Object nothing = null;
        Object alsoNothing = null;
        System.out.println(nothing == null);
        System.out.println(nothing == alsoNothing);
        System.out.println(first != nothing);

        int[] numbers = {7};
        int[] copy = {7};
        System.out.println(numbers == numbers);
        System.out.println(numbers == copy);
        System.out.println(numbers[0] == copy[0]);

        String shared = literal();
        System.out.println(shared == "equal");
        String built = new StringBuilder("equal").toString();
        System.out.println(shared == built);
        System.out.println(shared.equals(built));
        System.out.println(shared.concat("") == shared);
        System.out.println(shared.concat("!") == "equal!");
        System.out.println(String.valueOf(7) == "7");

        System.out.println(System.out == System.out);
        System.out.println(java.awt.Color.WHITE == java.awt.Color.WHITE);
        System.out.println(java.awt.Color.white == java.awt.Color.WHITE);
        System.out.println(java.awt.Color.WHITE == new java.awt.Color(255, 255, 255));
        System.out.println(args[0] == args[1]);
        System.out.println(args[0].equals(args[1]));
    }
}
