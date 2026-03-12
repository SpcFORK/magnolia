# gui-native-win

Win32-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, frameMs, updateOnDispatch)` — create native Win32 window state
- helper constants and icon handling utilities
- platform-specific frame batching to reuse device contexts

Layer options

- `options.layer2D` / `options.layer2d`: `auto` (default), `opengl` (staged), `ddraw`, or `gdi`
- `options.layer3D` / `options.layer3d`: `auto` (default), `d3d9`, `cpu`, or `none`

Window state fields

- `window.presenterBackend` — selected 2D presenter backend (`:ddraw` or `:gdi`)
- `window.layers.twoD` — requested/selected 2D layer plus capability details
- `window.layers.threeD` — requested/selected 3D layer plus capability details
- `window.layers.twoD.init` — DirectDraw init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.threeD.init` — Direct3D9 init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.twoD.fallback` / `window.layers.threeD.fallback` — populated when init errors trigger automatic fallback (`ddraw->gdi`, `d3d9->cpu`)
- `window.layers.twoD.opengl` — OpenGL capability probe for staged 2D selection on Windows.
- `window.layers.twoD.openglInit` — OpenGL initialization attempt result preserved even when runtime falls back to DirectDraw/GDI.
- `window.layers.twoD.staged` — metadata for staged (not yet active) layer selections.
- `window.layers.twoD.init.path` / `window.layers.twoD.init.interface` — shows whether `DirectDrawCreateEx` (`IDirectDraw7`) or legacy `DirectDrawCreate` (`IDirectDraw`) was used.
- `window.layers.twoD.init.primarySurfaceCreated` — indicates whether a primary DirectDraw surface bootstrap succeeded during layer initialization.
- `window.ddrawPrimarySurface` — cached DirectDraw primary surface pointer (0 when unavailable or after fallback).

Notes

- Prefer using the generic `GUI` facade; import this module when you need Win32-specific features.
- On Windows, the presenter probes `ddraw.dll` and enables a DDraw-assisted path when available; it falls back to the GDI blit path automatically when DDraw is unavailable.
- When DDraw is selected, initialization first attempts `DirectDrawCreateEx` and then falls back to legacy `DirectDrawCreate` before declaring init failure.
- If DirectDraw initialization is degraded (for example, object creation succeeds but primary surface bootstrap fails), the backend now automatically falls back to GDI for stable presentation.
- OpenGL 2D is currently staged above DirectDraw in selection priority when supported; it currently falls back to DirectDraw/GDI presenters until the active OpenGL present path is enabled.
- When DirectDraw has a primary surface, `endFrame` attempts present via `IDirectDrawSurface7::GetDC` and GDI blit to that surface. If this per-frame present fails, the backend immediately falls back to GDI blit path for subsequent frames.
- If Direct3D9 initialization is degraded (for example, release status is not healthy), the backend now automatically falls back to CPU for stable 3D behavior.
