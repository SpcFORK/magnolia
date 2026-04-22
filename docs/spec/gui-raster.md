# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-3dmath`

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

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `gui-lighting`

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

## Module: `gui-thread`

- `threadLib` · `import(...)`
### `CommandQueue()`

> returns `:object`

### `FrameFence(workerCount)`

> returns `:object`

### `WorkerPool(numWorkers)`

> returns `:object`

### `StateGuard()`

> returns `:object`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

### `AsyncLoader(cmdQueue)`

> returns `:object`

### `FrameScheduler(pool, cmdQueue)`

> returns `:object`

### `initWindowThreading(window, options)`

### `threadingEnabled?(window)`

### `commandQueue(window)`

### `workerPool(window)`

### `scheduler(window)`

### `stateGuard(window)`

### `flushThreadedCommands(window)`

### `destroyWindowThreading(window)`

> returns `?`

## Module: `lib\gui-raster.oak`

- `guiLighting` · `import(...)`
- `gui3dmath` · `import(...)`
- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_pointOutCode(window, p)`

> returns `:bool`

### `_lineVisible?(window, a, b)`

### `_min2(a, b)`

### `_max2(a, b)`

### `_triVisible?(window, a, b, c)`

> returns `:bool`

### `_triArea2x(pa, pb, pc)`

> returns `:int`

### `_clipPolyEdge(verts, n, getVal, limit, isMin)`

> returns `:list`

### `_clipPolyToViewport(window, verts, n)`

> returns `:list`

### `_clipAndFillTriangle(deps, window, p0, p1, p2, color)`

> returns `:bool`

### `_lerpX(pa, pb, y)`

### `_buildEdgeTable(p0, p1, p2)`

> returns `:object`

### `_drawScanline(deps, window, y, xa, xb, color)`

> returns `?`

### `_edgeSlope(pa, pb)`

> returns `:int`

### `_fillScanStepped(deps, window, y, splitY, maxY, xaTop, xaBot, xb, slopeATop, slopeABot, slopeB, color)`

### `_fillScan(deps, window, y, maxY, p0, p1, p2, color)`

### `_sortTriByY(p0, p1, p2)`

> returns `:list`

### `drawTriangleFilled(deps, window, p0, p1, p2, color, borderColor)`

### `drawTriangleFilledAA(deps, window, p0, p1, p2, color, borderColor, bgColor)`

### `_concatLists(left, right, i, out)`

### `_compactTrisInPlace(tris)`

### `_insertionSortRange(arr, lo, hi)`

### `_partition(arr, lo, hi)`

### `_sortDepthRange(arr, lo, hi)`

### `_slice(arr, start, end, acc)`

### `_sortDepthInPlace(tris, count)`

### `_sortDepth(tris)`

### `_vecSub(deps, a, b)`

### `_vecCross(deps, a, b)`

### `_vecDot(a, b)`

### `_vecLen(v)`

### `_vecNormalize(deps, v)`

- `_projParamsCache` · `?`
- `_projCacheW` · `?`
- `_projCacheH` · `?`
- `_projCacheFov` · `?`
- `_projCacheZ` · `?`
- `_projCacheMode` · `?`
- `_projCacheOrtho` · `?`
### `_projectionParams(window, camera)`

### `_projectPointFast(p, params)`

> returns `:object`

### `projectVertices(deps, window, verts, camera)`

### `_computeMeshBounds(verts, i, mnX, mxX, mnY, mxY, mnZ, mxZ)`

> returns `:object`

### `_meshBounds(verts)`

> returns `:object`

### `_ensureMeshBounds(mesh)`

### `_transformBoundsQuick(localBounds, transform)`

> returns `:object`

### `_sphereInFrustum?(bounds, params, farPlane)`

> returns `:bool`

### `computeLightParams(deps, light)`

### `faceShadeGeneric(deps, lp, pa3, pb3, pc3)`

### `_colorR(c)`

> returns `:bool`

### `_colorG(c)`

> returns `:bool`

### `_colorB(c)`

> returns `:bool`

### `_shadeColor(deps, color, intensity)`

### `_frontFacing?(pa, pb, pc)`

> returns `:bool`

### `_computeDepth(pa3, pb3, pc3, pd3)`

### `_buildTriangle(pa, pb, pc, depth, color)`

> returns `:object`

### `_shouldCullTriangle?(backfaceCulling, pa, pb, pc, window)`

> returns `:bool`

### `_processFace(deps, window, i, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, acc)`

### `_collectFaces(deps, window, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, i, acc)`

### `_collectFacesRange(deps, window, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, start, end, acc)`

### `_drawTriangles(deps, window, tris, i, count, drawn)`

### `drawMeshSolid(deps, window, mesh, transform, camera, color, light, borderColor)`

> returns `:object`

### `drawMeshWireframe(deps, window, mesh, transform, camera, color)`

> returns `:object`

### `_processFaceLit(deps, window, i, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, acc)`

### `_collectFacesLit(deps, window, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, i, acc)`

### `_collectFacesLitRange(deps, window, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, start, end, acc)`

### `drawMeshLit(deps, window, mesh, transform, camera, color, scene, material, borderColor)`

> returns `:object`

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

