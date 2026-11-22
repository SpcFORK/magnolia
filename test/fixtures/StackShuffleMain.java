// JVM stack-shuffle forms emitted by javac for field and array increments.
public final class StackShuffleMain {
    int value = 5;
    int[] ints = new int[]{10, 20};
    int counter = 0;

    int exercise() {
        int beforeField = this.value++;
        int afterField = ++this.value;
        int at = counter++;
        int beforeElement = this.ints[at]++;
        int afterElement = ++this.ints[at];
        System.out.println(beforeField);
        System.out.println(afterField);
        System.out.println(beforeElement);
        System.out.println(afterElement);
        return this.value + this.ints[at];
    }

    public static void main(String[] args) {
        System.out.println(new StackShuffleMain().exercise());
    }
}
