# Pack Utils Library (pack-utils)

## Overview

`libpack-utils` provides utilities for packing Oak programs into standalone executable binaries with embedded bundles and virtual filesystems.

## Import

```oak
packUtils := import('pack-utils')
{
    packBundle: packBundle
    packBundleWithVFS: packBundleWithVFS
    serializeVFS: serializeVFS
    deserializeVFS: deserializeVFS
} := import('pack-utils')
```

## Constants

### Magic Bytes

```oak
MagicBytes := 'oak \x19\x98\x10\x15'       // 8 bytes: identifies Oak bundles
VFSMagicBytes := 'vfs \x01\x00\x00\x00'    // 8 bytes: identifies VFS data
```

### Size Encoding

```oak
MaxBundleSizeLen := 24  // Bytes for bundle size (fits UINT64_MAX)
MaxVFSSizeLen := 24     // Bytes for VFS size
```

## Functions

### `packBundle(executable, oakBundle)`

Combines an executable with an Oak bundle into a packed binary format.

**Parameters:**
- `executable` - Binary executable data
- `oakBundle` - Oak source/bytecode bundle

**Returns:** Complete packed binary

**Format:** `[executable][oakBundle][bundleSize(24)][magic(8)]`

```oak
{ packBundle: packBundle } := import('pack-utils')

executable := readFile('oak-runtime')
oakBundle := compileBundle('main.oak')

packed := packBundle(executable, oakBundle)
writeFile('myapp', packed)
```

### `packBundleWithVFS(executable, oakBundle, vfsFiles)`

Packs executable, Oak bundle, and virtual filesystem into one binary.

**Parameters:**
- `executable` - Binary executable data  
- `oakBundle` - Oak source/bytecode bundle
- `vfsFiles` - Object mapping virtual paths to file contents

**Returns:** Complete packed binary with VFS

**Format:** `[executable][oakBundle][bundleSize(24)][magic(8)][vfsData][vfsSize(24)][vfsMagic(8)]`

```oak
{ packBundleWithVFS: packBundleWithVFS } := import('pack-utils')

executable := readFile('oak-runtime')
oakBundle := compileBundle('main.oak')
vfsFiles := {
    'templates/index.html': readFile('templates/index.html')
    'static/style.css': readFile('static/style.css')
}

packed := packBundleWithVFS(executable, oakBundle, vfsFiles)
writeFile('myapp', packed)
```

### `serializeVFS(vfsFiles)`

Converts VFS files object to JSON string.

**Parameters:**
- `vfsFiles` - Object mapping paths to contents

**Returns:** JSON string

```oak
{ serializeVFS: serializeVFS } := import('pack-utils')

vfs := {
    'config.json': '{"port": 8080}'
    'template.html': '<html>...</html>'
}

serialized := serializeVFS(vfs)
// => '{"config.json":"{\\"port\\": 8080}","template.html":"<html>...</html>"}'
```

### `deserializeVFS(vfsData)`

Converts JSON string back to VFS files object.

**Parameters:**
- `vfsData` - JSON string

**Returns:** VFS files object

```oak
{ deserializeVFS: deserializeVFS } := import('pack-utils')

json := '{"config.json":"{\\"port\\": 8080}"}'
vfs := deserializeVFS(json)
// => { 'config.json': '{"port": 8080}' }

// Handles edge cases
deserializeVFS('') // => {}
deserializeVFS(?) // => {}
```

## Binary Format Structures

### Basic Pack Format

```
┌─────────────────┬──────────────┬──────────────┬──────────┐
│   Executable    │  Oak Bundle  │ Bundle Size  │  Magic   │
│   (variable)    │  (variable)  │   (24 bytes) │ (8 bytes)│
└─────────────────┴──────────────┴──────────────┴──────────┘
```

### Pack with VFS Format

```
┌──────────┬────────┬────────┬───────┬─────────┬──────────┬──────────┐
│Executable│ Bundle │ Size   │ Magic │VFS Data │VFS Size  │VFS Magic │
│(variable)│(var)   │(24b)   │(8b)   │(var)    │(24b)     │(8b)      │
└──────────┴────────┴────────┴───────┴─────────┴──────────┴──────────┘
```

