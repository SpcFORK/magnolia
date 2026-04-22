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

## Module: `lib\mermaid.oak`

- `fs` · `import(...)`
- `path` · `import(...)`
- `FlowchartLR` · `:lr`
- `FlowchartTD` · `:td`
- `FlowchartRL` · `:rl`
- `FlowchartBT` · `:bt`
- `SequenceDiagram` · `:sequence`
- `ClassDiagram` · `:classDiagram`
- `StateDiagram` · `:stateDiagram`
- `PieChart` · `:pie`
- `Gantt` · `:gantt`
- `ERDiagram` · `:erDiagram`
- `GitGraph` · `:gitGraph`
- `ShapeRect` · `:rect`
- `ShapeRound` · `:round`
- `ShapeStadium` · `:stadium`
- `ShapeCircle` · `:circle`
- `ShapeRhombus` · `:rhombus`
- `ShapeHexagon` · `:hexagon`
- `ShapeTrapezoid` · `:trapezoid`
- `ShapeDefault` · `:default`
- `EdgeArrow` · `:arrow`
- `EdgeDotted` · `:dotted`
- `EdgeThick` · `:thick`
- `EdgeOpen` · `:open`
### `graph(direction)`

> returns `:object`

### `node(g, id, label, shape)`

### `edge(g, from, to, label, style)`

### `subgraph(g, id, label, builderFn)`

### `raw(g, line)`

### `setTitle(g, t)`

### `setTheme(g, t)`

### `escapeLabel(s)`

### `renderShape(id, label, shape)`

### `renderEdge(e)`

### `renderSubgraph(sg, indent)`

### `render(g)`

### `renderFlowchart(g)`

### `seqMessage(g, from, to, label, lineType)`

### `seqNote(g, position, actor, text)`

### `seqActivate(g, actor)`

### `seqDeactivate(g, actor)`

### `seqLoop(g, label, builderFn)`

### `seqAlt(g, label, builderFn, elseLabel, elseBuilderFn)`

### `pieSlice(g, label, value)`

### `renderHTML(g)`

> returns `:string`

### `save(g, filepath)`

### `saveHTML(g, filepath)`

### `saveImage(g, filepath, opts)`

### `inferFormat(filepath)`

> returns `:string`

### `depGraph(moduleNodes, entryPath)`

### `extractImportEdges(g, fromId, node, sanitize)`

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

