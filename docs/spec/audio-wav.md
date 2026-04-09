# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-wav.oak`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_f32ToU8(s)`

### `_u8ToF32(s)`

### `_byteStr(blist)`

### `_u16LE(n)`

> returns `:list`

### `_i16LE(n)`

### `_u32LE(n)`

> returns `:list`

### `_readU8(s, off)`

### `_readU16LE(s, off)`

### `_readI16LE(s, off)`

### `_readU32LE(s, off)`

### `_readI32LE(s, off)`

### `_wavSampleBytes(sample, bitDepth)`

### `_wavDataChunk(samples, bitDepth)`

### `_wavHeader(fileLen, sampleRate, channels, bitDepth, byteRate, blockAlign, dataLen)`

> returns `:string`

### `wav(samples, sampleRate, channels, bitDepth)`

### `_wavHeaderValid(data)`

> returns `:bool`

### `_wavFindDataChunk(data, off)`

> returns `?`

### `_wavReadSample(data, off, bitDepth)`

### `_wavReadSamples(data, dataOff, totalSamples, bytesPerSample, bitDepth)`

### `pbatchWav(specs)`

### `pbatchParseWav(dataList)`

### `parseWav(data)`

> returns `?`

