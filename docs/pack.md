# Pack Library (pack)

## Overview

`libpack` packages Oak programs as standalone, self-contained executables by bundling source code, dependencies, and resources into a single binary.

## Import

```oak
pack := import('pack')
{ pack: pack, configure: configure } := import('pack')
```

## Features

- **Standalone binaries**: Creates executable files with embedded Oak runtime
- **Dependency bundling**: Automatically includes all imported modules
- **Virtual filesystem**: Embeds resources accessible at runtime
- **Cross-platform**: Generates executables for target platforms
- **Zero dependencies**: Resulting binary requires no external files

## Functions

### `configure(config)`

Configures the pack system.

**Parameters (config object):**
- `entry` - Entry point file (required)
- `output` - Output executable path (required)
- `includes` - Additional modules to bundle
- `includeVFS` - Files/directories for virtual filesystem
- `log` - Logging function
- `abort` - Error handler

```oak
{ configure: configure } := import('pack')

configure({
    entry: 'src/main.oak'
    output: 'myapp'
    includeVFS: 'assets/'
})
```

### `pack()`

Executes the packing process.

```oak
{ configure: configure, pack: pack } := import('pack')

configure({
    entry: 'app.oak'
    output: 'dist/myapp'
})

pack()
println('Executable created!')
```

## Virtual Filesystem Patterns

### Include Single Directory

```oak
configure({
    entry: 'main.oak'
    output: 'app'
    includeVFS: 'templates/'
})

// Runtime access:
// Virtual.readFile('templates/index.html')
```

### Include Multiple Directories

```oak
configure({
    entry: 'main.oak'
    output: 'app'
    includeVFS: [
        'templates/'
        'static/css/'
        'static/js/'
    ]
})
```

### Include with Custom Paths

```oak
configure({
    entry: 'main.oak'
    output: 'app'
    includeVFS: [
        'views:templates/'       // templates/ → views/
        'assets:static/'         // static/ → assets/
    ]
})

// Runtime: Virtual.readFile('views/index.html')
```

## Usage Examples

### Simple Executable

```oak
{ configure: configure, pack: pack } := import('pack')

configure({
    entry: 'hello.oak'
    output: 'hello'
})

pack()
```

### Web Server with Assets

```oak
{ configure: configure, pack: pack } := import('pack')

configure({
    entry: 'server.oak'
    output: 'webserver'
    includes: ['lib/http.oak', 'lib/router.oak']
    includeVFS: [
        'static/css/'
        'static/js/'
        'templates/'
    ]
})

pack()
println('Packaged web server created!')
```

### CLI Tool with Config

```oak
{ configure: configure, pack: pack } := import('pack')
{ printf: printf } := import('fmt')

configure({
    entry: 'src/cli.oak'
    output: 'bin/mytool'
    includeVFS: 'config/'
    log: fn(msg) {
        printf('[PACK] {{0}}', msg)
    }
})

pack()
```

### Cross-Platform Build Script

```oak
{ configure: configure, pack: pack } := import('pack')
{ printf: printf } := import('fmt')

platforms := [
    { name: 'linux-x64', output: 'dist/app-linux' }
    { name: 'windows-x64', output: 'dist/app-windows.exe' }
    { name: 'darwin-x64', output: 'dist/app-macos' }
]

each(platforms, fn(platform) {
    printf('Building for {{0}}...', platform.name)
    
    configure({
        entry: 'src/main.oak'
        output: platform.output
        includeVFS: 'resources/'
    })
    
    pack()
    printf('✓ {{0}} complete', platform.name)
})
```

## Accessing Embedded Resources

Files included via `includeVFS` are accessible through the `Virtual` library:

```oak
// At build time
configure({
    includeVFS: 'templates/'
})

// At runtime in main.oak
Virtual := import('Virtual')

template := Virtual.readFile('templates/index.html')
if template != ? -> {
    println('Template loaded: ' + string(len(template)) + ' bytes')
}
```

## Pack Workflow

1. **Parse** entry point
2. **Resolve** dependencies (imports)
3. **Collect** VFS files from `includeVFS`
4. **Bundle** all Oak code into single AST
5. **Embed** runtime + code + VFS data
6. **Generate** executable binary
7. **Write** output file with execute permissions

## Platform-Specific Outputs

### Linux/macOS

```oak
configure({
    entry: 'app.oak'
    output: 'myapp'  // Creates executable 'myapp'
})
```

### Windows

```oak
configure({
    entry: 'app.oak'
    output: 'myapp.exe'  // Creates 'myapp.exe'
})
```

## Build vs Pack

**`build`**: Compiles to JavaScript/WebAssembly/Oak source
**`pack`**: Creates standalone native executable

```oak
// build - outputs JavaScript
build.configure({ entry: 'app.oak', output: 'app.js', web?: true })
build.compile()

// pack - outputs executable binary
pack.configure({ entry: 'app.oak', output: 'app' })
pack.pack()
```

## Implementation Notes

- Uses `build` library internally for bundling
- Embeds Oak runtime into output binary
- Sets executable permissions on Unix platforms
- Recursively processes directories in `includeVFS`
- Preserves directory structure in virtual filesystem
- Uses `pack-utils` for platform-specific operations

## Limitations

- Output file size includes runtime + all dependencies
- No compression of embedded resources
- Cannot update VFS files after packing
- Limited to including text files (binary files may have issues)
- No code signing or notarization
- No custom runtime flags

## Use Cases

- **CLI tools**: Self-contained command-line utilities
- **Web servers**: Embed templates and static assets
- **Scripts**: Distribute Oak scripts as executables
- **Games**: Bundle assets with game logic
- **Installers**: Create setup programs

## Performance Considerations

- Pack time scales with number of files in VFS
- Runtime VFS access is in-memory (fast)
- Larger VFS increases binary size
- First load extracts VFS to memory

## See Also

- [build.md](build.md) - Oak compiler and bundler
- [virtual-fs.md](virtual-fs.md) - Virtual filesystem documentation
- [pack-utils.md](pack-utils.md) - Pack utility functions
- `Virtual` library - Runtime VFS access
