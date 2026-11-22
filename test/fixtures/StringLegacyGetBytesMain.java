// Owned Java 8 fixture for deprecated String.getBytes(II[BI)V.
// Exact ASCII byte copying, JDK 11 range ordering, null, and atomic failures.
public final class StringLegacyGetBytesMain {
    private static void trial(String name, int start, int end, byte[] dst, int at) {
        System.out.println(name);
        try {
            "ABCDE".getBytes(start, end, dst, at);
            System.out.println("ok");
        } catch (StringIndexOutOfBoundsException e) {
            System.out.println("bounds");
        } catch (NullPointerException e) {
            System.out.println("null");
        }
        if (dst == null) {
            System.out.println("no-array");
        } else {
            for (int i = 0; i < dst.length; i++) System.out.println(dst[i]);
        }
    }
    public static void main(String[] args) {
        trial("copy", 1, 4, new byte[]{9, 9, 9, 9, 9}, 1);
        trial("empty-end", 5, 5, new byte[0], 0);
        trial("source-negative-null", -1, 6, null, -1);
        trial("source-too-long-null", 0, 6, null, 0);
        trial("source-reversed-null", 2, 1, null, 0);
        trial("null-destination", 0, 1, null, 0);
        trial("null-zero", 0, 0, null, 0);
        trial("dest-negative", 0, 1, new byte[]{7, 7}, -1);
        trial("dest-short", 0, 5, new byte[]{8, 8}, 0);
        trial("dest-past-end-zero", 5, 5, new byte[]{6, 6}, 3);
        trial("source-past-end-zero", 9, 9, new byte[0], 0);
        trial("valid-zero", 1, 1, new byte[0], 0);
    }
}
