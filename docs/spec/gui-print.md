# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-print.oak`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `PD_ALLPAGES` · `0`
- `PD_SELECTION` · `1`
- `PD_PAGENUMS` · `2`
- `PD_NOSELECTION` · `4`
- `PD_NOPAGENUMS` · `8`
- `PD_COLLATE` · `16`
- `PD_PRINTTOFILE` · `32`
- `PD_PRINTSETUP` · `64`
- `PD_NOWARNING` · `128`
- `PD_RETURNDC` · `256`
- `PD_RETURNIC` · `512`
- `PD_RETURNDEFAULT` · `1024`
- `PD_SHOWHELP` · `2048`
- `PD_USEDEVMODECOPIES` · `262144`
- `PD_DISABLEPRINTTOFILE` · `524288`
- `PD_HIDEPRINTTOFILE` · `1048576`
- `PD_CURRENTPAGE` · `4194304`
- `DI_APPBANDING` · `1`
- `_PRINTDLGW_SIZE` — constant
- `_PD_OFF_HWNDOWNER` · `8`
- `_PD_OFF_HDEVMODE` · `16`
- `_PD_OFF_HDEVNAMES` · `24`
- `_PD_OFF_HDC` · `32`
- `_PD_OFF_FLAGS` · `40`
- `_PD_OFF_FROMPAGE` · `44`
- `_PD_OFF_TOPAGE` · `46`
- `_PD_OFF_MINPAGE` · `48`
- `_PD_OFF_MAXPAGE` · `50`
- `_PD_OFF_NCOPIES` · `52`
- `_DOCINFOW_SIZE` — constant
- `_DI_OFF_DOCNAME` — constant
- `_DI_OFF_OUTPUT` — constant
- `_DI_OFF_DATATYPE` — constant
- `_DI_OFF_FWTYPE` — constant
### `_writeU16(base, value)`

### `_readU16(base)`

### `showPrintDialog(options)`

> returns `:object`

### `startDoc(hDC, docName, outputFile)`

> returns `:object`

### `startPage(hDC)`

> returns `:bool`

### `endPage(hDC)`

> returns `:bool`

### `endDoc(hDC)`

> returns `:bool`

### `abortDoc(hDC)`

### `deleteDC(hDC)`

### `printTextOut(hDC, x, y, text)`

### `printMoveTo(hDC, x, y)`

### `printLineTo(hDC, x, y)`

### `printRectangle(hDC, left, top, right, bottom)`

### `printEllipse(hDC, left, top, right, bottom)`

### `printSetFont(hDC, height, weight, italic, fontName)`

### `printSetTextColor(hDC, r, g, b)`

### `printSetBkMode(hDC, mode)`

### `printSetPen(hDC, style, width, r, g, b)`

### `printDeleteObject(hObj)`

- `DEVCAP_HORZRES` · `8`
- `DEVCAP_VERTRES` · `10`
- `DEVCAP_LOGPIXELSX` · `88`
- `DEVCAP_LOGPIXELSY` · `90`
- `DEVCAP_PHYSICALWIDTH` · `110`
- `DEVCAP_PHYSICALHEIGHT` · `111`
- `DEVCAP_PHYSICALOFFSETX` · `112`
- `DEVCAP_PHYSICALOFFSETY` · `113`
### `getDeviceCaps(hDC, index)`

### `getPrinterPageSize(hDC)`

> returns `:object`

### `createPreviewDC(width, height)`

> returns `:object`

### `destroyPreviewDC(preview)`

### `printToFile(outputPath, docName, renderFn)`

> returns `:object`

### `printDocument(options, renderPageFn)`

> returns `:object`

