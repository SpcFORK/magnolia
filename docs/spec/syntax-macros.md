# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

## Module: `lib\syntax-macros.oak`

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

