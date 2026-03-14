# LZW Compression Library (compression-lzw)

## Overview

`compression-lzw` implements dictionary-based LZW compression for Oak strings.

It emits a text-safe payload with:
- An alphabet header (hex codepoints)
- A comma-separated integer code stream

Use this codec when data contains repeated substrings or phrases.

## Import

```oak
lzw := import('compression-lzw')
{ compress: compress, decompress: decompress } := import('compression-lzw')
```

## Constants

```oak
Magic := 'LZW1:'
```

## Packet Format

Compressed payload format:

```text
LZW1:<hex-char,hex-char,...>|<code,code,code,...>
```

- Left side is the initial dictionary alphabet.
- Right side is the emitted LZW code stream.

## API

### `compress(input)`

Compresses an input string with LZW dictionary expansion.

Returns `:error` if the dictionary would exceed 65535 entries.

```oak
lzw := import('compression-lzw')

packed := lzw.compress('TOBEORNOTTOBEORTOBEORNOT')
```

### `decompress(encoded)`

Decompresses payloads produced by `compress`.

Returns `:error` for malformed packets, invalid headers/codes, or unsupported dictionary growth.

```oak
lzw := import('compression-lzw')

source := 'TOBEORNOTTOBEORTOBEORNOT'
packed := lzw.compress(source)
restored := lzw.decompress(packed)

restored = source // => true
```

### `compressed?(encoded)`

Returns true if payload starts with `Magic`.

```oak
lzw := import('compression-lzw')

lzw.compressed?('plain') // => false
lzw.compressed?(lzw.compress('aaaaabbbb')) // => true
```

## Notes

- Effective on repetitive phrase patterns.
- Can be larger than input for short or highly random strings.
- Output is text-safe and deterministic for the same input.
