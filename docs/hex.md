# lib/hex

`lib/hex.oak` is a minimal hexadecimal helper module shared by the low-level
encoders in this tree (base64url, Java-source emission, byte tooling). It
depends on `libstd` only.

## Exports

Encoding (lowercase output):

- `hexEncodeNibble(n)` — one lowercase hex digit for `n` (0–15).
- `hexEncodeByte(b)` — two-character lowercase encoding of `b` (0–255).
- `bytesToHex(bytes)` — hex string for a list of byte values.

Decoding (accepts either case):

- `hexDecodeNibble(c)` — value of one hex digit, or `?` for a non-digit.
- `hexDecodePair(s, i)` — byte at offsets `i`, `i + 1` of `s`, or `?`.
- `hexToBytes(s)` — list of bytes; requires even length, returns `?` on any
  invalid digit.

## Notes

- `hexToBytes` is total: malformed input yields `?` rather than a runtime
  error from object-lookup arithmetic.
- `std`'s own `toHex`/`fromHex` remain in place (`std` cannot import `hex` —
  `hex` itself imports `std`, so that would be an import cycle). They handle
  number-scale integers rather than byte lists.