## Usage Examples

### Simple Packing

```oak
{ packBundle: packBundle } := import('pack-utils')

// Read runtime executable
runtime := readFile('build/oak')

// Compile Oak code to bundle
bundle := buildOakBundle('src/main.oak')

// Pack together
packed := packBundle(runtime, bundle)

// Write executable
writeFile('dist/myapp', packed)
```

### Packing with Assets

```oak
{ packBundleWithVFS: packBundleWithVFS } := import('pack-utils')

runtime := readFile('oak-runtime')
bundle := buildOakBundle('app.oak')

// Collect assets
vfs := {
    'assets/logo.png': readFile('assets/logo.png')
    'assets/style.css': readFile('assets/style.css')
    'templates/index.html': readFile('templates/index.html')
}

packed := packBundleWithVFS(runtime, bundle, vfs)
writeFile('myapp', packed)
```

### VFS Serialization

```oak
{ serializeVFS: serializeVFS, deserializeVFS: deserializeVFS } := import('pack-utils')

// Serialize VFS
vfs := {
    'data.txt': 'Hello, World!'
    'config.json': '{"debug": true}'
}

serialized := serializeVFS(vfs)
println('Serialized size: ' + string(len(serialized)))

// Deserialize later
restored := deserializeVFS(serialized)
println(restored.('data.txt')) // => 'Hello, World!'
```

### Size Encoding

```oak
{ encodeBundleSize: encodeBundleSize } := import('pack-utils')

// Bundle sizes are encoded as padded strings
encoded := encodeBundleSize(12345)
println(encoded) // => '                   12345' (24 chars)

encoded := encodeBundleSize(9876543210)
println(encoded) // => '           9876543210' (24 chars)
```

### Cross-Platform Packaging

```oak
{ packBundleWithVFS: packBundleWithVFS } := import('pack-utils')

fn packForPlatform(platform) {
    runtime := readFile('runtimes/' + platform + '/oak')
    bundle := buildOakBundle('app.oak')
    vfs := collectAssets()
    
    packed := packBundleWithVFS(runtime, bundle, vfs)
    
    outputName := if platform {
        'windows' -> 'app.exe'
        _ -> 'app'
    }
    
    writeFile('dist/' + platform + '/' + outputName, packed)
}

each(['windows', 'linux', 'darwin'], packForPlatform)
```

## Magic Bytes Details

### Oak Bundle Magic

```
'oak ' (ASCII) + hex bytes 19 98 10 15
Represents: "oak" + date 1998-10-15 (Oak's birthdate?)
Total: 8 bytes
```

### VFS Magic

```
'vfs ' (ASCII) + hex bytes 01 00 00 00
Represents: "vfs" + version 1.0.0.0
Total: 8 bytes
```

## Unpacking Process

At runtime, the Oak executable:

1. **Reads magic bytes** from end of file
2. **Verifies** magic bytes match
3. **Reads size** (24 bytes before magic)
4. **Extracts bundle** using size
5. **Checks for VFS magic** (if present)
6. **Extracts VFS data** if VFS magic found
7. **Deserializes VFS** to object
8. **Loads bundle** and executes

## Implementation Notes

- Size fields are right-padded with spaces
- VFS data is JSON-encoded
- Binary format is platform-independent
- File sizes stored as decimal strings
- 24-byte size field supports bundles up to ~1 yottabyte

## Limitations

- No compression of bundle or VFS
- No encryption
- No version checking
- VFS stored as JSON (not binary efficient)
- Size fields are ASCII (not binary integers)

## Use Cases

- **Application distribution**: Single-file apps
- **Embedded assets**: Bundle resources with code
- **Portable tools**: Self-contained utilities
- **Deployment**: Simplified distribution

## Security Considerations

- No signature validation
- No tamper detection
- Executable permissions must be set separately
- VFS data is plaintext (JSON)

## See Also

- [pack.md](pack.md) - Pack library
- [build.md](build.md) - Build system
- [virtual-fs.md](virtual-fs.md) - Virtual filesystem
- `json` library - JSON encoding/decoding
