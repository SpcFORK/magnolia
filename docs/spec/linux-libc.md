# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\linux-libc.oak`

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

