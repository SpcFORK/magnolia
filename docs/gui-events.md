# GUI Events Library (gui-events)

## Overview

`gui-events` contains event bus wiring, key-event extraction, and Win32 message
helpers used by `GUI`.

## Import

```oak
events := import('gui-events')
```

## Event bus helpers

- `_ensureEventBus(window)`
- `eventBus(window)`
- `on(window, event, handler)`
- `once(window, event, handler)`
- `off(window, event, tokenOrHandler)`
- `emit(window, event, payload, onDone?)`
- `listenerCount(window, event)`
- `clearListeners(window, event)`
- `publish(window, event, payload)`

## Lifecycle subscriptions

- `onDispatch`, `onceDispatch`
- `onRunStart`, `onceRunStart`
- `onIdle`, `onceIdle`
- `onFrame`, `onceFrame`
- `onClosing`, `onceClosing`
- `onClosed`, `onceClosed`

## Key and message helpers

- `onKeyDownEvent(window, handler)`
- `onceKeyDownEvent(window, handler)`
- `onKeyUpEvent(window, handler)`
- `onceKeyUpEvent(window, handler)`
- `formMsgType`, `formMsgWParam`, `formMsgLParam`
- `formLoWord`, `formHiWord`
- `formEventContext(window)`
- `formInRect?(mx, my, x, y, w, h)`

## Win32 dispatch subscriptions

- `onMouseMove`
- `onLButtonDown`
- `onLButtonUp`
- `onRButtonDown`
- `onRButtonUp`
- `onKeyDown`
- `onKeyUp`
- `onChar`
- `onResize`
- `isResizeDispatch?(window, step)`

## Constants

Message offsets and key/message constants are exposed for `GUI` internals,
including `MSG_OFF_*`, `FORM_WM_*`, `GUI_WM_*`, and `FORM_VK_*`.
