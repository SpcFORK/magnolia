# Windows Interop Library (windows)

## Overview

`windows` provides thin Win32 bindings on top of Magnolia's `sys` wrappers.

It is intended for:

- resolving and calling `kernel32.dll` / `ntdll.dll` / `psapi.dll` exports
- resolving and calling GUI/system DLL exports like `user32.dll` / `gdi32.dll`
- querying process and module handles
- basic virtual memory and remote process memory APIs

All public calls are OS-gated and return structured error objects on non-Windows hosts.

## Split modules

The `windows` facade is composed of these focused modules:

- [windows-constants](windows-constants.md)
- [windows-flags](windows-flags.md)
- [windows-core](windows-core.md)
- [windows-loader](windows-loader.md)
- [windows-kernel](windows-kernel.md)
- [windows-gdi](windows-gdi.md)
- [windows-windowing](windows-windowing.md)
- [windows-net](windows-net.md)
- [windows-registry](windows-registry.md)

## Import

```oak
windows := import('windows')
```

## Constants

### DLL names

- `Kernel32`
- `Ntdll`
- `Psapi`
- `User32`
- `Gdi32`
- `Advapi32`
- `Shell32`
- `Ole32`
- `Ws2_32`
- `Comctl32`
- `Wininet`
- `D3d9`
- `Ddraw`
- `Vulkan1`
- `Msvcrt`
- `Ucrtbase`
- `Vcruntime140`
- `ActionCenter`
- `Aclui`, `Acledit`
- `Acppage`, `Acproxy`
- `Adprovider`
- `Aeinv`, `Aepic`
- `Amstream`
- `Adsldp`, `Adsnt`
- `Adtschema`, `Adsldpc`, `Adsmsext`
- `Adhsvc`
- `Advapi32res`, `Advpack`
- `Aeevts`
- `Apds`
- `Winhttp`, `Urlmon`
- `Crypt32`, `Bcrypt`, `Secur32`
- `Comdlg32`, `Oleaut32`
- `Imm32`
- `Shlwapi`, `Shcore`
- `UxTheme`, `Dwmapi`
- `Version`, `Setupapi`, `Netapi32`
- `Winmm`, `Avrt`, `Mmdevapi`, `Dsound`
- `Mfplat`, `Mfreadwrite`, `Mfuuid`
- `Taskschd`, `Wevtapi`
- `Wlanapi`, `Mpr`, `Spoolss`, `Wtsapi32`, `Rasapi32`
- `Msi`, `Wimgapi`, `Cabinet`, `Apphelp`
- `Wer`, `Faultrep`, `Dbghelp`, `Dbgeng`
- `Pdh`, `Iphlpapi`, `Wscapi`, `Sensapi`
- `Ncrypt`, `Cryptui`, `Wintrust`, `Samlib`
- `Netshell`, `Fwpuclnt`, `Dnsapi`, `Nlaapi`, `Httpapi`
- `Rpcrt4`, `Srpapi`, `Sxs`
- `Msvcirt`
- `ApiSetPrefix` (prefix for ApiSet compatibility stubs like `api-ms-win-*.dll`)
- `D3dx9Prefix` (prefix for legacy `d3dx9_*.dll` helper DLL family)
- `MsvcpPrefix`, `VcruntimePrefix`, `AtlPrefix`, `MfcPrefix`, `VcompPrefix`

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

### FormatMessage flags

- `FORMAT_MESSAGE_IGNORE_INSERTS`
- `FORMAT_MESSAGE_FROM_SYSTEM`
- `ERROR_SUCCESS`

### Winsock constants

- `AF_INET`
- `SOCK_STREAM`, `SOCK_DGRAM`
- `IPPROTO_TCP`, `IPPROTO_UDP`
- `SOCKET_ERROR`, `INVALID_SOCKET`
- `SD_RECEIVE`, `SD_SEND`, `SD_BOTH`

### WinINet constants

- `INTERNET_OPEN_TYPE_PRECONFIG`
- `INTERNET_OPEN_TYPE_DIRECT`
- `INTERNET_OPEN_TYPE_PROXY`
- `INTERNET_DEFAULT_HTTP_PORT`
- `INTERNET_DEFAULT_HTTPS_PORT`
- `INTERNET_SERVICE_HTTP`

