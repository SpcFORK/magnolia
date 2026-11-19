import java.awt.Point;

// JDK oracle for the bounded System.arraycopy VM-library transition.
// Exercises primitive types, overlap in both directions, reference identity,
// reference covariance, null elements, and a valid end-position zero copy.
// Optional arguments deliberately trigger rejected operations without relying
// on exception handlers the Magnolia JVM profile has not implemented.
public final class ArrayCopyMain {
    public static void main(String[] args) {
        int[] ints = {1, 2, 3, 4, 5};
        System.arraycopy(ints, 0, ints, 1, 4);
        for (int i = 0; i < ints.length; i++) System.out.println(ints[i]);
        System.arraycopy(ints, 1, ints, 0, 4);
        System.out.println(ints[0]);
        System.out.println(ints[3]);
        System.out.println(ints[4]);

        byte[] signed = {-128, 0, 127};
        byte[] signedCopy = new byte[3];
        System.arraycopy(signed, 0, signedCopy, 0, 3);
        System.out.println(signedCopy[0]);
        System.out.println(signedCopy[2]);
        char[] chars = {'A', '\u00ff'};
        char[] charCopy = new char[2];
        System.arraycopy(chars, 0, charCopy, 0, 2);
        System.out.println((int) charCopy[0]);
        System.out.println((int) charCopy[1]);
        short[] shorts = {-32768, 32767};
        short[] shortCopy = new short[2];
        System.arraycopy(shorts, 0, shortCopy, 0, 2);
        System.out.println(shortCopy[0]);
        System.out.println(shortCopy[1]);
        boolean[] flags = {true, false};
        boolean[] flagCopy = new boolean[2];
        System.arraycopy(flags, 0, flagCopy, 0, 2);
        System.out.println(flagCopy[0]);
        System.out.println(flagCopy[1]);

        Object[] source = {"Oak", "JS", null};
        String[] strings = new String[3];
        System.arraycopy(source, 0, strings, 0, 3);
        System.out.println(strings[0] == source[0]);
        System.out.println(strings[0].equals("Oak"));
        System.out.println(strings[2] == null);
        Object[] back = new Object[3];
        System.arraycopy(strings, 0, back, 0, 3);
        System.out.println(back[1] == source[1]);
        Point[] points = {new Point(1, 2), null};
        Object[] objects = new Object[2];
        System.arraycopy(points, 0, objects, 0, 2);
        System.out.println(objects[0] == points[0]);
        System.out.println(objects[1] == null);
        System.arraycopy(ints, ints.length, ints, ints.length, 0);
        System.out.println(ints[4]);

        if (args.length != 0) {
            int[] sample = {1, 2};
            if (args[0].equals("range")) System.arraycopy(sample, 1, sample, 0, 2);
            if (args[0].equals("negative")) System.arraycopy(sample, 0, sample, 0, -1);
            if (args[0].equals("null")) System.arraycopy(null, 0, sample, 0, 0);
            if (args[0].equals("type")) System.arraycopy(sample, 0, new byte[2], 0, 0);
            if (args[0].equals("store")) {
                Object[] mixed = {"ok", new Point(1, 2)};
                String[] onlyStrings = new String[2];
                System.arraycopy(mixed, 0, onlyStrings, 0, 2);
            }
        }
    }
}
