# Win32 Icon Helpers (gui-native-win-icons)

## Overview

`gui-native-win-icons` loads and applies window icons on Windows. It supports loading icons from system icon IDs, file paths, or pre-resolved Win32 HICON handles, and automatically derives a small icon from the large icon when only one is provided.

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Icon configuration is normally done through the `GUI.createWindow` options object.

## Import

```oak
icons := import('gui-native-win-icons')
{ resolveWindowIcons: resolveWindowIcons, applyWindowIcons: applyWindowIcons } := import('gui-native-win-icons')
```

## Icon Specification Format

An icon spec for `opts.icon`, `opts.taskbarIcon`, `opts.windowIcon`, `opts.iconBig`, or `opts.iconSmall` may be:

| Type      | Interpretation                                      |
|-----------|-----------------------------------------------------|
| `?`       | No icon (use system default)                        |
| integer   | System icon resource ID (e.g. `IDI_APPLICATION = 32512`) |
| string    | File path to a `.ico` file                          |
| object `{ handle }`  | Pre-resolved Win32 HICON handle            |
| object `{ id }`      | System icon resource ID                    |
| object `{ path }`    | `.ico` file path                           |

## Functions

### `resolveWindowIcons(opts)`

Loads big (32×32) and small (16×16) icons from the spec values in `opts`. When only a large icon is available, a small copy is automatically derived with `CopyImage`. Returns `{ big: hicon|0, small: hicon|0 }`.

```oak
icons := resolveWindowIcons({
    icon: 'app.ico'
    taskbarIcon: 'taskbar.ico'
})
// => { big: hBigIcon, small: hSmallIcon }
```

### `applyWindowIcons(hwnd, icons, classAtom)`

Sends `WM_SETICON` messages to `hwnd` for both the big and small icon handles, and updates the window class icon via `SetClassLongPtrW`. Also invalidates the window border with `SetWindowPos` so the title bar refreshes.

```oak
applyWindowIcons(window.hwnd, icons, window.classAtom)
```

## Constants

| Constant          | Value  | Win32 Name             |
|-------------------|--------|------------------------|
| `WM_SETICON`      | 128    | `WM_SETICON`           |
| `ICON_SMALL`      | 0      | Small title-bar icon   |
| `ICON_BIG`        | 1      | Large task-bar icon    |
| `IMAGE_ICON`      | 1      | `LoadImage` type flag  |
| `LR_DEFAULTSIZE`  | 64     | Use system default size |
| `LR_LOADFROMFILE` | 16     | Load from `.ico` file   |
