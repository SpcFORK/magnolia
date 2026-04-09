# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-dialogs.oak`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_ptrSize` — constant
### `_writePtr(address, value)`

### `_readUtf16Str(addr)`

### `_zeros(n)`

- `OFN_SIZE` — constant
- `OFN_FILEBUFSIZE` · `520`
- `OFN_PATHMUSTEXIST` · `2048`
- `OFN_FILEMUSTEXIST` · `4096`
- `OFN_OVERWRITEPROMPT` · `2`
- `OFN_NOCHANGEDIR` · `8`
- `OFN_EXPLORER` · `524288`
### `_buildFilter(filter)`

### `_fileDialog(hwnd, title, filter, flags, apiName)`

> returns `:object`

### `openFileDialog(options)`

### `saveFileDialog(options)`

- `CC_SIZE` — constant
- `CC_RGBINIT` · `1`
- `CC_FULLOPEN` · `2`
### `chooseColor(options)`

> returns `:object`

- `BI_SIZE` — constant
- `BIF_RETURNONLYFSDIRS` · `1`
- `BIF_NEWDIALOGSTYLE` · `64`
### `pickFolder(options)`

> returns `:object`

- `CF_SIZE` — constant
- `LOGFONT_SIZE` · `92`
- `CF_SCREENFONTS` · `1`
- `CF_EFFECTS` · `256`
- `CF_INITTOLOGFONTSTRUCT` · `64`
### `chooseFont(options)`

> returns `:object`

