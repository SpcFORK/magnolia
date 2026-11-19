# Compression Library: ANS Codec (compression-ans)

## Overview

`compression-ans` implements **context-adaptive Asymmetric Numeral Systems**, the
integer-exact form of the arithmetic-coding idea from 3Blue1Brown's *Information
Theory, part 3*. It is the only codec in this repository that exploits context
when pricing symbols, which is what Shannon's theorem actually ties
compressibility to.

For comparison, on the same 327-byte passage of the US Constitution:

| codec | output | note |
|---|---|---|
| Huffman | 3721 bytes | context-free, fixed code per character |
| LZW | 745 bytes | context-free, dictionary phrases |
| **ANS** | **222 bytes** | per-context probability tables |

That is **70% smaller than LZW** and **94% smaller than Huffman**, with an exact
round-trip guaranteed by the checksum in the frame.

## Why context matters

The existing codecs all assign a character a fixed price. If `q` costs six bits
after `e` and six bits after `zzzz`, the model is throwing away everything it
knows. Shannon's result says the optimal cost of a message is
`-log2 P(message)`, and `P` is contextual: after `the`, `e` is nearly free;
after `xq`, almost nothing is.

ANS turns that into an algorithm. A probability table divides the state space
into slots, and folding the integer state through a slot is exactly the
"zoom into an interval" animation from the video - except the interval is
`[2^23, 2^32)` in integer space, so no bit of precision is ever lost.

## Import

```oak
ans := import('compression-ans')

// or through the facade
compression := import('compression')
compression.compress(text, :ans)
```

## API

### `compress(input, order?)`

Encodes `input` (any byte string) into a self-describing frame.

- `order` may be `0`, `1`, or `2` to pin the Markov model order.
- Omitted or `?`: all three orders are tried and the smallest frame wins.
- Returns `:error` only if no order can encode the input (does not currently
  happen for any byte string; the escape valve makes the coder total).

### `decompress(encoded)`

Restores the original byte string, or returns `:error`. Malformed input -
wrong magic, bad hex, truncated frame, non-canonical state, length or checksum
mismatch - is always an error, never a silently wrong string.

### `compressed?(encoded)`

Reports whether `encoded` starts with the ANS frame magic.

### `stats(input)`

Returns the numbers the video uses to compare an encoding against theory:

```oak
ans.stats(text)
// {
//   bytes: 5130
//   shannonBits: 19968              // order-0 Shannon information
//   shannonBitsPerChar: 3.89        // that, per character
//   frameBytes: 898                 // what actually shipped
//   frameBitsPerChar: 1.40          // beating order-0 Shannon via context
//   order0Bits: 19968               // alias of shannonBits
// }
```

Note `frameBitsPerChar` can be *lower* than `shannonBitsPerChar`. That is not a
violation of Shannon: the order-0 entropy is the floor for a model that ignores
context, and the coder uses context. The honest floor is the model's own
cross-entropy, which the frame achieves up to per-event rounding.

## Frame format

```
'ANS1:' <ver:2> <order:2> <inLen:8> <csum:4> <rawLen:6> <handle>
```

Fixed-width hex fields (status bytes), then a raw payload:

- `inLen` - original byte length, so decoding stops exactly
- `csum` - 16-bit rolling checksum over the original input
- `rawLen` - how many escaped bytes sit at the end of the handle
- `handle` - rANS final state (4 bytes, high byte first) followed by the
  renormalisation bytes, then the raw escaped bytes

A minimum frame is 35 bytes (31-byte header + 4-byte state), so ANS is not
worth using on strings shorter than roughly 100 bytes. Use RLE there.

## How it works

1. **Prior.** A small English-prose prior lives in the library, not the
   payload - free side information, the same idea as zstd's trained
   dictionaries. Both encoder and decoder blend it into every table.
2. **Epochs.** Every `RebuildPeriod` (64) symbols, the model rebuilds from the
   last `WindowSize` (160) bytes blended with the prior at weight 4. The decoder
   has the same bytes because it has decoded them.
3. **Tables.** Each context gets a frequency table normalised so slot counts
   sum to exactly `L = 2^10` or `2^12`, with `1/32` of the slots reserved for
   the escape symbol.
4. **Folding.** For symbol with frequency `f` and cumulative `c`:
   renormalise while `x >= f * (2^31 / L)` by flushing low bytes, then
   `x = (x // f) * L + (x % f) + c`.
