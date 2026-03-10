# Virtual File System (VFS) for Oak

## Overview

The Virtual File System (VFS) provides an in-memory file system that can be embedded in packed Oak binaries, enabling cross-platform file access without relying on the underlying operating system's file system.

## Features

- **In-Memory Storage**: Files are stored in memory as a simple key-value object
- **Cross-Platform**: Works on all platforms (Windows, Linux, macOS, Web/WASM)
- **Path Normalization**: Automatically normalizes paths (converts backslashes to forward slashes)
- **Embeddable**: Can be serialized as JSON and embedded in packed binaries
- **Drop-in Replacement**: Provides similar API to the standard `fs` module

## API

### Creating a Virtual FS

```oak
Virtual := import('Virtual')

// Create empty VFS
vfs := Virtual.createVirtualFS({})

// Create VFS with initial files
vfs := Virtual.createVirtualFS({
    'config.json': '{"version": "1.0"}'
    'data/test.txt': 'test data'
})
```

### VFS Methods

#### `readFile(path, withFile?)`

Read file contents from the VFS.

```oak
// Synchronous
content := vfs.readFile('config.json')

// Asynchronous
vfs.readFile('config.json') fn(content) {
    // handle content
}
```

#### `writeFile(path, content, withResult?)`

Write file contents to the VFS.

```oak
// Synchronous
vfs.writeFile('output.txt', 'Hello World')

// Asynchronous
vfs.writeFile('output.txt', 'Hello World') fn(result) {
    // handle result
}
```

#### `exists?(path)`

Check if a file exists in the VFS.

```oak
if vfs.exists?('config.json') {
    true -> // file exists
}
```

#### `statFile(path, withStat?)`

Get file statistics.

```oak
// Synchronous
stat := vfs.statFile('config.json')
// Returns { name: path, size: length, dir?: false, mod: 0 }

// Asynchronous
vfs.statFile('config.json') fn(stat) {
    // handle stat
}
```

#### `listFiles(withList?)`

List all files in the VFS.

```oak
// Synchronous
files := vfs.listFiles()

// Asynchronous
vfs.listFiles() fn(files) {
    // handle files
}
```

#### `deleteFile(path, withResult?)`

Delete a file from the VFS.

```oak
// Synchronous
vfs.deleteFile('temp.txt')

// Asynchronous
vfs.deleteFile('temp.txt') fn(result) {
    // handle result
}
```

#### `getFiles()`

Get all files as an object (for serialization).

```oak
allFiles := vfs.getFiles()
```

## File Bundling

The VFS supports embedding files directly into Oak bundles and packed executables using the `--includeVFS` flag.

### Using `--includeVFS` with `oak build`

Bundle files into an Oak script or JavaScript bundle:

```bash
# Bundle a single file
oak build --entry app.oak --output bundle.oak --includeVFS data.txt

# Bundle multiple files (comma-separated)
oak build --entry app.oak --output bundle.oak --includeVFS config.json,data.txt

# Bundle with custom target names (using target:source syntax)
oak build --entry app.oak --output bundle.oak --includeVFS config:config.json,readme:README.md

# Bundle entire directories (recursively includes all files)
oak build --entry app.oak --output bundle.oak --includeVFS data:./data-dir
```

When files are bundled with `oak build`, they are available via the global `__Oak_VFS` variable:

```oak
// Access bundled files
if __Oak_VFS? {
    true -> {
        content := __Oak_VFS.readFile('data.txt')
        // Use the content...
    }
    _ -> println('No VFS data bundled')
}
```

### Using `--includeVFS` with `oak pack`

Embed files into standalone executables:

```bash
# Pack with embedded files
oak pack --entry app.oak --output app --includeVFS config.json,assets:./assets

# The packed executable will contain the embedded files
./app
```

Files are embedded in the executable and automatically available at runtime.

### Spec Format

The `--includeVFS` flag accepts specifications in the following formats:

- **Simple file**: `file.txt` - bundles `file.txt` as `file.txt` in the VFS
- **Target:source**: `config:myconfig.json` - bundles `myconfig.json` as `config` in the VFS
- **Directory**: `assets:./assets-dir` - recursively bundles all files from `assets-dir` directory with the prefix `assets/`
- **Multiple specs**: Comma-separated list like `file1.txt,config:file2.json`

### Example

Create a simple app that reads bundled configuration:

