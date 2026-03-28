# Math Base (math-base)

## Overview

`libmath-base` provides foundational math primitives used by the math sub-modules (`math-geo`, `math-stats`). It exists to break circular dependencies — sub-modules can import `math-base` for core functions like `sqrt`, `abs`, and `sign` without depending on the full `math` module.

For most application code, prefer importing `math` directly, which re-exports everything from `math-base`.

## Import

```oak
mb := import('math-base')
{ sqrt: sqrt, abs: abs, sign: sign, Pi: Pi } := import('math-base')
```

## Constants

### `Pi`

The mathematical constant π.

```oak
Pi // => 3.14159265358979323846264338327950288419716939937510
```

### `E`

The mathematical constant e (base of natural logarithm).

```oak
E // => 2.71828182845904523536028747135266249775724709369995
```

## Functions

### `sign(n)`

Returns `-1` for negative numbers, `1` for zero and positive numbers.

```oak
sign(-5)  // => -1
sign(0)   // => 1
sign(42)  // => 1
```

### `abs(n)`

Returns the absolute value of a number.

```oak
abs(-7)   // => 7
abs(3.14) // => 3.14
```

### `sqrt(n)`

Returns the principal square root of a non-negative number. Returns `?` for negative inputs.

```oak
sqrt(16) // => 4
sqrt(2)  // => 1.4142135623730951
sqrt(-1) // => ?
```

## See Also

- [math.md](math.md) — Full math library (re-exports math-base)
- [math-geo.md](math-geo.md) — Geometry and trigonometry
- [math-stats.md](math-stats.md) — Statistical functions
