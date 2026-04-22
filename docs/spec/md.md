# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

## Module: `fs`

- `ReadBufSize` · `1048576`
### `readFileSync(path)`

> returns `?`

### `readFileAsync(path, withFile)`

### `readFile(path, withFile)`

### `writeFileSyncWithFlag(path, file, flag)`

> returns `?`

### `writeFileAsyncWithFlag(path, file, flag, withEnd)`

### `writeFile(path, file, withEnd)`

### `appendFile(path, file, withEnd)`

### `statFileSync(path)`

> returns `?`

### `statFileAsync(path, withStat)`

### `statFile(path, withStat)`

### `listFilesSync(path)`

> returns `?`

### `listFilesAsync(path, withFiles)`

### `listFiles(path, withFiles)`

### `readFiles(paths)`

### `readFilesConcurrent(paths, maxOpen)`

### `writeFiles(pairs)`

### `statFiles(paths)`

## Module: `http`

- `sort` · `import(...)`
- `json` · `import(...)`
### `queryEncode(params)`

### `queryDecode(params)`

### `_encodeChar(uri?)`

> **thunk** returns `:function`

### `percentEncode(s)`

### `percentEncodeURI(s)`

### `_hex?(c)`

> returns `:bool`

### `percentDecode(s)`

### `cs Router()`

> returns `:object`

- `MimeTypes` · `{17 entries}`
### `mimeForPath(path)`

- `NotFound` · `{2 entries}`
- `MethodNotAllowed` · `{2 entries}`
### `_hdr(attrs)`

### `cs Server()`

> returns `:object`

### `handleStaticUnsafe(path)`

> **thunk** returns `:function`

### `handleStatic(path)`

### `pbatchQueryEncode(paramSets)`

### `pbatchQueryDecode(queryStrings)`

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

## Module: `lib\md.oak`

### `cs Reader(s)`

> returns `:object`

### `uword?(c)`

> returns `:bool`

### `tokenizeText(line)`

### `unifyTextNodes(nodes, joiner)`

### `parseText(tokens)`

### `uListItemLine?(line)`

> returns `:bool`

### `oListItemLine?(line)`

> returns `:bool`

### `listItemLine?(line)`

> returns `:bool`

### `tableLine?(line)`

> returns `:bool`

### `tableSepLine?(line)`

### `trimUListGetLevel(reader)`

### `trimOListGetLevel(reader)`

### `lineNodeType(line)`

> returns `?`

### `parse(text)`

### `parseDoc(lineReader)`

### `parseHeader(nodeType, lineReader)`

> returns `:object`

### `parseBlockQuote(lineReader)`

> returns `:object`

### `parseCodeBlock(lineReader)`

> returns `:object`

### `parseRawHTML(lineReader)`

> returns `:object`

### `parseList(lineReader, listType)`

> returns `:object`

### `parseTableRow(line)`

### `parseTableAlign(sepLine)`

### `parseTable(lineReader)`

> returns `:object`

### `parseParagraph(lineReader)`

> returns `:object`

### `compile(nodes)`

### `wrap(tag, node)`

> returns `:string`

### `sanitizeAttr(attr)`

### `sanitizeURL(url)`

> returns `:string`

### `compileNode(node)`

### `transform(text)`

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

