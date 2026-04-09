# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-mesh.oak`

- `threadLib` · `import(...)`
### `Vec3(x, y, z)`

> returns `:object`

### `Mesh(vertices, edges)`

> returns `:object`

### `GridMesh(size, step)`

### `AxesMesh(length)`

### `_vecKey(v)`

### `_edgeKey(a, b)`

### `_addVertex(vertices, indexByKey, v)`

> returns `:object`

### `_addEdge(edges, edgeSet, a, b)`

> returns `:object`

### `_voxelMeshSub(voxels, voxelSize, i, vertices, indexByKey, edges, edgeSet)`

### `VoxelMesh(voxels, voxelSize)`

### `PlaneMesh(width, depth, subdivisionsW, subdivisionsD)`

> returns `:object`

### `PyramidMesh(base, height)`

> returns `:object`

### `_cylinderRingVerts(cx, cy, cz, radius, segments, i, out)`

### `CylinderMesh(radius, height, segments)`

> returns `:object`

### `ConeMesh(radius, height, segments)`

> returns `:object`

### `SphereMesh(radius, segments, rings)`

> returns `:object`

### `TorusMesh(majorRadius, minorRadius, majorSegments, minorSegments)`

> returns `:object`

### `VoxelGrid(options)`

> returns `:object`

### `HemisphereMesh(radius, segments, rings)`

> returns `:object`

### `WedgeMesh(width, height, depth)`

> returns `:object`

### `TubeMesh(outerRadius, innerRadius, height, segments)`

> returns `:object`

### `ArrowMesh(shaftRadius, shaftHeight, headRadius, headHeight, segments)`

> returns `:object`

### `PrismMesh(radius, height, sides)`

> returns `:object`

### `StairsMesh(steps, width, stepHeight, stepDepth)`

### `IcosphereMesh(radius)`

> returns `:object`

### `parallelMeshGenerate(specs)`

