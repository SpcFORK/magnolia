// Owned Java 8 bytecode for a 256-entry indexed palette's snapshot readback.
// Not an arbitrary display/alpha palette or the original Flash DisplayList.
import java.awt.image.IndexColorModel;

public final class FivePaletteMain {
    public static void main(String[] ignored) {
        byte[] red = new byte[256], green = new byte[256], blue = new byte[256];
        red[0] = -1; green[0] = -128; blue[0] = 7;
        red[1] = 33; green[1] = 44; blue[1] = 55;
        red[255] = 9; green[255] = 10; blue[255] = 11;
        IndexColorModel model = new IndexColorModel(8, 256, red, green, blue);
        red[0] = 1; green[0] = 2; blue[0] = 3;
        System.out.println(model.getMapSize());
        byte[] copied = new byte[257];
        copied[256] = 99;
        model.getReds(copied);
        System.out.println(copied[0]); System.out.println(copied[1]);
        System.out.println(copied[255]); System.out.println(copied[256]);
        model.getGreens(copied);
        System.out.println(copied[0]); System.out.println(copied[1]);
        System.out.println(copied[255]); System.out.println(copied[256]);
        model.getBlues(copied);
        System.out.println(copied[0]); System.out.println(copied[1]);
        System.out.println(copied[255]); System.out.println(copied[256]);
        byte[] shortTarget = {88};
        try {
            model.getReds(shortTarget);
            System.out.println("BAD-short");
        } catch (ArrayIndexOutOfBoundsException expected) {
            // JDK copies the prefix before throwing; this is observable.
            System.out.println(shortTarget[0]);
        }
        try {
            model.getGreens(null);
            System.out.println("BAD-null");
        } catch (NullPointerException expected) {
            System.out.println("null-caught");
        }
    }

    // For generated-Java refusal, select the getter itself without the
    // constructor/byte-array operations in the selected method's bytecode.
    public static void paletteSize(IndexColorModel model) {
        System.out.println(model.getMapSize());
    }

    // Valid JDK palette with different bit depth/map size, outside the VM's
    // *existing* 8-bit/256-entry construction/pixel profile.
    public static void otherPalette() {
        byte[] channel = new byte[16];
        IndexColorModel model = new IndexColorModel(4, 16,
            channel, channel, channel);
        System.out.println(model.getMapSize());
    }

    // The model is not a general ColorModel getRGB replacement.
    public static void indexedRGB() {
        byte[] channel = new byte[256];
        IndexColorModel model = new IndexColorModel(8, 256,
            channel, channel, channel);
        System.out.println(model.getRGB(0));
    }
}