### Registry constants

- `HKEY_CLASSES_ROOT`
- `HKEY_CURRENT_USER`
- `HKEY_LOCAL_MACHINE`
- `HKEY_USERS`
- `HKEY_CURRENT_CONFIG`
- `KEY_QUERY_VALUE`, `KEY_SET_VALUE`, `KEY_CREATE_SUB_KEY`, `KEY_ENUMERATE_SUB_KEYS`, `KEY_READ`
- `KEY_WRITE`
- `REG_SZ`, `REG_DWORD`, `REG_QWORD`

### Window class/style/message constants

- `CS_VREDRAW`, `CS_HREDRAW`
- `WS_OVERLAPPED`, `WS_CAPTION`, `WS_SYSMENU`, `WS_THICKFRAME`
- `WS_MINIMIZEBOX`, `WS_MAXIMIZEBOX`, `WS_VISIBLE`, `WS_CLIPSIBLINGS`, `WS_CLIPCHILDREN`, `WS_OVERLAPPEDWINDOW`
- `CW_USEDEFAULT`
- `WM_CREATE`, `WM_DESTROY`, `WM_PAINT`, `WM_CLOSE`, `WM_QUIT`, `WM_COMMAND`
- `SW_HIDE`, `SW_SHOW`
- `PM_NOREMOVE`, `PM_REMOVE`
- `MB_OK`, `MB_ICONERROR`, `MB_ICONWARNING`, `MB_ICONINFORMATION`
- `IDC_ARROW`, `IDI_APPLICATION`

## Helpers

### `isWindows?()`

Returns true when the host OS is Windows.

### `wstr(s)`

Converts a string to UTF-16 bytes with trailing null terminator (for `*W` APIs).

### `cstr(s)`

Converts a string to ANSI bytes with trailing null terminator.

### `makeWord(low, high)`

Builds a 16-bit value commonly used for Winsock version negotiation. Example:
`makeWord(2, 2)` for Winsock 2.2.

## Resolution and Dispatch

### `resolve(symbol)`

Resolves `symbol` from `kernel32.dll`.

### `resolveIn(library, symbol)`

Resolves `symbol` from a specific library.

### `call(target, args...)`

Calls a resolved proc or address through `sys.call`.

### `kernel32(symbol, args...)`
### `ntdll(symbol, args...)`
### `ntNative(symbol, args...)`
### `psapi(symbol, args...)`

Resolve + call convenience wrappers for each library.

### `loadDll(library)`

Loads a DLL by name/path using `LoadLibraryW` and memoizes the module handle.

Returns:

- `{type: :ok, handle: <int>, library: <string>}` on success
- `{type: :error, ...}` on failure

### `resolveInLoaded(library, symbol)`

Ensures a DLL is loaded, then resolves a symbol via `GetProcAddress`.

Returns:

- `{type: :ok, proc: <int>, handle: <int>, library: <string>, symbol: <string>}`
- `{type: :error, ...}`

### `callIn(library, symbol, args...)`

Load + resolve + call helper for arbitrary DLLs.

