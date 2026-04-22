# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gpu`

### `cudaLibraries()`

> returns `:list`

### `openclLibraries()`

> returns `:list`

### `backends()`

> returns `:list`

### `_resolveFirst(libraries, symbol, i)`

> returns `:object`

### `resolve(library, symbol)`

### `call(procOrAddress, args...)`

### `cuda(symbol)`

### `opencl(symbol)`

### `_scanBackend(backend)`

> returns `:object`

### `scan()`

### `available()`

## Module: `lib\gpus.oak`

- `gpu` · `import(...)`
- `_defaultOrder` · `[2 items]`
### `_normalizeOrder(order)`

> returns `:list`

### `_findBackend(backends, name, i)`

> returns `?`

### `select(order)`

### `_resolveWithBackend(backendName, symbol)`

### `resolve(symbol, order)`

### `call(procOrAddress, args...)`

### `invoke(symbol, args...)`

### `scan()`

### `available()`

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

