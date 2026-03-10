# Bundle Utils Library (bundle-utils)

## Overview

`libbundle-utils` provides utilities for module path normalization and common prefix calculation used during the bundling process.

## Import

```oak
bundleUtils := import('bundle-utils')
{ commonPrefix: commonPrefix, normalizeModulePath: normalizeModulePath } := import('bundle-utils')
```

## Functions

### `commonPrefix(paths)`

Finds the longest common prefix path among a list of file paths.

**Parameters:**
- `paths` - List of file path strings

**Returns:** Common prefix string, or `''` if fewer than 2 paths

```oak
{ commonPrefix: commonPrefix } := import('bundle-utils')

paths := [
    '/home/user/project/src/main.oak'
    '/home/user/project/src/utils.oak'
    '/home/user/project/lib/helpers.oak'
]

prefix := commonPrefix(paths)
// => '/home/user/project/'
```

### `normalizeModulePath(path, allPaths)`

Converts an absolute file path to a module identifier by removing the common prefix.

**Parameters:**
- `path` - Absolute file path to normalize
- `allPaths` - List of all module paths (used to find common prefix)

**Returns:** Normalized module path

```oak
{ normalizeModulePath: normalizeModulePath } := import('bundle-utils')

allPaths := [
    '/project/src/main.oak'
    '/project/src/utils.oak'
    '/project/lib/helpers.oak'
]

normalized := normalizeModulePath('/project/src/main.oak', allPaths)
// => 'src/main.oak'

normalized := normalizeModulePath('/project/lib/helpers.oak', allPaths)
// => 'lib/helpers.oak'
```

## Usage Examples

### Module Path Normalization

```oak
{ commonPrefix: commonPrefix, normalizeModulePath: normalizeModulePath } := import('bundle-utils')

modulePaths := [
    '/Users/dev/myapp/src/index.oak'
    '/Users/dev/myapp/src/router.oak'
    '/Users/dev/myapp/lib/utils.oak'
]

// Find common prefix
prefix := commonPrefix(modulePaths)
println('Common prefix: ' + prefix)
// => 'Common prefix: /Users/dev/myapp/'

// Normalize each path
normalizedPaths := modulePaths |> map(fn(path) {
    normalizeModulePath(path, modulePaths)
})

each(normalizedPaths, println)
// => 'src/index.oak'
// => 'src/router.oak'
// => 'lib/utils.oak'
```

### Build System Integration

```oak
{ normalizeModulePath: normalizeModulePath } := import('bundle-utils')

fn bundleModules(entryPoint) {
    // Collect all module paths
    modulePaths := resolveAllImports(entryPoint)
    
    // Normalize for clean module names
    modules := modulePaths |> map(fn(path) {
        code := readFile(path)
        ast := parseCode(code)
        moduleName := normalizeModulePath(path, modulePaths)
        
        [moduleName, ast]
    })
    
    wrapBundle(modules, normalizeModulePath(entryPoint, modulePaths))
}
```

### Cross-Platform Path Handling

```oak
{ normalizeModulePath: normalizeModulePath } := import('bundle-utils')

// Works with both forward and backslashes
windowsPaths := [
    'C:\\project\\src\\main.oak'
    'C:\\project\\src\\utils.oak'
]

unixPaths := [
    '/home/user/project/src/main.oak'
    '/home/user/project/src/utils.oak'
]

normalizeModulePath(windowsPaths.0, windowsPaths)
// => 'src\\main.oak' or 'src/main.oak'

normalizeModulePath(unixPaths.0, unixPaths)
// => 'src/main.oak'
```

### Single Module Edge Case

```oak
{ commonPrefix: commonPrefix, normalizeModulePath: normalizeModulePath } := import('bundle-utils')

// With single path, prefix is empty (avoids stripping the entire path)
singlePath := ['/project/main.oak']

prefix := commonPrefix(singlePath)
// => ''

normalized := normalizeModulePath(singlePath.0, singlePath)
// => '/project/main.oak' (unchanged)
```

## Implementation Details

### Prefix Calculation Algorithm

1. Compare paths character by character
2. Stop at first mismatch
3. Return longest common substring
4. Empty string if < 2 paths

### Path Normalization

1. Calculate common prefix from all paths
2. Remove prefix from target path
3. Strip leading slash/backslash if present
4. Return relative path

### Edge Cases

```oak
// Empty prefix
commonPrefix([]) // => ''
commonPrefix(['/single/path.oak']) // => ''

// No common prefix
commonPrefix(['/a/b.oak', '/x/y.oak']) // => ''

// Full path match
commonPrefix(['/same/path.oak', '/same/path.oak']) // => '/same/path.oak'

// Trailing slashes handled
normalizeModulePath('/proj/src/main.oak', ['/proj/'])
// => 'src/main.oak'
```

## Use Cases

- **Module bundling**: Clean module identifiers
- **Import resolution**: Relative path calculation
- **Build systems**: Normalize output paths
- **Cross-platform**: Handle path separators

## Limitations

- Compares paths lexically (doesn't resolve symlinks)
- Doesn't canonicalize paths (../ not resolved)
- Path separator handling is basic
- No filesystem access (pure string manipulation)

## See Also

- [bundle-ast.md](bundle-ast.md) - AST bundling
- [build.md](build.md) - Build system
- `path` library - Path manipulation utilities
