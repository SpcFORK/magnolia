# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

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

