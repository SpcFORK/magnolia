# Audio WAV Helpers (audio-wav)

## Overview

`audio-wav` provides WAV file encoding helpers used internally by `libaudio`. It converts normalised floating-point sample lists to standard PCM WAV binary data at configurable sample rates, channel counts, and bit depths.

This module is typically used through `libaudio`'s `wav` function but can be imported directly for low-level WAV generation.

## Import

```oak
wavLib := import('audio-wav')
// or destructure
{ wav: wav } := import('audio-wav')
```

## Functions

### `wav(samples, sampleRate?, channels?, bitDepth?)`

Encodes a list of normalised floating-point samples (`[-1, 1]`) as a binary WAV file and returns the full file string. All optional parameters have sensible defaults.

| Parameter    | Default | Notes                          |
|--------------|---------|--------------------------------|
| `sampleRate` | 44100   | Hz                             |
| `channels`   | 1       | Mono. Pass 2 for stereo.       |
| `bitDepth`   | 16      | Supported: 8, 16, 32          |

```oak
samples := [0.0, 0.5, 1.0, 0.5, 0.0, -0.5, -1.0, -0.5]
data := wav(samples, 44100, 1, 16)
writeFile('out.wav', data)
```

For stereo, interleave left and right samples: `[L0, R0, L1, R1, ...]`.

```oak
stereoSamples := [0.5, -0.5, 0.8, -0.8]
data := wav(stereoSamples, 44100, 2, 16)
```

## Format Details

The output conforms to standard PCM WAV (format tag 1, little-endian). The file begins with the 44-byte RIFF/WAVE header followed by the data chunk. Samples are clipped to [-1, 1] before encoding.

| Bit depth | Encoding        |
|-----------|-----------------|
| 8         | Unsigned byte   |
| 16        | Signed 16-bit   |
| 32        | Signed 32-bit   |
## Parallel Batch Operations

### `pbatchWav(specs)`

Encodes multiple sample lists into WAV byte strings in parallel. Each spec is `{ samples, sampleRate?, channels?, bitDepth? }`.

```oak
pbatchWav([{ samples: s1 }, { samples: s2, bitDepth: 32 }])
```

### `pbatchParseWav(dataList)`

Decodes multiple WAV byte strings in parallel.

```oak
pbatchParseWav([wavData1, wavData2])
// => [{ sampleRate, channels, bitDepth, samples }, ...]
```