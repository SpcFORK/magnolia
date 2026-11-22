import javax.microedition.midlet.MIDlet;
import org.microemu.MIDletAccess;
import org.microemu.MIDletBridge;
import org.microemu.app.Headless;

// JDK-only independent reference: call actual MicroEmulator MIDletAccess
// callbacks (not our owned model). For this owned fixture only, from a
// bounded process with a private user.home and in-memory RMS.
public final class ReferenceAMS {
    public static void main(String[] args) throws Exception {
        String name = args.length == 0 ? "FiniteLifecycleMIDlet" : args[0];
        if (!"FiniteLifecycleMIDlet".equals(name) &&
            !"UnsupportedPauseMIDlet".equals(name) &&
            !"NotifiedPauseMIDlet".equals(name))
            throw new IllegalArgumentException("unapproved owned fixture");
        Headless.main(new String[]{"--rms", "memory", name});
        MIDlet midlet = MIDletBridge.getCurrentMIDlet();
        // The emulator intentionally loads application bytes in its own
        // classloader, so JVM instanceof against this harness is not valid.
        if (midlet == null || !name.equals(midlet.getClass().getName()))
            throw new AssertionError("reference did not create the fixture MIDlet");
        MIDletAccess access = MIDletBridge.getMIDletAccess(midlet);
        if (access == null) throw new AssertionError("missing reference AMS access");
        access.pauseApp();
        access.startApp();
        access.destroyApp(true);
        System.out.println("REFERENCE:completed");
        // MicroEmulator's headless scheduler may retain background threads.
        // This oracle owns a short-lived dedicated process, never a server.
        System.exit(0);
    }
}
