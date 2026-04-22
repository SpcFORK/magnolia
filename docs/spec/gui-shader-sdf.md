# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-color`

### `_clampByte(value)`

### `_clampOpacity(value)`

### `rgb(r, g, b)`

> returns `:bool`

### `colorR(color)`

### `colorG(color)`

### `colorB(color)`

### `opacity(color, amount, background)`

### `rgba(r, g, b, a, background)`

## Module: `gui-shader-color`

- `guiColor` · `import(...)`
- `m` · `import(...)`
- `threadLib` · `import(...)`
### `packRGB(r, g, b)`

### `unpackRGB(c)`

> returns `:object`

### `colorR(c)`

> returns `:bool`

### `colorG(c)`

> returns `:bool`

### `colorB(c)`

> returns `:bool`

### `mix(a, b, t)`

### `mix3(a, b, c, t)`

### `brighten(c, amount)`

### `darken(c, amount)`

### `invert(c)`

### `grayscale(c)`

### `overlay(fg, bg, alpha)`

### `hsl2rgb(h, s, l)`

### `rgb2hsl(c)`

> returns `:object`

### `hsv2rgb(h, s, v)`

### `rgb2hsv(c)`

> returns `:object`

### `floatStr(v)`

### `cosinePalette(t, a, b, c, d)`

### `contrast(c, amount)`

### `sepia(c)`

### `blendMultiply(a, b)`

### `blendScreen(a, b)`

### `blendAdd(a, b)`

## Module: `gui-shader-math`

- `threadLib` · `import(...)`
- `PI` · `3.14159265358979`
- `TAU` · `6.28318530717959`
- `HALF_PI` · `1.5707963267949`
- `E` · `2.71828182845905`
- `DEG2RAD` — constant
- `RAD2DEG` — constant
- `SQRT2` · `1.4142135623731`
### `fract(x)`

### `mod(x, y)`

### `sign(x)`

> returns `:int`

### `abs2(x)`

> returns `:int`

### `clamp(x, lo, hi)`

### `saturate(x)`

### `lerpFloat(a, b, t)`

### `inverseLerp(a, b, x)`

> returns `:float`

### `remap(x, inLo, inHi, outLo, outHi)`

### `step(edge, x)`

> returns `:int`

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

> returns `:float`

### `easeInCubic(t)`

### `easeOutCubic(t)`

### `easeInOutCubic(t)`

> returns `:float`

### `easeInSine(t)`

> returns `:float`

### `easeOutSine(t)`

### `easeInOutSine(t)`

> returns `:int`

### `easeInExpo(t)`

> returns `:float`

### `easeOutExpo(t)`

> returns `:float`

### `easeOutElastic(t)`

> returns `:float`

### `easeOutBounce(t)`

> returns `:float`

### `vec2(x, y)`

> returns `:object`

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

> returns `:object`

### `fromPolar(r, theta)`

### `vec3(x, y, z)`

> returns `:object`

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

## Module: `gui-shader-noise`

- `m` · `import(...)`
- `threadLib` · `import(...)`
### `hash(seed)`

### `hash2(a, b)`

### `hash3(a, b, c)`

### `noise2D(x, y)`

### `fbm(x, y, octaves?)`

### `noiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `fbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

## Module: `lib\gui-shader-sdf.oak`

- `m` · `import(...)`
- `col` · `import(...)`
- `noise` · `import(...)`
- `threadLib` · `import(...)`
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

> returns `:int`

### `dots(x, y, spacing, radius)`

> returns `:int`

### `voronoi(x, y, scale_)`

> returns `:object`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

