# gui-native-win

Win32-specific window lifecycle, event loop, and drawing helpers.

Key exports

- `createWindowState(title, width, height, options, className, frameMs, updateOnDispatch)` — create native Win32 window state
- `showWindow(window)` / `hideWindow(window)` — show or hide a native window
- `moveWindow(window, x, y)` — move a native window without resizing
- `resizeWindow(window, width, height)` — resize native window bounds
- `setFullscreen(window, enabled)` — maximize/restore fullscreen-like mode
- helper constants and icon handling utilities
- platform-specific frame batching to reuse device contexts

Parameters

- `options` - full GUI options object (layer selection, icons, and related backend flags)
- `className` - explicit Win32 class name to register for this window
- `frameMs` - target frame step in milliseconds
- `updateOnDispatch` - whether dispatch events should also drive frame updates

Layer options

- `options.layer2D` / `options.layer2d`: `auto` (default), `vulkan`, `opengl`, `ddraw`, or `gdi`
- `options.vulkanAuto`: opt-in Vulkan selection in `auto` mode (default: `false`)
- `options.dllLoadMode`: `sync` (default) or `async` for background DLL capability probing
- `options.layer3D` / `options.layer3d`: `auto` (default), `d3d9`, `cpu`, or `none`

Window state fields

- `window.className` — registered Win32 class name used by the window instance
- `window.presenterBackend` — selected 2D presenter backend (`:vulkan`, `:opengl`, `:ddraw`, or `:gdi`)
- `window.layers.twoD` — requested/selected 2D layer plus capability details
- `window.layers.threeD` — requested/selected 3D layer plus capability details
- `window.layers.twoD.init` — DirectDraw init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.threeD.init` — Direct3D9 init attempt result (`:ok`, `:warn`, `:error`, `:skipped`)
- `window.layers.twoD.fallback` / `window.layers.threeD.fallback` — populated when init errors trigger automatic fallback (`ddraw->gdi`, `d3d9->cpu`)
- `window.layers.twoD.vulkan` — Vulkan capability probe for 2D selection on Windows.
- `window.layers.twoD.vulkanInit` — Vulkan initialization attempt result preserved even when runtime uses OpenGL/DirectDraw/GDI for presentation.
- `window.layers.twoD.opengl` — OpenGL capability probe for staged 2D selection on Windows.
- `window.layers.twoD.openglInit` — OpenGL initialization attempt result preserved even when runtime falls back to DirectDraw/GDI.
- `window.layers.twoD.staged` — metadata for staged (not yet active) layer selections.
- `window.layers.twoD.staged.bootstrapped` — indicates whether staged backend bootstrap completed before another presenter became active.
- `window.layers.twoD.staged.presenter` — active stable presenter used while the staged backend remains inactive.
- `window.layers.twoD.init.path` / `window.layers.twoD.init.interface` — shows whether `DirectDrawCreateEx` (`IDirectDraw7`) or legacy `DirectDrawCreate` (`IDirectDraw`) was used.
- `window.layers.twoD.init.primarySurfaceCreated` — indicates whether a primary DirectDraw surface bootstrap succeeded during layer initialization.
- `window.ddrawPrimarySurface` — cached DirectDraw primary surface pointer (0 when unavailable or after fallback).
- `window.vulkanInstance` — cached Vulkan instance pointer (0 when unavailable or after cleanup).
- `window.vulkanSurface` — cached Vulkan Win32 surface pointer (0 when unavailable or after cleanup).
- `window.vulkanPhysicalDevice` — selected Vulkan physical device handle (0 when unavailable or after cleanup).
- `window.vulkanQueueFamily` — selected Vulkan graphics+present queue family index (`-1` when unavailable).
- `window.ready` — `true` after deferred layer bootstrap completes.
- `window._dllProbePending` — `true` while async DLL probing is still running.
- `window._dllLoadMode` — normalized DLL probe mode (`sync` or `async`).

Notes

- Prefer using the generic `GUI` facade; import this module when you need Win32-specific features.
- On Windows, the presenter probes `ddraw.dll` and enables a DDraw-assisted path when available; it falls back to the GDI blit path automatically when DDraw is unavailable.
- When DDraw is selected, initialization first attempts `DirectDrawCreateEx` and then falls back to legacy `DirectDrawCreate` before declaring init failure.
- If DirectDraw initialization is degraded (for example, object creation succeeds but primary surface bootstrap fails), the backend now automatically falls back to GDI for stable presentation.
- Vulkan 2D is available alongside OpenGL and can be requested explicitly with `layer2D: 'vulkan'`.
- Default 2D auto selection favors stable presenters (`opengl -> ddraw -> gdi`). Set `options.vulkanAuto: true` to allow Vulkan selection through `auto`.
- Vulkan support initializes an instance with `VK_KHR_surface` + `VK_KHR_win32_surface`, validates at least one physical device, and bootstraps a Win32 surface before activation; presentation still uses the existing OpenGL/DDraw/GDI present paths.
- Vulkan activation now also requires finding a queue family that supports both graphics and presentation for the created Win32 surface.
- Even when Vulkan bootstrap succeeds, Magnolia currently stages Vulkan state and uses a stable fallback presenter for actual frame display until swapchain-based Vulkan present is implemented.
- OpenGL-selected windows currently keep context/bootstrap state but present frames through the stable GDI blit path while full OpenGL present is staged.
- `SwapBuffers` is intentionally avoided in this staged path; on double-buffered pixel formats it swaps front/back buffers and can expose an unrendered (white) back buffer when actual OpenGL drawing is not yet active.
- Window classes for Vulkan/OpenGL selection now include `CS_OWNDC` to improve device-context stability on Windows.
- The registered Win32 window class uses a null background brush (`hbrBackground = 0`) so frame clears are controlled by the renderer and do not flash the system background between presents.
- When DirectDraw has a primary surface, `endFrame` attempts present via `IDirectDrawSurface7::GetDC` and GDI blit to that surface. If this per-frame present fails, the backend immediately falls back to GDI blit path for subsequent frames.
- If Direct3D9 initialization is degraded (for example, release status is not healthy), the backend now automatically falls back to CPU for stable 3D behavior.
- With `options.dllLoadMode: 'async'`, capability DLL probing (`opengl32.dll`, `vulkan-1.dll`, `ddraw.dll`, `d3d9.dll`) runs in a background task so `createWindow(...)` can return faster.

Related samples

- `samples/windows-higher-renderers.oak` — probes Vulkan/OpenGL/DDraw/D3D capability, initialization state, staged metadata, and presenter fallback.
- `samples/windows-2d-layer-hotload.oak` — recreates the window every 2 seconds to cycle Windows 2D layer requests while keeping one animated scene alive.
- `samples/windows-2d-layer-hotload-game.oak` — recreates the window every 2 seconds while preserving the bouncing-box game scene across Windows 2D layer requests.
- `samples/gui-sample.oak` — cross-platform quickstart; on Windows it displays the active presenter and fallback state in the scene.
- `samples/gui-game.oak` — bouncing-box sample with frame-rate independent motion and Windows presenter-state overlay.
