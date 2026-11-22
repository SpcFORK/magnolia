# lib/sample

`lib/sample.oak` is the shared **sample-conversion primitive** library. It
defines the exact float-to-fixed-point conversions between a normalised
float sample in **[-1, 1]** and the fixed-width integer representations
used by PCM audio containers, in one place so the audio codecs cannot
drift on the scaling constants.

This module was abstracted from the binary sample encoders/decoders of the
audio codecs (`lib/audio-wav.oak`, `lib/audio-ogg.oak`, `lib/audio-aiff.oak`,
`lib/audio-au.oak`, `lib/audio-raw.oak`) and the audio facade
(`lib/audio.oak`). Each audio lib now imports `sample` and re-stubs its
former local helpers as thin delegates, so the conversion semantics are
defined once and every consumer gets the same bit-exact behaviour.

## Primitives

All operate on a single sample value; no dependencies (pure arithmetic).

**Clamping**
- `_clipF32(s)` — clamp a float sample to the normalised [-1, 1] range.

**Signed conversions** (float [-1,1] ↔ signed fixed-width)
- `_f32ToI8(s)`, `_i8ToF32(s)` — 8-bit signed (-128..127).
- `_f32ToI16(s)`, `_i16ToF32(s)` — 16-bit signed (scale 32767 / 32768).
- `_f32ToI32(s)`, `_i32ToF32(s)` — 32-bit signed (scale 2147483647 / 2147483648).

**Unsigned 8-bit conversion** (float [-1,1] ↔ 0..255, 128 = silence)
- `_f32ToU8(s)`, `_u8ToF32(s)`.

## Why it is a primitive

- The scaling constants (32767, 32768, 2147483647, 2147483648, 127, 128,
  127.5, 255) are the exact values each codec previously carried inline;
  centralising them removes the risk of a codec drifting from the others.
- `_f32ToI8`'s saturating clamp and `_clipF32`'s range guard are defined
  once, so every codec applies identical bounds.
- New consumers (packing, other formats) can reuse the primitives instead
  of duplicating the conversion math.