### `user32(symbol, args...)`
### `gdi32(symbol, args...)`
### `advapi32(symbol, args...)`
### `shell32(symbol, args...)`
### `ole32(symbol, args...)`
### `ws2_32(symbol, args...)`
### `comctl32(symbol, args...)`
### `wininet(symbol, args...)`
### `d3d9(symbol, args...)`
### `ddraw(symbol, args...)`
### `vulkan1(symbol, args...)`
### `d3dx9Dll(suffix)`
### `d3dx9(suffix, symbol, args...)`
### `msvcrt(symbol, args...)`
### `ucrtbase(symbol, args...)`
### `vcruntime140(symbol, args...)`
### `actionCenter(symbol, args...)`
### `aclui(symbol, args...)`
### `acledit(symbol, args...)`
### `acppage(symbol, args...)`
### `acproxy(symbol, args...)`
### `adprovider(symbol, args...)`
### `aeinv(symbol, args...)`
### `aepic(symbol, args...)`
### `amstream(symbol, args...)`
### `adsldp(symbol, args...)`
### `adsnt(symbol, args...)`
### `adtschema(symbol, args...)`
### `adsldpc(symbol, args...)`
### `adsmsext(symbol, args...)`
### `adhsvc(symbol, args...)`
### `advapi32res(symbol, args...)`
### `advpack(symbol, args...)`
### `aeevts(symbol, args...)`
### `apds(symbol, args...)`
### `winhttp(symbol, args...)`
### `urlmon(symbol, args...)`
### `crypt32(symbol, args...)`
### `bcrypt(symbol, args...)`
### `secur32(symbol, args...)`
### `comdlg32(symbol, args...)`
### `oleaut32(symbol, args...)`
### `imm32(symbol, args...)`
### `shlwapi(symbol, args...)`
### `shcore(symbol, args...)`
### `uxTheme(symbol, args...)`
### `dwmapi(symbol, args...)`
### `versionDll(symbol, args...)`
### `setupapi(symbol, args...)`
### `netapi32(symbol, args...)`
### `winmm(symbol, args...)`
### `avrt(symbol, args...)`
### `mmdevapi(symbol, args...)`
### `dsound(symbol, args...)`
### `mfplat(symbol, args...)`
### `mfreadwrite(symbol, args...)`
### `mfuuid(symbol, args...)`
### `taskschd(symbol, args...)`
### `wevtapi(symbol, args...)`
### `wlanapi(symbol, args...)`
### `mpr(symbol, args...)`
### `spoolss(symbol, args...)`
### `wtsapi32(symbol, args...)`
### `rasapi32(symbol, args...)`
### `msi(symbol, args...)`
### `wimgapi(symbol, args...)`
### `cabinet(symbol, args...)`
### `apphelp(symbol, args...)`
### `wer(symbol, args...)`
### `faultrep(symbol, args...)`
### `dbghelp(symbol, args...)`
### `dbgeng(symbol, args...)`
### `pdh(symbol, args...)`
### `iphlpapi(symbol, args...)`
### `wscapi(symbol, args...)`
### `sensapi(symbol, args...)`
### `ncrypt(symbol, args...)`
### `cryptui(symbol, args...)`
### `wintrust(symbol, args...)`
### `samlib(symbol, args...)`
### `netshell(symbol, args...)`
### `fwpuclnt(symbol, args...)`
### `dnsapi(symbol, args...)`
### `nlaapi(symbol, args...)`
### `httpapi(symbol, args...)`
### `rpcrt4(symbol, args...)`
### `srpapi(symbol, args...)`
### `sxs(symbol, args...)`
### `msvcirt(symbol, args...)`

### `msvcpDll(suffix)` / `msvcpFamily(suffix, symbol, args...)`
### `vcruntimeDll(suffix)` / `vcruntimeFamily(suffix, symbol, args...)`
### `atlDll(suffix)` / `atlFamily(suffix, symbol, args...)`
### `mfcDll(suffix)` / `mfcFamily(suffix, symbol, args...)`
### `vcompDll(suffix)` / `vcompFamily(suffix, symbol, args...)`

Dynamic helpers for runtime library families such as `msvcp*.dll`,
`vcruntime*.dll`, `atl*.dll`, `mfc*.dll`, and `vcomp*.dll`.

- pass a numeric/variant suffix like `'140'` or `'140_1'` to build names
- pass a full filename ending in `.dll` to use it directly

### `apiSetDll(contract)`
### `apiSet(contract, symbol, args...)`

Convenience wrappers that route through `callIn(...)`.

## Added Feature Guide

This section focuses on practical usage for the expanded Windows surface.

### 1) Safe Availability Pattern For Optional DLLs

Some DLLs are optional by Windows edition, installed feature, or runtime package.
Probe first, then call APIs only when loaded.

```oak
windows := import('windows')

probe := windows.loadDll(windows.Taskschd)
if probe.type = :ok {
    true -> {
        println('taskschd loaded: ' + string(probe.handle))
    }
    _ -> {
        # Commonly 126 when the module is unavailable on this machine.
        println('taskschd unavailable: ' + string(probe))
    }
}
```

### 2) Generic Dispatch For Any Export

