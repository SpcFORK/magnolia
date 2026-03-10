# Build Library (build)

## Overview

`libbuild` is Oak's compiler and bundler, transforming Oak source code into standalone executables or JavaScript bundles for various target platforms.

## Import

```oak
build := import('build')
{ configure: configure, compile: compile } := import('build')
```

## Features

- **Multi-target compilation**: Oak native, JavaScript (web), WebAssembly
- **Module bundling**: Combines multiple Oak files into single output
- **Dependency resolution**: Automatically includes imported modules
- **Virtual filesystem**: Embeds resources into executables
- **Minification**: Optimizes output for production
- **Tree shaking**: Removes unused code

## Functions

### `configure(config)`

Configures the build system with compilation options.

**Parameters (config object):**
- `entry` - Entry point file path (required)
- `output` - Output file path (required)
- `web?` - Compile to JavaScript (boolean)
- `wasm?` - Compile to WebAssembly (boolean)
- `includes` - List of additional modules to include
- `includeVFS` - Files to embed in virtual filesystem
- `log` - Logging function (default: `printf`)
- `abort` - Error handler function

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

### Oak Native

Default compilation target, produces Oak bytecode.

```oak
configure({
    entry: 'main.oak'
    output: 'program.oak'
})
```

### JavaScript (Web)

Transpiles Oak to JavaScript for browser/Node.js.

```oak
configure({
    entry: 'main.oak'
    output: 'bundle.js'
    web?: true
})
```

### WebAssembly

Compiles Oak to WebAssembly (experimental).

```oak
configure({
    entry: 'main.oak'
    output: 'program.wasm'
    wasm?: true
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

1. **Parse** entry point and dependencies
2. **Analyze** AST for semantic information
3. **Transform** AST for target platform
4. **Bundle** modules into single tree
5. **Codegen** target code (Oak/JS/Wasm)
6. **Optimize** output (tree shaking, minification)
7. **Write** output file

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

## Limitations

- No incremental compilation (rebuilds all files)
- Limited error reporting (line/column info may be approximate)
- WebAssembly target is experimental
- No source maps generated
- Tree shaking requires explicit imports

## See Also

- [pack.md](pack.md) - Create standalone executables
- [transpile.md](transpile.md) - AST transformation middleware
- [bundle-ast.md](bundle-ast.md) - Module bundling
- [virtual-fs.md](virtual-fs.md) - Virtual filesystem documentation
