# gui-resolution

Virtual / design resolution and scaling for the Magnolia GUI system.

## Overview

Games and applications often want to render at a fixed **logical resolution**
(e.g. 320×180 for pixel art) while displaying in a larger window
(e.g. 1280×720).  `gui-resolution` provides this capability transparently:

1. The back-buffer is allocated at the **design** (logical) dimensions.
2. All drawing calls operate in the logical coordinate space.
3. At frame present time, the back-buffer is scaled up to the physical
   window using `StretchBlt` (GDI backend).
4. Mouse / input coordinates are automatically mapped from physical to
   logical so game code doesn't need to convert.

## Quick Start

```oak
gui := import('GUI')

// Option A: via createWindow options
window := gui.createWindow('My Game', 1280, 720, {
    designWidth: 320
    designHeight: 180
    scaleMode: 'fit'
    pixelPerfect: true
})

// Option B: set after window creation
window := gui.createWindow('My Game', 1280, 720, {})
gui.setDesignResolution(window, 320, 180, {
    scaleMode: 'fit'
    pixelPerfect: true
})

// All drawing uses 320×180 coordinates
gui.fillRect(window, 0, 0, 320, 180, gui.rgb(20, 20, 40))
gui.drawText(window, 10, 10, 'Hello!', gui.rgb(255, 255, 255))

// window.width = 320, window.height = 180 (logical)
// gui.physicalWidth(window) = 1280 (actual window)
```

## Scale Modes

| Mode        | Behaviour                                            |
|-------------|------------------------------------------------------|
| `'fit'`     | Maintain aspect ratio, add letterbox / pillarbox bars |
| `'fill'`    | Maintain aspect ratio, crop edges to fill window      |
| `'stretch'` | Stretch to fill, may distort aspect ratio             |

## Pixel-Perfect Mode

When `pixelPerfect: true` is set with `'fit'` mode, the scale factor is
rounded down to the nearest integer (1×, 2×, 3×, …).  This ensures every
logical pixel maps to exactly N×N physical pixels with no interpolation
artefacts — ideal for pixel-art games.

## API Reference

### GUI.oak exports

| Function | Description |
|----------|-------------|
| `setDesignResolution(window, w, h, opts)` | Set the logical resolution |
| `clearDesignResolution(window)` | Remove virtual resolution (1:1) |
| `hasDesignResolution?(window)` | Check if a design resolution is active |
| `designWidth(window)` | Logical width |
| `designHeight(window)` | Logical height |
| `physicalWidth(window)` | Actual window client width |
| `physicalHeight(window)` | Actual window client height |
| `resolutionScaleX(window)` | Horizontal scale factor |
| `resolutionScaleY(window)` | Vertical scale factor |
| `resolutionOffsetX(window)` | Horizontal letterbox offset (pixels) |
| `resolutionOffsetY(window)` | Vertical letterbox offset (pixels) |
| `physicalToLogical(window, px, py)` | Convert physical → logical coords |
| `logicalToPhysical(window, lx, ly)` | Convert logical → physical coords |

### Options for setDesignResolution / createWindow

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `designWidth` | int | — | Logical width in px |
| `designHeight` | int | — | Logical height in px |
| `scaleMode` | string | `'fit'` | `'fit'`, `'fill'`, or `'stretch'` |
| `pixelPerfect` | bool | `false` | Integer scale factors only |

## Mouse Coordinate Mapping

When a design resolution is active, all mouse event handlers
(`onMouseMove`, `onLButtonDown`, `onLButtonUp`, `onRButtonDown`,
`onRButtonUp`) and `formEventContext` automatically deliver coordinates
in the logical space.  No manual conversion is needed in game code.

For raw physical coordinates, access the Win32 message LPARAM directly.

## Implementation Notes

- The GDI back-buffer (`frameHdc` / `frameBitmap`) is created at the
  logical resolution.
- `endFrame` → `presentFrameViaGdi` uses `StretchBlt` to scale from
  the logical back-buffer to the physical window DC.
- Letterbox bars are painted with `PatBlt(BLACKNESS)`.
- `SetStretchBltMode` is set to `COLORONCOLOR` for pixel-perfect or
  `HALFTONE` for smooth scaling.
