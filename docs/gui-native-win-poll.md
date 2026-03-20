# Win32 Message Poll Helpers (gui-native-win-poll)

## Overview

`gui-native-win-poll` wraps `PeekMessage` / `TranslateMessage` / `DispatchMessage` into a single Oak function used by the `gui-native-win` event loop. It handles the closed-window detection case where `PeekMessage` returns `errno 1400` (invalid window handle).

This module is part of the native Windows GUI backend and is not intended for direct use in application code. Message polling is managed by the `GUI` module's event loop.

## Import

```oak
poll := import('gui-native-win-poll')
{ pollWindowMessages: pollWindowMessages } := import('gui-native-win-poll')
```

## Functions

### `pollWindowMessages(window)`

Calls `PeekMessage` with `PM_REMOVE` on the message buffer pre-allocated in `window.msgBuf`. Returns a tagged result object:

| Return value                        | Condition                                              |
|-------------------------------------|--------------------------------------------------------|
| `{ type: :dispatch, detail: msg }`  | A message was retrieved and dispatched successfully.   |
| `{ type: :idle, detail: msg }`      | No message was available (`PeekMessage` returned 0).   |
| `{ type: :closed }`                 | `errno 1400` — window handle is invalid; sets `window.closed = true`. |

```oak
result := pollWindowMessages(window)
if result.type {
    :dispatch -> handleMessage(result.detail)
    :closed   -> printf('window closed')
    _ -> ?  // idle, continue loop
}
```

**Note:** The function calls `TranslateMessage` and `DispatchMessage` only when a message is actually retrieved; virtual key codes are translated to `WM_CHAR` messages before dispatch.
