// JDK-valid counterexample. append(char) is still outside Magnolia's
// deliberately bounded append(int/String/boolean) profile.
public final class StringBuilderUnsupportedChar {
    public static void main(String[] args) {
        System.out.println(new StringBuilder("c").append('Z').toString());
    }
}
