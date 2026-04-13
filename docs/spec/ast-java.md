# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ast-java.oak`

- `JavaReserved` · `[55 items]`
### `_tok()`

> returns `:object`

### `javaNull()`

> returns `:object`

### `javaInt(val)`

> returns `:object`

### `javaDouble(val)`

> returns `:object`

### `javaString(val)`

> returns `:object`

### `javaBool(val)`

> returns `:object`

### `javaChar(val)`

> returns `:object`

### `javaIdentifier(name)`

> returns `:object`

### `javaArray(elems)`

> returns `:object`

### `javaNewArray(elementType, elems)`

> returns `:object`

### `javaMap(entries)`

> returns `:object`

### `javaMapEntry(key, val)`

> returns `:object`

### `javaBinary(op, left, right)`

> returns `:object`

### `javaUnary(op, operand)`

> returns `:object`

### `javaAssignment(left, right)`

> returns `:object`

### `javaDot(object, field)`

> returns `:object`

### `javaIndex(object, index)`

> returns `:object`

### `javaCall(callee, args)`

> returns `:object`

### `javaMethodCall(object, method, args)`

> returns `:object`

### `javaNew(className, args)`

> returns `:object`

### `javaLambda(params, body)`

> returns `:object`

### `javaMethod(modifiers, returnType, name, params, body)`

> returns `:object`

### `javaVarDecl(javaType, name, init)`

> returns `:object`

### `javaReturn(value)`

> returns `:object`

### `javaBlock(stmts)`

> returns `:object`

### `javaIf(test, consequent, alternate)`

> returns `:object`

### `javaFor(init, test, update, body)`

> returns `:object`

### `javaForEach(javaType, name, iterable, body)`

> returns `:object`

### `javaWhile(test, body)`

> returns `:object`

### `javaCast(javaType, expr)`

> returns `:object`

### `javaTernary(test, consequent, alternate)`

> returns `:object`

### `javaInstanceOf(expr, javaType)`

> returns `:object`

### `javaClass(name, members)`

> returns `:object`

### `_mapJavaOp(op)`

> returns `:atom`

### `_mapJavaUnaryOp(op)`

> returns `:atom`

### `_javaIdentToOak(name)`

> returns `:string`

### `transpileToOak(node)`

> returns `:object`

### `walkJava(node, visitor)`

> returns `?`

### `renderJava(node)`

> returns `:string`

