# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ast-go.oak`

- `GoReserved` · `[25 items]`
- `GoBuiltins` · `[41 items]`
### `_tok()`

> returns `:object`

### `goNil()`

> returns `:object`

### `goInt(val)`

> returns `:object`

### `goFloat(val)`

> returns `:object`

### `goString(val)`

> returns `:object`

### `goRawString(val)`

> returns `:object`

### `goBool(val)`

> returns `:object`

### `goRune(val)`

> returns `:object`

### `goIdentifier(name)`

> returns `:object`

### `goSliceLit(elems)`

> returns `:object`

### `goArrayLit(size, elems)`

> returns `:object`

### `goMapLit(entries)`

> returns `:object`

### `goMapEntry(key, val)`

> returns `:object`

### `goStructLit(typeName, fields)`

> returns `:object`

### `goStructField(name, val)`

> returns `:object`

### `goBinary(op, left, right)`

> returns `:object`

### `goUnary(op, operand)`

> returns `:object`

### `goAssignment(left, right)`

> returns `:object`

### `goShortDecl(left, right)`

> returns `:object`

### `goVarDecl(name, goType, init)`

> returns `:object`

### `goConstDecl(name, goType, init)`

> returns `:object`

### `goDot(object, field)`

> returns `:object`

### `goIndex(object, index)`

> returns `:object`

### `goSliceExpr(object, low, high, max)`

> returns `:object`

### `goCall(callee, args)`

> returns `:object`

### `goMethodCall(object, method, args)`

> returns `:object`

### `goFunc(name, params, results, body)`

> returns `:object`

### `goFuncParam(name, goType)`

> returns `:object`

### `goFuncLit(params, results, body)`

> returns `:object`

### `goReturn(values)`

> returns `:object`

### `goBlock(stmts)`

> returns `:object`

### `goIf(init, cond, consequent, alternate)`

> returns `:object`

### `goFor(init, cond, post, body)`

> returns `:object`

### `goForRange(key, value, iterable, body)`

> returns `:object`

### `goSwitch(tag, cases)`

> returns `:object`

### `goSwitchCase(exprs, body)`

> returns `:object`

### `goTypeSwitch(assign, cases)`

> returns `:object`

### `goSelect(cases)`

> returns `:object`

### `goSelectCase(comm, body)`

> returns `:object`

### `goGo(call)`

> returns `:object`

### `goDefer(call)`

> returns `:object`

### `goChanSend(ch, value)`

> returns `:object`

### `goChanRecv(ch)`

> returns `:object`

### `goMake(goType, args)`

> returns `:object`

### `goTypeAssert(expr, goType)`

> returns `:object`

### `goCast(goType, expr)`

> returns `:object`

### `goStruct(name, fields)`

> returns `:object`

### `goInterface(name, methods)`

> returns `:object`

### `goTypeDecl(name, underlying)`

> returns `:object`

### `goImport(path, alias)`

> returns `:object`

### `goPackage(name)`

> returns `:object`

### `goComment(text)`

> returns `:object`

### `goMultiReturn(values)`

> returns `:object`

### `goMultiAssign(lefts, rights)`

> returns `:object`

### `goBlank()`

> returns `:object`

### `goEllipsis(expr)`

> returns `:object`

### `goPointerType(base)`

> returns `:object`

### `goAddrOf(expr)`

> returns `:object`

### `goDeref(expr)`

> returns `:object`

### `_mapGoOp(op)`

> returns `:atom`

### `_mapGoUnaryOp(op)`

> returns `:atom`

### `_goIdentToOak(name)`

> returns `:string`

### `transpileToOak(node)`

> returns `:object`

