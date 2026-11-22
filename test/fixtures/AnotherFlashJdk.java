// Independent JDK 11 runner. The VM overwrites -Djava.version at startup,
// so setProperty supplies an explicit reference value only to the JDK fixture.
// Magnolia receives it solely through the virtualJavaVersion host option.
public final class AnotherFlashJdk {
    public static void main(String[] args) throws Exception {
        if (args.length != 1) throw new IllegalArgumentException("one mode required");
        System.setProperty("java.version", "11.0.99");
        switch (args[0]) {
            case "fields": AnotherFlashMain.fields(); break;
            case "inconsistentFields": AnotherFlashMain.inconsistentFields(); break;
            case "missingField": AnotherFlashMain.missingField(); break;
            case "directRead": AnotherFlashMain.directRead(); break;
            case "output": AnotherFlashMain.output(); break;
            case "largeCapacity": AnotherFlashMain.largeCapacity(); break;
            case "tooManyWrites": AnotherFlashMain.tooManyWrites(); break;
            case "nullOutput": AnotherFlashMain.nullOutput(); break;
            case "outputOverload": AnotherFlashMain.outputOverload(); break;
            case "property": AnotherFlashMain.property(); break;
            case "unknownProperty": AnotherFlashMain.unknownProperty(); break;
            case "stringSearch": AnotherFlashMain.stringSearch(); break;
            case "longSearch": AnotherFlashMain.longSearch(); break;
            case "wrongSearchOverload": AnotherFlashMain.wrongSearchOverload(); break;
            case "angles": AnotherFlashMain.angles(); break;
            case "arbitraryAngle": AnotherFlashMain.arbitraryAngle(); break;
            case "infiniteAngle": AnotherFlashMain.infiniteAngle(); break;
            case "otherTrig": AnotherFlashMain.otherTrig(); break;
            default: throw new IllegalArgumentException("unknown mode " + args[0]);
        }
    }
}
