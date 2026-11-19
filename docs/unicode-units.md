# lib/unicode-units

`lib/unicode-units.oak` converts between UTF-8/modified-UTF-8 byte sequences
and UTF-16 code units, in Oak. It depends on `libstd` only.

## Why it exists

Two front-ends in this tree historically carried duplicate private decoders:

- the Java source emitter (`lib/applet-java.oak`) needed to decode the bytes
  of Oak strings (plain UTF-8, plus JVM modified UTF-8 for class-file data)
  into the UTF-16 code units Java source requires — Java's Unicode escapes
  are processed before lexing, so ASCII quotes, backslashes and newlines must
  never be emitted as `\u0022` / `\u005c` / `\u000a`;
- the ES3/ES5 tokenizer (`lib/js-tokenize.oak`) needed the inverse: encoding
  a codepoint as standard UTF-8 bytes.

Both now import this module.

## Exports

`utf8Units(text)` — decodes an Oak string (byte-indexed UTF-8, or modified
UTF-8) into UTF-16 code units.

- a C0 80 modified-UTF-8 NUL decodes to `U+0000`;
- a standard four-byte sequence is re-encoded as a UTF-16 surrogate pair
  (matching Java `String` storage);
- malformed input returns `{type: :error, message: ...}` instead of producing
  replacement guesses; success returns `{type: :ok, value: [unit, ...]}`.

`utf8Bytes(cp)` — encodes one Unicode codepoint as standard UTF-8 bytes.
Non-BMP planes use the four-byte form. If the range is invalid, the result is
undefined — validate inputs with the predicates in `lib/str.oak` or check the
returned length.

## Scope

This module is intentionally minimal. It is **not** a general Unicode
library: no normalization, no grapheme clusters, no case mapping.
