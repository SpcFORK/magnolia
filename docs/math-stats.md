# Math Statistics (math-stats)

## Overview

`math-stats` provides summary statistics and numeric utility functions. It is used internally by `libmath` and can also be imported directly.

## Import

```oak
stats := import('math-stats')
// or destructure
{ mean: mean, median: median, stddev: stddev, clamp: clamp, round: round } := import('math-stats')
```

## Aggregation Functions

### `sum(xs...)`

Returns the sum of all arguments. Accepts a list or individual values.

```oak
sum(1, 2, 3)      // => 6
sum([1, 2, 3]...) // => 6
```

### `prod(xs...)`

Returns the product of all arguments.

```oak
prod(2, 3, 4) // => 24
```

### `min(xs...)`

Returns the minimum value among all arguments.

```oak
min(3, 1, 4, 1, 5) // => 1
```

### `max(xs...)`

Returns the maximum value among all arguments.

```oak
max(3, 1, 4, 1, 5) // => 5
```

## Statistical Functions

### `mean(xs)`

Returns the arithmetic mean of list `xs`. Returns `?` for an empty list.

```oak
mean([1, 2, 3, 4])  // => 2.5
mean([])             // => ?
```

### `median(xs)`

Returns the median of list `xs` (sorts the list first). For even-length lists returns the average of the two middle values. Returns `?` for an empty list.

```oak
median([3, 1, 4, 1, 5]) // => 3
median([1, 2, 3, 4])     // => 2.5
```

### `stddev(xs)`

Returns the population standard deviation of list `xs`. Returns `?` for an empty list.

```oak
stddev([2, 4, 4, 4, 5, 5, 7, 9]) // => 2.0
```

## Numeric Utilities

### `clamp(x, a, b)`

Returns `x` clamped to the range `[a, b]`.

```oak
clamp(10, 0, 5)  // => 5
clamp(-3, 0, 5)  // => 0
clamp(3, 0, 5)   // => 3
```

### `round(n, decimals?)`

Rounds `n` to `decimals` decimal places (default: 0). Uses round-half-away-from-zero.

```oak
round(3.14159, 2) // => 3.14
round(2.5)        // => 3
round(-2.5)       // => -3
round(123.456, 0) // => 123
```

## Parallel Batch Operations

### `pbatchMean(datasets)`

Computes mean of multiple datasets in parallel.

```oak
pbatchMean([[1, 2, 3], [4, 5, 6]])  // => [2, 5]
```

### `pbatchStddev(datasets)`

Computes standard deviation of multiple datasets in parallel.

```oak
pbatchStddev([[1, 2, 3], [4, 5, 6]])
```
