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

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `gui-draw`

- `windows` · `import(...)`
- `linux` · `import(...)`
- `guiFonts` · `import(...)`
- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_asBool(value)`

> returns `:bool`

### `_fontKey(fontSpec)`

### `_windowsDeleteCachedFont(window)`

> returns `?`

### `_ensureWindowsFont(window)`

> returns `:int`

### `_ensureLinuxFont(window, gcHandle)`

> returns `:int`

### `_webFontString(fontSpec)`

### `_windowsDeleteCachedPens(window)`

> returns `:int`

### `_windowsDeleteCachedBrushes(window)`

> returns `:int`

### `releaseResources(window)`

> returns `:int`

### `invalidateDrawCaches(window)`

> returns `:int`

### `drawText(window, x, y, text, color, defaultColor)`

### `textWidth(window, text)`

### `_estimateTextWidth(window, text)`

### `fillRect(window, x, y, width, height, color, defaultColor, borderColor)`

- `_maxCacheSize` · `256`
### `_evictPenCache(window)`

> returns `?`

### `_evictBrushCache(window)`

> returns `?`

### `_getCachedPen(window, hdcValue, useColor)`

### `_getCachedBrush(window, useColor)`

- `_nullPenHandle` · `0`
### `_getNullPen()`

### `pushMask(window, x, y, w, h)`

### `popMask(window)`

### `drawLine(window, x1, y1, x2, y2, color, defaultColor)`

### `_setupFillGDI(window, hdcValue, fillColor, borderColor)`

> returns `:bool`

### `fillEllipse(window, cx, cy, rx, ry, fillColor, borderColor)`

### `fillRoundedRect(window, x, y, width, height, radius, fillColor, borderColor)`

### `_writeI32(address, value)`

### `_ensurePointBuf(window, n)`

### `fillPolygon(window, pts, fillColor, borderColor)`

### `drawPolylineNative(window, pts, color)`

### `drawLinesBatch(window, segs, color, defaultColor)`

### `DrawQueue()`

> returns `:object`

### `parallelFillRects(window, coords, computeFn, numWorkers)`

## Module: `gui-fonts`

- `guiThread` · `import(...)`
### `_asBool(v)`

> returns `:bool`

- `_platformCache` — constant
### `_platformId()`

### `isWindows?()`

### `isLinux?()`

- `_winMod` · `?`
- `_linuxMod` · `?`
### `_win()`

### `_lnx()`

- `FW_THIN` · `100`
- `FW_EXTRALIGHT` · `200`
- `FW_LIGHT` · `300`
- `FW_NORMAL` · `400`
- `FW_MEDIUM` · `500`
- `FW_SEMIBOLD` · `600`
- `FW_BOLD` · `700`
- `FW_EXTRABOLD` · `800`
- `FW_HEAVY` · `900`
### `defaultFontSpec()`

> returns `:object`

### `createFont(spec)`

### `deleteFont(fontResult)`

> returns `:int`

### `_webFontString(spec)`

> returns `?`

### `fontKey(spec)`

> returns `?`

### `cachedFont(windowOrDisplay, spec)`

### `releaseCachedFonts(windowOrDisplay)`

### `measureText(windowOrDisplay, spec, text)`

### `selectFont(args)`

### `getTextMetrics(hdc)`

### `fontLineHeight(hdcOrFontStruct)`

### `buildXLFD(spec)`

### `webFontString(spec)`

## Module: `gui-shader`

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

## Module: `gui-shader-codegen`

- `threadLib` · `import(...)`
### `glslVersion(ver?)`

> returns `:string`

### `glslPrecision(prec?, type?)`

> returns `:string`

### `glslStdUniforms()`

> returns `:string`

### `glslUniform(type, name)`

> returns `:string`

### `glslUniforms(uniforms)`

### `glslIn(type, name)`

> returns `:string`

### `glslOut(type, name)`

> returns `:string`

### `glslQuadVertex()`

### `glslQuadVertexCompat()`

> returns `:string`

### `glslFragmentWrap(body, version?)`

### `glslMathLib()`

> returns `:string`

### `hlslStdCBuffer()`

> returns `:string`

### `hlslCBuffer(name, uniforms)`

### `hlslQuadVertex()`

> returns `:string`

### `hlslFragmentWrap(body)`

### `hlslMathLib()`

> returns `:string`

### `submitWebGL(window, fragSource, vertSource?)`

> returns `:object`

### `drawWebGL(window, clearR?, clearG?, clearB?)`

### `renderWebGL(window, fragSource)`

### `compileGLSL(source, stage?, outputPath?)`

### `compileHLSL(source, profile?, entry?, outputPath?)`

### `compileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `assembleGLSL(opts)`

