# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `json`

### `esc(c)`

> returns `:string`

### `uEscape(c)`

> returns `:string`

### `uParse(uc)`

### `escape(s)`

### `serialize(c)`

> returns `:string`

### `cs Reader(s)`

> returns `:object`

### `parseNull(r)`

> returns `?`

### `parseString(r)`

### `parseNumber(r)`

> returns `:atom`

### `parseTrue(r)`

> returns `:bool`

### `parseFalse(r)`

> returns `:bool`

### `parseList(r)`

### `parseObject(r)`

### `_parseReader(r)`

> returns `:atom`

### `parse(s)`

## Module: `lib\p2p.oak`

- `thread` · `import(...)`
- `json` · `import(...)`
- `ws` · `import(...)`
### `_noop(_)`

> returns `?`

### `_packet(kind, body)`

### `_sendPacket(socket, kind, body)`

### `_terminalPeerError?(evt)`

> returns `:bool`

### `_decodePacket(evt)`

> returns `:object`

### `_findPeer(peers, peerId)`

### `_peerSummaries(peers, excludeId)`

### `_peerSockets(peers, excludeId)`

### `_fanout(state, mutex, excludeId, kind, body)`

### `_removePeer(state, mutex, peerId)`

### `_sendDirect(state, mutex, toPeerId, fromPeerId, channel, payload)`

### `_peerEvent(packetType, body)`

> returns `:object`

### `_hostError(onEvent, peerId, error)`

### `_hostLoop(state, mutex, peerId, socket, onEvent)`

### `_acceptPeer(state, mutex, host, path, socket, req, onEvent)`

### `Host(host, path, onEvent)`

> returns `:object`

### `join(url, peerId, onEvent, meta)`

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

## Module: `websocket`

### `connect(url, headers)`

### `send(socket, data, opcode)`

### `recv(socket)`

### `close(socket, code, reason)`

### `listen(host, path, onConnect)`

### `Socket(socket)`

> returns `:object`

- `Opcode` · `{5 entries}`
- `CloseCode` · `{8 entries}`
