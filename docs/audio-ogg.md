# Audio OGG Helpers (audio-ogg)

## Overview

`audio-ogg` provides OGG container encoding and parsing helpers for `libaudio`. It implements the OGG bitstream framing format (RFC 3533) with uncompressed signed 16-bit little-endian PCM payload.

While real-world OGG files typically carry Vorbis or Opus compressed audio, this pure-Oak implementation uses raw PCM so that no external codec library is needed. Files produced by this module can be round-tripped through `ogg()`/`parseOgg()` and converted to standard formats via the other libaudio modules.

This module is typically used through `libaudio`'s `ogg` and `parseOgg` functions but can be imported directly.

## Import

```oak
oggLib := import('audio-ogg')
// or destructure
{ ogg: ogg, parseOgg: parseOgg } := import('audio-ogg')
```

## Functions

### `ogg(samples, sampleRate?, channels?, bitDepth?)`

Encodes a list of normalised floating-point samples (`[-1, 1]`) into an OGG container with PCM payload and returns the full file string.

| Parameter    | Default | Notes                          |
|--------------|---------|--------------------------------|
| `sampleRate` | 44100   | Hz                             |
| `channels`   | 1       | Mono. Pass 2 for stereo.       |
| `bitDepth`   | 16      | Currently only 16-bit supported |

```oak
samples := sine(440, 44100, 44100)
data := ogg(samples, 44100, 1, 16)
writeFile('out.ogg', data)
```

For stereo, interleave left and right samples: `[L0, R0, L1, R1, ...]`.

### `parseOgg(data)`

Parses an OGG binary string with PCM payload and returns a metadata object, or `?` on error.

```oak
r := parseOgg(fileBytes)
r.sampleRate  // => 44100
r.channels    // => 1
r.bitDepth    // => 16
r.samples     // => list of normalised floats
```

Only OGG files produced by this module (codec id `'PCM     '`) are supported. Standard Vorbis/Opus OGG files will return `?`.

### `pbatchOgg(specs)`

Encodes multiple sample lists into OGG byte strings in parallel. Each spec is `{ samples, sampleRate?, channels?, bitDepth? }`.

### `pbatchParseOgg(dataList)`

Decodes multiple OGG byte strings in parallel.

## Format Details

### OGG Page Structure

Each OGG page has the following header:

| Offset | Size | Field              | Description                            |
|--------|------|--------------------|----------------------------------------|
| 0      | 4    | Capture pattern    | `OggS`                                |
| 4      | 1    | Version            | 0                                      |
| 5      | 1    | Header type        | 0x02=BOS, 0x04=EOS, 0x00=continuation |
| 6      | 8    | Granule position   | Absolute sample position (LE i64)      |
| 14     | 4    | Serial number      | Stream serial (LE u32)                 |
| 18     | 4    | Page sequence      | Page counter (LE u32)                  |
| 22     | 4    | CRC-32             | Checksum (LE u32)                      |
| 26     | 1    | Segment count      | Number of segment table entries        |
| 27     | N    | Segment table      | N bytes, each 0-255                    |
| 27+N   | M    | Segment data       | Concatenated segment payloads          |

### Codec Identification Header (BOS page payload)

| Offset | Size | Field       | Description              |
|--------|------|-------------|--------------------------|
| 0      | 8    | Codec ID    | `PCM     ` (padded)      |
| 8      | 2    | Channels    | LE u16                   |
| 10     | 4    | Sample rate | LE u32                   |
| 14     | 2    | Bit depth   | LE u16                   |

### Data Pages

Audio samples are encoded as signed 16-bit little-endian PCM and split across multiple OGG pages, each containing up to 32 000 samples. The last data page has the EOS header flag set.

### CRC-32

OGG uses CRC-32 with polynomial `0x04C11DB7`. The checksum is computed over the entire page with the CRC field zeroed, then patched into bytes 22-25.

## Limitations

- Only 16-bit PCM is supported (matching the most common use case)
- This is **not** Vorbis/Opus — files are uncompressed and will not play in standard media players expecting OGG Vorbis
- For interop with standard OGG players, encode to WAV first and use an external tool to convert to OGG Vorbis
