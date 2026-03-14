# Compression Library (compression)

## Overview

`libcompression` provides multiple lossless string codecs:

- RLE for repeated character runs
- Huffman coding for frequency-based bit packing
- LZW for dictionary-based phrase compression

The default `compress()` and `decompress()` functions keep using RLE for backward compatibility.
The implementation is split internally across `compression-rle`, `compression-huffman`, and `compression-lzw`, while `compression` remains the public facade.

Codec-specific docs:
- [RLE codec](compression-rle.md)
- [Huffman codec](compression-huffman.md)
- [LZW codec](compression-lzw.md)

Benchmark sample:
- `samples/compression-benchmark.oak`

## Import

```oak
compression := import('compression')
{
    compress: compress
    decompress: decompress
    huffmanCompress: huffmanCompress
    lzwCompress: lzwCompress
} := import('compression')
```

## Generic API

### `compress(input, algorithm?)`

Compresses `input` with the requested codec.

Supported algorithms:
- `:rle` or `'rle'`
- `:huffman` or `'huffman'`
- `:lzw` or `'lzw'`

If `algorithm` is omitted, RLE is used.

```oak
compression.compress('aaaaabbbb')
compression.compress('BANANA BANDANA', :huffman)
compression.compress('TOBEORNOTTOBEORTOBEORNOT', :lzw)
```

### `decompress(encoded, algorithm?)`

Decompresses content previously produced by the selected codec.

If `algorithm` is omitted, RLE is used.

Returns `:error` for malformed data or unknown algorithms.

```oak
compressed := compression.compress('aaaaabbbb')
compression.decompress(compressed)

compressed := compression.compress('BANANA BANDANA', :huffman)
compression.decompress(compressed, :huffman)
```

### `compressed?(encoded, algorithm?)`

Checks whether a value appears to be compressed by a supported codec.

- Without `algorithm`, returns true for any supported format.
- With `algorithm`, checks only that codec.

```oak
compression.compressed?(value)
compression.compressed?(value, :huffman)
compression.compressed?(value, :lzw)
```

## RLE Codec

RLE is the default codec and remains compatible with the original API.

### Constants

```oak
Marker := '\x1d'
RunPrefix := '#'
RunSep := ':'
```

### `rleCompress(input)`

Encodes repeated runs using an escape-safe token format.

Rules:
- Literal characters are emitted as-is.
- `Marker` is escaped as `Marker + Marker`.
- Runs may be emitted as `Marker + '#' + <count> + ':' + <char>`.
- Runs are used only when shorter than the literal form.

### `rleDecompress(encoded)`

Decodes an RLE stream. Returns `:error` for malformed tokens.

### `rleCompressed?(encoded)`

Checks whether a string contains the RLE marker form.

```oak
packed := compression.rleCompress('aaaaabbbbcc')
restored := compression.rleDecompress(packed)
```

## Huffman Codec

Huffman compression builds a character frequency table, derives a binary tree, and stores the resulting bitstream as a text-safe payload.

### Constants

```oak
HuffmanMagic := 'HUF1:'
```

### `huffmanCompress(input)`

Compresses `input` using a Huffman tree derived from the input character frequencies.

Implementation notes:
- The output stores a serialized frequency header followed by `|` and the encoded bitstring.
- Header entries are encoded as `<hex-codepoint>:<count>` and joined with commas.
- Small inputs may grow due to the header cost.

### `huffmanDecompress(encoded)`

Decompresses Huffman data created by `huffmanCompress()`.

Returns `:error` for invalid packets, malformed frequency headers, or malformed bitstreams.

### `huffmanCompressed?(encoded)`

Returns true when the payload starts with `HuffmanMagic`.

```oak
source := 'BANANA BANDANA BANANA BANDANA'
packed := compression.huffmanCompress(source)
restored := compression.huffmanDecompress(packed)
```

## LZW Codec

LZW compression builds a dictionary of repeated phrases and emits integer codes.

### Constants

```oak
LZWMagic := 'LZW1:'
```

### `lzwCompress(input)`

Compresses `input` using an LZW dictionary.

Implementation notes:
- The output stores a serialized alphabet header followed by `|` and a comma-separated integer code list.
- Alphabet entries are encoded as hexadecimal codepoints.
- Returns `:error` if the dictionary would exceed 65535 entries.

### `lzwDecompress(encoded)`

Decompresses data created by `lzwCompress()`.

Returns `:error` for malformed packets, invalid headers, invalid code streams, or unsupported dictionary growth.

### `lzwCompressed?(encoded)`

Returns true when the payload starts with `LZWMagic`.

```oak
source := 'TOBEORNOTTOBEORTOBEORNOT'
packed := compression.lzwCompress(source)
restored := compression.lzwDecompress(packed)
```

## Examples

### Compare Codecs

```oak
compression := import('compression')

source := 'BANANA BANDANA BANANA BANDANA'

rle := compression.compress(source)
huffman := compression.compress(source, :huffman)
lzw := compression.compress(source, :lzw)

println(len(rle))
println(len(huffman))
println(len(lzw))
```

### Safe Dispatch

```oak
compression := import('compression')

fn decodePacket(packet, codec) if codec {
    :rle -> compression.decompress(packet, :rle)
    :huffman -> compression.decompress(packet, :huffman)
    :lzw -> compression.decompress(packet, :lzw)
    _ -> :error
}
```

## Notes

- All codecs are lossless.
- RLE is best for repeated character runs.
- Huffman is best when the data has a skewed character distribution.
- LZW is best when the data contains recurring phrases or substrings.
- Huffman and LZW output formats are self-framed with magic prefixes plus explicit text-safe payload sections.