### `assembleHLSL(opts)`

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

## Module: `gui-shader-sdf`

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

## Module: `gui-thread`

- `threadLib` · `import(...)`
### `CommandQueue()`

> returns `:object`

### `FrameFence(workerCount)`

> returns `:object`

### `WorkerPool(numWorkers)`

> returns `:object`

### `StateGuard()`

> returns `:object`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

### `AsyncLoader(cmdQueue)`

> returns `:object`

### `FrameScheduler(pool, cmdQueue)`

> returns `:object`

### `initWindowThreading(window, options)`

### `threadingEnabled?(window)`

### `commandQueue(window)`

### `workerPool(window)`

### `scheduler(window)`

### `stateGuard(window)`

### `flushThreadedCommands(window)`

### `destroyWindowThreading(window)`

> returns `?`

## Module: `lib\gui-video.oak`

- `video` · `import(...)`
- `guiColor` · `import(...)`
- `guiShader` · `import(...)`
- `threadLib` · `import(...)`
### `frameToBuffer(frame)`

### `bufferToFrame(buf)`

### `renderFrame(drawCtx, window, frame, ox, oy)`

### `renderFrameScaled(drawCtx, window, frame, ox, oy, scale)`

### `bufferGrayscale(buf)`

### `bufferInvert(buf)`

### `bufferThreshold(buf, t)`

### `bufferBlend(bufA, bufB, alpha)`

### `bufferDiff(bufA, bufB)`

### `captureBuffer(buf)`

### `frameToBmpPixels(frame)`

## Module: `linux`

- `sys` · `import(...)`
## Module: `linux-constants`

- `LibC` · `[4 items]`
- `LibDL` · `[4 items]`
- `LibX11` · `[4 items]`
- `PROT_NONE` · `0`
- `PROT_READ` · `1`
- `PROT_WRITE` · `2`
- `PROT_EXEC` · `4`
- `MAP_SHARED` · `1`
- `MAP_PRIVATE` · `2`
- `MAP_FIXED` · `16`
- `MAP_ANONYMOUS` · `32`
- `O_RDONLY` · `0`
- `O_WRONLY` · `1`
- `O_RDWR` · `2`
- `O_CREAT` · `64`
- `O_TRUNC` · `512`
- `O_APPEND` · `1024`
- `SEEK_SET` · `0`
- `SEEK_CUR` · `1`
- `SEEK_END` · `2`
- `F_OK` · `0`
- `X_OK` · `1`
- `W_OK` · `2`
- `R_OK` · `4`
- `RTLD_LAZY` · `1`
- `RTLD_NOW` · `2`
- `RTLD_GLOBAL` · `256`
- `RTLD_LOCAL` · `0`
- `_SC_PAGESIZE` · `30`
- `KeyPressMask` · `1`
- `ButtonPressMask` · `4`
- `ExposureMask` · `32768`
- `StructureNotifyMask` · `131072`
- `KeyPress` · `2`
- `Expose` · `12`
- `DestroyNotify` · `17`
- `ClientMessage` · `33`
## Module: `linux-core`

- `sys` · `import(...)`
### `isLinux?()`

### `cstr(s)`

### `_readCString(ptr, maxLen)`

### `_platformError(apiName)`

> returns `:object`

### `_resolveFirst(libraries, symbol, i)`

> returns `:object`

### `_callResolved(resolved, args...)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `libc(symbol, args...)`

### `libdl(symbol, args...)`

## Module: `linux-libc`

- `sys` · `import(...)`
### `currentProcessId()`

### `parentProcessId()`

### `pageSize()`

### `errno()`

### `strerror(errorCode)`

### `lastErrorMessage()`

### `getLastError()`

### `formatMessage(errorCode)`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `getuid()`

### `geteuid()`

### `getgid()`

### `getegid()`

### `gethostname(bufferPtr, size)`

### `getcwd(bufferPtr, size)`

### `chdir(path)`

### `access(path, mode)`

### `openFile(path, flags, mode)`

### `closeFile(fd)`

### `closeHandle(handle)`

### `readFileDescriptor(fd, bufferPtr, count)`

### `writeFileDescriptor(fd, bufferPtr, count)`

### `seek(fd, offset, whence)`

### `unlink(path)`

### `mmap(addr, length, prot, flags, fd, offset)`

### `munmap(addr, length)`

### `mprotect(addr, length, prot)`

### `allocPages(size, prot)`

### `freePages(address, size)`

### `protectPages(address, size, prot)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `_compatNotImplemented(apiName)`

> returns `:object`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `linux-loader`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_libraryKey(library)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `dlopen(path, flags)`

### `dlsym(handle, symbol)`

### `dlclose(handle)`

### `loadLibrary(path)`

