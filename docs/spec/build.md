# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\build.oak`

- `std` · `import(...)`
- `slice` — constant
- `default` — constant
- `clone` — constant
- `map` — constant
- `each` — constant
- `reduce` — constant
- `append` — constant
- `contains?` — constant
- `merge` — constant
- `once` — constant
- `keys` — constant
- `str` · `import(...)`
- `join` — constant
- `replace` — constant
- `split` — constant
- `startsWith?` — constant
- `sort` · `import(...)`
- `sort!` — constant
- `fs` · `import(...)`
- `readFile` — constant
- `writeFile` — constant
- `statFile` — constant
- `fmt` · `import(...)`
- `printf` — constant
- `format` — constant
- `path` · `import(...)`
- `dir` — constant
- `resolve` — constant
- `json` · `import(...)`
- `encodeJSON` — constant
- `syntax` · `import(...)`
- `bundleUtils` · `import(...)`
- `transform` · `import(...)`
- `transpile` · `import(...)`
- `packModule` · `import(...)`
- `buildIncludes` · `import(...)`
- `buildIdent` · `import(...)`
- `buildAst` · `import(...)`
- `buildAnalyze` · `import(...)`
- `buildRender` · `import(...)`
- `buildConfig` · `import(...)`
- `buildImports` · `import(...)`
- `threadLib` · `import(...)`
- `pmap` — constant
- `runtimeCodegen` · `import(...)`
- `OakNativeRuntime` — constant
- `OakJSRuntime` — constant
- `Log` — constant
### `Abort(msg)`

- `Entry` · `?`
- `Web?` · `false`
- `Wasm?` · `false`
- `Ast?` · `false`
- `Bin?` · `false`
- `Doc?` · `false`
- `Ts?` · `false`
- `Lua?` · `false`
- `Java?` · `false`
- `Graph?` · `false`
- `Target` · `'oak'`
- `Output` · `?`
- `Includes` · `[]`
- `IncludeVFS` · `?`
- `VFSFiles` · `{}`
- `AbsoluteEntry` · `?`
### `requireOpt(flag, value)`

- `parseInclude` — constant
- `parseIncludes` — constant
### `applyConfig(config)`

- `Configured` · `?`
### `configure(config)`

### `compile()`

- `ModuleNodes` · `{}`
- `ImportCallNode` · `{5 entries}`
- `ImportAssignmentNode` · `{5 entries}`
### `cachedParse(path, text)`

### `addImportsFromFile(path)`

### `addImportsFromSource(path, file)`

### `formatIdent(name, key)`

### `wrapModule(block)`

### `bundleCommonPrefix(paths)`

### `wrapBundle(modules, entryModuleName)`

### `analyzeNode(node)`

### `renderOakBundle(bundleNode)`

### `renderJSBundle(bundleNode)`

### `renderWasmBundle(bundleNode)`

### `renderAstBundle(bundleNode)`

### `renderBinBundle(bundleNode)`

### `renderDocBundle(bundleNode)`

### `renderTsBundle(bundleNode)`

### `renderLuaBundle(bundleNode)`

### `renderJavaBundle(bundleNode)`

### `renderGoBundle(bundleNode)`

### `renderNativeBundle(bundleNode)`

### `renderGraphBundle(bundleNode)`

### `renderBundle(bundleNode)`

### `run(config)`

