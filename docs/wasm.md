# WebAssembly (WASM) Support

Oak now has experimental support for compiling to WebAssembly. This allows Oak programs to run in any environment that supports WebAssembly, including browsers, Node.js, and other WASM runtimes.

## Overview

The WASM implementation in Oak works in two stages:

1. **Build Stage**: `oak build --wasm` compiles Oak source to WebAssembly Text (WAT) format
2. **Runtime Stage**: The WASM module can then be executed by a WASM host that provides the `oak.run` function

## Building to WASM

To compile an Oak program to WebAssembly:

```bash
oak build --entry main.oak --output program.wat --wasm
```

This generates a `.wat` (WebAssembly Text) file that embeds your Oak source code in the WASM linear memory and exports several functions:

- `main()` - Entry point, calls `__oak_run_bundle()`
- `run()` - The bundled module runner
- `bundle_ptr()` - Returns the memory offset of the bundled source code
- `bundle_len()` - Returns the length of the bundled source code

## Running WASM Modules

### Prerequisites

Since the Oak build system outputs `.wat` (text format), you need to convert it to binary `.wasm` format:

**Option 1: Using WABT (WebAssembly Binary Toolkit)**

```bash
# Install WABT from https://github.com/WebAssembly/wabt
wat2wasm program.wat -o program.wasm
```

**Option 2: Using wasm-interp (included in WABT)**

```bash
# For testing and debugging
wasm-interp program.wat
```

### Executing with Oak

Once you have a binary `.wasm` file, you can execute it using:

```bash
oak wasm program.wasm
```

Or directly by passing the file as an argument:

```bash
oak program.wasm
```

## How It Works

When Oak executes a WASM module:

1. The WASM host runtime is initialized
2. An `oak` module is loaded that exports the `run` function
3. The `run` function is the interface between WASM and the Oak interpreter:
   - Takes a pointer (`i32`) and length (`i32`) of source code in WASM linear memory
   - Reads the source code from memory
   - Executes it using the full Oak interpreter
   - Returns an exit code
4. The WASM module's entrypoint (typically `main()`) is called
5. Results and output are captured and displayed

## Current Limitations

- The runtime currently only supports binary WASM format (`.wasm` files)
- WAT (text) files must be converted to binary using external tools
- Graphics/GUI rendering requires appropriate middleware setup
- Future versions may include:
  - Direct WAT parsing support
  - Browser-based WASM hosts
  - Node.js WASM integration
  - Optimized WASM output

## Example

```oak
// hello.oak
println('Hello from WASM!')
println('1 + 2 = ' << (1 + 2))
```

Build it:

```bash
oak build --entry hello.oak --output hello.wat --wasm
wat2wasm hello.wat -o hello.wasm
```

Run it:

```bash
oak hello.wasm
```

Output:

```
Hello from WASM!
1 + 2 = 3
```

## Future Directions

- **Binary WASM Output**: Integrate WAT-to-WASM conversion into the build system
- **Browser Support**: Implement JavaScript interop for browser-based WASM
- **WASI Support**: Expand WASI bindings for file I/O and system access
- **Optimization**: Bytecode compilation and optimization for WASM targets