### `freeLibrary(handle)`

### `procAddress(module, symbol)`

### `_loadDllCandidate(candidates, i, flags)`

> returns `:object`

### `loadDll(library)`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `x11(symbol, args...)`

## Module: `linux-windowing`

- `sys` · `import(...)`
### `_r1Positive?(res)`

> returns `:bool`

### `callOk?(res)`

> returns `:bool`

### `_zeros(n, acc)`

### `xEventSize()`

> returns `:int`

### `createXEventBuffer()`

### `xEventType(eventPtr)`

### `openDisplay(displayName)`

### `closeDisplay(display)`

### `defaultScreen(display)`

### `rootWindow(display, screen)`

### `blackPixel(display, screen)`

### `whitePixel(display, screen)`

### `createSimpleWindow(display, parent, x, y, width, height, borderWidth, border, background)`

### `destroyWindow(display, window)`

### `storeName(display, window, title)`

### `selectInput(display, window, eventMask)`

### `mapWindow(display, window)`

### `unmapWindow(display, window)`

### `moveWindow(display, window, x, y)`

### `resizeWindow(display, window, width, height)`

### `displayWidth(display, screen)`

### `displayHeight(display, screen)`

### `createGC(display, drawable, valueMask, values)`

### `freeGC(display, gc)`

### `setForeground(display, gc, color)`

### `drawLine(display, window, gc, x1, y1, x2, y2)`

### `fillRectangle(display, window, gc, x, y, width, height)`

### `drawString(display, window, gc, x, y, text)`

### `flush(display)`

### `pending(display)`

### `nextEvent(display, eventPtr)`

### `_openWindowFromRoot(display, screen, root, black, white, title, width, height)`

> returns `:object`

### `_openWindowFromScreen(display, screen, title, width, height)`

### `_openWindowFromDisplay(display, title, width, height)`

### `openDefaultWindow(title, width, height)`

### `closeWindow(state)`

> returns `:int`

### `pumpWindowEvent(display, eventPtr)`

> returns `:object`

### `runWindowLoop(display, eventPtr)`

> returns `:int`

## Module: `math`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

### `orient(x0, y0, x1, y1)`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

### `median(xs)`

### `stddev(xs)`

### `round(n, decimals)`

## Module: `math-base`

- `Pi` · `3.141592653589793`
- `E` · `2.718281828459045`
### `sign(n)`

> returns `:int`

### `abs(n)`

### `sqrt(n)`

## Module: `math-geo`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

> returns `:list`

### `orient(x0, y0, x1, y1)`

> returns `:int`

## Module: `math-stats`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

> returns `?`

### `median(xs)`

> returns `?`

### `stddev(xs)`

### `pbatchMean(datasets)`

### `pbatchStddev(datasets)`

### `round(n, decimals)`

## Module: `sort`

### `sort!(xs, pred)`

### `sort(xs, pred)`

### `_mergeSorted(a, b, pred)`

### `psort(xs, pred)`

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

## Module: `sys`

### `_isObject?(v)`

### `ok?(result)`

> returns `:bool`

### `error?(result)`

> returns `:bool`

### `resolve(library, symbol)`

> returns `:object`

### `call(target, args...)`

### `resolveAndCall(library, symbol, args...)`

### `valueOr(result, fallback)`

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

## Module: `video`

### `clampByte(n)`

> returns `:int`

### `clampUnit(n)`

> returns `:int`

### `_min(a, b)`

### `_max(a, b)`

### `_pixelIndexRaw(width, channels, x, y, ch)`

### `pixelIndex(frame, x, y, ch)`

### `frame(width, height, pixels, channels)`

> returns `:object`

### `blank(width, height, channels, value)`

### `cloneFrame(f)`

> returns `:object`

### `getPixel(f, x, y)`

### `setPixel(f, x, y, pixel)`

> returns `:object`

### `mapPixels(f, mapper)`

> returns `:object`

### `_luma(pixel)`

### `toGrayscale(f)`

### `invert(f, maxValue)`

### `threshold(f, t)`

### `rgbToYuv(pixel)`

> returns `:list`

### `yuvToRgb(pixel)`

> returns `:list`

### `crop(f, x, y, width, height)`

### `resizeNearest(f, newWidth, newHeight)`

### `blend(a, b, alpha)`

### `frameDiff(a, b)`

### `mapFrames(frames, f)`

### `sampleFrame(frames, timeSeconds, fps)`

> returns `?`

### `pmapPixels(f, mapper, numWorkers)`

> returns `:object`

### `presizeNearest(f, newWidth, newHeight, numWorkers)`

### `pmapFrames(frames, f)`

## Module: `windows`

- `sys` · `import(...)`
## Module: `windows-constants`

