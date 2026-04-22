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

## Module: `lib\gui-accessibility.oak`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `ROLE_SYSTEM_TITLEBAR` · `1`
- `ROLE_SYSTEM_MENUBAR` · `2`
- `ROLE_SYSTEM_SCROLLBAR` · `3`
- `ROLE_SYSTEM_GRIP` · `4`
- `ROLE_SYSTEM_SOUND` · `5`
- `ROLE_SYSTEM_CURSOR` · `6`
- `ROLE_SYSTEM_CARET` · `7`
- `ROLE_SYSTEM_ALERT` · `8`
- `ROLE_SYSTEM_WINDOW` · `9`
- `ROLE_SYSTEM_CLIENT` · `10`
- `ROLE_SYSTEM_MENUPOPUP` · `11`
- `ROLE_SYSTEM_MENUITEM` · `12`
- `ROLE_SYSTEM_TOOLTIP` · `13`
- `ROLE_SYSTEM_APPLICATION` · `14`
- `ROLE_SYSTEM_DOCUMENT` · `15`
- `ROLE_SYSTEM_PANE` · `16`
- `ROLE_SYSTEM_CHART` · `17`
- `ROLE_SYSTEM_DIALOG` · `18`
- `ROLE_SYSTEM_BORDER` · `19`
- `ROLE_SYSTEM_GROUPING` · `20`
- `ROLE_SYSTEM_SEPARATOR` · `21`
- `ROLE_SYSTEM_TOOLBAR` · `22`
- `ROLE_SYSTEM_STATUSBAR` · `23`
- `ROLE_SYSTEM_TABLE` · `24`
- `ROLE_SYSTEM_COLUMNHEADER` · `25`
- `ROLE_SYSTEM_ROWHEADER` · `26`
- `ROLE_SYSTEM_COLUMN` · `27`
- `ROLE_SYSTEM_ROW` · `28`
- `ROLE_SYSTEM_CELL` · `29`
- `ROLE_SYSTEM_LINK` · `30`
- `ROLE_SYSTEM_HELPBALLOON` · `31`
- `ROLE_SYSTEM_CHARACTER` · `32`
- `ROLE_SYSTEM_LIST` · `33`
- `ROLE_SYSTEM_LISTITEM` · `34`
- `ROLE_SYSTEM_OUTLINE` · `35`
- `ROLE_SYSTEM_OUTLINEITEM` · `36`
- `ROLE_SYSTEM_PAGETAB` · `37`
- `ROLE_SYSTEM_PROPERTYPAGE` · `38`
- `ROLE_SYSTEM_INDICATOR` · `39`
- `ROLE_SYSTEM_GRAPHIC` · `40`
- `ROLE_SYSTEM_STATICTEXT` · `41`
- `ROLE_SYSTEM_TEXT` · `42`
- `ROLE_SYSTEM_PUSHBUTTON` · `43`
- `ROLE_SYSTEM_CHECKBUTTON` · `44`
- `ROLE_SYSTEM_RADIOBUTTON` · `45`
- `ROLE_SYSTEM_COMBOBOX` · `46`
- `ROLE_SYSTEM_DROPLIST` · `47`
- `ROLE_SYSTEM_PROGRESSBAR` · `48`
- `ROLE_SYSTEM_DIAL` · `49`
- `ROLE_SYSTEM_HOTKEYFIELD` · `50`
- `ROLE_SYSTEM_SLIDER` · `51`
- `ROLE_SYSTEM_SPINBUTTON` · `52`
- `ROLE_SYSTEM_DIAGRAM` · `53`
- `ROLE_SYSTEM_ANIMATION` · `54`
- `ROLE_SYSTEM_EQUATION` · `55`
- `ROLE_SYSTEM_BUTTONDROPDOWN` · `56`
- `ROLE_SYSTEM_BUTTONMENU` · `57`
- `ROLE_SYSTEM_BUTTONDROPDOWNGRID` · `58`
- `ROLE_SYSTEM_WHITESPACE` · `59`
- `ROLE_SYSTEM_PAGETABLIST` · `60`
- `ROLE_SYSTEM_CLOCK` · `61`
- `ROLE_SYSTEM_SPLITBUTTON` · `62`
- `STATE_SYSTEM_NORMAL` · `0`
- `STATE_SYSTEM_UNAVAILABLE` · `1`
- `STATE_SYSTEM_SELECTED` · `2`
- `STATE_SYSTEM_FOCUSED` · `4`
- `STATE_SYSTEM_PRESSED` · `8`
- `STATE_SYSTEM_CHECKED` · `16`
- `STATE_SYSTEM_MIXED` · `32`
- `STATE_SYSTEM_READONLY` · `64`
- `STATE_SYSTEM_HOTTRACKED` · `128`
- `STATE_SYSTEM_DEFAULT` · `256`
- `STATE_SYSTEM_EXPANDED` · `512`
- `STATE_SYSTEM_COLLAPSED` · `1024`
- `STATE_SYSTEM_BUSY` · `2048`
- `STATE_SYSTEM_INVISIBLE` · `32768`
- `STATE_SYSTEM_OFFSCREEN` · `65536`
- `STATE_SYSTEM_SIZEABLE` · `131072`
- `STATE_SYSTEM_MOVEABLE` · `262144`
- `STATE_SYSTEM_FOCUSABLE` · `1048576`
- `STATE_SYSTEM_SELECTABLE` · `2097152`
- `STATE_SYSTEM_LINKED` · `4194304`
- `STATE_SYSTEM_TRAVERSED` · `8388608`
- `STATE_SYSTEM_MULTISELECTABLE` · `16777216`
- `STATE_SYSTEM_HASPOPUP` · `1073741824`
- `EVENT_SYSTEM_SOUND` · `1`
- `EVENT_SYSTEM_ALERT` · `2`
- `EVENT_SYSTEM_FOREGROUND` · `3`
- `EVENT_SYSTEM_MENUSTART` · `4`
- `EVENT_SYSTEM_MENUEND` · `5`
- `EVENT_SYSTEM_MENUPOPUPSTART` · `6`
- `EVENT_SYSTEM_MENUPOPUPEND` · `7`
- `EVENT_SYSTEM_CAPTURESTART` · `8`
- `EVENT_SYSTEM_CAPTUREEND` · `9`
- `EVENT_SYSTEM_MOVESIZESTART` · `10`
- `EVENT_SYSTEM_MOVESIZEEND` · `11`
- `EVENT_SYSTEM_CONTEXTHELPSTART` · `12`
- `EVENT_SYSTEM_CONTEXTHELPEND` · `13`
- `EVENT_SYSTEM_DRAGDROPSTART` · `14`
- `EVENT_SYSTEM_DRAGDROPEND` · `15`
- `EVENT_SYSTEM_DIALOGSTART` · `16`
- `EVENT_SYSTEM_DIALOGEND` · `17`
- `EVENT_SYSTEM_SCROLLINGSTART` · `18`
- `EVENT_SYSTEM_SCROLLINGEND` · `19`
- `EVENT_SYSTEM_SWITCHSTART` · `20`
- `EVENT_SYSTEM_SWITCHEND` · `21`
- `EVENT_OBJECT_CREATE` · `32768`
- `EVENT_OBJECT_DESTROY` · `32769`
- `EVENT_OBJECT_SHOW` · `32770`
- `EVENT_OBJECT_HIDE` · `32771`
- `EVENT_OBJECT_REORDER` · `32772`
- `EVENT_OBJECT_FOCUS` · `32773`
- `EVENT_OBJECT_SELECTION` · `32774`
- `EVENT_OBJECT_SELECTIONADD` · `32775`
- `EVENT_OBJECT_SELECTIONREMOVE` · `32776`
- `EVENT_OBJECT_SELECTIONWITHIN` · `32777`
- `EVENT_OBJECT_STATECHANGE` · `32778`
- `EVENT_OBJECT_LOCATIONCHANGE` · `32779`
- `EVENT_OBJECT_NAMECHANGE` · `32780`
- `EVENT_OBJECT_DESCRIPTIONCHANGE` · `32781`
- `EVENT_OBJECT_VALUECHANGE` · `32782`
- `EVENT_OBJECT_PARENTCHANGE` · `32783`
- `EVENT_OBJECT_HELPCHANGE` · `32784`
- `EVENT_OBJECT_DEFACTIONCHANGE` · `32785`
- `EVENT_OBJECT_ACCELERATORCHANGE` · `32786`
- `EVENT_OBJECT_INVOKED` · `32787`
- `EVENT_OBJECT_TEXTSELECTIONCHANGED` · `32788`
- `EVENT_OBJECT_CONTENTSCROLLED` · `32789`
- `EVENT_OBJECT_LIVEREGIONCHANGED` · `32793`
- `WM_GETOBJECT` · `61`
- `OBJID_SELF` · `0`
- `OBJID_WINDOW` · `0`
- `OBJID_CLIENT` — constant
- `OBJID_NATIVEOM` — constant
- `_UIA_UIAROOT_OBJECTID` — constant
- `_oleaccProcCache` · `{}`
- `Oleacc` · `'oleacc.dll'`
### `oleacc(symbol, args...)`

