# gui-2d

2D math, camera helpers, transforms, and simple 2D drawing utilities used by the GUI stack.

Key exports

- `Vec2(x, y)` — 2D vector constructor
- `Rect2(x, y, width, height)` — axis-aligned rectangle
- `Transform2D(options)` — translation/rotation/scale transform
- `Camera2D(options)` — simple 2D camera with `worldToScreen2D`/`screenToWorld2D`
- 2D draw helpers: `drawRect2D`, `drawCircle2D`, `drawPolyline2D`, `drawPolygon2D`
- `drawEllipse2D(deps, window, cx, cy, rx, ry, color, filled?, borderColor?)` — ellipse with independent radii
- `drawArc2D(deps, window, cx, cy, radius, startAngle, endAngle, color)` — arc segment (degrees)
- `drawTriangle2D(deps, window, x1, y1, x2, y2, x3, y3, color, filled?, borderColor?)` — triangle shorthand
- `drawRoundedRect2D(deps, window, x, y, w, h, radius, color, filled?, borderColor?)` — rectangle with rounded corners
- `drawStar2D(deps, window, cx, cy, outerR, innerR, points, color, filled?, borderColor?)` — regular star polygon
- `drawBezier2D(deps, window, points, color, steps?)` — quadratic (3 pts) or cubic (4 pts) bezier curve
- `drawRing2D(deps, window, cx, cy, outerR, innerR, color)` — filled annulus (ring)

## Border Color

All fillable 2D shapes (`drawRect2D`, `drawCircle2D`, `drawPolygon2D`, `drawEllipse2D`, `drawTriangle2D`, `drawRoundedRect2D`, `drawStar2D`) accept an optional `borderColor` parameter as their last argument. When provided on a filled shape, the shape is first filled with `color`, then its outline is drawn with `borderColor`.

```oak
gui := import('GUI')

// Filled red rectangle with a white border
gui.drawRect2D(window, 50, 50, 200, 100, gui.rgb(255, 0, 0), true, gui.rgb(255, 255, 255))

// Filled blue circle with a yellow border
gui.drawCircle2D(window, 200, 200, 60, gui.rgb(0, 0, 255), true, gui.rgb(255, 255, 0))

// Without borderColor — works exactly as before
gui.drawRect2D(window, 50, 50, 200, 100, gui.rgb(255, 0, 0), true)
```

Usage

Import and use as supporting utilities for GUI drawing and UI coordinate transforms.

```oak
g2 := import('gui-2d')
cam := g2.Camera2D({x:0, y:0, zoom:1})
screen := g2.worldToScreen2D({x:10,y:5}, cam, window)
```
