# Shader Math Primitives (gui-shader-math)

## Overview

`gui-shader-math` provides pure-arithmetic math primitives for the CPU shader
engine. It has no external imports — all functions are self-contained and operate
on plain numbers or `{ x, y }` / `{ x, y, z }` vector dicts.

## Import

```oak
m := import('gui-shader-math')
```

All symbols are also re-exported by `gui-shader` and through the `GUI` facade.

## Constants

| Name | Value | Description |
|------|-------|-------------|
| `PI` | 3.14159… | Pi |
| `TAU` | 6.28318… | 2π |
| `HALF_PI` | 1.57079… | π/2 |
| `E` | 2.71828… | Euler's number |
| `DEG2RAD` | π/180 | Degrees → radians multiplier |
| `RAD2DEG` | 180/π | Radians → degrees multiplier |
| `SQRT2` | 1.41421… | √2 |

## Scalar Math

| Function | Description |
|----------|-------------|
| `fract(x)` | Fractional part of x |
| `mod(x, y)` | Modulo (floor-based) |
| `sign(x)` | Returns -1, 0, or 1 |
| `abs2(x)` | Absolute value |
| `clamp(x, lo, hi)` | Clamp x to [lo, hi] |
| `saturate(x)` | Clamp to [0, 1] |
| `lerpFloat(a, b, t)` | Linear interpolation |
| `lerp(a, b, t)` | Alias for `lerpFloat` |
| `inverseLerp(a, b, x)` | Inverse of lerp — returns t |
| `remap(x, inLo, inHi, outLo, outHi)` | Map x from one range to another |
| `step(edge, x)` | 0 if x < edge, else 1 |
| `smoothstep(e0, e1, x)` | Hermite interpolation (3t²−2t³) |
| `smootherstep(e0, e1, x)` | Perlin's improved smoothstep (6t⁵−15t⁴+10t³) |
| `min2(a, b)` | Minimum of two values |
| `max2(a, b)` | Maximum of two values |
| `sqr(x)` | x² |
| `sqrt(x)` | Square root via `pow(x, 0.5)` |
| `atan2(y, x)` | Two-argument arctangent |
| `pingpong(t, length)` | Triangle wave between 0 and length |
| `degToRad(d)` | Degrees to radians |
| `radToDeg(r)` | Radians to degrees |

## Easing Functions

All easing functions take `t` in [0, 1] and return the eased value.

| Function | Curve |
|----------|-------|
| `easeInQuad(t)` | Quadratic ease-in |
| `easeOutQuad(t)` | Quadratic ease-out |
| `easeInOutQuad(t)` | Quadratic ease-in-out |
| `easeInCubic(t)` | Cubic ease-in |
| `easeOutCubic(t)` | Cubic ease-out |
| `easeInOutCubic(t)` | Cubic ease-in-out |
| `easeInSine(t)` | Sine ease-in |
| `easeOutSine(t)` | Sine ease-out |
| `easeInOutSine(t)` | Sine ease-in-out |
| `easeInExpo(t)` | Exponential ease-in |
| `easeOutExpo(t)` | Exponential ease-out |
| `easeOutElastic(t)` | Elastic overshoot |
| `easeOutBounce(t)` | Bouncing ease-out |

## 2D Vector Math

Vectors are `{ x, y }` dicts.

| Function | Description |
|----------|-------------|
| `vec2(x, y)` | Create a 2D vector |
| `dot2(a, b)` | Dot product |
| `length2(v)` | Magnitude |
| `distance2(a, b)` | Distance between two points |
| `normalize2(v)` | Unit vector |
| `rotate2(v, angle)` | Rotate by angle (radians) |
| `scale2(v, s)` | Scalar multiply |
| `add2(a, b)` | Component-wise add |
| `sub2(a, b)` | Component-wise subtract |
| `lerp2(a, b, t)` | Component-wise lerp |
| `negate2(v)` | Negate both components |
| `abs2v(v)` | Absolute value per component |
| `min2v(a, b)` | Component-wise min |
| `max2v(a, b)` | Component-wise max |
| `floor2(v)` | Floor per component |
| `fract2(v)` | Fract per component |
| `reflect2(v, n)` | Reflect v across normal n |
| `toPolar(v)` | Convert to `{ r, theta }` |
| `fromPolar(r, theta)` | Convert from polar to Cartesian |

## 3D Vector Math

Vectors are `{ x, y, z }` dicts.

| Function | Description |
|----------|-------------|
| `vec3(x, y, z)` | Create a 3D vector |
| `add3(a, b)` | Component-wise add |
| `sub3(a, b)` | Component-wise subtract |
| `scale3(v, s)` | Scalar multiply |
| `dot3(a, b)` | Dot product |
| `length3(v)` | Magnitude |
| `distance3(a, b)` | Distance between two points |
| `normalize3(v)` | Unit vector |
| `cross3(a, b)` | Cross product |
| `lerp3(a, b, t)` | Component-wise lerp |
| `negate3(v)` | Negate all components |
| `reflect3(v, n)` | Reflect v across normal n |
