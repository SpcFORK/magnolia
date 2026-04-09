# API Documentation

_Auto-generated from Magnolia source._

---

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