- `Kernel32` · `'kernel32.dll'`
- `Ntdll` · `'ntdll.dll'`
- `Psapi` · `'psapi.dll'`
- `User32` · `'user32.dll'`
- `Gdi32` · `'gdi32.dll'`
- `Advapi32` · `'advapi32.dll'`
- `Shell32` · `'shell32.dll'`
- `Ole32` · `'ole32.dll'`
- `Ws2_32` · `'ws2_32.dll'`
- `Comctl32` · `'comctl32.dll'`
- `Wininet` · `'wininet.dll'`
- `OpenGL32` · `'opengl32.dll'`
- `Vulkan1` · `'vulkan-1.dll'`
- `D3d9` · `'d3d9.dll'`
- `D3d11` · `'d3d11.dll'`
- `Dxgi` · `'dxgi.dll'`
- `Ddraw` · `'ddraw.dll'`
- `Msvcrt` · `'msvcrt.dll'`
- `Ucrtbase` · `'ucrtbase.dll'`
- `Vcruntime140` · `'vcruntime140.dll'`
- `ActionCenter` · `'actioncenter.dll'`
- `Aclui` · `'aclui.dll'`
- `Acledit` · `'acledit.dll'`
- `Acppage` · `'acppage.dll'`
- `Acproxy` · `'acproxy.dll'`
- `Adprovider` · `'adprovider.dll'`
- `Aeinv` · `'aeinv.dll'`
- `Aepic` · `'aepic.dll'`
- `Amstream` · `'amstream.dll'`
- `Adsldp` · `'adsldp.dll'`
- `Adsnt` · `'adsnt.dll'`
- `Adtschema` · `'adtschema.dll'`
- `Adsldpc` · `'adsldpc.dll'`
- `Adsmsext` · `'adsmsext.dll'`
- `Adhsvc` · `'adhsvc.dll'`
- `Advapi32res` · `'advapi32res.dll'`
- `Advpack` · `'advpack.dll'`
- `Aeevts` · `'aeevts.dll'`
- `Apds` · `'apds.dll'`
- `Winhttp` · `'winhttp.dll'`
- `Urlmon` · `'urlmon.dll'`
- `Crypt32` · `'crypt32.dll'`
- `Bcrypt` · `'bcrypt.dll'`
- `Secur32` · `'secur32.dll'`
- `Comdlg32` · `'comdlg32.dll'`
- `Oleaut32` · `'oleaut32.dll'`
- `Imm32` · `'imm32.dll'`
- `Shlwapi` · `'shlwapi.dll'`
- `Shcore` · `'shcore.dll'`
- `UxTheme` · `'uxtheme.dll'`
- `Dwmapi` · `'dwmapi.dll'`
- `Version` · `'version.dll'`
- `Setupapi` · `'setupapi.dll'`
- `Netapi32` · `'netapi32.dll'`
- `Winmm` · `'winmm.dll'`
- `Avrt` · `'avrt.dll'`
- `Mmdevapi` · `'mmdevapi.dll'`
- `Dsound` · `'dsound.dll'`
- `Mfplat` · `'mfplat.dll'`
- `Mfreadwrite` · `'mfreadwrite.dll'`
- `Mfuuid` · `'mfuuid.dll'`
- `Taskschd` · `'taskschd.dll'`
- `Wevtapi` · `'wevtapi.dll'`
- `Wlanapi` · `'wlanapi.dll'`
- `Mpr` · `'mpr.dll'`
- `Spoolss` · `'spoolss.dll'`
- `Wtsapi32` · `'wtsapi32.dll'`
- `Rasapi32` · `'rasapi32.dll'`
- `Msi` · `'msi.dll'`
- `Wimgapi` · `'wimgapi.dll'`
- `Cabinet` · `'cabinet.dll'`
- `Apphelp` · `'apphelp.dll'`
- `Wer` · `'wer.dll'`
- `Faultrep` · `'faultrep.dll'`
- `Dbghelp` · `'dbghelp.dll'`
- `Dbgeng` · `'dbgeng.dll'`
- `Pdh` · `'pdh.dll'`
- `Iphlpapi` · `'iphlpapi.dll'`
- `Wscapi` · `'wscapi.dll'`
- `Sensapi` · `'sensapi.dll'`
- `Ncrypt` · `'ncrypt.dll'`
- `Cryptui` · `'cryptui.dll'`
- `Wintrust` · `'wintrust.dll'`
- `Samlib` · `'samlib.dll'`
- `Netshell` · `'netshell.dll'`
- `Fwpuclnt` · `'fwpuclnt.dll'`
- `Dnsapi` · `'dnsapi.dll'`
- `Nlaapi` · `'nlaapi.dll'`
- `Httpapi` · `'httpapi.dll'`
- `Rpcrt4` · `'rpcrt4.dll'`
- `Srpapi` · `'srpapi.dll'`
- `Sxs` · `'sxs.dll'`
- `Msvcirt` · `'msvcirt.dll'`
- `ApiSetPrefix` · `'api-ms-win-'`
- `D3dx9Prefix` · `'d3dx9_'`
- `MsvcpPrefix` · `'msvcp'`
- `VcruntimePrefix` · `'vcruntime'`
- `AtlPrefix` · `'atl'`
- `MfcPrefix` · `'mfc'`
- `VcompPrefix` · `'vcomp'`
## Module: `windows-core`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `makeWord(low, high)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `kernel32(symbol, args...)`

