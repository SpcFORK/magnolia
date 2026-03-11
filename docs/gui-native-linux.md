# gui-native-linux

X11/Linux-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, frameMs, updateOnDispatch)` — create X11 window state
- helpers for display/context management

Notes

- Prefer using the generic `GUI` facade; import this module when you need Linux-specific features.
