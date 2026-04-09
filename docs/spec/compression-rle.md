# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\compression-rle.oak`

- `Marker` · `''`
- `RunPrefix` · `'#'`
- `RunSep` · `':'`
### `_int?(n)`

### `_digitCount(n)`

### `_literalLen(c, count)`

### `_runLen(count)`

> returns `:int`

### `_appendLiteral(out, c)`

### `_appendRun(out, c, count)`

### `_flush(out, c, count)`

### `_readDigits(s, start)`

### `_repeatChar(c, count)`

### `compress(input)`

> returns `:string`

### `decompress(encoded)`

### `compressed?(encoded)`

