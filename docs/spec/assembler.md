# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\assembler.oak`

- `ARCH_X86_64` · `:x86_64`
- `ARCH_ARM64` · `:arm64`
- `x86_64Regs` · `{32 entries}`
- `arm64Regs` · `{}`
### `_initARM64Regs()`

> returns `:int`

- `TOK_LABEL` · `:label`
- `TOK_MNEMONIC` · `:mnemonic`
- `TOK_REG` · `:register`
- `TOK_IMM` · `:immediate`
- `TOK_MEM` · `:memory`
- `TOK_SYMBOL` · `:symbol`
- `TOK_COMMA` · `:comma`
- `TOK_NEWLINE` · `:newline`
- `TOK_COMMENT` · `:comment`
- `TOK_DIRECTIVE` · `:directive`
### `tokenizeAsm(source)`

### `assembleX86_64(source)`

> returns `:object`

### `assembleARM64(source)`

> returns `:object`

### `assemble(source, arch)`

