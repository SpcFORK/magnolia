# Windows GDI Library (windows-gdi)

## Overview

`windows-gdi` provides Win32 painting and GDI wrappers used by `windows` and
native GUI backends.

## Import

```oak
wgdi := import('windows-gdi')
```

## Exports

- `beginPaint`, `endPaint`
- `getDC`, `releaseDC`
- `getStockObject`, `selectObject`
- `setBkMode`, `setTextColor`, `textOut`
- `createFont`
- `rectangle`, `ellipse`
- `createSolidBrush`, `deleteObject`
