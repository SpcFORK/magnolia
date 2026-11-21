// Owned Java 8 fixture: exact member families seen in the original FlashApi.ZIP.
// Not a Flash movie, audio backend, display peer, or proprietary bytecode.
import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.DataOutputStream;
import java.io.InputStream;

public final class AnotherFlashMain extends ByteArrayInputStream {
    private AnotherFlashMain(byte[] bytes) { super(bytes); }

    public static void fields() throws Exception {
        byte[] original = {1, -2, 3, 4};
        new AnotherFlashMain(original).fieldProbe(original);
    }
    private void fieldProbe(byte[] original) throws Exception {
        System.out.println(super.count);
        System.out.println(super.pos);
        System.out.println(super.buf[1] & 255);
        byte[] substitute = {-1, 7, 9, 11};
        super.buf = substitute;
        super.pos = 1;
        super.count = 3;
        InputStream inherited = this; // symbolic owner InputStream.read([B)
        byte[] destination = new byte[4];
        System.out.println(inherited.read(destination));
        System.out.println(destination[0] & 255);
        System.out.println(destination[1] & 255);
        System.out.println(super.pos);
        substitute[1] = 42; // buf remains an alias, never a snapshot
        super.pos = 1;
        super.count = 2;
        System.out.println(inherited.read(destination));
        System.out.println(destination[0] & 255);
        super.count = Integer.MAX_VALUE; // SoundMixer does this; fields alone permit it
        System.out.println(super.count);
        System.out.println(original[0]);
    }
    public static void inconsistentFields() throws Exception {
        new AnotherFlashMain(new byte[] {1, 2}).inconsistentFieldRead();
    }
    private void inconsistentFieldRead() throws Exception {
        super.count = Integer.MAX_VALUE;
        InputStream inherited = this;
        inherited.read(new byte[2]); // JDK succeeds here; bounded VM refuses unsafe count
    }
    public static void missingField() {
        new AnotherFlashMain(new byte[] {1}).readMark();
    }
    private void readMark() {
        System.out.println(super.mark); // JDK-valid, deliberately not admitted
    }
    public static void directRead() throws Exception {
        ByteArrayInputStream input = new ByteArrayInputStream(new byte[] {1, 2});
        System.out.println(input.read()); // unadmitted primitive overload
    }

    public static void output() throws Exception {
        ByteArrayOutputStream memory = new ByteArrayOutputStream(0);
        DataOutputStream writer = new DataOutputStream(memory);
        int[] header = {779316836, 0, -1, 1, 8000, 1, 0};
        for (int word : header) writer.writeInt(word);
        byte[] first = memory.toByteArray();
        byte[] second = memory.toByteArray();
        System.out.println(first.length);
        System.out.println(first[0] & 255);
        System.out.println(first[1] & 255);
        System.out.println(first[2] & 255);
        System.out.println(first[3] & 255);
        System.out.println(first[11] & 255);
        System.out.println(first[19] & 255);
        first[0] = 0;
        System.out.println(second[0] & 255); // independent toByteArray snapshot
        writer.writeInt(0x01020304);
        System.out.println(second.length);
        System.out.println(memory.toByteArray().length);
    }
    public static void largeCapacity() {
        new ByteArrayOutputStream(65537);
        System.out.println("jdk-accepted-capacity");
    }
    public static void tooManyWrites() throws Exception {
        ByteArrayOutputStream memory = new ByteArrayOutputStream(0);
        DataOutputStream writer = new DataOutputStream(memory);
        for (int i = 0; i < 16385; ++i) writer.writeInt(i);
        System.out.println(memory.toByteArray().length);
    }
    public static void nullOutput() throws Exception {
        DataOutputStream writer = new DataOutputStream(null);
        writer.writeInt(7); // JDK throws NPE; VM refuses unmodeled target
    }
    public static void outputOverload() throws Exception {
        DataOutputStream writer = new DataOutputStream(new ByteArrayOutputStream(0));
        writer.writeShort(42); // no arbitrary DataOutputStream writer surface
    }

    public static void property() {
        String version = System.getProperty("java.version");
        System.out.println(version);
        System.out.println(version == System.getProperty("java.version"));
        System.out.println("SLAV".concat(version).indexOf("SLAV1.0"));
    }
    public static void unknownProperty() {
        System.out.println(System.getProperty("user.home") != null);
    }
    public static void stringSearch() {
        System.out.println("SLAV1.0".indexOf("SLAV1.0"));
        System.out.println("SLAV11.0.99".indexOf("SLAV1.0"));
        System.out.println("ABABA".indexOf("BAB"));
        System.out.println("".indexOf(""));
        System.out.println("abc".indexOf(""));
        System.out.println("abcd".indexOf("xyz"));
        try { "abc".indexOf((String) null); }
        catch (NullPointerException e) { System.out.println("null-caught"); }
    }
    public static void longSearch() {
        System.out.println("short".indexOf("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"));
    }
    public static void wrongSearchOverload() {
        System.out.println("abc".indexOf("b", 1));
    }

    public static void angles() {
        double[] ys = {0.0, -0.0, 0.0, -0.0, 1.0, -1.0, 1.0, -1.0, 2.0, -2.0, 1.0, -1.0};
        double[] xs = {0.0, 0.0, -0.0, -0.0, 0.0, -0.0, 1.0, 1.0, -2.0, -2.0, -1.0, -1.0};
        for (int i = 0; i < ys.length; ++i)
            System.out.println(Double.doubleToRawLongBits(Math.atan2(ys[i], xs[i])));
    }
    public static void arbitraryAngle() {
        System.out.println(Double.doubleToRawLongBits(Math.atan2(1.0, 2.0)));
    }
    public static void infiniteAngle() {
        System.out.println(Double.doubleToRawLongBits(Math.atan2(Double.POSITIVE_INFINITY, 2.0)));
    }
    public static void otherTrig() {
        System.out.println(Math.sin(0.25)); // no general host libm
    }
}
