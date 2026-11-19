# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `image-ppm`

### `clampByte(n)`

> returns `:int`

### `_mkHeader(magic, width, height, maxVal)`

### `_mkBuf(header)`

> returns `:object`

### `ppm(width, height, pixels)`

### `ppmPlain(width, height, pixels)`

### `pgm(width, height, pixels)`

### `pgmPlain(width, height, pixels)`

### `pbm(width, height, pixels)`

### `pbmPlain(width, height, pixels)`

### `_skipWS(data, idx)`

### `_readToken(data, idx)`

### `_readHeader(data)`

> returns `:object`

### `decodePPM(data)`

> returns `:object`

### `decodePGM(data)`

> returns `:object`

## Module: `lib\rmqr.oak`

- `ppm` · `import(...)`
- `M` · `:M`
- `H` · `:H`
- `MinimizeWidth` · `:minimizeWidth`
- `MinimizeHeight` · `:minimizeHeight`
- `Balanced` · `:balanced`
- `Numeric` · `:numeric`
- `Alphanumeric` · `:alphanumeric`
- `Byte` · `:byte`
- `Kanji` · `:kanji`
- `_Inf` · `1000000`
- `_QuietZone` · `2`
- `_AlphaChars` · `'0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ $%*+-./:'`
- `_Versions` · `[32 items]`
- `_AlignmentXs` · `{6 entries}`
### `_version(name)`

### `_modeIndex(mode)`

> returns `:int`

### `_bitString(n, width)`

### `_bitLength(n)`

> returns `:int`

### `_bch(data)`

### `_alphaValue(c)`

### `_numeric?(c)`

### `_modeValid?(mode, data)`

### `_grid(width, height, value)`

### `_dataBits(info, ecc)`

### `_blockDefinitions(info, ecc)`

### `_dataCodewordCount(defs)`

### `_numericPayload(data)`

### `_alphanumericPayload(data)`

### `_bytePayload(data)`

### `_segmentBits(segment, info)`

> returns `?`

### `_segmentsBits(segments, info)`

### `_bitsToDataCodewords(bits, count)`

### `_gfMultiply(a, b)`

### `_gfPow2(n)`

### `_generator(degree)`

### `_reedSolomon(data, degree)`

### `_blocks(data, definitions)`

### `_interleave(blocks, key)`

### `_finalCodewords(blocks)`

### `_set(grid, x, y, value)`

### `_get(grid, x, y)`

### `_putFinder(grid, width, height)`

### `_putFinderSub(grid, width, height)`

### `_putCornerFinders(grid, width, height)`

### `_alignmentCenters(width)`

### `_putAlignment(grid, width, height)`

### `_putTiming(grid, width, height)`

### `_putFormat(grid, width, height, format)`

### `_placeData(grid, width, height, words, remainderBits)`

### `_applyMask(grid, area, width, height)`

### `segment(data, mode)`

> returns `:object`

### `make(versionName, ecc, segments)`

> returns `?`

### `_autoSegments(data, info, ecc)`

> returns `?`

### `fit(data, ecc, strategy)`

### `matrix(code, quietZone)`

> returns `?`

### `ascii(code, quietZone)`

> returns `?`

### `pbm(code, scale, quietZone)`

> returns `?`

### `versions()`

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

- `__std_builtin_int` — constant
- `__std_builtin_float` — constant
- `__std_builtin_string` — constant
- `__std_builtin_type` — constant
- `__std_builtin_len` — constant
- `__std_builtin_char` — constant
- `__std_builtin_codepoint` — constant
- `__std_builtin_print` — constant
- `__std_builtin_keys` — constant
- `__std_builtin_atom` — constant
- `__std_builtin_name` — constant
- `__std_builtin_rand` — constant
- `__std_builtin_srand` — constant
- `__std_builtin_input` — constant
- `__std_builtin_exit` — constant
- `__std_builtin_args` — constant
- `__std_builtin_addr` — constant
- `__std_builtin_bits` — constant
- `__std_builtin_memread` — constant
- `__std_builtin_memwrite` — constant
- `__std_builtin_utf16` — constant
- `__std_builtin_go` — constant
- `__std_builtin_wait` — constant
- `__std_builtin_exec` — constant
- `__std_builtin_env` — constant
- `__std_builtin_write` — constant
- `__std_builtin_read` — constant
- `__std_builtin_interpreter` — constant
- `__std_builtin_pointer` — constant
- `__std_builtin_log` — constant
- `__std_builtin_time` — constant
- `__std_builtin_nanotime` — constant
- `int` — constant
- `float` — constant
- `string` — constant
- `type` — constant
- `len` — constant
- `print` — constant
- `time` — constant
- `nanotime` — constant
- `char` — constant
- `codepoint` — constant
- `keys` — constant
- `atom` — constant
- `name` — constant
- `rand` — constant
- `srand` — constant
- `input` — constant
- `args` — constant
- `addr` — constant
- `bits` — constant
- `memread` — constant
- `memwrite` — constant
- `utf16` — constant
- `exec` — constant
- `env` — constant
- `write` — constant
- `read` — constant
- `interpreter` — constant
- `pointer` — constant
- `log` — constant
- `wait` — constant
- `__std_builtin_sin` — constant
- `__std_builtin_cos` — constant
- `__std_builtin_tan` — constant
- `__std_builtin_asin` — constant
- `__std_builtin_acos` — constant
- `__std_builtin_atan` — constant
- `__std_builtin_pow` — constant
- `sin` — constant
- `cos` — constant
- `tan` — constant
- `asin` — constant
- `acos` — constant
- `atan` — constant
- `pow` — constant
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