### `ntdll(symbol, args...)`

### `ntNative(symbol, args...)`

### `psapi(symbol, args...)`

## Module: `windows-flags`

- `PROCESS_TERMINATE` · `1`
- `PROCESS_VM_READ` · `16`
- `PROCESS_VM_WRITE` · `32`
- `PROCESS_VM_OPERATION` · `8`
- `PROCESS_QUERY_INFORMATION` · `1024`
- `PROCESS_QUERY_LIMITED_INFORMATION` · `4096`
- `PROCESS_ALL_ACCESS` · `2035711`
- `MEM_COMMIT` · `4096`
- `MEM_RESERVE` · `8192`
- `MEM_DECOMMIT` · `16384`
- `MEM_RELEASE` · `32768`
- `PAGE_NOACCESS` · `1`
- `PAGE_READONLY` · `2`
- `PAGE_READWRITE` · `4`
- `PAGE_EXECUTE` · `16`
- `PAGE_EXECUTE_READ` · `32`
- `PAGE_EXECUTE_READWRITE` · `64`
- `FORMAT_MESSAGE_IGNORE_INSERTS` · `512`
- `FORMAT_MESSAGE_FROM_SYSTEM` · `4096`
- `ERROR_SUCCESS` · `0`
- `AF_INET` · `2`
- `SOCK_STREAM` · `1`
- `SOCK_DGRAM` · `2`
- `IPPROTO_TCP` · `6`
- `IPPROTO_UDP` · `17`
- `SOCKET_ERROR` — constant
- `INVALID_SOCKET` — constant
- `SD_RECEIVE` · `0`
- `SD_SEND` · `1`
- `SD_BOTH` · `2`
- `INTERNET_OPEN_TYPE_PRECONFIG` · `0`
- `INTERNET_OPEN_TYPE_DIRECT` · `1`
- `INTERNET_OPEN_TYPE_PROXY` · `3`
- `INTERNET_DEFAULT_HTTP_PORT` · `80`
- `INTERNET_DEFAULT_HTTPS_PORT` · `443`
- `INTERNET_SERVICE_HTTP` · `3`
- `HKEY_CLASSES_ROOT` · `2147483648`
- `HKEY_CURRENT_USER` · `2147483649`
- `HKEY_LOCAL_MACHINE` · `2147483650`
- `HKEY_USERS` · `2147483651`
- `HKEY_CURRENT_CONFIG` · `2147483653`
- `KEY_QUERY_VALUE` · `1`
- `KEY_SET_VALUE` · `2`
- `KEY_CREATE_SUB_KEY` · `4`
- `KEY_ENUMERATE_SUB_KEYS` · `8`
- `KEY_READ` · `131097`
- `KEY_WRITE` · `131078`
- `REG_SZ` · `1`
- `REG_DWORD` · `4`
- `REG_QWORD` · `11`
- `CS_VREDRAW` · `1`
- `CS_HREDRAW` · `2`
- `CS_DBLCLKS` · `8`
- `CS_OWNDC` · `32`
- `WS_OVERLAPPED` · `0`
- `WS_CAPTION` · `12582912`
- `WS_SYSMENU` · `524288`
- `WS_THICKFRAME` · `262144`
- `WS_MINIMIZEBOX` · `131072`
- `WS_MAXIMIZEBOX` · `65536`
- `WS_VISIBLE` · `268435456`
- `WS_CLIPSIBLINGS` · `67108864`
- `WS_CLIPCHILDREN` · `33554432`
- `WS_OVERLAPPEDWINDOW` · `13565952`
- `CW_USEDEFAULT` — constant
- `WS_POPUP` · `2147483648`
- `WS_EX_APPWINDOW` · `262144`
- `GWL_STYLE` — constant
- `GWL_EXSTYLE` — constant
- `SM_CXSCREEN` · `0`
- `SM_CYSCREEN` · `1`
- `HWND_TOP` · `0`
- `HWND_TOPMOST` — constant
- `HWND_NOTOPMOST` — constant
- `WM_CREATE` · `1`
- `WM_DESTROY` · `2`
- `WM_PAINT` · `15`
- `WM_CLOSE` · `16`
- `WM_QUIT` · `18`
- `WM_COMMAND` · `273`
- `SW_HIDE` · `0`
- `SW_MAXIMIZE` · `3`
- `SW_SHOW` · `5`
- `SW_RESTORE` · `9`
- `PM_NOREMOVE` · `0`
- `PM_REMOVE` · `1`
- `MB_OK` · `0`
- `MB_ICONERROR` · `16`
- `MB_ICONWARNING` · `48`
- `MB_ICONINFORMATION` · `64`
- `IDC_ARROW` · `32512`
- `IDI_APPLICATION` · `32512`
## Module: `windows-gdi`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
- `_szBuf` · `bits(...)`
### `isWindows?()`