### `accNode(id, name, role, bounds, state, children)`

> returns `:object`

### `accTree(window)`

### `accRegister(window, id, name, role, bounds, state)`

### `accUnregister(window, id)`

### `accSetName(window, id, name)`

### `accSetState(window, id, state)`

### `accSetValue(window, id, value)`

### `accSetBounds(window, id, bounds)`

### `accSetDescription(window, id, desc)`

### `accSetDefaultAction(window, id, action)`

### `accFocus(window, id)`

### `accSelection(window, id)`

### `notifyAccEvent(window, event, childId)`

### `enableAccessibility(window)`

> returns `:bool`

### `accRegisterButton(window, id, label, x, y, w, h)`

### `accRegisterCheckbox(window, id, label, x, y, w, h, checked?)`

### `accRegisterRadio(window, id, label, x, y, w, h, selected?)`

### `accRegisterTextField(window, id, label, x, y, w, h, value)`

### `accRegisterSlider(window, id, label, x, y, w, h, value)`

### `accRegisterProgressBar(window, id, label, x, y, w, h, value)`

### `accRegisterTab(window, id, label, x, y, w, h, selected?)`

### `accRegisterListItem(window, id, label, x, y, w, h, selected?)`

### `accRegisterLink(window, id, label, x, y, w, h)`

