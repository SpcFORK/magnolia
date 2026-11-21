# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `Virtual`

- `builtinKeys` — constant
- `syntax` · `import(...)`
- `path` · `import(...)`
- `fs` · `import(...)`
- `syntax` · `import(...)`
- `vbc` · `import(...)`
### `cs VirtualMachine()`

> returns `:object`

### `createVM()`

### `cs StandardVM()`

### `createStandardVM()`

### `cs VirtualContext(initialScope)`

### `createVirtualContext(initialScope)`

### `cs VirtualFS(initialFiles)`

> returns `:object`

### `createVirtualFS(initialFiles)`

### `createVirtualFSWithFS(basePath, patterns)`

## Module: `Virtual-Bytecode`

- `std` · `import(...)`
- `syntax` · `import(...)`
- `wasmVM` · `import(...)`
- `fmt` · `import(...)`
- `printf` — constant
- `str` · `import(...)`
- `join` — constant
- `strSplit` — constant
- `builtinKeys` — constant
- `builtinValues` — constant
- `__jsReflectApply` — constant
- `__jsSetTimeout` — constant
- `__jsTryCall` — constant
### `vmDebugLog(msg)`

- `MODE_WASM` · `:wasm`
- `MODE_GO` · `:go`
- `CONST_STRING` · `0`
- `CONST_ATOM` · `1`
- `CONST_FLOAT` · `2`
- `I32_SIGN` · `2147483648`
- `I32_WRAP` · `4294967296`
- `_STACK_CAPACITY` · `2048`
### `rangeN(n)`

> returns `:list`

### `bytesFromString(s)`

### `bytesFromAny(v)`

### `readU16(bs, at)`

> returns `:bool`

### `readU32(bs, at)`

> returns `:bool`

### `readI32(bs, at)`

### `bytesToString(bs, start, count)`

> returns `:string`

### `normalizeConstEntry(entry, mode)`

> returns `:object`

### `decodeConstantPoolBytes(raw)`

> returns `:list`

### `decodeFunctionTableBytes(raw)`

> returns `:list`

### `decodeNameListBytes(raw)`

> returns `:list`

### `parseMBCBinary(raw)`

> returns `?`

### `detectMode(raw)`

### `normalizeFunctions(rawFunctions)`

### `normalizeChunk(raw)`

> returns `:object`

### `defaultOpcodes(mode)`

> returns `:object`

### `stackPush(vm, v)`

### `stackPop(vm)`

> returns `?`

### `stackPeek(vm)`

> returns `?`

### `currentFrameIndex(vm)`

### `currentFrame(vm)`

### `findTopLocal(vm, name)`

### `setCurrentFrame(vm, frame)`

### `currentScope(vm)`

### `scopeGetByName(scope, name)`

> returns `?`

### `scopeSetByName(scope, name, val)`

> returns `:bool`

### `scopeAtDepth(scope, depth)`

### `scopeGetByDepthSlot(scope, depth, slot)`

> returns `?`

### `scopeSetByDepthSlot(scope, depth, slot, val)`

> returns `:bool`

### `fetchU8(vm)`

> returns `:bool`

### `fetchU16(vm)`

### `fetchI32(vm)`

### `constValueAt(vm, idx)`

> returns `?`

### `keyToString(k)`

### `vmGetProp(obj, key)`

> returns `?`

### `vmSetProp(obj, key, val)`

### `isTruthy(v)`

### `deepEq(a, b)`

> returns `:bool`

### `vmAdd(a, b)`

### `vmSub(a, b)`

### `vmMul(a, b)`

### `vmDiv(a, b)`

> returns `?`

### `vmMod(a, b)`

> returns `?`

### `vmPow(a, b)`

> returns `?`

### `vmCompare(op, a, b)`

> returns `:bool`

### `vmConcat(a, b)`

### `builtinAppend(x, y)`

### `callBuiltin(idx, args)`

> returns `?`

### `makeFrame(vm, closure, args, returnPC)`

> returns `:object`

### `invokeCallable(vm, callee, args)`

### `wrapClosureArgs(vm, args)`

### `popArgs(vm, arity)`

> returns `:list`

### `runChunk(rawChunk, opts)`

> returns `?`

### `compileAst(ast)`

### `compileSource(source)`

### `runSource(source, opts)`

- `run` — constant
- `_defaultImportModules` · `{5 entries}`
### `makeDefaultImportFn()`

> **thunk** returns `:function`

- `defaultImportFn` · `makeDefaultImportFn(...)`
## Module: `ast-transform`