### `wstr(s)`

- `_gdiProcCache` · `{}`
- `_userProcCache` · `{}`
### `_cachedGdi32(symbol, args...)`

### `_cachedUser32(symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `beginPaint(hwnd, paintStructPtr)`

### `endPaint(hwnd, paintStructPtr)`

### `getDC(hwnd)`

### `releaseDC(hwnd, hdc)`

### `getStockObject(objectIndex)`

### `selectObject(hdc, gdiObject)`

### `setBkMode(hdc, mode)`

### `setTextColor(hdc, colorRef)`

### `textOut(hdc, x, y, text)`

### `createFont(height, width, escapement, orientation, weight, italic, underline, strikeOut, charSet, outPrecision, clipPrecision, quality, pitchAndFamily, faceName)`

### `rectangle(hdc, left, top, right, bottom)`

### `ellipse(hdc, left, top, right, bottom)`

### `createSolidBrush(colorRef)`

### `getTextExtentPoint32(hdc, text)`

> returns `:object`

### `deleteObject(gdiObject)`

## Module: `windows-kernel`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `_utf16leToString(buf)`

### `wstr(s)`

### `cstr(s)`

### `kernel32(symbol, args...)`

### `statusOk?(res)`

> returns `:bool`

### `ptrSize()`

> returns `:int`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `getLastError()`

### `formatMessage(errorCode)`

### `lastErrorMessage()`

### `currentProcessId()`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `loadLibrary(path)`

### `freeLibrary(module)`

### `procAddress(module, symbol)`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `closeHandle(handle)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `windows-loader`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

- `_loaderProcCache` · `{}`
### `_cachedCallIn(library, symbol, args...)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `loadDll(library)`

> returns `:object`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `advapi32(symbol, args...)`

### `shell32(symbol, args...)`

### `ole32(symbol, args...)`

### `ws2_32(symbol, args...)`

### `comctl32(symbol, args...)`

### `wininet(symbol, args...)`

### `opengl32(symbol, args...)`

### `vulkan1(symbol, args...)`

### `d3d9(symbol, args...)`

### `ddraw(symbol, args...)`

### `d3d11(symbol, args...)`

### `dxgi(symbol, args...)`

### `directDrawCreateEx(guidPtr, outPtr, iidPtr, outerUnknown)`

### `directDrawCreate(guidPtr, outPtr, outerUnknown)`

### `direct3DCreate9(sdkVersion)`

### `d3dx9Dll(suffix)`

### `d3dx9(suffix, symbol, args...)`

### `apiSetDll(contract)`

### `apiSet(contract, symbol, args...)`

### `msvcrt(symbol, args...)`

### `ucrtbase(symbol, args...)`

### `vcruntime140(symbol, args...)`

### `actionCenter(symbol, args...)`

### `aclui(symbol, args...)`

### `acledit(symbol, args...)`

### `acppage(symbol, args...)`

### `acproxy(symbol, args...)`

### `adprovider(symbol, args...)`

### `aeinv(symbol, args...)`

### `aepic(symbol, args...)`

### `amstream(symbol, args...)`

### `adsldp(symbol, args...)`

### `adsnt(symbol, args...)`

### `adtschema(symbol, args...)`

### `adsldpc(symbol, args...)`

### `adsmsext(symbol, args...)`

### `adhsvc(symbol, args...)`

### `advapi32res(symbol, args...)`

### `advpack(symbol, args...)`

### `aeevts(symbol, args...)`

### `apds(symbol, args...)`

### `winhttp(symbol, args...)`

### `urlmon(symbol, args...)`

### `crypt32(symbol, args...)`

### `bcrypt(symbol, args...)`

### `secur32(symbol, args...)`

### `comdlg32(symbol, args...)`

### `oleaut32(symbol, args...)`

### `imm32(symbol, args...)`

### `shlwapi(symbol, args...)`

