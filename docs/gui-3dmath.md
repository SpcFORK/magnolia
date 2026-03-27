# gui-3dmath

3D math utilities: vectors, rotations, projection helpers and common transforms used by the 3D renderer.

Key exports

- `Vec3(x,y,z)` — 3D vector
- `Vec3Add`, `Vec3Sub`, `Vec3Scale`, `Vec3Dot`, `Vec3Cross`, `Vec3Length`, `Vec3Normalize`, `Vec3Distance`, `Vec3Lerp`, `Vec3Negate`, `Vec3Reflect` — vector operations
- `degToRad(deg)` — degrees to radians
- rotation helpers: `_rotateX`, `_rotateY`, `_rotateZ`
- projection helpers for perspective/orthographic mapping
- `Mat4Identity`, `Mat4Translate`, `Mat4Scale`, `Mat4RotateX/Y/Z`, `Mat4Multiply`, `Mat4TransformPoint` — 4x4 matrix transforms
- `Quat`, `QuatFromAxisAngle`, `QuatMultiply`, `QuatNormalize`, `QuatRotateVector`, `QuatToMat4`, `QuatSlerp` — quaternion rotation

Usage

Used by `gui-render` and `gui-raster` to transform and project 3D geometry to 2D screen space.

```oak
g3 := import('gui-3dmath')
v := g3.Vec3(1,2,3)
rot := g3._rotateY(v, g3.degToRad(45))
```