### `accRegisterStaticText(window, id, label, x, y, w, h)`

### `accRegisterGroup(window, id, label, x, y, w, h)`

### `accRegisterTreeItem(window, id, label, x, y, w, h, expanded?)`

### `accRegisterTable(window, id, label, x, y, w, h)`

### `accRegisterTableCell(window, id, label, x, y, w, h)`

### `accRegisterDropdown(window, id, label, x, y, w, h, expanded?)`

### `accAnnounce(window, id, text)`

### `accDump(window)`

### `accVerify(window)`

> returns `:object`

### `accRoleName(role)`

> returns `:string`

### `accStateName(state)`

> returns `:string`

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

## Module: `windows`

- `sys` · `import(...)`
## Module: `windows-constants`

- `Kernel32` · `'kernel32.dll'`
- `Ntdll` · `'ntdll.dll'`
- `Psapi` · `'psapi.dll'`
- `User32` · `'user32.dll'`
- `Gdi32` · `'gdi32.dll'`
- `Advapi32` · `'advapi32.dll'`
- `Shell32` · `'shell32.dll'`
- `Ole32` · `'ole32.dll'`
- `Ws2_32` · `'ws2_32.dll'`
- `Comctl32` · `'comctl32.dll'`
- `Wininet` · `'wininet.dll'`
- `OpenGL32` · `'opengl32.dll'`
- `Vulkan1` · `'vulkan-1.dll'`
- `D3d9` · `'d3d9.dll'`
- `D3d11` · `'d3d11.dll'`
- `Dxgi` · `'dxgi.dll'`
- `Ddraw` · `'ddraw.dll'`
- `Msvcrt` · `'msvcrt.dll'`
- `Ucrtbase` · `'ucrtbase.dll'`
- `Vcruntime140` · `'vcruntime140.dll'`
- `ActionCenter` · `'actioncenter.dll'`
- `Aclui` · `'aclui.dll'`
- `Acledit` · `'acledit.dll'`
- `Acppage` · `'acppage.dll'`
- `Acproxy` · `'acproxy.dll'`
- `Adprovider` · `'adprovider.dll'`
- `Aeinv` · `'aeinv.dll'`
- `Aepic` · `'aepic.dll'`
- `Amstream` · `'amstream.dll'`
- `Adsldp` · `'adsldp.dll'`
- `Adsnt` · `'adsnt.dll'`
- `Adtschema` · `'adtschema.dll'`
- `Adsldpc` · `'adsldpc.dll'`
- `Adsmsext` · `'adsmsext.dll'`
- `Adhsvc` · `'adhsvc.dll'`
- `Advapi32res` · `'advapi32res.dll'`
- `Advpack` · `'advpack.dll'`
- `Aeevts` · `'aeevts.dll'`
- `Apds` · `'apds.dll'`
- `Winhttp` · `'winhttp.dll'`
- `Urlmon` · `'urlmon.dll'`
- `Crypt32` · `'crypt32.dll'`
- `Bcrypt` · `'bcrypt.dll'`
- `Secur32` · `'secur32.dll'`
- `Comdlg32` · `'comdlg32.dll'`
- `Oleaut32` · `'oleaut32.dll'`
- `Imm32` · `'imm32.dll'`
- `Shlwapi` · `'shlwapi.dll'`
- `Shcore` · `'shcore.dll'`
- `UxTheme` · `'uxtheme.dll'`
- `Dwmapi` · `'dwmapi.dll'`
- `Version` · `'version.dll'`
- `Setupapi` · `'setupapi.dll'`
- `Netapi32` · `'netapi32.dll'`
- `Winmm` · `'winmm.dll'`
- `Avrt` · `'avrt.dll'`
- `Mmdevapi` · `'mmdevapi.dll'`
- `Dsound` · `'dsound.dll'`
- `Mfplat` · `'mfplat.dll'`
- `Mfreadwrite` · `'mfreadwrite.dll'`
- `Mfuuid` · `'mfuuid.dll'`
- `Taskschd` · `'taskschd.dll'`
- `Wevtapi` · `'wevtapi.dll'`
- `Wlanapi` · `'wlanapi.dll'`
- `Mpr` · `'mpr.dll'`
- `Spoolss` · `'spoolss.dll'`
- `Wtsapi32` · `'wtsapi32.dll'`
- `Rasapi32` · `'rasapi32.dll'`
- `Msi` · `'msi.dll'`
- `Wimgapi` · `'wimgapi.dll'`
- `Cabinet` · `'cabinet.dll'`
- `Apphelp` · `'apphelp.dll'`
- `Wer` · `'wer.dll'`
- `Faultrep` · `'faultrep.dll'`
- `Dbghelp` · `'dbghelp.dll'`
- `Dbgeng` · `'dbgeng.dll'`
- `Pdh` · `'pdh.dll'`
- `Iphlpapi` · `'iphlpapi.dll'`
- `Wscapi` · `'wscapi.dll'`
- `Sensapi` · `'sensapi.dll'`
- `Ncrypt` · `'ncrypt.dll'`
- `Cryptui` · `'cryptui.dll'`
- `Wintrust` · `'wintrust.dll'`
- `Samlib` · `'samlib.dll'`
- `Netshell` · `'netshell.dll'`
- `Fwpuclnt` · `'fwpuclnt.dll'`
- `Dnsapi` · `'dnsapi.dll'`
- `Nlaapi` · `'nlaapi.dll'`
- `Httpapi` · `'httpapi.dll'`
- `Rpcrt4` · `'rpcrt4.dll'`
- `Srpapi` · `'srpapi.dll'`
- `Sxs` · `'sxs.dll'`
- `Msvcirt` · `'msvcirt.dll'`
- `ApiSetPrefix` · `'api-ms-win-'`
- `D3dx9Prefix` · `'d3dx9_'`
- `MsvcpPrefix` · `'msvcp'`
- `VcruntimePrefix` · `'vcruntime'`
- `AtlPrefix` · `'atl'`
- `MfcPrefix` · `'mfc'`
- `VcompPrefix` · `'vcomp'`
## Module: `windows-core`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `makeWord(low, high)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `kernel32(symbol, args...)`

