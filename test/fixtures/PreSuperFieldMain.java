// Java 8 verifier fixture. Byte-level tests move the `own` or `inherited`
// putfield ahead of super() without changing the method's code length.
// JVMS allows the current class's directly declared instance field, but
// rejects the inherited field even if its symbolic owner is this subclass.
class PreSuperFieldBase {
    public int inherited;
    public PreSuperFieldBase() { inherited = 3; }
}

public final class PreSuperFieldMain extends PreSuperFieldBase {
    public int own;

    public PreSuperFieldMain(int value) {
        super();
        own = value;
        inherited = value + 1;
    }

    public static void main(String[] args) {
        System.out.println("before");
        PreSuperFieldMain item = new PreSuperFieldMain(7);
        System.out.println(item.own);
        System.out.println(item.inherited);
    }
}
