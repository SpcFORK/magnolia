# Build Library (build)

## Overview

`libbuild` is Oak's compiler and bundler, transforming Oak source code into standalone executables or JavaScript bundles for various target platforms.

## Import

```oak
build := import('build')
{ configure: configure, compile: compile } := import('build')
```

## Features

- **Multi-target compilation**: Oak native, JavaScript (web), WebAssembly, AST JSON, bytecode binary, TypeScript, Lua 5.4, Java 17+
- **Documentation generation**: Auto-generate Markdown API docs from source
- **Module bundling**: Combines multiple Oak files into single output
- **Dependency resolution**: Automatically includes imported modules
- **Virtual filesystem**: Embeds resources into executables
- **Transpiler middleware**: Pluggable AST-to-AST transpilers
- **Constant-fold optimizer**: Optional compile-time evaluation
- **Minification**: Optimizes output for production
- **Tree shaking**: Removes unused code

## Functions

### `configure(config)`

Configures the build system with compilation options.

**Parameters (config object):**
- `entry` - Entry point file path (required). Probes extensions `.oak`, `.ok`, `.mag`, `.mg` if no extension given
- `output` - Output file path (required)
- `web?` - Compile to JavaScript (boolean)
- `wasm?` - Compile to WebAssembly (boolean)
- `ast?` - Emit bundled AST as JSON (boolean)
- `bin?` - Emit compact bytecode binary `.mgb` (boolean)
- `doc?` - Generate Markdown API documentation (boolean)
- `ts?` - Compile to TypeScript (boolean)
- `lua?` - Transpile to Lua 5.4 (boolean)
- `java?` - Transpile to Java 17+ (boolean)
- `includes` - List of additional modules to include
- `includeVFS` - Files to embed in virtual filesystem
- `optimize?` - Enable constant-fold optimizer (boolean, default: `true`)
- `transpile?` - Enable transpiler middleware (boolean, default: `true`)
- `transpileVerbose?` - Verbose transpiler logging (boolean, default: `false`)
- `transpilers` - List of extra transpiler objects to register
- `log` - Logging function (default: `printf`)
- `abort` - Error handler function

> Target flags are mutually exclusive — only one may be set at a time. If none is set, the default Oak native target is used.

```oak
{ configure: configure } := import('build')

configure({
    entry: 'src/main.oak'
    output: 'dist/app.js'
    web?: true
    includes: ['lib/utils.oak', 'lib/helpers.oak']
    includeVFS: 'assets/'
})
```

### `compile()`

Executes the build process based on configuration.

```oak
{ configure: configure, compile: compile } := import('build')

configure({ entry: 'main.oak', output: 'out.js', web?: true })
compile()
```

## Build Targets

### Oak Native (default)

Bundles all modules into a single native Oak source file with a runtime preamble and module system. Self-contained; runs under the Oak/Magnolia interpreter.

```oak
configure({
    entry: 'main.oak'
    output: 'program.oak'
})
```

### JavaScript (Web)

Transpiles Oak to JavaScript for browser/Node.js/Deno. Escapes reserved words, handles `?`/`!` in identifiers, and adds trampoline-based tail-call elimination.

```oak
configure({
    entry: 'main.oak'
    output: 'bundle.js'
    web?: true
})
```

### WebAssembly

Compiles Oak to WebAssembly text format (experimental).

```oak
configure({
    entry: 'main.oak'
    output: 'program.wasm'
    wasm?: true
})
```

### AST (JSON)

Serializes the bundled AST as JSON for external tooling, analysis, or code generation pipelines.

```oak
configure({
    entry: 'main.oak'
    output: 'bundle.json'
    ast?: true
})
```

### Binary (Bytecode)

Compiles the bundle to a compact bytecode binary format (`.mgb`). Uses a `MGbc` header with version `2`, containing a constant pool, function table, and top-level names. Designed for the Go VM with fast startup.

```oak
configure({
    entry: 'main.oak'
    output: 'program.mgb'
    bin?: true
})
```

### Documentation (Markdown)

Generates Markdown API documentation by walking the bundled AST. Extracts module names, function signatures (including rest args), and constant declarations.

```oak
configure({
    entry: 'lib/mylib.oak'
    output: 'docs/mylib.md'
    doc?: true
})
```

### TypeScript

Wraps JavaScript output with TypeScript type stubs (`OakVal`, `OakFn`, `OakList`, `OakObject`), prefixed with `@ts-nocheck`. Useful for TypeScript interop.

```oak
configure({
    entry: 'main.oak'
    output: 'bundle.ts'
    ts?: true
})
```

### Lua 5.4

Full AST-to-Lua 5.4 transpiler. Includes a Lua runtime preamble (`__oak_eq`, `__oak_push`, `__oak_acc`, `__oak_modularize`, etc.), reserved word escaping, and `if`-expression-to-IIFE translation.

