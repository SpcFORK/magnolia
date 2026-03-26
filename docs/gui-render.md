# gui-render

High-level 3D renderer API used by the GUI stack. Provides `Renderer3D` convenience object plus helpers to render wireframe and solid meshes.

Key exports

- `Renderer3D(deps, window, options)` — creates a renderer with methods like `renderMesh`, `renderMeshSolid`, `renderGrid`, `renderVoxels`, etc.
- `drawMeshWireframe(window, mesh, transform?, camera?, color?)`
- `drawMeshSolid(window, mesh, transform?, camera?, color?)`
- Camera and transform fields documented in `docs/gui.md` (camera/options section)

Renderer3D methods for new mesh shapes

- `renderSphere(radius?, segments?, rings?, transform?, color?)` / `renderSphereSolid(...)`
- `renderPyramid(base?, height?, transform?, color?)` / `renderPyramidSolid(...)`
- `renderCylinder(radius?, height?, segments?, transform?, color?)` / `renderCylinderSolid(...)`
- `renderCone(radius?, height?, segments?, transform?, color?)` / `renderConeSolid(...)`
- `renderTorus(majorR?, minorR?, majSegs?, minSegs?, transform?, color?)` / `renderTorusSolid(...)`
- `renderPlane(width?, depth?, sw?, sd?, transform?, color?)` / `renderPlaneSolid(...)`

All new mesh render methods support both wireframe and solid modes and cache meshes by parameters.

Usage

```oak
gr := import('gui-render')
renderer := gr.Renderer3D(deps, window, { camera: { z: 5, fov: 90 } })
renderer.renderMesh(mesh)
```