### `ntdll(symbol, args...)`

### `ntNative(symbol, args...)`

### `psapi(symbol, args...)`

## Module: `windows-flags`

- `PROCESS_TERMINATE` · `1`
- `PROCESS_VM_READ` · `16`
- `PROCESS_VM_WRITE` · `32`
- `PROCESS_VM_OPERATION` · `8`
- `PROCESS_QUERY_INFORMATION` · `1024`
- `PROCESS_QUERY_LIMITED_INFORMATION` · `4096`
- `PROCESS_ALL_ACCESS` · `2035711`
- `MEM_COMMIT` · `4096`
- `MEM_RESERVE` · `8192`
- `MEM_DECOMMIT` · `16384`
- `MEM_RELEASE` · `32768`
- `PAGE_NOACCESS` · `1`
- `PAGE_READONLY` · `2`
- `PAGE_READWRITE` · `4`
- `PAGE_EXECUTE` · `16`
- `PAGE_EXECUTE_READ` · `32`
- `PAGE_EXECUTE_READWRITE` · `64`
- `FORMAT_MESSAGE_IGNORE_INSERTS` · `512`
- `FORMAT_MESSAGE_FROM_SYSTEM` · `4096`
- `ERROR_SUCCESS` · `0`
- `AF_INET` · `2`
- `SOCK_STREAM` · `1`
- `SOCK_DGRAM` · `2`
- `IPPROTO_TCP` · `6`
- `IPPROTO_UDP` · `17`
- `SOCKET_ERROR` — constant
- `INVALID_SOCKET` — constant
- `SD_RECEIVE` · `0`
- `SD_SEND` · `1`
- `SD_BOTH` · `2`
- `INTERNET_OPEN_TYPE_PRECONFIG` · `0`
- `INTERNET_OPEN_TYPE_DIRECT` · `1`
- `INTERNET_OPEN_TYPE_PROXY` · `3`
- `INTERNET_DEFAULT_HTTP_PORT` · `80`
- `INTERNET_DEFAULT_HTTPS_PORT` · `443`
- `INTERNET_SERVICE_HTTP` · `3`
- `HKEY_CLASSES_ROOT` · `2147483648`
- `HKEY_CURRENT_USER` · `2147483649`
- `HKEY_LOCAL_MACHINE` · `2147483650`
- `HKEY_USERS` · `2147483651`
- `HKEY_CURRENT_CONFIG` · `2147483653`
- `KEY_QUERY_VALUE` · `1`
- `KEY_SET_VALUE` · `2`
- `KEY_CREATE_SUB_KEY` · `4`
- `KEY_ENUMERATE_SUB_KEYS` · `8`
- `KEY_READ` · `131097`
- `KEY_WRITE` · `131078`
- `REG_SZ` · `1`
- `REG_DWORD` · `4`
- `REG_QWORD` · `11`
- `CS_VREDRAW` · `1`
- `CS_HREDRAW` · `2`
- `CS_DBLCLKS` · `8`
- `CS_OWNDC` · `32`
- `WS_OVERLAPPED` · `0`
- `WS_CAPTION` · `12582912`
- `WS_SYSMENU` · `524288`
- `WS_THICKFRAME` · `262144`
- `WS_MINIMIZEBOX` · `131072`
- `WS_MAXIMIZEBOX` · `65536`
- `WS_VISIBLE` · `268435456`
- `WS_CLIPSIBLINGS` · `67108864`
- `WS_CLIPCHILDREN` · `33554432`
- `WS_OVERLAPPEDWINDOW` · `13565952`
- `CW_USEDEFAULT` — constant
- `WS_POPUP` · `2147483648`
- `WS_EX_APPWINDOW` · `262144`
- `GWL_STYLE` — constant
- `GWL_EXSTYLE` — constant
- `SM_CXSCREEN` · `0`
- `SM_CYSCREEN` · `1`
- `HWND_TOP` · `0`
- `HWND_TOPMOST` — constant
- `HWND_NOTOPMOST` — constant
- `WM_CREATE` · `1`
- `WM_DESTROY` · `2`
- `WM_PAINT` · `15`
- `WM_CLOSE` · `16`
- `WM_QUIT` · `18`
- `WM_COMMAND` · `273`
- `SW_HIDE` · `0`
- `SW_MAXIMIZE` · `3`
- `SW_SHOW` · `5`
- `SW_RESTORE` · `9`
- `PM_NOREMOVE` · `0`
- `PM_REMOVE` · `1`
- `MB_OK` · `0`
- `MB_ICONERROR` · `16`
- `MB_ICONWARNING` · `48`
- `MB_ICONINFORMATION` · `64`
- `IDC_ARROW` · `32512`
- `IDI_APPLICATION` · `32512`
## Module: `windows-gdi`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
- `_szBuf` · `bits(...)`
### `isWindows?()`

