# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-aiff.oak`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI8(s)`

> returns `:int`

### `_i8ToF32(s)`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_byteStr(blist)`

### `_u16BE(n)`

> returns `:list`

### `_i16BE(n)`

### `_u32BE(n)`

> returns `:list`

### `_i32BE(n)`

### `_readU8(s, off)`

### `_readI8(s, off)`

### `_readU16BE(s, off)`

### `_readI16BE(s, off)`

### `_readU32BE(s, off)`

### `_readI32BE(s, off)`

### `_encodeExtended80(val)`

### `_decodeExtended80(s, off)`

### `_aiffSampleBytes(sample, bitDepth)`

### `_aiffDataChunk(samples, bitDepth)`

### `aiff(samples, sampleRate, channels, bitDepth)`

> returns `:string`

### `_aiffReadSample(data, off, bitDepth)`

### `_aiffReadSamples(data, dataOff, totalSamples, bytesPerSample, bitDepth)`

### `_aiffFindChunk(data, off, endOff, tag)`

> returns `?`

### `parseAiff(data)`

> returns `?`

### `pbatchAiff(specs)`

### `pbatchParseAiff(dataList)`

