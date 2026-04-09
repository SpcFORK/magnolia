# gui-filedrop — Drag-and-Drop File Handling

`import('gui-filedrop')` provides Win32 file-drop event handling (`WM_DROPFILES`) and OLE drag-and-drop support for dragging files into windows.

## Quick Start

```oak
fd := import('gui-filedrop')
gui := import('GUI')

window := gui.createWindow('Drop Target', 640, 480, {})

// Simple file drop
fd.enableFileDrop(window)
fd.onFileDrop(window, fn(files) {
    println('Dropped: ' + string(len(files)) + ' files')
    each(files, fn(f) println('  ' + f))
})

// OLE drag-and-drop with visual feedback
fd.enableOleDrop(window)
state := fd.dragDropState()
fd.onDragOver(window, state, fn(x, y) {
    // Draw drop indicator at (x, y)
})
fd.onOleDrop(window, fn(info) {
    println('Dropped at (' + string(info.x) + ', ' + string(info.y) + ')')
    each(info.files, fn(f) println('  ' + f))
})
```

## API Reference

### `enableFileDrop(window)`

Enables drag-and-drop on a window via `DragAcceptFiles`.

### `disableFileDrop(window)`

Disables drag-and-drop on a window.

### `onFileDrop(window, handler)`

Subscribes to file-drop events. The handler receives a list of file path strings.

### `enableOleDrop(window)`

Enables OLE drag-and-drop on a window for richer drop data.

### `disableOleDrop(window)`

Disables OLE drag-and-drop.

### `onOleDrop(window, handler)`

Subscribes to OLE drop events. The handler receives `{files, x, y, effect}`.

### `dragDropState()`

Creates a mutable state tracker for drag-over visual feedback.

### `onDragOver(window, state, handler)`

Tracks drag-over position so applications can display visual feedback during a drag operation.

## Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `WM_DROPFILES` | 563 | Windows message for file drop |
| `DROPEFFECT_NONE` | 0 | No drop effect |
| `DROPEFFECT_COPY` | 1 | Copy drop effect |
| `DROPEFFECT_MOVE` | 2 | Move drop effect |
| `DROPEFFECT_LINK` | 4 | Link drop effect |