### `wstr(s)`

- `_gdiProcCache` · `{}`
- `_userProcCache` · `{}`
### `_cachedGdi32(symbol, args...)`

### `_cachedUser32(symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `beginPaint(hwnd, paintStructPtr)`

### `endPaint(hwnd, paintStructPtr)`

### `getDC(hwnd)`

### `releaseDC(hwnd, hdc)`

### `getStockObject(objectIndex)`

### `selectObject(hdc, gdiObject)`

### `setBkMode(hdc, mode)`

### `setTextColor(hdc, colorRef)`

### `textOut(hdc, x, y, text)`

### `createFont(height, width, escapement, orientation, weight, italic, underline, strikeOut, charSet, outPrecision, clipPrecision, quality, pitchAndFamily, faceName)`

### `rectangle(hdc, left, top, right, bottom)`

### `ellipse(hdc, left, top, right, bottom)`

### `createSolidBrush(colorRef)`

### `getTextExtentPoint32(hdc, text)`

> returns `:object`

### `deleteObject(gdiObject)`

## Module: `windows-kernel`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `_utf16leToString(buf)`

### `wstr(s)`

### `cstr(s)`

### `kernel32(symbol, args...)`

### `statusOk?(res)`

> returns `:bool`

### `ptrSize()`

> returns `:int`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `getLastError()`