Use `callIn(dll, symbol, args...)` for quick one-off calls without creating a
new wrapper.

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        # Beep(DWORD freq, DWORD duration)
        beep := windows.callIn(windows.Kernel32, 'Beep', 880, 120)
        println(string(beep))
    }
}
```

### 3) Runtime Family Helpers (`msvcp*`, `vcruntime*`, `atl*`, `mfc*`, `vcomp*`)

Use `*Dll(suffix)` to build a DLL name and `*Family(suffix, symbol, args...)`
to dispatch through that generated filename.

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        println(windows.msvcpDll('140'))       # msvcp140.dll
        println(windows.vcruntimeDll('140'))   # vcruntime140.dll
        println(windows.atlDll('100'))         # atl100.dll
        println(windows.mfcDll('140'))         # mfc140.dll
        println(windows.vcompDll('140'))       # vcomp140.dll

        # Example: ask any runtime family DLL for DllCanUnloadNow when available.
        unload := windows.msvcpFamily('140', 'DllCanUnloadNow')
        println(string(unload))
    }
}
```

You can also pass a full filename (ending in `.dll`) as the suffix argument.

### 4) ApiSet Contract Helpers

Use `apiSetDll(contract)` for contract name generation and `apiSet(...)` for
direct call attempts.

```oak
windows := import('windows')

name := windows.apiSetDll('file-l1-1-0')
println(name) # api-ms-win-core-file-l1-1-0.dll

probe := windows.loadDll(name)
println(string(probe))
```

### 5) D3DX9 Family Helpers

Use `d3dx9Dll(suffix)` for `d3dx9_*.dll` names and `d3dx9(...)` for dispatch.

```oak
windows := import('windows')

println(windows.d3dx9Dll('43'))
attempt := windows.d3dx9('43', 'D3DXCheckVersion', 32, 43)
println(string(attempt))
```

### 6) Winsock + IPv4 Struct Convenience

Use `sockaddrIn(ip, port)` to build a compatible `sockaddr_in` byte buffer for
`connectSocket` and `bindSocket`.

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        wsaData := bits([0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0])
        startup := windows.wsaStartup(windows.makeWord(2, 2), addr(wsaData))
        if windows.callOk?(startup) {
            true -> {
                sock := windows.socket(windows.AF_INET, windows.SOCK_STREAM, windows.IPPROTO_TCP)
                sa := windows.sockaddrIn('93.184.216.34', 80)
                if sa.type = :ok {
                    true -> {
                        conn := windows.connectSocket(sock.r1, sa.ptr, sa.len)
                        println(string(conn))
                    }
                }
                windows.closeSocket(sock.r1)
                windows.wsaCleanup()
            }
        }
    }
}
```

### 7) WinINet One-Liner HTTP Fetch

Use `internetSimpleGet(url, agent?, chunkSize?)` for quick text fetches.

```oak
windows := import('windows')

result := windows.internetSimpleGet('https://example.com', 'Magnolia', 4096)
println(string(result))
```

### 8) Registry Read/Write Helpers

Use `regReadString` / `regReadDword` for reads and `regWriteString` /
`regWriteDword` for writes.

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        keyPath := 'Software\\MagnoliaDemo'

        writeN := windows.regWriteDword(windows.HKEY_CURRENT_USER, keyPath, 'Counter', 7)
        writeS := windows.regWriteString(windows.HKEY_CURRENT_USER, keyPath, 'Label', 'hello')
        println(string(writeN))
        println(string(writeS))

        readN := windows.regReadDword(windows.HKEY_CURRENT_USER, keyPath, 'Counter')
        readS := windows.regReadString(windows.HKEY_CURRENT_USER, keyPath, 'Label')
        println(string(readN))
        println(string(readS))
    }
}
```

### 9) Broad DLL Wrapper Surface

For explicitly wrapped libraries (`imm32`, `msvcirt`, `dnsapi`, `httpapi`,
`netapi32`, `mmdevapi`, `dbghelp`, and many others), usage is always the same:

```oak
windows := import('windows')

# ImmGetContext(HWND)
imm := windows.imm32('ImmGetContext', 0)
println(string(imm))

# Example no-arg call through another wrapper
dbg := windows.dbghelp('SymCleanup', 0)
println(string(dbg))
```

Use this pattern to keep module-specific code small:

1. Probe with `loadDll(...)` when availability may vary.
2. Call through the specific wrapper (`dnsapi`, `taskschd`, `wlanapi`, etc.).
3. Handle both syscall transport success and API-level result codes.

