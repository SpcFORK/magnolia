# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\build-render.oak`

- `std` · `import(...)`
- `map` — constant
- `reduce` — constant
- `append` — constant
- `keys` — constant
- `len` — constant
- `str` · `import(...)`
- `join` — constant
- `replace` — constant
- `sort` · `import(...)`
- `sort!` — constant
- `fmt` · `import(...)`
- `format` — constant
- `buildRenderNode` · `import(...)`
### `renderVfsOak(vfsFiles, encodeJSON)`

> returns `:string`

### `renderVfsJs(vfsFiles, encodeJSON)`

> returns `:string`

### `renderOakBundle(bundleNode, vfsFiles, encodeJSON, oakNativeRuntime, formatIdent, abort)`

> returns `:string`

### `renderJSBundle(bundleNode, vfsFiles, encodeJSON, oakJSRuntime, formatIdent, clone)`

> returns `:string`

### `renderBundle(bundleNode, target, renderOak, renderJS, renderWasm, renderAst, renderBin, renderDoc, renderTs, renderLua, renderJava, renderGraph)`

