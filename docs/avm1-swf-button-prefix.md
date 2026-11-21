# Uploaded SWF5: virtual original Play-button release

**Bounded result, not movie playback.** The unchanged uploaded SWF5's original `DefineButton2` character **10** has four independently decoded button records and one conditional action. The button is placed on sprite 11's `loaded` frame at depth 6. Its original on-release action consists of **six binary AVM1 records**: push zero and `_root`, `GetVariable`, push `play`, `CallMethod`, `Pop`, `End`. Magnolia executes those original bytes and records one virtual zero-argument **`_root.play()`** request with an empty final stack.

[Executable independent oracle](../avm1-swf-button-prefix-oracle.py), [machine-readable report](avm1-swf-button-prefix.json), native exit **0**, report SHA-256 `943b87f4db642b35ea668e4102591a95855e09e33fc9733a0e961a2c246657d1`. The independent Python code re-walks the full original container, locates the actual first-frame Button2 and sprite-11 loaded placement, parses all button-state matrices/color transforms and the condition block, and separately interprets the actual action bytes. It compares every emitted Oak field against that independent structural/stack result. Its structural understanding of `DefineButton2`/`BUTTONCONDACTION` follows the [SWF format specification](https://open-flash.github.io/mirrors/swf-spec-19.pdf); it is not a second Flash player's execution trace.

| Evidence | Value |
| --- | --- |
| Unchanged uploaded FWS5 SHA-256 | `2bde4d96b5b6695f8dd81afe005080c2bc6b9a14545f7cc8f35f49c71ac43cbc` |
| Button tag header / body | offsets **1859 / 1865**, 85 bytes; body SHA-256 `6988703abda47627e511189c63defc73f0ba275056e740cc3f1dd295cecdfd4f` |
| Four state records | `(state bits, character, depth)`: **(15,8,1), (1,9,2), (2,9,2), (12,9,2)**; full decoded transforms are in the report |
| Sprite-11 loaded frame | label offset **2247**, removes depths 2/3/6 and places button 10 at depth **6**, offset **2260** |
| Conditional event | `OverDownToOverUp` (`08 00` condition bytes), explicitly **synthetic**, not hit-tested |
| Action | offset **1918**, 32 bytes, SHA-256 `796f9f741e88f58dd858314ff71f3a922db3b9e7aed9242ba0cbd5a298a3d34c`; visited instruction offsets **0,19,20,29,30,31** |

The caller first executes the existing first-frame preloader against a **synthetic** completed `1/1` byte snapshot, verifies its pending request for the `loaded` label, explicitly selects the declared loaded-frame button placement, and dispatches one synthetic release. The additional `runButton` admission is deliberately opt-in: the older `run` direct-preloader interface still refuses this new `_root.play` call. Six same-length structural/action corruptions, an unknown option, changed whole-source hash and the old-host refusal are checked before public output. The preloader direct-action oracle (seven frame-0 snapshots, two frame-1 direct calls, 24 refusals) and whole-first-frame oracle (four snapshots, nine refusals) were separately rerun with the new admission and exited **0**. A separately generated Node bundle also passed the original 32-byte direct action/legacy-refusal check when constructed byte-by-byte using `char`; this is only the action slice, not a whole-SWF or browser test.

**JS literal follow-up:** The earlier `00 7f 80 ff` probe exposed replacement characters from raw high bytes in generated JS. The later [web byte-compatibility work](web-compatibility-progress.md) emits ASCII byte escapes and verifies all 256 byte values and ordinary UTF-8 text in Node and Chrome. A [separate original-action literal oracle](../web-original-button-literal.py) now extracts the pinned 32-byte action from the unchanged upload and confirms the same six-record virtual trace in native, Node and real Chrome, with the legacy-host refusal. This is a synthetic direct action, not a physical click or whole-SWF playback.

This does **not** establish automatic progression to the `loaded` frame, real loading counts, the physical mouse position/hit area, a real event queue, later root frame entry, image/text rendering, or decoded MP3. The separate [Ruffle real-input reference](uploaded-swf5-ruffle-postinput.json) shows a real pointer changes the user's screen, but its pixels were **not** compared to this virtual host. No general AVM1, full movie or checklist acceptance is claimed.

Reproduce from the checkout with the unchanged upload held outside it:

```sh
go build -o /var/tmp/magnolia-button .
python3 avm1-swf-button-prefix-oracle.py \
  --swf /external/asdf.swf.txt --magnolia /var/tmp/magnolia-button \
  --output /external/new-button-report.json
```

The Oak CLI `avm1-swf-button-prefix.oak` and embedded `avm1-swf-button-prefix` library are deliberately unpinned on their own; only the external oracle SHA-checks the entire upload before invoking the native host. Original source bytes remain outside the checkout.
