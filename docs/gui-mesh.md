# gui-mesh

Mesh and voxel builders used by the 3D renderer and voxel utilities.

Key exports

- `Vec3(x,y,z)` — vertex constructor
- `Mesh(vertices?, edges?)` — wireframe mesh
- `CubeMesh(size?)`, `GridMesh(size?, step?)`, `AxesMesh(length?)`
- `VoxelMesh(voxels, voxelSize?)`, `VoxelGrid(options?)`
- `SphereMesh(radius?, segments?, rings?)` — UV sphere
- `PyramidMesh(base?, height?)` — four-sided pyramid
- `CylinderMesh(radius?, height?, segments?)` — cylinder along Y axis
- `ConeMesh(radius?, height?, segments?)` — cone with apex at top
- `TorusMesh(majorRadius?, minorRadius?, majorSegments?, minorSegments?)` — torus (donut)
- `PlaneMesh(width?, depth?, subdivisionsW?, subdivisionsD?)` — flat XZ plane

Notes

- Mesh objects are `{ vertices: [...], edges: [...], faces?: [...] }` used by `gui-render`.
- New mesh types include `faces` arrays for solid rendering.
- `VoxelGrid` provides a mutable voxel set with helper methods and `toMesh()`.
