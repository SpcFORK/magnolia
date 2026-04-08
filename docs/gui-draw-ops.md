# gui-draw-ops — Unified Drawing API

`import('gui-draw-ops')` provides a single `draw()` entry point that replaces
dozens of separate drawing functions with one composable, options-object-based
call.  Instead of remembering which function handles which combination of
shape + fill + AA + lighting + dash style, you describe what you want once:

```oak
gui := import('GUI')
{ rgb: rgb } := import('gui-color')

bg := rgb(30, 30, 30)

// Unified draw() — one function for everything
gui.draw(window, {
    shape: :circle
    cx: 200, cy: 200, r: 50
    color: rgb(255, 100, 100)
    filled: true
    aa: true
    bgColor: bg
})
```

The old separate functions (`drawCircle2D`, `drawCircleFilledAA`,
`drawMeshLit`, etc.) remain available and unchanged.  `draw()` delegates to
them internally based on the options you provide.

## Quick Start

### 2D Shapes

```oak
// Anti-aliased filled rounded rect
gui.draw(window, {
    shape: :roundedRect
    x: 50, y: 50, width: 200, height: 80
    radius: 12
    color: rgb(100, 200, 255)
    filled: true
    aa: true
    bgColor: bg
})

// Dashed line
gui.draw(window, {
    shape: :line
    x1: 10, y1: 10, x2: 300, y2: 200
    color: rgb(255, 255, 100)
    dash: { len: 8, gap: 4 }
})

// Thick line
gui.draw(window, {
    shape: :line
    x1: 10, y1: 50, x2: 300, y2: 50
    color: rgb(255, 255, 255)
    thickness: 3
})

// Star with border
gui.draw(window, {
    shape: :star
    cx: 400, cy: 300
    outerR: 60, innerR: 25, points: 5
    color: rgb(255, 200, 50)
    filled: true
    borderColor: rgb(200, 150, 0)
})
```

### 3D Meshes

```oak
camera := { z: 5, fov: 90 }

// Basic solid mesh (simple directional light)
gui.draw(window, {
    shape: :mesh
    mesh: gui.CubeMesh(2)
    transform: { rx: 0.5, ry: 0.8, rz: 0 }
    camera: camera
    color: rgb(100, 150, 255)
    light: { x: 0.3, y: 0.7, z: -1, ambient: 0.2, diffuse: 0.8 }
})

// Fully lit mesh (Blinn-Phong multi-light)
scene := gui.LightScene({
    lights: [
        gui.DirectionalLight({ x: 0.5, y: -1, z: 0.3, intensity: 1 })
        gui.AmbientLight({ intensity: 0.15 })
        gui.PointLight({ px: 2, py: 3, pz: -1, intensity: 0.8, radius: 8 })
    ]
    globalAmbient: 0.05
})
mat := gui.Material({ diffuse: 0.9, specular: 0.4, shininess: 64 })

gui.draw(window, {
    shape: :mesh
    mesh: gui.SphereMesh(1.5, 16, 12)
    transform: { rx: 0, ry: 0.3, rz: 0, tx: 2 }
    camera: camera
    color: rgb(220, 180, 140)
    lighting: scene       // <- triggers lit rendering
    material: mat
    borderColor: rgb(50, 40, 30)
})

// Wireframe mesh
gui.draw(window, {
    shape: :meshWire
    mesh: gui.CubeMesh(2)
    transform: { rx: 0.3, ry: 0.6, rz: 0 }
    camera: camera
    color: rgb(0, 255, 128)
})
```

### Batch Drawing

```oak
gui.drawBatch(window, [
    { shape: :rect, x: 10, y: 10, width: 100, height: 50, color: rgb(255, 0, 0), filled: true }
    { shape: :circle, cx: 200, cy: 100, r: 30, color: rgb(0, 255, 0), filled: true, aa: true, bgColor: bg }
    { shape: :line, x1: 300, y1: 50, x2: 400, y2: 150, color: rgb(0, 0, 255), aa: true, bgColor: bg }
    { shape: :text, x: 10, y: 200, text: 'Hello!', color: rgb(255, 255, 255) }
])
```

## API Reference

### `draw(window, op)`

Draws a single primitive. `op` is an object with a required `shape` field plus
shape-specific geometry and optional effect modifiers.

### `drawBatch(window, ops)`

Draws a list of `op` objects in sequence.

---

## Shape Reference

### Common Fields (all shapes)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `shape` | atom | *required* | Shape type (see below) |
| `color` | int | `0` (black) | Fill/stroke color (packed RGB) |
| `borderColor` | int | `?` | Outline color |
| `filled` | bool | `true` | Fill the shape (vs outline only) |
| `aa` | bool | `false` | Enable anti-aliasing |
| `bgColor` | int | `0` | Background color for AA blending |

### `:line`

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `x1, y1` | num | `0` | Start point |
| `x2, y2` | num | `0` | End point |
| `thickness` | num | `?` | Thick line width (>1 to activate) |
| `dash` | obj | `?` | `{ len: 6, gap: 4 }` for dashed |

