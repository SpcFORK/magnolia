// Interface call-site linking, subclass overrides and interface casts.
interface Named {
    int value();
}
interface Tuned extends Named { }
class BaseValue implements Tuned {
    public int value() { return 2; }
}
class ChangedValue extends BaseValue {
    @Override public int value() { return 7; }
}
public class InterfaceMain {
    static int score(Named receiver) { return receiver.value() + 1; }
    public static void main(String[] args) {
        Named a = new BaseValue();
        Tuned b = new ChangedValue();
        System.out.println(score(a));
        System.out.println(score(b));
        System.out.println(b.value());
        System.out.println(a instanceof Named);
        System.out.println(b instanceof Tuned);
        Object raw = b;
        System.out.println(((Named) raw).value());
    }
}
