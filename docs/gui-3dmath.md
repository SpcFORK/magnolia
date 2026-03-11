# gui-3dmath

3D math utilities: vectors, rotations, projection helpers and common transforms used by the 3D renderer.

Key exports

- `Vec3(x,y,z)` — 3D vector
- `degToRad(deg)` — degrees to radians
- rotation helpers: `_rotateX`, `_rotateY`, `_rotateZ`
- projection helpers for perspective/orthographic mapping

Usage

Used by `gui-render` and `gui-raster` to transform and project 3D geometry to 2D screen space.

```oak
g3 := import('gui-3dmath')
v := g3.Vec3(1,2,3)
rot := g3._rotateY(v, g3.degToRad(45))
```
