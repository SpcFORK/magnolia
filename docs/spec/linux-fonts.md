# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\linux-fonts.oak`

- `sys` · `import(...)`
### `loadQueryFont(display, name)`

> returns `:object`

### `loadFont(display, name)`

> returns `:object`

### `setFont(display, gc, fontId)`

### `freeFont(display, fontStruct)`

### `unloadFont(display, fontId)`

### `textWidth(fontStruct, text)`

> returns `:object`

- `XFONTSTRUCT_FID_OFF` · `8`
- `XFONTSTRUCT_DIRECTION_OFF` · `16`
- `XFONTSTRUCT_ASCENT_OFF` · `88`
- `XFONTSTRUCT_DESCENT_OFF` · `92`
### `fontStructAscent(fontStruct)`

### `fontStructDescent(fontStruct)`

### `fontStructHeight(fontStruct)`

### `fontStructFid(fontStruct)`

### `fontMetrics(display, name)`

> returns `:object`

### `measureText(display, fontName, text)`

> returns `:object`

- `_fontCacheMap` · `{}`
### `_fontCacheKey(displayPtr, fontName)`

### `cachedFont(display, fontName)`

### `releaseCachedFonts(display)`

> returns `:int`

### `releaseAllCachedFonts()`

> returns `:int`

- `FONT_FIXED` · `'fixed'`
- `FONT_CURSOR` · `'cursor'`
- `FONT_6X13` · `'6x13'`
- `FONT_7X13` · `'7x13'`
- `FONT_8X13` · `'8x13'`
- `FONT_9X15` · `'9x15'`
- `FONT_10X20` · `'10x20'`
- `FONT_COURIER_14` · `'-*-courier-medium-r-normal--14-*-*-*-*-*-iso8859-1'`
- `FONT_COURIER_B14` · `'-*-courier-bold-r-normal--14-*-*-*-*-*-iso8859-1'`
- `FONT_HELV_12` · `'-*-helvetica-medium-r-normal--12-*-*-*-*-*-iso8859-1'`
- `FONT_HELV_B14` · `'-*-helvetica-bold-r-normal--14-*-*-*-*-*-iso8859-1'`
- `FONT_TIMES_14` · `'-*-times-medium-r-normal--14-*-*-*-*-*-iso8859-1'`
- `FONT_TIMES_B14` · `'-*-times-bold-r-normal--14-*-*-*-*-*-iso8859-1'`
### `buildXLFD(spec)`

> returns `:string`

