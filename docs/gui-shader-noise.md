# Shader Noise & Hashing (gui-shader-noise)

## Overview

`gui-shader-noise` provides pseudo-random hashing and value noise functions for
the CPU shader engine. It depends only on `gui-shader-math`.

## Import

```oak
noise := import('gui-shader-noise')
```

All symbols are also re-exported by `gui-shader` and through the `GUI` facade.

## Hashing

| Function | Description |
|----------|-------------|
| `hash(seed)` | Pseudo-random float in [0,1) from a single seed |
| `hash2(a, b)` | Pseudo-random float from two inputs |
| `hash3(a, b, c)` | Pseudo-random float from three inputs |

The hash functions use a `fract(sin(…) * 43758.5453)` approach, matching the
common GPU-shader hashing idiom.

## Noise

| Function | Signature | Description |
|----------|-----------|-------------|
| `noise2D` | `(x, y)` | Smooth value noise in [0,1] at position (x, y). Uses bilinear interpolation of hashed grid corners with a Hermite smoothing curve. |
| `fbm` | `(x, y, octaves?)` | Fractal Brownian motion — layers `noise2D` at multiple frequencies. Default 4 octaves. Each octave halves amplitude and doubles frequency. |
