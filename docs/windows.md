# Windows Interop Library (windows)

## Overview

`windows` provides thin Win32 bindings on top of Magnolia's `sys` wrappers.

It is intended for:

- resolving and calling `kernel32.dll` / `ntdll.dll` / `psapi.dll` exports
- querying process and module handles
- basic virtual memory and remote process memory APIs

All public calls are OS-gated and return structured error objects on non-Windows hosts.

## Import

```oak
windows := import('windows')
```

## Constants

### DLL names

- `Kernel32`
- `Ntdll`
- `Psapi`

### Process access flags

- `PROCESS_TERMINATE`
- `PROCESS_VM_READ`
- `PROCESS_VM_WRITE`
- `PROCESS_VM_OPERATION`
- `PROCESS_QUERY_INFORMATION`
- `PROCESS_QUERY_LIMITED_INFORMATION`
- `PROCESS_ALL_ACCESS`

### Memory flags

- `MEM_COMMIT`
- `MEM_RESERVE`
- `MEM_DECOMMIT`
- `MEM_RELEASE`

### Page protection flags

- `PAGE_NOACCESS`
- `PAGE_READONLY`
- `PAGE_READWRITE`
- `PAGE_EXECUTE`
- `PAGE_EXECUTE_READ`
- `PAGE_EXECUTE_READWRITE`

## Helpers

### `isWindows?()`

Returns true when the host OS is Windows.

### `wstr(s)`

Converts a string to UTF-16 bytes with trailing null terminator (for `*W` APIs).

### `cstr(s)`

Converts a string to ANSI bytes with trailing null terminator.

## Resolution and Dispatch

### `resolve(symbol)`

Resolves `symbol` from `kernel32.dll`.

### `resolveIn(library, symbol)`

Resolves `symbol` from a specific library.

### `call(target, args...)`

Calls a resolved proc or address through `sys.call`.

### `kernel32(symbol, args...)`
### `ntdll(symbol, args...)`
### `psapi(symbol, args...)`

Resolve + call convenience wrappers for each library.

## Process and Module APIs

### `getLastError()`

Returns Win32 `GetLastError()` value (or `-1` on non-Windows).

### `currentProcessId()`

Returns `GetCurrentProcessId()` (or `-1` on non-Windows).

### `currentProcess()`

Returns `GetCurrentProcess()` pseudo-handle (or `0` on non-Windows).

### `moduleHandle(name)`

Calls `GetModuleHandleW`.

- pass `?` to get the current executable module handle
- pass a DLL/module name string to resolve that module

### `imageBase()`

Convenience helper for current executable image base (`moduleHandle(?)`).

## Library Management APIs

### `loadLibrary(path)`

Calls `LoadLibraryW`.

### `freeLibrary(module)`

Calls `FreeLibrary`.

### `procAddress(module, symbol)`

Calls `GetProcAddress`.

## Handle and Memory APIs

### `openProcess(desiredAccess, inheritHandle, processId)`
### `closeHandle(handle)`

Wrap `OpenProcess` and `CloseHandle`.

### `virtualAlloc(baseAddress, size, allocationType, protection)`
### `virtualFree(address, size, freeType)`
### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

Wrap virtual memory management APIs.

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`
### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`
### `virtualQuery(address, mbiBufferPtr, mbiSize)`

Wrap process memory and memory region query APIs.

## Example

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        pid := windows.currentProcessId()
        base := windows.imageBase()
        println('PID: ' + string(pid))
        println('Image base: ' + string(base))
    }
    _ -> println('windows library is inactive on this host')
}
```

## Notes

- Interop calls are unsafe by nature; validate pointers and sizes.
- A successful syscall-like response can be `:ok` or `:success`.
- Prefer building higher-level wrappers for application logic.

## Related

- [System Interop Wrappers](./sys.md)
- [Go Runtime and System Interop](./go.md)
- [Thread Library](./thread.md)
