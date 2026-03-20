# GUI Color Library (gui-color)

## Overview

`gui-color` provides packed RGB color helpers used by `GUI` and other GUI
submodules.

## Import

```oak
colors := import('gui-color')
```

## Exports

- `_clampByte(value)` - clamps to `0..255`
- `_clampOpacity(value)` - clamps to `0..1`
- `rgb(r, g, b)` - packs channels into one integer
- `colorR(color)` - extracts red channel
- `colorG(color)` - extracts green channel
- `colorB(color)` - extracts blue channel
- `opacity(color, amount, background?)` - alpha-composes foreground over background
- `rgba(r, g, b, a, background?)` - convenience alpha composition wrapper

## Notes

- `GUI` re-exports these helpers, so most callers can keep using `import('GUI')`.
- Colors are stored as packed integers for backend portability.
