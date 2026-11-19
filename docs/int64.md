# lib/int64

`lib/int64.oak` implements exact signed **64-bit integers** in Oak. It is a
general-purpose primitive: the JVM execution layer uses it for `long`
arithmetic, and any other Oak program can too. It depends on `libstd` only.

## What it represents

A value is a plain Oak object with the interchange shape shared with the
bounded JVM execution layer and its generated Java `WideLong`:

```oak
{jvmLong: true, hi: <signed int32>, lo: <unsigned int32>}
```

- `hi` is the signed high 32-bit limb (`-2147483648` … `2147483647`).
- `lo` is the unsigned low 32-bit limb (`0` … `4294967295`).

All arithmetic is limb-exact; no JS `Number` floats beyond the safe integer
range are ever involved.

## Exports

`u32(n)`, `s32(n)` — wrap any number into unsigned/signed 32-bit limb form.
`long?(v)` — predicate for the interchange shape.
`long(hi, lo)` — construct a value (limbs are normalized).
`add(a, b)`, `sub(a, b)`, `mul(a, b)` — wrapping two's-complement arithmetic
(`add`/`sub`/`mul` wrap modulo 2^64, matching the JVM).
`neg(a)` — two's-complement negation (`MIN_VALUE` wraps to itself).
`cmp(a, b)` — `-1`, `0` or `1`.
`and(a, b)`, `or(a, b)`, `xor(a, b)` — bitwise operators (limb-exact).
`shiftLeft(a, n)`, `shiftRightUnsigned(a, n)`, `shiftRightSigned(a, n)` —
shift count masked to its low six bits, as the JVM requires;
`shiftRightUnsigned` is logical, `shiftRightSigned` is arithmetic.
`divRem(a, b)` — truncated division, JVM sign rules; returns
`{type: :ok, quotient, remainder}` or `{type: :error, message: 'long division by zero'}`.
`MIN_VALUE / -1` wraps to `MIN_VALUE`.
`toText(value)` — decimal formatting without a 64-bit or floating-point
intermediate; at most twenty digits.

## Relationship to the JVM execution layer

`lib/applet-runtime.oak` owns the JVM-side policy: argument validation,
raising `java/lang/ArithmeticException` on division by zero, and the
`lshl`/`lshr`/`lushr` opcode dispatch. It delegates all limb arithmetic to
`import('int64')`, so generated standalone Oak (which inlines this module,
see `docs/applet.md`) keeps the same behavior.

## Status and verification

- On 2026-11-16, during the extraction into this module, the multiply path
  was found to be **broken as it existed** in the unvetted in-tree addition:
  the 16-bit-limb schoolbook dropped the carry out of the high 16-bit digit
  (`a1*b0 + carry`) and packed the wrong limbs into the 64-bit result, so
  `lmul` produced garbage for essentially all non-trivial operands (a Python
  limb model of the original code was wrong on 20,000/20,000 random vectors).
  The fix is minimal (carry the limb, pack `[hi32, w2*65536 + w0]`) and is
  covered below.
- A Python model of this module's exact algorithms (limb ops, shift masking,
  truncated division, decimal formatting) was checked against
  arbitrary-precision arithmetic on ~30k generated vectors plus the edge set
  (`MIN_VALUE`, `MIN_VALUE/-1` wraps to `MIN_VALUE`, `MIN_VALUE/1`, sign-of-
  remainder = dividend sign): **zero mismatches**. This validates the limb
  algorithms as extracted, **not** a JDK oracle, **not** an Oak-executed test
  run, and **not** the frame gate's admission of these opcodes (the gate
  still does not admit them). It remains unvalidated assistance; the
  keep/revert decision in [docs/applet.md](docs/applet.md) and
  [applet.md](applet.md) still governs.

## Portability notes

- Every operation is constant-time by construction (fixed upper bound of
  limb steps), with no host floating-point involvement for arithmetic.
- `divRem` is a bit-by-bit unsigned-magnitude algorithm (64 iterations);
  it is correct, not fast.
