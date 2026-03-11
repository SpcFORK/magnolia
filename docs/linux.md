# Linux Interop Library (linux)

## Overview

`linux` provides Linux-oriented native bindings using Magnolia's `sys` wrapper layer.

It supports:

- resolving symbols from `libc` / `libdl`
- process and runtime helpers (`getpid`, `getppid`, `sysconf`)
- errno and strerror helpers
- virtual memory wrappers (`mmap`, `munmap`, `mprotect`)
- convenience page allocation wrappers
- dynamic loader wrappers (`dlopen`, `dlsym`, `dlclose`)

All public calls are OS-gated and return structured error objects on non-Linux hosts.

## Import

Both names are supported:

```oak
linux := import('linux')
// or
linux := import('Linux')
```

## Constants

### Library candidates

- `LibC`
- `LibDL`

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

### `dlopen()` flags

- `RTLD_LAZY`
- `RTLD_NOW`
- `RTLD_GLOBAL`
- `RTLD_LOCAL`

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

## Process and Runtime APIs

### `currentProcessId()`

Calls `getpid` (`-1` on non-Linux).

### `parentProcessId()`

Calls `getppid` (`-1` on non-Linux).

### `pageSize()`

Calls `sysconf(_SC_PAGESIZE)` (`-1` on non-Linux).

### `errno()`

Returns the current thread-local errno value (`-1` on non-Linux or lookup failure).

### `strerror(errorCode)`

Returns a best-effort error string for a numeric errno code (or `?` on failure).

### `lastErrorMessage()`

Convenience helper for `strerror(errno())`.

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

## Dynamic Loader APIs

### `dlopen(path, flags)`
### `dlsym(handle, symbol)`
### `dlclose(handle)`

Thin wrappers over dynamic loader APIs.

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

## Notes

- Interop calls are unsafe by nature; validate pointers and lengths.
- Symbol availability varies by distro and runtime environment.
- Keep wrappers narrow and validate arguments near call boundaries.

## Related

- [System Interop Wrappers](./sys.md)
- [Go Runtime and System Interop](./go.md)
- [Thread Library](./thread.md)