### `shcore(symbol, args...)`

### `uxTheme(symbol, args...)`

### `dwmapi(symbol, args...)`

### `versionDll(symbol, args...)`

### `setupapi(symbol, args...)`

### `netapi32(symbol, args...)`

### `winmm(symbol, args...)`

### `avrt(symbol, args...)`

### `mmdevapi(symbol, args...)`

### `dsound(symbol, args...)`

### `mfplat(symbol, args...)`

### `mfreadwrite(symbol, args...)`

### `mfuuid(symbol, args...)`

### `taskschd(symbol, args...)`

### `wevtapi(symbol, args...)`

### `wlanapi(symbol, args...)`

### `mpr(symbol, args...)`

### `spoolss(symbol, args...)`

### `wtsapi32(symbol, args...)`

### `rasapi32(symbol, args...)`

### `msi(symbol, args...)`

### `wimgapi(symbol, args...)`

### `cabinet(symbol, args...)`

### `apphelp(symbol, args...)`

### `wer(symbol, args...)`

### `faultrep(symbol, args...)`

### `dbghelp(symbol, args...)`

### `dbgeng(symbol, args...)`

### `pdh(symbol, args...)`

### `iphlpapi(symbol, args...)`

### `wscapi(symbol, args...)`

### `sensapi(symbol, args...)`

### `ncrypt(symbol, args...)`

### `cryptui(symbol, args...)`

### `wintrust(symbol, args...)`

### `samlib(symbol, args...)`

### `netshell(symbol, args...)`

### `fwpuclnt(symbol, args...)`

### `dnsapi(symbol, args...)`

### `nlaapi(symbol, args...)`

### `httpapi(symbol, args...)`

### `rpcrt4(symbol, args...)`

### `srpapi(symbol, args...)`

### `sxs(symbol, args...)`

### `msvcirt(symbol, args...)`

### `_familyDll(prefix, suffix)`

### `msvcpDll(suffix)`

### `msvcpFamily(suffix, symbol, args...)`

### `vcruntimeDll(suffix)`

### `vcruntimeFamily(suffix, symbol, args...)`

### `atlDll(suffix)`

### `atlFamily(suffix, symbol, args...)`

### `mfcDll(suffix)`

### `mfcFamily(suffix, symbol, args...)`

### `vcompDll(suffix)`

### `vcompFamily(suffix, symbol, args...)`

## Module: `windows-net`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `callValueOrZero(res)`

### `ws2_32(symbol, args...)`

### `wininet(symbol, args...)`

### `wsLastError()`

### `wsaStartup(version, wsaDataPtr)`

### `wsaCleanup()`

### `socket(af, socketType, protocol)`

### `bindSocket(sock, sockaddrPtr, sockaddrLen)`

### `connectSocket(sock, sockaddrPtr, sockaddrLen)`

### `listenSocket(sock, backlog)`

### `acceptSocket(sock, addrOutPtr, addrLenInOutPtr)`

### `sendSocket(sock, bufferPtr, size, flags)`

### `recvSocket(sock, bufferPtr, size, flags)`

### `shutdownSocket(sock, how)`

### `closeSocket(sock)`

### `htons(value)`

### `htonl(value)`

### `inetAddr(ipv4)`

### `internetOpen(agent, accessType, proxy, proxyBypass, flags)`

### `internetConnect(hInternet, serverName, serverPort, username, password, service, flags, context)`

### `internetOpenUrl(hInternet, url, headers, headersLen, flags, context)`

### `internetReadFile(hFile, outBufferPtr, bytesToRead, bytesReadOutPtr)`

### `internetCloseHandle(hInternet)`

### `_bytesToString(raw)`

### `sockaddrIn(ipv4, port)`

> returns `:object`

### `_internetReadAll(hInternetFile, chunkBuf, bytesReadBuf, chunkSize, out)`

> returns `:object`

### `internetSimpleGet(url, agent, chunkSize)`

## Module: `windows-registry`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `ptrSize()`

> returns `:int`

### `_zeros(n, acc)`

### `_statusOk?(res)`

> returns `:bool`

### `_ptrRead(address)`

### `_utf16leToString(buf)`

### `wstr(s)`

### `advapi32(symbol, args...)`

### `regCloseKey(hKey)`

### `regOpenKeyEx(rootKey, subKey, options, samDesired, outKeyPtr)`

### `regCreateKeyEx(rootKey, subKey, reserved, className, options, samDesired, securityAttributesPtr, outKeyPtr, dispositionOutPtr)`

### `regQueryValueEx(hKey, valueName, reserved, typeOutPtr, dataOutPtr, dataLenInOutPtr)`

### `regSetValueEx(hKey, valueName, reserved, valueType, dataPtr, dataLen)`

