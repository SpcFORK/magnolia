# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\linux-windowing.oak`

- `sys` · `import(...)`
### `_r1Positive?(res)`

> returns `:bool`

### `callOk?(res)`

> returns `:bool`

### `_zeros(n, acc)`

### `xEventSize()`

> returns `:int`

### `createXEventBuffer()`

### `xEventType(eventPtr)`

### `openDisplay(displayName)`

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

### `createGC(display, drawable, valueMask, values)`

### `freeGC(display, gc)`

### `setForeground(display, gc, color)`

### `drawLine(display, window, gc, x1, y1, x2, y2)`

### `fillRectangle(display, window, gc, x, y, width, height)`

### `drawString(display, window, gc, x, y, text)`

### `flush(display)`

### `pending(display)`

### `nextEvent(display, eventPtr)`

### `_openWindowFromRoot(display, screen, root, black, white, title, width, height)`

> returns `:object`

### `_openWindowFromScreen(display, screen, title, width, height)`

### `_openWindowFromDisplay(display, title, width, height)`

### `openDefaultWindow(title, width, height)`

### `closeWindow(state)`

> returns `:int`

### `pumpWindowEvent(display, eventPtr)`

> returns `:object`

### `runWindowLoop(display, eventPtr)`

> returns `:int`

