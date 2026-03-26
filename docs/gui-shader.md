# CPU Shader Engine (gui-shader)

## Overview

`gui-shader` is a CPU-based shader engine and facade module that re-exports all
shader sub-modules and provides the core shader runtime: constructors, a shader
registry, pixel-level and line-level renderers, gradient/column helpers, a pixel
buffer, and multi-pass composition.

### Sub-modules

| Module | Docs | Purpose |
|---|---|---|
| gui-shader-math | [gui-shader-math.md](gui-shader-math.md) | Constants, math helpers, easing, 2D/3D vectors |
| gui-shader-color | [gui-shader-color.md](gui-shader-color.md) | Packed RGB, HSL/HSV, blending |
| gui-shader-noise | [gui-shader-noise.md](gui-shader-noise.md) | Hashing, noise2D, fbm |
| gui-shader-sdf | [gui-shader-sdf.md](gui-shader-sdf.md) | SDF primitives, boolean ops, patterns |
| gui-shader-codegen | [gui-shader-codegen.md](gui-shader-codegen.md) | GLSL/HLSL codegen, WebGL, offline compilation |

All public symbols from sub-modules are re-exported at the top level, so a
single import gives access to the full API.

## Import

```oak
sh := import('gui-shader')
```

## Quick Start

```oak
sh := import('gui-shader')
gui := import('GUI')

// 1. Define a fragment function: (x, y, w, h, time, uniforms) -> color
myFragment := fn(x, y, w, h, t, u) {
    sh.hsl2rgb(sh.fract(t * 0.1 + float(x) / float(w)), 1.0, 0.5)
}

// 2. Create a shader
s := sh.Shader(myFragment, { resolution: 4 })

// 3. Render inside your draw loop
sh.render(window, s, 0, 0, 320, 240)
```

---

## Shader Core

### `cs Shader(fragment?, opts?)`

Constructs a new shader object and auto-registers it in the global registry.

**Parameters:**

| Name | Type | Description |
|---|---|---|
| `fragment?` | `fn(x, y, w, h, t, u)` or `?` | Per-pixel fragment function. Returns a packed color. |
| `opts?` | object or `?` | Optional configuration (see below). |

**Options (`opts`):**

| Key | Type | Default | Description |
|---|---|---|---|
| `draw` | `fn(window, x, y, w, h, t, uniforms)` or `?` | `?` | Custom batch-draw function. When set, `render` calls this instead of the built-in line renderer. |
| `resolution` | int | `4` | Pixel step size for the built-in renderers. Lower values = higher quality, slower. |
| `uniforms` | object | `{}` | User-defined uniform values passed to fragment/draw functions. |

**Returns:** A shader object with these fields:

| Field | Description |
|---|---|
| `fragment` | The fragment function. |
| `draw` | The custom draw function (or `?`). |
| `resolution` | Current pixel resolution step. |
| `uniforms` | Uniform values object. |
| `_startNano` | Internal clock start (nanoseconds). |
| `_pausedAt` | Pause timestamp (or `?` if running). |
| `_pauseAccum` | Accumulated pause duration. |
| `_frameStart` | Frame start timestamp for dt calculation. |
| `_dt` | Delta time of last frame (seconds). |
| `_frameCount` | Number of completed frames. |

### `fn elapsed(shader)`

Returns the elapsed active time in **seconds** (float), excluding any paused
duration.

### `fn pause(shader)`

Pauses the shader clock. No-op if already paused.

### `fn resume(shader)`

Resumes the shader clock. Accumulated pause time is subtracted from `elapsed`.
No-op if not paused.

### `fn reset(shader)`

Resets the shader's clock to zero and clears pause state.

### `fn isRunning(shader)`

Returns `true` if the shader is not paused, `false` otherwise.

### `fn beginFrame(shader)`

Marks the start of a frame. Call before your draw logic.

### `fn endFrame(shader)`

Marks the end of a frame. Computes `_dt` and increments `_frameCount`. Returns
the frame's delta time in seconds.

### `fn dt(shader)`

Returns the delta time of the most recent frame (seconds).

### `fn frameCount(shader)`

Returns the total number of completed frames.

### `fn setUniform(shader, key, value)`

Sets a uniform value: `shader.uniforms.(key) <- value`.

### `fn getUniform(shader, key)`

Returns `shader.uniforms.(key)`.

### `fn setResolution(shader, res)`

Updates the shader's pixel resolution step.

---

## Shader Registry

A module-level list that tracks all live shader instances.

### `fn registeredCount()`

Returns the number of currently registered shaders.

### `fn unregisterShader(shader)`

Removes a specific shader from the registry.

### `fn clearAll()`

Resets every registered shader's clock and frame count to zero without removing
them from the registry.

### `fn destroyAll()`

Resets all shaders and empties the registry.

---

## Render Dispatch

### `fn render(window, shader, x, y, w, h)`

Primary render entry point. Computes `elapsed(shader)`, then:

- If `shader.draw` is set, calls `shader.draw(window, x, y, w, h, t, uniforms)`.
- Otherwise, falls back to `renderShaderLines`.

---

## Multi-Pass Composition

### `cs ShaderPass(shader, x, y, w, h)`

Bundles a shader with a screen-space rectangle for multi-pass rendering.

| Field | Description |
|---|---|
| `shader` | Reference to a `Shader` instance. |
| `x`, `y` | Top-left origin of the pass. |
| `w`, `h` | Dimensions of the pass. |

### `fn composePasses(window, passes)`

Renders an array of `ShaderPass` objects in order, calling `render` for each.

---

