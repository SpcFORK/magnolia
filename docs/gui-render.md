# gui-render

High-level 3D renderer API used by the GUI stack. Provides `Renderer3D` convenience object plus helpers to render wireframe and solid meshes.

Key exports

- `Renderer3D(deps, window, options)` — creates a renderer with methods like `renderMesh`, `renderMeshSolid`, `renderGrid`, `renderVoxels`, etc.
- `drawMeshWireframe(window, mesh, transform?, camera?, color?)`
- `drawMeshSolid(window, mesh, transform?, camera?, color?)`
- Camera and transform fields documented in `docs/gui.md` (camera/options section)

Usage

```oak
gr := import('gui-render')
renderer := gr.Renderer3D(deps, window, { camera: { z: 5, fov: 90 } })
renderer.renderMesh(mesh)
```
