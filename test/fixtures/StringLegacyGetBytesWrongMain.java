// This distinct overload remains unsupported and must fail preflight.
public final class StringLegacyGetBytesWrongMain {
    public static void main(String[] args) {
        System.out.println("prefix-must-not-appear");
        byte[] data = "ASCII".getBytes();
        System.out.println(data.length);
    }
}
