# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-native-win-ddraw.oak`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `COM_RELEASE` · `2`
- `IDirectDraw7_CreateSurface` · `6`
- `IDirectDraw7_SetCooperativeLevel` · `20`
- `IDirectDrawSurface7_GetDC` · `17`
- `IDirectDrawSurface7_ReleaseDC` · `26`
- `DDSCL_NORMAL` · `8`
- `DDSD_CAPS` · `1`
- `DDSCAPS_PRIMARYSURFACE` · `512`
- `SRCCOPY` · `13369376`
### `_iidDirectDraw7()`

### `_ddrawCreatePrimarySurface(ddrawObj)`

> returns `:object`

### `_ddrawPresentViaPrimarySurface(window)`

> returns `:object`

### `initDdraw2DLayer(window)`

> returns `:object`

### `presentFrameViaDdraw(window)`

> returns `:object`

### `releasePrimarySurface(window)`

> returns `:int`

