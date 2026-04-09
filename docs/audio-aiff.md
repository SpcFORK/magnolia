# Audio AIFF Helpers (audio-aiff)

## Overview

`audio-aiff` provides AIFF (Audio Interchange File Format) encoding and parsing helpers for `libaudio`. AIFF is Apple's standard uncompressed audio format, using big-endian byte order and signed integer PCM. It supports 8-bit, 16-bit, and 32-bit sample depths.

This module is typically used through `libaudio`'s `aiff` and `parseAiff` functions but can be imported directly.

## Import

```oak
aiffLib := import('audio-aiff')
// or destructure
{ aiff: aiff, parseAiff: parseAiff } := import('audio-aiff')
```

## Functions

### `aiff(samples, sampleRate?, channels?, bitDepth?)`

Encodes a list of normalised floating-point samples (`[-1, 1]`) as a binary AIFF file and returns the full file string.

| Parameter    | Default | Notes                          |
|--------------|---------|--------------------------------|
| `sampleRate` | 44100   | Hz                             |
| `channels`   | 1       | Mono. Pass 2 for stereo.       |
| `bitDepth`   | 16      | Supported: 8, 16, 32          |

```oak
samples := sine(440, 44100, 44100)
data := aiff(samples, 44100, 1, 16)
writeFile('out.aiff', data)
```

For stereo, interleave left and right samples: `[L0, R0, L1, R1, ...]`.

### `parseAiff(data)`

Parses an AIFF binary string and returns a metadata object, or `?` on error.

```oak
r := parseAiff(fileBytes)
r.sampleRate  // => 44100
r.channels    // => 1
r.bitDepth    // => 16
r.samples     // => list of normalised floats
```

### `pbatchAiff(specs)`

Encodes multiple sample lists into AIFF byte strings in parallel. Each spec is `{ samples, sampleRate?, channels?, bitDepth? }`.

```oak
results := pbatchAiff([
    { samples: s1 }
    { samples: s2, bitDepth: 32 }
])
```

### `pbatchParseAiff(dataList)`

Decodes multiple AIFF byte strings in parallel.

```oak
results := pbatchParseAiff([aiffData1, aiffData2])
```

## Format Details

The output conforms to standard AIFF (big-endian, signed integer PCM). The file structure is:

- **FORM** container chunk with type `AIFF`
- **COMM** chunk: channels, frame count, bit depth, sample rate (80-bit extended float)
- **SSND** chunk: sound data with offset and block size fields

| Bit depth | Encoding          |
|-----------|-------------------|
| 8         | Signed 8-bit      |
| 16        | Signed 16-bit BE  |
| 32        | Signed 32-bit BE  |

Samples are clipped to `[-1, 1]` before encoding. Note that unlike WAV, AIFF uses **signed** encoding for 8-bit samples (WAV uses unsigned).
