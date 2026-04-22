# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

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

## Module: `lib\gui-native-linux.oak`

- `linux` · `import(...)`
- `guiThread` · `import(...)`
### `createWindowState(title, width, height, frameMs, updateOnDispatch)`

> returns `:object`

### `showWindow(window)`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `_displaySize(window)`

> returns `:object`

### `setFullscreen(window, enabled)`

### `setTitle(window, title)`

### `lockResize(window, locked)`

### `poll(window)`

> returns `:object`

### `close(window)`

> returns `:int`

### `sleepFrame(frameMs)`

## Module: `linux`

- `sys` · `import(...)`
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

## Module: `linux-libc`

- `sys` · `import(...)`
### `currentProcessId()`

### `parentProcessId()`

### `pageSize()`

### `errno()`

### `strerror(errorCode)`

### `lastErrorMessage()`

### `getLastError()`

### `formatMessage(errorCode)`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `getuid()`

### `geteuid()`

### `getgid()`

### `getegid()`

### `gethostname(bufferPtr, size)`

### `getcwd(bufferPtr, size)`

### `chdir(path)`

### `access(path, mode)`

### `openFile(path, flags, mode)`

### `closeFile(fd)`

### `closeHandle(handle)`

### `readFileDescriptor(fd, bufferPtr, count)`

### `writeFileDescriptor(fd, bufferPtr, count)`

### `seek(fd, offset, whence)`

### `unlink(path)`

### `mmap(addr, length, prot, flags, fd, offset)`

### `munmap(addr, length)`

### `mprotect(addr, length, prot)`

### `allocPages(size, prot)`

### `freePages(address, size)`

### `protectPages(address, size, prot)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `_compatNotImplemented(apiName)`

> returns `:object`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `linux-loader`

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

## Module: `linux-windowing`

- `sys` · `import(...)`
### `_r1Positive?(res)`

> returns `:bool`

### `callOk?(res)`

> returns `:bool`

### `_zeros(n, acc)`

### `xEventSize()`

> returns `:int`

### `createXEventBuffer()`

### `xEventType(eventPtr)`

### `openDisplay(displayName)`

### `closeDisplay(display)`

### `defaultScreen(display)`

### `rootWindow(display, screen)`

### `blackPixel(display, screen)`

### `whitePixel(display, screen)`

### `createSimpleWindow(display, parent, x, y, width, height, borderWidth, border, background)`

### `destroyWindow(display, window)`

### `storeName(display, window, title)`

### `selectInput(display, window, eventMask)`

### `mapWindow(display, window)`

### `unmapWindow(display, window)`

### `moveWindow(display, window, x, y)`

### `resizeWindow(display, window, width, height)`

### `displayWidth(display, screen)`

### `displayHeight(display, screen)`

### `createGC(display, drawable, valueMask, values)`

### `freeGC(display, gc)`

### `setForeground(display, gc, color)`

### `drawLine(display, window, gc, x1, y1, x2, y2)`

### `fillRectangle(display, window, gc, x, y, width, height)`

### `drawString(display, window, gc, x, y, text)`

### `flush(display)`

### `pending(display)`

### `nextEvent(display, eventPtr)`

### `_openWindowFromRoot(display, screen, root, black, white, title, width, height)`

> returns `:object`

### `_openWindowFromScreen(display, screen, title, width, height)`

### `_openWindowFromDisplay(display, title, width, height)`

### `openDefaultWindow(title, width, height)`

### `closeWindow(state)`

> returns `:int`

### `pumpWindowEvent(display, eventPtr)`

> returns `:object`

### `runWindowLoop(display, eventPtr)`

> returns `:int`

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

## Module: `writes`

### `_b0(v)`

### `_b1(v)`

### `_b2(v)`

### `_b3(v)`

### `_b4(v)`

### `_b5(v)`

### `_b6(v)`

### `_b7(v)`

### `readU32(address)`

### `writeU32(address, value)`

### `readU64(address)`

### `writeU64(address, value)`

