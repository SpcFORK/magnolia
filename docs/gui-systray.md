# gui-systray — System Tray Icons

`import('gui-systray')` provides Win32 system tray (notification area) icon support using `Shell_NotifyIconW` with balloon tip notifications and minimize-to-tray helpers.

## Quick Start

```oak
tray := import('gui-systray')
gui := import('GUI')

window := gui.createWindow('Tray Demo', 640, 480, {})

// Add tray icon
tray.addTrayIcon(window, 1, 'My App', window.hIcon)

// Handle tray events
tray.onTrayEvent(window, fn(event) if event {
    :lbuttondblclk -> tray.restoreFromTray(window)
    :rbuttonup -> println('Right-clicked tray icon')
})

// Show a balloon notification
tray.showBalloonTip(window, 1, 'Hello', 'App is running in the tray', 1)

// Minimize to tray
tray.minimizeToTray(window)
```

## API Reference

### `addTrayIcon(window, id, tooltip, hIcon)`

Adds a system tray icon. Returns `true` on success.

### `removeTrayIcon(window, id)`

Removes a system tray icon. Returns `true` on success.

### `updateTrayTooltip(window, id, tooltip)`

Updates the tooltip text shown when hovering the tray icon.

### `showBalloonTip(window, id, title, message, iconFlag)`

Shows a balloon notification. `iconFlag`: `0` = none, `1` = info, `2` = warning, `3` = error.

### `onTrayEvent(window, handler)`

Subscribes to tray icon events. The handler receives an event type atom: `:lbuttondown`, `:lbuttonup`, `:rbuttondown`, `:rbuttonup`, `:lbuttondblclk`.

### `minimizeToTray(window)`

Hides the window (minimize to tray).

### `restoreFromTray(window)`

Shows the window and brings it to the foreground.

## Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `NIM_ADD` | 0 | Add icon |
| `NIM_MODIFY` | 1 | Modify icon |
| `NIM_DELETE` | 2 | Delete icon |
| `NIF_MESSAGE` | 1 | Callback message flag |
| `NIF_ICON` | 2 | Icon handle flag |
| `NIF_TIP` | 4 | Tooltip flag |
| `NIF_INFO` | 16 | Balloon tip flag |
| `WM_TRAYICON` | 32769 | Tray callback message id |
