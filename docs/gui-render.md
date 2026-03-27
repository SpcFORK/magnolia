# gui-render

High-level 3D renderer API used by the GUI stack. Provides `Renderer3D` convenience object plus helpers to render wireframe and solid meshes.

Key exports

- `Renderer3D(deps, window, options)` — creates a renderer with methods like `renderMesh`, `renderMeshSolid`, `renderGrid`, `renderVoxels`, etc.
- `drawMeshWireframe(window, mesh, transform?, camera?, color?)`
- `drawMeshSolid(window, mesh, transform?, camera?, color?, light?, borderColor?)`
- Camera and transform fields documented in `docs/gui.md` (camera/options section)

Renderer3D methods for new mesh shapes

- `renderSphere(radius?, segments?, rings?, transform?, color?)` / `renderSphereSolid(..., borderColor?)`
- `renderPyramid(base?, height?, transform?, color?)` / `renderPyramidSolid(..., borderColor?)`
- `renderCylinder(radius?, height?, segments?, transform?, color?)` / `renderCylinderSolid(..., borderColor?)`
- `renderCone(radius?, height?, segments?, transform?, color?)` / `renderConeSolid(..., borderColor?)`
- `renderTorus(majorR?, minorR?, majSegs?, minSegs?, transform?, color?)` / `renderTorusSolid(..., borderColor?)`
- `renderPlane(width?, depth?, sw?, sd?, transform?, color?)` / `renderPlaneSolid(..., borderColor?)`
- `renderHemisphere(radius?, segments?, rings?, transform?, color?)` / `renderHemisphereSolid(..., borderColor?)`
- `renderWedge(width?, height?, depth?, transform?, color?)` / `renderWedgeSolid(..., borderColor?)`
- `renderTube(outerR?, innerR?, height?, segments?, transform?, color?)` / `renderTubeSolid(..., borderColor?)`
- `renderArrow(shaftR?, shaftH?, headR?, headH?, segments?, transform?, color?)` / `renderArrowSolid(..., borderColor?)`
- `renderPrism(radius?, height?, sides?, transform?, color?)` / `renderPrismSolid(..., borderColor?)`
- `renderStairs(steps?, width?, stepH?, stepD?, transform?, color?)` / `renderStairsSolid(..., borderColor?)`
- `renderIcosphere(radius?, transform?, color?)` / `renderIcosphereSolid(..., borderColor?)`

All new mesh render methods support both wireframe and solid modes and cache meshes by parameters.

## Border Color

All `*Solid` render methods accept an optional `borderColor` as their last parameter. When provided, the mesh is rendered solid first, then a wireframe edge overlay is drawn on top using `borderColor`. This creates a visible edge outline over filled 3D shapes.

```oak
gui := import('GUI')
renderer := gui.Renderer3D(window, { camera: { z: 5, fov: 90 } })

// Solid blue cube with white wireframe edges
renderer.renderCubeSolid(2, transform, gui.rgb(0, 80, 200), gui.rgb(255, 255, 255))

// Solid sphere without border — works as before
renderer.renderSphereSolid(1, 16, 12, transform, gui.rgb(200, 50, 50))
```

Usage

```oak
gr := import('gui-render')
renderer := gr.Renderer3D(deps, window, { camera: { z: 5, fov: 90 }, drawDistance: 100 })
renderer.renderMesh(mesh)
```

## Draw Distance / Far-Plane Culling

Set `drawDistance` in options (or call `renderer.setDrawDistance(dist)`) to cull meshes whose bounding sphere is entirely beyond that distance from the camera. A value of `0` (default) disables the far-plane limit.

```oak
renderer := gr.Renderer3D(deps, window, { drawDistance: 50 })
// or at runtime:
renderer.setDrawDistance(100)
```

## Rendering Optimizations

The 3D pipeline includes several automatic optimizations:

- **Frustum culling** — each `drawMesh*` call computes a bounding sphere and skips the mesh entirely if it falls outside the view frustum (behind near plane, beyond far plane, or off-screen)
- **Sub-pixel rejection** — triangles that project to less than 1 pixel of screen area are skipped
- **Adaptive depth sort** — painter sort uses in-place insertion sort for small face batches (≤16 triangles), reducing allocation overhead
- **Backface culling** — enabled by default (`camera.backfaceCulling`), skips triangles facing away from the camera
- **Mesh caching** — all primitive meshes are cached by parameters to avoid repeated geometry construction
