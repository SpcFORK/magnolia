# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

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

