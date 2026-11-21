// JDK -Xverify:all/Xvfb sanity for owned Applet fixtures, never a GUI pixel.
public final class FontMetricsHelveticaJdk {
    public static void main(String[] args) {
        if (args.length != 1) throw new IllegalArgumentException("one case");
        switch (args[0]) {
            case "good": new FontMetricsHelveticaMain().init(); break;
            case "font": new FontMetricsWrongFontMain().init(); break;
            case "control": new FontMetricsControlMain().init(); break;
            case "table": new FontMetricsWidthTableMain().init(); break;
            default: throw new IllegalArgumentException(args[0]);
        }
    }
}
