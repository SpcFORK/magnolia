# GUI Middleware Library (GUI)

## Overview

`GUI` provides a simple cross-platform middleware layer for lightweight window work.

Backends:

- Windows via `windows` (Win32)
- Linux via `linux` (X11)
- Web/JS runtimes with a Canvas/WebGL command middleware state

This library is intentionally small and focuses on a consistent API shape.

Because Magnolia does not currently expose direct DOM built-ins, the web backend
uses a command queue model. You can optionally pass `options.webBridge` to
`createWindow(...)` to forward recorded Canvas/WebGL operations to host-side
JavaScript.

## Performance

The GUI renderer includes several performance optimizations for 2D and 3D rendering. See **[gui-performance.md](gui-performance.md)** for:

- Projection parameter caching (40–50% reduction in vertex overhead)
- Deferred lighting for culled faces (20–30% faster backface culling)
- Renderer mesh caching (100% bypass of repeated allocations)
- Circle outline trig optimization (95% reduction in per-segment trigonometry)
- Benchmarking guide and best practices

Expected overall improvement: **5–15% FPS gain** on typical hardware.

## Import

```oak
gui := import('GUI')
// alias also supported:
gui := import('gui')
```

## Backend helpers

### `backend()`

Returns one of these atoms:

- `:windows`
- `:linux`
- `:web`
- `:unknown`

### `isWindows?()`
### `isLinux?()`
### `isWeb?()`

Boolean helpers derived from `backend()`.

### `rgb(r, g, b)`

Builds a packed RGB integer color value.

## Window lifecycle

### `createWindow(title?, width?, height?, options?)`

Creates a window state object.

Useful `options` fields:

- `frameMs` - target idle frame step in milliseconds (default: `16`)
- `maxFrameDtMs` - clamps per-frame `dt` after stalls (default: `250`)
- `updateOnDispatch` - whether input/message dispatch should also trigger
    `onFrame` (default: `false`)
- Windows icon options:
    - `icon` - base icon spec used for both big/small icon when specific values
        are not provided
    - `taskbarIcon` or `iconBig` - big icon (taskbar/alt-tab)
    - `windowIcon` or `iconSmall` - small icon (titlebar)

Icon spec can be:

- an integer resource ID (for example `32512` for `IDI_APPLICATION`)
- a string path to an `.ico` file
- an object: `{ id: <int> }`, `{ path: <string> }`, or `{ handle: <int> }`

- Windows: registers a default class and creates a top-level window.
- Linux: creates a default X11 window.
- Web: creates logical state with no native host window API calls.

Returns:

- `{type: :ok, ...windowState}` on success
- `{type: :error, error: <string>, detail: <any>}` on failure

### `show(window)`

Shows/maps the window where supported.

### `setTitle(window, title)`

Sets the title where supported.

### `beginFrame(window)` / `endFrame(window)`

Optional frame batching helpers.

- On Windows, these reuse one device context (DC) across multiple draw calls
    and flush/release once per frame, which reduces flicker and intermittent
    draw stalls.
- Windows layer selection options are available via `createWindow(..., options)`:
    - `options.layer2D`: `auto` (default), `vulkan`, `opengl`, `ddraw`, `gdi`
    - `options.vulkanAuto`: allow `auto` to select Vulkan when available (default `false`)
        - `options.layer3D`: `auto` (default), `d3d9`, `cpu`, `none`
    Selected capabilities are exposed on `window.layers` when using the Windows backend.
    Vulkan selection currently bootstraps Vulkan instance/surface state but still uses a stable fallback presenter for frame display.
- On Linux/Web they are safe no-ops for API consistency.

### `close(window)`

Closes/release window resources.

- Windows: destroys the window and unlocks the pinned thread.
- Linux: destroys window and closes display.
- Web: marks the middleware state as closed.

## Web Canvas + WebGL

Web backend windows include these fields:

- `window.canvas`
- `window.webgl`
- `window.messages` (recorded middleware operations)

### Constants

- `GL_COLOR_BUFFER_BIT`
- `GL_DEPTH_BUFFER_BIT`
- `GL_TRIANGLES`

