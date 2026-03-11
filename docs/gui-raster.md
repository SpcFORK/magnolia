# gui-raster

Rasterization primitives: triangle fill (scanline), backface culling, clipping, and per-pixel shading helpers used by the renderer.

Key exports

- `drawTriangleFilled(window, p0, p1, p2, color?)` — scanline triangle fill
- clipping helpers and viewport/codes utilities
- simple shading/lighting helpers consumed by `gui-render` for solid rendering

Notes

- This module implements the low-level raster algorithms used to fill triangles and perform painter sorting when needed.
