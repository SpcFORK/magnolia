// Isolated javac must receive SourceUtil.java explicitly; it cannot search the workspace.
public class SourceMain {
    public static void main(String[] args) {
        System.out.println(SourceUtil.answer());
    }
}
