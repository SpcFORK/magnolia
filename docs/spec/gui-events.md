# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `async/event-bus`

- `_eventKeyCache` · `{}`
### `_eventKey(name)`

### `cs EventBus()`

> returns `:object`

### `create()`

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `gui-input`

- `VK_BACK` · `8`
- `VK_TAB` · `9`
- `VK_CLEAR` · `12`
- `VK_RETURN` · `13`
- `VK_SHIFT` · `16`
- `VK_CONTROL` · `17`
- `VK_ALT` · `18`
- `VK_PAUSE` · `19`
- `VK_CAPSLOCK` · `20`
- `VK_ESCAPE` · `27`
- `VK_SPACE` · `32`
- `VK_PAGEUP` · `33`
- `VK_PAGEDOWN` · `34`
- `VK_END` · `35`
- `VK_HOME` · `36`
- `VK_LEFT` · `37`
- `VK_UP` · `38`
- `VK_RIGHT` · `39`
- `VK_DOWN` · `40`
- `VK_INSERT` · `45`
- `VK_DELETE` · `46`
- `VK_0` · `48`
- `VK_1` · `49`
- `VK_2` · `50`
- `VK_3` · `51`
- `VK_4` · `52`
- `VK_5` · `53`
- `VK_6` · `54`
- `VK_7` · `55`
- `VK_8` · `56`
- `VK_9` · `57`
- `VK_A` · `65`
- `VK_B` · `66`
- `VK_C` · `67`
- `VK_D` · `68`
- `VK_E` · `69`
- `VK_F` · `70`
- `VK_G` · `71`
- `VK_H` · `72`
- `VK_I` · `73`
- `VK_J` · `74`
- `VK_K` · `75`
- `VK_L` · `76`
- `VK_M` · `77`
- `VK_N` · `78`
- `VK_O` · `79`
- `VK_P` · `80`
- `VK_Q` · `81`
- `VK_R` · `82`
- `VK_S` · `83`
- `VK_T` · `84`
- `VK_U` · `85`
- `VK_V` · `86`
- `VK_W` · `87`
- `VK_X` · `88`
- `VK_Y` · `89`
- `VK_Z` · `90`
- `VK_NUMPAD0` · `96`
- `VK_NUMPAD1` · `97`
- `VK_NUMPAD2` · `98`
- `VK_NUMPAD3` · `99`
- `VK_NUMPAD4` · `100`
- `VK_NUMPAD5` · `101`
- `VK_NUMPAD6` · `102`
- `VK_NUMPAD7` · `103`
- `VK_NUMPAD8` · `104`
- `VK_NUMPAD9` · `105`
- `VK_MULTIPLY` · `106`
- `VK_ADD` · `107`
- `VK_SEPARATOR` · `108`
- `VK_SUBTRACT` · `109`
- `VK_DECIMAL` · `110`
- `VK_DIVIDE` · `111`
- `VK_F1` · `112`
- `VK_F2` · `113`
- `VK_F3` · `114`
- `VK_F4` · `115`
- `VK_F5` · `116`
- `VK_F6` · `117`
- `VK_F7` · `118`
- `VK_F8` · `119`
- `VK_F9` · `120`
- `VK_F10` · `121`
- `VK_F11` · `122`
- `VK_F12` · `123`
- `VK_NUMLOCK` · `144`
- `VK_SCROLLLOCK` · `145`
- `VK_OEM_SEMICOLON` · `186`
- `VK_OEM_PLUS` · `187`
- `VK_OEM_COMMA` · `188`
- `VK_OEM_MINUS` · `189`
- `VK_OEM_PERIOD` · `190`
- `VK_OEM_SLASH` · `191`
- `VK_OEM_TILDE` · `192`
- `VK_OEM_LBRACKET` · `219`
- `VK_OEM_BACKSLASH` · `220`
- `VK_OEM_RBRACKET` · `221`
- `VK_OEM_QUOTE` · `222`
- `VK_LSHIFT` · `160`
- `VK_RSHIFT` · `161`
- `VK_LCONTROL` · `162`
- `VK_RCONTROL` · `163`
- `VK_LALT` · `164`
- `VK_RALT` · `165`
### `isLetterKey?(vk)`

