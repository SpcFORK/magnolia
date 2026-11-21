// Owned fixture: exception identity only. It does not construct or fetch URLs.
import java.io.IOException;
import java.net.MalformedURLException;

public final class MalformedUrlExceptionMain {
    static void checked() throws IOException {
        throw new MalformedURLException("bad address");
    }
    static void generic() throws IOException {
        throw new IOException("ordinary");
    }
    public static void main(String[] args) {
        try {
            checked();
        } catch (MalformedURLException x) {
            System.out.println("specific:" + x.getMessage());
            System.out.println(x instanceof IOException);
            System.out.println(x instanceof Exception);
        } catch (IOException x) {
            System.out.println("wrong handler");
        }
        try {
            checked();
        } catch (IOException x) {
            System.out.println("parent:" + x.getMessage());
        }
        try {
            generic();
        } catch (MalformedURLException x) {
            System.out.println("wrong subtype");
        } catch (IOException x) {
            System.out.println("sibling:" + x.getMessage());
        }
        System.out.println(new MalformedURLException().getMessage() == null);
    }
}
