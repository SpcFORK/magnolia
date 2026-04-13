<h1>
    <img width="24" height="24" alt="🌳" src="https://github.com/user-attachments/assets/22fadd8f-707e-4279-8ea0-63bd3da6fdba" />
    <em><b>⠀Magnolia</b></em>⠀🌸
</h1>

**Magnolia** is an expressive, dynamically typed programming language based on [Oak](https://oaklang.org/). It extends Oak with a cross-platform GUI system, audio/DSP processing, full networking stack, concurrency primitives, multi-target compilation (JS, WASM, Lua, TypeScript), GPU interop, AST macros, a self-hosted bytecode VM, and 140+ standard library modules — while keeping the simplicity and elegance of the original language.

> **Highlights:** 3.7x faster bytecode VM · Cross-platform GUI (Win32/X11/WebGL) with Vulkan · CPU shader engine · P2P mesh networking · Thread pools & async event bus · 9 compilation targets · Self-hosted bytecode VM · Mermaid dependency graphs · 67 sample programs

---

## Table of Contents

- [At a Glance](#at-a-glance)
- [Getting Started](#getting-started)
  - [Install](#install)
  - [Quick Start](#quick-start)
  - [CLI Reference](#cli-reference)
- [Language Overview](#language-overview)
  - [Types](#types)
  - [Functions](#functions)
  - [Operators](#operators)
  - [Control Flow](#control-flow)
  - [Classes](#classes)
  - [Async Sugar (`with`)](#async-sugar-with)
- [Features](#features)
  - [Cross-Platform GUI](#-cross-platform-gui)
  - [Audio Processing](#-audio-processing)
  - [Networking Stack](#-networking-stack)
  - [Concurrency](#-thread-library--async-event-bus)
  - [Build System & Multi-Target Compilation](#-build-system--multi-target-compilation)
  - [Virtual Machines & Runtime Codegen](#-virtual-bytecode-vm)
  - [Transpile Middleware & AST Macros](#-transpile-middleware--ast-macros)
  - [Virtual File System & Packed Binaries](#-virtual-file-system--packed-binaries)
  - [Compression & Serialization](#-compression--serialization)
  - [Image & Video](#-image--video)
  - [Crypto & Data Protection](#-crypto--data-protection)
  - [GPU Computing](#-gpu-computing)
  - [Math Extensions](#-math-extensions)
  - [Platform-Native Bindings](#-platform-native-bindings)
  - [Go Runtime & System Interop](#-go-runtime-and-system-interop)
  - [Enhanced Error Display](#-enhanced-error-display)
  - [Code Generation & Runtime Evaluation](#-code-generation-and-runtime-evaluation)
- [Performance & Benchmarks](#performance--benchmarks)
- [Samples](#samples)
- [Editor Support](#editor-support)
- [Project Structure](#project-structure)
- [Development](#development)
- [Testing](#testing)
- [Documentation Index](#documentation-index)
- [Known Limitations](#known-limitations)

---

## At a Glance

```js
// FizzBuzz with pattern matching
std := import('std')

fn fizzbuzz(n) if [n % 3, n % 5] {
    [0, 0] -> 'FizzBuzz'
    [0, _] -> 'Fizz'
    [_, 0] -> 'Buzz'
    _ -> string(n)
}

std.range(1, 101) |> std.each(fn(n) {
    std.println(fizzbuzz(n))
})
```

```js
// Async file I/O
std := import('std')
fs := import('fs')

with fs.readFile('./file.txt') fn(file) if file {
    ? -> std.println('Could not read file!')
    _ -> print(file)
}
```

```js
// HTTP server with routing
std := import('std')
fmt := import('fmt')
http := import('http')

server := http.Server()
with server.route('/hello/:name') fn(params) {
    fn(req, end) if req.method {
        'GET' -> end({
            status: 200
            body: fmt.format('Hello, {{ 0 }}!'
                std.default(params.name, 'World'))
        })
        _ -> end(http.MethodNotAllowed)
    }
}
server.start(9999)
```

---

## Getting Started

### Install

Magnolia is installed from source. You need [Go](https://go.dev/) 1.26+.

**Unix / macOS:**
```sh
make install
```

**Windows:**
```bat
build.bat
```

**Any platform (direct):**
```sh
go build .
```

**Run without installing:**
```sh
go run . <file-or-command>
```

> **Note:** `go build .` produces `magnolia` (or `magnolia.exe`). `make install` installs as `oak`. Use whichever name matches your setup.

### Quick Start

```sh
magnolia repl                    # Interactive REPL
magnolia samples/hello.oak       # Run a file
magnolia eval "1 + 2 * 3"       # Evaluate an expression
magnolia help                    # Show CLI help
```

### CLI Reference

| Command | Description |
|---------|-------------|
| `magnolia <file>` | Run a Magnolia/Oak source file (tree-walking interpreter) |
| `magnolia --bytecode <file>` | Run with the bytecode VM (up to 7x faster) |
| `magnolia --executable <file>` | Run a packed binary bundle (`.mb`/`.mgb`/`.magb`) |
| `magnolia repl` | Start the interactive REPL |
| `magnolia eval "<expr>"` | Evaluate a single expression |
| `magnolia pipe` | Read and evaluate from stdin |
| `magnolia build --entry <file> --output <file>` | Bundle/compile a program |
| `magnolia fmt --fix <files>` | Format source files |
| `magnolia cat <files>` | Concatenate and print files |
| `magnolia pack <file>` | Create a packed self-extracting binary |
| `magnolia help` | Show help text |
| `magnolia version` | Print version info |

**Execution mode flags:**

| Flag | Short | Description |
|------|-------|-------------|
| `--normal` | `-n` | Force tree-walking interpreter |
| `--bytecode` | `-b` | Force bytecode VM |
| `--executable` | `-x` | Run packed binary |

---

## Language Overview

### Types

Magnolia has 8 primitive and 3 complex types:

```js
?        // null, also "()"
_        // "empty" value, equal to anything
1, 2, 3  // integers (64-bit)
3.14     // floats (64-bit)
true     // booleans
'hello'  // strings (mutable byte arrays)
:error   // atoms (immutable interned strings)
pointer(0) // pointers

[1, :number]    // list (vector-backed)
{ a: 'hello' }  // object (hash map)
fn(a, b) a + b  // function (closure)
```

Notable details:
- No implicit type casting, except ints cast up to floats in arithmetic
- Strings are mutable byte arrays (like Lua) — use atoms for immutable strings
- Lists use vector backing — append and index are O(1), clone is O(n)
- Equality on lists and objects is deep equality; no identity equality exists

### Functions

Define functions with `fn`. Name is optional; `()` can be omitted when there are no arguments:

```js
fn double(n) 2 * n
fn speak {
    println('Hello!')
}
```

### Operators

**Assignment:**

```js
a := 1              // local assignment (creates new binding)
[b, c] := [2, 3]    // destructuring assignment
```

**Nonlocal assignment** — walks up scopes to update an existing binding:

```js
n := 10
{
    n <- 30          // updates n in the outer scope
}
n // 30
```

**Arithmetic:** `+`, `-`, `*`, `/`, `%`, `**` (exponentiation)

**Bitwise:** `&` (AND), `|` (OR), `^` (XOR), `~` (NOT), `<<` (left shift), `>>` (right shift)

**Push** (`<<` in list/string context) — mutates and returns:

```js
list := [1, 2, 3]
list << 4 << 5 // [1, 2, 3, 4, 5]
```

**Pipe** (`|>`) — passes the left value as the first argument to the right:

```js
range(10) |> filter(prime?) |> each(println)
10 |> add(20) |> add(3) // 33
```

### Control Flow

Magnolia's `if` is a pattern-matching expression, not a boolean test:

```js
fn pluralize(word, count) if count {
    1 -> word
    2 -> 'a pair of ' + word
    _ -> word + 's'
}
```

Combined with safe tail recursion, this makes Magnolia Turing-complete.

### Classes

The `cs` keyword provides syntactic sugar for constructor functions:

```js
cs Pair(left, right) {
    {
        left: left
        right: right
    }
}
Pair(1, 2).right // 2

cs Counter(start) {
    {
        value: start
        add: fn(delta) start + delta
    }
}
Counter(4).add(3) // 7

// Variadic parameters
cs Bag(items...,) {
    items
}
len(Bag(1, 2, 3)) // 3

// Assignment-only body — methods target instance fields
cs LocalCounter {
    a := 2
    set := fn {
        a <- 3
    }
}
```

Key features: constructor sugar, parameter capture, closure over state, variadic support, assignment-only bodies. Under the hood, classes are just functions that return objects.

### Async Sugar (`with`)

The `with` expression places the trailing callback as the last argument:

```js
with readFile('./path') fn(file) {
    println(file)
}
// desugars to:
readFile('./path', fn(file) {
    println(file)
})
```

For the full language specification, see [docs/spec.md](docs/spec.md).

---

## Features

Magnolia extends Oak with a broad set of capabilities spanning GUI, audio, networking, concurrency, compilation targets, and systems programming.

### 🖼️ Cross-Platform GUI

A full GUI middleware for Windows (Win32/GDI/OpenGL/Vulkan), Linux (X11), and the web (Canvas/WebGL). Create windows, draw 2D shapes, handle input events, render text with custom fonts, and run CPU shaders — all from Oak.

```js
gui := import('GUI')

window := gui.createWindow('My App', 800, 600)
gui.onFrame(window, fn {
    gui.fillRect(window, 10, 10, 200, 100, gui.rgb(60, 120, 220))
    gui.drawText(window, 20, 40, 'Hello Magnolia!', gui.rgb(255, 255, 255))
})
gui.run(window)
```

Highlights:
- **2D primitives**: rectangles, circles, polygons, bezier curves, ellipses, arcs, stars, rings, capsules, rounded rects, arrows, spirals
- **Camera systems**: world-to-screen mapping for scrolling/zooming
- **Virtual resolution**: render at a logical size and scale to the physical window (`fit`, `fill`, `stretch`, `pixelPerfect`)
- **CPU shader engine**: per-pixel fragment shaders with noise (Perlin, fBm), SDFs, HSL/HSV color, easing functions, multi-pass composition, and GLSL/HLSL codegen
- **Presenter backends**: GDI, DirectDraw, OpenGL, Vulkan (Windows); X11 (Linux); Canvas/WebGL (web)
- **Font & text**: TTF metrics, cached font handles, text extent measurement

```js
sh := import('gui-shader')
gui := import('GUI')

// Rainbow shader
rainbow := sh.Shader(fn(x, y, w, h, t, _) {
    sh.hsl2rgb(sh.fract(t * 0.1 + float(x) / float(w)), 1.0, 0.5)
}, { resolution: 4 })

gui.shaderRender(window, rainbow, 0, 0, 320, 240)
```

See [docs/gui.md](docs/gui.md), [docs/gui-2d.md](docs/gui-2d.md), [docs/gui-shader.md](docs/gui-shader.md), and [docs/gui-resolution.md](docs/gui-resolution.md).

---

### 🔊 Audio Processing

Full-featured audio with PCM sample handling, WAV encoding, oscillators, DSP transforms, ADSR envelopes, and FFT spectral analysis.

```js
audio := import('audio')

// Generate a 440 Hz sine wave, write as WAV
samples := audio.sine(440, 44100, 2.0)
data := audio.wav(samples, 44100, 1, 16)
writeFile('tone.wav', data)
```

- WAV encoding: 8/16/32-bit, mono/stereo, configurable sample rates (CD 44.1 kHz, DVD 48 kHz, HD 96 kHz)
- DSP: window functions (Hann, Hamming), FIR convolution, filter kernel generation
- FFT/IFFT: radix-2 Cooley-Tukey, magnitude/phase extraction
- ADSR envelope shaping for synthesized sounds

See [docs/audio.md](docs/audio.md), [docs/audio-wav.md](docs/audio-wav.md), [docs/audio-dsp.md](docs/audio-dsp.md), and [docs/audio-fft.md](docs/audio-fft.md).

---

### 🌐 Networking Stack

A complete networking toolkit: HTTP servers with routing, WebSockets, raw TCP sockets, email (SMTP/IMAP), peer-to-peer relay mesh, and local-network game discovery.

```js
// HTTP server
http := import('http')
server := http.Server()
server.route('/api/:resource', fn(params) fn(req, end) {
    end({ status: 200, body: 'Resource: ' + params.resource })
})
server.start(8080)
```

```js
// P2P mesh relay
p2p := import('p2p')
host := p2p.Host('0.0.0.0:9411', '/mesh', fn(evt) println(evt))
peer := p2p.join('ws://localhost:9411/mesh', 'alice', fn(evt) {
    if evt.type { :ready -> peer.broadcast({ msg: 'hello everyone' }) }
})
```

- **HTTP**: URL routing with path parameters, static file serving, URL encoding/decoding
- **WebSocket**: client and server helpers, message opcodes, bidirectional communication
- **TCP Sockets**: raw TCP/TLS streams, SNI support, certificate verification
- **SMTP / IMAP**: email sending and mailbox reading with STARTTLS
- **P2P**: relay-based mesh with peer discovery, direct/broadcast messaging, channels
- **WLAN**: local-network game beacon scanning, subnet peer sweep

See [docs/http.md](docs/http.md), [docs/websocket.md](docs/websocket.md), [docs/p2p.md](docs/p2p.md), and [docs/wlan.md](docs/wlan.md).

---

### 🧵 Thread Library & Async Event Bus

High-level concurrency primitives: mutexes, semaphores, wait groups, thread pools, and a publish/subscribe event bus.

```js
thread := import('thread')

// Parallel map
results := thread.pmap([1, 2, 3, 4], fn(x) x * x)

// Mutex for safe shared state
mutex := thread.Mutex()
mutex.lock()
// critical section
mutex.unlock()

// Thread pool
pool := thread.Pool(4)
pool.submit(fn() {
    // work to be done
})
```

```js
// Async event bus
bus := import('async-event-bus')
eb := bus.create()
eb.on('player:join', fn(payload, _) println('joined: ' + payload.name))
eb.emit('player:join', { name: 'Alice' })
```

See [docs/thread.md](docs/thread.md) and [docs/async-event-bus.md](docs/async-event-bus.md).

---

### 🔧 Transpile Middleware & AST Macros

A plugin architecture for AST transformations during the build process, plus first-class AST macros for compile-time metaprogramming.

```js
build := import('build')
transpile := build.transpile

// Create custom transpiler
myTranspiler := transpile.createTranspiler(fn(node) {
    // Transform AST nodes
    node
})

build.run({
    entry: 'main.oak'
    transpilers: [myTranspiler]
})
```

```js
syntax := import('syntax')

// Define macro expanders
myMacro := syntax.Macro(fn(node) {
    // Transform AST nodes at parse time
    node
})
ast := syntax.parseWithMacros('(my-macro 1 2 3)', [myMacro])
```

See [docs/transpile.md](docs/transpile.md) and [docs/syntax.md](docs/syntax.md).

---

### 📁 Virtual File System & Packed Binaries

An in-memory file system that can be embedded in packed binaries for single-file deployment.

```js
Virtual := import('Virtual')

vfs := Virtual.createVirtualFS({
    'config.json': '{"version": "1.0"}'
    'data/test.txt': 'test data'
})

content := vfs.readFile('config.json')
vfs.writeFile('output.txt', 'Hello World')
```

The build system can embed a VFS directory into a compiled bundle, so your program ships with all its assets in one file. See [docs/virtual-fs.md](docs/virtual-fs.md) and [docs/pack.md](docs/pack.md).

---

### 🖥️ Virtual Bytecode VM

A self-hosted stack-based bytecode VM that compiles and executes Oak source code at runtime — enabling dynamic code evaluation, rule engines, and sandboxed execution without shelling out.

```js
vbc := import('Virtual-Bytecode')

result := vbc.run('fn fib(n) if n < 2 { true -> n, _ -> fib(n - 1) + fib(n - 2) }; fib(10)')
println(result) // => 55

// Or compile once, run many times
chunk := vbc.compileSource('x * x + 1')
vbc.runChunk(chunk, { globals: { x: 7 } }) // => 50
```

- 52-opcode instruction set (matches `bytecode.go` and `wasm-vm.oak`)
- Dual-mode: auto-detects Go-compiled or WASM-compiled chunks
- Custom global bindings and import resolvers
- Closures, tail calls, destructuring, pattern matching, rest args

See [docs/Virtual-Bytecode.md](docs/Virtual-Bytecode.md), [docs/Virtual.md](docs/Virtual.md), and [docs/wasm-vm.md](docs/wasm-vm.md).

---

### 🗜️ Compression & Serialization

Lossless compression codecs and compact binary serialization.

```js
compression := import('compression')
compressed := compression.huffmanCompress('hello world hello world')
original := compression.huffmanDecompress(compressed)

msgpack := import('msgpack')
packed := msgpack.serializeSafe({ name: 'alice', scores: [98, 87, 95] })
```

- **RLE**: run-length encoding for simple repeated patterns
- **Huffman**: frequency-optimized variable-length bit packing
- **LZW**: dictionary-based compression for repeated phrases
- **MessagePack**: compact binary serialization (smaller/faster than JSON)

See [docs/compression.md](docs/compression.md) and [docs/msgpack.md](docs/msgpack.md).

---

### 🖼️ Image & Video

BMP image encoding and frame-based video containers for pixel-stream pipelines.

```js
bmp := import('bmp')
pixels := [255, 0, 0, 0, 255, 0, 0, 0, 255] // 3 RGB pixels
data := bmp(3, 1, pixels)
writeFile('rgb.bmp', data)
```

- **BMP**: 24-bit image encoding with automatic row padding
- **ICO**: Windows icon (.ico) file generation
- **Video**: frame containers with per-frame pixel buffers (width, height, channels)

See [docs/bmp.md](docs/bmp.md), [docs/ico.md](docs/ico.md), and [docs/video.md](docs/video.md).

---

### 🔐 Crypto & Data Protection

Cryptographic primitives and data integrity helpers.

```js
crypto := import('crypto')

id := crypto.uuid()          // RFC 4122 v4
hash := crypto.sha256('data')
safe := crypto.randomBytes(32)
```

- **Crypto**: UUID v4, SHA-256, cryptographically secure random bytes/ints, constant-time comparison, session tokens
- **Data protection**: parity checks, XOR/additive checksums, CRC16-CCITT, CRC32, Hamming distance, LDPC syndrome validation

See [docs/crypto.md](docs/crypto.md) and [docs/dataprot.md](docs/dataprot.md).

---

### 🎮 GPU Computing

Low-level helpers for GPU interop with CUDA and OpenCL support:

```js
gpu := import('gpu')

// Scan available GPU backends
backends := gpu.scan()

// Call CUDA functions
cudaInit := gpu.cuda('cuInit')
gpu.call(cudaInit, 0)

// Or OpenCL
clGetPlatforms := gpu.opencl('clGetPlatformIDs')
```

See [docs/gpu.md](docs/gpu.md) and [docs/gpus.md](docs/gpus.md).

---

### 📐 Math Extensions

Extended math libraries for geometry, statistics, and foundational primitives.

```js
geo := import('math-geo')
stats := import('math-stats')

geo.hypot(0, 0, 3, 4)       // => 5 (Euclidean distance)
geo.bearing(0, 0, 1, 1)     // angle in radians

stats.mean([1, 2, 3, 4, 5]) // => 3
stats.stddev([2, 4, 4, 4, 5, 5, 7, 9]) // standard deviation
```

- **math-geo**: Euclidean distance, coordinate scaling, polar-to-Cartesian, bearing/orientation
- **math-stats**: mean, median, standard deviation, sum, product, clamp, min/max aggregation
- **math-base**: dependency-free core (`Pi`, `E`, `sqrt`, `abs`, `sign`) used by sub-modules

See [docs/math.md](docs/math.md), [docs/math-geo.md](docs/math-geo.md), [docs/math-stats.md](docs/math-stats.md), and [docs/math-base.md](docs/math-base.md).

---

### 🔩 Go Runtime and System Interop

Built-ins for low-level host interop, including goroutines/channels, runtime metadata, foreign procedure calls, and raw memory access:

```js
info := ___runtime_sys_info()
println(info.os + '/' + info.arch)

ch := make_chan(1)
go(fn {
    chan_send(ch, ___runtime_go_version())
})

evt := chan_recv(ch)
println(evt.data)
```

See [docs/go.md](docs/go.md) for complete usage and safety notes.

---

### ⚙️ Code Generation and Runtime Evaluation

A runtime code generation library for dynamic code synthesis and evaluation:

```js
codegen := import('runtime-codegen')

// Generate code templates dynamically
template := codegen.template('fn add(a, b) a + b')

// Create custom code generators
generator := codegen.createGenerator(fn(type) 
    'fn process_' + type + '(x) x'
)

// Evaluate generated code at runtime
code := generator('string')
fn := codegen.eval(code)
```

See [docs/runtime-codegen.md](docs/runtime-codegen.md).

---

### 🌍 Build System & Multi-Target Compilation

Magnolia's build system bundles multi-file programs into single deployable artifacts, with **9 compilation targets**:

| Target | Flag | Description |
|--------|------|-------------|
| Oak native | *(default)* | Bundled `.oak` source |
| JavaScript | `--web` | Browser / Node.js / Deno |
| WebAssembly | `--wasm` | `.wat` text format or embedded bytecode VM |
| JSON AST | `--ast` | Serialized AST representation |
| Bytecode binary | `--bin` | Pre-compiled bytecode bundle |
| Documentation | `--doc` | Generated docs |
| TypeScript | `--ts` | TypeScript output |
| Lua | `--lua` | Lua transpilation |
| Dependency graph | `--graph` | Mermaid diagram of module imports |

```sh
magnolia build --entry src/main.oak --output dist/bundle.oak         # Native
magnolia build --entry src/app.oak --output dist/bundle.js --web     # JavaScript
magnolia build --entry src/main.oak --output dist/program.wat --wasm  # WebAssembly
magnolia build --entry src/main.oak --output dist/deps.mmd --graph    # Dependency graph
```

Features: module bundling, dependency resolution, tree-shaking, code minification, virtual filesystem embedding, and transpile middleware hooks.

See [docs/build.md](docs/build.md), [docs/wasm.md](docs/wasm.md), and [docs/mermaid.md](docs/mermaid.md).

---

### 🖥️ Platform-Native Bindings

Direct access to platform APIs for systems programming:

- **Windows**: kernel32, ntdll, user32, gdi32, psapi DLL exports; virtual memory (mmap/munmap/mprotect); registry access; GUI integration; WLAN scanning
- **Linux**: libc, libdl, libX11 symbol resolution; process helpers (getpid, sysconf); X11 windowing; dynamic library loading (dlopen/dlsym/dlclose)

```js
// Windows
win := import('windows')
hModule := win.kernel32('GetModuleHandleW', 0)

// Linux
linux := import('Linux')
pid := linux.getpid()
```

See [docs/windows.md](docs/windows.md) and [docs/linux.md](docs/linux.md).

### 🎨 Enhanced Error Display

Beautiful, color-coded error messages with source code context:

```
╭─ Runtime Error ───────────────────────────────────────────
│
│ File: test.oak
│ Position: [4:8]
│
│ Division by zero
│
│ Context:
│    2 │ x := 10
│    3 │ y := 20
│    4 │ z := x / 0
│      │        ^
│    5 │ 
│    6 │ println(z)
╰───────────────────────────────────────────────────────────
```

See [docs/error-display.md](docs/error-display.md).

---

## Performance & Benchmarks

Magnolia offers three execution modes with different performance profiles:

| Mode | Flag | Description |
|------|------|-------------|
| Tree-walking interpreter | `--normal` / `-n` | Default. AST-walking evaluator |
| Bytecode VM | `--bytecode` / `-b` | 52-opcode stack-based VM, significantly faster |
| Packed binary | `--executable` / `-x` | Pre-compiled bytecode bundles |

### Benchmark Results

A/B engine test across 17 workloads (median of 5 runs, tree-walking mode with `bytecode()` / `interpreter()` engine-switching builtins):

| Benchmark | Tree-walk (ms) | Bytecode (ms) | Speedup |
|-----------|----------------|---------------|---------|
| int arithmetic (200k) | 210.3 | 36.8 | **5.7x** |
| float arithmetic (200k) | 270.5 | 50.1 | **5.4x** |
| fib(20) naive | 13.1 | 4.1 | **3.2x** |
| closure create+call (10k) | 21.4 | 4.5 | **4.8x** |
| compose chain (10k) | 22.8 | 5.6 | **4.1x** |
| fizzbuzz classify (20k) | 40.8 | 11.1 | **3.7x** |
| sieve of Eratosthenes (10k) | 56.5 | 28.8 | **2.0x** |
| map (10k) | 20.2 | 16.6 | 1.2x |
| filter (10k) | 23.9 | 21.1 | 1.1x |
| sort 2k random ints | 84.4 | 73.4 | 1.2x |
| string concat (2k) | 3.1 | 2.0 | 1.5x |
| object build (3k keys) | 5.0 | 1.0 | **5.0x** |
| object read (3k keys) | 4.0 | 1.0 | **4.0x** |

**Bytecode VM wins 15 of 17 benchmarks**, with up to **5.7x speedup** on arithmetic-heavy workloads and **3–5x** on function call / object access patterns.

### Optimizations Implemented

1. **Constant folding** — arithmetic, string concat, boolean logic, and equality on literals resolved at compile time.
2. **Memory allocation audit** — pprof-guided pass yielding **2.4x speedup** on fib(20). Key wins: `sync.Pool` for scope mutexes, pre-sized scope maps, direct fn-call fast path, index-based tokenizer.
3. **Bytecode VM** — 52-opcode stack-based VM with closures, destructuring, pattern matching, pipes, and upvalue mutation. **3.7x speedup** on fib(30), **1.9x** on loop-heavy workloads.

### Optimization Roadmap

- Register-based bytecode VM (cut push/pop overhead)
- NaN-boxing / tagged-pointer value representation
- Inline caching for property access
- Escape analysis and stack allocation
- Lazy / copy-on-write strings and lists
- Parallel GC or arenas for short-lived scopes
- JIT compilation of hot bytecode paths
- SIMD-accelerated string / list builtins

---

## Samples

The [samples/](samples/) directory contains **67 example programs** covering the full breadth of Magnolia's capabilities:

| Category | Examples |
|----------|----------|
| **Core language** | `hello.oak`, `fizzbuzz.oak`, `fib.oak`, `tailcall.oak` |
| **GUI & Graphics** | `gui-sample.oak`, `gui-2d.oak`, `gui-3d.oak`, `gui-fonts.oak`, `gui-game.oak`, `gui-graphing.oak`, `gui-lighting.oak` |
| **GUI Forms** | `gui-form-login.oak`, `gui-form-settings.oak`, `gui-form-wizard.oak`, `gui-form-dashboard.oak` |
| **Audio & Video** | `audio-demo.oak`, `gui-audio.oak` |
| **Networking** | `fileserver.oak`, `p2p.oak`, `p2p-cli.oak`, `listen-multiport.oak` |
| **Concurrency** | `thread-examples.oak` |
| **Virtual FS** | `vfs-example.oak`, `vfs-bundle-example.oak` |
| **Windows native** | `windows-2d-layer-hotload.oak`, `windows-d3d9.oak`, `windows-dll-bindings.oak`, `windows-registry.oak`, +11 more |
| **Linux native** | `linux-draw.oak`, `linux-interop.oak`, `linux-window.oak` |
| **Data & utilities** | `json-examples.oak`, `crypto-uuid.oak`, `datetime-examples.oak`, `compression-benchmark.oak`, `sort-examples.oak` |
| **Performance** | `perf-ab-test.oak`, `perf-bench.oak` |
| **Misc** | `transpile-examples.oak`, `pointers-bits.oak`, `go-interop.oak`, `shell-example.oak`, `md-parser.oak` |

---

## Editor Support

| Editor | Plugin | Location |
|--------|--------|----------|
| **VS Code** | Oak/Magnolia syntax highlighting | [tools/oak-vscode/](tools/oak-vscode/) (also available as [oak-vscode.zip](tools/oak-vscode.zip)) |
| **Vim** | Syntax file | [tools/oak.vim](tools/oak.vim) |

---

## Project Structure

```
magnolia/
├── *.go                  # Interpreter core (tokenizer, parser, evaluator, bytecode VM)
├── cmd/                  # Built-in CLI commands (build, fmt, help, pack, shell, etc.)
├── lib/                  # Standard library (~144 modules)
│   ├── GUI.oak           #   GUI entry point
│   ├── gui-*.oak         #   GUI subsystems (2D, shaders, fonts, events, native backends)
│   ├── audio*.oak        #   Audio processing (PCM, WAV, DSP, FFT)
│   ├── http.oak          #   HTTP server & routing
│   ├── websocket.oak     #   WebSocket support
│   ├── thread.oak        #   Concurrency primitives
│   ├── build*.oak        #   Build system & bundler
│   ├── mermaid.oak       #   Mermaid diagram generation & module graphs
│   ├── syntax*.oak       #   Parsing, macros, AST transforms
│   ├── windows*.oak      #   Windows platform bindings
│   ├── Linux*.oak        #   Linux platform bindings
│   └── ...               #   crypto, compression, math, etc.
├── docs/                 # Comprehensive documentation (~107 files)
├── samples/              # Example programs (~67 files)
├── test/                 # Unit & generative tests (~50 test files)
├── tools/                # Editor plugins (VS Code, Vim)
├── www/                  # Website source
├── build/                # Build output
├── Makefile              # Unix build/test/install targets
└── build.bat             # Windows build script
```

---

## Development

### Build & Test Commands

**Unix (Make):**

| Command | Description |
|---------|-------------|
| `make run` | Build and start the REPL |
| `make tests` / `make t` | Run Go test suite |
| `make test-oak` / `make tk` | Run Magnolia standard library tests |
| `make test-bundle` | Test via the build/bundle system |
| `make test-js` | Run tests compiled to JavaScript (Node.js) |
| `make test-wasm` | Run WebAssembly target tests |
| `make test-pack` | Test pack functionality |
| `make build` | Cross-platform release builds |
| `make build-linux` / `make build-darwin` / `make build-windows` | OS-specific builds |
| `make fmt` / `make f` | Format files with unstaged changes |
| `make install` | Install as `oak` + Vim syntax file |
| `make site` | Build website bundle |
| `make site-gen` | Rebuild static site docs |

**Windows:**

If `-race` commands fail due to CGO/toolchain constraints, use non-race alternatives:

```sh
go test .
go run . test/main.oak
```

---

## Testing

Magnolia has two kinds of tests:

- **Unit tests** — assertion-based tests built on the `libtest` library, plus Go-level tests in `eval_test.go`. Run with `make tests` (Go) and `magnolia test/main.oak` (Magnolia).
- **Generative / fuzz tests** — procedurally generated inputs validating parsers, date/time algorithms, and other complex behaviors. Located in `test/generative/`.

Both sets run entirely in "userland" without invoking the interpreter separately.

---

## Documentation Index

Comprehensive docs for every module live in [docs/](docs/). The [docs/spec/](docs/spec/) directory contains **184 detailed specification files** mirroring the stdlib — one per module — covering every function signature, data structure, and usage pattern.

<details>
<summary><strong>Language Specification</strong> (184 spec files)</summary>

The [docs/spec/](docs/spec/) directory provides per-module API specifications organized by category:

- **Syntax & AST** — [syntax.md](docs/spec/syntax.md) · [syntax-tokenize.md](docs/spec/syntax-tokenize.md) · [syntax-parse.md](docs/spec/syntax-parse.md) · [syntax-print.md](docs/spec/syntax-print.md) · [syntax-macros.md](docs/spec/syntax-macros.md) · [syntaxfmt.md](docs/spec/syntaxfmt.md) · [ast-analyze.md](docs/spec/ast-analyze.md) · [ast-transform.md](docs/spec/ast-transform.md)
- **Build & Bundling** — [build.md](docs/spec/build.md) · [build-analyze.md](docs/spec/build-analyze.md) · [build-ast.md](docs/spec/build-ast.md) · [build-config.md](docs/spec/build-config.md) · [build-ident.md](docs/spec/build-ident.md) · [build-imports.md](docs/spec/build-imports.md) · [build-includes.md](docs/spec/build-includes.md) · [build-render.md](docs/spec/build-render.md) · [build-render-node.md](docs/spec/build-render-node.md) · [bundle-ast.md](docs/spec/bundle-ast.md) · [bundle-utils.md](docs/spec/bundle-utils.md) · [pack.md](docs/spec/pack.md) · [pack-utils.md](docs/spec/pack-utils.md)
- **Audio** — [audio.md](docs/spec/audio.md) · [audio-aiff.md](docs/spec/audio-aiff.md) · [audio-au.md](docs/spec/audio-au.md) · [audio-wav.md](docs/spec/audio-wav.md) · [audio-ogg.md](docs/spec/audio-ogg.md) · [audio-raw.md](docs/spec/audio-raw.md) · [audio-dsp.md](docs/spec/audio-dsp.md) · [audio-fft.md](docs/spec/audio-fft.md) · [audio-complex.md](docs/spec/audio-complex.md)
- **AI & Machine Learning** — [ai.md](docs/spec/ai.md) · [ai-data.md](docs/spec/ai-data.md) · [ai-decode.md](docs/spec/ai-decode.md) · [ai-linalg.md](docs/spec/ai-linalg.md) · [ai-ml.md](docs/spec/ai-ml.md) · [ai-nn.md](docs/spec/ai-nn.md) · [ai-optim.md](docs/spec/ai-optim.md) · [ai-text.md](docs/spec/ai-text.md) · [ai-vec.md](docs/spec/ai-vec.md)
- **GUI Core** — [GUI.md](docs/spec/GUI.md) · [gui-2d.md](docs/spec/gui-2d.md) · [gui-3dmath.md](docs/spec/gui-3dmath.md) · [gui-audio.md](docs/spec/gui-audio.md) · [gui-canvas.md](docs/spec/gui-canvas.md) · [gui-color.md](docs/spec/gui-color.md) · [gui-common.md](docs/spec/gui-common.md) · [gui-draw.md](docs/spec/gui-draw.md) · [gui-graph.md](docs/spec/gui-graph.md) · [gui-lighting.md](docs/spec/gui-lighting.md) · [gui-loop.md](docs/spec/gui-loop.md) · [gui-menus.md](docs/spec/gui-menus.md) · [gui-mesh.md](docs/spec/gui-mesh.md) · [gui-raster.md](docs/spec/gui-raster.md) · [gui-render.md](docs/spec/gui-render.md) · [gui-resolution.md](docs/spec/gui-resolution.md) · [gui-theme.md](docs/spec/gui-theme.md) · [gui-thread.md](docs/spec/gui-thread.md) · [gui-video.md](docs/spec/gui-video.md) · [gui-web.md](docs/spec/gui-web.md)
- **GUI Shaders** — [gui-shader.md](docs/spec/gui-shader.md) · [gui-shader-codegen.md](docs/spec/gui-shader-codegen.md) · [gui-shader-color.md](docs/spec/gui-shader-color.md) · [gui-shader-math.md](docs/spec/gui-shader-math.md) · [gui-shader-noise.md](docs/spec/gui-shader-noise.md) · [gui-shader-sdf.md](docs/spec/gui-shader-sdf.md)
- **GUI Input & Interaction** — [gui-input.md](docs/spec/gui-input.md) · [gui-form.md](docs/spec/gui-form.md) · [gui-events.md](docs/spec/gui-events.md) · [gui-gamepad.md](docs/spec/gui-gamepad.md) · [gui-filedrop.md](docs/spec/gui-filedrop.md) · [gui-clipboard.md](docs/spec/gui-clipboard.md) · [gui-dialogs.md](docs/spec/gui-dialogs.md) · [gui-accessibility.md](docs/spec/gui-accessibility.md) · [gui-systray.md](docs/spec/gui-systray.md) · [gui-fonts.md](docs/spec/gui-fonts.md)
- **GUI Rendering Backends** — [gui-aa.md](docs/spec/gui-aa.md) · [gui-print.md](docs/spec/gui-print.md) · [gui-leak-detect.md](docs/spec/gui-leak-detect.md) · [gui-gpu-info.md](docs/spec/gui-gpu-info.md) · [gui-draw-ops.md](docs/spec/gui-draw-ops.md) · [gui-test.md](docs/spec/gui-test.md)
- **GUI Native Windows** — [gui-native-win.md](docs/spec/gui-native-win.md) · [gui-native-win-close.md](docs/spec/gui-native-win-close.md) · [gui-native-win-frame.md](docs/spec/gui-native-win-frame.md) · [gui-native-win-icons.md](docs/spec/gui-native-win-icons.md) · [gui-native-win-opengl.md](docs/spec/gui-native-win-opengl.md) · [gui-native-win-poll.md](docs/spec/gui-native-win-poll.md) · [gui-native-win-present.md](docs/spec/gui-native-win-present.md) · [gui-native-win-ddraw.md](docs/spec/gui-native-win-ddraw.md) · [gui-native-win-d3d11.md](docs/spec/gui-native-win-d3d11.md) · [gui-native-win-vulkan.md](docs/spec/gui-native-win-vulkan.md) · [gui-native-win-vulkan-init.md](docs/spec/gui-native-win-vulkan-init.md) · [gui-native-win-vulkan-swapchain.md](docs/spec/gui-native-win-vulkan-swapchain.md) · [gui-native-win-vulkan-present.md](docs/spec/gui-native-win-vulkan-present.md) · [gui-native-win-vulkan-constants.md](docs/spec/gui-native-win-vulkan-constants.md) · [gui-native-win-probe.md](docs/spec/gui-native-win-probe.md)
- **GUI Native Linux** — [gui-native-linux.md](docs/spec/gui-native-linux.md)
- **Data Formats** — [data.md](docs/spec/data.md) · [data-csv.md](docs/spec/data-csv.md) · [data-ini.md](docs/spec/data-ini.md) · [data-toml.md](docs/spec/data-toml.md) · [data-xml.md](docs/spec/data-xml.md) · [data-yaml.md](docs/spec/data-yaml.md) · [json.md](docs/spec/json.md) · [md.md](docs/spec/md.md) · [msgpack.md](docs/spec/msgpack.md)
- **Image Formats** — [image.md](docs/spec/image.md) · [image-bmp.md](docs/spec/image-bmp.md) · [image-ico.md](docs/spec/image-ico.md) · [image-ppm.md](docs/spec/image-ppm.md) · [image-qoi.md](docs/spec/image-qoi.md) · [image-tga.md](docs/spec/image-tga.md)
- **Compression** — [compression.md](docs/spec/compression.md) · [compression-huffman.md](docs/spec/compression-huffman.md) · [compression-lzw.md](docs/spec/compression-lzw.md) · [compression-rle.md](docs/spec/compression-rle.md)
- **Networking** — [socket.md](docs/spec/socket.md) · [websocket.md](docs/spec/websocket.md) · [http.md](docs/spec/http.md) · [p2p.md](docs/spec/p2p.md) · [email.md](docs/spec/email.md) · [email-smtp.md](docs/spec/email-smtp.md) · [email-imap.md](docs/spec/email-imap.md) · [email-pop.md](docs/spec/email-pop.md) · [smtp.md](docs/spec/smtp.md) · [imap.md](docs/spec/imap.md) · [pop.md](docs/spec/pop.md) · [WLAN.md](docs/spec/WLAN.md)
- **Crypto & Security** — [crypto.md](docs/spec/crypto.md) · [dataprot.md](docs/spec/dataprot.md)
- **Math** — [math.md](docs/spec/math.md) · [math-base.md](docs/spec/math-base.md) · [math-geo.md](docs/spec/math-geo.md) · [math-stats.md](docs/spec/math-stats.md)
- **Runtime & VMs** — [Virtual.md](docs/spec/Virtual.md) · [Virtual-Bytecode.md](docs/spec/Virtual-Bytecode.md) · [VirtualToken.md](docs/spec/VirtualToken.md) · [wasm-vm.md](docs/spec/wasm-vm.md) · [wasm-vm-runtime.md](docs/spec/wasm-vm-runtime.md) · [runtime-native.md](docs/spec/runtime-native.md) · [runtime-js.md](docs/spec/runtime-js.md) · [runtime-codegen.md](docs/spec/runtime-codegen.md) · [codegen-common.md](docs/spec/codegen-common.md)
- **Platform — Windows** — [windows.md](docs/spec/windows.md) · [windows-constants.md](docs/spec/windows-constants.md) · [windows-core.md](docs/spec/windows-core.md) · [windows-flags.md](docs/spec/windows-flags.md) · [windows-kernel.md](docs/spec/windows-kernel.md) · [windows-gdi.md](docs/spec/windows-gdi.md) · [windows-fonts.md](docs/spec/windows-fonts.md) · [windows-windowing.md](docs/spec/windows-windowing.md) · [windows-registry.md](docs/spec/windows-registry.md) · [windows-net.md](docs/spec/windows-net.md) · [windows-loader.md](docs/spec/windows-loader.md) · [win-common.md](docs/spec/win-common.md)
- **Platform — Linux** — [Linux.md](docs/spec/Linux.md) · [linux-constants.md](docs/spec/linux-constants.md) · [linux-core.md](docs/spec/linux-core.md) · [linux-libc.md](docs/spec/linux-libc.md) · [linux-fonts.md](docs/spec/linux-fonts.md) · [linux-loader.md](docs/spec/linux-loader.md) · [linux-windowing.md](docs/spec/linux-windowing.md)
- **Core & Utilities** — [std.md](docs/spec/std.md) · [str.md](docs/spec/str.md) · [fmt.md](docs/spec/fmt.md) · [fs.md](docs/spec/fs.md) · [path.md](docs/spec/path.md) · [cli.md](docs/spec/cli.md) · [shell.md](docs/spec/shell.md) · [sort.md](docs/spec/sort.md) · [random.md](docs/spec/random.md) · [datetime.md](docs/spec/datetime.md) · [bitwise.md](docs/spec/bitwise.md) · [test.md](docs/spec/test.md) · [debug.md](docs/spec/debug.md) · [sys.md](docs/spec/sys.md) · [thread.md](docs/spec/thread.md) · [async-event-bus.md](docs/spec/async-event-bus.md) · [video.md](docs/spec/video.md) · [writes.md](docs/spec/writes.md) · [codecols.md](docs/spec/codecols.md) · [gpu.md](docs/spec/gpu.md) · [gpus.md](docs/spec/gpus.md) · [transpile.md](docs/spec/transpile.md)

</details>

<details>
<summary><strong>GUI & Graphics</strong> (58 docs)</summary>

[gui.md](docs/gui.md) · [gui-2d.md](docs/gui-2d.md) · [gui-3dmath.md](docs/gui-3dmath.md) · [gui-aa.md](docs/gui-aa.md) · [gui-accessibility.md](docs/gui-accessibility.md) · [gui-audio.md](docs/gui-audio.md) · [gui-canvas.md](docs/gui-canvas.md) · [gui-clipboard.md](docs/gui-clipboard.md) · [gui-color.md](docs/gui-color.md) · [gui-common.md](docs/gui-common.md) · [gui-dialogs.md](docs/gui-dialogs.md) · [gui-draw.md](docs/gui-draw.md) · [gui-draw-ops.md](docs/gui-draw-ops.md) · [gui-events.md](docs/gui-events.md) · [gui-filedrop.md](docs/gui-filedrop.md) · [gui-fonts.md](docs/gui-fonts.md) · [gui-form.md](docs/gui-form.md) · [gui-gamepad.md](docs/gui-gamepad.md) · [gui-gpu-info.md](docs/gui-gpu-info.md) · [gui-graph.md](docs/gui-graph.md) · [gui-input.md](docs/gui-input.md) · [gui-leak-detect.md](docs/gui-leak-detect.md) · [gui-lighting.md](docs/gui-lighting.md) · [gui-loop.md](docs/gui-loop.md) · [gui-menus.md](docs/gui-menus.md) · [gui-mesh.md](docs/gui-mesh.md) · [gui-native-linux.md](docs/gui-native-linux.md) · [gui-native-win.md](docs/gui-native-win.md) · [gui-native-win-close.md](docs/gui-native-win-close.md) · [gui-native-win-d3d11.md](docs/gui-native-win-d3d11.md) · [gui-native-win-ddraw.md](docs/gui-native-win-ddraw.md) · [gui-native-win-frame.md](docs/gui-native-win-frame.md) · [gui-native-win-icons.md](docs/gui-native-win-icons.md) · [gui-native-win-opengl.md](docs/gui-native-win-opengl.md) · [gui-native-win-poll.md](docs/gui-native-win-poll.md) · [gui-native-win-present.md](docs/gui-native-win-present.md) · [gui-native-win-probe.md](docs/gui-native-win-probe.md) · [gui-native-win-vulkan.md](docs/gui-native-win-vulkan.md) · [gui-native-win-vulkan-constants.md](docs/gui-native-win-vulkan-constants.md) · [gui-native-win-vulkan-init.md](docs/gui-native-win-vulkan-init.md) · [gui-native-win-vulkan-present.md](docs/gui-native-win-vulkan-present.md) · [gui-native-win-vulkan-swapchain.md](docs/gui-native-win-vulkan-swapchain.md) · [gui-print.md](docs/gui-print.md) · [gui-raster.md](docs/gui-raster.md) · [gui-render.md](docs/gui-render.md) · [gui-resolution.md](docs/gui-resolution.md) · [gui-shader.md](docs/gui-shader.md) · [gui-shader-codegen.md](docs/gui-shader-codegen.md) · [gui-shader-color.md](docs/gui-shader-color.md) · [gui-shader-math.md](docs/gui-shader-math.md) · [gui-shader-noise.md](docs/gui-shader-noise.md) · [gui-shader-sdf.md](docs/gui-shader-sdf.md) · [gui-systray.md](docs/gui-systray.md) · [gui-test.md](docs/gui-test.md) · [gui-theme.md](docs/gui-theme.md) · [gui-thread.md](docs/gui-thread.md) · [gui-video.md](docs/gui-video.md) · [gui-web.md](docs/gui-web.md)

</details>

<details>
<summary><strong>Audio</strong> (9 docs)</summary>

[audio.md](docs/audio.md) · [audio-aiff.md](docs/audio-aiff.md) · [audio-au.md](docs/audio-au.md) · [audio-complex.md](docs/audio-complex.md) · [audio-dsp.md](docs/audio-dsp.md) · [audio-fft.md](docs/audio-fft.md) · [audio-ogg.md](docs/audio-ogg.md) · [audio-raw.md](docs/audio-raw.md) · [audio-wav.md](docs/audio-wav.md)

</details>

<details>
<summary><strong>Networking</strong> (12 docs)</summary>

[http.md](docs/http.md) · [websocket.md](docs/websocket.md) · [socket.md](docs/socket.md) · [p2p.md](docs/p2p.md) · [smtp.md](docs/smtp.md) · [imap.md](docs/imap.md) · [pop.md](docs/pop.md) · [wlan.md](docs/wlan.md) · [email.md](docs/email.md) · [email-imap.md](docs/email-imap.md) · [email-pop.md](docs/email-pop.md) · [email-smtp.md](docs/email-smtp.md)

</details>

<details>
<summary><strong>Build System & Targets</strong> (17 docs)</summary>

[build.md](docs/build.md) · [build-analyze.md](docs/build-analyze.md) · [build-ast.md](docs/build-ast.md) · [build-config.md](docs/build-config.md) · [build-ident.md](docs/build-ident.md) · [build-imports.md](docs/build-imports.md) · [build-includes.md](docs/build-includes.md) · [build-render.md](docs/build-render.md) · [build-render-node.md](docs/build-render-node.md) · [bundle-ast.md](docs/bundle-ast.md) · [bundle-utils.md](docs/bundle-utils.md) · [mermaid.md](docs/mermaid.md) · [target-ast.md](docs/target-ast.md) · [target-bin.md](docs/target-bin.md) · [target-doc.md](docs/target-doc.md) · [target-java.md](docs/target-java.md) · [target-lua.md](docs/target-lua.md) · [target-ts.md](docs/target-ts.md)

</details>

<details>
<summary><strong>Virtual Machines & Code Gen</strong> (11 docs)</summary>

[Virtual-Bytecode.md](docs/Virtual-Bytecode.md) · [Virtual.md](docs/Virtual.md) · [VirtualToken.md](docs/VirtualToken.md) · [wasm-vm.md](docs/wasm-vm.md) · [wasm-vm-runtime.md](docs/wasm-vm-runtime.md) · [wasm.md](docs/wasm.md) · [runtime-codegen.md](docs/runtime-codegen.md) · [runtime-js.md](docs/runtime-js.md) · [runtime-native.md](docs/runtime-native.md) · [codegen-common.md](docs/codegen-common.md) · [engine-switching.md](docs/engine-switching.md)

</details>

<details>
<summary><strong>Concurrency & Events</strong> (2 docs)</summary>

[thread.md](docs/thread.md) · [async-event-bus.md](docs/async-event-bus.md)

</details>

<details>
<summary><strong>Syntax & Macros</strong> (8 docs)</summary>

[syntax.md](docs/syntax.md) · [syntax-parse.md](docs/syntax-parse.md) · [syntax-print.md](docs/syntax-print.md) · [syntax-tokenize.md](docs/syntax-tokenize.md) · [syntax-macros.md](docs/syntax-macros.md) · [syntaxfmt.md](docs/syntaxfmt.md) · [ast-analyze.md](docs/ast-analyze.md) · [ast-transform.md](docs/ast-transform.md)

</details>

<details>
<summary><strong>Platform & System</strong> (23 docs)</summary>

[windows.md](docs/windows.md) · [windows-constants.md](docs/windows-constants.md) · [windows-core.md](docs/windows-core.md) · [windows-flags.md](docs/windows-flags.md) · [windows-fonts.md](docs/windows-fonts.md) · [windows-gdi.md](docs/windows-gdi.md) · [windows-kernel.md](docs/windows-kernel.md) · [windows-loader.md](docs/windows-loader.md) · [windows-net.md](docs/windows-net.md) · [windows-registry.md](docs/windows-registry.md) · [windows-windowing.md](docs/windows-windowing.md) · [win-common.md](docs/win-common.md) · [linux.md](docs/linux.md) · [linux-constants.md](docs/linux-constants.md) · [linux-core.md](docs/linux-core.md) · [linux-fonts.md](docs/linux-fonts.md) · [linux-libc.md](docs/linux-libc.md) · [linux-loader.md](docs/linux-loader.md) · [linux-windowing.md](docs/linux-windowing.md) · [gpu.md](docs/gpu.md) · [gpus.md](docs/gpus.md) · [go.md](docs/go.md) · [sys.md](docs/sys.md)

</details>

<details>
<summary><strong>Data, Crypto & Utilities</strong> (32 docs)</summary>

[data.md](docs/data.md) · [data-csv.md](docs/data-csv.md) · [data-ini.md](docs/data-ini.md) · [data-toml.md](docs/data-toml.md) · [data-xml.md](docs/data-xml.md) · [data-yaml.md](docs/data-yaml.md) · [compression.md](docs/compression.md) · [compression-huffman.md](docs/compression-huffman.md) · [compression-lzw.md](docs/compression-lzw.md) · [compression-rle.md](docs/compression-rle.md) · [msgpack.md](docs/msgpack.md) · [json.md](docs/json.md) · [md.md](docs/md.md) · [crypto.md](docs/crypto.md) · [dataprot.md](docs/dataprot.md) · [bmp.md](docs/bmp.md) · [ico.md](docs/ico.md) · [ppm.md](docs/ppm.md) · [qoi.md](docs/qoi.md) · [tga.md](docs/tga.md) · [image.md](docs/image.md) · [image-bmp.md](docs/image-bmp.md) · [image-ico.md](docs/image-ico.md) · [image-ppm.md](docs/image-ppm.md) · [image-qoi.md](docs/image-qoi.md) · [image-tga.md](docs/image-tga.md) · [video.md](docs/video.md) · [datetime.md](docs/datetime.md) · [math.md](docs/math.md) · [math-geo.md](docs/math-geo.md) · [math-stats.md](docs/math-stats.md) · [math-base.md](docs/math-base.md)

</details>

<details>
<summary><strong>Core Language & Misc</strong> (23 docs)</summary>

[spec.md](docs/spec.md) · [cs.md](docs/cs.md) · [std.md](docs/std.md) · [str.md](docs/str.md) · [fmt.md](docs/fmt.md) · [fs.md](docs/fs.md) · [path.md](docs/path.md) · [cli.md](docs/cli.md) · [shell.md](docs/shell.md) · [shell-simple.md](docs/shell-simple.md) · [debug.md](docs/debug.md) · [error-display.md](docs/error-display.md) · [bitwise.md](docs/bitwise.md) · [transpile.md](docs/transpile.md) · [virtual-fs.md](docs/virtual-fs.md) · [pack.md](docs/pack.md) · [pack-utils.md](docs/pack-utils.md) · [ai.md](docs/ai.md) · [codecols.md](docs/codecols.md) · [random.md](docs/random.md) · [sort.md](docs/sort.md) · [test.md](docs/test.md) · [writes.md](docs/writes.md)

</details>

---

## Known Limitations

Magnolia is under active development. Some features are experimental or have known limitations:

- **Virtual (self-hosting) interpreter**: The Virtual library provides a self-hosted Oak interpreter written in Magnolia itself. Still being stabilized — may not support all language features yet.
- **Channel operations**: Low-level channel primitives for async communication are under development. Some edge cases in async patterns may not be fully supported.
- **Class inheritance**: Classes with constructors and static members are supported; multiple inheritance syntax is still being refined.
- **Memory operations**: Low-level `memread`/`memwrite` primitives are exposed for systems programming but require careful usage.

For the latest updates, check the [GitHub issues](https://github.com/SpcFORK/magnolia/issues) and [documentation](docs/).
