# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-fonts.oak`

- `guiThread` · `import(...)`
### `_asBool(v)`

> returns `:bool`

- `_platformCache` — constant
### `_platformId()`

### `isWindows?()`

### `isLinux?()`

- `_winMod` · `?`
- `_linuxMod` · `?`
### `_win()`

### `_lnx()`

- `FW_THIN` · `100`
- `FW_EXTRALIGHT` · `200`
- `FW_LIGHT` · `300`
- `FW_NORMAL` · `400`
- `FW_MEDIUM` · `500`
- `FW_SEMIBOLD` · `600`
- `FW_BOLD` · `700`
- `FW_EXTRABOLD` · `800`
- `FW_HEAVY` · `900`
### `defaultFontSpec()`

> returns `:object`

### `createFont(spec)`

### `deleteFont(fontResult)`

> returns `:int`

### `_webFontString(spec)`

> returns `?`

### `fontKey(spec)`

> returns `?`

### `cachedFont(windowOrDisplay, spec)`

### `releaseCachedFonts(windowOrDisplay)`

### `measureText(windowOrDisplay, spec, text)`

### `selectFont(args)`

### `getTextMetrics(hdc)`

### `fontLineHeight(hdcOrFontStruct)`

### `buildXLFD(spec)`

### `webFontString(spec)`

