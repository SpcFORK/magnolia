// Independent JDK oracle for the *original*, unmodified JVM 45.3 fliplay.class.
// Build with that class on javac's classpath; Magnolia takes the resulting
// harness .class and original fliplay.class, not a Java reimplementation.
public final class FliplayCoreMain {
    public static void main(String[] args) {
        fliplay player = new fliplay();
        byte[] bytes = {(byte) 0xff, (byte) 0x80, 0x34, 0x12};
        System.out.println(player.get_ubyte(bytes, 0));
        System.out.println(player.get_ubyte(bytes, 1));
        System.out.println(player.get_ushort(bytes, 0));
        System.out.println(player.get_ushort(bytes, 1));
        System.out.println(player.get_ulong(bytes, 0));
        System.out.println(player.get_ulong(new byte[] {-1, -1, -1, -1}, 0));
    }
}
