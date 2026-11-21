// Java 8 fixture for an opt-in, bounded 320x240 indexed-pixel work array.
// `javac --release 8 -g:none PixelFrameMain.java` reproduces the .class.
// This tests array sizes/budgets; it is not an AWT image or movie player.
public final class PixelFrameMain {
    public static void main(String[] args) {
        if (args.length != 0) {
            // The JDK accepts these allocations; Magnolia must refuse before
            // allocating more than its explicit cumulative 1 Mi cell budget.
            byte[][] many = new byte[9][];
            for (int i = 0; i < many.length; i++)
                many[i] = new byte[131072];
            System.out.println("JDK-accepted allocation exceeds Magnolia budget");
            return;
        }
        byte[] pixels = new byte[320 * 240];
        pixels[0] = (byte) 0xff;
        pixels[320 * 120 + 160] = 7;
        pixels[pixels.length - 1] = 42;
        System.out.println(pixels.length);
        System.out.println(pixels[0] & 255);
        System.out.println(pixels[320 * 120 + 160] & 255);
        System.out.println(pixels[pixels.length - 1] & 255);
    }
}
