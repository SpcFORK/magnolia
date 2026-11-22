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

### `_screenPositionFor(size, w, h, position, margin)`

> returns `:object`

- `POSITION_CENTER` · `:center`
- `POSITION_PRIMARY_CENTER` · `:primaryCenter`
- `POSITION_MONITOR_CENTER` · `:monitorCenter`
- `POSITION_TOP_LEFT` · `:topLeft`
- `POSITION_TOP_RIGHT` · `:topRight`
- `POSITION_BOTTOM_LEFT` · `:bottomLeft`
- `POSITION_BOTTOM_RIGHT` · `:bottomRight`
- `POSITION_TOP_CENTER` · `:topCenter`
- `POSITION_BOTTOM_CENTER` · `:bottomCenter`
- `POSITION_LEFT_CENTER` · `:leftCenter`
- `POSITION_RIGHT_CENTER` · `:rightCenter`
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

- `__std_builtin_int` — constant
- `__std_builtin_float` — constant
- `__std_builtin_string` — constant
- `__std_builtin_type` — constant
- `__std_builtin_len` — constant
- `__std_builtin_char` — constant
- `__std_builtin_codepoint` — constant
- `__std_builtin_print` — constant
- `__std_builtin_keys` — constant
- `__std_builtin_atom` — constant
- `__std_builtin_name` — constant
- `__std_builtin_rand` — constant
- `__std_builtin_srand` — constant
- `__std_builtin_input` — constant
- `__std_builtin_exit` — constant
- `__std_builtin_args` — constant
- `__std_builtin_addr` — constant
- `__std_builtin_bits` — constant
- `__std_builtin_memread` — constant
- `__std_builtin_memwrite` — constant
- `__std_builtin_utf16` — constant
- `__std_builtin_go` — constant
- `__std_builtin_wait` — constant
- `__std_builtin_exec` — constant
- `__std_builtin_env` — constant
- `__std_builtin_write` — constant
- `__std_builtin_read` — constant
- `__std_builtin_interpreter` — constant
- `__std_builtin_pointer` — constant
- `__std_builtin_log` — constant
- `__std_builtin_time` — constant
- `__std_builtin_nanotime` — constant
- `int` — constant
- `float` — constant
- `string` — constant
- `type` — constant
- `len` — constant
- `print` — constant
- `time` — constant
- `nanotime` — constant
- `char` — constant
- `codepoint` — constant
- `keys` — constant
- `atom` — constant
- `name` — constant
- `rand` — constant
- `srand` — constant
- `input` — constant
- `args` — constant
- `addr` — constant
- `bits` — constant
- `memread` — constant
- `memwrite` — constant
- `utf16` — constant
- `exec` — constant
- `env` — constant
- `write` — constant
- `read` — constant
- `interpreter` — constant
- `pointer` — constant
- `log` — constant
- `wait` — constant
- `__std_builtin_sin` — constant
- `__std_builtin_cos` — constant
- `__std_builtin_tan` — constant
- `__std_builtin_asin` — constant
- `__std_builtin_acos` — constant
- `__std_builtin_atan` — constant
- `__std_builtin_pow` — constant
- `sin` — constant
- `cos` — constant
- `tan` — constant
- `asin` — constant
- `acos` — constant
- `atan` — constant
- `pow` — constant
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

### `pmapWithPool(list, fnToRun, pool)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

