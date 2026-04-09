# gui-clipboard — Win32 Clipboard Access

`import('gui-clipboard')` provides Win32 clipboard read/write for text data using `CF_UNICODETEXT` format with UTF-16LE encoding.

## Quick Start

```oak
clip := import('gui-clipboard')

// Write text to clipboard
clip.clipboardSetText('Hello, world!')

// Check and read
if clip.clipboardHasText() -> {
    text := clip.clipboardGetText()
    println(text)
}
```

## API Reference

### `clipboardGetText()`

Reads Unicode text from the clipboard. Returns the string or `?` on failure.

### `clipboardSetText(text)`

Writes Unicode text to the clipboard. Returns `true` on success.

### `clipboardHasText()`

Returns `true` if the clipboard currently contains text data.

## Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `CF_UNICODETEXT` | 13 | Clipboard format for Unicode text |
| `GMEM_MOVEABLE` | 2 | Global memory allocation flag |
