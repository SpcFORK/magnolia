# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ast-lua.oak`

- `LuaReserved` · `[22 items]`
### `_tok()`

> returns `:object`

### `luaNil()`

> returns `:object`

### `luaNumber(val)`

> returns `:object`

### `luaString(val)`

> returns `:object`

### `luaBool(val)`

> returns `:object`

### `luaIdentifier(name)`

> returns `:object`

### `luaTable(entries)`

> returns `:object`

### `luaTableEntry(key, val)`

> returns `:object`

### `luaBinary(op, left, right)`

> returns `:object`

### `luaUnary(op, operand)`

> returns `:object`

### `luaAssignment(left, right, local?)`

> returns `:object`

### `luaIndex(table, key)`

> returns `:object`

### `luaDot(table, field)`

> returns `:object`

### `luaCall(callee, args)`

> returns `:object`

### `luaMethodCall(object, method, args)`

> returns `:object`

### `luaFunction(name, params, body)`

> returns `:object`

### `luaReturn(values)`

> returns `:object`

### `luaBlock(stmts)`

> returns `:object`

### `luaIf(test, consequent, elseifs, alternate)`

> returns `:object`

### `luaElseIf(test, body)`

> returns `:object`

### `luaForNumeric(var, start, stop, step, body)`

> returns `:object`

### `luaForGeneric(vars, iterators, body)`

> returns `:object`

### `luaWhile(test, body)`

> returns `:object`

### `luaRepeatUntil(body, test)`

> returns `:object`

### `luaConcat(left, right)`

> returns `:object`

### `luaVararg()`

> returns `:object`

### `luaLength(operand)`

> returns `:object`

### `_mapLuaOp(op)`

> returns `:atom`

### `_mapLuaUnaryOp(op)`

> returns `:atom`

### `transpileToOak(node)`

> returns `:object`

### `walkLua(node, visitor)`

> returns `?`

### `renderLua(node)`

> returns `:string`

