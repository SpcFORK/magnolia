# Math Library (math)

## Overview

`libmath` implements basic arithmetic and algebraic functions for Oak, including trigonometry, statistics, and geometric calculations.

For coordinate and angle functions, the coordinate plane is Cartesian with +x to the east and +y to the north. Angles are measured in radians from the +x axis, counterclockwise.

## Import

```oak
math := import('math')
// or destructure specific functions
{ sqrt: sqrt, sin: sin, cos: cos, Pi: Pi } := import('math')
```

## Constants

### `Pi`

The mathematical constant π (pi).

```oak
Pi // => 3.14159265358979323846264338327950288419716939937510
```

### `E`

The mathematical constant e (base of natural logarithm).

```oak
E // => 2.71828182845904523536028747135266249775724709369995
```

## Basic Arithmetic

### `sign(n)`

Returns `-1` for negative numbers, `1` for zero and positive numbers.

```oak
sign(-5) // => -1
sign(0) // => 1
sign(42) // => 1
```

### `abs(n)`

Returns the absolute value of a number.

```oak
abs(-5) // => 5
abs(3.14) // => 3.14
abs(0) // => 0
```

### `sqrt(n)`

Returns the principal square root of a non-negative number. Returns `?` for negative numbers.

```oak
sqrt(16) // => 4
sqrt(2) // => 1.4142135623730951
sqrt(-1) // => ?
```

### `round(n, decimals)`

Rounds `n` to the nearest `decimals`-th decimal place. Negative `decimals` returns `n` unchanged. Default is 0 decimals.

```oak
round(3.14159, 2) // => 3.14
round(3.14159, 0) // => 3
round(3.5) // => 4
round(-2.7, 0) // => -3
round(123.456, -1) // => 123.456
```

## Geometry

### `hypot(x0, y0, x1, y1)`

Returns the Euclidean distance between two points `(x0, y0)` and `(x1, y1)`. If `x1` and `y1` are omitted (or `?`), the second point defaults to the origin `(0, 0)`.

```oak
hypot(3, 4, 0, 0) // => 5 (distance from origin)
hypot(3, 4) // => 5 (same, defaults to origin)
hypot(0, 0, 3, 4) // => 5
hypot(1, 1, 4, 5) // => 5
```

### `bearing(x, y, d, t)`

Returns the point `[x', y']` at the end of a line segment starting at `(x, y)`, extending distance `d` at angle `t` (radians).

```oak
bearing(0, 0, 5, 0) // => [5, 0] (5 units east)
bearing(0, 0, 5, Pi/2) // => [0, 5] (5 units north)
bearing(10, 10, 10, Pi) // => [0, 10] (10 units west)
```

### `orient(x0, y0, x1, y1)`

Returns the angle (in radians, range `(-Pi, Pi]`) of the line from `(x0, y0)` to `(x1, y1)`. If `x1` and `y1` are omitted, assumes the origin as the starting point.

Equivalent to `atan2(y, x)` (note reversed argument order).

```oak
orient(1, 0) // => 0 (pointing east)
orient(0, 1) // => Pi/2 ~= 1.5708 (pointing north)
orient(-1, 0) // => Pi ~= 3.14159 (pointing west)
orient(0, -1) // => -Pi/2 ~= -1.5708 (pointing south)

orient(0, 0, 3, 3) // => Pi/4 ~= 0.7854 (45° northeast)
```

## Scaling and Clamping

### `scale(x, a, b, c, d)`

Maps value `x` from range `[a, b]` to range `[c, d]`. If `c` and `d` are omitted, maps to `[0, 1]`. Works even if `x` is outside `[a, b]`.

```oak
scale(5, 0, 10, 0, 100) // => 50
scale(5, 0, 10) // => 0.5 (map to [0, 1])
scale(15, 0, 10, 0, 100) // => 150 (extrapolates)
scale(-5, 0, 10, 0, 100) // => -50
```

### `clamp(x, a, b)`

Constrains `x` to the range `[a, b]`. Returns `x` if within bounds, otherwise returns the nearest bound.

