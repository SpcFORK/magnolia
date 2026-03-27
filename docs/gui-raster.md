# gui-raster

Rasterization primitives: triangle fill (scanline), backface culling, frustum culling, clipping, and per-pixel shading helpers used by the renderer.

Key exports

- `drawTriangleFilled(window, p0, p1, p2, color?, borderColor?)` — scanline triangle fill; when `borderColor` is provided, triangle edges are drawn on top of the fill
- `drawMeshSolid(window, mesh, transform?, camera?, color?, light?, borderColor?)` — solid mesh rendering with optional wireframe edge overlay
- clipping helpers and viewport/codes utilities
- simple shading/lighting helpers consumed by `gui-render` for solid rendering

## Optimizations

- **O(1) frustum culling** — mesh bounding spheres are computed once in local space and cached on the mesh object (`_localBounds`). At render time only the sphere center is transformed and the radius scaled, making the frustum check O(1) instead of O(N) per frame.
- **Combined transform + project pass** — `drawMeshSolid`, `drawMeshLit`, and `drawMeshWireframe` use `transformAndProjectVertices` from `gui-3dmath` to produce both world-space and screen-space vertex arrays in a single loop, eliminating the intermediate array allocation.
- **Scanline clipping** — `_drawScanline` and `_fillScan` clip scanlines to the window bounds, skipping rows and pixels outside the viewport to avoid out-of-bounds draws.
- **Sub-pixel triangle rejection** — triangles whose projected screen area is less than 1 pixel are skipped during face collection, avoiding unnecessary scanline fill.
- **Insertion sort for small batches** — depth sorting uses in-place insertion sort for ≤16 triangles, avoiding allocation overhead of merge sort at the base case.
- **Far-plane culling** — when `camera.farPlane` is set to a positive value, meshes whose bounding sphere is entirely beyond that distance are culled.
- **Skip compact pass** — `drawMeshLit` skips the `_compactTris` pass since its face processor never produces null entries. `_insertionSortDepth` also handles null entries defensively.

Camera fields consumed by culling:

- `farPlane` (default `0` = disabled) — maximum view distance for frustum culling
- `backfaceCulling` (default `true`) — enable/disable backface culling

Notes

- This module implements the low-level raster algorithms used to fill triangles and perform painter sorting when needed.
