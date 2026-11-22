// Owned Java 8 oracle for exactly the Class.forName(String) probe in the
// SHA-pinned original Flash ScriptPlayer. This is not a general class loader.
public final class FlashImageClassMain {
    private static final String DECODER = "sun.awt.image.ImageDecoder";

    public static void positive() throws ClassNotFoundException {
        Class<?> first = Class.forName(DECODER);
        // Same value from a distinct String object: the lookup is not an
        // accidental comparison of the caller's String reference identity.
        Class<?> again = Class.forName(new String(DECODER));
        System.out.println(first.getName());
        System.out.println(first == again);
        System.out.println(first != null);
        System.out.println(first.getName() == again.getName());
        System.out.println(first.getName() == DECODER);
    }

    // The archive itself tests "SLAV" + System.getProperty("java.version")
    // against "SLAV1.0" before loading this exact ImageDecoder class.
    public static void archiveBranch() throws ClassNotFoundException {
        String version = System.getProperty("java.version");
        if (("SLAV" + version).indexOf("SLAV1.0") >= 0) {
            Class<?> decoder = Class.forName(DECODER);
            System.out.println("reference-branch");
            System.out.println(decoder.getName());
        } else {
            System.out.println("different-version");
        }
    }

    public static void notFound() {
        try {
            Class.forName("sun.audio.AudioPlayer");
            System.out.println("BUG-audio-found");
        } catch (ClassNotFoundException expected) {
            System.out.println(expected.getMessage());
        }
    }

    public static void nullName() {
        try {
            Class.forName(null);
            System.out.println("BUG-null-lookup");
        } catch (NullPointerException expected) {
            System.out.println("null-name");
        } catch (ClassNotFoundException unexpected) {
            System.out.println("BUG-null-class-not-found");
        }
    }

    public static void nullReceiver() throws ClassNotFoundException {
        Class<?> none = null;
        try {
            System.out.println(none.getName());
        } catch (NullPointerException expected) {
            System.out.println("null-receiver");
        }
    }

    // These are valid, *observed* on JDK 11, and out of the VM's exact
    // archive-derived lookup scope. Fail closed, never invent a classpath.
    public static void otherLoadedClass() throws ClassNotFoundException {
        System.out.println(Class.forName("java.lang.String").getName());
    }
    public static void otherDecoder() throws ClassNotFoundException {
        System.out.println(Class.forName("sun.awt.image.JPEGImageDecoder").getName());
    }
    public static void otherOverload() throws ClassNotFoundException {
        System.out.println(Class.forName(DECODER, false, null).getName());
    }
    public static void otherMember() throws ClassNotFoundException {
        System.out.println(Class.forName(DECODER).getSuperclass().getName());
    }
    public static void tooManyProbes() throws ClassNotFoundException {
        for (int i = 0; i < 17; ++i) Class.forName(DECODER);
        System.out.println("jdk-seventeen");
    }
    public static void longName() {
        try {
            Class.forName("not.present.this.is.a.deliberately.very.long.class.name.that.is.over.sixty.four.bytes");
            System.out.println("BUG-long-found");
        } catch (ClassNotFoundException expected) {
            System.out.println("jdk-long-missing");
        }
    }
}
