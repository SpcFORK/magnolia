# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ast-ts.oak`

- `astJs` · `import(...)`
- `jsNull` — constant
- `jsUndefined` — constant
- `jsNumber` — constant
- `jsString` — constant
- `jsBool` — constant
- `jsIdentifier` — constant
- `jsArray` — constant
- `jsObject` — constant
- `jsObjectEntry` — constant
- `jsBinary` — constant
- `jsUnary` — constant
- `jsAssignment` — constant
- `jsProperty` — constant
- `jsIndex` — constant
- `jsCall` — constant
- `jsFunction` — constant
- `jsArrowFunction` — constant
- `jsReturn` — constant
- `jsBlock` — constant
- `jsIf` — constant
- `jsVar` — constant
- `jsTernary` — constant
- `jsSpread` — constant
### `_tok()`

> returns `:object`

### `tsTypeAnnotation(typeExpr)`

> returns `:object`

### `tsTypeRef(name)`

> returns `:object`

### `tsArrayType(elementType)`

> returns `:object`

### `tsUnionType(types)`

> returns `:object`

### `tsIntersectionType(types)`

> returns `:object`

### `tsGenericType(name, typeArgs)`

> returns `:object`

### `tsFunctionType(params, returnType)`

> returns `:object`

### `tsParam(name, typeAnnotation)`

> returns `:object`

### `tsInterface(name, members, extends)`

> returns `:object`

### `tsPropertySig(name, typeAnnotation, optional?)`

> returns `:object`

### `tsMethodSig(name, params, returnType)`

> returns `:object`

### `tsTypeAlias(name, typeExpr)`

> returns `:object`

### `tsEnum(name, members)`

> returns `:object`

### `tsEnumMember(name, value)`

> returns `:object`

### `tsAsExpr(expr, typeExpr)`

> returns `:object`

### `tsNonNullExpr(expr)`

> returns `:object`

### `tsTypedVar(kind, name, typeAnnotation, init)`

> returns `:object`

### `tsTypedFunction(name, params, returnType, body)`

> returns `:object`

### `renderType(typeNode)`

> returns `:string`

### `transpileToOak(node)`

> returns `:object`

### `walkTS(node, visitor)`

> returns `?`

### `renderTS(node)`

> returns `:string`