5. **Decoding** inverts each step and pulls a byte whenever `x < 2^23`.
6. **Escapes.** A byte the table has never seen is coded through the reserved
   tail and its 8 raw bits go to the raw plane, so every byte string
   round-trips.

## Tuning (measured, not guessed)

The model constants were swept independently, then searched jointly, then
validated on a held-out corpus with no overlap with the tuning set.

| WindowSize / WindowWeight / EscDivisor | tuning bytes | held-out bytes | time |
|---|---|---|---|
| 160 / 4 / 32 — first hand-picked values | 929 | 960 | ~7.6 s |
| **256 / 32 / 24 — shipped** | **833** | **872** | **~7.5 s** |
| 384 / 24 / 24 | 832 | 870 | ~8.2 s |
| 512 / 48 / 24 | 826 | 868 | ~8.2 s |

The original hand-picked constants were **10% worse than measured ones at
identical speed**. The larger windows buy 2–4 more bytes and cost 10% more wall
time, so they were rejected: same ratio-per-second is the operating point.

`RebuildPeriod` has the sharpest trade-off. Smaller means better ratio and
superlinear cost:

| RebuildPeriod | total bytes | time |
|---|---|---|
| 8 | 844 | 16.7 s |
| 16 | 849 | 11.8 s |
| 32 | 867 | 8.3 s |
| **64 (shipped)** | **928 → 833 after retune** | ~7.5 s |
| 128 | 1099 | 3.7 s |
| 512 | 1861 | 1.8 s |

`EscDivisor` is the one constant with a **correctness** cliff, not just a cost:
values below ~16 start losing round-trips outright (`EscDivisor=2` produced 4
failures on the tuning corpus, because the reserved escape band stops being
able to represent the escape symbol's own state transitions).

The hyper-parameter sweep harness used during development
(`tools/compression-ans-sweep.oak`) has been ablated from this tree; the
constants above are the measured result it produced.

### The build gotcha that invalidated the first sweep

Libraries are embedded into the binary with `//go:embed` at **build** time.
Editing `lib/compression-ans.oak` and re-running a script measures the *old*
library until `go build` runs again. The original sweep harness missed this,
and reported a suspiciously perfect 928 bytes for all 60 parameter
combinations - an all-identical result that should have read as "the harness
is broken", not "the parameters do not matter". Any re-measurement must
rebuild between variants.

## Invariants (and the bugs they caught)

These are worth knowing because breaking any one produces corruption that
looks like success:

- **Renorm threshold.** The encoder must fold at `f * (2^31 / L)`, not
  `f * (2^32 / L)`. The wider bound still fits in 32 bits, but leaves states
  the decoder's `x < 2^23` rule never asks bytes for. A 20,000-trial
  randomized simulation distinguishes them; hand-checks do not.
- **Slot accounting.** `sum(freqs) + esc = L`. If the escape reserve is carved
  out before normalisation instead of after, some slots belong to no symbol and
  decode into garbage.
- **Causal models.** Epoch models may only look at bytes *before* the epoch
  boundary. Any peek at future input desynchronises encoder and decoder.
- **Order of flushed bytes.** The encoder walks input in reverse, so both the
  renormalisation bytes and the raw escape bytes must be reversed before
  writing.
- **No early returns.** In Oak, `if cond -> :error` evaluates to `:error` and
  continues. Guards must delegate to another function or nest the happy path.
- **Escape carve-out must exist.** After normalising slot counts to `L`, the
  `esc` slots have to be *subtracted from the largest symbol* so the multiset
  still sums to exactly `L`. Deleting this step (which a careless refactor did)
  makes `sum(freqs) + esc > L`, leaves slots belonging to no symbol, and
  decodes into garbage. The test suite asserts `sum(freqs) + esc = L` for this
  reason.
- **Memoise by reference, not by value.** Oak passes objects by value, so a
  cache written inside a callee mutates a copy. Cache state must live in a
  list (or other shared reference) owned by the caller.

## Choosing a codec

| data | use |
|---|---|
| natural language, code, JSON | **ans** |
| long runs of one byte | rle (ANS pays ~1/8 extra per escape) |
| short strings (<100 bytes) | rle, or nothing |
| incompressible binary | ans (degrades to raw size, plus header) |

## See also

- [compression](compression.md) - the facade and codec selection
- [compression-rle](compression-rle.md), [compression-huffman](compression-huffman.md), [compression-lzw](compression-lzw.md)
- `samples/compression-ans-demo.oak` - side-by-side comparison with the other codecs
