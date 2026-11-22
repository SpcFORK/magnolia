// Owned Java 8 bytecode for exact InputStream.read(byte[]) on a local
// file: URL and a mutable ByteArrayInputStream. This is not Flash playback.
import java.applet.Applet;
import java.io.ByteArrayInputStream;
import java.io.DataInputStream;
import java.io.IOException;
import java.io.InputStream;
import java.net.URL;

public final class BulkStreamApplet extends Applet {
    private void readFile(String label, URL base) throws IOException {
        InputStream stream = new URL(base, "movie.bin").openStream();
        byte[] target = {11, 11, 11, 11};
        System.out.println(label);
        System.out.println(stream.read(new byte[0])); // FileInputStream: zero before EOF.
        System.out.println(stream.read(target));
        System.out.println(target[0]);
        System.out.println(target[1]);
        System.out.println(target[2]);
        System.out.println(target[3]); // Unwritten sentinel must remain.
        System.out.println(stream.read(new byte[0])); // FileInputStream: zero at EOF.
        System.out.println(stream.read(target)); // -1 at EOF on a nonempty target.
        try {
            stream.read(null);
            System.out.println("BAD-null");
        } catch (NullPointerException expected) {
            System.out.println("null-caught");
        }
        stream.close();
    }

    private void readArray() throws IOException {
        byte[] source = {7, -2, 5, 6};
        InputStream stream = new ByteArrayInputStream(source);
        byte[] target = {9, 9};
        System.out.println("array");
        System.out.println(stream.read(new byte[0])); // Before EOF: zero.
        System.out.println(stream.read(target));
        System.out.println(target[0]);
        System.out.println(target[1]);
        source[2] = -128; // The underlying array was not copied.
        System.out.println(stream.read(target));
        System.out.println(target[0]);
        System.out.println(target[1]);
        System.out.println(stream.read(new byte[0])); // At EOF: -1 here.
        System.out.println(stream.read(target));
    }

    public void probe() throws IOException {
        readFile("document", ((Applet) this).getDocumentBase());
        readFile("code", ((Applet) this).getCodeBase());
        readArray();
    }

    // JDK-valid resource wrapper: check the older DataInputStream read
    // overload's zero-length EOF rule on a file: source separately.
    public void wrappedZero() throws IOException {
        InputStream raw = new URL(((Applet) this).getDocumentBase(),
                                  "movie.bin").openStream();
        DataInputStream stream = new DataInputStream(raw);
        byte[] target = new byte[4];
        System.out.println(stream.read(target, 0, 4));
        System.out.println(stream.read(target, 0, 0));
        raw.close();
    }

    // JDK System.arraycopy handles source/destination overlap in this
    // ByteArrayInputStream implementation. The bounded host must also
    // snapshot these few bytes before writing into the same source array.
    public void overlap() throws IOException {
        byte[] source = {1, 2, 3, 4};
        InputStream stream = new ByteArrayInputStream(source);
        stream.read(new byte[1]);
        System.out.println(stream.read(source));
        System.out.println(source[0]);
        System.out.println(source[1]);
        System.out.println(source[2]);
        System.out.println(source[3]);
    }

    // Selected under a too-small VM step budget: no output/cursor mutation
    // may occur if the complete copy cannot be metered before it starts.
    public void budgetRead() throws IOException {
        byte[] source = new byte[256];
        source[0] = 1;
        source[255] = 4;
        InputStream stream = new ByteArrayInputStream(source);
        byte[] target = new byte[256];
        System.out.println(stream.read(target));
    }

    public void closedResource() throws IOException {
        InputStream stream = new URL(((Applet) this).getCodeBase(),
                                     "movie.bin").openStream();
        stream.close();
        try {
            stream.read(new byte[1]);
            System.out.println("BAD-closed");
        } catch (IOException expected) {
            System.out.println("closed-caught");
        }
    }

    // A DataInputStream is an InputStream too, but this deliberately narrow
    // host contract does not substitute its separate one-array override.
    public void unmodeledSubtype() throws IOException {
        InputStream stream = new DataInputStream(
            new ByteArrayInputStream(new byte[] {3}));
        System.out.println(stream.read(new byte[1]));
    }

    // Legal on the JDK, but the wider InputStream.read(byte[],int,int)
    // symbolic member remains outside the narrow one-array contract.
    public void unsupportedOverload() throws IOException {
        InputStream stream = new ByteArrayInputStream(new byte[] {1});
        System.out.println(stream.read(new byte[1], 0, 1));
    }

    // javac emits a different constant-pool owner here, although the JDK
    // behavior is identical. Do not admit it by method name alone.
    public void unsupportedOwner() throws IOException {
        ByteArrayInputStream stream = new ByteArrayInputStream(new byte[] {1});
        System.out.println(stream.read(new byte[1]));
    }
}
