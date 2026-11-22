import java.awt.Event;
import java.lang.reflect.Field;

/** JDK 11 reference using a no-display Applet instance, never a browser UI. */
@SuppressWarnings("deprecation")
public final class LegacyEventJdk {
    private static LegacyEventMain noWindow() throws Exception {
        Class<?> clazz = Class.forName("sun.misc.Unsafe");
        Field field = clazz.getDeclaredField("theUnsafe");
        field.setAccessible(true);
        Object unsafe = field.get(null);
        return (LegacyEventMain) clazz
            .getMethod("allocateInstance", Class.class)
            .invoke(unsafe, LegacyEventMain.class);
    }
    public static void main(String[] args) throws Exception {
        LegacyEventMain app = noWindow();
        int[][] records = {
            {401, 65, 0, 0}, {403, 66, 0, 0},
            {402, 70, 0, 0}, {501, 0, 9, 7},
            {506, 0, 10, 8}, {504, 0, 5, 6},
            {1001, 0, 0, 0}, {777, 0, 0, 0}, {2032, 0, 0, 0}
        };
        for (int[] r : records) {
            Event e = new Event(app, r[0], null);
            e.key = r[1]; e.x = r[2]; e.y = r[3];
            boolean handled = app.handleEvent(e);
            System.out.println("EVENT " + r[0] + " " + (handled ? 1 : 0) +
                " " + app.hits + " " + app.checksum);
        }
    }
}
