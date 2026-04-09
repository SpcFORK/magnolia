# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\Virtual-Bytecode.oak`

- `std` · `import(...)`
- `syntax` · `import(...)`
- `wasmVM` · `import(...)`
- `fmt` · `import(...)`
- `printf` — constant
- `str` · `import(...)`
- `join` — constant
- `builtinKeys` — constant
- `builtinValues` — constant
- `MODE_WASM` · `:wasm`
- `MODE_GO` · `:go`
- `CONST_STRING` · `0`
- `CONST_ATOM` · `1`
- `CONST_FLOAT` · `2`
- `I32_SIGN` · `2147483648`
- `I32_WRAP` · `4294967296`
### `rangeN(n)`

> returns `:list`

### `bytesFromString(s)`

### `bytesFromAny(v)`

### `readU16(bs, at)`

> returns `:bool`

### `readU32(bs, at)`

> returns `:bool`

### `readI32(bs, at)`

### `bytesToString(bs, start, count)`

> returns `:string`

### `normalizeConstEntry(entry, mode)`

> returns `:object`

### `decodeConstantPoolBytes(raw)`

> returns `:list`

### `decodeFunctionTableBytes(raw)`

> returns `:list`

### `detectMode(raw)`

### `normalizeFunctions(rawFunctions)`

### `normalizeChunk(raw)`

> returns `:object`

### `defaultOpcodes(mode)`

> returns `:object`

### `stackPush(vm, v)`

### `stackPop(vm)`

> returns `?`

### `stackPeek(vm)`

> returns `?`

### `currentFrameIndex(vm)`

### `currentFrame(vm)`

### `setCurrentFrame(vm, frame)`

### `currentScope(vm)`

### `scopeGetByName(scope, name)`

> returns `?`

### `scopeSetByName(scope, name, val)`

> returns `:bool`

### `scopeAtDepth(scope, depth)`

### `scopeGetByDepthSlot(scope, depth, slot)`

> returns `?`

### `scopeSetByDepthSlot(scope, depth, slot, val)`

> returns `:bool`

### `fetchU8(vm)`

> returns `:bool`

### `fetchU16(vm)`

### `fetchI32(vm)`

### `constValueAt(vm, idx)`

> returns `?`

### `keyToString(k)`

### `vmGetProp(obj, key)`

> returns `?`

### `vmSetProp(obj, key, val)`

### `isTruthy(v)`

### `deepEq(a, b)`

> returns `:bool`

### `vmAdd(a, b)`

### `vmSub(a, b)`

### `vmMul(a, b)`

### `vmDiv(a, b)`

> returns `?`

### `vmMod(a, b)`

> returns `?`

### `vmPow(a, b)`

### `vmCompare(op, a, b)`

> returns `:bool`

### `vmConcat(a, b)`

### `builtinAppend(x, y)`

### `callBuiltin(idx, args)`

> returns `?`

### `makeFrame(vm, closure, args, returnPC)`

> returns `:object`

### `invokeCallable(vm, callee, args)`

### `popArgs(vm, arity)`

### `runChunk(rawChunk, opts)`

### `compileAst(ast)`

### `compileSource(source)`

### `runSource(source, opts)`

- `run` — constant
- `_defaultImportModules` · `{5 entries}`
### `makeDefaultImportFn()`

> **thunk** returns `:function`

- `defaultImportFn` · `makeDefaultImportFn(...)`
