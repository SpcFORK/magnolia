# WASM VM — Bytecode Virtual Machine for Oak

The WASM VM compiles Oak programs to bytecodes and generates a
self-contained WebAssembly module with an embedded bytecode interpreter.
Unlike the previous `--wasm` target (which embeds Oak source text and
delegates to a host-provided `oak.run` import), the VM approach executes
Oak programs **directly in WASM** — no external Oak interpreter required.

## Architecture

```
Oak AST  ──(wasm-vm.oak)──►  Bytecodes + Constants + Function Table
                                        │
                                        ▼
                              ┌───────────────────┐
                              │  WASM Module       │
                              │  ┌───────────────┐ │
                              │  │ Data segments  │ │  ← bytecodes, constants, fn table
                              │  ├───────────────┤ │
                              │  │ VM Interpreter │ │  ← stack-based dispatch loop
                              │  │ (WAT)          │ │
                              │  ├───────────────┤ │
                              │  │ Built-ins      │ │  ← print, len, type, etc.
                              │  ├───────────────┤ │
                              │  │ Heap allocator │ │  ← bump allocator for values
                              │  └───────────────┘ │
                              │                     │
                              │  Host imports:      │
                              │    env.print_string │  ← only I/O
                              │    env.exit         │
                              └───────────────────┘
```

## Usage

The VM target is available via `renderWasmVMBundle` in `runtime-codegen`:

```oak
runtimeCodegen := import('runtime-codegen')
watSource := runtimeCodegen.renderWasmVMBundle(bundleNode)
```

The generated WAT requires only two host imports:

| Import | Signature | Description |
|--------|-----------|-------------|
| `env.print_string` | `(param i32 i32)` | Print bytes at (ptr, len) to stdout |
| `env.exit` | `(param i32)` | Terminate with exit code |

## Bytecode Format

### Opcodes (0–51, contiguous for `br_table` dispatch)

| Code | Name | Operands | Description |
|------|------|----------|-------------|
| 0 | `HALT` | — | Stop execution |
| 1 | `NOP` | — | No operation |
| 2 | `CONST_NULL` | — | Push null |
| 3 | `CONST_EMPTY` | — | Push empty (`_`) |
| 4 | `CONST_TRUE` | — | Push `true` |
| 5 | `CONST_FALSE` | — | Push `false` |
| 6 | `CONST_INT` | i32 | Push integer literal |
| 7 | `CONST_FLOAT` | u16 | Push float from constant pool |
| 8 | `CONST_STRING` | u16 | Push string from constant pool |
| 9 | `CONST_ATOM` | u16 | Push atom from constant pool |
| 10 | `POP` | — | Discard top of stack |
| 11 | `DUP` | — | Duplicate top of stack |
| 12 | `LOAD_LOCAL` | u16 | Push local variable |
| 13 | `STORE_LOCAL` | u16 | Pop into local variable |
| 14 | `LOAD_UPVAL` | u16, u16 | Load from enclosing scope (depth, index) |
| 15 | `STORE_UPVAL` | u16, u16 | Store to enclosing scope |
| 16–21 | Arithmetic | — | `ADD SUB MUL DIV MOD POW` |
| 22 | `NEG` | — | Unary negate |
| 23–26 | Bitwise | — | `BAND BOR BXOR BRSHIFT` |
| 27–32 | Comparison | — | `EQ NEQ GT LT GEQ LEQ` |
| 33 | `NOT` | — | Logical not |
| 34 | `CONCAT` | — | String push arrow (`<<`) |
| 35 | `MAKE_LIST` | u16 | Create list from N stack items |
| 36 | `MAKE_OBJECT` | u16 | Create object from N key-value pairs |
| 37 | `GET_PROP` | — | Property access (obj, key → value) |
| 38 | `SET_PROP` | — | Property set (obj, key, val → obj) |
| 39 | `JUMP` | i32 | Unconditional jump |
| 40 | `JUMP_FALSE` | i32 | Jump if top is falsy |
| 41 | `CLOSURE` | u16 | Create closure from function template |
| 42 | `CALL` | u8 | Call function with N arguments |
| 43 | `RETURN` | — | Return from function |
| 44 | `TAIL_CALL` | u8 | Tail call optimization |
| 45 | `BUILTIN` | u16, u8 | Call built-in function |
| 46 | `IMPORT` | — | Module import (bundled modules resolve at compile time) |
| 47 | `DEEP_EQ` | — | Deep structural equality |
| 48 | `SWAP` | — | Swap top two stack entries |
| 49 | `MATCH_JUMP` | i32 | Pattern match: pop target, compare with TOS, jump if no match |
| 50 | `SCOPE_PUSH` | — | Push scope level (reserved) |
| 51 | `SCOPE_POP` | — | Pop scope level (reserved) |

### Value Representation

Each value on the stack is 8 bytes: `[type:i32][payload:i32]`

| Tag | Type | Payload |
|-----|------|---------|
| 0 | null | 0 |
| 1 | empty | 0 |
| 2 | int | i32 value |
| 3 | float | heap pointer to f64 string |
| 4 | bool | 0 or 1 |
| 5 | string | heap pointer `[len:i32][bytes]` |
| 6 | atom | heap pointer (interned string) |
| 7 | list | heap pointer `[len:i32][cap:i32][items_ptr:i32]` |
| 8 | object | heap pointer `[count:i32][cap:i32][entries_ptr:i32]` |
| 9 | function | heap pointer `[fn_idx:i32][scope_csp:i32]` |

## Memory Layout

```
┌────────────────────┐  0x0000
│  Value Stack       │  16 KB (2048 slots × 8 bytes)
├────────────────────┤  0x4000
│  Call Stack        │  16 KB (frames for function calls)
├────────────────────┤  0x8000
│  Bytecode          │  Variable length
├────────────────────┤
│  Constant Pool     │  Variable length
├────────────────────┤
│  Function Table    │  Variable length (9 bytes per template)
├────────────────────┤
│  Heap              │  Grows upward (bump allocator)
└────────────────────┘
```

## Built-in Functions

| Index | Name | Description |
|-------|------|-------------|
| 0 | `print` | Print value to stdout |
| 1 | `len` | Length of string/list/object |
| 2 | `type` | Type tag as atom |
| 3 | `string` | Convert to string |
| 4 | `int` | Convert to integer |
| 5 | `float` | Convert to float |
| 6 | `codepoint` | Character to codepoint |
| 7 | `char` | Codepoint to character |
| 8 | `keys` | Object keys |
| 9 | `values` | Object values |
| 10 | `slice` | Slice string/list |
| 11 | `append` | Append to list |
| 12 | `wait` | Sleep/pause |
| 13 | `exit` | Terminate program |

## Debugging

The `disassemble` function in `wasm-vm` produces human-readable bytecode listings:

```
0: CONST_INT 42
5: STORE_LOCAL @0
8: LOAD_LOCAL @0
11: BUILTIN print(1)
16: HALT
```

## Files

- [lib/wasm-vm.oak](../lib/wasm-vm.oak) — Bytecode compiler (AST → opcodes)
- [lib/wasm-vm-runtime.oak](../lib/wasm-vm-runtime.oak) — WAT VM runtime generator
- [lib/runtime-codegen.oak](../lib/runtime-codegen.oak) — Integration (`renderWasmVMBundle`)
