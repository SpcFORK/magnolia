<h1>
    <img width="24" height="24" alt="🌳" src="https://github.com/user-attachments/assets/22fadd8f-707e-4279-8ea0-63bd3da6fdba" />
    <em><b>⠀Magnolia</b></em>⠀🌸
</h1>

**Magnolia** is an expressive, dynamically typed programming language based on [Oak](https://oaklang.org/). It extends Oak with a cross-platform GUI system, audio/DSP processing, full networking stack, concurrency primitives, multi-target compilation (JS, WASM, Lua, TypeScript), GPU interop, AST macros, a self-hosted bytecode VM, and 140+ standard library modules — while keeping the simplicity and elegance of the original language.

> **Highlights:** 3.7x faster bytecode VM · Cross-platform GUI (Win32/X11/WebGL) with Vulkan · CPU shader engine · P2P mesh networking · Thread pools & async event bus · 8 compilation targets · Self-hosted bytecode VM · 67 sample programs

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

### 🌍 Multi-Target Compilation

Magnolia programs can be compiled to multiple targets from a single codebase:

```sh
# Native Oak bundle
magnolia build --entry src/main.oak --output dist/bundle.oak

# JavaScript (browser / Node.js / Deno)
magnolia build --entry src/app.oak --output dist/bundle.js --web

# WebAssembly
magnolia build --entry src/main.oak --output dist/program.wat --wasm
```

The build system supports module bundling, dependency resolution, tree-shaking, code minification, and virtual filesystem embedding. See [docs/build.md](docs/build.md) and [docs/wasm.md](docs/wasm.md).

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

## Overview

Magnolia has 8 primitive and 3 complex types.

```js
?        // null, also "()"
_        // "empty" value, equal to anything
1, 2, 3  // integers
3.14     // floats
true     // booleans
'hello'  // strings
:error   // atoms
pointer(0) // pointers

[1, :number]    // list
{ a: 'hello' }  // objects
fn(a, b) a + b  // functions
```

These types mostly behave as you'd expect. Some notable details:

- There is no implicit type casting between any types, except during arithmetic operations when ints may be cast up to floats.
- Both ints and floats are full 64-bit.
- Strings are mutable byte arrays, also used for arbitrary data storage in memory, like in Lua. For immutable strings, use atoms.
- Lists are backed by a vector data structure -- appending and indexing is cheap, but cloning is not
- For lists and objects, equality is defined as deep equality. There is no identity equality in Magnolia.

We define a function in Magnolia with the `fn` keyword. A name is optional, and if given, will define that function in that scope. If there are no arguments, the `()` may be omitted.

```js
fn double(n) 2 * n
fn speak {
    println('Hello!')
}
```

Besides the normal set of arithmetic operators, Magnolia has a few strange operators.

- The **assignment operator** `:=` binds values on the right side to names on the left, potentially by destructuring an object or list. For example:

    ```js
    a := 1              // a is 1
    [b, c] := [2, 3]    // b is 2, c is 3
    d := double(a)      // d is 2
    ```
- The **nonlocal assignment operator** `<-` binds values on the right side to names on the left, but only when those variables already exist. If the variable doesn't exist in the current scope, the operator ascends up parent scopes until it reaches the global scope to find the last scope where that name was bound.

    ```js
    n := 10
    m := 20
    {
        n <- 30
        m := 40
    }
    n // 30
    m // 20
    ```
- **Arithmetic operators**: `+`, `-`, `*`, `/`, `%` for basic math, and `**` for exponentiation.

    ```js
    2 + 3       // 5
    10 / 3      // 3.333...
    2 ** 8      // 256 (2 to the power of 8)
    ```

- **Bitwise operators** for low-level bit manipulation: `&` (AND), `|` (OR), `^` (XOR), `~` (NOT), `<<` (left shift), and `>>` (right shift).

    ```js
    0xFF & 0x0F        // 15
    0x01 | 0x02        // 3
    5 ^ 3              // 6
    ~10                // -11
    4 << 2             // 16 (multiply by 2^2)
    16 >> 2            // 4 (divide by 2^2)
    ```

- The **push operator** `<<` (in list/string context) pushes values onto the end of a string or a list, mutating it, and returns the changed string or list.

    ```js
    str := 'Hello '
    str << 'World!' // 'Hello World!'

    list := [1, 2, 3]
    list << 4
    list << 5 << 6 // [1, 2, 3, 4, 5, 6]
    ```
- The **pipe operator** `|>` takes a value on the left and makes it the first argument to a function call on the right.

    ```js
    // print 2n for every prime n in range [0, 10)
    range(10) |> filter(prime?) |>
        each(double) |> each(println)

    // adding numbers
    fn add(a, b) a + b
    10 |> add(20) |> add(3) // 33
    ```

Magnolia uses one main construct for control flow -- the `if` match expression. Unlike a traditional `if` expression, which can only test for truthy and falsy values, Magnolia's `if` acts like a sophisticated switch-case, comparing values until the right match is reached.

```js
fn pluralize(word, count) if count {
    1 -> word
    2 -> 'a pair of ' + word
    _ -> word + 's'
}
```

This match expression, combined with safe tail recursion, makes Magnolia Turing-complete.

Magnolia also provides **class syntax sugar** for creating constructor functions with the `cs` keyword. Classes are syntactic sugar that make it easier to create objects with shared state and methods.

```js
// Class without parameters
cs Empty {
    {}
}
type(Empty()) // :object

// Class constructor parameters are captured in body
cs Pair(left, right) {
    {
        left: left
        right: right
    }
}
Pair(1, 2).right // 2

// Class methods can close over constructor state
cs Counter(start) {
    {
        value: start
        add: fn(delta) start + delta
    }
}
Counter(4).add(3) // 7

// Classes support variadic parameters
cs Bag(items...,) {
    items
}
len(Bag(1, 2, 3)) // 3

// Assignment-only class bodies build instance fields directly
cs LocalCounter {
    a := 2
    set := fn {
        a <- 3
    }
}
counter := LocalCounter()
counter.set()
counter.a // 3
```

Key features of classes:
- **Constructor sugar**: Classes without parameters act as simple constructor functions that return objects
- **Parameter capture**: Constructor parameters are available in the class body and can be used to initialize object properties
- **Closure over state**: Methods defined in the class body can close over constructor parameters
- **Variadic support**: Classes support variadic parameters using the `...` syntax
- **Assignment-only sugar**: In assignment-only class bodies, method reads/writes target the constructed instance fields
- **Return value**: Classes with an empty block body (`{}`) return `?` (null), while classes with an object expression return that object

Under the hood, classes are simply functions that return objects, but the `cs` syntax provides a cleaner way to define object constructors with shared behavior.

Lastly, because callback-based asynchronous concurrency is common in Magnolia, there's special syntax sugar, the `with` expression, to help. The `with` syntax sugar de-sugars like this.

```js
with readFile('./path') fn(file) {
    println(file)
}

// desugars to
readFile('./path', fn(file) {
    println(file)
})
```

For a more detailed description of the language, see the [work-in-progress language spec](docs/spec.md).

For Magnolia-specific features, see:
- [GUI middleware](docs/gui.md) · [2D drawing](docs/gui-2d.md) · [Shaders](docs/gui-shader.md) · [Resolution scaling](docs/gui-resolution.md) · [Events](docs/gui-events.md)
- [Audio](docs/audio.md) · [WAV encoding](docs/audio-wav.md) · [DSP](docs/audio-dsp.md) · [FFT](docs/audio-fft.md)
- [HTTP server](docs/http.md) · [WebSockets](docs/websocket.md) · [TCP sockets](docs/socket.md) · [P2P mesh](docs/p2p.md)
- [SMTP](docs/smtp.md) · [IMAP](docs/imap.md) · [WLAN discovery](docs/wlan.md)
- [Compression (RLE/Huffman/LZW)](docs/compression.md) · [MessagePack](docs/msgpack.md)
- [BMP images](docs/bmp.md) · [ICO icons](docs/ico.md) · [Video frames](docs/video.md)
- [Crypto (UUID, SHA-256, random)](docs/crypto.md) · [Data protection (CRC, checksums)](docs/dataprot.md)
- [Virtual Bytecode VM](docs/Virtual-Bytecode.md) · [Virtual interpreter](docs/Virtual.md) · [VirtualToken constructors](docs/VirtualToken.md)
- [Virtual File System](docs/virtual-fs.md) · [Pack/bundle](docs/pack.md)
- [Build system](docs/build.md) · [WASM target](docs/wasm.md) · [WASM VM](docs/wasm-vm.md)
- [Thread library](docs/thread.md) · [Async event bus](docs/async-event-bus.md)
- [Transpile middleware](docs/transpile.md) · [Syntax and macros](docs/syntax.md)
- [Code generation](docs/runtime-codegen.md) · [Go runtime interop](docs/go.md) · [System interop (`sys`)](docs/sys.md)
- [GPU](docs/gpu.md) · [Multi-backend GPU helpers (`gpus`)](docs/gpus.md)
- [Math](docs/math.md) · [Geometry](docs/math-geo.md) · [Statistics](docs/math-stats.md) · [Math base](docs/math-base.md)
- [Bitwise and pointer helpers](docs/bitwise.md) · [Classes (`cs`)](docs/cs.md) · [String manipulation](docs/str.md)
- [Windows platform](docs/windows.md) · [Linux platform](docs/linux.md)
- [Error display](docs/error-display.md) · [Debug helpers](docs/debug.md) · [JSON](docs/json.md) · [DateTime](docs/datetime.md)
- Example programs in [samples/](samples/) including GUI, threading, transpilation, VFS, and pointer/bitwise examples

### Builds and deployment

While the Magnolia interpreter can run programs and modules directly from source code on the file system, Magnolia also offers a build tool, `build`, which can _bundle_ a Magnolia program distributed across many files into a single "bundle" source file. The same command can also cross-compile Magnolia bundles into JavaScript bundles, to run in the browser or in JavaScript environments like Node.js and Deno. This allows Magnolia programs to be deployed and distributed as single-file programs, both on the server and in the browser.

To build a new bundle, we can simply pass an "entrypoint" to the program.

```sh
magnolia build --entry src/main.oak --output dist/bundle.oak
```

Compiling to JavaScript works similarly, but with the `--web` flag, which turns on JavaScript cross-compilation.

```sh
magnolia build --entry src/app.js.oak --output dist/bundle.js --web
```

The bundler and compiler are built on top of past work with the [September](https://github.com/thesephist/september) toolchain for Ink, but slightly re-architected to support bundling and multiple compilation targets. In the future, the goal of `build` is to become a lightly optimizing compiler and potentially help yield a `compile` command that could package the interpreter and a Magnolia bundle into a single executable binary. For more information, see `magnolia help build` (or `oak help build`, depending on your install name).

### Performance

Magnolia inherits Oak's performance characteristics. As of September 2021, Oak is about 5-6x slower than Python 3.9 on pure function call and number-crunching overhead (assessed by a basic `fib(30)` benchmark). These are worst-case estimates — the language's simpler data structures narrow the gap on realistic programs. Being as fast as Python and Ruby is a good long-term goal.

Runtime performance is not the primary concern today; correctness and a pleasant developer experience come first. That said, significant work has already landed, and there's a clear roadmap ahead.

#### Done

1. ~~Constant folding and propagation on the AST.~~ **Implemented** — constant arithmetic, string concatenation, boolean logic, unary operations, and equality comparisons on literals are resolved at compile time.
2. ~~Memory allocation audit and optimization pass.~~ **Implemented** — pprof-guided pass yielding a **2.4x speedup** on fib(20) (35.9ms → 14.6ms) and ~6.6% fewer heap allocations. Key wins: `sync.Pool` for scope mutexes, pre-sized scope maps (`newScopeN`), direct fn-call fast path (no args slice / thunk wrapping for the common case), index-based tokenizer (no per-token `[]rune` allocator), and parser capacity hints.
3. ~~Bytecode VM.~~ **Implemented** — 52-opcode stack-based VM accessible via `--bytecode`. Supports closures, recursion, destructuring, pattern matching, pipes, and upvalue mutation. **3.7x speedup** on fib(30), **1.9x** on loop-heavy workloads vs. the tree-walking interpreter.

#### TODO

6. **Register-based bytecode VM** — introduce a register-based design to cut push/pop overhead side-by-side with the stack-based dispatch and enable easier peephole optimizations and more.
11. **SIMD-accelerated string / list builtins** — use platform SIMD intrinsics for `find`, `map`, `filter`, and bulk string operations on contiguous memory.
10. **JIT compilation of hot bytecode paths** — profile the bytecode VM at runtime, identify hot traces, and compile them to native machine code via a lightweight JIT backend.
9. **Parallel GC or arenas for short-lived scopes** — reduce stop-the-world GC pauses by giving hot inner loops their own allocation arena that can be freed in bulk.
7. **Escape analysis and stack allocation** — identify closures and objects that don't outlive their creating scope and allocate them on the stack instead of the heap.
8. **Lazy / copy-on-write strings and lists** — defer cloning until mutation, reducing allocation pressure for pipe-heavy functional code.
4. **Inline caching for property access** — cache the offset of frequently accessed object keys so repeated lookups (e.g., `obj.x` inside a loop) skip the map hash.
5. **NaN-boxing or tagged-pointer value representation** — pack small ints, bools, and atoms into a single 64-bit word to eliminate per-value heap allocations and improve cache locality.

## Development

Magnolia (ab)uses GNU Make to run development workflows and tasks.

- `make run` compiles and runs the Magnolia binary, which opens an interactive REPL
- `make fmt` or `make f` runs the formatter over files with unstaged changes (currently via `oak fmt --changes --fix` for compatibility)
- `make tests` or `make t` runs the Go test suite for the Magnolia language and interpreter
- `make test-oak` or `make tk` runs the Magnolia test suite, which tests the standard libraries
- `make test-bundle` runs the Magnolia test suite, bundled via the `build` command
- `make test-js` runs the Magnolia test suite on Node.js, compiled with the `--web` target
- `make build` generates release builds of Magnolia for various operating systems; `make build-<OS>` builds for a specific OS
- `make install` installs the interpreter on your `$GOPATH` as `oak`, and re-installs the Vim syntax file
- `make site` builds a Magnolia bundle for the website, and `make site-w` does it on every file save
- `make site-gen` rebuilds the statically generated parts of the Magnolia website, like the standard library documentation

On Windows, if `-race` commands fail in local development due to CGO/toolchain constraints, use non-race alternatives:

- `go test .`
- `go run . test/main.oak`

To try Magnolia by building from source, clone the repository and run `make install` (or simply `go build .`).

## Known Limitations

Magnolia is under active development. Some features are experimental or have known limitations:

- **Virtual (self-hosting) interpreter**: The Virtual library provides a self-hosted Oak interpreter written in Magnolia itself, enabling dynamic code evaluation at runtime. This feature is still being stabilized and may not support all language features yet.

- **Channel operations and async communication**: Low-level channel primitives for asynchronous communication are under development. Some edge cases in async patterns may not be fully supported.

- **Class inheritance**: While classes with constructors and static members are supported, multiple inheritance syntax is still being refined.

- **Memory operations**: Low-level memory read/write primitives (`memread`, `memwrite`) are exposed for systems programming but require careful usage.

For the latest updates and progress on these features, please check the [GitHub issues](https://github.com/SpcFORK/magnolia/issues) and [documentation](docs/).

## Unit and generative tests

The Magnolia repository has two kinds of tests: unit tests and generative/fuzz tests. **Unit tests** are just what they sound like -- tests validated with assertions -- and are built on the `libtest` Magnolia library with the exception of Go tests in `eval_test.go`. **Generative tests** include fuzz tests, and are tests that run some pre-defined behavior of functions through a much larger body of procedurally generated inputs, for validating behavior that's difficult to validate manually like correctness of parsers and `libdatetime`'s date/time conversion algorithms.

Both sets of tests are written and run entirely in the "userland" of Magnolia, without invoking the interpreter separately. Unit tests live in `./test` and are run with `./test/main.oak`; generative tests are in `test/generative`, and can be run manually.