> returns `:bool`

### `isDigitKey?(vk)`

> returns `:bool`

### `isNumpadKey?(vk)`

> returns `:bool`

### `isFunctionKey?(vk)`

> returns `:bool`

### `isArrowKey?(vk)`

> returns `:bool`

### `isModifierKey?(vk)`

> returns `:bool`

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

## Module: `lib\gui-events.oak`

- `windows` · `import(...)`
- `eventBusLib` · `import(...)`
- `guiInput` · `import(...)`
- `guiThread` · `import(...)`
### `emitThreadSafe(window, event, payload)`

### `emitFromWorker(window, event, payload)`

- `_mapResult` · `{2 entries}`
### `_mapMouseXY(window, physX, physY)`

### `_ensureEventBus(window)`

### `eventBus(window)`

### `on(window, event, handler)`

> returns `?`

### `once(window, event, handler)`

> returns `?`

### `off(window, event, tokenOrHandler)`

> returns `:int`

### `emit(window, event, payload, onDone)`

> returns `:int`

### `listenerCount(window, event)`

> returns `:int`

### `clearListeners(window, event)`

> returns `:bool`

### `publish(window, event, payload)`

> returns `:int`

### `onDispatch(window, handler)`

### `onceDispatch(window, handler)`

### `onRunStart(window, handler)`

### `onceRunStart(window, handler)`

### `onIdle(window, handler)`

### `onceIdle(window, handler)`

### `onFrame(window, handler)`

### `onceFrame(window, handler)`

### `onClosing(window, handler)`

### `onceClosing(window, handler)`

### `onClosed(window, handler)`

### `onceClosed(window, handler)`

- `MSG_OFF_TYPE` · `8`
- `MSG_OFF_WPARAM` · `16`
- `MSG_OFF_LPARAM` · `24`
- `FORM_WM_MOUSEMOVE` · `512`
- `FORM_WM_LBUTTONDOWN` · `513`
- `FORM_WM_LBUTTONUP` · `514`
- `FORM_WM_RBUTTONDOWN` · `516`
- `FORM_WM_RBUTTONUP` · `517`
- `FORM_WM_MBUTTONDOWN` · `519`
- `FORM_WM_MBUTTONUP` · `520`
- `FORM_WM_MOUSEWHEEL` · `522`
- `FORM_WM_LBUTTONDBLCLK` · `515`
- `FORM_WM_KEYDOWN` · `256`
- `FORM_WM_KEYUP` · `257`
- `FORM_WM_CHAR` · `258`
- `GUI_WM_SIZE` · `5`
- `GUI_WM_PAINT` · `15`
- `GUI_WM_ERASEBKGND` · `20`
- `GUI_WM_SIZING` · `532`
- `GUI_WM_WINDOWPOSCHANGED` · `71`
- `GUI_WM_ENTERSIZEMOVE` · `561`
- `GUI_WM_EXITSIZEMOVE` · `562`
- `GUI_WM_MOUSEHOVER` · `673`
- `GUI_WM_MOUSELEAVE` · `675`
- `GUI_WM_DPICHANGED` · `736`
- `WM_IME_STARTCOMPOSITION` · `269`
- `WM_IME_ENDCOMPOSITION` · `270`
- `WM_IME_COMPOSITION` · `271`
- `WM_IME_SETCONTEXT` · `641`
- `WM_IME_NOTIFY` · `642`
- `WM_IME_CHAR` · `646`
- `GCS_COMPSTR` · `8`
- `GCS_RESULTSTR` · `2048`
- `GCS_COMPATTR` · `16`
- `GCS_CURSORPOS` · `128`
- `TME_HOVER` · `1`
- `TME_LEAVE` · `2`
- `TME_SIZEOF` · `24`
- `FORM_VK_BACK` — constant
- `FORM_VK_TAB` — constant
- `FORM_VK_RETURN` — constant
- `FORM_VK_SHIFT` — constant
- `FORM_VK_CONTROL` — constant
- `FORM_VK_ALT` — constant
- `FORM_VK_ESCAPE` — constant
- `MK_LBUTTON` · `1`
- `MK_RBUTTON` · `2`
- `MK_SHIFT` · `4`
- `MK_CONTROL` · `8`
- `MK_MBUTTON` · `16`
### `modShift?(wp)`