**Priority**: `dash` > `thickness` > `aa` > plain

### `:rect`

| Field | Type | Default |
|-------|------|---------|
| `x, y` | num | `0` |
| `width, height` | num | `0` |

### `:circle`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `r` (or `radius`) | num | `0` |

### `:ellipse`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `rx, ry` | num | `0` |

### `:triangle`

Two input modes:

**Point objects:** `p0, p1, p2` — each `{ x, y }`

**Scalars:** `x1, y1, x2, y2, x3, y3`

### `:roundedRect`

| Field | Type | Default |
|-------|------|---------|
| `x, y` | num | `0` |
| `width, height` | num | `0` |
| `radius` | num | `4` |

### `:polygon`

| Field | Type | Default |
|-------|------|---------|
| `points` | list | `[]` |

### `:polyline`

| Field | Type | Default |
|-------|------|---------|
| `points` | list | `[]` |
| `closed` | bool | `false` |

### `:arc`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `r` (or `radius`) | num | `0` |
| `startAngle` | num | `0` |
| `endAngle` | num | `360` |

### `:star`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `outerR, innerR` | num | `0` |
| `points` | int | `5` |

### `:bezier`

| Field | Type | Default |
|-------|------|---------|
| `points` | list | `[]` |
| `steps` | int | `?` (auto) |

### `:ring`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `outerR, innerR` | num | `0` |

### `:cross`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `size` | num | `0` |
| `thickness` | num | `2` |

### `:diamond`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `width, height` | num | `0` |

### `:arrow`

| Field | Type | Default |
|-------|------|---------|
| `x1, y1, x2, y2` | num | `0` |
| `headSize` | num | `?` (auto) |

### `:capsule`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `width, height` | num | `0` |

### `:sector`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `r` (or `radius`) | num | `0` |
| `startAngle, endAngle` | num | `0, 360` |

### `:regularPolygon`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `r` (or `radius`) | num | `0` |
| `sides` | int | `6` |

### `:spiral`

| Field | Type | Default |
|-------|------|---------|
| `cx, cy` | num | `0` |
| `startRadius` | num | `0` |
| `growth` | num | `1` |
| `turns` | num | `3` |

### `:grid`

| Field | Type | Default |
|-------|------|---------|
| `spacing` | num | `32` |
| `originX, originY` | num | `0` |

### `:text`

| Field | Type | Default |
|-------|------|---------|
| `x, y` | num | `0` |
| `text` | string | `''` |

### `:mesh` (3D solid or lit)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `mesh` | obj | `?` | Mesh object (vertices, faces, edges) |
| `transform` | obj | `?` | `{rx, ry, rz, scale, tx, ty, tz}` |
| `camera` | obj | `?` | `{z, fov, mode, orthoScale, ...}` |
| `light` | obj | `?` | Simple directional light |
| `lighting` | obj | `?` | LightScene for Blinn-Phong (*) |
| `material` | obj | `?` | Material for lit rendering (*) |

(*) When `lighting` is provided, the renderer uses `drawMeshLit` (full
Blinn-Phong multi-light).  Otherwise it uses `drawMeshSolid` with the simple
`light` parameter.

### `:meshWire` (3D wireframe)

| Field | Type | Default |
|-------|------|---------|
| `mesh` | obj | `?` |
| `transform` | obj | `?` |
| `camera` | obj | `?` |

## How It Works

`draw()` is a thin dispatcher that reads `op.shape` and routes to the
appropriate existing drawing function, selecting the best variant based on
the effect flags present:

```
op.shape = :circle
  + op.aa = true, op.filled = true  → drawCircleFilledAA(...)
  + op.aa = true, op.filled = false → drawCircleOutlineAA(...)
  + op.aa = false                   → drawCircle2D(...)

op.shape = :line
  + op.dash present                 → drawDashedLine2D(...)
  + op.thickness > 1                → drawThickLine2D(...)
  + op.aa = true                    → drawLineAA(...)
  + none of above                   → drawLine(...)

op.shape = :mesh
  + op.lighting present             → drawMeshLit(...)
  + op.lighting absent              → drawMeshSolid(...)
```

No new rendering code is added — the dispatcher composes the existing
primitives.  The deps object is allocated once (lazily on first call) and
reused for all subsequent `draw()` calls.

## Migration Guide

The old API remains fully functional.  To adopt `draw()`, replace:

```oak
// Before
drawCircle2D(window, 100, 100, 40, red, true, ?)
drawCircleFilledAA(window, 100, 100, 40, red, bg)
drawMeshLit(window, mesh, transform, camera, white, scene, material, ?)

// After
draw(window, { shape: :circle, cx: 100, cy: 100, r: 40, color: red, filled: true })
draw(window, { shape: :circle, cx: 100, cy: 100, r: 40, color: red, filled: true, aa: true, bgColor: bg })
draw(window, { shape: :mesh, mesh: mesh, transform: transform, camera: camera, color: white, lighting: scene, material: material })
```
