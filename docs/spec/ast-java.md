# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

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