### `createCanvas(window, id?, options?)`

Configures canvas metadata (`id`, `width`, `height`, `dpr`) and records
`:canvas_create`.

### `initWebGL(window, contextName?, attrs?)`

Initializes WebGL middleware context metadata and records `:webgl_init`.

### `webglCreateShader(window, shaderType, source)`

Records shader creation and returns `{type: :ok, shader: {...}}`.

### `webglCreateProgram(window, vertexShader, fragmentShader)`

Records program creation and returns `{type: :ok, program: {...}}`.

### `webglUseProgram(window, program)`

Sets the active program and records `:webgl_use_program`.

### `webglClearColor(window, r, g, b, a)`

Sets clear color state and records `:webgl_clear_color`.

### `webglViewport(window, x, y, width, height)`

Sets viewport state and records `:webgl_viewport`.

### `webglClear(window, mask?)`

Records `:webgl_clear` (default mask is `GL_COLOR_BUFFER_BIT`).

### `webglDrawArrays(window, mode?, first?, count)`

Records `:webgl_draw_arrays` (default mode is `GL_TRIANGLES`).

### `webglFlush(window)`

Records `:webgl_flush` and returns queued commands.

## Event and loop helpers

### `poll(window)`

Pumps one backend event step and returns an event object.

Common result shapes:

- `{type: :dispatch, ...}`
- `{type: :idle}`
- `{type: :closed}`
- `{type: :error, ...}`

### `run(window, onEvent?, onFrame?)`

Runs a simple loop:

- calls `poll(window)`
- dispatches `onEvent(window, evt)` on `:dispatch`
- calls `onFrame(window, dt)` when a frame is due on `:dispatch` and `:idle`
- applies `frameMs` pacing and clamps `dt` to `maxFrameDtMs` after long stalls
- on Windows, marks resize/paint-related dispatches as urgent so the next frame is not delayed
- exits on `:closed`

## Drawing helpers

### `drawText(window, x, y, text)`

Draws text in the target window.

- Windows: `TextOutW`
- Linux: `XDrawString`
- Web: records a logical `:text` draw op in `window.messages`

### `fillRect(window, x, y, width, height, color?)`

Fills a rectangle.

- Windows: GDI brush + rectangle draw
- Linux: X11 foreground + `XFillRectangle`
- Web: records a logical `:rect` draw op in `window.messages`

### `drawLine(window, x1, y1, x2, y2, color?)`

Draws a line segment in the target window.

- Windows: GDI pen + `MoveToEx`/`LineTo`
- Linux: X11 `XDrawLine`
- Web: records a logical `:line` draw op in `window.messages`

## 2D Tooling

GUI now includes a 2D helper suite for math, transforms, camera mapping, and
shape drawing.

### 2D math and geometry

- `Vec2(x, y)`
- `Rect2(x, y, width, height)`
- `vec2Add(a, b)`
- `vec2Sub(a, b)`
- `vec2Scale(v, s)`
- `vec2Dot(a, b)`
- `vec2Len(v)`
- `vec2Normalize(v)`
- `rectTranslate(rect, dx, dy)`
- `rectContains(rect, point)`
- `rectIntersects(a, b)`

### 2D transforms and camera

- `Transform2D(options)` fields:
    - `tx`, `ty` translation
    - `r` rotation in degrees
    - `sx`, `sy` scale (or `scale` for uniform)
- `applyTransform2D(point, transform)`
- `Camera2D(options)` fields: `x`, `y`, `zoom`
- `worldToScreen2D(point, camera, window)`
- `screenToWorld2D(point, camera, window)`

### 2D drawing primitives

- `drawRect2D(window, x, y, width, height, color?, filled?)`
- `drawCircle2D(window, cx, cy, radius, color?, filled?)`
- `drawPolyline2D(window, points, color?, closed?)`
- `drawPolygon2D(window, points, color?, filled?)`
- `drawGrid2D(window, spacing?, color?, originX?, originY?)`

## 3D Renderer

