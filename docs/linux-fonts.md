# linux-fonts

X11 core font loading, querying, metrics, measurement, and caching for Linux.

Uses Xlib's built-in font functions (`XLoadQueryFont`, `XLoadFont`,
`XSetFont`, `XTextWidth`, `XFreeFont`, etc.) to provide an Oak-friendly
font API on Linux.

## Dependencies

- `sys`, `writes`, `linux-core`, `linux-loader`

## Xlib Font Wrappers

### `loadQueryFont(display, name) -> result`

Loads and queries an X11 core font by name (XLFD or alias).
Returns `{ type: :ok, fontStruct: <ptr> }`.

### `loadFont(display, name) -> result`

Loads an X11 font and returns its Font XID.
Returns `{ type: :ok, fontId }`.

### `setFont(display, gc, fontId)`

Sets the font for a graphics context.

### `freeFont(display, fontStruct)`

Releases an XFontStruct.

### `unloadFont(display, fontId)`

Unloads a font loaded via `loadFont`.

## Text Measurement

### `textWidth(fontStruct, text) -> result`

Pixel width via `XTextWidth`. Returns `{ type: :ok, width }`.

## XFontStruct Field Accessors

### `fontStructAscent(fontStruct) -> int`
### `fontStructDescent(fontStruct) -> int`
### `fontStructHeight(fontStruct) -> int`
### `fontStructFid(fontStruct) -> int`

## Higher-Level Helpers

### `fontMetrics(display, name) -> result`

Loads a font and returns its metrics:
`{ type: :ok, fontStruct, fontId, ascent, descent, height }`

### `measureText(display, fontName, text) -> result`

Loads a font, measures the text, and returns:
`{ type: :ok, fontStruct, width, height, ascent, descent }`

The font remains loaded — call `freeFont` when done.

## Font Cache

### `cachedFont(display, fontName) -> result`

Returns a cached XFontStruct, loading on first access.

### `releaseCachedFonts(display)`

Frees all cached fonts for a given display.

### `releaseAllCachedFonts()`

Frees every cached font entry globally.

## Well-Known Font Names

| Constant          | Value                                                  |
|-------------------|--------------------------------------------------------|
| `FONT_FIXED`      | `'fixed'`                                             |
| `FONT_CURSOR`     | `'cursor'`                                            |
| `FONT_6X13`       | `'6x13'`                                              |
| `FONT_9X15`       | `'9x15'`                                              |
| `FONT_10X20`      | `'10x20'`                                             |
| `FONT_COURIER_14` | `'-*-courier-medium-r-normal--14-*-...-iso8859-1'`    |
| `FONT_HELV_12`    | `'-*-helvetica-medium-r-normal--12-*-...-iso8859-1'`  |
| `FONT_TIMES_14`   | `'-*-times-medium-r-normal--14-*-...-iso8859-1'`      |

### `buildXLFD(spec) -> string`

Constructs an X Logical Font Description from a spec:

```oak
xlfd := lf.buildXLFD({
    family: 'courier'
    weight: 'bold'
    slant: 'r'
    pixelSize: 14
    registry: 'iso8859'
    encoding: '1'
})
// -> '-*-courier-bold-r-*-*-14-*-*-*-*-*-iso8859-1'
```

## Example

```oak
lf := import('linux-fonts')

// Load a font and measure text
metrics := lf.fontMetrics(display, '9x15')
if metrics.type = :ok -> {
    w := lf.textWidth(metrics.fontStruct, 'Hello, world!')
    // w.width is the pixel width
    lf.freeFont(display, metrics.fontStruct)
}

// Using the cache
cached := lf.cachedFont(display, 'fixed')
if cached.type = :ok -> {
    lf.setFont(display, gc, lf.fontStructFid(cached.fontStruct))
}
lf.releaseCachedFonts(display)
```
