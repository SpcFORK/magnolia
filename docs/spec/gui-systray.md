# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-systray.oak`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `NIM_ADD` · `0`
- `NIM_MODIFY` · `1`
- `NIM_DELETE` · `2`
- `NIM_SETVERSION` · `4`
- `NIF_MESSAGE` · `1`
- `NIF_ICON` · `2`
- `NIF_TIP` · `4`
- `NIF_INFO` · `16`
- `_NOTIFYICONDATA_SIZE` · `956`
- `WM_TRAYICON` · `32769`
### `_zeros(n)`

### `_writeUtf16At(baseAddr, text, maxBytes)`

- `_NID_OFF_HWND` · `8`
- `_NID_OFF_UID` · `16`
- `_NID_OFF_FLAGS` · `20`
- `_NID_OFF_CALLBACKMSG` · `24`
- `_NID_OFF_HICON` · `32`
- `_NID_OFF_TIP` · `40`
- `_NID_OFF_INFO` · `304`
- `_NID_OFF_INFOTITLE` · `820`
- `_NID_OFF_INFOFLAGS` · `948`
### `addTrayIcon(window, id, tooltip, hIcon)`

> returns `:bool`

### `removeTrayIcon(window, id)`

> returns `:bool`

### `updateTrayTooltip(window, id, tooltip)`

> returns `:bool`

### `showBalloonTip(window, id, title, message, iconFlag)`

> returns `:bool`

- `TRAY_WM_LBUTTONDOWN` · `513`
- `TRAY_WM_LBUTTONUP` · `514`
- `TRAY_WM_RBUTTONDOWN` · `516`
- `TRAY_WM_RBUTTONUP` · `517`
- `TRAY_WM_LBUTTONDBLCLK` · `515`
### `onTrayEvent(window, handler)`

### `minimizeToTray(window)`

### `restoreFromTray(window)`

