# gui-mesh

Mesh and voxel builders used by the 3D renderer and voxel utilities.

Key exports

- `Vec3(x,y,z)` — vertex constructor
- `Mesh(vertices?, edges?)` — wireframe mesh
- `CubeMesh(size?)`, `GridMesh(size?, step?)`, `AxesMesh(length?)`
- `VoxelMesh(voxels, voxelSize?)`, `VoxelGrid(options?)`

Notes

- Mesh objects are simple `{ vertices: [...], edges: [...] }` used by `gui-render`.
- `VoxelGrid` provides a mutable voxel set with helper methods and `toMesh()`.
