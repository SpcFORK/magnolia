# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\dataprot.oak`

### `_byte?(n)`

> returns `:bool`

### `_byteList(data)`

### `_bitValue(v)`

> returns `:int`

### `_bitVector(data)`

### `_bitMatrix(matrix)`

### `_popcount8(n)`

### `_targetParity(odd)`

> returns `:int`

### `_crc16Byte(crc, b, poly)`

### `_crc32Byte(crc, b, poly)`

### `_rowParity(row, bits)`

### `_bitColumn(rows, idx)`

### `_sameBits(a, b)`

> returns `:bool`

### `_flipBit(bits, idx)`

### `_formatBits(bits, template)`

### `parity(data)`

> returns `:atom`

### `parityBit(data, odd)`

> returns `:atom`

### `parityValid?(data, checkBit, odd)`

> returns `:atom`

### `xorChecksum(data)`

> returns `:atom`

### `sumChecksum8(data)`

> returns `:atom`

### `sumChecksum16(data)`

> returns `:atom`

### `sumChecksum32(data)`

> returns `:atom`

### `crc16Ccitt(data, seed, poly)`

> returns `:atom`

### `crc32(data, seed, poly, finalXor)`

> returns `:atom`

### `hammingDistance(a, b)`

> returns `:atom`

### `ldpcSyndrome(word, parityMatrix)`

> returns `:atom`

### `ldpcValid?(word, parityMatrix)`

> returns `:atom`

### `ldpcCheck(word, parityMatrix)`

> returns `:atom`

### `ldpcCandidates(word, parityMatrix)`

> returns `:atom`

### `ldpcCorrect(word, parityMatrix)`

> returns `:atom`

### `pbatchCrc32(payloads)`

### `pbatchLdpcCheck(words, parityMatrix)`

