# Shader Color Helpers (gui-shader-color)

## Overview

`gui-shader-color` provides packed-RGB color manipulation for the CPU shader
engine: channel extraction, blending, HSL/HSV conversion, cosine palettes,
contrast, and sepia. It depends only on `gui-color` (for `rgb` packing) and
`gui-shader-math`.

## Import

```oak
col := import('gui-shader-color')
```

All symbols are also re-exported by `gui-shader` and through the `GUI` facade.

## Packed RGB

Colors are stored as single integers: `r | (g << 8) | (b << 16)`.

| Function | Description |
|----------|-------------|
| `packRGB(r, g, b)` | Pack r, g, b [0–255] into a single int |
| `unpackRGB(c)` | Returns `{ r, g, b }` from a packed color |
| `colorR(c)` | Extract red channel [0–255] |
| `colorG(c)` | Extract green channel [0–255] |
| `colorB(c)` | Extract blue channel [0–255] |

## Blending & Interpolation

| Function | Description |
|----------|-------------|
| `mix(a, b, t)` | Linearly interpolate between two packed colors (t in [0,1]) |
| `mix3(a, b, c, t)` | Interpolate across three colors: a→b→c as t goes 0→0.5→1 |
| `overlay(fg, bg, alpha)` | Blend foreground over background with alpha [0,1] |

## Brightness & Tone

| Function | Description |
|----------|-------------|
| `brighten(c, amount)` | Add amount [0–255] to each channel (clamped) |
| `darken(c, amount)` | Subtract amount from each channel |
| `invert(c)` | Complement color (255 − each channel) |
| `grayscale(c)` | Convert to grayscale using luminance weights (0.299/0.587/0.114) |
| `contrast(c, amount)` | Adjust contrast around the midpoint |
| `sepia(c)` | Apply a sepia tone filter |

## Color Space Conversion

| Function | Description |
|----------|-------------|
| `hsl2rgb(h, s, l)` | HSL (all [0,1]) → packed RGB |
| `rgb2hsl(c)` | Packed RGB → `{ h, s, l }` (all [0,1]) |
| `hsv2rgb(h, s, v)` | HSV (all [0,1]) → packed RGB |
| `rgb2hsv(c)` | Packed RGB → `{ h, s, v }` (all [0,1]) |

## Advanced

| Function | Description |
|----------|-------------|
| `floatStr(v)` | Format a 0–255 channel value as a GLSL float literal string |
| `cosinePalette(t, a, b, c, d)` | Generate a color from iq's cosine palette technique. `a`, `b`, `c`, `d` are `{ x, y, z }` vectors. |
