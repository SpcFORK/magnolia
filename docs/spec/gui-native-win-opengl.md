# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-native-win-opengl.oak`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `PFD_DRAW_TO_WINDOW` · `4`
- `PFD_SUPPORT_OPENGL` · `32`
- `PFD_DOUBLEBUFFER` · `1`
- `GL_COLOR_BUFFER_BIT` · `16384`
- `GL_BGRA_EXT` · `32993`
- `GL_UNSIGNED_BYTE` · `5121`
- `GL_UNPACK_ALIGNMENT` · `3317`
- `GL_RGBA8` · `32856`
- `GL_RENDERBUFFER` · `36161`
- `GL_READ_FRAMEBUFFER` · `36008`
- `GL_DRAW_FRAMEBUFFER` · `36009`
- `GL_COLOR_ATTACHMENT0` · `36064`
- `GL_NEAREST` · `9728`
- `GL_LINEAR` · `9729`
- `DIB_RGB_COLORS` · `0`
- `BI_RGB` · `0`
- `_glProcs` · `{}`
### `_glGetProc(name)`

### `_glCall(proc, args...)`

### `_initOpenGLWithContext(window, hdcHandle, chosen)`

> returns `:object`

### `initOpenGL2DLayer(window)`

### `_ensureGlPixelBuffer(window)`

### `_ensureGlBmiHeader(window)`

### `_ensureGlFbo(window, dw, dh)`

> returns `:bool`

### `_presentGlScaled(window, hdc, ctx, r)`

> returns `:object`

### `presentFrameOpenGL(window)`

> returns `:object`

### `releaseGlPixelBuffer(window)`

> returns `:int`

