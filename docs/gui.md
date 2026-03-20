# GUI Middleware Library (GUI)

## Overview

`GUI` provides a simple cross-platform middleware layer for lightweight window work.

Backends:

- Windows via `windows` (Win32)
- Linux via `linux` (X11)
- Web/JS runtimes with a Canvas/WebGL command middleware state

This library is intentionally small and focuses on a consistent API shape.

## Split modules

GUI internals are split into focused modules that can also be imported directly:

- [gui-color](gui-color.md)
- [gui-events](gui-events.md)
- [gui-graph](gui-graph.md)
- [gui-form](gui-form.md)
- [gui-loop](gui-loop.md)

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

### `rgba(r, g, b, a, background?)`

Builds a packed RGB color by composing an RGBA-style input over
`background` using alpha `a` in `[0, 1]`.

This is equivalent to `opacity(rgb(r, g, b), a, background)`.

### `opacity(color, amount, background?)`

Pre-composes `color` over `background` using an opacity in `[0, 1]` and
returns a packed RGB integer that can be passed to the native drawing helpers.

This is useful on native backends that currently draw with opaque packed colors
instead of full RGBA surfaces.

## Window lifecycle

### `createWindow(title?, width?, height?, options?)`

Creates a window state object.

Useful `options` fields:

- `frameMs` - target idle frame step in milliseconds (default: `16`)
- `maxFrameDtMs` - clamps per-frame `dt` after stalls (default: `250`)
- `updateOnDispatch` - whether input/message dispatch should also trigger
    `onFrame` (default: `false`)
- `className` (Windows) - override the Win32 class name used during
    registration. Defaults to a generated unique class name.
- `dllLoadMode` (Windows) - DLL probe strategy for renderer capability checks:
    `sync` (default) probes during `createWindow`; `async` starts probing in
    the background so window creation stays responsive.
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

### `hide(window)`

Hides/unmaps the window where supported.

### `move(window, x, y)`

Moves the window top-left position where supported.

- Windows: uses Win32 `SetWindowPos` without resizing.
- Linux: uses X11 `XMoveWindow`.
- Web: records `:window_move` middleware op and updates window state.

### `resize(window, width, height)`

Resizes the window where supported.

- Windows: uses Win32 `SetWindowPos` with current position.
- Linux: uses X11 `XResizeWindow`.
- Web: updates middleware state and records `:window_resize`.

### `scale(window, scaleX, scaleY?)`

Scales the current window size by multipliers and applies `resize(...)`.

- `scaleY` is optional; when omitted it uses `scaleX`.
- Output dimensions are clamped to at least `1x1`.

### `fullscreen(window, enabled?)`

Toggles fullscreen-like behavior (default `enabled = true`).

- Windows: maximize/restore via `ShowWindow`.
- Linux: best-effort fullscreen by moving to `(0,0)` and resizing to display dimensions.
- Web: updates middleware fullscreen state and records `:window_fullscreen`.

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
    - `options.dllLoadMode`: `sync` (default) or `async` for background DLL probing
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

### Event-Bus integration

Each `:ok` window now carries a lazily initialized event bus at `window.eventBus`.
The GUI module exposes thin helpers around this bus:

- `eventBus(window)` returns the bus instance
- `on(window, event, handler)` subscribes and returns a token
- `once(window, event, handler)` one-shot subscription, returns a token
- `off(window, event, tokenOrHandler)` unsubscribes and returns removed count
- `emit(window, event, payload, onDone?)` emits sync/async based on callback presence
- `listenerCount(window, event)` returns active listener count
- `clearListeners(window, event?)` clears one event or all listeners
- `onDispatch(window, fn(step){...})` convenience wrapper over `on(window, :dispatch, ...)`
- `onceDispatch(window, fn(step){...})` one-shot dispatch subscription
- `onKeyDownEvent(window, fn(key, evt){...})` filters dispatch events to `:keyDown`
- `onKeyUpEvent(window, fn(key, evt){...})` filters dispatch events to `:keyUp`
- `onceKeyDownEvent(window, fn(key, evt){...})` one-shot `:keyDown` subscription
- `onceKeyUpEvent(window, fn(key, evt){...})` one-shot `:keyUp` subscription

