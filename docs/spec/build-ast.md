# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `build-includes`

- `std` · `import(...)`
- `default` — constant
- `map` — constant
- `filter` — constant
- `reduce` — constant
- `str` · `import(...)`
- `split` — constant
- `endsWith?` — constant
- `path` · `import(...)`
- `resolve` — constant
- `SupportedExts` · `[4 items]`
### `hasOakExt?(path)`

### `resolveImportFile(basePath, checkFn)`

### `normalizeIncludePath(pathSpec)`

### `parseInclude(spec)`

> returns `:object`

### `parseIncludes(specs)`

## Module: `bundle-utils`

- `std` · `import(...)`
- `slice` — constant
- `reduce` — constant
### `pairPrefix(left, right)`

### `commonPrefix(paths)`

> returns `:string`

### `normalizeModulePath(path, allPaths)`

## Module: `lib\build-ast.oak`

- `std` · `import(...)`
- `slice` — constant
- `map` — constant
- `each` — constant
- `filter` — constant
- `append` — constant
- `contains?` — constant
- `reduce` — constant
- `sort` · `import(...)`
- `sort!` — constant
- `path` · `import(...)`
- `dir` — constant
- `resolve` — constant
- `bundleUtils` · `import(...)`
- `buildIncludes` · `import(...)`
- `SupportedExts` — constant
### `wrapModule(block)`

> returns `:object`

### `bundleCommonPrefix(paths)`

> returns `:string`

### `wrapBundle(modules, entryModuleName, includes, importCallNode)`

> returns `:object`

## Module: `path`

### `abs?(path)`

### `rel?(path)`

### `_lastSlash(path)`

> returns `:int`

### `dir(path)`

### `base(path)`

### `cut(path)`

> returns `:list`

### `clean(path)`

> returns `:string`

### `join(parts...)`

### `split(path)`

> returns `:list`

### `resolve(path, base)`

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

