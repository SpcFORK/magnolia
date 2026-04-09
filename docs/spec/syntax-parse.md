# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\syntax-parse.oak`

- `std` · `import(...)`
- `fromHex` — constant
- `slice` — constant
- `append` — constant
- `last` — constant
- `filter` — constant
- `map` — constant
- `each` — constant
- `str` · `import(...)`
- `strContains?` — constant
- `fmt` · `import(...)`
- `format` — constant
### `cloneNameSet(set)`

### `addPatternBindings(shadowed, node)`

> returns `:bool`

### `rewriteClassSugarAssignmentLeft(node, visibleFields, allFields, shadowed, selfName, isLocal)`

### `rewriteClassSugarNode(node, visibleFields, allFields, shadowed, selfName)`

### `classBodyFromAssignmentBlock(body, reservedNames)`

> returns `:list`

### `wrapBodyWithSelfVar(body, reservedNames)`

> returns `:object`

### `Parser(tokens)`

> returns `:object`

### `parse(text)`

