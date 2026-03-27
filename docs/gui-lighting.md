# gui-lighting

Multi-light Blinn-Phong shading system for 3D scene rendering. Integrates with `gui-raster` and `gui-render` (via `Renderer3D`) to support directional, point, spot, and ambient lights with per-channel colored illumination and specular highlights.

## Light Types

### `DirectionalLight(options?)`

Parallel rays from a direction (sun-like). No distance attenuation.

| Option     | Default | Description                        |
|------------|---------|------------------------------------|
| `x`        | `0`     | Direction X component              |
| `y`        | `-1`    | Direction Y component              |
| `z`        | `0`     | Direction Z component              |
| `intensity`| `1`     | Brightness multiplier (0–10)       |
| `colorR`   | `255`   | Red channel of light color (0–255) |
| `colorG`   | `255`   | Green channel (0–255)              |
| `colorB`   | `255`   | Blue channel (0–255)               |

### `PointLight(options?)`

Radial light from a position with inverse-square distance attenuation.

| Option     | Default | Description                        |
|------------|---------|------------------------------------|
| `px`       | `0`     | Position X                         |
| `py`       | `2`     | Position Y                         |
| `pz`       | `0`     | Position Z                         |
| `intensity`| `1`     | Brightness multiplier (0–10)       |
| `radius`   | `10`    | Maximum influence radius           |
| `colorR`   | `255`   | Red channel (0–255)                |
| `colorG`   | `255`   | Green channel (0–255)              |
| `colorB`   | `255`   | Blue channel (0–255)               |

### `SpotLight(options?)`

Cone of light from a position with angular and distance falloff.

| Option     | Default | Description                            |
|------------|---------|----------------------------------------|
| `px`       | `0`     | Position X                             |
| `py`       | `3`     | Position Y                             |
| `pz`       | `0`     | Position Z                             |
| `x`        | `0`     | Spot direction X                       |
| `y`        | `-1`    | Spot direction Y                       |
| `z`        | `0`     | Spot direction Z                       |
| `intensity`| `1`     | Brightness multiplier (0–10)           |
| `radius`   | `10`    | Maximum influence radius               |
| `angle`    | `30`    | Half-angle of the cone in degrees (1–90) |
| `penumbra` | `5`     | Soft-edge penumbra width in degrees (0–45) |
| `colorR`   | `255`   | Red channel (0–255)                    |
| `colorG`   | `255`   | Green channel (0–255)                  |
| `colorB`   | `255`   | Blue channel (0–255)                   |

### `AmbientLight(options?)`

Uniform omnidirectional illumination.

| Option     | Default | Description                        |
|------------|---------|------------------------------------|
| `intensity`| `0.2`   | Brightness multiplier (0–1)        |
| `colorR`   | `255`   | Red channel (0–255)                |
| `colorG`   | `255`   | Green channel (0–255)              |
| `colorB`   | `255`   | Blue channel (0–255)               |

## Material

### `Material(options?)`

Defines surface shading properties.

| Option      | Default | Description                                  |
|-------------|---------|----------------------------------------------|
| `ambient`   | `0.2`   | Ambient response coefficient (0–1)           |
| `diffuse`   | `0.8`   | Diffuse (Lambertian) coefficient (0–1)       |
| `specular`  | `0.3`   | Specular (Blinn-Phong) coefficient (0–1)     |
| `shininess` | `32`    | Specular exponent / glossiness (1–256)       |
| `emissiveR` | `0`     | Self-illumination red channel (0–255)        |
| `emissiveG` | `0`     | Self-illumination green channel (0–255)      |
| `emissiveB` | `0`     | Self-illumination blue channel (0–255)       |

## LightScene

### `LightScene(options?)`

A collection of lights applied together.

| Option          | Default | Description                           |
|-----------------|---------|---------------------------------------|
| `lights`        | `[]`    | Initial array of light objects        |
| `globalAmbient` | `0.1`   | Base ambient illumination level (0–1) |

### `addLight(scene, light)` → scene
Add a light to the scene. Mutates and returns the scene.

