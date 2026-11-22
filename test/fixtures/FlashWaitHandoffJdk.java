// Independent JDK 11 -Xverify:all harness. Only FlashWaitHandoffMain.class is
// sent to Magnolia; join/interruption here are reference-side test controls.
public final class FlashWaitHandoffJdk {
    private static void joined() throws Exception {
        FlashWaitHandoffMain.worker.join(2000L);
        if (FlashWaitHandoffMain.worker.isAlive())
            throw new AssertionError("bounded JDK fixture did not terminate");
    }

    public static void main(String[] args) throws Exception {
        switch (args[0]) {
            case "positive":
                FlashWaitHandoffMain.positive();
                joined();
                break;
            case "notOwned":
                FlashWaitHandoffMain.notOwned();
                break;
            case "nullWait":
                FlashWaitHandoffMain.nullWait();
                break;
            case "timedWait":
                FlashWaitHandoffMain.timedWait();
                break;
            case "wrongNotifier":
            case "noNotifier":
            case "priorNotification":
                if (args[0].equals("wrongNotifier"))
                    FlashWaitHandoffMain.wrongNotifier();
                else if (args[0].equals("priorNotification"))
                    FlashWaitHandoffMain.priorNotification();
                else FlashWaitHandoffMain.noNotifier();
                FlashWaitHandoffMain.worker.join(50L);
                if (!FlashWaitHandoffMain.worker.isAlive())
                    throw new AssertionError("JDK wait returned without notice");
                FlashWaitHandoffMain.worker.interrupt();
                joined();
                break;
            case "twoNotifiers":
                FlashWaitHandoffMain.twoNotifiers();
                joined();
                break;
            case "noMonitorStartInTask":
                FlashWaitHandoffMain.noMonitorStartInTask();
                joined();
                break;
            default:
                throw new AssertionError("unknown test mode");
        }
    }
}
