# Bytecode Binary Output Target (`--bin`)

The `--bin` build target compiles Oak programs to a compact bytecode binary format that can be loaded directly by the Go bytecode VM for fast startup.

## Usage

```sh
# Build to .mgb
oak build --entry main.oak --output app.mgb --bin

# Run directly (auto-detected by .mgb extension)
oak app.mgb

# Or explicitly with --bytecode flag
oak --bytecode app.mgb
```

## Binary Format

The format consists of a header followed by four sections:

```
Header:
  Magic:   4 bytes  "MGbc"
  Version: 2 bytes  u16 little-endian (currently 2)

Bytecode Section:
  Length:  4 bytes  u32 little-endian
  Data:    N bytes  raw bytecode instructions

Constant Pool Section:
  Length:  4 bytes  u32 little-endian
  Data:    N bytes  serialized constants
    Count: 4 bytes u32 LE
    Per entry:
      Type: 1 byte (0=string, 1=atom, 2=float)
      Len:  4 bytes u32 LE
      Data: N bytes (float stored as decimal string)

Function Table Section:
  Length:  4 bytes  u32 little-endian
  Data:    N bytes  serialized function templates
    Count: 4 bytes u32 LE
    Per entry:
      Offset:     4 bytes u32 LE (bytecode offset)
      Arity:      2 bytes u16 LE
      LocalCount: 2 bytes u16 LE
      HasRestArg: 1 byte  (0 or 1)
      NameCount:  2 bytes u16 LE
      Per local name:
        Len:  2 bytes u16 LE
        Data: N bytes

Top-Level Names Section:
  Length:  4 bytes  u32 little-endian
  Data:    N bytes  serialized top-level local names
    Count: 4 bytes u32 LE
    Per name:
      Len:  2 bytes u16 LE
      Data: N bytes
```

## Benefits

- **Fast startup**: No parsing or compilation needed at runtime
- **Smaller files**: Binary is more compact than source text
- **Source protection**: Bytecode is not human-readable
- **Distribution**: Single file contains everything needed to run

## Implementation

The bytecode compiler from `wasm-vm.oak` is reused. This is the same compiler that generates bytecodes for the WASM VM target, so opcode compatibility with `bytecode.go` is maintained.

The runtime preamble from `runtime-native.oak` is compiled into the bytecode so that the `.mgb` file is self-contained — module registration and import resolution functions are part of the bytecode itself.
