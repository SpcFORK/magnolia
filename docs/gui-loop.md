# GUI Loop Library (gui-loop)

## Overview

`gui-loop` centralizes frame timing and scheduling logic for `GUI` runtime loops.

## Import

```oak
loop := import('gui-loop')
```

## Exports

- `frameIntervalNt(window)` - target frame interval in nanoseconds
- `frameMaxDtNt(window)` - maximum clamped frame delta in nanoseconds
- `markUrgentFrameIfResizeDispatch(window, step)` - marks immediate redraw on resize-related Win32 dispatch events
- `maybeRunFrame(window, onFrame, force, publish)` - runs frame callback when due and emits frame event through `publish`
- `sleepUntilNextFrame(window)` - backend-specific idle sleep until next frame

## Notes

- `GUI` uses this module internally and re-exports compatibility wrappers.
