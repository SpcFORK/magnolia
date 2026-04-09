# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-ogg.oak`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_byteStr(blist)`

### `_u16LE(n)`

> returns `:list`

### `_u32LE(n)`

> returns `:list`

### `_i16LE(n)`

### `_readU8(s, off)`

### `_readU16LE(s, off)`

### `_readU32LE(s, off)`

### `_readI16LE(s, off)`

### `_pow2(n)`

### `_xor32(a, b)`

### `_crc32OggTable()`

- `_crcTable` · `_crc32OggTable(...)`
### `_crc32Ogg(data)`

### `_granuleBytes(n)`

### `_segmentTable(payloadLen)`

### `_oggPage(headerType, granule, serial, pageSeq, payload)`

### `_oggPcmIdHeader(sampleRate, channels, bitDepth)`

> returns `:string`

- `_maxPageSamples` · `32000`
### `_encodeSamples(samples, bitDepth)`

### `ogg(samples, sampleRate, channels, bitDepth)`

### `_readOggPageHeader(data, off)`

> returns `:object`

### `_readPcmSamples(data, off, nBytes)`

### `parseOgg(data)`

> returns `:object`

### `pbatchOgg(specs)`

### `pbatchParseOgg(dataList)`