- `std` · `import(...)`
- `clone` — constant
- `entries` — constant
### `wrapBlock(nodes)`

> returns `:object`

### `wrapModule(block)`

> returns `:object`

### `formatIdentForWeb(name, key)`

> returns `:string`

### `formatIdentForOak(name)`

> returns `:string`

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

## Module: `lib\pack.oak`

- `std` · `import(...)`
- `default` — constant
- `append` — constant
- `slice` — constant
- `map` — constant
- `each` — constant
- `merge` — constant
- `keys` — constant
- `fmt` · `import(...)`
- `printf` — constant
- `format` — constant
- `fs` · `import(...)`
- `readFile` — constant
- `writeFile` — constant
- `statFile` — constant
- `listFiles` — constant
- `str` · `import(...)`
- `join` — constant
- `split` — constant
- `replace` — constant
- `startsWith?` — constant
- `trimStart` — constant
- `trimEnd` — constant
- `path` · `import(...)`
- `dir` — constant
- `resolve` — constant
- `base` — constant
- `Virtual` · `import(...)`
- `shell` · `import(...)`
- `packUtils` · `import(...)`
- `transpile` · `import(...)`
### `requireOpt(flag, value, abort)`

### `buildArgs(entry, output, includes)`

### `windowsDrivePath?(part)`

> returns `:bool`

### `normalizeVFSTargetName(name)`

### `parseVFSIncludeSpec(spec)`

> returns `:list`

### `joinVFSPath(basePath, relPath)`

### `collectVFSFiles(includeSpecs, log)`

- `Configured` · `?`
### `configure(config)`

### `pack()`

### `run(config)`

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

- `__math_builtin_sin` — constant
- `__math_builtin_cos` — constant
- `__math_builtin_tan` — constant
- `__math_builtin_asin` — constant
- `__math_builtin_acos` — constant
- `__math_builtin_atan` — constant
- `__math_builtin_pow` — constant
- `__math_builtin_log` — constant
- `sin` — constant
- `cos` — constant
- `tan` — constant
- `asin` — constant
- `acos` — constant
- `atan` — constant
- `pow` — constant
- `log` — constant
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

## Module: `msgpack`

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

## Module: `pack-utils`

- `MagicBytes` · `'oak �'`
- `MaxBundleSizeLen` · `24`
### `encodeBundleSize(bundleSize)`

### `packBundle(executable, oakBundle)`

- `VFSMagicBytes` · `'vfs    '`
- `MaxVFSSizeLen` · `24`
### `encodeVFSSize(vfsSize)`

### `serializeVFS(vfsFiles)`

> returns `:string`

### `deserializeVFS(vfsData)`

> returns `:object`

### `packBundleWithVFS(executable, oakBundle, vfsFiles)`

### `extractVFSFromPacked(packedData)`

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

## Module: `shell`

- `path` · `import(...)`
- `fs` · `import(...)`
### `tokenizeLine(line)`

### `splitPipeline(tokens)`

### `parseRedirections(tokens)`

### `globMatch?(pattern, name)`

### `hasGlobChars?(s)`

> returns `:bool`

### `normalizePath(currentDir, targetPath)`

### `cs Shell(vfs?)`

> returns `:object`

### `createShell(vfs?)`

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

### `pmapWithPool(list, fnToRun, pool)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

## Module: `transpile`

- `std` · `import(...)`
- `map` — constant
- `each` — constant
- `clone` — constant
- `merge` — constant
- `filter` — constant
- `reduce` — constant
- `fmt` · `import(...)`
- `printf` — constant
- `format` — constant
- `transform` · `import(...)`
- `TranspileRegistry` · `[]`
- `TranspileConfig` · `{3 entries}`
### `configure(config)`

### `registerTranspiler(transpiler)`

### `clearTranspilers()`

> returns `:list`

### `applyTranspiler(node, transpiler)`

### `transpileNode(node)`

### `walkNode(node, visitor)`

### `_mkInt(pos, value)`

> returns `:object`

### `_mkFloat(pos, value)`

> returns `:object`

### `_mkBool(pos, value)`

> returns `:object`

### `_mkString(pos, value)`

> returns `:object`

### `_isConst?(n)`

> returns `:bool`

### `_isNum?(n)`

> returns `:bool`

### `_isInt?(n)`

### `optimizeConstants(node)`

### `removeDebugCalls(node)`

### `createTranspiler(visitor)`

> **thunk** returns `:function`

## Module: `wasm-vm`

