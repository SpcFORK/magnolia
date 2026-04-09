# Audio Raw PCM Helpers (audio-raw)

## Overview

`audio-raw` provides headerless raw PCM encoding and parsing helpers for `libaudio`. Raw PCM is the simplest audio representation: just sample bytes with no header or metadata. All format parameters must be specified explicitly.

This is useful for inter-process piping, hardware I/O, embedding audio data in custom containers, and working with tools like `ffmpeg -f s16le`.

This module is typically used through `libaudio`'s `rawEncode` and `rawDecode` functions but can be imported directly.

## Import

```oak
rawLib := import('audio-raw')
// or destructure
{ rawEncode: rawEncode, rawDecode: rawDecode } := import('audio-raw')
```

## Functions

### `rawEncode(samples, opts?)`

Encodes a list of normalised floating-point samples (`[-1, 1]`) into a raw PCM binary string with no header.

| Option     | Default | Notes                                            |
|------------|---------|--------------------------------------------------|
| `bitDepth` | 16      | Supported: 8, 16, 32                            |
| `signed`   | true    | Signed vs unsigned (only affects 8-bit)          |
| `endian`   | `:le`   | `:le` (little-endian) or `:be` (big-endian)     |

```oak
// 16-bit little-endian signed (CD-quality raw)
data := rawEncode(samples, { bitDepth: 16, endian: :le })

// 8-bit unsigned (compatible with many simple audio tools)
data := rawEncode(samples, { bitDepth: 8, signed: false })

// 32-bit big-endian signed
data := rawEncode(samples, { bitDepth: 32, endian: :be })
```

### `rawDecode(data, opts?)`

Decodes a raw PCM binary string into normalised float samples. Returns `{ samples }`.

Options are the same as `rawEncode`:

| Option     | Default | Notes                                            |
|------------|---------|--------------------------------------------------|
| `bitDepth` | 16      | Supported: 8, 16, 32                            |
| `signed`   | true    | Signed vs unsigned (only affects 8-bit)          |
| `endian`   | `:le`   | `:le` (little-endian) or `:be` (big-endian)     |

```oak
r := rawDecode(data, { bitDepth: 16, endian: :le })
r.samples  // => list of normalised floats
```

### `pbatchRawEncode(specs)`

Encodes multiple sample lists in parallel. Each spec is `{ samples, opts? }`.

### `pbatchRawDecode(specs)`

Decodes multiple raw PCM byte strings in parallel. Each spec is `{ data, opts? }`.

## Usage with External Tools

Raw PCM is commonly piped to/from external audio tools:

```bash
# Generate raw PCM with Magnolia, play with ffplay
magnolia gen.oak | ffplay -f s16le -ar 44100 -ac 1 -

# Record raw PCM with ffmpeg, process with Magnolia
ffmpeg -f alsa -i default -f s16le -ar 44100 -ac 1 pipe:1 | magnolia process.oak
```

## Format Notes

- No header is written or expected — the caller must track sample rate, channels, bit depth, and endianness externally.
- 8-bit audio can be signed (`[-128, 127]`) or unsigned (`[0, 255]`). 16-bit and 32-bit are always signed.
- Endianness only matters for 16-bit and 32-bit. 8-bit samples are single bytes.
