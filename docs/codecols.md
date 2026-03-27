# Column Counter Library (codecols)

## Overview

`libcodecols` provides column-counting and histogram rendering for source code analysis. Designed after Rasmus Andersson's [linelen_hist.sh](https://gist.github.com/rsms/36bda3b5c8ab83d951e45ed788a184f4). It computes per-line column widths, frequency distributions, summary statistics (including percentiles), and renders Unicode histogram charts.

## Import

```oak
codecols := import('codecols')
```

## Quick Start

```oak
codecols := import('codecols')

// All-in-one analysis
result := codecols.analyze(readFile('main.c'))
codecols.renderHistogram(result.freqs, {})
println('mean:', result.stats.mean)
println('p90:', result.stats.p90)

// Or use report() for a formatted summary
codecols.report(readFile('main.c'), { histoWidth: 40 })
```

## Functions

### `columnCounts(lines)`

Takes a list of line strings and returns a list of column widths. Empty lines are excluded. Each width is rounded up to the nearest even number.

**Parameters:**
- `lines` — list of strings (one per line)

**Returns:** list of integers

```oak
codecols.columnCounts(['hello', '', 'hi'])
// => [6, 2]
```

### `frequencies(cols)`

Takes a list of column counts and returns a frequency map `{ columnWidth: count }`.

**Parameters:**
- `cols` — list of integers from `columnCounts()`

**Returns:** object mapping column width strings to occurrence counts

```oak
codecols.frequencies([4, 4, 6, 8])
// => { '4': 2, '6': 1, '8': 1 }
```

### `stats(cols)`

Computes summary statistics for a list of column counts.

**Parameters:**
- `cols` — list of integers from `columnCounts()`

**Returns:**
```oak
{
    min: 4           // smallest column count
    max: 120         // largest column count
    mean: 42.5       // average (rounded to 2 decimal places)
    median: 38       // median value
    total: 8500      // sum of all column counts
    lineCount: 200   // number of non-empty lines
    p75: 60          // 75th percentile
    p90: 80          // 90th percentile
    p95: 100         // 95th percentile
    p99: 118         // 99th percentile
}
```

### `percentile(sortedVals, p)`

Computes the p-th percentile of a pre-sorted list of numbers.

**Parameters:**
- `sortedVals` — sorted list of numbers
- `p` — percentile in range `[0, 100]`

**Returns:** number

```oak
codecols.percentile([1, 2, 3, 4, 5, 6, 7, 8, 9, 10], 90)
// => 9
```

### `histo(n)`

Renders a Unicode histogram bar of a given length using 1/8th block character precision.

**Parameters:**
- `n` — bar length in 1/8th-character units (e.g. 24 = 3 full blocks)

**Returns:** string of Unicode block characters

```oak
codecols.histo(24) // => '███'
codecols.histo(5)  // => '▋'
```

### `renderHistogram(freqs, opts)`

Prints a formatted histogram table to stdout.

**Parameters:**
- `freqs` — frequency map from `frequencies()`
- `opts` — optional config object:
  - `maxCols` — maximum column number to display (default: auto from data)
  - `histoWidth` — maximum bar width in characters (default: 60)
  - `step` — column grouping step size (default: 2)

```oak
codecols.renderHistogram(freqs, { histoWidth: 40, step: 4 })
```

### `formatHistogram(freqs, opts)`

Returns the histogram as a list of strings instead of printing to stdout. Takes the same parameters as `renderHistogram()`.

**Returns:** list of strings (header + one per row)

```oak
lines := codecols.formatHistogram(freqs, {})
```

### `analyze(text)`

Performs a full column analysis on raw source text. Combines `columnCounts()`, `frequencies()`, and `stats()`.

**Parameters:**
- `text` — raw source code string

**Returns:**
```oak
{
    cols: [...]      // list of per-line column counts
    freqs: { ... }   // frequency map
    stats: { ... }   // summary statistics
}
```

### `report(text, opts)`

All-in-one: analyzes source text, prints the histogram, and prints a summary line. Returns the analysis result.

**Parameters:**
- `text` — raw source code string
- `opts` — histogram options (same as `renderHistogram()`)

**Returns:** same as `analyze()`
