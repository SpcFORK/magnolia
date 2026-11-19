// Supported primitive/reference array opcodes must match Java 8 behavior.
public class ValidArraysMain {
    public static void main(String[] args) {
        byte[] bytes = new byte[1];
        bytes[0] = (byte) 200;
        System.out.println(bytes[0]);
        boolean[] flags = new boolean[1];
        flags[0] = true;
        System.out.println(flags[0]);
        int[] numbers = new int[2];
        numbers[0] = 7;
        System.out.println(numbers.length);
        System.out.println(numbers[0]);
        String[] words = new String[1];
        words[0] = "ok";
        System.out.println(words[0]);
        Object[] refs = words;
        System.out.println(refs instanceof Object[]);
    }
}
