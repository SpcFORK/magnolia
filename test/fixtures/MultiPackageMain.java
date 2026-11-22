// Explicitly staged sources share a basename but live in different packages.
public class MultiPackageMain {
    public static void main(String[] args) {
        System.out.println(a.Peer.value() + b.Peer.value());
    }
}
