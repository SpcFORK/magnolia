import java.awt.image.BufferedImage;
import java.io.File;
import java.security.MessageDigest;
import javax.imageio.ImageIO;

// Independent JDK 11 PNG decoder for generated *synthetic* filter/bit-depth
// controls only. Original Acotel PNG parity uses real MicroEmulator instead.
public final class IndexedPngJdkReference {
    public static void main(String[] args) {
        try {
            if (args.length != 1 || args[0].indexOf("://") >= 0)
                throw new IllegalArgumentException("exactly one local PNG file");
            BufferedImage image = ImageIO.read(new File(args[0]));
            if (image == null || image.getWidth() < 1 ||
                    image.getWidth() > 512 || image.getHeight() < 1 ||
                    image.getHeight() > 256 ||
                    (long) image.getWidth() * image.getHeight() > 65536)
                throw new AssertionError("out-of-profile PNG");
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            int transparent = 0;
            for (int y = 0; y < image.getHeight(); ++y)
                for (int x = 0; x < image.getWidth(); ++x) {
                    int pixel = image.getRGB(x, y);
                    digest.update((byte)(pixel >>> 24));
                    digest.update((byte)(pixel >>> 16));
                    digest.update((byte)(pixel >>> 8));
                    digest.update((byte)pixel);
                    if ((pixel >>> 24) == 0) ++transparent;
                }
            StringBuilder hex = new StringBuilder();
            for (byte b : digest.digest()) {
                int value = b & 255;
                hex.append(Character.forDigit(value >>> 4, 16));
                hex.append(Character.forDigit(value & 15, 16));
            }
            System.out.println("JDKPNG:" + image.getWidth() + "x" +
                    image.getHeight() + ":sha=" + hex + ":transparent=" + transparent);
        } catch (Throwable error) {
            error.printStackTrace(System.err);
            System.exit(2);
        }
    }
}
