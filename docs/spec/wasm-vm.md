# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

## Module: `lib\wasm-vm.oak`

- `std` · `import(...)`
- `map` — constant
- `each` — constant
- `reduce` — constant
- `append` — constant
- `indexOf` — constant
- `len` — constant
- `clone` — constant
- `slice` — constant
- `str` · `import(...)`
- `join` — constant
- `fmt` · `import(...)`
- `format` — constant
- `OP_HALT` · `0`
- `OP_NOP` · `1`
- `OP_CONST_NULL` · `2`
- `OP_CONST_EMPTY` · `3`
- `OP_CONST_TRUE` · `4`
- `OP_CONST_FALSE` · `5`
- `OP_CONST_INT` · `6`
- `OP_CONST_FLOAT` · `7`
- `OP_CONST_STRING` · `8`
- `OP_CONST_ATOM` · `9`
- `OP_POP` · `10`
- `OP_DUP` · `11`
- `OP_LOAD_LOCAL` · `12`
- `OP_STORE_LOCAL` · `13`
- `OP_LOAD_UPVAL` · `14`
- `OP_STORE_UPVAL` · `15`
- `OP_ADD` · `16`
- `OP_SUB` · `17`
- `OP_MUL` · `18`
- `OP_DIV` · `19`
- `OP_MOD` · `20`
- `OP_POW` · `21`
- `OP_NEG` · `22`
- `OP_BAND` · `23`
- `OP_BOR` · `24`
- `OP_BXOR` · `25`
- `OP_BRSHIFT` · `26`
- `OP_EQ` · `27`
- `OP_NEQ` · `28`
- `OP_GT` · `29`
- `OP_LT` · `30`
- `OP_GEQ` · `31`
- `OP_LEQ` · `32`
- `OP_NOT` · `33`
- `OP_CONCAT` · `34`
- `OP_MAKE_LIST` · `35`
- `OP_MAKE_OBJECT` · `36`
- `OP_GET_PROP` · `37`
- `OP_SET_PROP` · `38`
- `OP_JUMP` · `39`
- `OP_JUMP_FALSE` · `40`
- `OP_CLOSURE` · `41`
- `OP_CALL` · `42`
- `OP_RETURN` · `43`
- `OP_TAIL_CALL` · `44`
- `OP_BUILTIN` · `45`
- `OP_IMPORT` · `46`
- `OP_IMPORT_DYN` · `47`
- `OP_DEEP_EQ` · `48`
- `OP_SWAP` · `49`
- `OP_MATCH_JUMP` · `50`
- `OP_SCOPE_PUSH` · `51`
- `OP_SCOPE_POP` · `52`
- `OP_CALL_SPREAD` · `53`
- `NUM_OPCODES` · `54`
- `OpcodeName` · `[54 items]`
- `BUILTIN_PRINT` · `0`
- `BUILTIN_LEN` · `1`
- `BUILTIN_TYPE` · `2`
- `BUILTIN_STRING` · `3`
- `BUILTIN_INT` · `4`
- `BUILTIN_FLOAT` · `5`
- `BUILTIN_CODEPOINT` · `6`
- `BUILTIN_CHAR` · `7`
- `BUILTIN_KEYS` · `8`
- `BUILTIN_VALUES` · `9`
- `BUILTIN_SLICE` · `10`
- `BUILTIN_APPEND` · `11`
- `BUILTIN_WAIT` · `12`
- `BUILTIN_EXIT` · `13`
- `BuiltinName` · `[14 items]`
- `TAG_NULL` · `0`
- `TAG_EMPTY` · `1`
- `TAG_INT` · `2`
- `TAG_FLOAT` · `3`
- `TAG_BOOL` · `4`
- `TAG_STRING` · `5`
- `TAG_ATOM` · `6`
- `TAG_LIST` · `7`
- `TAG_OBJECT` · `8`
- `TAG_FUNCTION` · `9`
### `createCompiler()`

> returns `:object`

### `emitByte(c, b)`

### `emitU16(c, v)`

### `emitI32(c, v)`

### `patchI32(c, offset, v)`

> returns `:bool`

### `currentOffset(c)`

### `addConstant(c, entry)`

### `addString(c, s)`

### `addAtom(c, name)`

### `addFloat(c, v)`

### `resolveLocal(c, name)`

### `declareLocal(c, name)`

### `resolveUpvalue(c, name)`

### `resolveBuiltin(name)`

### `compileNode(c, node)`

### `compileAssignment(c, node)`

### `compileIfExpr(c, node)`

### `compileFunction(c, node)`

### `compileGenericCall(c, node)`

### `compileFnCall(c, node)`

### `compileProgram(node)`

### `serializeBytecodeToString(c)`

### `serializeConstantPool(c)`

### `serializeFunctionTable(c)`

### `serializeTopLevelNames(c)`

### `disassemble(c)`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `str`

### `checkRange(lo, hi)`

> **thunk** returns `:function`

### `upper?(c)`

> returns `:bool`

### `lower?(c)`

> returns `:bool`

### `digit?(c)`

> returns `:bool`

### `space?(c)`

> returns `:bool`

### `letter?(c)`

> returns `:bool`

### `word?(c)`

> returns `:bool`

### `join(strings, joiner)`

> returns `:string`

### `startsWith?(s, prefix)`

### `endsWith?(s, suffix)`

### `_matchesAt?(s, substr, idx)`

> returns `:bool`

### `indexOf(s, substr)`

### `rindexOf(s, substr)`

### `contains?(s, substr)`

### `cut(s, sep)`

> returns `:list`

### `lower(s)`

### `upper(s)`

### `_replaceNonEmpty(s, old, new)`

### `replace(s, old, new)`

### `_splitNonEmpty(s, sep)`

### `split(s, sep)`

### `_extend(pad, n)`

### `padStart(s, n, pad)`

### `padEnd(s, n, pad)`

### `_trimStartSpace(s)`

### `_trimStartNonEmpty(s, prefix)`

### `trimStart(s, prefix)`

### `_trimEndSpace(s)`

### `_trimEndNonEmpty(s, suffix)`

### `trimEnd(s, suffix)`

### `trim(s, part)`