> returns `:bool`

### `modCtrl?(wp)`

> returns `:bool`

### `modAlt?(window)`

> returns `:bool`

### `keyShiftDown?(window)`

> returns `:bool`

### `keyCtrlDown?(window)`

> returns `:bool`

### `keyAltDown?(window)`

### `formMsgType(window)`

### `formMsgWParam(window)`

### `formMsgLParam(window)`

### `formLoWord(v)`

> returns `:bool`

### `formHiWord(v)`

> returns `:bool`

### `_cacheDispatchContext(window)`

### `_clearDispatchContext(window)`

> returns `?`

### `formEventContext(window)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

> returns `:bool`

### `_vkKeyName(vk)`

- `_noKeyMatch` · `{1 entries}`
### `_extractKeyEvent(window, evt, expectedMsgType, expectedEvtType)`

> returns `:object`

### `onKeyDownEvent(window, handler)`

### `onceKeyDownEvent(window, handler)`

### `onKeyUpEvent(window, handler)`

### `onceKeyUpEvent(window, handler)`

### `onMouseMove(window, handler)`

### `onLButtonDown(window, handler)`

### `onLButtonUp(window, handler)`

### `onRButtonDown(window, handler)`

### `onRButtonUp(window, handler)`

### `onMButtonDown(window, handler)`

### `onMButtonUp(window, handler)`

### `onMouseWheel(window, handler)`

### `onLButtonDblClk(window, handler)`

### `_trackMouseEvent(window, flags)`

### `enableMouseTracking(window)`

### `onMouseHover(window, handler)`

### `onMouseLeave(window, handler)`

### `onKeyDown(window, handler)`

### `onKeyUp(window, handler)`

### `onChar(window, handler)`

### `onResize(window, handler)`

### `onDpiChanged(window, handler)`

### `_imeGetString(hwnd, gcsFlag)`

### `_imeGetCursorPos(hwnd)`

### `onImeStartComposition(window, handler)`

### `onImeEndComposition(window, handler)`

### `onImeComposition(window, handler)`

### `setImePosition(window, x, y)`