### `formatMessage(errorCode)`

### `lastErrorMessage()`

### `currentProcessId()`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `loadLibrary(path)`

### `freeLibrary(module)`

### `procAddress(module, symbol)`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `closeHandle(handle)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `windows-loader`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

- `_loaderProcCache` · `{}`
### `_cachedCallIn(library, symbol, args...)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `loadDll(library)`

> returns `:object`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `advapi32(symbol, args...)`

### `shell32(symbol, args...)`

### `ole32(symbol, args...)`

### `ws2_32(symbol, args...)`

### `comctl32(symbol, args...)`

### `wininet(symbol, args...)`

### `opengl32(symbol, args...)`

### `vulkan1(symbol, args...)`

### `d3d9(symbol, args...)`

### `ddraw(symbol, args...)`

### `d3d11(symbol, args...)`

### `dxgi(symbol, args...)`

### `directDrawCreateEx(guidPtr, outPtr, iidPtr, outerUnknown)`

### `directDrawCreate(guidPtr, outPtr, outerUnknown)`

### `direct3DCreate9(sdkVersion)`

### `d3dx9Dll(suffix)`

### `d3dx9(suffix, symbol, args...)`

### `apiSetDll(contract)`

### `apiSet(contract, symbol, args...)`

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

### `_familyDll(prefix, suffix)`

