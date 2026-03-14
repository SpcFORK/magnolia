# RLE Compression Library (compression-rle)

## Overview

`compression-rle` implements escape-safe run-length encoding (RLE) for Oak strings.

Use this codec when your data has repeated character runs, for example:
- Log lines with repeated delimiters
- Padding-heavy text
- Simple, highly repetitive payloads

## Import

```oak
rle := import('compression-rle')
{ compress: compress, decompress: decompress } := import('compression-rle')
```

## Constants

```oak
Marker := '\x1d'
RunPrefix := '#'
RunSep := ':'
```

## API

### `compress(input)`

Compresses a string with run-length tokens.

Rules:
- Literal characters are emitted as-is.
- Marker is escaped as `Marker + Marker`.
- Repeated runs may be emitted as `Marker + '#' + <count> + ':' + <char>`.
- Runs are emitted only when shorter than literal output.

```oak
rle := import('compression-rle')

rle.compress('aaaaabbbbcc')
// => '\x1d#5:a\x1d#4:bcc'

rle.compress('\x1d\x1d')
// => '\x1d\x1d\x1d\x1d'
```

### `decompress(encoded)`

Decompresses RLE payloads produced by `compress`.

Returns `:error` for malformed token sequences.

```oak
rle := import('compression-rle')

rle.decompress('\x1d#6:x')
// => 'xxxxxx'

rle.decompress('\x1d#0:a')
// => :error
```

### `compressed?(encoded)`

Returns true if the string contains the RLE marker pattern.

```oak
rle := import('compression-rle')

rle.compressed?('plain text') // => false
rle.compressed?(rle.compress('aaaa')) // => true
```

## Notes

- RLE is very fast and easy to debug.
- It is usually ineffective for non-repetitive text.
- This codec is used as the default in `compression.compress()`.