## Winsock APIs (ws2_32)

### `wsLastError()`
### `wsaStartup(version, wsaDataPtr)`
### `wsaCleanup()`

Core Winsock lifecycle helpers.

### `socket(af, socketType, protocol)`
### `bindSocket(sock, sockaddrPtr, sockaddrLen)`
### `connectSocket(sock, sockaddrPtr, sockaddrLen)`
### `listenSocket(sock, backlog)`
### `acceptSocket(sock, addrOutPtr, addrLenInOutPtr)`
### `sendSocket(sock, bufferPtr, size, flags)`
### `recvSocket(sock, bufferPtr, size, flags)`
### `shutdownSocket(sock, how)`
### `closeSocket(sock)`

Thin socket wrappers over WinSock calls.

### `htons(value)`
### `htonl(value)`
### `inetAddr(ipv4)`
### `sockaddrIn(ipv4, port)`

Network byte-order and IPv4 conversion helpers.

`sockaddrIn` returns `{type: :ok, buffer, ptr, len}` for passing to
`connectSocket`/`bindSocket`.

## WinINet APIs (wininet)

### `internetOpen(agent, accessType, proxy, proxyBypass, flags)`
### `internetConnect(hInternet, serverName, serverPort, username, password, service, flags, context)`
### `internetOpenUrl(hInternet, url, headers, headersLen, flags, context)`
### `internetReadFile(hFile, outBufferPtr, bytesToRead, bytesReadOutPtr)`
### `internetCloseHandle(hInternet)`
### `internetSimpleGet(url, agent?, chunkSize?)`

Thin wrappers for high-level HTTP/FTP style network handles.

## Registry APIs (advapi32)

### `statusOk?(res)`

Returns true when a syscall result is successful and has `ERROR_SUCCESS` status
(`r1 = 0`), which is common for Win32 registry functions.

### `regCloseKey(hKey)`
### `regOpenKeyEx(rootKey, subKey, options, samDesired, outKeyPtr)`
### `regCreateKeyEx(rootKey, subKey, reserved, className, options, samDesired, securityAttributesPtr, outKeyPtr, dispositionOutPtr)`
### `regQueryValueEx(hKey, valueName, reserved, typeOutPtr, dataOutPtr, dataLenInOutPtr)`
### `regSetValueEx(hKey, valueName, reserved, valueType, dataPtr, dataLen)`
### `regDeleteValue(hKey, valueName)`
### `regReadDword(rootKey, subKey, valueName)`
### `regWriteDword(rootKey, subKey, valueName, value)`
### `regReadString(rootKey, subKey, valueName)`
### `regWriteString(rootKey, subKey, valueName, value)`

Thin registry wrappers for opening, creating, querying, setting, and deleting
registry values.

## Windowing APIs (user32)

### `registerClassEx(wndClassExPtr)`

Calls `RegisterClassExW`.

### `createWindowEx(exStyle, className, windowName, style, x, y, width, height, parent, menu, instance, param)`

Calls `CreateWindowExW` with UTF-16 conversion for class and window names.

### `defWindowProc(hwnd, msg, wParam, lParam)`
### `showWindow(hwnd, cmdShow)`
### `updateWindow(hwnd)`
### `destroyWindow(hwnd)`
### `postQuitMessage(exitCode)`

Core window lifecycle helpers.

### `getMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax)`
### `peekMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax, removeMsg)`
### `translateMessage(msgPtr)`
### `dispatchMessage(msgPtr)`
### `isWindow(hwnd)`
### `waitMessage()`

Message-loop helpers. `msgPtr` should point to a `MSG`-compatible buffer.

### `callOk?(res)`

Returns true for a successful syscall result, and also treats some Win32
interop cases with `r1 > 0` as truthy.

### `noMessage?(res)`

Returns true when a peek/get call produced no queued message.

### `windowAlive?(hwnd)`

Returns true when `IsWindow(hwnd)` indicates the handle is still valid.

### `msgStructSize()`

Returns platform-correct `MSG` struct byte size (`48` on 64-bit, `28` on
32-bit targets).

### `createMsgBuffer()`

Allocates and returns a zero-initialized `MSG`-compatible byte buffer.

### `pumpWindowMessage(hwnd, msgPtr)`

