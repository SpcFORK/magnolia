package com.macromedia.flash;

// Independent JDK oracle for the *original*, unmodified FlashApi.ZIP Matrix.
// The two Unicode escapes below are the JVM method identifiers in that class,
// not new algorithms: U+0177 is integer absolute value; U+0178 approximates
// vector magnitude. Calling either also runs Matrix's original <clinit>.
public final class FlashMatrixCoreMain {
    public static void main(String[] args) {
        System.out.println(Matrix.\u0177(-23));
        System.out.println(Matrix.\u0177(-2147483648));
        System.out.println(Matrix.\u0177(2147483647));
        System.out.println(Matrix.\u0178(8, -6));
        System.out.println(Matrix.\u0178(-50, 25));
        System.out.println(Matrix.\u0178(0, 0));
    }
}
