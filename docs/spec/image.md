# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `image-bmp`

### `bytes(parts)`

### `intToBytes(n, width)`

### `hexsplit(n)`

### `pixelBytes(pixel)`

### `bmp(width, height, pixels)`

## Module: `image-ico`

### `bytes(parts)`

### `intToBytes(n, width)`

### `hexsplit(n)`

### `pixelBytes(pixel)`

### `ico(width, height, pixels)`

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

## Module: `image-qoi`

### `bytes(parts)`

### `clampByte(n)`

> returns `:int`

### `_mod(a, b)`

### `intToBytesBE(n, width)`

- `_OP_INDEX` · `0`
- `_OP_DIFF` · `64`
- `_OP_LUMA` · `128`
- `_OP_RUN` · `192`
- `_OP_RGB` · `254`
- `_OP_RGBA` · `255`
### `_hash(r, g, b, a)`

### `_signed(v)`

### `_mkIndex()`

### `qoi(width, height, pixels, channels)`

### `decodeQOI(data)`

> returns `?`

## Module: `image-tga`

### `bytes(parts)`

### `intToBytes(n, width)`

### `clampByte(n)`

> returns `:int`

### `_mkBuf()`

> returns `:object`

### `_writeHeader(appd, w, h, bpp, imageType)`

### `_writePixel(appd, px, bpp)`

### `_pixelEqual(a, b, bpp, bytesPerPixel)`

### `tga(width, height, pixels, bpp)`

### `tgaRLE(width, height, pixels, bpp)`

### `decodeTGA(data)`

> returns `:object`

## Module: `lib\image.oak`

- `bmp` · `import(...)`
- `ico` · `import(...)`
- `ppm` · `import(...)`
- `tga` · `import(...)`
- `qoi` · `import(...)`
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

