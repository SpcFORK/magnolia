# gui-print — Win32 Printing

`import('gui-print')` provides Win32 print dialog and GDI print rendering supporting page-by-page printing, print preview, PDF/XPS export via print-to-file, and printer capability queries.

## Quick Start

```oak
pr := import('gui-print')

// High-level: show dialog and print
pr.printDocument({}, fn(hDC, pageNum, pageSize) {
    pr.printSetFont(hDC, 24, 400, false, 'Arial')
    pr.printTextOut(hDC, 100, 100, 'Hello from Magnolia!')
    // Return true to print another page, false to stop
    false
})

// Print to file (PDF via system PDF printer)
pr.printToFile('output.pdf', 'My Document', fn(hDC, pageNum, pageSize) {
    pr.printTextOut(hDC, 50, 50, 'Page ' + string(pageNum))
    false
})
```

## Print Dialog

### `showPrintDialog(options)`

Displays the system print dialog. Returns `{type: :ok, hDC, fromPage, toPage, copies, flags}` or `{type: :cancel}`.

## Print Job Control

### `startDoc(hDC, docName, outputFile)`

Begins a print job. Returns `{type: :ok, jobId}` or `{type: :error}`.

### `startPage(hDC)`

Begins a new page.

### `endPage(hDC)`

Ends the current page.

### `endDoc(hDC)`

Ends the print job.

### `abortDoc(hDC)`

Cancels the print job.

### `deleteDC(hDC)`

Releases a printer device context.

## Drawing Primitives

### `printTextOut(hDC, x, y, text)`

Draws text at `(x, y)`.

### `printMoveTo(hDC, x, y)`

Moves the current drawing position.

### `printLineTo(hDC, x, y)`

Draws a line from the current position to `(x, y)`.

### `printRectangle(hDC, left, top, right, bottom)`

Draws a rectangle.

### `printEllipse(hDC, left, top, right, bottom)`

Draws an ellipse.

## GDI Object Management

### `printSetFont(hDC, height, weight, italic, fontName)`

Creates and selects a font. Returns the old font handle.

### `printSetTextColor(hDC, r, g, b)`

Sets the text color.

### `printSetBkMode(hDC, mode)`

Sets background mode: `1` = TRANSPARENT, `2` = OPAQUE.

### `printSetPen(hDC, style, width, r, g, b)`

Creates and selects a pen. Returns the old pen handle.

### `printDeleteObject(hObj)`

Deletes a GDI object (pen, brush, font).

## Device Capabilities

### `getDeviceCaps(hDC, index)`

Queries printer device capabilities by index.

### `getPrinterPageSize(hDC)`

Returns the printable area and DPI as `{horzRes, vertRes, logPixelsX, logPixelsY, physWidth, physHeight, offsetX, offsetY}`.

## Preview & Export

### `createPreviewDC(width, height)`

Creates an off-screen compatible DC for print preview rendering.

### `destroyPreviewDC(preview)`

Releases preview DC resources.

### `printToFile(outputPath, docName, renderFn)`

Prints to a file using the system's print-to-file mechanism.

### `printDocument(options, renderPageFn)`

High-level helper that shows the print dialog and calls `renderPageFn(hDC, pageNum, pageSize)` for each page. Return `false` from the callback to stop.

## Constants

### Print Dialog Flags

| Constant | Value |
|----------|-------|
| `PD_ALLPAGES` | 0 |
| `PD_SELECTION` | 1 |
| `PD_PAGENUMS` | 2 |
| `PD_PRINTTOFILE` | 32 |
| `PD_RETURNDC` | 256 |

### Device Capability Indices

| Constant | Value | Description |
|----------|-------|-------------|
| `DEVCAP_HORZRES` | 8 | Printable width in pixels |
| `DEVCAP_VERTRES` | 10 | Printable height in pixels |
| `DEVCAP_LOGPIXELSX` | 88 | Horizontal DPI |
| `DEVCAP_LOGPIXELSY` | 90 | Vertical DPI |
| `DEVCAP_PHYSICALWIDTH` | 110 | Physical page width |
| `DEVCAP_PHYSICALHEIGHT` | 111 | Physical page height |
| `DEVCAP_PHYSICALOFFSETX` | 112 | Non-printable left margin |
| `DEVCAP_PHYSICALOFFSETY` | 113 | Non-printable top margin |
