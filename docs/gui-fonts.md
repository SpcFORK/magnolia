# gui-fonts

Cross-platform font abstraction layer for Magnolia GUI applications.

Dispatches to `windows-fonts` or `linux-fonts` depending on the runtime
platform, providing a unified API for font creation, caching, measurement,
and lifecycle management.

## Font Spec

All functions accept a **font spec** object:

```oak
spec := {
    family: 'Segoe UI'   // Win32 font family or X11 font name / XLFD
    size: 16              // point size (Windows) / pixel size (Linux)
    weight: 400           // 100-900; 400 = normal, 700 = bold
    italic: false
    underline: false
    strikeOut: false
}
```

On the `:web` backend the spec may also include a raw `css` field:

```oak
spec := { css: 'italic bold 14px monospace' }
```

## Constants

| Name           | Value | Description     |
|----------------|-------|-----------------|
| `FW_THIN`      | 100   | Thin weight     |
| `FW_EXTRALIGHT`| 200   | Extra-light     |
| `FW_LIGHT`     | 300   | Light           |
| `FW_NORMAL`    | 400   | Normal / Regular|
| `FW_MEDIUM`    | 500   | Medium          |
| `FW_SEMIBOLD`  | 600   | Semi-bold       |
| `FW_BOLD`      | 700   | Bold            |
| `FW_EXTRABOLD` | 800   | Extra-bold      |
| `FW_HEAVY`     | 900   | Heavy / Black   |

## Functions

### `defaultFontSpec()`

Returns a font spec object with sensible per-platform defaults.

### `createFont(spec) -> result`

Creates a platform-native font.

- **Windows**: returns `{ type: :ok, handle, family, size, weight }`
- **Linux**: returns `{ type: :ok, fontStruct, ... }`
- **Web**: returns `{ type: :ok, css, family, size }`

### `deleteFont(fontResult)`

Releases resources associated with a font returned by `createFont`.

### `fontKey(spec) -> string`

Returns a canonical cache key string for the given spec.

### `cachedFont(windowOrDisplay, spec) -> result`

Returns a cached font handle, creating one if not already cached.

### `releaseCachedFonts(windowOrDisplay)`

Frees all cached fonts.

### `measureText(windowOrDisplay, spec, text) -> result`

Measures text dimensions. Returns `{ type: :ok, width, height }`.

### `selectFont(args)`

Selects a font into a drawing context.

- **Windows**: `{ hdc: hdc, handle: fontHandle }`
- **Linux**: `{ display: dpy, gc: gc, fontId: fid }`

### `getTextMetrics(hdc) -> result`

Returns TEXTMETRIC fields for the currently selected font (Windows only).

### `fontLineHeight(hdcOrFontStruct) -> result`

Returns `{ type: :ok, lineHeight: N }`.

### `buildXLFD(spec) -> string`

Constructs an X Logical Font Description string (Linux only).

## Example

```oak
fonts := import('gui-fonts')

// Create a bold font
font := fonts.createFont({ family: 'Consolas', size: 14, weight: fonts.FW_BOLD })
if font.type = :ok -> {
    // Use font.handle (Windows) or font.fontStruct (Linux)
    // ...
    fonts.deleteFont(font)
}
```
