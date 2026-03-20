# GUI Graph Library (gui-graph)

## Overview

`gui-graph` provides graph and chart drawing helpers used by `GUI`.

All drawing helpers receive an injected drawing context as their first argument:

- `ctx.drawLine`
- `ctx.fillRect`
- `ctx.drawText`
- `ctx.rgb`

## Import

```oak
graphs := import('gui-graph')
```

## Exports

- `graphRange(values, options?)`
- `graphMapX(index, count, x, width)`
- `graphMapY(value, y, height, range)`
- `drawGraphAxes(ctx, window, x, y, width, height, options?)`
- `drawLineGraph(ctx, window, x, y, width, height, values, options?)`
- `drawBarGraph(ctx, window, x, y, width, height, values, options?)`
- `drawSparkline(ctx, window, x, y, width, height, values, options?)`

## Notes

- `GUI` re-exports window-first wrappers (`drawLineGraph(...)`, etc.).
- Option objects support min/max overrides, colors, grid visibility, and labels.
