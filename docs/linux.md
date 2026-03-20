# Linux Interop Library (linux)

## Overview

`linux` provides Linux-oriented native bindings using Magnolia's `sys` wrapper layer.

It supports:

- resolving symbols from `libc` / `libdl`
- resolving and calling symbols from `libX11`
- process and runtime helpers (`getpid`, `getppid`, `sysconf`)
- errno and strerror helpers
- virtual memory wrappers (`mmap`, `munmap`, `mprotect`)
- convenience page allocation wrappers
- dynamic loader wrappers (`dlopen`, `dlsym`, `dlclose`)
- X11 window creation, event loops, and basic painting

All public calls are OS-gated and return structured error objects on non-Linux hosts.

## Import

Both names are supported:

```oak
linux := import('linux')
// or
linux := import('Linux')
```

## Module Map

The Linux interop surface is split into focused modules and re-exported by
`linux`/`Linux`.

- [linux-constants](linux-constants.md): constants and shared-library candidate lists
- [linux-core](linux-core.md): OS guard, C-string helpers, symbol resolution/call dispatch
- [linux-loader](linux-loader.md): `dlopen`/`dlsym` helpers and cached call-by-symbol
- [linux-windowing](linux-windowing.md): X11 display/window/event/drawing helpers
- [linux-libc](linux-libc.md): process, errno, file descriptor, and virtual-memory wrappers

You can import these modules directly when you only need part of the Linux API:

```oak
constants := import('linux-constants')
core := import('linux-core')
loader := import('linux-loader')
x11 := import('linux-windowing')
libc := import('linux-libc')
```

## Constants

### Library candidates

- `LibC`
- `LibDL`
- `LibX11`

### Memory protection flags

- `PROT_NONE`
- `PROT_READ`
- `PROT_WRITE`
- `PROT_EXEC`

### `mmap` flags

- `MAP_SHARED`
- `MAP_PRIVATE`
- `MAP_FIXED`
- `MAP_ANONYMOUS`

### `open()` flags

- `O_RDONLY`
- `O_WRONLY`
- `O_RDWR`
- `O_CREAT`
- `O_TRUNC`
- `O_APPEND`

### `lseek()` constants

- `SEEK_SET`
- `SEEK_CUR`
- `SEEK_END`

### `access()` mode constants

- `F_OK`
- `R_OK`
- `W_OK`
- `X_OK`

### `dlopen()` flags

- `RTLD_LAZY`
- `RTLD_NOW`
- `RTLD_GLOBAL`
- `RTLD_LOCAL`

### X11 event masks and types

- `KeyPressMask`
- `ButtonPressMask`
- `ExposureMask`
- `StructureNotifyMask`
- `KeyPress`
- `Expose`
- `DestroyNotify`
- `ClientMessage`

## Helpers

### `isLinux?()`

Returns true when the host OS is Linux.

### `cstr(s)`

Appends a null terminator to `s` for C-style API calls.

### `readU32(address)` / `writeU32(address, value)`
### `readU64(address)` / `writeU64(address, value)`

Little-endian typed integer helpers built on top of `memread`/`memwrite`.

## Resolution and Dispatch

### `resolve(symbol)`

Resolves `symbol` from the first available libc candidate.

### `resolveIn(library, symbol)`

Resolves `symbol` from an explicit library.

### `call(target, args...)`

Calls a resolved proc or raw address via `sys.call`.

### `libc(symbol, args...)`
### `libdl(symbol, args...)`

Resolve + call convenience wrappers for libc/libdl candidate sets.

### `loadDll(library)`

Loads and memoizes a shared library handle using `dlopen`.

- `library` can be a single string path/name, or a candidate list.
- returns `{type: :ok, handle: <int>, library: <string>}` on success.

### `resolveInLoaded(library, symbol)`

Loads a shared library (if needed) and resolves `symbol` via `dlsym`.

### `callIn(library, symbol, args...)`

Load + resolve + call helper for Linux shared libraries.

### `x11(symbol, args...)`

Calls `symbol` from the first available `libX11` candidate.

## Process and Runtime APIs

### `currentProcessId()`

Calls `getpid` (`-1` on non-Linux).

### `currentProcess()`

Compatibility alias for `currentProcessId()`.

### `moduleHandle(name)`

Compatibility helper for Windows-style module-handle lookup:

- pass `?` to get the current process image handle via `dlopen(NULL, ...)`
- pass a shared library name/path to load and return a handle

### `imageBase()`

Compatibility alias for `moduleHandle(?)` value extraction.

### `parentProcessId()`

Calls `getppid` (`-1` on non-Linux).

### `pageSize()`

Calls `sysconf(_SC_PAGESIZE)` (`-1` on non-Linux).

### `errno()`

Returns the current thread-local errno value (`-1` on non-Linux or lookup failure).

### `strerror(errorCode)`

Returns a best-effort error string for a numeric errno code (or `?` on failure).

### `getLastError()`

Compatibility alias for `errno()`.

### `formatMessage(errorCode)`

Compatibility alias for `strerror(errorCode)`.

### `lastErrorMessage()`

Convenience helper for `strerror(errno())`.

### `getuid()` / `geteuid()`
### `getgid()` / `getegid()`

Returns real/effective UID and GID via libc wrappers.

### `gethostname(bufferPtr, size)`

Calls `gethostname(2)` with caller-managed memory.

### `getcwd(bufferPtr, size)`

Calls `getcwd(3)` with caller-managed memory.

### `chdir(path)`

Changes process working directory.

### `access(path, mode)`

