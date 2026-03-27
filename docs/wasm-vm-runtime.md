# WASM VM Runtime (wasm-vm-runtime)

## Overview

`wasm-vm-runtime` generates a complete, self-contained WebAssembly Text (WAT)
module that contains a bytecode interpreter for Oak programs. The generated
module embeds compiled bytecodes, a constant pool, and a function template table
in linear memory, and implements a stack-based VM that executes them without
requiring an external Oak interpreter.

This is the companion module to `wasm-vm`, which handles bytecode compilation.
Together they form the `--wasm` VM pipeline.

## Import

```oak
rt := import('wasm-vm-runtime')
```

Typically used indirectly via `runtime-codegen.renderWasmVMBundle`.

## Host Imports

The generated WASM module requires only two host functions:

| Import | Signature | Description |
|--------|-----------|-------------|
| `env.print_string` | `(param i32 i32)` | Write bytes at (ptr, len) to stdout |
| `env.exit` | `(param i32)` | Terminate execution with exit code |

## Memory Layout

The generated module organizes linear memory as follows (offsets computed
per-program):

| Region | Size | Description |
|--------|------|-------------|
| Value stack | 16 KB (2048 slots) | 8 bytes per entry: type (i32) + payload (i32) |
| Call/frame stack | 16 KB | Call frames for function invocation |
| Bytecode | variable | Compiled Oak program bytecodes |
| Constant pool | variable | String/number literals |
| Function table | variable | Function templates (offset, arity, locals count) |
| Heap | 64 KB initial, growable | Bump-allocated, grows upward on demand |

## API

### `generateVMWat(bytecodeWAT, bytecodeLen, constpoolWAT, constpoolLen, fntableWAT, fntableLen)`

Produces the complete WAT source string for the VM module. Parameters are
WAT-escaped data strings and their byte lengths for each of the three data
segments.

The generated module includes:
- Memory with computed page count
- Data segments for bytecode, constant pool, and function table
- Global registers: `$pc` (program counter), `$sp` (value stack pointer),
  `$csp` (call stack pointer), `$hp` (heap pointer)
- Bump allocator with automatic memory growth
- Value stack operations (`push_val`, `pop_type`, `pop_val`, `peek_type`, `peek_val`)
- Call stack operations for function frames
- `br_table`-based opcode dispatch loop (opcodes 0–51)
- Built-in function handlers (print, len, type, string conversion, arithmetic)
- Exported inspection globals: `vstack_base`, `cstack_base`, `bytecode_base`,
  `constpool_base`, `fntable_base`, `heap_start`

### `generateDispatchTable()`

Builds the WAT `br_table` dispatch block covering all 52 opcodes (0–51). Each
opcode handler is emitted as a nested block with a `(br $dispatch)` loop-back.
Opcodes include `HALT`, `NOP`, constants, arithmetic, comparison, string
operations, object/list manipulation, control flow, function calls (with tail
call optimization), and scope management.

### `generateSetLocalAtHelper()`

Emits a WAT helper function that sets a local variable at a given index within
a function frame's locals region.
