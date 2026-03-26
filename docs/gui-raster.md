# gui-raster

Rasterization primitives: triangle fill (scanline), backface culling, clipping, and per-pixel shading helpers used by the renderer.

Key exports

- `drawTriangleFilled(window, p0, p1, p2, color?, borderColor?)` — scanline triangle fill; when `borderColor` is provided, triangle edges are drawn on top of the fill
- `drawMeshSolid(window, mesh, transform?, camera?, color?, light?, borderColor?)` — solid mesh rendering with optional wireframe edge overlay
- clipping helpers and viewport/codes utilities
- simple shading/lighting helpers consumed by `gui-render` for solid rendering

Notes

- This module implements the low-level raster algorithms used to fill triangles and perform painter sorting when needed.
