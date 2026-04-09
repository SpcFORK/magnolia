# gui-dialogs — Win32 Common Dialogs

`import('gui-dialogs')` provides Win32 common dialog wrappers for file open/save, folder picking, color selection, and font chooser.

## Quick Start

```oak
dlg := import('gui-dialogs')

// Open file
result := dlg.openFileDialog({
    title: 'Open Image'
    filter: 'Images\0*.png;*.jpg\0All Files\0*.*\0'
})
if result.type = :ok -> println('Selected: ' + result.path)

// Save file
result := dlg.saveFileDialog({ title: 'Save As', defaultExt: 'txt' })

// Pick a color
result := dlg.chooseColor({ initialColor: [255, 0, 0] })
if result.type = :ok -> println(result.color)

// Pick a folder
result := dlg.pickFolder({ title: 'Choose output folder' })

// Choose a font
result := dlg.chooseFont({})
if result.type = :ok -> println(result.name + ' ' + string(result.size))
```

## API Reference

### `openFileDialog(options)`

Displays a file open dialog. Returns `{type: :ok, path}` or `{type: :cancel}`.

**Options:**
- `title` — dialog title
- `filter` — file type filter string (null-separated pairs)
- `initialDir` — starting directory

### `saveFileDialog(options)`

Displays a file save dialog. Returns `{type: :ok, path}` or `{type: :cancel}`.

**Options:**
- `title` — dialog title
- `filter` — file type filter string
- `defaultExt` — default file extension
- `initialDir` — starting directory

### `chooseColor(options)`

Displays a color chooser dialog. Returns `{type: :ok, color}` or `{type: :cancel}`.

**Options:**
- `initialColor` — `[r, g, b]` starting color
- `fullOpen` — show custom color controls immediately

### `pickFolder(options)`

Displays a folder picker dialog. Returns `{type: :ok, path}` or `{type: :cancel}`.

**Options:**
- `title` — dialog title

### `chooseFont(options)`

Displays a font chooser dialog. Returns `{type: :ok, name, size, weight, italic, color}` or `{type: :cancel}`.

## Constants

### File Dialog Flags

| Constant | Value |
|----------|-------|
| `OFN_PATHMUSTEXIST` | 2048 |
| `OFN_FILEMUSTEXIST` | 4096 |
| `OFN_OVERWRITEPROMPT` | 2 |
| `OFN_EXPLORER` | 524288 |

### Color Dialog Flags

| Constant | Value |
|----------|-------|
| `CC_RGBINIT` | 1 |
| `CC_FULLOPEN` | 2 |

### Font Dialog Flags

| Constant | Value |
|----------|-------|
| `CF_SCREENFONTS` | 1 |
| `CF_EFFECTS` | 256 |
| `CF_INITTOLOGFONTSTRUCT` | 64 |