Lifecycle aliases:

- `onRunStart/onceRunStart(window, handler)` for `:runStart`
- `onIdle/onceIdle(window, handler)` for `:idle`
- `onFrame/onceFrame(window, handler)` for `:frame`
- `onClosing/onceClosing(window, handler)` for `:closing`
- `onClosed/onceClosed(window, handler)` for `:closed`

Built-in events emitted by GUI runtime:

- `:runStart` when `run(...)` begins on a window
- `:dispatch` for each polled dispatch step
- `:idle` for each idle step
- `:frame` whenever `onFrame(window, dt)` is invoked (payload includes `dt` and `timeNs`)
- `:closing` when `close(window)` begins
- `:closed` after `close(window)` completes or when run loop receives closed step

### Typed dispatch helpers (`On<Event>`)

For native Windows message handling, GUI exposes typed subscribers that wrap
`on(window, :dispatch, ...)` and filter by message type for you.

- `onMouseMove(window, fn(mx, my) {...})`
- `onLButtonDown(window, fn(mx, my) {...})`
- `onLButtonUp(window, fn(mx, my) {...})`
- `onRButtonDown(window, fn(mx, my) {...})`
- `onRButtonUp(window, fn(mx, my) {...})`
- `onKeyDown(window, fn(vk) {...})`
- `onKeyUp(window, fn(vk) {...})`
- `onChar(window, fn(code) {...})`
- `onResize(window, fn(width, height) {...})`

Each helper returns the same subscription token as `on(...)`, so you can remove
handlers with `off(window, :dispatch, token)`.

## Drawing helpers

### `drawText(window, x, y, text, color?)`

Draws text in the target window.

Use `setFont(window, fontSpec)` to control font family/size/weight used by
subsequent text draws.

When `color` is omitted, GUI uses its default light text color. You can pass a
packed RGB color from `rgb(...)` or a pre-composed color from
`opacity(...)`.

- Windows: `TextOutW`
- Linux: `XDrawString`
- Web: records a logical `:text` draw op in `window.messages` (includes `font`
    CSS value when set)

### `setFont(window, fontSpec)`

Sets the active text font for a window.

`fontSpec` fields:

- `family` - font family name (default: `Segoe UI` on Windows)
- `size` - font size in px (default: `16`)
- `weight` - numeric weight (default: `400`, bold is usually `700`)
- `italic` - boolean
- `underline` - boolean
- `strikeOut` - boolean
- `css` - optional web CSS font shorthand override (web backend only)

Example:

```oak
gui.setFont(window, {
    family: 'Consolas'
    size: 18
    weight: 700
})
gui.drawText(window, 24, 32, 'Font-aware text', gui.rgb(248, 232, 242))
```

### `clearFont(window)`

Clears the custom font and returns to backend defaults.

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

## Graphing helpers

GUI now includes lightweight graph drawing helpers built on top of
`drawLine(...)`, `fillRect(...)`, and `drawText(...)`.

### `graphRange(values, options?)`

Computes a display range for numeric series.

- `options.min` optional fixed minimum
- `options.max` optional fixed maximum
- `options.padding` optional padding applied to both sides

Returns `{min, max, span}`.

### `graphMapX(index, count, x, width)`

Maps a sample index to an X coordinate in a graph rect.

### `graphMapY(value, y, height, range)`

Maps a numeric value to a Y coordinate in a graph rect using a
`graphRange(...)` result.

### `drawGraphAxes(window, x, y, width, height, options?)`

Draws graph background, grid, and border axes.

- `options.gridColor`
- `options.axisColor`
- `options.backgroundColor`
- `options.xTicks` (default `5`)
- `options.yTicks` (default `4`)
- `options.showGrid` (default `true`)

### `drawLineGraph(window, x, y, width, height, values, options?)`

Draws a line graph with optional points and min/max labels.

- `options.lineColor`
- `options.pointColor`
- `options.showPoints` (default `true`)
- `options.showLabels` (default `true`)
- `options.min`, `options.max`, `options.rangePadding`
- `options.axis` nested options forwarded to `drawGraphAxes(...)`

