# Math Geometry (math-geo)

## Overview

`math-geo` provides 2D coordinate geometry helpers: Euclidean distance, range scaling, polar-to-Cartesian conversion, and angle calculation. It is used internally by `libmath` and can also be imported directly.

For all angular functions: angles are in radians, measured counterclockwise from the +x axis.

## Import

```oak
geo := import('math-geo')
// or destructure
{ hypot: hypot, scale: scale, bearing: bearing, orient: orient } := import('math-geo')
```

## Constants

### `Pi`

The mathematical constant π.

```oak
Pi // => 3.14159265358979323846264338327950288419716939937510
```

## Functions

### `hypot(x0, y0, x1?, y1?)`

Returns the Euclidean distance between `(x0, y0)` and `(x1, y1)`. When `x1` and `y1` are omitted the distance from the origin is returned.

```oak
hypot(3, 4)        // => 5.0   (distance from origin)
hypot(0, 0, 3, 4)  // => 5.0
hypot(1, 1, 4, 5)  // => 5.0
```

### `scale(x, a, b, c?, d?)`

Maps `x` from the input range `[a, b]` to the output range `[c, d]`. When `c` and `d` are omitted the result is the normalised value in `[0, 1]`.

```oak
scale(5, 0, 10)         // => 0.5   (normalised)
scale(5, 0, 10, 0, 100) // => 50.0  (mapped to [0, 100])
scale(0, -1, 1, 0, 255) // => 127.5
```

### `bearing(x, y, d, t)`

Returns a new point `[x', y']` located `d` units away from `(x, y)` at angle `t` radians.

```oak
bearing(0, 0, 1, 0)    // => [1, 0]      (east)
bearing(0, 0, 1, Pi/2) // => [0, 1]      (north)
```

### `orient(x0, y0, x1?, y1?)`

Returns the angle in radians from `(x0, y0)` to `(x1, y1)`. When `x1` and `y1` are omitted, treats `(x0, y0)` as a direction vector from the origin.

```oak
orient(1, 0)        // => 0.0       (+x axis)
orient(0, 1)        // => ~1.5708   (π/2, +y axis)
orient(0, 0, 1, 1)  // => ~0.7854   (π/4, northeast)
```
