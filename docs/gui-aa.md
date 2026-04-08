# gui-aa — Anti-Aliased Drawing Primitives

`import('gui-aa')` provides software anti-aliased (AA) drawing routines that
work on **all backends** (Windows GDI, Linux X11, Web).  The module builds on
top of the existing `drawLine` primitive by blending edge pixels using
`gui-color.opacity()` for sub-pixel coverage.

These functions are also exposed on the main `GUI` facade as convenience
wrappers (e.g. `gui.drawLineAA(...)`).

## Quick Start

```oak
gui := import('GUI')
{ rgb: rgb } := import('gui-color')

// Inside your render loop:
bg := rgb(30, 30, 30)
gui.fillRect(window, 0, 0, window.width, window.height, bg)

// Anti-aliased line
gui.drawLineAA(window, 50, 50, 300, 200, rgb(255, 100, 100), bg)

// Anti-aliased filled circle
gui.drawCircleFilledAA(window, 200, 200, 60, rgb(100, 200, 255), bg)

// Anti-aliased circle outline
gui.drawCircleOutlineAA(window, 400, 200, 40, rgb(255, 255, 100), bg)

// Anti-aliased filled ellipse
gui.drawEllipseFilledAA(window, 300, 300, 80, 40, rgb(100, 255, 100), bg)

// Anti-aliased rounded rectangle
gui.drawRoundedRectFilledAA(window, 50, 350, 200, 80, 12, rgb(200, 150, 255), bg)

// Anti-aliased filled triangle
gui.drawTriangleFilledAA(window,
    { x: 500, y: 50 }, { x: 600, y: 200 }, { x: 420, y: 180 },
    rgb(255, 200, 100), bg)
```

## API Reference

All functions accept a `bgColor` parameter — the background color used for
alpha blending at edges.  Pass the color you cleared the background with for
best results.  Defaults to black (`rgb(0, 0, 0)`) when omitted.

### GUI Facade (import('GUI'))

| Function | Description |
|----------|-------------|
| `drawLineAA(window, x1, y1, x2, y2, color, bgColor)` | Xiaolin Wu anti-aliased line |
| `drawCircleFilledAA(window, cx, cy, radius, color, bgColor)` | Filled circle with smooth edges |
| `drawCircleOutlineAA(window, cx, cy, radius, color, bgColor)` | Circle outline with smooth edges |
| `drawEllipseFilledAA(window, cx, cy, rx, ry, color, bgColor)` | Filled ellipse with smooth edges |
| `drawRoundedRectFilledAA(window, x, y, w, h, radius, color, bgColor)` | Rounded rect with smooth edges |
| `drawTriangleFilledAA(window, p0, p1, p2, color, bgColor)` | Filled triangle with smooth edges |

### Low-Level (import('gui-aa'))

Each function takes a `deps` object as its first argument (must contain at
least `{ drawLine: fn(...) }`):

| Function | Description |
|----------|-------------|
| `drawLineAA(deps, window, x1, y1, x2, y2, color, bgColor)` | Xiaolin Wu line algorithm |
| `drawCircleFilledAA(deps, window, cx, cy, radius, color, bgColor)` | Scanline circle with edge blending |
| `drawCircleOutlineAA(deps, window, cx, cy, radius, color, bgColor)` | Per-pixel distance ring |
| `drawEllipseFilledAA(deps, window, cx, cy, rx, ry, color, bgColor)` | Scanline ellipse with edge blending |
| `drawRoundedRectFilledAA(deps, window, x, y, w, h, r, color, bgColor)` | Scanline rounded rect |
| `drawTriangleFilledAA(deps, window, p0, p1, p2, color, bgColor)` | Per-pixel SDF triangle |

### Rasterizer (import('gui-raster'))

The rasterizer gains a new function alongside `drawTriangleFilled`:

| Function | Description |
|----------|-------------|
| `drawTriangleFilledAA(deps, window, p0, p1, p2, color, borderColor, bgColor)` | Triangle fill + AA edge lines |

When `deps.drawLineAA` is provided, edges use Xiaolin Wu blending.  Otherwise
falls back to the standard aliased `drawLine` for borders.

## How It Works

### Lines — Xiaolin Wu Algorithm

The classic Xiaolin Wu algorithm draws each pixel along the line's major axis,
splitting coverage between two adjacent pixels perpendicular to the axis.  Each
pixel's opacity is proportional to its distance from the mathematical line
center, producing smooth 1-pixel-wide anti-aliased lines.

### Circles & Ellipses — Scanline Edge Blending

For each horizontal scanline crossing the shape, the exact fractional x-extent
is computed analytically.  Interior pixels are drawn at full opacity.  The two
edge pixels (left and right) receive an alpha value derived from a
`smoothstep()` blend based on sub-pixel coverage, producing smooth curved
boundaries without per-pixel distance evaluation.

### Triangles — Signed Distance Field

The AA triangle rasterizer computes the signed distance from each pixel to all
three edges.  Pixels deep inside the triangle get full opacity; pixels near
edges get a `smoothstep()` blend proportional to their distance from the
nearest edge, creating smooth boundaries.

### Rounded Rectangles

Combines the scanline approach with corner radius math.  Each row computes the
horizontal inset from the rounded corner, then applies the same left/right
edge blending.  Top and bottom edges also get vertical blending.

## Performance Notes

- **Lines**: ~2x the pixel writes of a standard Bresenham line (two blended
  pixels per step instead of one).  Fast for typical UI line counts.
- **Circles/Ellipses**: Same scanline count as non-AA fill, plus 2 extra
  single-pixel draws per scanline for edge blending.  Negligible overhead.
- **Triangles (SDF)**: Per-pixel evaluation over the bounding box — slower
  than the standard scanline fill for large triangles.  Best for UI elements
  and small/medium triangles.  For 3D mesh rendering with many triangles,
  prefer the standard `drawTriangleFilled` with post-process AA or FXAA.
- **bgColor**: Passing the correct background color is important for visual
  quality.  Blending against the wrong color produces visible halos.

## Background Color

The `bgColor` parameter determines what color edge pixels blend toward.  For
opaque backgrounds, pass the background's RGB value.  For layered/transparent
rendering, you may need to read back the actual pixel color, which is not
currently supported (would require a framebuffer read).  In practice, passing
the fill color of the area behind the shape works well for most UI scenarios.