### `msvcpDll(suffix)`

### `msvcpFamily(suffix, symbol, args...)`

### `vcruntimeDll(suffix)`

### `vcruntimeFamily(suffix, symbol, args...)`

### `atlDll(suffix)`

### `atlFamily(suffix, symbol, args...)`

### `mfcDll(suffix)`

### `mfcFamily(suffix, symbol, args...)`

### `vcompDll(suffix)`

### `vcompFamily(suffix, symbol, args...)`

## Module: `windows-net`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `callValueOrZero(res)`

### `ws2_32(symbol, args...)`

### `wininet(symbol, args...)`

### `wsLastError()`

### `wsaStartup(version, wsaDataPtr)`

### `wsaCleanup()`

### `socket(af, socketType, protocol)`

### `bindSocket(sock, sockaddrPtr, sockaddrLen)`

### `connectSocket(sock, sockaddrPtr, sockaddrLen)`

### `listenSocket(sock, backlog)`

### `acceptSocket(sock, addrOutPtr, addrLenInOutPtr)`

### `sendSocket(sock, bufferPtr, size, flags)`

### `recvSocket(sock, bufferPtr, size, flags)`

### `shutdownSocket(sock, how)`

### `closeSocket(sock)`

### `htons(value)`

### `htonl(value)`

### `inetAddr(ipv4)`

### `internetOpen(agent, accessType, proxy, proxyBypass, flags)`

### `internetConnect(hInternet, serverName, serverPort, username, password, service, flags, context)`

### `internetOpenUrl(hInternet, url, headers, headersLen, flags, context)`

### `internetReadFile(hFile, outBufferPtr, bytesToRead, bytesReadOutPtr)`

### `internetCloseHandle(hInternet)`

### `_bytesToString(raw)`

### `sockaddrIn(ipv4, port)`

> returns `:object`

### `_internetReadAll(hInternetFile, chunkBuf, bytesReadBuf, chunkSize, out)`

> returns `:object`

### `internetSimpleGet(url, agent, chunkSize)`

## Module: `windows-registry`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `ptrSize()`

> returns `:int`

### `_zeros(n, acc)`

### `_statusOk?(res)`

> returns `:bool`

### `_ptrRead(address)`

### `_utf16leToString(buf)`

### `wstr(s)`

### `advapi32(symbol, args...)`

### `regCloseKey(hKey)`

### `regOpenKeyEx(rootKey, subKey, options, samDesired, outKeyPtr)`

### `regCreateKeyEx(rootKey, subKey, reserved, className, options, samDesired, securityAttributesPtr, outKeyPtr, dispositionOutPtr)`

### `regQueryValueEx(hKey, valueName, reserved, typeOutPtr, dataOutPtr, dataLenInOutPtr)`

### `regSetValueEx(hKey, valueName, reserved, valueType, dataPtr, dataLen)`

### `regDeleteValue(hKey, valueName)`

### `regDeleteTree(rootKey, subKey)`

### `regReadDword(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteDword(rootKey, subKey, valueName, value)`

> returns `:object`

### `regReadString(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteString(rootKey, subKey, valueName, value)`

> returns `:object`

## Module: `windows-windowing`

- `sys` · `import(...)`
### `_toI32(u)`

