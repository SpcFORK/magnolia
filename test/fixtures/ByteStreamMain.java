// Java 8/JDK 11 oracle for the deliberately small, in-memory stream island.
// No disk, URL, ZIP, skip, mark/reset, single-byte read, or full java.io API.
import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.InputStream;
import java.io.IOException;

public final class ByteStreamMain {
    public static void main(String[] args) throws IOException {
        if (args.length > 0 && args[0].equals("budget")) {
            byte[] input = new byte[1024];
            byte[] output = new byte[1024];
            DataInputStream limited = new DataInputStream(new ByteArrayInputStream(input));
            System.out.println(limited.read(output, 0, 1024));
            return;
        }
        byte[] input = {0, (byte) 255, 3, 4};
        ByteArrayInputStream raw = new ByteArrayInputStream(input);
        InputStream parent = raw;
        DataInputStream data = new DataInputStream(parent);
        byte[] output = {99, 99, 99, 99, 99, 99};
        System.out.println(raw instanceof InputStream);
        System.out.println(data instanceof InputStream);
        System.out.println(data instanceof DataInputStream);
        System.out.println(raw == parent);
        System.out.println(data.read(output, 1, 0));
        System.out.println(data.read(output, 1, 2));
        System.out.println(output[0]);
        System.out.println(output[1]);
        System.out.println(output[2]);
        input[2] = 88; // ByteArrayInputStream keeps the original array reference.
        System.out.println(data.read(output, 3, 3)); // partial read: only 2 remain
        System.out.println(output[3]);
        System.out.println(output[4]);
        System.out.println(output[5]);
        System.out.println(data.read(output, 0, 1)); // EOF
        System.out.println(data.read(output, 0, 0)); // EOF takes precedence over zero length
        try {
            data.read(output, -1, 0); // bad range still throws at EOF
            System.out.println("range was accepted");
        } catch (IndexOutOfBoundsException expected) {
            System.out.println("range rejected");
        }
        try {
            data.read(null, 0, 0); // null still throws at EOF
            System.out.println("null was accepted");
        } catch (NullPointerException expected) {
            System.out.println("null rejected");
        }
        DataInputStream empty = new DataInputStream(new ByteArrayInputStream(new byte[0]));
        System.out.println(empty.read(new byte[2], 0, 0));
        DataInputStream second = new DataInputStream(new ByteArrayInputStream(input));
        System.out.println(second.read(output, 0, 4));
        System.out.println(output[2]);
        DataInputStream nullSource = new DataInputStream((InputStream) null);
        try {
            nullSource.read(output, 0, 1);
            System.out.println("missing source was accepted");
        } catch (NullPointerException expected) {
            System.out.println("missing source rejected");
        }
        try {
            new ByteArrayInputStream((byte[]) null);
            System.out.println("missing input was accepted");
        } catch (NullPointerException expected) {
            System.out.println("missing input rejected");
        }
    }
}
