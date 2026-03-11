# gui-draw

Cross-platform drawing primitives used by the GUI facade: text, lines, rectangles and simple pen/brush helpers.

Key exports

- `drawText(window, x, y, text)`
- `fillRect(window, x, y, width, height, color?)`
- `drawLine(window, x1, y1, x2, y2, color?)`

Behavior

- On native backends these map to platform drawing APIs (GDI/X11).
- On web backends they record logical draw ops into `window.messages` for host-side rendering.

Example

```oak
gd := import('gui-draw')
gd.fillRect(window, 10, 10, 200, 120, gui.rgb(200,80,40))
gd.drawText(window, 16, 28, 'Sample')
```
