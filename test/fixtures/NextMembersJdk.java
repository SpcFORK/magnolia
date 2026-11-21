// Independent JDK 11 -Xverify:all invocations of the exact owned class.
public final class NextMembersJdk {
    public static void main(String[] args) throws Exception {
        if (args.length != 1) throw new IllegalArgumentException("mode");
        switch (args[0]) {
            case "rectangle": NextMembersMain.rectangleOnly(); break;
            case "notify": NextMembersMain.notifyOnly(); break;
            case "yield": NextMembersMain.yieldOnly(); break;
            case "yield-bound": NextMembersMain.yieldOverBound(); break;
            case "notify-bound": NextMembersMain.notifyOverBound(); break;
            case "replace-rectangle": NextMembersMain.replaceRectangle(); break;
            case "mutate-rectangle": NextMembersMain.mutateRectangle(); break;
            case "notify-one": NextMembersMain.notifyOne(); break;
            case "real-wait": NextMembersMain.realWait(); break;
            case "sleep-without-task": NextMembersMain.sleepWithoutTask(); break;
            default: throw new IllegalArgumentException(args[0]);
        }
    }
}
