# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\compression-lzw.oak`

- `Magic` · `'LZW1:'`
- `HeaderSep` · `'|'`
- `ListSep` · `','`
### `_alphabet(input)`

### `_dict(alphabet)`

### `_validInts?(xs)`

### `_encodeAlphabet(alphabet)`

### `_decodeAlphabet(header)`

> returns `:list`

### `_encodeCodes(codes)`

### `_decodeCodes(header)`

> returns `:list`

### `compress(input)`

> returns `:atom`

### `decompress(encoded)`

> returns `:atom`

### `compressed?(encoded)`