```oak
// app.oak
{
    println: println
} := import('std')

if __Oak_VFS? {
    true -> {
        config := __Oak_VFS.readFile('config.json')
        println('Config: ' << config)
    }
}
```

Bundle it with a config file:

```bash
oak build --entry app.oak --output app.bundle.oak --includeVFS config.json
oak app.bundle.oak
```

Pack it as an executable:

```bash
oak pack --entry app.oak --output myapp --includeVFS config.json
./myapp
```
```

#### `setFiles(newFiles)`

Replace all files in the VFS.

```oak
vfs.setFiles({
    'new.txt': 'new content'
})
```

## Pack Integration

The VFS is integrated with the `oak pack` command to embed files in standalone binaries.

### Embedding Files

Use the `--include` flag to embed files in the packed binary:

```bash
# Embed a single file
oak pack --entry main.oak --output app --include config.json

# Embed a directory
oak pack --entry main.oak --output app --include static:./static

# Embed multiple files/directories
oak pack --entry main.oak --output app --include "config.json,data:./data,lib:./lib"
```

### Include Syntax

The `--include` flag accepts comma-separated specifications in these formats:

- `path` - Embed file/directory at the same path in VFS
- `name:path` - Embed file/directory at a different name in VFS

Examples:

- `config.json` → VFS path: `config.json`
- `data:./data` → VFS paths: `data/file1.txt`, `data/file2.txt`, etc.
- `lib/helper.oak:./src/helper.oak` → VFS path: `lib/helper.oak`

### Accessing VFS at Runtime

When a packed binary runs, it can use the VFS through a runtime-provided module:

```oak
// Access the embedded VFS (if available)
vfs := ___packed_vfs()

if vfs {
    ? -> // No VFS embedded
    _ -> {
        // Use VFS
        config := vfs.readFile('config.json')
        // ... use config
    }
}
```

## Implementation Details

### Storage Format

In packed binaries, VFS data is stored as JSON at the end of the file:

```
[executable][oak bundle][bundle size][magic bytes][vfs json][vfs size][vfs magic]
```

- **VFS JSON**: Serialized object mapping paths to file contents
- **VFS Size**: 24-byte padded size of the VFS JSON
- **VFS Magic**: 8 bytes: `"vfs \x01\x00\x00\x00"`

### Path Normalization

All paths are normalized to use forward slashes (`/`) and avoid double slashes:

```oak
vfs.writeFile('path\\to\\file.txt', 'data')
content := vfs.readFile('path/to/file.txt')  // Works!
```

### Serialization

The VFS uses JSON for serialization to ensure cross-platform compatibility:

```oak
packUtils := import('pack-utils')

// Serialize
jsonString := packUtils.serializeVFS(vfsFiles)

// Deserialize
vfsFiles := packUtils.deserializeVFS(jsonString)
```

## Use Cases

1. **Embedded Configuration**: Ship configuration files with the binary
2. **Static Assets**: Embed templates, images, or other assets
3. **Data Files**: Include reference data or lookup tables
4. **Cross-Platform Deployment**: Ensure files are available regardless of platform
5. **Single-File Distribution**: Package entire applications as a single executable

## Limitations

- Files are loaded entirely into memory
- No streaming for large files
- No directory operations (directories are implicit from paths)
- Read-write in memory only (changes not persisted back to packed binary)
- File metadata is minimal (no permissions, timestamps, etc.)

## Testing

Run VFS tests:

```bash
# Run all tests including VFS
make test-oak

# Run specific VFS tests
oak test/main.oak virtual pack
```

## Example: Packed Web Server

```oak
// main.oak - Simple web server with embedded files
http := import('http')
vfs := ___packed_vfs()

// Serve files from embedded VFS
http.serve('0.0.0.0:8080', fn(req) {
    path := if req.path = '/' {
        true -> 'index.html'
        _ -> req.path |> slice(1)  // remove leading /
    }
    
    if content := vfs.readFile('static/' + path) {
        ? -> { status: 404, body: 'Not Found' }
        _ -> { status: 200, body: content }
    }
})
```

Pack with embedded files:

```bash
oak pack --entry main.oak --output server --include static:./public
./server  # Run standalone server with embedded files
```

## Future Enhancements

Potential future improvements:

- Compression of embedded files
- Encryption of sensitive embedded data
- Lazy loading of large embedded files
- Directory traversal APIs
- Pattern matching for file selection
- Virtual file mounting/overlays
- Write-back to real file system option
