/** Orchestration ONLY: the original Fliplay constructor, init and loader run unmodified. */
public final class FliplayLiveLink {
  public static fliplay load() {
    fliplay player = new fliplay();
    player.init();
    player.load_fli();
    return player;
  }
}
