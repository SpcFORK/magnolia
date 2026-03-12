# gui-native-linux

X11/Linux-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, frameMs, updateOnDispatch)` — create X11 window state
- `showWindow(window)` / `hideWindow(window)` — map or unmap an X11 window
- `moveWindow(window, x, y)` — move an X11 window
- `resizeWindow(window, width, height)` — resize an X11 window
- `setFullscreen(window, enabled)` — best-effort fullscreen by resizing to display bounds
- helpers for display/context management

Notes

- Prefer using the generic `GUI` facade; import this module when you need Linux-specific features.
