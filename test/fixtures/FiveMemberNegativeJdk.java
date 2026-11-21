// JDK-valid operations deliberately outside the exact five-member VM island.
// Execute the original owned classfile methods; none is an original Flash run.
import java.awt.image.IndexColorModel;

public final class FiveMemberNegativeJdk {
    public static void main(String[] args) {
        if (args.length != 1) throw new IllegalArgumentException("mode");
        switch (args[0]) {
            case "palette-other": FivePaletteMain.otherPalette(); break;
            case "palette-rgb": FivePaletteMain.indexedRGB(); break;
            case "palette-getter": {
                byte[] c = new byte[256];
                FivePaletteMain.paletteSize(new IndexColorModel(8, 256, c, c, c));
                break;
            }
            case "chars-wide": FiveCharsMain.wideChar(); break;
            case "chars-oversized": FiveCharsMain.oversized(); break;
            case "chars-range": FiveCharsMain.rangeConstructor(); break;
            case "clock-other": FiveVirtualClockMain.otherClock(); break;
            default: throw new IllegalArgumentException(args[0]);
        }
    }
}
