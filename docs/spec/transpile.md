# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\transpile.oak`

- `std` · `import(...)`
- `map` — constant
- `each` — constant
- `clone` — constant
- `merge` — constant
- `filter` — constant
- `reduce` — constant
- `fmt` · `import(...)`
- `printf` — constant
- `format` — constant
- `transform` · `import(...)`
- `TranspileRegistry` · `[]`
- `TranspileConfig` · `{3 entries}`
### `configure(config)`

### `registerTranspiler(transpiler)`

### `clearTranspilers()`

> returns `:list`

### `applyTranspiler(node, transpiler)`

### `transpileNode(node)`

### `walkNode(node, visitor)`

### `_mkInt(pos, value)`

> returns `:object`

### `_mkFloat(pos, value)`

> returns `:object`

### `_mkBool(pos, value)`

> returns `:object`

### `_mkString(pos, value)`

> returns `:object`

### `_isConst?(n)`

> returns `:bool`

### `_isNum?(n)`

> returns `:bool`

### `_isInt?(n)`

### `optimizeConstants(node)`

### `removeDebugCalls(node)`

### `createTranspiler(visitor)`

> **thunk** returns `:function`

