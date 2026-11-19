// Java 8 verifier oracle for reference assignability at real call sites.
// Reproduce: javac --release 8 -g:none RefSubtypeMain.java
interface RefScore { int score(); }
class RefParent implements RefScore {
    final int value;
    RefParent(int value) { this.value = value; }
    public int score() { return value; }
}
class RefChild extends RefParent {
    RefChild(int value) { super(value); }
}
class RefSibling extends RefParent {
    RefSibling(int value) { super(value); }
}
class RefOther {
    final int value;
    RefOther(int value) { this.value = value; }
}
public final class RefSubtypeMain {
    static RefParent saved;
    static void stash(RefParent parent, RefOther other) { saved = parent; }
    static RefParent choose(RefParent parent, RefOther other) { return parent; }
    static RefParent merge(boolean first, RefChild child, RefSibling sibling) {
        return first ? child : sibling;
    }
    static int fromParent(RefParent value) {
        return value == null ? -1 : value.score();
    }
    static int fromScore(RefScore value) { return value.score(); }
    static int fromArray(RefParent[] values) { return values.length; }
    static int mergeArrays(boolean first, RefChild[] children, RefSibling[] siblings) {
        return fromArray(first ? children : siblings);
    }
    static int fromObject(Object value) {
        return value instanceof RefScore ? ((RefScore) value).score() : -7;
    }
    static int join(boolean first, RefChild child, RefOther other) {
        return fromObject(first ? child : other);
    }
    public static void main(String[] args) {
        RefParent parent = new RefParent(4);
        RefChild child = new RefChild(9);
        RefOther other = new RefOther(12);
        RefSibling sibling = new RefSibling(13);
        System.out.println(fromParent(parent));
        System.out.println(fromParent(child));
        System.out.println(fromParent(null));
        System.out.println(fromObject(other));
        System.out.println(fromScore(child));
        RefChild[] children = new RefChild[2];
        children[0] = child;
        System.out.println(fromArray(children));
        RefParent[] parents = children;
        System.out.println(parents[0].score());
        stash(child, other);
        System.out.println(fromParent(choose(saved, other)));
        System.out.println(merge(false, child, sibling).score());
        System.out.println(merge(true, child, sibling).score());
        System.out.println(mergeArrays(false, children, new RefSibling[3]));
        System.out.println(join(true, child, other));
        System.out.println(join(false, child, other));
    }
}
