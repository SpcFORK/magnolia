# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\msgpack.oak`

### `bytes(parts)`

### `packet(data)`

### `fromPacket(parts)`

### `Binary(data)`

> returns `:object`

### `binary?(value)`

> returns `:bool`

### `_u8(n)`

> returns `:list`

### `_u16be(n)`

> returns `:list`

### `_u32be(n)`

> returns `:list`

### `_encodeString(s)`

> returns `:atom`

### `_encodeBinary(data)`

> returns `:atom`

### `_encodeInt(n)`

> returns `:list`

### `_encodeList(xs)`

> returns `:atom`

### `_encodeObject(obj)`

> returns `:atom`

### `_encodeToPacket(value)`

### `_f64ToPacket(f)`

> returns `:list`

### `_f32ToPacket(f)`

> returns `:list`

### `_encodeFloat64(f)`

### `encodeFloat64(f)`

### `encodeFloat32(f)`

### `_readU8(pkt, i)`

### `_readU16(pkt, i)`

> returns `:bool`

### `_readU32(pkt, i)`

> returns `:bool`

### `_readI8(pkt, i)`

> returns `:atom`

### `_readI16(pkt, i)`

> returns `:atom`

### `_readI32(pkt, i)`

> returns `:atom`

### `_readBits64(pkt, i)`

### `_ieee754f64(bs)`

> returns `:float`

### `_ieee754f32(bs)`

> returns `:float`

### `_decode(pkt, i)`

> returns `:atom`

### `_decodeStr(pkt, i, n)`

> returns `:atom`

### `_decodeBin(pkt, i, n)`

> returns `:atom`

### `_decodeArray(pkt, i, n)`

### `_decodeMap(pkt, i, n)`

### `_slicePkt(pkt, a, b)`

### `parsePacket(pkt)`

> returns `:atom`

### `parse(data)`

### `serializeSafe(value)`

> returns `:atom`

### `serialize(value)`

### `serializePacket(value)`

> returns `:atom`

### `pbatchSerialize(values)`

### `pbatchParse(packets)`

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

## Module: `str`

### `checkRange(lo, hi)`

> **thunk** returns `:function`

### `upper?(c)`

> returns `:bool`

### `lower?(c)`

> returns `:bool`

### `digit?(c)`

> returns `:bool`

### `space?(c)`

> returns `:bool`

### `letter?(c)`

> returns `:bool`

### `word?(c)`

> returns `:bool`

### `join(strings, joiner)`

> returns `:string`

### `startsWith?(s, prefix)`

### `endsWith?(s, suffix)`

### `_matchesAt?(s, substr, idx)`

> returns `:bool`

### `indexOf(s, substr)`

### `rindexOf(s, substr)`

### `contains?(s, substr)`

### `cut(s, sep)`

> returns `:list`

### `lower(s)`

### `upper(s)`

### `_replaceNonEmpty(s, old, new)`

### `replace(s, old, new)`

### `_splitNonEmpty(s, sep)`

### `split(s, sep)`

### `_extend(pad, n)`

### `padStart(s, n, pad)`

### `padEnd(s, n, pad)`

### `_trimStartSpace(s)`

### `_trimStartNonEmpty(s, prefix)`

### `trimStart(s, prefix)`

### `_trimEndSpace(s)`

### `_trimEndNonEmpty(s, suffix)`

### `trimEnd(s, suffix)`

### `trim(s, part)`

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

