# Windows Windowing Library (windows-windowing)

## Overview

`windows-windowing` contains window creation, class registration, and message
loop helpers.

## Import

```oak
ww := import('windows-windowing')
```

## Window/class helpers

- `registerClassEx`, `registerWindowClassEx`, `registerDefaultWindowClass`
- `createWindowEx`, `createTopLevelWindow`
- `defWindowProc`
- `showWindow`, `updateWindow`, `destroyWindow`
- `setWindowText`, `messageBox`
- `loadCursor`, `loadIcon`
- `getWindowLongPtr`, `setWindowLongPtr`
- `getSystemMetrics`

## Message loop helpers

- `getMessage`, `peekMessage`, `translateMessage`, `dispatchMessage`
- `waitMessage`, `postQuitMessage`
- `createMsgBuffer`, `msgStructSize`
- `pumpWindowMessage`
- `runWindowLoop`, `runWindowLoopPeek`
- `isWindow`, `windowAlive?`
- `callOk?`, `noMessage?`

## Utility exports

- `ptrSize`, `writePtr`, `ptrInt`, `callValueOrZero`, `_zeros`