Runs one non-blocking `PeekMessageW` loop iteration and returns one of:

- `{type: :dispatch, detail: ...}` when a message was dispatched
- `{type: :idle, detail: ...}` when no message was pending
- `{type: :closed}` when the window is no longer valid
- `{type: :error, ...}` on unexpected call failure

### `runWindowLoopPeek(hwnd, msgPtr)`

Runs a close-aware `PeekMessageW` + `WaitMessage` loop until `hwnd` closes.
Returns `0` when the window closes, or an error object.

### `registerDefaultWindowClass(className)`

Registers a default top-level `WNDCLASSEXW` using `DefWindowProcW`.

This is a convenience helper for Magnolia samples where no custom Oak WNDPROC
callback is available.

### `createTopLevelWindow(className, title, width, height)`

Convenience helper that creates a visible overlapped top-level window using
`CreateWindowExW`.

### `runWindowLoop(hwnd)`

Runs the native `win_msg_loop` builtin for `hwnd` on the current thread.

For Win32 UI code, pair this with `lock_thread()` / `unlock_thread()` so
window creation and message dispatch stay on the same OS thread.

### `messageBox(hwnd, text, caption, msgType)`
### `setWindowText(hwnd, text)`
### `loadCursor(instance, cursorId)`
### `loadIcon(instance, iconId)`

Common Win32 UI helper calls.

### `beginPaint(hwnd, paintStructPtr)`
### `endPaint(hwnd, paintStructPtr)`
### `getDC(hwnd)`
### `releaseDC(hwnd, hdc)`

Painting and DC access helpers.

## Basic GDI APIs (gdi32)

### `getStockObject(objectIndex)`
### `selectObject(hdc, gdiObject)`
### `textOut(hdc, x, y, text)`
### `rectangle(hdc, left, top, right, bottom)`
### `ellipse(hdc, left, top, right, bottom)`
### `createSolidBrush(colorRef)`
### `deleteObject(gdiObject)`

Thin wrappers for common drawing primitives.

## Process and Module APIs

### `getLastError()`

Returns Win32 `GetLastError()` value (or `-1` on non-Windows).

### `formatMessage(errorCode)`

Best-effort user-readable message for a Win32 error code via `FormatMessageW`.

Returns a string on success and `?` on failure/non-Windows hosts.

### `lastErrorMessage()`

Convenience helper for `formatMessage(getLastError())`.

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

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`
### `virtualFreeEx(process, address, size, freeType)`
### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

Remote-process variants of virtual memory management/query APIs.

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`
### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`
### `virtualQuery(address, mbiBufferPtr, mbiSize)`

Wrap process memory and memory region query APIs.

### `readU32(address)` / `writeU32(address, value)`
### `readU64(address)` / `writeU64(address, value)`

Little-endian typed integer helpers built on top of `memread`/`memwrite`.

## Example

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        pid := windows.currentProcessId()
        base := windows.imageBase()
        println('PID: ' + string(pid))
        println('Image base: ' + string(base))
        println('Error 2 text: ' + string(windows.formatMessage(2)))
    }
    _ -> println('windows library is inactive on this host')
}
```

### Window Creation Example

See [samples/windows-window.oak](../samples/windows-window.oak) for a runnable
window sample.

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        lock_thread()

        className := 'MagnoliaWindowClass'
        windows.registerDefaultWindowClass(className)
        hwnd := windows.createTopLevelWindow(className, 'Magnolia Win32 Window', 800, 480)

        if hwnd.type = :ok & hwnd.r1 > 0 {
            true -> {
                windows.showWindow(hwnd.r1, windows.SW_SHOW)
                windows.updateWindow(hwnd.r1)

                msgBuf := windows.createMsgBuffer()
                windows.runWindowLoopPeek(hwnd.r1, addr(msgBuf))
            }
        }

        unlock_thread()
    }
}
```

### Drawing Example (Immediate GDI)

See [samples/windows-draw.oak](../samples/windows-draw.oak) for a runnable
example that:

- creates a visible top-level window
- draws text + simple shapes using `getDC`, `textOut`, `rectangle`, and `ellipse`
- enters a close-aware `PeekMessage`/`DispatchMessage` loop

### Networking Example (Winsock + WinINet)