### `regDeleteValue(hKey, valueName)`

### `regDeleteTree(rootKey, subKey)`

### `regReadDword(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteDword(rootKey, subKey, valueName, value)`

> returns `:object`

### `regReadString(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteString(rootKey, subKey, valueName, value)`

> returns `:object`

## Module: `windows-windowing`

- `sys` · `import(...)`
### `_toI32(u)`

- `_cachedDefProcAddr` · `?`
- `_cachedCursor` · `?`
- `_cachedIcon` · `?`
- `_cachedImageBase` · `?`
- `_regClassTimings` · `{}`
- `_registeredClasses` · `{}`
- `DefaultClassName` · `'MagnoliaGUIWindowClass'`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

- `_cachedPtrSize` — constant
### `ptrSize()`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `noMessage?(res)`

> returns `:bool`

- `_k32ProcCache` · `{}`
- `_u32ProcCache` · `{}`
- `_shcoreProcCache` · `{}`
- `_dwmapiProcCache` · `{}`
### `kernel32(symbol, args...)`

### `user32(symbol, args...)`

### `shcore(symbol, args...)`

### `dwmapi(symbol, args...)`

### `moduleHandle(name)`

### `imageBase()`

### `registerClassEx(wndClassExPtr)`

### `createWindowEx(exStyle, className, windowName, style, x, y, width, height, parent, menu, instance, param)`

### `defWindowProc(hwnd, msg, wParam, lParam)`

### `showWindow(hwnd, cmdShow)`

### `updateWindow(hwnd)`

### `getWindowLongPtr(hwnd, index)`

### `setWindowLongPtr(hwnd, index, value)`

### `getSystemMetrics(idx)`

### `destroyWindow(hwnd)`

### `postQuitMessage(exitCode)`

### `getMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax)`

### `peekMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax, removeMsg)`

### `translateMessage(msgPtr)`

### `dispatchMessage(msgPtr)`

### `isWindow(hwnd)`

### `waitMessage()`

### `windowAlive?(hwnd)`

> returns `:bool`

### `msgStructSize()`

> returns `:int`

### `createMsgBuffer()`

### `pumpWindowMessage(hwnd, msgPtr)`

> returns `:object`

### `loadCursor(instance, cursorId)`

### `loadIcon(instance, iconId)`

### `registerWindowClassEx(className, iconHandle, smallIconHandle, cursorHandle, classStyle)`

- `_initTimings` · `?`
### `getRegClassTimings()`

### `getInitTimings()`

### `registerDefaultWindowClass(className)`

### `createTopLevelWindow(className, title, width, height, style)`

### `runWindowLoop(hwnd)`

### `runWindowLoopPeek(hwnd, msgPtr)`

### `messageBox(hwnd, text, caption, msgType)`

### `setWindowText(hwnd, text)`

### `getCursorPos()`

> returns `:object`

### `getWindowRect(hwnd)`

> returns `:object`

### `getWindowPlacement(hwnd)`

> returns `:object`

- `DPI_AWARENESS_CONTEXT_UNAWARE` — constant
- `DPI_AWARENESS_CONTEXT_SYSTEM_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2` — constant
- `MDT_EFFECTIVE_DPI` · `0`
- `MDT_ANGULAR_DPI` · `1`
- `MDT_RAW_DPI` · `2`
### `setProcessDpiAwarenessContext(context)`

> returns `:bool`

### `getDpiForWindow(hwnd)`

### `getDpiForSystem()`

### `getDpiForMonitor(hMonitor, dpiType)`

> returns `:object`

- `MONITOR_DEFAULTTONULL` · `0`
- `MONITOR_DEFAULTTOPRIMARY` · `1`
- `MONITOR_DEFAULTTONEAREST` · `2`
### `monitorFromWindow(hwnd, flags)`

### `adjustWindowRectExForDpi(x, y, w, h, style, exStyle, dpi)`

> returns `:object`

### `enableNonClientDpiScaling(hwnd)`

> returns `:bool`

### `dpiScale(value, dpi)`

### `dpiUnscale(value, dpi)`

- `_MONITORINFOEX_SIZE` · `104`
### `getMonitorInfo(hMonitor)`

> returns `:object`

### `monitorFromPoint(x, y, flags)`

### `monitorFromRect(left, top, right, bottom, flags)`

### `getSystemMetricsForDpi(index, dpi)`

## Module: `writes`

### `_b0(v)`

### `_b1(v)`

### `_b2(v)`

### `_b3(v)`

### `_b4(v)`

### `_b5(v)`

### `_b6(v)`

### `_b7(v)`

### `readU32(address)`

### `writeU32(address, value)`

### `readU64(address)`

### `writeU64(address, value)`

