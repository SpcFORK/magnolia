# Data Protection Helpers (dataprot)

## Overview

`libdataprot` provides lightweight payload integrity helpers for Oak byte strings and byte lists.

The module covers three common needs:
- parity helpers for single-bit checks
- additive/XOR/CRC checksums for transport and storage validation
- LDPC parity-matrix checks for validating binary codewords

It is intentionally focused on integrity and validation. The LDPC functions do not implement belief-propagation decoding or code construction.

## Import

```oak
dataprot := import('dataprot')

{
    parity: parity
    parityBit: parityBit
    parityValid?: parityValid?
    xorChecksum: xorChecksum
    sumChecksum8: sumChecksum8
    sumChecksum16: sumChecksum16
    sumChecksum32: sumChecksum32
    crc16Ccitt: crc16Ccitt
    crc32: crc32
    hammingDistance: hammingDistance
    ldpcSyndrome: ldpcSyndrome
    ldpcValid?: ldpcValid?
    ldpcCheck: ldpcCheck
    ldpcCandidates: ldpcCandidates
    ldpcCorrect: ldpcCorrect
} := import('dataprot')
```

## Byte-Oriented Checks

All checksum functions accept either:
- a byte string
- a list of integers in the range `0..255`

Invalid payloads return `:error`.

### `parity(data)`

Returns the raw payload parity bit:
- `0` when the payload has even parity
- `1` when the payload has odd parity

```oak
{ parity: parity } := import('dataprot')

parity('ABC')
// => 1
```

### `parityBit(data, odd)`

Returns the check bit needed to make the total parity even by default, or odd when `odd` is `true`.

```oak
{ parityBit: parityBit } := import('dataprot')

parityBit('ABC')
// => 1  (append 1 for even total parity)

parityBit('ABC', true)
// => 0  (append 0 for odd total parity)
```

### `parityValid?(data, checkBit, odd)`

Checks whether a payload plus parity bit matches the requested convention.

```oak
{ parityValid?: parityValid? } := import('dataprot')

parityValid?('ABC', 1)
// => true

parityValid?('ABC', 0, true)
// => true
```

### `xorChecksum(data)`

Computes the XOR of all payload bytes.

```oak
{ xorChecksum: xorChecksum } := import('dataprot')

xorChecksum('ABC')
// => 64
```

### `sumChecksum8(data)`

Computes an additive checksum modulo $2^8$.

### `sumChecksum16(data)`

Computes an additive checksum modulo $2^{16}$.

### `sumChecksum32(data)`

Computes an additive checksum modulo $2^{32}$.

```oak
{
    sumChecksum8: sumChecksum8
    sumChecksum16: sumChecksum16
    sumChecksum32: sumChecksum32
} := import('dataprot')

sumChecksum8([255, 1])
// => 0

sumChecksum16([255, 255, 1])
// => 511

sumChecksum32([255, 255, 255, 255])
// => 1020
```

### `crc16Ccitt(data, seed, poly)`

Computes CRC-16/CCITT-FALSE by default:
- seed: `0xFFFF`
- polynomial: `0x1021`

```oak
{ crc16Ccitt: crc16Ccitt } := import('dataprot')

crc16Ccitt('123456789')
// => 10673  (0x29B1)
```

### `crc32(data, seed, poly, finalXor)`

Computes reflected IEEE CRC-32 by default:
- seed: `0xFFFFFFFF`
- polynomial: `0xEDB88320`
- final xor: `0xFFFFFFFF`

```oak
{ crc32: crc32 } := import('dataprot')

crc32('123456789')
// => 3421780262  (0xCBF43926)
```

### `hammingDistance(a, b)`

Counts differing bits between two equally-sized byte sequences.

Returns `:error` when the payload lengths differ or the inputs are not valid byte strings/lists.

```oak
{ hammingDistance: hammingDistance } := import('dataprot')

hammingDistance('A', 'C')
// => 1
```

## LDPC Checks

The LDPC helpers operate on binary vectors and parity-check matrices.

Accepted bit formats:
- lists of `0` and `1`
- lists of booleans
- strings containing only `'0'` and `'1'`

### `ldpcSyndrome(word, parityMatrix)`

Computes the syndrome vector $H \cdot x^T \bmod 2$.

Returns `:error` when the word or matrix contains invalid bits, or when a row width does not match the word length.

### `ldpcValid?(word, parityMatrix)`

Returns `true` when every syndrome bit is zero.

### `ldpcCheck(word, parityMatrix)`

Returns a structured result:

```oak
{
    valid: <bool>
    syndrome: <list[int]>
    failed: <list[int]>
    weight: <int>
}
```

`failed` contains the row indices of unsatisfied parity checks.

### `ldpcCandidates(word, parityMatrix)`

Returns the bit indices whose parity-check matrix column matches the current syndrome.

This is useful for single-bit fault localization in Hamming-like and sparse parity-check matrices.

### `ldpcCorrect(word, parityMatrix)`

Attempts a single-bit correction when the syndrome maps to exactly one matrix column.

Returns a structured result:

```oak
{
    valid: <bool>
    corrected: <bool>
    word: <same shape as input word>
    index: <int|null>
    syndrome: <list[int]>
    failed: <list[int]>
}
```

- `valid` is true when the returned word satisfies every parity check.
- `corrected` is true only when a single-bit correction was applied successfully.
- `index` is the corrected bit index, or `?` when no unique correction was available.

```oak
dataprot := import('dataprot')

H := [
    '1010101'
    '0110011'
    '0001111'
]

word := '1010101'
broken := '1010100'

dataprot.ldpcValid?(word, H)
// => true

dataprot.ldpcSyndrome(broken, H)
// => [1, 1, 1]

dataprot.ldpcCheck(broken, H)
// => { valid: false, syndrome: [1, 1, 1], failed: [0, 1, 2], weight: 3 }

dataprot.ldpcCandidates(broken, H)
// => [6]

dataprot.ldpcCorrect(broken, H)
// => { valid: true, corrected: true, word: '1010101', index: 6, syndrome: [0, 0, 0], failed: [] }
```

## Notes

- The additive checksum helpers are straight modular sums, not Internet checksum or Fletcher/Adler variants.
- `crc16Ccitt` and `crc32` expose optional tuning arguments so alternate seeds or polynomials can be reused without duplicating the core loop.
- `ldpcCorrect` performs only single-bit syndrome correction. It is not a general LDPC decoder and does not implement belief propagation or iterative soft decisions.