// JDK-valid original ArcCanvas.redraw call without a displayable peer.
// Magnolia may return only for a constructed detached, peerless Canvas;
// returning is not a queued repaint, actual geometry or arc pixel painting.
// Compile against the pinned ORIGINAL ArcCanvas.class outside the repository.
public final class ArcCanvasRedrawMain {
    public static void main(String[] args) {
        ArcCanvas canvas = new ArcCanvas();
        canvas.redraw(true, 30, 85);
        System.out.println("JDK Canvas repaint returned");
    }
}
