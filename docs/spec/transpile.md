# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `ast-transform`

- `std` · `import(...)`
- `clone` — constant
- `entries` — constant
### `wrapBlock(nodes)`

> returns `:object`

### `wrapModule(block)`

> returns `:object`

### `formatIdentForWeb(name, key)`

> returns `:string`

### `formatIdentForOak(name)`

> returns `:string`

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

## Module: `lib\transpile.oak`

- `std` · `import(...)`
- `map` — constant
- `each` — constant
- `clone` — constant
- `merge` — constant
- `filter` — constant
- `reduce` — constant
- `fmt` · `import(...)`
- `printf` — constant
- `format` — constant
- `transform` · `import(...)`
- `TranspileRegistry` · `[]`
- `TranspileConfig` · `{3 entries}`
### `configure(config)`

### `registerTranspiler(transpiler)`

### `clearTranspilers()`

> returns `:list`

### `applyTranspiler(node, transpiler)`

### `transpileNode(node)`

### `walkNode(node, visitor)`

### `_mkInt(pos, value)`

> returns `:object`

### `_mkFloat(pos, value)`

> returns `:object`

### `_mkBool(pos, value)`

> returns `:object`

### `_mkString(pos, value)`

> returns `:object`

### `_isConst?(n)`

> returns `:bool`

### `_isNum?(n)`

> returns `:bool`

### `_isInt?(n)`

### `optimizeConstants(node)`

### `removeDebugCalls(node)`

### `createTranspiler(visitor)`

> **thunk** returns `:function`

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

