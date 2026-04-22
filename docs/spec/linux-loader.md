# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\linux-loader.oak`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_libraryKey(library)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `dlopen(path, flags)`

### `dlsym(handle, symbol)`

### `dlclose(handle)`

### `loadLibrary(path)`

### `freeLibrary(handle)`

### `procAddress(module, symbol)`

### `_loadDllCandidate(candidates, i, flags)`

> returns `:object`

### `loadDll(library)`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `x11(symbol, args...)`

## Module: `linux-constants`

- `LibC` · `[4 items]`
- `LibDL` · `[4 items]`
- `LibX11` · `[4 items]`
- `PROT_NONE` · `0`
- `PROT_READ` · `1`
- `PROT_WRITE` · `2`
- `PROT_EXEC` · `4`
- `MAP_SHARED` · `1`
- `MAP_PRIVATE` · `2`
- `MAP_FIXED` · `16`
- `MAP_ANONYMOUS` · `32`
- `O_RDONLY` · `0`
- `O_WRONLY` · `1`
- `O_RDWR` · `2`
- `O_CREAT` · `64`
- `O_TRUNC` · `512`
- `O_APPEND` · `1024`
- `SEEK_SET` · `0`
- `SEEK_CUR` · `1`
- `SEEK_END` · `2`
- `F_OK` · `0`
- `X_OK` · `1`
- `W_OK` · `2`
- `R_OK` · `4`
- `RTLD_LAZY` · `1`
- `RTLD_NOW` · `2`
- `RTLD_GLOBAL` · `256`
- `RTLD_LOCAL` · `0`
- `_SC_PAGESIZE` · `30`
- `KeyPressMask` · `1`
- `ButtonPressMask` · `4`
- `ExposureMask` · `32768`
- `StructureNotifyMask` · `131072`
- `KeyPress` · `2`
- `Expose` · `12`
- `DestroyNotify` · `17`
- `ClientMessage` · `33`
## Module: `linux-core`

- `sys` · `import(...)`
### `isLinux?()`

### `cstr(s)`

### `_readCString(ptr, maxLen)`

### `_platformError(apiName)`

> returns `:object`

### `_resolveFirst(libraries, symbol, i)`

> returns `:object`

### `_callResolved(resolved, args...)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `libc(symbol, args...)`

### `libdl(symbol, args...)`

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

## Module: `sys`

### `_isObject?(v)`

### `ok?(result)`

> returns `:bool`

### `error?(result)`

> returns `:bool`

### `resolve(library, symbol)`

> returns `:object`

### `call(target, args...)`

### `resolveAndCall(library, symbol, args...)`

### `valueOr(result, fallback)`

