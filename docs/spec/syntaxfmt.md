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

## Module: `lib\syntaxfmt.oak`

- `syntax` · `import(...)`
### `formatContent(content)`

### `formatFile(path)`

> returns `:object`

### `formatFileInPlace(path)`

> returns `:object`

### `formatFileWithDiff(path)`

> returns `:object`

### `getChangedOakFiles()`

> returns `:object`

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

## Module: `syntax`

## Module: `syntax-macros`

- `std` · `import(...)`
- `default` — constant
- `clone` — constant
- `map` — constant
### `Macro(expand)`

> returns `:object`

### `macro?(x)`

> returns `:bool`

### `expandMacros(ast, macros)`

### `parseWithMacros(text, macros)`

## Module: `syntax-parse`

- `std` · `import(...)`
- `fromHex` — constant
- `slice` — constant
- `append` — constant
- `last` — constant
- `filter` — constant
- `map` — constant
- `each` — constant
- `str` · `import(...)`
- `strContains?` — constant
- `fmt` · `import(...)`
- `format` — constant
### `cloneNameSet(set)`

### `addPatternBindings(shadowed, node)`

> returns `:bool`

### `rewriteClassSugarAssignmentLeft(node, visibleFields, allFields, shadowed, selfName, isLocal)`

### `rewriteClassSugarNode(node, visibleFields, allFields, shadowed, selfName)`

### `classBodyFromAssignmentBlock(body, reservedNames)`

> returns `:list`

### `wrapBodyWithSelfVar(body, reservedNames)`

> returns `:object`

### `Parser(tokens)`

> returns `:object`

### `parse(text)`

## Module: `syntax-print`

- `std` · `import(...)`
- `default` — constant
- `range` — constant
- `take` — constant
- `first` — constant
- `each` — constant
- `map` — constant
- `str` · `import(...)`
- `cut` — constant
- `join` — constant
- `trimStart` — constant
- `trim` — constant
- `math` · `import(...)`
- `min` — constant
- `max` — constant
- `fmt` · `import(...)`
- `printf` — constant
### `Printer(tokens)`

> returns `:object`

### `print(text)`

## Module: `syntax-tokenize`

- `std` · `import(...)`
- `contains?` — constant
- `str` · `import(...)`
- `digit?` — constant
- `word?` — constant
- `space?` — constant
- `startsWith?` — constant
- `trimEnd` — constant
- `trim` — constant
- `fmt` · `import(...)`
- `format` — constant
### `shebang?(text)`

### `renderPos(pos)`

> returns `:string`

### `renderToken(token)`

### `Tokenizer(source)`

> returns `:object`

### `tokenize(text)`

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

