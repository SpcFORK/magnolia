# gui-native-win

Win32-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, frameMs, updateOnDispatch)` — create native Win32 window state
- helper constants and icon handling utilities
- platform-specific frame batching to reuse device contexts

Layer options

- `options.layer2D` / `options.layer2d`: `auto` (default), `vulkan`, `opengl`, `ddraw`, or `gdi`
- `options.layer3D` / `options.layer3d`: `auto` (default), `d3d9`, `cpu`, or `none`

Window state fields

- `window.presenterBackend` — selected 2D presenter backend (`:vulkan`, `:opengl`, `:ddraw`, or `:gdi`)
- `window.layers.twoD` — requested/selected 2D layer plus capability details
- `window.layers.threeD` — requested/selected 3D layer plus capability details
- `window.layers.twoD.init` — DirectDraw init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.threeD.init` — Direct3D9 init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.twoD.fallback` / `window.layers.threeD.fallback` — populated when init errors trigger automatic fallback (`ddraw->gdi`, `d3d9->cpu`)
- `window.layers.twoD.vulkan` — Vulkan capability probe for 2D selection on Windows.
- `window.layers.twoD.vulkanInit` — Vulkan initialization attempt result preserved even when runtime falls back to OpenGL/DirectDraw/GDI.
- `window.layers.twoD.opengl` — OpenGL capability probe for staged 2D selection on Windows.
- `window.layers.twoD.openglInit` — OpenGL initialization attempt result preserved even when runtime falls back to DirectDraw/GDI.
- `window.layers.twoD.staged` — metadata for staged (not yet active) layer selections.
- `window.layers.twoD.init.path` / `window.layers.twoD.init.interface` — shows whether `DirectDrawCreateEx` (`IDirectDraw7`) or legacy `DirectDrawCreate` (`IDirectDraw`) was used.
- `window.layers.twoD.init.primarySurfaceCreated` — indicates whether a primary DirectDraw surface bootstrap succeeded during layer initialization.
- `window.ddrawPrimarySurface` — cached DirectDraw primary surface pointer (0 when unavailable or after fallback).
- `window.vulkanInstance` — cached Vulkan instance pointer (0 when unavailable or after cleanup).

Notes

- Prefer using the generic `GUI` facade; import this module when you need Win32-specific features.
- On Windows, the presenter probes `ddraw.dll` and enables a DDraw-assisted path when available; it falls back to the GDI blit path automatically when DDraw is unavailable.
- When DDraw is selected, initialization first attempts `DirectDrawCreateEx` and then falls back to legacy `DirectDrawCreate` before declaring init failure.
- If DirectDraw initialization is degraded (for example, object creation succeeds but primary surface bootstrap fails), the backend now automatically falls back to GDI for stable presentation.
- Vulkan 2D is now available alongside OpenGL and sits at the top of auto selection priority (`vulkan -> opengl -> ddraw -> gdi`).
- Vulkan support initializes an instance and validates at least one physical device before activation; presentation still uses the existing OpenGL/DDraw/GDI present paths.
- Window classes for Vulkan/OpenGL selection now include `CS_OWNDC` to improve device-context stability on Windows.
- When DirectDraw has a primary surface, `endFrame` attempts present via `IDirectDrawSurface7::GetDC` and GDI blit to that surface. If this per-frame present fails, the backend immediately falls back to GDI blit path for subsequent frames.
- If Direct3D9 initialization is degraded (for example, release status is not healthy), the backend now automatically falls back to CPU for stable 3D behavior.