- `_cachedDefProcAddr` · `?`
- `_cachedCursor` · `?`
- `_cachedIcon` · `?`
- `_cachedImageBase` · `?`
- `_regClassTimings` · `{}`
- `_registeredClasses` · `{}`
- `DefaultClassName` · `'MagnoliaGUIWindowClass'`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

- `_cachedPtrSize` — constant
### `ptrSize()`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `noMessage?(res)`

> returns `:bool`

- `_k32ProcCache` · `{}`
- `_u32ProcCache` · `{}`
- `_shcoreProcCache` · `{}`
- `_dwmapiProcCache` · `{}`
### `kernel32(symbol, args...)`

### `user32(symbol, args...)`

### `shcore(symbol, args...)`

### `dwmapi(symbol, args...)`

### `moduleHandle(name)`

### `imageBase()`

### `registerClassEx(wndClassExPtr)`

### `createWindowEx(exStyle, className, windowName, style, x, y, width, height, parent, menu, instance, param)`

### `defWindowProc(hwnd, msg, wParam, lParam)`

### `showWindow(hwnd, cmdShow)`

### `updateWindow(hwnd)`

### `getWindowLongPtr(hwnd, index)`

### `setWindowLongPtr(hwnd, index, value)`

### `getSystemMetrics(idx)`

### `destroyWindow(hwnd)`

### `postQuitMessage(exitCode)`

### `getMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax)`

### `peekMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax, removeMsg)`

### `translateMessage(msgPtr)`

### `dispatchMessage(msgPtr)`

### `isWindow(hwnd)`

### `waitMessage()`

### `windowAlive?(hwnd)`

> returns `:bool`

### `msgStructSize()`

> returns `:int`

### `createMsgBuffer()`

### `pumpWindowMessage(hwnd, msgPtr)`

> returns `:object`

### `loadCursor(instance, cursorId)`

### `loadIcon(instance, iconId)`

### `registerWindowClassEx(className, iconHandle, smallIconHandle, cursorHandle, classStyle)`

- `_initTimings` · `?`
### `getRegClassTimings()`

### `getInitTimings()`

### `registerDefaultWindowClass(className)`

### `createTopLevelWindow(className, title, width, height, style)`

### `runWindowLoop(hwnd)`

### `runWindowLoopPeek(hwnd, msgPtr)`

### `messageBox(hwnd, text, caption, msgType)`

### `setWindowText(hwnd, text)`

### `getCursorPos()`

> returns `:object`

### `getWindowRect(hwnd)`

> returns `:object`

### `getWindowPlacement(hwnd)`

> returns `:object`

- `DPI_AWARENESS_CONTEXT_UNAWARE` — constant
- `DPI_AWARENESS_CONTEXT_SYSTEM_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2` — constant
- `MDT_EFFECTIVE_DPI` · `0`
- `MDT_ANGULAR_DPI` · `1`
- `MDT_RAW_DPI` · `2`
### `setProcessDpiAwarenessContext(context)`

> returns `:bool`

### `getDpiForWindow(hwnd)`

### `getDpiForSystem()`

### `getDpiForMonitor(hMonitor, dpiType)`

> returns `:object`

- `MONITOR_DEFAULTTONULL` · `0`
- `MONITOR_DEFAULTTOPRIMARY` · `1`
- `MONITOR_DEFAULTTONEAREST` · `2`
### `monitorFromWindow(hwnd, flags)`

### `adjustWindowRectExForDpi(x, y, w, h, style, exStyle, dpi)`

> returns `:object`

### `enableNonClientDpiScaling(hwnd)`

> returns `:bool`

### `dpiScale(value, dpi)`

### `dpiUnscale(value, dpi)`

- `_MONITORINFOEX_SIZE` · `104`
### `getMonitorInfo(hMonitor)`

> returns `:object`

### `monitorFromPoint(x, y, flags)`

### `monitorFromRect(left, top, right, bottom, flags)`

### `getSystemMetricsForDpi(index, dpi)`

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

