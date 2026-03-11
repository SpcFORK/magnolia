# gui-native-win

Win32-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, frameMs, updateOnDispatch)` — create native Win32 window state
- helper constants and icon handling utilities
- platform-specific frame batching to reuse device contexts

Notes

- Prefer using the generic `GUI` facade; import this module when you need Win32-specific features.
