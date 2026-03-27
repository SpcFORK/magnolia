# Shader SDF & Patterns (gui-shader-sdf)

## Overview

`gui-shader-sdf` provides signed distance field primitives, boolean operations,
fill/stroke helpers, and procedural patterns for the CPU shader engine. It
depends on `gui-shader-math`, `gui-shader-color`, and `gui-shader-noise`.

## Import

```oak
sdf := import('gui-shader-sdf')
```

All symbols are also re-exported by `gui-shader` and through the `GUI` facade.

## SDF Primitives

All primitives return a signed distance: negative inside, zero on the boundary,
positive outside.

| Function | Signature | Description |
|----------|-----------|-------------|
| `sdCircle` | `(px, py, cx, cy, r)` | Distance from (px,py) to circle centered at (cx,cy) with radius r |
| `sdBox` | `(px, py, cx, cy, hw, hh)` | Distance to axis-aligned box centered at (cx,cy) with half-extents (hw,hh) |
| `sdLine` | `(px, py, ax, ay, bx, by)` | Distance from point to line segment A→B |
| `sdRoundedBox` | `(px, py, cx, cy, hw, hh, r)` | Distance to rounded box with corner radius r |

## Fill & Stroke Helpers

| Function | Signature | Description |
|----------|-----------|-------------|
| `sdfFill` | `(d, color)` | Returns color if d ≤ 0, else `?` |
| `sdfSmoothFill` | `(d, color, bg, edge)` | Anti-aliased fill with smoothstep transition |
| `sdfStroke` | `(d, thickness, color)` | Returns color if |d| ≤ thickness/2 |
| `sdfGlow` | `(d, color, intensity, radius)` | Radial glow falloff around the SDF boundary |

## Boolean Operations

| Function | Description |
|----------|-------------|
| `sdUnion(d1, d2)` | Union (min) |
| `sdIntersect(d1, d2)` | Intersection (max) |
| `sdSubtract(d1, d2)` | Subtraction (d1 minus d2) |
| `sdSmoothUnion(d1, d2, k)` | Smooth union with blending radius k |
| `sdSmoothIntersect(d1, d2, k)` | Smooth intersection |
| `sdSmoothSubtract(d1, d2, k)` | Smooth subtraction |
| `sdAnnular(d, r)` | Annular ring (hollow shape) with width r |
| `sdRepeat2(px, py, cx, cy)` | Tiled repetition with cell size (cx, cy); returns the local position |

## Procedural Patterns

| Function | Signature | Returns |
|----------|-----------|---------|
| `checkerboard` | `(x, y, size)` | 0 or 1 (alternating squares) |
| `stripes` | `(x, y, angle, width)` | 0 or 1 (directional stripes) |
| `grid` | `(x, y, size, thickness)` | 0 or 1 (grid lines) |
| `dots` | `(x, y, spacing, radius)` | 0 or 1 (dot grid) |
| `voronoi` | `(x, y, scale)` | `{ dist, cellId }` — Voronoi cell distance and a per-cell random ID |
