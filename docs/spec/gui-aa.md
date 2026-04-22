# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-color`

### `_clampByte(value)`

### `_clampOpacity(value)`

### `rgb(r, g, b)`

> returns `:bool`

### `colorR(color)`

### `colorG(color)`

### `colorB(color)`

### `opacity(color, amount, background)`

### `rgba(r, g, b, a, background)`

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `lib\gui-aa.oak`

- `guiColor` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_abs(x)`

> returns `:int`

### `_floor(x)`

### `_fpart(x)`

### `_rfpart(x)`

> returns `:int`

### `_ipart(x)`

### `_round(x)`

### `_swap(a, b)`

> returns `:list`

### `_min(a, b)`

### `_max(a, b)`

### `_blendColor(fg, alpha, bg)`

### `_smoothstep(edge0, edge1, x)`

### `drawLineAA(deps, window, x0, y0, x1, y1, color, bgColor)`

### `drawCircleFilledAA(deps, window, cx, cy, radius, color, bgColor)`

### `drawCircleOutlineAA(deps, window, cx, cy, radius, color, bgColor)`

### `drawEllipseFilledAA(deps, window, cx, cy, rx, ry, color, bgColor)`

### `drawRoundedRectFilledAA(deps, window, x, y, width, height, radius, color, bgColor)`

### `_edgeDistSigned(px, py, ax, ay, bx, by)`

> returns `:int`

### `_triMin3(a, b, c)`

### `_triMax3(a, b, c)`

### `drawTriangleFilledAA(deps, window, p0, p1, p2, color, bgColor)`

### `drawCircleFilledAAParallel(deps, window, cx, cy, radius, color, bgColor, numWorkers)`

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

