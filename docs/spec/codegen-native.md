# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `fmt`

### `format(raw, values...)`

### `printf(raw, values...)`

## Module: `lib\codegen-native.oak`

- `ARCH_X86_64` · `:x86_64`
- `ARCH_ARM64` · `:arm64`
- `IR_CONST` · `:ir_const`
- `IR_LOAD` · `:ir_load`
- `IR_STORE` · `:ir_store`
- `IR_ADD` · `:ir_add`
- `IR_SUB` · `:ir_sub`
- `IR_MUL` · `:ir_mul`
- `IR_DIV` · `:ir_div`
- `IR_MOD` · `:ir_mod`
- `IR_NEG` · `:ir_neg`
- `IR_NOT` · `:ir_not`
- `IR_AND` · `:ir_and`
- `IR_OR` · `:ir_or`
- `IR_XOR` · `:ir_xor`
- `IR_SHL` · `:ir_shl`
- `IR_SHR` · `:ir_shr`
- `IR_EQ` · `:ir_eq`
- `IR_NEQ` · `:ir_neq`
- `IR_LT` · `:ir_lt`
- `IR_GT` · `:ir_gt`
- `IR_LEQ` · `:ir_leq`
- `IR_GEQ` · `:ir_geq`
- `IR_CALL` · `:ir_call`
- `IR_RET` · `:ir_ret`
- `IR_JMP` · `:ir_jmp`
- `IR_JZ` · `:ir_jz`
- `IR_JNZ` · `:ir_jnz`
- `IR_LABEL` · `:ir_label`
- `IR_ALLOC` · `:ir_alloc`
- `IR_PUSH` · `:ir_push`
- `IR_INDEX` · `:ir_index`
- `IR_SETINDEX` · `:ir_setindex`
- `IR_CLOSURE` · `:ir_closure`
- `IR_PHI` · `:ir_phi`
- `IR_NOP` · `:ir_nop`
### `irInst(op, dst, a, b, extra)`

> returns `:object`

### `irConst(dst, value)`

### `irLoad(dst, src)`

### `irStore(dst, src)`

### `irAdd(dst, a, b)`

### `irSub(dst, a, b)`

### `irMul(dst, a, b)`

### `irDiv(dst, a, b)`

### `irMod(dst, a, b)`

### `irNeg(dst, src)`

### `irNot(dst, src)`

### `irAnd(dst, a, b)`

### `irOr(dst, a, b)`

### `irXor(dst, a, b)`

### `irShl(dst, a, b)`

### `irShr(dst, a, b)`

### `irEq(dst, a, b)`

### `irNeq(dst, a, b)`

### `irLt(dst, a, b)`

### `irGt(dst, a, b)`

### `irLeq(dst, a, b)`

### `irGeq(dst, a, b)`

### `irCall(dst, fnReg, args)`

### `irRet(src)`

### `irJmp(label)`

### `irJz(cond, label)`

### `irJnz(cond, label)`

### `irLabel(name)`

### `irAlloc(dst, size)`

### `irPush(list, val)`

### `irIndex(dst, obj, key)`

### `irSetIndex(obj, key, val)`

### `irClosure(dst, fnRef, captures)`

### `irNop()`

### `createIRCompiler()`

> returns `:object`

### `createX86_64Generator()`

> returns `:object`

### `createARM64Generator()`

> returns `:object`

### `compileToNative(ast, arch)`

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

