# API Documentation

_Auto-generated from Magnolia source._

---

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
## Module: `gui-thread`

- `threadLib` · `import(...)`
### `CommandQueue()`

> returns `:object`

### `FrameFence(workerCount)`

> returns `:object`

### `WorkerPool(numWorkers)`

> returns `:object`

### `StateGuard()`

> returns `:object`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

### `AsyncLoader(cmdQueue)`

> returns `:object`

### `FrameScheduler(pool, cmdQueue)`

> returns `:object`

### `initWindowThreading(window, options)`

### `threadingEnabled?(window)`

### `commandQueue(window)`

### `workerPool(window)`

### `scheduler(window)`

### `stateGuard(window)`

### `flushThreadedCommands(window)`

### `destroyWindowThreading(window)`

> returns `?`

## Module: `gui-web`

- `guiThread` · `import(...)`
### `_isFn?(value)`

### `_nextResourceId(webgl)`

### `_webRecord(window, op)`

### `createWindowState(title, width, height, frameMs, updateOnDispatch, options)`

> returns `:object`

### `showWindow(window)`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `getPrimaryScreenSize(window)`

> returns `:object`

### `centerWindow(window)`

### `positionWindow(window, position, options)`

### `setFullscreen(window, enabled)`

### `lockResize(window, locked)`

### `createCanvas(window, id, options)`

### `initWebGL(window, contextName, attrs)`

> returns `:object`

### `webglCreateShader(window, shaderType, source)`

> returns `:object`

### `webglCreateProgram(window, vertexShader, fragmentShader)`

> returns `:object`

### `webglUseProgram(window, program)`

> returns `:object`

### `webglClearColor(window, r, g, b, a)`

> returns `:object`

### `webglViewport(window, x, y, width, height)`

> returns `:object`

### `webglClear(window, mask, colorBufferBit)`

> returns `:object`

### `webglDrawArrays(window, mode, first, count, trianglesMode)`

> returns `:object`

### `webglFlush(window)`

> returns `:object`

### `colorToCss(packed)`

> returns `:string`

### `beginFrame(window)`

### `endFrame(window)`

### `runFrameLoop(window, onFrame)`

### `stopFrameLoop(window)`

### `dispatchEvent(window, evt)`

### `pollEvents(window)`

### `setAlpha(window, alpha)`

### `drawImageData(window, x, y, w, h, pixelBytes)`

## Module: `json`

### `esc(c)`

> returns `:string`

### `uEscape(c)`

> returns `:string`

### `uParse(uc)`

### `escape(s)`

### `serialize(c)`

> returns `:string`

### `cs Reader(s)`

> returns `:object`

### `parseNull(r)`

> returns `?`

### `parseString(r)`

### `parseNumber(r)`

> returns `:atom`

### `parseTrue(r)`

> returns `:bool`

### `parseFalse(r)`

> returns `:bool`

### `parseList(r)`

### `parseObject(r)`

### `_parseReader(r)`

> returns `:atom`

### `parse(s)`

## Module: `lib\web-pipeline.oak`

- `wsModule` · `import(...)`
- `guiWeb` · `import(...)`
- `json` · `import(...)`
- `threadLib` · `import(...)`
### `isWeb?()`

### `isNode?()`

### `isBrowser?()`

- `DefaultContextAttrs` · `{6 entries}`
### `loadResource(url, options)`

### `fetchText(url, options)`

### `fetchJson(url, options)`

### `loadShader(window, url, shaderType)`

### `loadProgram(window, vertUrl, fragUrl)`

### `createContext(window, options)`

> returns `:object`

### `destroyContext(ctx)`

### `WebPipeline(window, options)`

> returns `:object`

### `pipelineResources(window, specs)`

- `Opcode` — constant
- `CloseCode` — constant
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

## Module: `web-fetch`

### `resolved(val)`

> returns `:object`

### `pending()`

> returns `:object`

### `fetch(data)`

## Module: `web-runtime`

### `jsRuntime?()`

### `node?()`

### `browser?()`

## Module: `websocket`

### `connect(url, headers)`

### `send(socket, data, opcode)`

### `recv(socket)`

### `close(socket, code, reason)`

### `listen(host, path, onConnect)`

### `Socket(socket)`

> returns `:object`

- `Opcode` · `{5 entries}`
- `CloseCode` · `{8 entries}`