### `removeLight(scene, index)` → scene
Remove the light at `index`. Mutates and returns the scene.

### `clearLights(scene)` → scene
Remove all lights. Mutates and returns the scene.

### `lightCount(scene)` → int
Number of lights in the scene.

## Geometry Helpers

### `faceNormal(pa, pb, pc)` → `[nx, ny, nz]`
Compute the unit normal of a triangle from three world-space vertices.

### `faceCenter(pa, pb, pc)` → `{x, y, z}`
Average position (centroid) of a triangle.

## Shading Functions

### `shadeFaceColor(baseColor, scene, material, pa3, pb3, pc3, camPos)` → packed RGB
Full per-channel Blinn-Phong shading. Applies every light in the scene to the triangle and returns a packed RGB color. Called through the GUI facade which provides `deps.rgb`.

### `shadeFaceIntensity(scene, material, pa3, pb3, pc3, camPos)` → float
Returns a single `[0..1]` greyscale intensity — useful for backward-compatible shading.

## Performance: Fast-Path API

For batch rendering, `drawMeshLit` (and all `Renderer3D` lit methods) automatically use the fast-path internally. Two additional functions are exported for advanced use:

### `prepareScene(scene, material, camPos)` → prepared
Pre-validates inputs, pre-normalizes directional/spot light directions, pre-computes cos values for spot cones, and caches `hasSpecular` flag to avoid per-face comparison. Call once per frame or per mesh to amortize setup cost.

### `shadeFaceColorFast(baseColor, prep, pa3, pb3, pc3)` → packedRGB
Uses a prepared scene object. Inlines all face normal, center, and view-direction math, and uses mutation-based light accumulation to avoid per-light object allocation. Roughly 2–3× fewer interpreter function calls per face compared to `shadeFaceColor`.

## Renderer3D Integration

`Renderer3D` accepts `lightScene` and `material` in its options and exposes:

- `setLightScene(scene)` — set/replace the active light scene
- `setMaterial(mat)` — set/replace the active material
- `renderMeshLit(mesh, transform, color?, material?, borderColor?)` — render any mesh with scene lighting
- `renderCubeLit(size, transform, color?, material?, borderColor?)`
- `renderSphereLit(radius, segs, rings, transform, color?, material?, borderColor?)`
- `renderPlaneLit(w, d, sw, sd, transform, color?, material?, borderColor?)`
- `renderCylinderLit(r, h, segs, transform, color?, material?, borderColor?)`
- `renderConeLit(r, h, segs, transform, color?, material?, borderColor?)`
- `renderTorusLit(majR, minR, majSegs, minSegs, transform, color?, material?, borderColor?)`
- `renderPyramidLit(base, h, transform, color?, material?, borderColor?)`
- `renderIcosphereLit(r, transform, color?, material?, borderColor?)`

Pass `materialOverride` to use a per-object material; otherwise the renderer's active material is used.

## Usage Example

```oak
gui := import('GUI')

window := gui.createWindow('Lit Scene', 960, 540)
gui.show(window)

scene := gui.LightScene({ globalAmbient: 0.15 })
gui.addLight(scene, gui.DirectionalLight({ x: 0.5, y: -0.8, z: -0.3, intensity: 0.9 }))
gui.addLight(scene, gui.PointLight({ px: 2, py: 3, pz: -1, intensity: 1.2, radius: 12, colorR: 255, colorG: 200, colorB: 150 }))
gui.addLight(scene, gui.AmbientLight({ intensity: 0.15 }))

mat := gui.Material({ specular: 0.5, shininess: 64 })

r := gui.Renderer3D(window, {
    lightScene: scene
    material: mat
})

gui.run(window, ?, fn(w) {
    gui.beginFrame(w)
    r.clear()
    r.renderCubeLit(2, { ry: 30, tx: -2 }, gui.rgb(200, 100, 100))
    r.renderSphereLit(1, 24, 12, { tx: 2 }, gui.rgb(100, 100, 200), gui.Material({ specular: 0.8, shininess: 128 }))
    gui.endFrame(w)
})
```