- `std` · `import(...)`
- `map` — constant
- `each` — constant
- `reduce` — constant
- `append` — constant
- `indexOf` — constant
- `len` — constant
- `clone` — constant
- `slice` — constant
- `str` · `import(...)`
- `join` — constant
- `fmt` · `import(...)`
- `format` — constant
- `OP_HALT` · `0`
- `OP_NOP` · `1`
- `OP_CONST_NULL` · `2`
- `OP_CONST_EMPTY` · `3`
- `OP_CONST_TRUE` · `4`
- `OP_CONST_FALSE` · `5`
- `OP_CONST_INT` · `6`
- `OP_CONST_FLOAT` · `7`
- `OP_CONST_STRING` · `8`
- `OP_CONST_ATOM` · `9`
- `OP_POP` · `10`
- `OP_DUP` · `11`
- `OP_LOAD_LOCAL` · `12`
- `OP_STORE_LOCAL` · `13`
- `OP_LOAD_UPVAL` · `14`
- `OP_STORE_UPVAL` · `15`
- `OP_ADD` · `16`
- `OP_SUB` · `17`
- `OP_MUL` · `18`
- `OP_DIV` · `19`
- `OP_MOD` · `20`
- `OP_POW` · `21`
- `OP_NEG` · `22`
- `OP_BAND` · `23`
- `OP_BOR` · `24`
- `OP_BXOR` · `25`
- `OP_BRSHIFT` · `26`
- `OP_EQ` · `27`
- `OP_NEQ` · `28`
- `OP_GT` · `29`
- `OP_LT` · `30`
- `OP_GEQ` · `31`
- `OP_LEQ` · `32`
- `OP_NOT` · `33`
- `OP_CONCAT` · `34`
- `OP_MAKE_LIST` · `35`
- `OP_MAKE_OBJECT` · `36`
- `OP_GET_PROP` · `37`
- `OP_SET_PROP` · `38`
- `OP_JUMP` · `39`
- `OP_JUMP_FALSE` · `40`
- `OP_CLOSURE` · `41`
- `OP_CALL` · `42`
- `OP_RETURN` · `43`
- `OP_TAIL_CALL` · `44`
- `OP_BUILTIN` · `45`
- `OP_IMPORT` · `46`
- `OP_IMPORT_DYN` · `47`
- `OP_DEEP_EQ` · `48`
- `OP_SWAP` · `49`
- `OP_MATCH_JUMP` · `50`
- `OP_SCOPE_PUSH` · `51`
- `OP_SCOPE_POP` · `52`
- `OP_CALL_SPREAD` · `53`
- `NUM_OPCODES` · `54`
- `OpcodeName` · `[54 items]`
- `BUILTIN_PRINT` · `0`
- `BUILTIN_LEN` · `1`
- `BUILTIN_TYPE` · `2`
- `BUILTIN_STRING` · `3`
- `BUILTIN_INT` · `4`
- `BUILTIN_FLOAT` · `5`
- `BUILTIN_CODEPOINT` · `6`
- `BUILTIN_CHAR` · `7`
- `BUILTIN_KEYS` · `8`
- `BUILTIN_VALUES` · `9`
- `BUILTIN_SLICE` · `10`
- `BUILTIN_APPEND` · `11`
- `BUILTIN_WAIT` · `12`
- `BUILTIN_EXIT` · `13`
- `BuiltinName` · `[14 items]`
- `TAG_NULL` · `0`
- `TAG_EMPTY` · `1`
- `TAG_INT` · `2`
- `TAG_FLOAT` · `3`
- `TAG_BOOL` · `4`
- `TAG_STRING` · `5`
- `TAG_ATOM` · `6`
- `TAG_LIST` · `7`
- `TAG_OBJECT` · `8`
- `TAG_FUNCTION` · `9`
### `createCompiler()`

> returns `:object`

### `emitByte(c, b)`

### `emitU16(c, v)`

### `emitI32(c, v)`

### `patchI32(c, offset, v)`

> returns `:bool`

### `currentOffset(c)`

### `addConstant(c, entry)`

### `addString(c, s)`

### `addAtom(c, name)`

### `addFloat(c, v)`

### `resolveLocal(c, name)`

### `declareLocal(c, name)`

### `resolveUpvalue(c, name)`

### `resolveBuiltin(name)`

### `compileNode(c, node)`

### `compileAssignment(c, node)`

### `compileIfExpr(c, node)`

### `compileFunction(c, node)`

### `compileGenericCall(c, node)`

### `compileFnCall(c, node)`

### `compileProgram(node)`

### `serializeBytecodeToString(c)`

### `serializeConstantPool(c)`

### `serializeFunctionTable(c)`

### `serializeTopLevelNames(c)`

### `disassemble(c)`