GUI includes a lightweight wireframe 3D renderer that works on native backends
and uses command recording on web backends.

### `Vec3(x, y, z)`

Constructs a 3D vector object.

### `CubeMesh(size?)`

Returns a cube wireframe mesh object:

- `vertices` list
- `edges` index pairs

### `Mesh(vertices?, edges?)`

Constructs a custom wireframe mesh object.

### `GridMesh(size?, step?)`

Builds an XZ-plane wireframe grid mesh.

### `AxesMesh(length?)`

Builds an XYZ axis wireframe mesh.

### `VoxelMesh(voxels, voxelSize?)`

Builds a wireframe mesh from voxel center positions.

- `voxels`: list of objects like `{x: 0, y: 1, z: 2}`
- `voxelSize`: cube size per voxel (default `1`)

### `VoxelGrid(options?)`

Mutable voxel set helper.

Options:

- `voxelSize`
- `voxels` (initial list)

Methods:

- `set(x, y, z, value?)`
- `get(x, y, z)`
- `clear()`
- `voxels()`
- `toMesh()`

### `drawMeshWireframe(window, mesh, transform?, camera?, color?)`

Projects and draws mesh edges as 2D lines.

### `drawTriangleFilled(window, p0, p1, p2, color?)`

Fills a projected triangle using scanline rasterization.

### `drawMeshSolid(window, mesh, transform?, camera?, color?)`

Renders solid faces (triangulated quads) with a simple painter sort by depth.

Transform fields:

- `tx`, `ty`, `tz`
- `rx`, `ry`, `rz` (degrees)
- `scale`

Camera fields:

- `z` (distance offset)
- `fov` (degrees)
- `mode` (`'perspective'` or `'orthographic'`)
- `orthoScale` (orthographic zoom scale)
- `backfaceCulling` (default `true`)

Lighting fields (for solid rendering):

- `x`, `y`, `z` (light direction)
- `ambient` (default `0.22`)
- `diffuse` (default `0.78`)

### `Renderer3D(window, options?)`

Creates a convenience renderer object.

Options:

- `camera`
- `background`
- `lineColor`

Returns object methods:

- `setCamera(camera)`
- `setLight(light)`
- `setProjection(mode, orthoScale?)`
- `clear()`
- `renderMeshSolid(mesh, transform?, color?)`
- `renderCubeSolid(size?, transform?, color?)`
- `renderMesh(mesh, transform?, color?)`
- `renderCube(size?, transform?, color?)`
- `renderGrid(size?, step?, transform?, color?)`
- `renderAxes(length?, transform?)`
- `renderVoxels(voxels, voxelSize?, transform?, color?)`
- `renderVoxelGrid(grid, transform?, color?)`

## Example

```oak
gui := import('GUI')

window := gui.createWindow('Magnolia GUI', 840, 520)
if window.type = :ok {
    gui.show(window)
    gui.fillRect(window, 24, 48, 320, 180, gui.rgb(46, 120, 226))
    gui.drawText(window, 24, 26, 'Hello from Magnolia GUI')

    gui.run(window, fn(win, evt) {
        // inspect events if needed
    }, fn(win, dt) {
        // optional frame callback with delta time in seconds
    })

    gui.close(window)
}
```

### WebGL Example (Middleware Queue)

```oak
gui := import('GUI')

window := gui.createWindow('WebGL Sample', 960, 540, {
    canvasId: 'app-canvas'
})

if window.type = :ok & gui.isWeb?() -> {
    gui.createCanvas(window, 'app-canvas', {
        width: 960
        height: 540
        dpr: 1
    })

    gui.initWebGL(window, 'webgl', {
        alpha: true
        antialias: true
        depth: true
    })

    vert := gui.webglCreateShader(window, :vertex, 'attribute vec2 aPos; void main(){ gl_Position = vec4(aPos, 0.0, 1.0); }')
    frag := gui.webglCreateShader(window, :fragment, 'precision mediump float; void main(){ gl_FragColor = vec4(0.08, 0.52, 0.95, 1.0); }')
    prog := gui.webglCreateProgram(window, vert.shader, frag.shader)

    gui.webglUseProgram(window, prog.program)
    gui.webglViewport(window, 0, 0, window.width, window.height)
    gui.webglClearColor(window, 0.02, 0.03, 0.08, 1.0)
    gui.webglClear(window, gui.GL_COLOR_BUFFER_BIT | gui.GL_DEPTH_BUFFER_BIT)
    gui.webglDrawArrays(window, gui.GL_TRIANGLES, 0, 3)

    queued := gui.webglFlush(window)
    println(string(queued))
}
```

