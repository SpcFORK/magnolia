// The JDK 11 -Xverify:all reference runner uses the very same owned class
// bytes that Magnolia executes; the java.version substitution is JDK-only.
public final class FlashImageClassJdk {
    public static void main(String[] args) throws Exception {
        switch (args[0]) {
            case "positive": FlashImageClassMain.positive(); break;
            case "archiveBranch":
                // JDK 11 replaces -Djava.version at boot; set it after startup.
                System.setProperty("java.version", "1.0.99");
                FlashImageClassMain.archiveBranch();
                break;
            case "notFound": FlashImageClassMain.notFound(); break;
            case "nullName": FlashImageClassMain.nullName(); break;
            case "nullReceiver": FlashImageClassMain.nullReceiver(); break;
            case "otherLoadedClass": FlashImageClassMain.otherLoadedClass(); break;
            case "otherDecoder": FlashImageClassMain.otherDecoder(); break;
            case "otherOverload": FlashImageClassMain.otherOverload(); break;
            case "otherMember": FlashImageClassMain.otherMember(); break;
            case "tooManyProbes": FlashImageClassMain.tooManyProbes(); break;
            case "longName": FlashImageClassMain.longName(); break;
            default: throw new AssertionError("unknown class probe " + args[0]);
        }
    }
}
