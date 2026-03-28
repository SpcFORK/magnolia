# Virtual Bytecode VM (Virtual-Bytecode)

## Overview

`Virtual-Bytecode` is a stack-based bytecode virtual machine written in Oak. It can execute bytecode chunks produced by both the Go bytecode compiler (`bytecode.go`, mode `:go`) and the self-hosted WASM compiler (`wasm-vm.oak`, mode `:wasm`). This enables Oak programs to compile and run Oak source code entirely at runtime, without shelling out to the host compiler.

## Import

```oak
vbc := import('Virtual-Bytecode')
// or destructure specific functions
{ runChunk: runChunk, compileSource: compileSource, run: run } := import('Virtual-Bytecode')
```

## Quick Start

```oak
vbc := import('Virtual-Bytecode')

// Compile and run Oak source code
result := vbc.run('1 + 2 * 3')
println(result) // => 7

// With a custom import resolver
result := vbc.run('std := import("std"); std.range(5)', {
    importFn: fn(name) import(name)
})
```

## Constants

### `MODE_WASM`

Atom `:wasm`. Identifies chunks produced by the WASM/self-hosted compiler.

### `MODE_GO`

Atom `:go`. Identifies chunks produced by the Go bytecode compiler.

## Functions

### `run(source, opts?)`

Compiles Oak source code to bytecode and executes it. Alias for `runSource`.

**Parameters:**
- `source` — Oak source code string
- `opts` — Optional object:
  - `globals` — Object of global variable bindings available to the program
  - `importFn` — `fn(name)` called when the program executes `import(name)`. Defaults to a resolver that provides `std`, `str`, `math`, `fmt`, and `syntax`.

**Returns:** The result of the last expression, or an error object `{ type: :error, message: ... }` on parse failure.

```oak
vbc := import('Virtual-Bytecode')

vbc.run('40 + 2') // => 42

vbc.run('x + y', {
    globals: { x: 10, y: 20 }
}) // => 30
```

### `compileSource(source)`

Parses and compiles Oak source code into a bytecode chunk without executing it.

**Parameters:**
- `source` — Oak source code string

**Returns:** A bytecode chunk object, or `{ type: :error, ... }` on parse failure.

```oak
chunk := vbc.compileSource('fn double(n) n * 2; double(21)')
result := vbc.runChunk(chunk)
println(result) // => 42
```

### `compileAst(ast)`

Compiles an already-parsed AST (or list of AST nodes) into a bytecode chunk.

**Parameters:**
- `ast` — A single AST node or a list of AST nodes

**Returns:** A bytecode chunk object.

```oak
syntax := import('syntax')
ast := syntax.parse('1 + 1')
chunk := vbc.compileAst(ast)
```

### `runChunk(rawChunk, opts?)`

Executes a pre-compiled bytecode chunk.

**Parameters:**
- `rawChunk` — A bytecode chunk object (from `compileSource`, `compileAst`, or an externally produced chunk)
- `opts` — Optional object:
  - `globals` — Global variable bindings
  - `importFn` — Import resolver function

**Returns:** The result of execution, or an error object `{ type: :error, message: ..., pc: ... }` on unknown opcode.

```oak
chunk := vbc.compileSource('2 ** 10')
vbc.runChunk(chunk) // => 1024

// Pass globals
chunk := vbc.compileSource('greeting << " world"')
vbc.runChunk(chunk, {
    globals: { greeting: 'hello' }
}) // => 'hello world'
```

### `normalizeChunk(raw)`

Normalizes a raw chunk from either the Go or WASM compiler into a unified internal representation.

**Parameters:**
- `raw` — A raw bytecode chunk object

**Returns:** Normalized chunk with fields:
- `mode` — `:go` or `:wasm`
- `code` — List of byte values
- `constants` — List of `{ type, value }` entries
- `functions` — List of function template objects
- `topLevelNames` — List of top-level variable name strings
- `metadata` — Metadata object

### `decodeConstantPoolBytes(raw)`

Decodes a binary-encoded constant pool (as produced by `bytecode.go`) into a list of constant entries.

**Parameters:**
- `raw` — String or byte list of the encoded constant pool

**Returns:** List of `{ type, value }` objects where type is `:string`, `:atom`, or `:float`.

### `decodeFunctionTableBytes(raw)`

Decodes a binary-encoded function table into a list of function template objects.

**Parameters:**
- `raw` — String or byte list of the encoded function table

**Returns:** List of function template objects with fields: `offset`, `arity`, `localCount`, `hasRestArg`, `name`, `localNames`.

## Bytecode Chunk Format

A normalized bytecode chunk contains:

| Field            | Type   | Description                                      |
|------------------|--------|--------------------------------------------------|
| `mode`           | atom   | `:go` or `:wasm` — identifies the compiler       |
| `code`           | list   | Flat list of byte values (the instruction stream) |
| `constants`      | list   | Constant pool entries                             |
| `functions`      | list   | Function template table                           |
| `topLevelNames`  | list   | Names of top-level local variables                |
| `metadata`       | object | Optional metadata                                 |

## Opcode Set

The VM implements 54 opcodes shared between Go and WASM modes. Opcode numbering differs slightly between modes (the VM auto-detects and uses the correct mapping).

