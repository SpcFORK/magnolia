# lib/bytes

`lib/bytes.oak` is the shared **byte-order primitive** library. It defines
the exact little/big-endian encoding and decoding of 8/16/32-bit unsigned
and signed integers in one place, so no language library has to carry a
private copy that can drift.

This module was abstracted from the binary encoders/decoders of the audio
codecs (`lib/audio-wav.oak`, `lib/audio-ogg.oak`, `lib/audio-aiff.oak`,
`lib/audio-au.oak`), the MessagePack codec (`lib/msgpack.oak`), the
style-bake binary format (`lib/style-bake.oak`) and the JVM class-file
machinery. Each language lib now imports `bytes` and re-stubs its former
local helpers as thin delegates, so the encoding semantics are defined
once and every consumer gets the same bit-exact behavior.

## Encoding

Encoders return **lists of byte integers** (0 .. 255):

`_u8(n)`, `_u16LE(n)`, `_u16BE(n)`, `_u32LE(n)`, `_u32BE(n)` — unsigned
word encoders. `_i16LE(n)`, `_i16BE(n)`, `_i32LE(n)`, `_i32BE(n)` —
signed word encoders, two's complement, wrap the corresponding unsigned
encoder over `n + 2^bits` for negatives.

`_byteStr(blist)` — converts a list of byte integers to a binary string
(via `char`), the representation the audio/msgpack libs build containers
with.

## Decoding

Decoders read from **binary strings** via `codepoint`, returning the
matching word:

`_readU8(s, off)`, `_readI8(s, off)`, `_readU16LE/BE(s, off)`,
`_readI16LE/BE(s, off)`, `_readU32LE/BE(s, off)`, `_readI32LE/BE(s, off)`.
Signed decoders convert the unsigned word with the standard sign extension
(`u > 2^(bits-1)-1 -> u - 2^bits`).

## Why it is a primitive

- No JS Number precision is lost: 16/32-bit words stay in integral limbs.
- The encoders and decoders are the exact same algorithms the private
  copies used, verified by the existing codec test suites.
- New consumers (Java output, wasm VM, packing) can reuse the primitives
  instead of duplicating word-order logic.

It depends on `std` only (`map`/`join` for `_byteStr`).
