# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-shader-codegen.oak`

- `threadLib` · `import(...)`
### `glslVersion(ver?)`

> returns `:string`

### `glslPrecision(prec?, type?)`

> returns `:string`

### `glslStdUniforms()`

> returns `:string`

### `glslUniform(type, name)`

> returns `:string`

### `glslUniforms(uniforms)`

### `glslIn(type, name)`

> returns `:string`

### `glslOut(type, name)`

> returns `:string`

### `glslQuadVertex()`

### `glslQuadVertexCompat()`

> returns `:string`

### `glslFragmentWrap(body, version?)`

### `glslMathLib()`

> returns `:string`

### `hlslStdCBuffer()`

> returns `:string`

### `hlslCBuffer(name, uniforms)`

### `hlslQuadVertex()`

> returns `:string`

### `hlslFragmentWrap(body)`

### `hlslMathLib()`

> returns `:string`

### `submitWebGL(window, fragSource, vertSource?)`

> returns `:object`

### `drawWebGL(window, clearR?, clearG?, clearB?)`

### `renderWebGL(window, fragSource)`

### `compileGLSL(source, stage?, outputPath?)`

### `compileHLSL(source, profile?, entry?, outputPath?)`

### `compileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `assembleGLSL(opts)`

### `assembleHLSL(opts)`

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

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