| Opcode          | Description                                   |
|-----------------|-----------------------------------------------|
| `HALT`          | Stop execution                                |
| `NOP`           | No operation                                  |
| `CONST_NULL`    | Push `?`                                      |
| `CONST_EMPTY`   | Push `_`                                      |
| `CONST_TRUE`    | Push `true`                                   |
| `CONST_FALSE`   | Push `false`                                  |
| `CONST_INT`     | Push i32 literal (inline)                     |
| `CONST_FLOAT`   | Push float from constant pool                 |
| `CONST_STRING`  | Push string from constant pool                |
| `CONST_ATOM`    | Push atom from constant pool                  |
| `POP`           | Discard top of stack                          |
| `DUP`           | Duplicate top of stack                        |
| `SWAP`          | Swap top two stack values                     |
| `LOAD_LOCAL`    | Load local variable by slot index             |
| `STORE_LOCAL`   | Store into local variable slot                |
| `LOAD_UPVAL`    | Load captured variable (by name or depth/slot)|
| `STORE_UPVAL`   | Store to captured variable                    |
| `ADD`–`POW`     | Arithmetic: add, sub, mul, div, mod, pow      |
| `NEG`           | Numeric negation                              |
| `BAND`, `BOR`, `BXOR`, `BRSHIFT` | Bitwise operations         |
| `EQ`, `NEQ`     | Equality / inequality                         |
| `DEEP_EQ`       | Structural deep equality                      |
| `GT`, `LT`, `GEQ`, `LEQ` | Comparison operators               |
| `NOT`           | Logical negation                              |
| `CONCAT`        | String/list concatenation (`<<`)              |
| `MAKE_LIST`     | Construct list from N stack values            |
| `MAKE_OBJECT`   | Construct object from N key-value pairs       |
| `GET_PROP`      | Property access                               |
| `SET_PROP`      | Property assignment                           |
| `JUMP`          | Unconditional jump                            |
| `JUMP_FALSE`    | Conditional jump (pop, jump if falsy)         |
| `MATCH_JUMP`    | Pattern match jump (for `if` branches)        |
| `CLOSURE`       | Create closure from function template         |
| `CALL`          | Call function with N arguments                |
| `CALL_SPREAD`   | Call with spread (`fn(args...)`)              |
| `TAIL_CALL`     | Tail-optimized call (reuses stack frame)      |
| `RETURN`        | Return from function                          |
| `BUILTIN`       | Call built-in function by index               |
| `IMPORT`        | Static import (constant pool name)            |
| `IMPORT_DYN`    | Dynamic import (name from stack)              |
| `SCOPE_PUSH`    | Push scope (no-op in this VM)                 |
| `SCOPE_POP`     | Pop scope (no-op in this VM)                  |

## Built-in Functions

The VM provides built-in functions accessible via the `BUILTIN` opcode:

| Index | Built-in      | Description                     |
|-------|---------------|---------------------------------|
| 0     | `print`       | Print a string to stdout        |
| 1     | `len`         | Length of string, list, or object |
| 2     | `type`        | Type atom of a value            |
| 3     | `string`      | Convert to string               |
| 4     | `int`         | Convert to integer              |
| 5     | `float`       | Convert to float                |
| 6     | `codepoint`   | First codepoint of a string     |
| 7     | `char`        | Character from codepoint        |
| 8     | `keys`        | Object keys                     |
| 9     | `values`      | Object values                   |
| 10    | `slice`       | Substring / sublist             |
| 11    | `append`      | Append to list or string        |
| 12    | `wait`        | Wait on a value                 |
| 13    | `exit`        | Exit with code                  |

## Upvalue Resolution

Upvalue (captured variable) resolution differs by mode:

- **Go mode** (`UPVAL_BY_NAME: true`): The operand is a u16 constant pool index containing the variable name. The VM walks the scope chain comparing names.
- **WASM mode** (`UPVAL_BY_NAME: true`): Same name-based strategy. Both modes currently use name-based upvalue resolution for compatibility.

## Dual-Mode Compatibility

The VM auto-detects whether a chunk was produced by the Go compiler or the WASM compiler:

1. If `raw.mode` is set, use it directly.
2. If `raw.topLevelNames` is present, assume Go mode.
3. If the first constant entry has a `.kind` field, assume Go mode; otherwise WASM mode.

Constant entries are normalized from Go format (`{ kind, str, f }`) to the unified format (`{ type, value }`).

## Example: Compile and Inspect

```oak
vbc := import('Virtual-Bytecode')

chunk := vbc.compileSource('fn add(a, b) a + b; add(3, 4)')
normalized := vbc.normalizeChunk(chunk)

println('Mode: ' << string(normalized.mode))
println('Constants: ' << string(len(normalized.constants)))
println('Functions: ' << string(len(normalized.functions)))

result := vbc.runChunk(chunk)
println('Result: ' << string(result)) // => 7
```

## See Also

- [Virtual.md](Virtual.md) — Tree-walking Oak interpreter
- [VirtualToken.md](VirtualToken.md) — AST node constructors
- [wasm-vm.md](wasm-vm.md) — WASM bytecode compiler
