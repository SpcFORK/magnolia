# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\windows-fonts.oak`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `wstr(s)`

### `gdi32(symbol, args...)`

### `user32(symbol, args...)`

- `FW_DONTCARE` · `0`
- `FW_THIN` · `100`
- `FW_EXTRALIGHT` · `200`
- `FW_LIGHT` · `300`
- `FW_NORMAL` · `400`
- `FW_MEDIUM` · `500`
- `FW_SEMIBOLD` · `600`
- `FW_BOLD` · `700`
- `FW_EXTRABOLD` · `800`
- `FW_HEAVY` · `900`
- `ANSI_CHARSET` · `0`
- `DEFAULT_CHARSET` · `1`
- `SYMBOL_CHARSET` · `2`
- `SHIFTJIS_CHARSET` · `128`
- `HANGUL_CHARSET` · `129`
- `GB2312_CHARSET` · `134`
- `CHINESEBIG5_CHARSET` · `136`
- `OEM_CHARSET` · `255`
- `JOHAB_CHARSET` · `130`
- `HEBREW_CHARSET` · `177`
- `ARABIC_CHARSET` · `178`
- `GREEK_CHARSET` · `161`
- `TURKISH_CHARSET` · `162`
- `VIETNAMESE_CHARSET` · `163`
- `THAI_CHARSET` · `222`
- `EASTEUROPE_CHARSET` · `238`
- `RUSSIAN_CHARSET` · `204`
- `BALTIC_CHARSET` · `186`
- `OUT_DEFAULT_PRECIS` · `0`
- `OUT_STRING_PRECIS` · `1`
- `OUT_STROKE_PRECIS` · `3`
- `OUT_TT_PRECIS` · `4`
- `OUT_DEVICE_PRECIS` · `5`
- `OUT_TT_ONLY_PRECIS` · `7`
- `OUT_OUTLINE_PRECIS` · `8`
- `CLIP_DEFAULT_PRECIS` · `0`
- `DEFAULT_QUALITY` · `0`
- `DRAFT_QUALITY` · `1`
- `PROOF_QUALITY` · `2`
- `NONANTIALIASED_QUALITY` · `3`
- `ANTIALIASED_QUALITY` · `4`
- `CLEARTYPE_QUALITY` · `5`
- `DEFAULT_PITCH` · `0`
- `FIXED_PITCH` · `1`
- `VARIABLE_PITCH` · `2`
- `FF_DONTCARE` · `0`
- `FF_ROMAN` · `16`
- `FF_SWISS` · `32`
- `FF_MODERN` · `48`
- `FF_SCRIPT` · `64`
- `FF_DECORATIVE` · `80`
- `OEM_FIXED_FONT` · `10`
- `ANSI_FIXED_FONT` · `11`
- `ANSI_VAR_FONT` · `12`
- `SYSTEM_FONT` · `13`
- `DEVICE_DEFAULT_FONT` · `14`
- `DEFAULT_GUI_FONT` · `17`
- `TA_NOUPDATECP` · `0`
- `TA_UPDATECP` · `1`
- `TA_LEFT` · `0`
- `TA_RIGHT` · `2`
- `TA_CENTER` · `6`
- `TA_TOP` · `0`
- `TA_BOTTOM` · `8`
- `TA_BASELINE` · `24`
- `TM_HEIGHT_OFF` · `0`
- `TM_ASCENT_OFF` · `4`
- `TM_DESCENT_OFF` · `8`
- `TM_INTERNAL_LEADING_OFF` · `12`
- `TM_EXTERNAL_LEADING_OFF` · `16`
- `TM_AVE_CHAR_WIDTH_OFF` · `20`
- `TM_MAX_CHAR_WIDTH_OFF` · `24`
- `TM_WEIGHT_OFF` · `28`
### `createFontEx(spec)`

> returns `:object`

### `deleteFont(handle)`

### `selectFont(hdc, fontHandle)`

### `getStockFont(stockIndex)`

> returns `:object`

### `getDefaultGuiFont()`

### `setTextAlign(hdc, flags)`

### `getTextMetrics(hdc)`

> returns `:object`

### `getTextExtent(hdc, text)`

> returns `:object`

- `_fontCacheMap` · `{}`
### `_fontCacheKey(spec)`

### `cachedFont(spec)`

### `releaseCachedFonts()`

> returns `:int`

### `removeCachedFont(spec)`

> returns `:int`

### `measureText(hwnd, spec, text)`

> returns `:object`

### `fontLineHeight(hdc)`

> returns `:object`

