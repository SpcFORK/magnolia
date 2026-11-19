// Deliberately tiny reference harness; the Mojang client JAR is provided at
// test time, outside this repository. This does NOT run the client entrypoint.
public final class BrandOracle {
    public static void main(String[] args) {
        System.out.println(net.minecraft.client.ClientBrandRetriever.getClientModName());
    }
}
