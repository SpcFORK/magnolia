# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\linker.oak`

- `FORMAT_ELF` · `:elf`
- `FORMAT_PE` · `:pe`
- `FORMAT_MACHO` · `:macho`
- `ELF_MAGIC` · `[4 items]`
- `ELFCLASS64` · `2`
- `ELFDATA2LSB` · `1`
- `EV_CURRENT` · `1`
- `ET_EXEC` · `2`
- `EM_X86_64` · `62`
- `EM_AARCH64` · `183`
- `PE_MAGIC` · `[2 items]`
- `PE_SIGNATURE` · `[4 items]`
- `IMAGE_FILE_MACHINE_AMD64` · `34404`
- `IMAGE_FILE_MACHINE_ARM64` · `43620`
### `createByteWriter()`

> returns `:object`

### `createSection(name, flags)`

> returns `:object`

### `linkELF(textCode, dataBytes, entrySymbol, arch)`

> returns `:object`

### `linkPE(textCode, dataBytes, entrySymbol, arch)`

> returns `:object`

### `createLinkerConfig(opts)`

> returns `:object`

### `link(config)`

### `compileAndLink(ast, format, arch)`

