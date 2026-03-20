# Linux Windowing Module (linux-windowing)

## Overview

`linux-windowing` provides X11-focused wrappers and helpers for creating a
simple window, drawing basic primitives/text, and running an event loop.

## Import

```oak
x11 := import('linux-windowing')
```

## Core Helpers

### `_r1Positive?(res)`

Returns `true` when `res` contains a positive `r1` field.

### `callOk?(res)`

Normalizes success checks for mixed result shapes used by interop calls.

### `xEventSize()`
### `createXEventBuffer()`
### `xEventType(eventPtr)`

Helpers for managing a raw XEvent-compatible buffer.

## Display and Window APIs

### `openDisplay(displayName?)`
### `closeDisplay(display)`
### `defaultScreen(display)`
### `rootWindow(display, screen)`
### `blackPixel(display, screen)`
### `whitePixel(display, screen)`
### `createSimpleWindow(display, parent, x, y, width, height, borderWidth, border, background)`
### `destroyWindow(display, window)`
### `storeName(display, window, title)`
### `selectInput(display, window, eventMask)`
### `mapWindow(display, window)`
### `unmapWindow(display, window)`
### `moveWindow(display, window, x, y)`
### `resizeWindow(display, window, width, height)`
### `displayWidth(display, screen)`
### `displayHeight(display, screen)`

## Drawing APIs

### `createGC(display, drawable, valueMask, values)`
### `freeGC(display, gc)`
### `setForeground(display, gc, color)`
### `drawLine(display, window, gc, x1, y1, x2, y2)`
### `fillRectangle(display, window, gc, x, y, width, height)`
### `drawString(display, window, gc, x, y, text)`
### `flush(display)`

## Event APIs

### `pending(display)`
### `nextEvent(display, eventPtr)`
### `pumpWindowEvent(display, eventPtr)`
### `runWindowLoop(display, eventPtr)`

`runWindowLoop` exits with `0` when an X11 `ClientMessage` or
`DestroyNotify` event is observed.

## High-Level Helpers

### `openDefaultWindow(title, width, height)`

Creates a top-level window and returns:

```oak
{
    type: :ok
    display: <int>
    window: <int>
    screen: <int>
    black: <int>
    white: <int>
}
```

### `closeWindow(state)`

Destroys and closes resources produced by `openDefaultWindow`.

## Example

See [samples/linux-window.oak](../samples/linux-window.oak) and
[samples/linux-draw.oak](../samples/linux-draw.oak) for runnable examples.
