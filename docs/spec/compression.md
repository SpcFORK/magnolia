# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\compression.oak`

- `rle` · `import(...)`
- `huffman` · `import(...)`
- `lzw` · `import(...)`
- `Marker` — constant
- `RunPrefix` — constant
- `RunSep` — constant
- `HuffmanMagic` — constant
- `LZWMagic` — constant
### `rleCompress(input)`

### `rleDecompress(encoded)`

### `rleCompressed?(encoded)`

### `huffmanCompress(input)`

### `huffmanDecompress(encoded)`

### `huffmanCompressed?(encoded)`

### `lzwCompress(input)`

### `lzwDecompress(encoded)`

### `lzwCompressed?(encoded)`

### `compress(input, algorithm)`

### `decompress(encoded, algorithm)`

### `compressed?(encoded, algorithm)`

> returns `:bool`

- `ParMagic` · `'MPAR'`
### `_splitBlocks(input, n)`

> returns `:list`

### `_joinBlocks(blocks)`

### `_encodeParallel(compressedBlocks, algByte)`

### `_decodeParallel(encoded)`

> returns `:list`

### `_algByte(algorithm)`

> returns `:string`

### `_compressorFor(algByte)`

### `_decompressorFor(algByte)`

### `pcompress(input, algorithm, numBlocks)`

> returns `:atom`

### `pdecompress(encoded)`

> returns `:atom`

### `pcompressed?(encoded)`

