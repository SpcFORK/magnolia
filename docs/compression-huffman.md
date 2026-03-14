# Huffman Compression Library (compression-huffman)

## Overview

`compression-huffman` implements Huffman coding for Oak strings.

It computes character frequencies, builds a binary coding tree, and emits:
- A text-safe serialized frequency header
- A bitstring payload (`0` / `1`)

Use this codec when character distribution is skewed.

## Import

```oak
huffman := import('compression-huffman')
{ compress: compress, decompress: decompress } := import('compression-huffman')
```

## Constants

```oak
Magic := 'HUF1:'
```

## Packet Format

Compressed payload format:

```text
HUF1:<hex-char:count,hex-char:count,...>|<bitstring>
```

- `<hex-char>` is the hex codepoint of the character.
- `<count>` is decimal occurrence count.
- `<bitstring>` contains only `0` and `1`.

## API

### `compress(input)`

Compresses an input string using a Huffman tree built from input frequencies.

```oak
huffman := import('compression-huffman')

packed := huffman.compress('BANANA BANDANA')
```

### `decompress(encoded)`

Decompresses payloads produced by `compress`.

Returns `:error` for malformed packets, invalid headers, or invalid bitstrings.

```oak
huffman := import('compression-huffman')

source := 'BANANA BANDANA'
packed := huffman.compress(source)
restored := huffman.decompress(packed)

restored = source // => true
```

### `compressed?(encoded)`

Returns true if payload starts with `Magic`.

```oak
huffman := import('compression-huffman')

huffman.compressed?('plain') // => false
huffman.compressed?(huffman.compress('abc')) // => true
```

## Notes

- Small inputs can grow because of header overhead.
- Best on text with strongly non-uniform character frequencies.
- Output is text-safe and easy to inspect.
