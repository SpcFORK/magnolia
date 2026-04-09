# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-shader.oak`

- `_draw` · `import(...)`
- `threadLib` · `import(...)`
- `m` · `import(...)`
- `col` · `import(...)`
- `noise` · `import(...)`
- `sdf` · `import(...)`
- `codegen` · `import(...)`
- `PI` — constant
- `TAU` — constant
- `HALF_PI` — constant
- `E` — constant
- `DEG2RAD` — constant
- `RAD2DEG` — constant
- `SQRT2` — constant
### `fract(x)`

### `mod(x, y)`

### `sign(x)`

### `abs2(x)`

### `clamp(x, lo, hi)`

### `saturate(x)`

### `lerpFloat(a, b, t)`

### `inverseLerp(a, b, x)`

### `remap(x, inLo, inHi, outLo, outHi)`

### `step(edge, x)`

### `smoothstep(edge0, edge1, x)`

### `smootherstep(edge0, edge1, x)`

### `min2(a, b)`

### `max2(a, b)`

### `sqr(x)`

### `sqrt(x)`

### `lerp(a, b, t)`

### `atan2(y, x)`

### `pingpong(t, length)`

### `degToRad(d)`

### `radToDeg(r)`

### `easeInQuad(t)`

### `easeOutQuad(t)`

### `easeInOutQuad(t)`

### `easeInCubic(t)`

### `easeOutCubic(t)`

### `easeInOutCubic(t)`

### `easeInSine(t)`

### `easeOutSine(t)`

### `easeInOutSine(t)`

### `easeInExpo(t)`

### `easeOutExpo(t)`

### `easeOutElastic(t)`

### `easeOutBounce(t)`

### `vec2(x, y)`

### `dot2(a, b)`

### `length2(v)`

### `distance2(a, b)`

### `normalize2(v)`

### `rotate2(v, angle)`

### `scale2(v, s)`

### `add2(a, b)`

### `sub2(a, b)`

### `lerp2(a, b, t)`

### `negate2(v)`

### `abs2v(v)`

### `min2v(a, b)`

### `max2v(a, b)`

### `floor2(v)`

### `fract2(v)`

### `reflect2(v, n)`

### `toPolar(v)`

### `fromPolar(r, theta)`

### `vec3(x, y, z)`

### `add3(a, b)`

### `sub3(a, b)`

### `scale3(v, s)`

### `dot3(a, b)`

### `length3(v)`

### `distance3(a, b)`

### `normalize3(v)`

### `cross3(a, b)`

### `lerp3(a, b, t)`

### `negate3(v)`

### `reflect3(v, n)`

### `packRGB(r, g, b)`

### `unpackRGB(c)`

### `colorR(c)`

### `colorG(c)`

### `colorB(c)`

### `mix(a, b, t)`

### `mix3(a, b, c, t)`

### `brighten(c, amount)`

### `darken(c, amount)`

### `invert(c)`

### `grayscale(c)`

### `overlay(fg, bg, alpha)`

### `hsl2rgb(h, s, l)`

### `rgb2hsl(c)`

### `hsv2rgb(h, s, v)`

### `rgb2hsv(c)`

### `floatStr(c)`

### `cosinePalette(t, a, b, c, d)`

### `contrast(c, amount)`

### `sepia(c)`

### `blendMultiply(a, b)`

### `blendScreen(a, b)`

### `blendAdd(a, b)`

### `hash(seed)`

### `hash2(a, b)`

### `hash3(a, b, c)`

### `noise2D(x, y)`

### `fbm(x, y, octaves?)`

### `noiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `fbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

### `sdCircle(px, py, cx, cy, r)`

### `sdBox(px, py, cx, cy, hw, hh)`

### `sdLine(px, py, ax, ay, bx, by)`

### `sdRoundedBox(px, py, cx, cy, hw, hh, r)`

### `sdfFill(d, color)`

### `sdfSmoothFill(d, color, bg, edge)`

### `sdfStroke(d, thickness, color)`

### `sdfGlow(d, color, intensity, radius)`

### `sdUnion(d1, d2)`

### `sdIntersect(d1, d2)`

### `sdSubtract(d1, d2)`

### `sdSmoothUnion(d1, d2, k)`

### `sdSmoothIntersect(d1, d2, k)`

### `sdSmoothSubtract(d1, d2, k)`

### `sdAnnular(d, r)`

### `sdRepeat2(px, py, cx, cy)`

### `checkerboard(x, y, size)`

### `stripes(x, y, angle, width)`

### `grid(x, y, size, thickness)`

### `dots(x, y, spacing, radius)`

### `voronoi(x, y, scale_)`

### `glslVersion(ver?)`

### `glslPrecision(prec?, type?)`

### `glslStdUniforms()`

### `glslUniform(type, name)`

### `glslUniforms(uniforms)`

### `glslIn(type, name)`

### `glslOut(type, name)`

### `glslQuadVertex()`

### `glslQuadVertexCompat()`

### `glslFragmentWrap(body, version?)`

### `glslMathLib()`

### `hlslStdCBuffer()`

### `hlslCBuffer(name, uniforms)`

### `hlslQuadVertex()`

### `hlslFragmentWrap(body)`

### `hlslMathLib()`

### `submitWebGL(window, fragSource, vertSource?)`

### `drawWebGL(window, clearR?, clearG?, clearB?)`

### `renderWebGL(window, fragSource)`

### `compileGLSL(source, stage?, outputPath?)`

### `compileHLSL(source, profile?, entry?, outputPath?)`

### `compileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `assembleGLSL(opts)`

### `assembleHLSL(opts)`

- `_registry` · `[]`
### `_registerShader(shader)`

> returns `?`

### `unregisterShader(shader)`

### `clearAll()`

### `destroyAll()`

> returns `:list`

### `registeredCount()`

### `cs Shader(fragment?, opts?)`

### `elapsed(shader)`

### `pause(shader)`

### `resume(shader)`

> returns `?`

### `reset(shader)`

> returns `?`

### `setUniform(shader, key, value)`

### `getUniform(shader, key)`

### `setResolution(shader, res)`

### `beginFrame(shader)`

### `endFrame(shader)`

### `dt(shader)`

### `isRunning(shader)`

### `frameCount(shader)`

### `render(window, shader, x, y, w, h)`

### `cs ShaderPass(shader, x, y, w, h)`

### `composePasses(window, passes)`

### `renderShader(window, shader, x, y, w, h)`

### `renderShaderLines(window, shader, x, y, w, h)`

### `renderGradientBands(window, gradientFn, x, y, w, h, time, bands?)`

### `renderGradient(window, gradientFn, x, y, w, h, time)`

### `renderHorizontalBands(window, gradientFn, x, y, w, h, time, bands?)`

### `renderColumns(window, columns, ox, oy, h)`

### `updateColumns(columns, h, t, rate?)`

### `createBuffer(w, h)`

> returns `:object`

### `clearBuffer(buf, color?)`

### `setPixel(buf, x, y, color)`

> returns `?`

### `getPixel(buf, x, y)`

> returns `:int`

### `renderBuffer(window, buf, ox, oy)`

### `renderShaderToBuffer(buf, shader)`

### `renderShaderToBufferParallel(buf, shader, numWorkers)`

### `renderParallel(window, shader, x, y, w, h, numWorkers)`

