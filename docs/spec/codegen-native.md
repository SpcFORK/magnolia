# API Documentation

_Auto-generated from Magnolia source._

---

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

