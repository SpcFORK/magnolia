# lib/ieee754

`lib/ieee754.oak` implements IEEE-754 binary32 (`float`) and binary64
(`double`) values in Oak. It is a general-purpose primitive: the JVM
execution layer uses it for `float`/`double` pool payloads, two-slot locals
and category-two operations, and any other Oak program can rely on the same
bit-exact semantics. It depends on `str` only (plus the built-in `pow`/`log`/`int`/`float` used below).

## What it represents

Tagged Oak objects with the interchange shape shared with the bounded JVM
execution layer and its generated Java `WideFloat`/`WideDouble`:

```oak
{jvmFloat: true, bits}          // f32 bits, 0 .. 4294967295
{jvmDouble: true, hi, lo}       // f64 bits as two 32-bit limbs
```

Neither negative zero nor NaN is represented by an untagged Oak number.

## Exports

`f32(bits)`, `f64(hi, lo)` — wrap raw bit patterns.
`float?(v)`, `double?(v)` — predicates.
`wide?(v)` — true for the double (category-two, two-slot) representation.
`sign(v)`, `kind(v)` — `:zero`, `:nan`, `:infinity` or `:finite` (subnormals report `:finite`).
`numeric(v)` — the value as a host number; only finite values and zero,
exact because every binary32 and the binary64 significand fits in binary64;
NaN/infinity return `?`.
`fromNumber(value, double, negativeZero?)` / `fromLong(hi, lo)` — construct from a
host number (with the sign-of-zero flag) or from exact 64-bit limbs.
`negative(v)`, `zero(double, negative?)`, `infinity(double, negative?)`,
`nan(double)` — constructors for special values (`double` picks the
representation).
`binary(op, a, b)` — the IEEE binary operations, selected by JVM opcode:
`98`/`99` fadd/dadd, `102`/`103` fsub/dsub, `106`/`107` fmul/dmul,
`110`/`111` fdiv/ddiv, `114`/`115` frem/drem; NaN propagates, signed zero
division returns infinity (never a host exception), and unmodeled ops return
`?`.
`compare(a, b, nanResult)` — `fcmpl`/`fcmpg`/`dcmpl`/`dcmpg` NaN semantics via the
`nanResult` argument (`-1` → NaN compares low, `1` → NaN compares high).
`toInt(v)` — truncation toward zero with JVM `NaN`/saturation semantics.
`text(v)` — decimal formatting for the finite/zero/infinity/NaN subset;
values outside the modeled subset return `?`.

## Relationship to the JVM execution layer

`lib/applet-runtime.oak` imports `ieee754` for constant-pool payloads and
category-two operations; the generated Java preamble (`WideFloat`/`WideDouble`)
is a separate implementation of the same shapes and must be kept in sync by
hand (there is no JDK oracle in CI).

## Status and verification

- `fromLong` (long→float, the double-rounding fix site) was model-checked
  against an exact integer/fraction oracle of round-half-to-even f32
  conversion on 20,132 in-domain int64 vectors including the `2^55 + 2^31 +
  1` midpoint case from the fix comment: **zero mismatches**. An earlier
  `int → float64 → float32` oracle "failed" 2 vectors at that exact midpoint —
  those were the double-rounding cases the fix exists to handle, and the
  exact oracle confirms the implementation is correctly rounded there.
- Still unverified: NaN payload policy details, the `binary` remainder on
  subnormals at the host-rounding edge, Oak execution of any of this, and the
  generated-Java mirror (`WideFloat`/`WideDouble`, kept in sync by hand).

## Notes

- Host `log`/`pow` are used only to classify exponent/significance and for
  `binary` arithmetic; exceptional cases (division by zero, NaN) never reach
  the host.
- Decimal formatting outside the modeled subset fails explicitly rather than
  silently printing a wrong value.
