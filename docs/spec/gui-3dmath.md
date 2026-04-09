# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-3dmath.oak`

- `threadLib` · `import(...)`
- `PI` · `3.141592653589793`
### `degToRad(deg)`

### `Vec3(x, y, z)`

> returns `:object`

### `_rotateX(v, radians)`

> returns `:object`

### `_rotateY(v, radians)`

> returns `:object`

### `_rotateZ(v, radians)`

> returns `:object`

### `transformPoint(v, transform)`

> returns `:object`

### `projectPoint(window, p, camera)`

> returns `:object`

### `transformVertices(vertices, transform, i, out)`

### `_transformVertex(v, scale, tx, ty, tz, cx, sx, cy, sy, cz, sz)`

> returns `:object`

### `transformVerticesBatch(meshes, transforms, i, out)`

### `transformAndProjectVertices(vertices, transform, projParams)`

### `Mat4Identity()`

> returns `:list`

### `Mat4Translate(x, y, z)`

> returns `:list`

### `Mat4Scale(x, y, z)`

> returns `:list`

### `Mat4RotateX(radians)`

> returns `:list`

### `Mat4RotateY(radians)`

> returns `:list`

### `Mat4RotateZ(radians)`

> returns `:list`

### `Mat4Multiply(a, b)`

### `Mat4TransformPoint(m, v)`

### `Quat(x, y, z, w)`

> returns `:object`

### `QuatFromAxisAngle(axis, radians)`

### `QuatMultiply(a, b)`

### `QuatNormalize(q)`

### `QuatRotateVector(q, v)`

### `QuatToMat4(q)`

> returns `:list`

### `QuatSlerp(a, b, t)`

### `Vec3Add(a, b)`

### `Vec3Sub(a, b)`

### `Vec3Scale(v, s)`

### `Vec3Dot(a, b)`

### `Vec3Cross(a, b)`

### `Vec3Length(v)`

### `Vec3Normalize(v)`

### `Vec3Distance(a, b)`

### `Vec3Lerp(a, b, t)`

### `Vec3Negate(v)`

### `Vec3Reflect(v, n)`

### `AABB3(minX, minY, minZ, maxX, maxY, maxZ)`

> returns `:object`

### `Sphere3D(cx, cy, cz, r)`

> returns `:object`

### `Plane3D(nx, ny, nz, d)`

> returns `:object`

### `pointInAABB3(point, box)`

> returns `:bool`

### `pointInSphere3D(point, sphere)`

### `aabb3Intersects(a, b)`

> returns `:bool`

### `sphere3DIntersects(a, b)`

### `sphereAABB3Intersects(sphere, box)`

### `planePointDistance(plane, point)`

### `planeClassifyPoint(plane, point)`

> returns `:atom`

### `raySphere3DIntersect(origin, dir, sphere)`

> returns `:object`

### `rayAABB3Intersect(origin, dir, box)`

> returns `:object`

### `rayPlane3DIntersect(origin, dir, plane)`

> returns `:object`

### `aabb3ClosestPoint(box, point)`

### `aabb3Union(a, b)`

### `aabb3Center(box)`

### `aabb3HalfExtents(box)`

### `parallelTransformVertices(vertices, transform, numWorkers)`

