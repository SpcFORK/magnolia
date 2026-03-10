# Magnolia 🌸

**Magnolia** is an expressive, dynamically typed programming language based on [Oak](https://oaklang.org/). It extends Oak with powerful new features including a transpile middleware system, virtual file system, advanced threading utilities, and GPU computing support, while maintaining the simplicity and elegance of the original language.

Here's an example Magnolia program.

```js
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

Magnolia has good support for asynchronous I/O. Here's how you read a file and print it.

```js
std := import('std')
fs := import('fs')

with fs.readFile('./file.txt') fn(file) if file {
    ? -> std.println('Could not read file!')
    _ -> print(file)
}
```

Magnolia also has a pragmatic standard library that comes built into the interpreter executable. For example, there's a built-in HTTP server and router in the `http` library.

```js
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

## Install

Magnolia is currently installed from source.

On Unix-like systems, build with Make:

```sh
make install
```

On Windows, use the provided build script:

```bat
build.bat
```

Or build directly with Go on any platform:

```sh
go build .
```

You can also run without installing:

```sh
go run . <file-or-command>
```

## What's New in Magnolia

Magnolia extends Oak with powerful new features for modern development:

### 🎨 Enhanced Error Display

Beautiful, color-coded error messages with source code context to help you quickly identify and fix issues:

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

See [error-display.md](docs/error-display.md) for more details.

### 🔧 Transpile Middleware

A plugin architecture for AST transformations during the build process. Write custom transpilers to transform your code at compile-time:

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

### 📁 Virtual File System

An in-memory file system that can be embedded in packed binaries, enabling true cross-platform deployment:

```js
Virtual := import('Virtual')

vfs := Virtual.createVirtualFS({
    'config.json': '{"version": "1.0"}'
    'data/test.txt': 'test data'
})

content := vfs.readFile('config.json')
vfs.writeFile('output.txt', 'Hello World')
```

### 🧵 Thread Library

High-level utilities for concurrent and parallel programming, including mutexes, semaphores, wait groups, and thread pools:

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

### ✨ AST Macros and Metaprogramming

The syntax library now supports AST macros for powerful compile-time code transformations:

```js
syntax := import('syntax')

// Define macro expanders
myMacro := syntax.Macro(fn(node) {
    // Transform AST nodes
    node
})

// Parse code with macro expansion
ast := syntax.parseWithMacros('(my-macro 1 2 3)', [myMacro])

// Recursively expand macros in AST
expanded := syntax.expandMacros(ast, [myMacro])
```

## Overview

Magnolia has 7 primitive and 3 complex types.

```js
?        // null, also "()"
_        // "empty" value, equal to anything
1, 2, 3  // integers
3.14     // floats
true     // booleans
'hello'  // strings
:error   // atoms

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
```

Key features of classes:
- **Constructor sugar**: Classes without parameters act as simple constructor functions that return objects
- **Parameter capture**: Constructor parameters are available in the class body and can be used to initialize object properties
- **Closure over state**: Methods defined in the class body can close over constructor parameters
- **Variadic support**: Classes support variadic parameters using the `...` syntax
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
- [Virtual File System documentation](docs/virtual-fs.md)
- [Transpile Middleware documentation](docs/transpile.md)
- [Code Generation documentation](docs/runtime-codegen.md)
- [Go Runtime and System Interop documentation](docs/go.md)
- [Classes (`cs`) documentation](docs/cs.md)
- [Syntax and Macros documentation](docs/syntax.md)
- [Advanced Build Features](docs/build.md)
- [String manipulation library](docs/str.md)
- Example programs in [samples/](samples/) including threading, transpilation, and VFS examples

### Builds and deployment

While the Magnolia interpreter can run programs and modules directly from source code on the file system, Magnolia also offers a build tool, `magnolia build`, which can _bundle_ a Magnolia program distributed across many files into a single "bundle" source file. `magnolia build` can also cross-compile Magnolia bundles into JavaScript bundles, to run in the browser or in JavaScript environments like Node.js and Deno. This allows Magnolia programs to be deployed and distributed as single-file programs, both on the server and in the browser.

To build a new bundle, we can simply pass an "entrypoint" to the program.

```sh
magnolia build --entry src/main.oak --output dist/bundle.oak
```

Compiling to JavaScript works similarly, but with the `--web` flag, which turns on JavaScript cross-compilation.

```sh
magnolia build --entry src/app.js.oak --output dist/bundle.js --web
```

The bundler and compiler are built on top of past work with the [September](https://github.com/thesephist/september) toolchain for Ink, but slightly re-architected to support bundling and multiple compilation targets. In the future, the goal of `magnolia build` is to become a lightly optimizing compiler and potentially help yield a `magnolia compile` command that could package the interpreter and a Magnolia bundle into a single executable binary. For more information on `magnolia build`, see `magnolia help build`.

### Performance

Magnolia inherits Oak's performance characteristics. As of September 2021, Oak is about 5-6x slower than Python 3.9 on pure function call and number-crunching overhead (assessed by a basic `fib(30)` benchmark). These figures are worst-case estimates -- because the language's data structures are far simpler than Python's, the ratios start to go down on more realistic complex programs. But nonetheless, this gives a good estimate of the kind of performance you can expect from Magnolia programs.

Runtime performance is not currently the primary concern; the primary concern is implementing a correct and pleasant interpreter that's fast _enough_ to write real apps with. Being as fast as Python and Ruby is a good long-term goal. Those languages run in production and receive continuous investments into performance tuning.

There are several immediately actionable things we can do to speed up Magnolia programs' runtime performance, though none are under works today. In order of increasing implementation complexity:

1. Basic compiler optimization techniques applied to the abstract syntax tree, like constant folding and propagation.
2. A thorough audit of the interpreter's memory allocation profile and a memory optimization pass (and the same for L1/L2 cache misses).
3. A bytecode VM that executes Magnolia compiled down to more compact and efficient bytecode rather than a syntax tree-walking interpreter.

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

To try Magnolia by building from source, clone the repository and run `make install` (or simply `go build .`).

## Known Limitations

Magnolia is under active development. Some features are experimental or have known limitations:

- **Bitwise right shift operator (`>>`)**: The right shift operator has a known syntax conflict with template syntax and is not fully functional. Users experiencing issues should use alternative approaches or workarounds.

- **Virtual (self-hosting) interpreter**: The Virtual library provides a self-hosted Oak interpreter written in Magnolia itself, enabling dynamic code evaluation at runtime. This feature is still being stabilized and may not support all language features yet.

- **Channel operations and async communication**: Low-level channel primitives for asynchronous communication are under development. Some edge cases in async patterns may not be fully supported.

- **Class inheritance**: While classes with constructors and static members are supported, multiple inheritance syntax is still being refined.

- **Memory operations**: Low-level memory read/write primitives (`memread`, `memwrite`) are exposed for systems programming but require careful usage.

For the latest updates and progress on these features, please check the [GitHub issues](https://github.com/SpcFORK/magnolia/issues) and [documentation](docs/).

## Unit and generative tests

The Magnolia repository has two kinds of tests: unit tests and generative/fuzz tests. **Unit tests** are just what they sound like -- tests validated with assertions -- and are built on the `libtest` Magnolia library with the exception of Go tests in `eval_test.go`. **Generative tests** include fuzz tests, and are tests that run some pre-defined behavior of functions through a much larger body of procedurally generated inputs, for validating behavior that's difficult to validate manually like correctness of parsers and `libdatetime`'s date/time conversion algorithms.

Both sets of tests are written and run entirely in the "userland" of Magnolia, without invoking the interpreter separately. Unit tests live in `./test` and are run with `./test/main.oak`; generative tests are in `test/generative`, and can be run manually.
