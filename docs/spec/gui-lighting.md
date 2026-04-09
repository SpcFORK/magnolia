# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-lighting.oak`

- `threadLib` · `import(...)`
- `PI` · `3.141592653589793`
### `_len3(x, y, z)`

### `_norm3(x, y, z)`

> returns `:list`

### `_dot(ax, ay, az, bx, by, bz)`

### `DirectionalLight(options)`

> returns `:object`

### `PointLight(options)`

> returns `:object`

### `SpotLight(options)`

> returns `:object`

### `AmbientLight(options)`

> returns `:object`

### `Material(options)`

> returns `:object`

### `LightScene(options)`

> returns `:object`

### `addLight(scene, light)`

### `removeLight(scene, index)`

### `clearLights(scene)`

### `lightCount(scene)`

### `faceNormal(pa, pb, pc)`

### `faceCenter(pa, pb, pc)`

> returns `:object`

### `_shadeDirectional(light, mat, n0, n1, n2, vdx, vdy, vdz, acc)`

### `_shadePoint(light, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, acc)`

### `_shadeSpot(light, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, acc)`

### `_shadeAmbient(light, mat, acc)`

### `_accumulate(lights, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, i, acc)`

### `shadeFaceColor(deps, baseColor, scene, material, pa3, pb3, pc3, camPos)`

### `shadeFaceIntensity(scene, material, pa3, pb3, pc3, camPos)`

### `prepareScene(scene, material, camPos)`

> returns `:object`

### `_prepareLights(lights, i, acc)`

### `shadeFaceColorFast(rgbFn, baseColor, prep, pa3, pb3, pc3)`

### `shadeFacesBatchParallel(rgbFn, faces, prepared, numWorkers)`