Calls `access(2)`; use `F_OK`, `R_OK`, `W_OK`, `X_OK`.

### `openFile(path, flags, mode?)`
### `closeFile(fd)`
### `closeHandle(handle)`
### `readFileDescriptor(fd, bufferPtr, count)`
### `writeFileDescriptor(fd, bufferPtr, count)`
### `seek(fd, offset, whence)`
### `unlink(path)`

Thin wrappers over POSIX fd APIs (`open`, `close`, `read`, `write`, `lseek`,
`unlink`).

- `openFile(..., mode?)` defaults mode to decimal `420` (`0644`) when omitted.
- `closeHandle(handle)` is a compatibility alias for `closeFile(handle)`.

## Virtual Memory APIs

### `mmap(addr, length, prot, flags, fd, offset)`
### `munmap(addr, length)`
### `mprotect(addr, length, prot)`

Thin wrappers over libc memory APIs.

### `allocPages(size, prot?)`

Convenience wrapper over `mmap` with anonymous private mapping.

- defaults `prot` to `PROT_READ | PROT_WRITE` when omitted
- uses `MAP_PRIVATE | MAP_ANONYMOUS`

### `freePages(address, size)`

Convenience wrapper over `munmap`.

### `protectPages(address, size, prot)`

Convenience wrapper over `mprotect`.

### `virtualAlloc(baseAddress, size, allocationType, protection)`
### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`
### `virtualFree(address, size, freeType)`
### `virtualFreeEx(process, address, size, freeType)`
### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`
### `virtualQuery(address, mbiBufferPtr, mbiSize)`
### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`
### `openProcess(desiredAccess, inheritHandle, processId)`
### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`
### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

Compatibility aliases for Windows-style virtual-memory APIs:

- `virtualAlloc(...)` maps to `allocPages(size, protection)`
- `virtualFree(...)` maps to `freePages(address, size)`
- `virtualProtect(...)` maps to `protectPages(address, size, newProtect)`
- `virtualAllocEx(...)`, `virtualFreeEx(...)`, `virtualQuery(...)`, `virtualQueryEx(...)`, `openProcess(...)`, `readProcessMemory(...)`, and `writeProcessMemory(...)` are exported compatibility stubs and currently return `{type: :error, ...}` on Linux.
- Stub errors include `api`, `platform`, and `supported` fields for consistent cross-platform capability checks.

## Dynamic Loader APIs

### `dlopen(path, flags)`
### `dlsym(handle, symbol)`
### `dlclose(handle)`

Thin wrappers over dynamic loader APIs.

### `loadLibrary(path)`
### `procAddress(module, symbol)`
### `freeLibrary(handle)`

Compatibility aliases for Windows-style loader naming:

- `loadLibrary(path)` calls `dlopen(path, RTLD_NOW | RTLD_LOCAL)`
- `procAddress(module, symbol)` calls `dlsym(module, symbol)`
- `freeLibrary(handle)` calls `dlclose(handle)`

## X11 Windowing and Drawing APIs

### Basic display and window APIs

### `openDisplay(displayName?)`
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

### Drawing APIs

### `createGC(display, drawable, valueMask, values)`
### `freeGC(display, gc)`
### `setForeground(display, gc, color)`
### `drawLine(display, window, gc, x1, y1, x2, y2)`
### `fillRectangle(display, window, gc, x, y, width, height)`
### `drawString(display, window, gc, x, y, text)`
### `flush(display)`

### Event APIs and loop helpers

### `pending(display)`
### `nextEvent(display, eventPtr)`
### `xEventSize()`
### `createXEventBuffer()`
### `xEventType(eventPtr)`
### `pumpWindowEvent(display, eventPtr)`
### `runWindowLoop(display, eventPtr)`

`runWindowLoop(...)` exits with `0` on `ClientMessage` or `DestroyNotify`.

### High-level window helpers

### `openDefaultWindow(title, width, height)`

Creates a simple top-level X11 window and returns:

```oak
{
    type: :ok
    display: <int>
    window: <int>
    screen: <int>
    black: <int>
    white: <int>
}
```

### `closeWindow(state)`

Destroys/tears down a window created by `openDefaultWindow(...)`.

## Example

```oak
linux := import('linux')

if linux.isLinux?() {
    true -> {
        pid := linux.currentProcessId()
        pagesize := linux.pageSize()
        println('PID: ' + string(pid))
        println('Page size: ' + string(pagesize))
        println('ENOENT text: ' + string(linux.strerror(2)))
    }
    _ -> println('linux library is inactive on this host')
}
```

## X11 Window Example

See [samples/linux-window.oak](../samples/linux-window.oak) for a runnable X11
window sample and [samples/linux-draw.oak](../samples/linux-draw.oak) for basic
drawing.

```oak
linux := import('Linux')

if linux.isLinux?() {
    true -> {
        win := linux.openDefaultWindow('Magnolia Linux Window', 800, 480)
        if win.type = :ok {
            true -> {
                eventBuf := linux.createXEventBuffer()
                linux.runWindowLoop(win.display, addr(eventBuf))
                linux.closeWindow(win)
            }
        }
    }
}
```

## Notes

- Interop calls are unsafe by nature; validate pointers and lengths.
- Symbol availability varies by distro and runtime environment.
- Keep wrappers narrow and validate arguments near call boundaries.
- X11 calls require a running X server and a valid `DISPLAY`.

## Related

- [System Interop Wrappers](./sys.md)
- [Go Runtime and System Interop](./go.md)
- [Thread Library](./thread.md)
