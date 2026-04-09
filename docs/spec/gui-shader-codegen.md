# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-shader-codegen.oak`

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