## Module reference

The GUI middleware is implemented across several focused modules. See the module pages for API details and examples.

- [gui-2d](gui-2d.md) — 2D math, transforms, camera mapping, and 2D drawing helpers.
- [gui-3dmath](gui-3dmath.md) — 3D math, vector/rotation utilities and projections.
- [gui-draw](gui-draw.md) — Cross-platform drawing primitives (text, lines, rects).
- [gui-mesh](gui-mesh.md) — Mesh/voxel builders and helpers used by the renderer.
- [gui-render](gui-render.md) — High-level 3D renderer and `Renderer3D` convenience API.
- [gui-raster](gui-raster.md) — Scanline rasterization, triangle fill, and culling.
- [gui-native-win](gui-native-win.md) — Win32-specific window lifecycle and helpers.
- [gui-native-linux](gui-native-linux.md) — X11-specific window lifecycle and helpers.
- [gui-web](gui-web.md) — Canvas/WebGL middleware and command-queue helpers for web runtimes.


## Samples

- `samples/gui-sample.oak` - cross-platform GUI quickstart
- `samples/gui-game.oak` - bouncing-box mini game using GUI middleware
- `samples/gui-3d.oak` - rotating wireframe cube using GUI 3D renderer

## Module Split

- `import('GUI')` remains the primary API and is backward compatible.
- Mesh and voxel builders are factored into `import('gui-mesh')` internally.
- The renderer factory is factored into `import('gui-render')` internally.
- Transform/projection math is factored into `import('gui-3dmath')` internally.
- Shading, culling, and raster routines are factored into `import('gui-raster')` internally.
- Web Canvas/WebGL middleware queue logic is factored into `import('gui-web')` internally.
- Win32 lifecycle/frame internals are factored into `import('gui-native-win')` internally.
- Linux lifecycle/event-loop internals are factored into `import('gui-native-linux')` internally.
- Cross-platform drawing primitives are factored into `import('gui-draw')` internally.
- 2D math/shape tooling is factored into `import('gui-2d')` internally.
- Optional direct mesh module imports:
    - `Mesh(vertices, edges)`
    - `GridMesh(size, step)`
    - `AxesMesh(length)`
    - `VoxelMesh(voxels, voxelSize)`
    - `VoxelGrid(options)`
- Optional direct renderer module import:
    - `Renderer3D(deps, window, options)`
- Optional direct 3D math module imports:
    - `degToRad(deg)`
    - `Vec3(x, y, z)`
    - `transformPoint(v, transform)`
    - `projectPoint(window, p, camera)`
    - `transformVertices(vertices, transform, i, out)`
- Optional direct raster module imports:
    - `drawTriangleFilled(deps, window, p0, p1, p2, color)`
    - `drawMeshSolid(deps, window, mesh, transform, camera, color, light)`
    - `drawMeshWireframe(deps, window, mesh, transform, camera, color)`
- Optional direct web module imports:
    - `createWindowState(title, width, height, frameMs, updateOnDispatch, options)`
    - `createCanvas(window, id, options)`
    - `initWebGL(window, contextName, attrs)`
    - `webglCreateShader(window, shaderType, source)`
    - `webglCreateProgram(window, vertexShader, fragmentShader)`
    - `webglUseProgram(window, program)`
    - `webglClearColor(window, r, g, b, a)`
    - `webglViewport(window, x, y, width, height)`
    - `webglClear(window, mask, colorBufferBit)`
    - `webglDrawArrays(window, mode, first, count, trianglesMode)`
    - `webglFlush(window)`
