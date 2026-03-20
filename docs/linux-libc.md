# Linux Libc Module (linux-libc)

## Overview

`linux-libc` groups process/runtime, fd, path, and virtual-memory helpers that
call into libc on Linux hosts.

## Import

```oak
libc := import('linux-libc')
```

## Process and Runtime APIs

### `currentProcessId()`
### `parentProcessId()`
### `pageSize()`

Convenience wrappers for `getpid`, `getppid`, and `sysconf(_SC_PAGESIZE)`.

### `errno()`
### `strerror(errorCode)`
### `lastErrorMessage()`

Thread-local error accessors.

### `getuid()` / `geteuid()`
### `getgid()` / `getegid()`

Real/effective UID and GID wrappers.

## Path and FD APIs

### `gethostname(bufferPtr, size)`
### `getcwd(bufferPtr, size)`
### `chdir(path)`
### `access(path, mode)`

Host/path query helpers.

### `openFile(path, flags, mode?)`
### `closeFile(fd)`
### `readFileDescriptor(fd, bufferPtr, count)`
### `writeFileDescriptor(fd, bufferPtr, count)`
### `seek(fd, offset, whence)`
### `unlink(path)`

Thin wrappers over POSIX-style file descriptor operations.

- `openFile(..., mode?)` defaults `mode` to decimal `420` (`0644`) when omitted.

## Virtual Memory APIs

### `mmap(addr, length, prot, flags, fd, offset)`
### `munmap(addr, length)`
### `mprotect(addr, length, prot)`

Direct libc mappings.

### `allocPages(size, prot?)`

Convenience wrapper around `mmap` using anonymous private mapping.

- default protection: `PROT_READ | PROT_WRITE`
- mapping flags: `MAP_PRIVATE | MAP_ANONYMOUS`

### `freePages(address, size)`
### `protectPages(address, size, prot)`

Convenience wrappers over `munmap` and `mprotect`.

## Return Behavior

- On non-Linux hosts, OS-gated calls return `-1`, `?`, or a structured
  `{type: :error, ...}` object depending on the API.
- On Linux hosts, results mirror underlying syscall/libc semantics via Magnolia's
  `sys` interop result shape.
