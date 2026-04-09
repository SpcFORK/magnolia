# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\compression-huffman.oak`

- `Magic` · `'HUF1:'`
- `HeaderSep` · `'|'`
- `PairSep` · `','`
- `ValueSep` · `':'`
### `_freqPairs(input)`

### `_nodeKey(node)`

### `_buildTree(freqPairs)`

### `_buildCodes(tree)`

### `_validBits?(bits)`

### `_decodeBits(tree, bits)`

> returns `:string`

### `_encodePair(pair)`

### `_decodePairs(header)`

> returns `:list`

### `compress(input)`

### `decompress(encoded)`

> returns `:atom`

### `compressed?(encoded)`

