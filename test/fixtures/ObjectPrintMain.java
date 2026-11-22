// Supported Object printing is limited to null, String and text builders.
// User-defined toString() must not be replaced with a host object's debug dump.
public class ObjectPrintMain {
    static class Label {
        public String toString() { return "label"; }
    }

    public static void main(String[] args) {
        Object word = "word";
        System.out.println(word);
        Object absent = null;
        System.out.println(absent);
        Object builder = new StringBuffer("buffer");
        System.out.println(builder);
        Object custom = new Label();
        System.out.println(custom);
    }
}