```oak
clamp(5, 0, 10) // => 5
clamp(-5, 0, 10) // => 0
clamp(15, 0, 10) // => 10
clamp(7, 3, 20) // => 7
```

## Statistics

### `sum(xs...)`

Returns the sum of all given values.

```oak
sum(1, 2, 3) // => 6
sum(10, 20) // => 30
sum() // => 0
```

### `prod(xs...)`

Returns the product of all given values.

```oak
prod(2, 3, 4) // => 24
prod(5, 10) // => 50
prod() // => 1
```

### `min(xs...)`

Returns the minimum value.

```oak
min(5, 2, 8, 1) // => 1
min(10, -5) // => -5
```

### `max(xs...)`

Returns the maximum value.

```oak
max(5, 2, 8, 1) // => 8
max(10, -5) // => 10
```

### `mean(xs)`

Returns the arithmetic mean (average) of a list. Returns `?` for empty lists.

```oak
mean([1, 2, 3, 4, 5]) // => 3
mean([10, 20, 30]) // => 20
mean([]) // => ?
```

### `median(xs)`

Returns the median (middle value) of a list. For even-length lists, returns the mean of the two middle values. Returns `?` for empty lists.

```oak
median([1, 2, 3, 4, 5]) // => 3
median([1, 2, 3, 4]) // => 2.5
median([5, 1, 3, 2, 4]) // => 3 (sorts internally)
median([]) // => ?
```

### `stddev(xs)`

Returns the population standard deviation of a list. Returns `?` for empty lists.

```oak
stddev([1, 2, 3, 4, 5]) // => 1.4142135623730951
stddev([10, 10, 10]) // => 0
stddev([]) // => ?
```

## Examples

### Distance Between Points

```oak
{
    hypot: hypot
    sqrt: sqrt
} := import('math')

// Distance from (0, 0) to (3, 4)
distance := hypot(3, 4) // => 5

// Distance between two arbitrary points
fn distance(x1, y1, x2, y2) hypot(x1, y1, x2, y2)
distance(1, 1, 4, 5) // => 5
```

### Circle Calculations

```oak
{
    Pi: Pi
    cos: cos
    sin: sin
} := import('math')

radius := 10
circumference := 2 * Pi * radius // => 62.83...
area := Pi * pow(radius, 2) // => 314.15...

// Point on circle at angle theta
fn pointOnCircle(cx, cy, r, theta) [
    cx + r * cos(theta)
    cy + r * sin(theta)
]
```

### Statistical Analysis

```oak
{
    mean: mean
    median: median
    stddev: stddev
    min: min
    max: max
} := import('math')

data := [23, 45, 67, 12, 89, 34, 56]

println('Mean: ' + string(mean(data))) // => 46.57...
println('Median: ' + string(median(data))) // => 45
println('Std Dev: ' + string(stddev(data))) // => 24.29...
println('Range: ' + string(min(data...)) + ' to ' + string(max(data...)))
// => Range: 12 to 89
```

### Mapping Values

```oak
{ scale: scale, clamp: clamp } := import('math')

// Map sensor reading (0-1023) to voltage (0-5V)
voltage := scale(512, 0, 1023, 0, 5) // => 2.5

// Map temperature to color gradient, bounded
temp := 75
colorValue := scale(temp, 0, 100, 0, 255) |> clamp(0, 255)
```

### Polar Coordinates

```oak
{
    bearing: bearing
    orient: orient
    hypot: hypot
    Pi: Pi
} := import('math')

// Convert polar to cartesian
[x, y] := bearing(0, 0, 10, Pi/4)
// x ≈ 7.07, y ≈ 7.07

// Convert cartesian to polar
r := hypot(x, y)
theta := orient(x, y)
```

## Notes

- All trigonometric functions use radians, not degrees
- To convert degrees to radians: `radians := degrees * Pi / 180`
- To convert radians to degrees: `degrees := radians * 180 / Pi`
- Oak provides built-in functions: `sin()`, `cos()`, `tan()`, `asin()`, `acos()`, `atan()`, `pow()`, `int()`, `float()`
- Statistical functions operate on lists except `sum()` and `prod()` which use variadic arguments
- `orient()` is essentially `atan2()` with a more intuitive name for geometric operations
