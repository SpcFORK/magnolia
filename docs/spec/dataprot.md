# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\dataprot.oak`

### `_byte?(n)`

> returns `:bool`

### `_byteList(data)`

### `_bitValue(v)`

> returns `:int`

### `_bitVector(data)`

### `_bitMatrix(matrix)`

### `_popcount8(n)`

### `_targetParity(odd)`

> returns `:int`

### `_crc16Byte(crc, b, poly)`

### `_crc32Byte(crc, b, poly)`

### `_rowParity(row, bits)`

### `_bitColumn(rows, idx)`

### `_sameBits(a, b)`

> returns `:bool`

### `_flipBit(bits, idx)`

### `_formatBits(bits, template)`

### `parity(data)`

> returns `:atom`

### `parityBit(data, odd)`

> returns `:atom`

### `parityValid?(data, checkBit, odd)`

> returns `:atom`

### `xorChecksum(data)`

> returns `:atom`

### `sumChecksum8(data)`

> returns `:atom`

### `sumChecksum16(data)`

> returns `:atom`

### `sumChecksum32(data)`

> returns `:atom`

### `crc16Ccitt(data, seed, poly)`

> returns `:atom`

### `crc32(data, seed, poly, finalXor)`

> returns `:atom`

### `hammingDistance(a, b)`

> returns `:atom`

### `ldpcSyndrome(word, parityMatrix)`

> returns `:atom`

### `ldpcValid?(word, parityMatrix)`

> returns `:atom`

### `ldpcCheck(word, parityMatrix)`

> returns `:atom`

### `ldpcCandidates(word, parityMatrix)`

> returns `:atom`

### `ldpcCorrect(word, parityMatrix)`

> returns `:atom`

### `pbatchCrc32(payloads)`

### `pbatchLdpcCheck(words, parityMatrix)`

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

