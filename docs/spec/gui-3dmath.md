# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

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

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