```oak
configure({
    entry: 'main.oak'
    output: 'bundle.lua'
    lua?: true
})
```

### Java 17+

Full AST-to-Java transpiler. Generates a `public class OakBundle` with a `main()` entry point, HashMap-based module system, and `Function<Object[], Object>` lambdas for closures.

```oak
configure({
    entry: 'main.oak'
    output: 'OakBundle.java'
    java?: true
})
```

## Include Patterns

### Include Specific Modules

```oak
configure({
    entry: 'main.oak'
    output: 'app.js'
    includes: ['lib/math.oak', 'lib/utils.oak']
})
```

### Include with Aliases

```oak
configure({
    entry: 'main.oak'
    output: 'app.js'
    includes: [
        { name: 'utils', path: 'lib/utilities.oak' }
        { name: 'helpers', path: 'lib/help.oak' }
    ]
})
```

### Include Virtual Filesystem

Embed files accessible via `Virtual` library:

```oak
configure({
    entry: 'main.oak'
    output: 'app.js'
    includeVFS: 'assets/'
})
```

## Usage Examples

### Simple Build

```oak
{ configure: configure, compile: compile } := import('build')

configure({
    entry: 'src/app.oak'
    output: 'dist/app.js'
    web?: true
})

compile()
println('Build complete!')
```

### Build with Dependencies

```oak
{ configure: configure, compile: compile } := import('build')

configure({
    entry: 'main.oak'
    output: 'bundle.js'
    web?: true
    includes: 'lib/http.oak,lib/json.oak'
})

compile()
```

### Build Script with Logging

```oak
{ configure: configure, compile: compile } := import('build')
{ printf: printf } := import('fmt')

configure({
    entry: 'src/main.oak'
    output: 'dist/app.js'
    web?: true
    log: fn(msg) {
        printf('[BUILD] {{0}}', msg)
    }
    abort: fn(msg) {
        printf('[ERROR] {{0}}', msg)
        exit(1)
    }
})

printf('Starting build...')
compile()
printf('Build successful!')
```

## Build Process

1. **Configure** — apply config, validate mutually exclusive target flags
2. **Resolve** imports from includes and entry point (probes `.oak`/`.ok`/`.mag`/`.mg`)
3. **Compile** modules in parallel (`pmap`) — parse → transform → transpile → analyze → wrap
4. **Bundle** compiled modules into single AST tree
5. **Render** bundle for the selected target (Oak/JS/Wasm/AST/Bin/Doc/TS/Lua/Java)
6. **Write** output file

## Virtual Filesystem

Files included via `includeVFS` become accessible through the `Virtual` library:

```oak
// Build config
configure({
    includeVFS: 'templates/'  // Includes templates/*.html
})

// Runtime access
Virtual := import('Virtual')
template := Virtual.readFile('templates/index.html')
```

## Implementation Notes

- Uses `syntax` library for parsing
- Applies AST transformations via `ast-transform`
- Bundles modules via `bundle-ast`
- Transpiles via `transpile` middleware
- Resolves file paths relative to entry point

## CLI Usage

```
oak build --entry [src] --output [dest] [options]
```

**Target flags** (mutually exclusive, default is Oak):

| Flag | Output | Description |
|------|--------|-------------|
| _(none)_ | `.oak` | Bundled native Oak source |
| `--web` | `.js` | JavaScript (browsers, Node.js, Deno) |
| `--wasm` | `.wasm` | WebAssembly text format (experimental) |
| `--ast` | `.json` | Bundled AST as JSON |
| `--bin` | `.mgb` | Compact bytecode binary (Go VM) |
| `--doc` | `.md` | Markdown API documentation |
| `--ts` | `.ts` | TypeScript with type stubs |
| `--lua` | `.lua` | Lua 5.4 |
| `--java` | `.java` | Java 17+ |

**Options:**
- `--entry` — Entrypoint for the bundle
- `--output` / `-o` — Output file path
- `--include` — Comma-separated list of modules to include explicitly

## Limitations

- No incremental compilation (rebuilds all files)
- Limited error reporting (line/column info may be approximate)
- WebAssembly target is experimental
- No source maps generated
- Tree shaking requires explicit imports

## Parallel Module Compilation

The build pipeline now uses `pmap` from the `thread` library to compile modules in parallel. Module keys are sorted and then each module is parsed, transformed, transpiled, analyzed, and wrapped concurrently, improving build times on multi-core systems.

## See Also

- [pack.md](pack.md) - Create standalone executables
- [transpile.md](transpile.md) - AST transformation middleware
- [bundle-ast.md](bundle-ast.md) - Module bundling
- [virtual-fs.md](virtual-fs.md) - Virtual filesystem documentation
