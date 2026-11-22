// Owned, finite harness for unchanged DS_BlobsMenu._mth0b nearest-neighbor
// image resampling bytecode. Direct method only: no run(), ImageProducer,
// Graphics, worker thread, window, real Applet image or gameplay.
public final class DSBlobsScaleMain extends DS_BlobsMenu {
    private static final int[] SOURCE = {
        0xff010203, 0xff102030, 0x11223344, 0x00000000,
        0x7f00ff80, 0xffadbeef, 0x00abcdef, 0xff332211,
        0xff000000, 0xffffffff, 0xff765432, 0x01234567,
    };

    private void sample(String name, int width, int height) {
        int[] dest = new int[width * height];
        for (int i = 0; i < dest.length; i++) dest[i] = 0x12345678;
        int[] returned = _mth0b(dest, width, height, SOURCE, 4, 3);
        // Deliberately exercise Java's StringBuilder.append(boolean) from
        // unchanged javac bytecode, not an owned string-literal workaround.
        System.out.println(name + "|identity=" + (returned == dest) +
                "|length=" + dest.length);
        for (int y = 0; y < height; y++) {
            String row = name + "|" + y;
            for (int x = 0; x < width; x++) {
                row += "|" + dest[y * width + x];
            }
            System.out.println(row);
        }
    }

    public void init() {
        sample("enlarge", 8, 6);
        sample("shrink", 3, 2);
        sample("reshape", 7, 5);
        sample("identical", 4, 3);
        int check = 0;
        for (int i = 0; i < SOURCE.length; i++) check = check * 31 + SOURCE[i];
        System.out.println("sourceCheck|" + check);
        StringBuffer buffer = new StringBuffer("buffer|");
        StringBuffer same = buffer.append(true);
        same.append(false);
        System.out.println(buffer.toString() + "|identity=" + (same == buffer));
    }
    public static void main(String[] args) { new DSBlobsScaleMain().init(); }
}
