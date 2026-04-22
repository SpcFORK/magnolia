# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `ai-linalg`

### `matZeros(rows, cols)`

### `matIdentity(n)`

### `matTranspose(mat)`

> returns `:list`

### `matMul(a, b)`

### `matAdd(a, b)`

### `matScale(mat, scalar)`

### `matDiag(mat)`

### `matFromDiag(diag)`

### `matTrace(mat)`

### `outerProduct(v1, v2)`

### `matHadamard(a, b)`

### `matFlatten(mat)`

### `matReshape(flat, rows, cols)`

### `pmatMul(a, b)`

## Module: `ai-vec`

### `dotProduct(v1, v2)`

### `magnitude(v)`

### `normalize(v)`

> returns `?`

### `cosineSimilarity(v1, v2)`

> returns `?`

### `euclideanDistance(v1, v2)`

### `manhattanDistance(v1, v2)`

### `chebyshevDistance(v1, v2)`

> returns `:int`

### `hammingDistance(v1, v2)`

### `vecAdd(v1, v2)`

### `vecSub(v1, v2)`

### `vecScale(v, scalar)`

### `minMaxScale(data)`

> returns `:object`

### `zScore(data)`

> returns `:object`

### `mean(data)`

> returns `:int`

### `variance(data)`

> returns `:int`

### `stddev(data)`

### `entropy(probabilities)`

### `ema(data, alpha)`

> returns `:list`

### `pbatchNormalize(vectors)`

### `pbatchCosineSimilarity(vectors, query)`

### `cosineSimilarityMatrix(vectors)`

### `distanceMatrix(vectors, distFn)`

## Module: `lib\ai-data.oak`

### `ModelConfig(params)`

### `oneHotEncode(label, numClasses)`

### `oneHotDecode(v)`

> returns `?`

### `shuffleIndices(n, seed)`

### `trainTestSplit(data, labels, ratio, seed)`

> returns `:object`

### `kFoldIndices(n, k)`

### `batchify(items, batchSize)`

### `padSequence(sequence, targetLength, padToken)`

### `labelEncode(labels)`

### `polynomialFeatures(features, degree)`

### `accuracy(predicted, actual)`

> returns `:int`

### `precisionScore(predicted, actual, positiveLabel)`

### `recallScore(predicted, actual, positiveLabel)`

### `f1Score(predicted, actual, positiveLabel)`

> returns `:int`

### `confusionMatrix(predicted, actual, numClasses)`

### `rSquared(actual, predicted)`

> returns `:int`

### `addNoise(v, scale, seed)`

### `mixup(sample1, sample2, lambda)`

### `quantize(values, numLevels)`

### `dequantize(levels, numLevels)`

### `trainLoop(params, data, batchSize, epochs, updateFn)`

## Module: `math`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

### `orient(x0, y0, x1, y1)`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

### `median(xs)`

### `stddev(xs)`

### `round(n, decimals)`

## Module: `math-base`

- `Pi` · `3.141592653589793`
- `E` · `2.718281828459045`
### `sign(n)`

> returns `:int`

### `abs(n)`

### `sqrt(n)`

## Module: `math-geo`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

> returns `:list`

### `orient(x0, y0, x1, y1)`

> returns `:int`

## Module: `math-stats`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

> returns `?`

### `median(xs)`

> returns `?`

### `stddev(xs)`

### `pbatchMean(datasets)`

### `pbatchStddev(datasets)`

### `round(n, decimals)`

## Module: `sort`

### `sort!(xs, pred)`

### `sort(xs, pred)`

### `_mergeSorted(a, b, pred)`

### `psort(xs, pred)`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

