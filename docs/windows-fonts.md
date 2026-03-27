# windows-fonts

Win32 GDI font creation, metrics, measurement, caching, and management.

Builds on top of `windows-gdi`'s low-level `CreateFontW` wrapper and exposes
a higher-level API with named parameters, caching, and metric accessors.

## Dependencies

- `sys`, `writes`, `windows-constants`

## Constants

### Font Weights

`FW_DONTCARE` (0), `FW_THIN` (100), `FW_EXTRALIGHT` (200), `FW_LIGHT` (300),
`FW_NORMAL` (400), `FW_MEDIUM` (500), `FW_SEMIBOLD` (600), `FW_BOLD` (700),
`FW_EXTRABOLD` (800), `FW_HEAVY` (900)

### Charsets

`ANSI_CHARSET` (0), `DEFAULT_CHARSET` (1), `SYMBOL_CHARSET` (2),
`SHIFTJIS_CHARSET` (128), `HANGUL_CHARSET` (129), `OEM_CHARSET` (255), ...

### Quality

`DEFAULT_QUALITY` (0), `DRAFT_QUALITY` (1), `PROOF_QUALITY` (2),
`NONANTIALIASED_QUALITY` (3), `ANTIALIASED_QUALITY` (4),
`CLEARTYPE_QUALITY` (5)

### Pitch & Family

`DEFAULT_PITCH` (0), `FIXED_PITCH` (1), `VARIABLE_PITCH` (2),
`FF_ROMAN` (16), `FF_SWISS` (32), `FF_MODERN` (48), `FF_SCRIPT` (64),
`FF_DECORATIVE` (80)

### Stock Font Indices

`OEM_FIXED_FONT` (10), `ANSI_FIXED_FONT` (11), `ANSI_VAR_FONT` (12),
`SYSTEM_FONT` (13), `DEVICE_DEFAULT_FONT` (14), `DEFAULT_GUI_FONT` (17)

### Text Alignment

`TA_LEFT` (0), `TA_RIGHT` (2), `TA_CENTER` (6), `TA_TOP` (0),
`TA_BOTTOM` (8), `TA_BASELINE` (24)

## Functions

### `createFontEx(spec) -> result`

Creates a GDI font from a spec object. Returns
`{ type: :ok, handle, family, size, weight }`.

Spec fields (all optional except `family`):

| Field           | Default           |
|-----------------|-------------------|
| `family`        | `'Segoe UI'`      |
| `size`          | `16`              |
| `weight`        | `FW_NORMAL`       |
| `italic`        | `false`           |
| `underline`     | `false`           |
| `strikeOut`     | `false`           |
| `charSet`       | `DEFAULT_CHARSET` |
| `quality`       | `CLEARTYPE_QUALITY`|
| `pitch`         | `DEFAULT_PITCH`   |
| `escapement`    | `0`               |
| `orientation`   | `0`               |
| `outPrecision`  | `OUT_DEFAULT_PRECIS` |
| `clipPrecision` | `CLIP_DEFAULT_PRECIS`|

### `deleteFont(handle)`

Deletes a GDI font object.

### `selectFont(hdc, fontHandle) -> prevFont`

Selects a font into a device context; returns the previous font handle.

### `getStockFont(stockIndex) -> result`

Returns a stock font handle.

### `getDefaultGuiFont() -> result`

Shorthand for `getStockFont(DEFAULT_GUI_FONT)`.

### `setTextAlign(hdc, flags)`

Sets text alignment flags on a DC.

### `getTextMetrics(hdc) -> result`

Reads `TEXTMETRICW` for the selected font. Returns:
`{ type: :ok, height, ascent, descent, internalLeading, externalLeading, aveCharWidth, maxCharWidth, weight }`

### `getTextExtent(hdc, text) -> result`

Pixel size of a string: `{ type: :ok, cx, cy }`.

### `cachedFont(spec) -> result`

Returns a cached GDI font, creating one on first call.

### `releaseCachedFonts()`

Deletes all cached font handles.

### `removeCachedFont(spec)`

Removes and deletes a single cached entry.

### `measureText(hwnd, spec, text) -> result`

Measures text without requiring a pre-acquired DC.
Returns `{ type: :ok, width, height }`.

### `fontLineHeight(hdc) -> result`

Returns `{ type: :ok, lineHeight }` (height + external leading).

## Example

```oak
wf := import('windows-fonts')

font := wf.createFontEx({
    family: 'Consolas'
    size: 14
    weight: wf.FW_BOLD
    quality: wf.CLEARTYPE_QUALITY
})

if font.type = :ok -> {
    // select into DC, draw, etc.
    wf.deleteFont(font.handle)
}
```
