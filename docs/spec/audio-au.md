# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-au.oak`

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

### `_u32BE(n)`

> returns `:list`

### `_readU8(s, off)`

### `_readI8(s, off)`

### `_readU32BE(s, off)`

### `_readI16BE(s, off)`

### `_readI32BE(s, off)`

- `_AU_FMT_PCM8` · `2`
- `_AU_FMT_PCM16` · `3`
- `_AU_FMT_PCM32` · `5`
### `_auSampleBytes(sample, bitDepth)`

### `_auDataChunk(samples, bitDepth)`

### `au(samples, sampleRate, channels, bitDepth)`

### `_auReadSample(data, off, encoding)`

### `_auBytesPerSample(encoding)`

> returns `:int`

### `_auEncodingBitDepth(encoding)`

> returns `:int`

### `_auReadSamples(data, dataOff, totalSamples, bytesPerSample, encoding)`

### `parseAu(data)`

> returns `?`

### `pbatchAu(specs)`

### `pbatchParseAu(dataList)`

