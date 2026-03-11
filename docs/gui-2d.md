# gui-2d

2D math, camera helpers, transforms, and simple 2D drawing utilities used by the GUI stack.

Key exports

- `Vec2(x, y)` — 2D vector constructor
- `Rect2(x, y, width, height)` — axis-aligned rectangle
- `Transform2D(options)` — translation/rotation/scale transform
- `Camera2D(options)` — simple 2D camera with `worldToScreen2D`/`screenToWorld2D`
- 2D draw helpers: `drawRect2D`, `drawCircle2D`, `drawPolyline2D`, `drawPolygon2D`

Usage

Import and use as supporting utilities for GUI drawing and UI coordinate transforms.

```oak
g2 := import('gui-2d')
cam := g2.Camera2D({x:0, y:0, zoom:1})
screen := g2.worldToScreen2D({x:10,y:5}, cam, window)
```
