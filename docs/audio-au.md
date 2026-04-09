# Audio AU Helpers (audio-au)

## Overview

`audio-au` provides AU (Sun/NeXT) audio format encoding and parsing helpers for `libaudio`. AU is a simple big-endian audio format originating from Sun Microsystems. It supports 8-bit, 16-bit, and 32-bit signed linear PCM.

This module is typically used through `libaudio`'s `au` and `parseAu` functions but can be imported directly.

## Import

```oak
auLib := import('audio-au')
// or destructure
{ au: au, parseAu: parseAu } := import('audio-au')
```

## Functions

### `au(samples, sampleRate?, channels?, bitDepth?)`

Encodes a list of normalised floating-point samples (`[-1, 1]`) as a binary AU file and returns the full file string.

| Parameter    | Default | Notes                          |
|--------------|---------|--------------------------------|
| `sampleRate` | 44100   | Hz                             |
| `channels`   | 1       | Mono. Pass 2 for stereo.       |
| `bitDepth`   | 16      | Supported: 8, 16, 32          |

```oak
samples := sine(440, 44100, 44100)
data := au(samples, 44100, 1, 16)
writeFile('out.au', data)
```

For stereo, interleave left and right samples: `[L0, R0, L1, R1, ...]`.

### `parseAu(data)`

Parses an AU binary string and returns a metadata object, or `?` on error.

```oak
r := parseAu(fileBytes)
r.sampleRate  // => 44100
r.channels    // => 1
r.bitDepth    // => 16
r.samples     // => list of normalised floats
```

Only linear PCM encodings are supported (AU encoding IDs 2, 3, 5). Compressed formats (mu-law, A-law, etc.) return `?`.

### `pbatchAu(specs)`

Encodes multiple sample lists into AU byte strings in parallel. Each spec is `{ samples, sampleRate?, channels?, bitDepth? }`.

### `pbatchParseAu(dataList)`

Decodes multiple AU byte strings in parallel.

## Format Details

The AU format has a simple header (minimum 28 bytes):

| Offset | Size | Field          | Description                      |
|--------|------|----------------|----------------------------------|
| 0      | 4    | Magic          | `.snd` (0x2e736e64)             |
| 4      | 4    | Data offset    | Byte offset to sample data       |
| 8      | 4    | Data size      | Size of sample data in bytes     |
| 12     | 4    | Encoding       | 2=8-bit, 3=16-bit, 5=32-bit    |
| 16     | 4    | Sample rate    | Samples per second               |
| 20     | 4    | Channels       | Number of channels               |
| 24     | 4    | Annotation     | Padding / annotation (optional)  |

All header fields are big-endian unsigned 32-bit integers. Sample data follows the header. A data size of `0xFFFFFFFF` means "unknown" and the parser reads to end-of-file.

| Bit depth | Encoding           | AU ID |
|-----------|--------------------|-------|
| 8         | Signed 8-bit       | 2     |
| 16        | Signed 16-bit BE   | 3     |
| 32        | Signed 32-bit BE   | 5     |