### `drawBarGraph(window, x, y, width, height, values, options?)`

Draws a bar graph.

- `options.barColor`
- `options.barBorderColor`
- `options.barGap` (default `2`)
- `options.showLabels` (default `true`)
- `options.min`, `options.max`, `options.rangePadding`
- `options.axis` nested options forwarded to `drawGraphAxes(...)`

### `drawSparkline(window, x, y, width, height, values, options?)`

Draws a compact line-only graph for dashboards and small stat cards.

- `options.lineColor`
- `options.backgroundColor`
- `options.min`, `options.max`, `options.rangePadding`

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
    bg := gui.rgb(12, 18, 30)
    gui.fillRect(window, 0, 0, window.width, window.height, bg)
    gui.fillRect(window, 24, 48, 320, 180, gui.rgba(46, 120, 226, 0.9, bg))
    gui.drawText(window, 24, 26, 'Hello from Magnolia GUI')
    gui.drawText(window, 24, 250, 'Close the window to exit.', gui.rgba(248, 232, 242, 0.75, bg))

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

### Logging Sample

```oak
gui := import('GUI')

window := gui.createWindow('GUI Logging Sample', 860, 460, {
    frameMs: 16
    layer2D: 'auto'
})

if window.type = :ok {
    state := { tick: 0, entries: [] }
    gui.show(window)

    gui.run(window, fn(_, evt) {
        if evt.type = :dispatch & (state.tick % 120 = 0) ->
            state.entries <- state.entries << ('EVENT tick=' + string(state.tick))
    }, fn(win, _dt) {
        state.tick <- state.tick + 1
        if state.tick % 90 = 0 ->
            state.entries <- state.entries << ('TRACE frame=' + string(state.tick))

        if gui.beginFrame(win).type = :ok -> {
            gui.fillRect(win, 0, 0, win.width, win.height, gui.rgb(18, 22, 34))
            gui.drawText(win, 20, 24, 'GUI Logging Sample')
            gui.drawText(win, 20, 52, 'entries=' + string(len(state.entries)))
            gui.endFrame(win)
        }
    })

    gui.close(window)
}
```

See `samples/gui-logging.oak` for a full runnable version with stdout logging.

### Test Sample

```oak
{ new: new } := import('test')

fn formatLogLine(ts, level, message) '[' + string(ts) + '] ' + level + ' ' + message
fn shouldLogTick?(tick, every) (every > 0) & (tick % every = 0)

suite := new('GUI Logging Sample Tests')
suite.eq('format line', formatLogLine(42, 'INFO', 'ready'), '[42] INFO ready')
suite.eq('tick schedule', [shouldLogTick?(90, 90), shouldLogTick?(91, 90)], [true, false])
suite.report()
suite.exit()
```

See `samples/test_gui_logging.oak` for the complete test sample.

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

- `samples/gui-sample.oak` - cross-platform GUI quickstart; on Windows it also displays the requested 2D layer, active presenter, and fallback state
- `samples/gui-game.oak` - bouncing-box mini game using GUI middleware with frame-rate independent motion and Windows presenter-state overlay
- `samples/gui-logging.oak` - GUI logging sample with frame/event logs rendered in-window and mirrored to stdout
- `samples/test_gui_logging.oak` - focused test sample for GUI logging helper functions using the `test` library
- `samples/gui-3d.oak` - rotating wireframe cube using GUI 3D renderer
- `samples/windows-higher-renderers.oak` - Windows renderer probe for Vulkan/OpenGL/DDraw/D3D capability and staged fallback metadata
- `samples/windows-2d-layer-hotload.oak` - Windows-only hotload demo that recreates the window every 2 seconds to cycle `gdi`, `ddraw`, `opengl`, and `vulkan` requests while preserving one animated scene
- `samples/windows-2d-layer-hotload-game.oak` - Windows-only hotload demo that recreates the window every 2 seconds while preserving the bouncing-box game scene across `gdi`, `ddraw`, `opengl`, and `vulkan` requests

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