- `WM_TOUCH` · `576`
- `_TOUCHINPUT_SIZE` · `40`
- `TOUCHEVENTF_MOVE` · `1`
- `TOUCHEVENTF_DOWN` · `2`
- `TOUCHEVENTF_UP` · `4`
- `TOUCHEVENTF_INRANGE` · `8`
- `TOUCHEVENTF_PRIMARY` · `16`
- `TOUCHEVENTF_NOCOALESCE` · `32`
- `TOUCHEVENTF_PEN` · `64`
- `TOUCHEVENTF_PALM` · `128`
- `TOUCHINPUTMASKF_CONTACTAREA` · `4`
- `TOUCHINPUTMASKF_EXTRAINFO` · `2`
- `TOUCHINPUTMASKF_TIMEFROMSYSTEM` · `1`
- `TWF_FINETOUCH` · `1`
- `TWF_WANTPALM` · `2`
- `WM_POINTERDOWN` · `582`
- `WM_POINTERUP` · `583`
- `WM_POINTERUPDATE` · `581`
- `WM_POINTERENTER` · `585`
- `WM_POINTERLEAVE` · `586`
- `WM_POINTERCAPTURECHANGED` · `588`
- `WM_POINTERWHEEL` · `590`
- `WM_POINTERHWHEEL` · `591`
- `PT_POINTER` · `1`
- `PT_TOUCH` · `2`
- `PT_PEN` · `3`
- `PT_MOUSE` · `4`
- `PT_TOUCHPAD` · `5`
- `POINTER_FLAG_NONE` · `0`
- `POINTER_FLAG_NEW` · `1`
- `POINTER_FLAG_INRANGE` · `2`
- `POINTER_FLAG_INCONTACT` · `4`
- `POINTER_FLAG_FIRSTBUTTON` · `16`
- `POINTER_FLAG_SECONDBUTTON` · `32`
- `POINTER_FLAG_PRIMARY` · `8192`
- `POINTER_FLAG_DOWN` · `65536`
- `POINTER_FLAG_UPDATE` · `131072`
- `POINTER_FLAG_UP` · `262144`
- `PEN_FLAG_BARREL` · `1`
- `PEN_FLAG_INVERTED` · `2`
- `PEN_FLAG_ERASER` · `4`
- `PEN_MASK_PRESSURE` · `1`
- `PEN_MASK_ROTATION` · `2`
- `PEN_MASK_TILT_X` · `4`
- `PEN_MASK_TILT_Y` · `8`
### `registerTouchWindow(window, flags)`

### `unregisterTouchWindow(window)`

### `_parseTouchInputs(window, count)`

### `onTouch(window, handler)`

### `enableTouchInput(window, options)`

### `_pointerIdFromWParam(wp)`

> returns `:bool`

### `_getPointerType(pointerId)`

- `_POINTER_INFO_SIZE` · `96`
### `_getPointerInfo(pointerId)`

> returns `:object`

- `_POINTER_PEN_INFO_SIZE` · `120`
### `_getPointerPenInfo(pointerId)`

> returns `:object`

### `_makePointerEvent(window, pointerId)`

### `onPointerDown(window, handler)`

### `onPointerUp(window, handler)`

### `onPointerUpdate(window, handler)`

### `onPointerEnter(window, handler)`

### `onPointerLeave(window, handler)`

### `enablePointerInput(window)`

- `WM_INPUT` · `255`
- `HID_USAGE_PAGE_GENERIC` · `1`
- `HID_USAGE_PAGE_GAME` · `5`
- `HID_USAGE_PAGE_LED` · `8`
- `HID_USAGE_PAGE_BUTTON` · `9`
- `HID_USAGE_GENERIC_POINTER` · `1`
- `HID_USAGE_GENERIC_MOUSE` · `2`
- `HID_USAGE_GENERIC_JOYSTICK` · `4`
- `HID_USAGE_GENERIC_GAMEPAD` · `5`
- `HID_USAGE_GENERIC_KEYBOARD` · `6`
- `HID_USAGE_GENERIC_KEYPAD` · `7`
- `HID_USAGE_GENERIC_MULTI_AXIS` · `8`
- `RIDEV_REMOVE` · `1`
- `RIDEV_INPUTSINK` · `256`
- `RIDEV_NOLEGACY` · `48`
- `RIDEV_DEVNOTIFY` · `8192`
- `RID_INPUT` · `268435459`
- `RID_HEADER` · `268435461`
- `RIM_TYPEMOUSE` · `0`
- `RIM_TYPEKEYBOARD` · `1`
- `RIM_TYPEHID` · `2`
- `_RAWINPUTDEVICE_SIZE` — constant
### `registerRawInputDevice(window, usagePage, usage, flags)`

### `unregisterRawInputDevice(usagePage, usage)`

- `_RAWINPUTHEADER_SIZE` — constant
### `_getRawInputData(lParam)`

> returns `:object`

### `onRawInput(window, handler)`

### `enableRawMouse(window)`

### `enableRawKeyboard(window)`

### `enableRawGamepad(window)`

### `enableRawJoystick(window)`

### `isResizeDispatch?(window, step)`

> returns `:bool`

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