See [samples/windows-network.oak](../samples/windows-network.oak) for a runnable
network sample that:

- initializes Winsock with `wsaStartup(makeWord(2, 2), ...)`
- builds a `sockaddr_in` struct via `sockaddrIn(ip, port)`
- performs a simple HTTP GET via `internetSimpleGet(...)`

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        wsaData := bits([...])
        startup := windows.wsaStartup(windows.makeWord(2, 2), addr(wsaData))
        if windows.callOk?(startup) {
            true -> {
                response := windows.internetSimpleGet('https://example.com', 'Magnolia', 2048)
                println(string(response))
                windows.wsaCleanup()
            }
        }
    }
}
```

### Registry Example (Advapi32 Helpers)

See [samples/windows-registry.oak](../samples/windows-registry.oak) for a
runnable registry sample that:

- reads a known system value with `regReadString(...)`
- writes and reads back `REG_DWORD` and `REG_SZ` values under `HKEY_CURRENT_USER`

```oak
windows := import('windows')

if windows.isWindows?() {
    true -> {
        product := windows.regReadString(
            windows.HKEY_LOCAL_MACHINE
            'SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion'
            'ProductName'
        )
        println(string(product))
    }
}
```

### DLL Coverage Example (Requested Set)

See [samples/windows-dll-bindings.oak](../samples/windows-dll-bindings.oak) for
a runnable sample that:

- probes the full explicitly requested DLL set using `loadDll(...)`
- prints counts of loaded vs failed modules
- classifies `errno = 126` as unavailable modules (often SKU/feature-dependent)
- prints detailed failures only for non-126 hard errors

### Dynamic Family Example (ApiSet + D3DX9)

See [samples/windows-dll-families.oak](../samples/windows-dll-families.oak) for
a runnable sample that:

- generates `api-ms-win-*.dll` names through `apiSetDll(...)`
- generates `d3dx9_*.dll` names through `d3dx9Dll(...)`
- probes those generated names with `loadDll(...)`

### D3D9 Program Sample

See [samples/windows-d3d9.oak](../samples/windows-d3d9.oak) for a runnable D3D9
bootstrap sample that:

- calls `Direct3DCreate9(D3D_SDK_VERSION)` through `d3d9(...)`
- invokes `IDirect3D9` COM methods via vtable pointer dispatch
- prints adapter count and releases the D3D object

Note: `Release` can report `errno = 31` in this runtime bridge even when the
call path is otherwise valid; the sample reports this as a benign runtime quirk.

## Current Limitation

`registerClassEx(...)` is available, but Magnolia currently does not expose a
native callback bridge for passing an Oak function as `WNDPROC`.

That means custom message handling (for example explicit `WM_PAINT` handlers)
still requires runtime support beyond the current `sysproc/syscall` surface.

## Threading and UI Loop Checklist

Use this sequence for stable Win32 window behavior in Magnolia:

1. Gate execution with `windows.isWindows?()`.
2. Call `lock_thread()` before creating any UI objects.
3. Register a class (`registerDefaultWindowClass(...)` or custom class setup).
4. Create the window (`createTopLevelWindow(...)` or `createWindowEx(...)`).
5. Show and update the window (`showWindow`, then `updateWindow`).
6. Allocate a `MSG` buffer with `createMsgBuffer()`.
7. Run a message loop on the same locked thread:
    - `runWindowLoopPeek(hwnd, addr(msgBuf))` for close-aware peek/wait flow, or
    - `runWindowLoop(hwnd)` for native built-in loop flow.
8. Exit loop when closed (`0` result for helper loops).
9. Call `unlock_thread()` when UI work is complete.

Common symptoms when steps are skipped:

- Missing `lock_thread()`: window can freeze or become non-responsive.
- No running message loop: window appears but does not process input/close.
- Loop on a different thread than creation: undefined behavior and stale handles.

## Notes

- Interop calls are unsafe by nature; validate pointers and sizes.
- A successful syscall-like response can be `:ok` or `:success`.
- Prefer building higher-level wrappers for application logic.
- On Windows UI code, lock to a single OS thread during create/show/loop flow.

## Related

- [System Interop Wrappers](./sys.md)
- [Go Runtime and System Interop](./go.md)
- [Thread Library](./thread.md)
