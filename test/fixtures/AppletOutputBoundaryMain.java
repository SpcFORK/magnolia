// Owned control: a real Java guest exception must preserve earlier JVM
// println output. The host's own resource-budget failures, unlike this
// Java throw, must refuse without publishing a partial success trace.
public final class AppletOutputBoundaryMain {
    public void init() {
        System.out.println("guest-prefix");
        throw new IllegalArgumentException("guest-exception");
    }
}
