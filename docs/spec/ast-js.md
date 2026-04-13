# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ast-js.oak`

- `JSReserved` · `[39 items]`
### `_tok()`

> returns `:object`

### `jsNull()`

> returns `:object`

### `jsUndefined()`

> returns `:object`

### `jsNumber(val)`

> returns `:object`

### `jsString(val)`

> returns `:object`

### `jsBool(val)`

> returns `:object`

### `jsIdentifier(name)`

> returns `:object`

### `jsArray(elems)`

> returns `:object`

### `jsObject(entries)`

> returns `:object`

### `jsObjectEntry(key, val)`

> returns `:object`

### `jsBinary(op, left, right)`

> returns `:object`

### `jsUnary(op, operand)`

> returns `:object`

### `jsAssignment(left, right)`

> returns `:object`

### `jsProperty(object, property)`

> returns `:object`

### `jsIndex(object, index)`

> returns `:object`

### `jsCall(callee, args)`

> returns `:object`

### `jsFunction(name, params, body)`

> returns `:object`

### `jsArrowFunction(params, body)`

> returns `:object`

### `jsReturn(value)`

> returns `:object`

### `jsBlock(stmts)`

> returns `:object`

### `jsIf(test, consequent, alternate)`

> returns `:object`

### `jsVar(kind, name, init)`

> returns `:object`

### `jsTernary(test, consequent, alternate)`

> returns `:object`

### `jsTemplateLiteral(parts)`

> returns `:object`

### `jsSpread(argument)`

> returns `:object`

### `_mapJSOp(op)`

> returns `:atom`

### `_mapJSUnaryOp(op)`

> returns `:atom`

### `_jsIdentToOak(name)`

> returns `:string`

### `transpileToOak(node)`

> returns `:object`

### `walkJS(node, visitor)`

> returns `?`

### `renderJS(node)`

> returns `:string`

