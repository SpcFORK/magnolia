# Magnolia Surfaces (Pocket-grammar integration)

Magnolia adopts the **PocketJS surface grammar** without embedding QuickJS or Rust:

```
Runtime = ⟨ Core, Surface, Guest(Magnolia), Host ⟩
```

> Don’t share a world ontology. Share the *grammar* for defining world ontologies.

## Laws

1. **State lives in cores; guests hold mirrors.**  
   Hot-path layout/paint never re-enters guest code.
2. **Intent = ops; facts = events.**  
   Everything crossing the boundary is enumerable and versioned (`ui-surface`).
3. **One guest turn per fixed host tick** (when using `ui-sim`).  
   Frame content is a pure function of tick index + input.
4. **Surfaces stay closed and small.**  
   Op/prop codes are append-only; never renumber.
5. **Determinism is first-class for UI cores** via `ui-sim` (virtual clock + input tape + DrawList hashes).

## Modules

| Module | Role |
| --- | --- |
| `ui-surface` | Single source of truth: ops, props, enums, DrawList opcodes, fingerprint |
| `ui-core` | Retained node tree, flex subset layout, style resolve, tick, DrawList emit |
| `ui-sdk` | Guest helpers: `View` / `Text`, mirror tree, `frame()` |
| `style-bake` | Utility class compiler + MGST `styles.bin` encode/decode + AST collectors |
| `font-bake` | Codepoint atlas plan for pack/VFS |
| `transpile-ui` | Two-pass collect → bake hooks for the build pipeline |
| `gui-drawlist` | Present DrawList through existing GUI backends |
| `ui-sim` | Fixed virtual clock, `after()`, effect queue, golden traces |

## Minimal example

```oak
std := import('std')
surface := import('ui-surface')
uiSdk := import('ui-sdk')
uiCore := import('ui-core')

app := uiSdk.createApp({ width: 480, height: 272 })
box := uiSdk.View(app, surface.ROOT_ID, 'flex flex-col gap-4 p-4 bg-slate-50')
uiSdk.Text(app, box, 'Hello', 'text-xl text-slate-950')
dl := uiSdk.frame(app)
std.println(uiCore.drawListHash(dl))
```

See `samples/ui-surface-demo.oak`.

## Two-pass build (transpile)

```oak
transpileUi := import('transpile-ui')
// Pass 1: collect class strings + codepoints from AST
state := transpileUi.collect(ast)
// Pass 1b: bake styles.bin + font atlas plan
baked := transpileUi.bake(state)
// Pass 2: existing build/pack embeds baked.stylesBin under transpileUi.VFS_STYLES
```

## Deterministic simulation

```oak
uiSim := import('ui-sim')
trace := uiSim.runScenario({
    width: 480
    height: 272
    hz: 60
    frames: 30
    setup: fn(app, sim) { /* build UI */ }
    script: [
        { at: 0.5, click: { x: 40, y: 40 } }
    ]
})
// trace.finalHash / trace.hashes are golden-stable
```

## DrawList → host

```oak
gui := import('GUI')
drawlist := import('gui-drawlist')
// window := gui.createWindow(...)
// drawlist.presentGUI(gui, window, dl)
```

## Compatibility

Legacy `GUI` / `gui-form` immediate-mode drawing is unchanged. New apps opt into the retained surface stack. Do not merge both ontologies into one giant API.

## Drift guard

`surface.fingerprint()` encodes the surface revision. Tests assert op/prop knowledge and encode/decode roundtrips for `styles.bin`. When extending the surface, **append only** and update tests.

## Implementation notes

### Prop resolution cache (ui-core)

`_resolveProps` merges defaults ← style-bake table entry ← node props. It is called
several times per node per layout/draw/hitTest pass (measure, place, emit, test),
so the core keeps an `epoch` counter and a per-node-id `resolveCache`. Every
mutating op (`createNode`, `destroyNode`, `insertBefore`, `removeChild`,
`setStyle`, `setProp`, `setText`, `loadStyles`) bumps the epoch and clears the
cache. Property resolution is therefore computed at most once per node per
quiescent frame, and guest mirrors can rely on stable object identity within a
frame.

### Text measurement cache (ui-core)

`_measureText(text, fontSize)` uses a bounded module-level memo (256 entries,
flush-when-full). The advance model is a pure function of `(fontSize, text)`,
so sharing the cache across cores is safe; it exists because flex measure
passes call it for every text node twice per frame.

### Pointer hit testing (ui-core)

`hitTest` walks in paint order (parent, then children in order, last sibling
paints last). The node under the pointer that was painted last wins, matching
the emission order used by `draw`. Cache/layout state is not required for
correctness, only for speed; `layoutDirty?` still forces a re-layout before
testing.

### Perf expectations

The reference hot path (flex column of N view/text children, M setProps per
frame) performs one `_resolveProps` per node per epoch instead of
O(passes × nodes). `drawListHash` remains the source of truth for golden
comparisons and is intentionally unchanged.
