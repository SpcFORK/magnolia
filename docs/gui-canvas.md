# gui-canvas — Floating Windowless Canvases

`gui-canvas` provides lightweight floating off-screen drawing surfaces that can
be positioned, layered, and composited onto a parent window.  Canvases reuse
the same drawing API as windows — all existing primitives (`fillRect`,
`drawText`, `drawLine`, `pushMask`, etc.) work on canvas objects unchanged.

## Quick start

```oak
gui := import('GUI')

window := gui.createWindow('Demo', 800, 600)

// Create a floating canvas at (50, 50), 200×100 pixels
panel := gui.createCanvas(window, {
    x: 50, y: 50
    width: 200, height: 100
    zIndex: 1
})

gui.run(window, ?, fn(win, dt) {
    gui.beginFrame(win)
    gui.fillRect(win, 0, 0, 800, 600, gui.rgb(20, 20, 30))

    // Draw into the floating canvas
    gui.beginCanvas(panel)
    gui.fillRect(panel, 0, 0, 200, 100, gui.rgb(40, 40, 60))
    gui.drawText(panel, 10, 10, 'Floating panel', gui.rgb(255, 255, 255))
    gui.endCanvas(panel)

    gui.endFrame(win)  // composites all canvases automatically
})
```

## API Reference

### Creation & destruction

| Function | Description |
|---|---|
| `createCanvas(window, options)` | Create a floating canvas on `window`. Returns a canvas object. |
| `destroyCanvas(canvas)` | Release GDI resources and remove from parent. |
| `destroyAllCanvases(window)` | Destroy all canvases on a window (called automatically on close). |

**Options:**

| Field | Type | Default | Description |
|---|---|---|---|
| `x` | int | 0 | X position on parent window |
| `y` | int | 0 | Y position on parent window |
| `width` | int | 200 | Canvas width in pixels |
| `height` | int | 100 | Canvas height in pixels |
| `visible` | bool | true | Whether the canvas is composited |
| `zIndex` | int | 0 | Layer order (higher = on top) |
| `opacity` | int | 255 | 0–255 whole-surface alpha (255 = opaque) |
| `transparentColor` | color/? | ? | Color-key for transparency (skips pixels of this color) |

### Drawing

| Function | Description |
|---|---|
| `beginCanvas(canvas)` | Prepare canvas for drawing (allocates surface on first call). |
| `endCanvas(canvas)` | Finish drawing to canvas (flushes GDI). |

Between `beginCanvas` / `endCanvas`, use **any** GUI draw function with the
canvas object in place of a window:

```oak
gui.beginCanvas(canvas)
gui.fillRect(canvas, 0, 0, w, h, bg)
gui.drawText(canvas, 10, 10, text, fg)
gui.drawLine(canvas, 0, 0, w, h, fg)
gui.pushMask(canvas, 5, 5, w - 10, h - 10)
// ... clipped drawing ...
gui.popMask(canvas)
gui.endCanvas(canvas)
```

### Manipulation

| Function | Description |
|---|---|
| `moveCanvas(canvas, x, y)` | Reposition the canvas. |
| `resizeCanvas(canvas, w, h)` | Change canvas dimensions (surface reallocated on next `beginCanvas`). |
| `setCanvasVisible(canvas, vis)` | Show/hide the canvas. |
| `setCanvasZIndex(canvas, z)` | Change layer order. |
| `setCanvasOpacity(canvas, alpha)` | Set whole-surface alpha (0–255). |
| `setCanvasTransparentColor(canvas, color)` | Set/clear the color-key for transparency. |

### Queries & hit testing

| Function | Description |
|---|---|
| `isCanvas?(obj)` | Returns `true` if `obj` is a canvas. |
| `canvases(window)` | List of all canvases on a window. |
| `canvasCount(window)` | Number of canvases. |
| `canvasAt(window, px, py)` | Topmost visible canvas at window coordinates, or `?`. |
| `canvasHitTest?(canvas, px, py)` | Is `(px, py)` inside the canvas? |
| `canvasToLocal(canvas, px, py)` | Convert window coords to canvas-local `{ x, y }`. |

## Compositing

Canvases are composited automatically during `gui.endFrame(window)`, after user
drawing but before presenting to screen.  Visible canvases with `_dirty = true`
(set by `beginCanvas`) are sorted by `zIndex` and blitted onto the window's
back-buffer.

**Compositing modes** (Windows):

| Condition | Method | Performance |
|---|---|---|
| `opacity = 255`, no transparent color | `BitBlt` (SRCCOPY) | Fastest |
| `transparentColor` set | `TransparentBlt` | Moderate |
| `opacity < 255` | `AlphaBlend` | Moderate |

On Linux, canvases use X11 Pixmaps composited via `XCopyArea`.
On Web, canvas draw messages are wrapped in clip/translate and appended to the
parent's message list.

## Architecture

A canvas object has the same interface fields as a window (`type: :ok`,
`backend`, `frameHdc`, `width`, `height`), allowing the existing cross-platform
drawing functions in `gui-draw.oak` to operate on canvases without any
modification.  Each canvas maintains its own GDI pen/brush/font caches.

```
┌─────────────────────────────────────────┐
│              Parent Window              │
│  frameHdc  ← user draws here           │
│                                         │
│  ┌─────────┐  ┌──────────────┐         │
│  │ Canvas A │  │   Canvas B   │         │
│  │ z=0      │  │   z=1        │         │
│  │ frameHdc │  │   frameHdc   │         │
│  └─────────┘  └──────────────┘         │
│                                         │
│  endFrame() composites A then B         │
│  onto window.frameHdc, then presents    │
└─────────────────────────────────────────┘
```
