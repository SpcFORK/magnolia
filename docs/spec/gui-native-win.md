# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-native-win.oak`

- `_nwImportT0` · `nanotime(...)`
- `windows` · `import(...)`
- `_nwImportT1` · `nanotime(...)`
- `guiNativeWinPresent` · `import(...)`
- `guiNativeWinIcons` · `import(...)`
- `guiNativeWinFrame` · `import(...)`
- `guiNativeWinPoll` · `import(...)`
- `guiNativeWinClose` · `import(...)`
- `guiNativeWinDdraw` · `import(...)`
- `guiNativeWinD3d11` · `import(...)`
- `guiNativeWinVulkan` · `import(...)`
- `guiNativeWinOpenGL` · `import(...)`
- `guiNativeWinProbe` · `import(...)`
- `guiCanvas` · `import(...)`
- `_nwImportT2` · `nanotime(...)`
- `guiThread` · `import(...)`
- `_nwImportT3` · `nanotime(...)`
- `_nwImportTimings` · `{4 entries}`
### `getNativeWinImportTimings()`

- `D3D_SDK_VERSION` · `32`
- `COM_RELEASE` · `2`
- `_windowThreadLockCount` · `0`
### `_acquireWindowThreadLock()`

### `_releaseWindowThreadLock()`

### `_startDllProbeAsync(state)`

### `_startD3d9ProbeAsync(state)`

### `_applyDllProbeResult(state)`

### `_init2DLayer(window)`

### `_init3DLayer(window)`

> returns `:object`

- `SRCCOPY` · `13369376`
- `SWP_NOSIZE` · `1`
- `SWP_NOMOVE` · `2`
- `SWP_NOZORDER` · `4`
- `SWP_NOACTIVATE` · `16`
- `SWP_FRAMECHANGED` · `32`
- `SWP_NOOWNERZORDER` · `512`
### `_fallbackPresenterAfterVulkanFailure(state)`

> returns `:atom`

### `_handleVulkan2DInit(state)`

### `_handleOpenGL2DInit(state)`

> returns `:object`

### `_handleD3d11Init(state)`

> returns `:object`

### `_finalize2DInit(state)`

### `_finalize3DInit(state)`

> returns `:object`

### `_runLayerInit(state)`

### `_ensureLayerInit(window)`

> returns `:object`

### `createWindowState(title, width, height, options, className, frameMs, updateOnDispatch)`

### `createWindowAsync(title, width, height, options, className, frameMs, updateOnDispatch)`

### `awaitWindow(future)`

- `_showTimings` · `?`
### `showWindow(window)`

### `getShowTimings()`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `setFullscreen(window, enabled)`

### `lockResize(window, locked)`

### `setAlwaysOnTop(window, enabled)`

- `WS_EX_LAYERED` · `524288`
- `LWA_ALPHA` · `2`
- `LWA_COLORKEY` · `1`
### `setWindowOpacity(window, alpha)`

### `setWindowColorKey(window, colorKey)`

### `removeLayeredStyle(window)`

### `beginDrag(window)`

### `updateDrag(window)`

### `endDrag(window)`

> returns `?`

### `setTitle(window, title)`

### `setIcon(window, iconSpec)`

### `_ensureLayerInitForFrame(window)`

### `beginFrame(window)`

### `endFrame(window)`

> returns `:int`

### `poll(window)`

> returns `:object`

### `close(window)`

> returns `:int`

- `_windowRegistry` · `[]`
### `registerWindow(window)`

### `unregisterWindow(window)`

### `allWindows()`

### `pollAllWindows()`

### `closeAllWindows()`

> returns `:list`

### `anyWindowOpen?()`

### `saveWindowState(window)`

> returns `:object`

### `restoreWindowState(window, state)`

- `_ERROR_ALREADY_EXISTS` · `183`
### `acquireSingleInstance(name)`

> returns `:object`

### `releaseSingleInstance(inst)`

- `_ptrSize` — constant
- `TBPF_NOPROGRESS` · `0`
- `TBPF_INDETERMINATE` · `1`
- `TBPF_NORMAL` · `2`
- `TBPF_ERROR` · `4`
- `TBPF_PAUSED` · `8`
### `_createTaskbarList3()`

### `_comCall(pInterface, vtableIdx, args...)`

- `_taskbarList3` · `?`
- `_taskbarInited` · `false`
### `_getTaskbar()`

### `setTaskbarProgress(window, completed, total)`

### `setTaskbarProgressState(window, flags)`

### `createOwnedWindow(parent, title, width, height, options)`

### `showModalDialog(parent, title, width, height, options, setupFn)`

### `closeModalDialog(dialog)`

### `setWindowOwner(child, parent)`

### `getWindowMonitor(window)`

### `centerOnMonitor(window)`

### `moveToMonitor(window, hMonitor)`

### `getWindowDpi(window)`

### `extendFrameIntoClientArea(window, margins)`

### `enableGlassSheet(window)`

- `HTCLIENT` · `1`
- `HTCAPTION` · `2`
- `HTSYSMENU` · `3`
- `HTMINBUTTON` · `8`
- `HTMAXBUTTON` · `9`
- `HTLEFT` · `10`
- `HTRIGHT` · `11`
- `HTTOP` · `12`
- `HTTOPLEFT` · `13`
- `HTTOPRIGHT` · `14`
- `HTBOTTOM` · `15`
- `HTBOTTOMLEFT` · `16`
- `HTBOTTOMRIGHT` · `17`
- `HTCLOSE` · `20`
- `_WM_NCHITTEST` · `132`
### `onNcHitTest(window, handler)`

### `customChromeHitTest(mx, my, width, height, borderSize, captionHeight)`

### `setWindowRgn(window, hRgn, redraw)`

### `createRoundRectRgn(left, top, right, bottom, rx, ry)`

### `_escapeXml(s)`

### `_escapePsString(s)`

### `showToastNotification(title, message, options)`

### `showToastWithFallback(title, message, options)`

### `addJumpListTask(title, path, arguments, iconPath, iconIndex, description)`

### `clearJumpList()`

### `addJumpListRecentFile(filePath)`

### `registerFileAssociation(extension, progId, description, command, iconPath)`

### `unregisterFileAssociation(extension, progId)`

### `refreshShellAssociations()`

### `addSearchFolder(folderPath, scope)`

### `searchFiles(query, maxResults)`

### `searchFilesWithProperty(query, property, maxResults)`

