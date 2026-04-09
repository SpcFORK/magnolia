# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-native-win-probe.oak`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_probeCacheReady` · `false`
- `_probeCache` · `?`
- `_probeDdrawTiming` · `0`
- `_probeD3d9Timing` · `0`
- `_probeD3d11Timing` · `0`
- `_probeOpenGLTiming` · `0`
- `_probeVulkanTiming` · `0`
### `_probeDdraw()`

### `_probeD3d9()`

### `_probeD3d11()`

### `_probeOpenGL()`

### `_probeVulkan()`

### `_pendingProbe(library, backend)`

> returns `:object`

### `_pendingDdrawProbe()`

> returns `:object`

### `pendingProbeSet()`

> returns `:object`

### `probeAllDlls()`

### `probe2DGraphicsStack()`

> returns `:object`

### `probe2DGdiOnly()`

> returns `:object`

### `probeNoGpu()`

> returns `:object`

### `select2DLayer(opts, opengl, ddraw, vulkan, d3d11)`

> returns `:atom`

### `getProbeTimings()`

> returns `:object`

### `select3DLayer(opts, d3d9)`

> returns `:atom`

