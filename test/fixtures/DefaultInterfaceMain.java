// Explicit negative fixture: JVM default interface methods are not yet emulated.
interface DefaultSpeaker {
    default int value() { return 4; }
}
public class DefaultInterfaceMain implements DefaultSpeaker {
    public static void main(String[] args) {
        DefaultSpeaker speaker = new DefaultInterfaceMain();
        System.out.println(speaker.value());
    }
}