## Built-in Renderers

### `fn renderShader(window, shader, x, y, w, h)`

Full per-pixel renderer. Iterates every `resolution`-sized cell, calls
`shader.fragment(rx, ry, w, h, t, u)`, and draws a filled rectangle for each
non-null result.

### `fn renderShaderLines(window, shader, x, y, w, h)`

Line-based renderer (faster). Samples the fragment at the horizontal midpoint of
each row and draws a full-width horizontal line in that color.

### `fn renderGradientBands(window, gradientFn, x, y, w, h, time, bands?)`

Renders a vertical gradient by dividing the height into `bands` (default 16)
horizontal bands, sampling `gradientFn(bandMidY, h, time)` for each.

### `fn renderGradient(window, gradientFn, x, y, w, h, time)`

Shorthand for `renderGradientBands` with 16 bands.

### `fn renderHorizontalBands(window, gradientFn, x, y, w, h, time, bands?)`

Like `renderGradientBands` but splits into vertical columns instead.

### `fn renderColumns(window, columns, ox, oy, h)`

Draws an array of rain-style columns. Each column object has `{ x, headY,
length, color }`. Handles wrap-around when `headY - length < 0`.

### `fn updateColumns(columns, h, t, rate?)`

Advances each column's `headY` based on elapsed time, speed, and wrap phase.
Default scroll rate is 40.0.

---

## Pixel Buffer

An off-screen pixel buffer for software rendering.

### `fn createBuffer(w, h)`

Returns `{ width, height, data: {} }`.

### `fn clearBuffer(buf, color?)`

Fills every pixel in the buffer with `color` (default `0`).

### `fn setPixel(buf, x, y, color)`

Sets a single pixel. No-op if out of bounds.

### `fn getPixel(buf, x, y)`

Returns the color at `(x, y)`, or `0` if out of bounds or unset.

### `fn renderBuffer(window, buf, ox, oy)`

Blits the pixel buffer to screen at offset `(ox, oy)` using 1×1 filled rects.

### `fn renderShaderToBuffer(buf, shader)`

Evaluates `shader.fragment` for every cell in the buffer and writes results via
`setPixel`.

---

## Re-exported Sub-module API

All functions below are re-exported from the corresponding sub-module and
available directly on the `gui-shader` import.

### Math (gui-shader-math)

**Constants:** `PI`, `TAU`, `HALF_PI`, `E`, `DEG2RAD`, `RAD2DEG`, `SQRT2`

**Helpers:**
`fract`, `mod`, `sign`, `abs2`, `clamp`, `saturate`, `lerpFloat`,
`inverseLerp`, `remap`, `step`, `smoothstep`, `smootherstep`, `min2`, `max2`,
`sqr`, `sqrt`, `lerp`, `atan2`, `pingpong`, `degToRad`, `radToDeg`

**Easing:**
`easeInQuad`, `easeOutQuad`, `easeInOutQuad`, `easeInCubic`, `easeOutCubic`,
`easeInOutCubic`, `easeInSine`, `easeOutSine`, `easeInOutSine`, `easeInExpo`,
`easeOutExpo`, `easeOutElastic`, `easeOutBounce`

**2D Vectors:**
`vec2`, `dot2`, `length2`, `distance2`, `normalize2`, `rotate2`, `scale2`,
`add2`, `sub2`, `lerp2`, `negate2`, `abs2v`, `min2v`, `max2v`, `floor2`,
`fract2`, `reflect2`, `toPolar`, `fromPolar`

**3D Vectors:**
`vec3`, `add3`, `sub3`, `scale3`, `dot3`, `length3`, `distance3`, `normalize3`,
`cross3`, `lerp3`, `negate3`, `reflect3`

### Color (gui-shader-color)

`packRGB`, `unpackRGB`, `colorR`, `colorG`, `colorB`, `mix`, `mix3`,
`brighten`, `darken`, `invert`, `grayscale`, `overlay`, `hsl2rgb`, `rgb2hsl`,
`hsv2rgb`, `rgb2hsv`, `cosinePalette`, `contrast`, `sepia`, `blendMultiply`,
`blendScreen`, `blendAdd`

### Noise (gui-shader-noise)

`hash`, `hash2`, `hash3`, `noise2D`, `fbm`

### SDF (gui-shader-sdf)

**Primitives:** `sdCircle`, `sdBox`, `sdLine`, `sdRoundedBox`

**Fill/Stroke:** `sdfFill`, `sdfSmoothFill`, `sdfStroke`, `sdfGlow`

**Boolean Ops:** `sdUnion`, `sdIntersect`, `sdSubtract`, `sdSmoothUnion`,
`sdSmoothIntersect`, `sdSmoothSubtract`, `sdAnnular`

**Repetition:** `sdRepeat2`

**Patterns:** `checkerboard`, `stripes`, `grid`, `dots`, `voronoi`

### Codegen (gui-shader-codegen)

**GLSL:** `glslVersion`, `glslPrecision`, `glslStdUniforms`, `glslUniform`,
`glslUniforms`, `glslIn`, `glslOut`, `glslQuadVertex`, `glslQuadVertexCompat`,
`glslFragmentWrap`, `glslMathLib`, `assembleGLSL`

**HLSL:** `hlslStdCBuffer`, `hlslCBuffer`, `hlslQuadVertex`,
`hlslFragmentWrap`, `hlslMathLib`, `assembleHLSL`

**WebGL:** `submitWebGL`, `drawWebGL`, `renderWebGL`

**Compilation:** `compileGLSL`, `compileHLSL`, `compileDXC`
