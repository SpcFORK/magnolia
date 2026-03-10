# Path Library (path)

## Overview

`libpath` implements utilities for working with UNIX-style file system paths and URIs. It provides path manipulation, normalization, and resolution functions.

## Import

```oak
path := import('path')
// or destructure specific functions
{ join: join, clean: clean, dir: dir, base: base } := import('path')
```

## Path Classification

### `abs?(path)`

Returns true if the path is absolute (starts with `/`).

```oak
{ abs?: abs? } := import('path')

abs?('/usr/local/bin') // => true
abs?('/home/user') // => true
abs?('relative/path') // => false
abs?('./current') // => false
```

### `rel?(path)`

Returns true if the path is relative (does not start with `/`).

```oak
{ rel?: rel? } := import('path')

rel?('relative/path') // => true
rel?('./current') // => true
rel?('../parent') // => true
rel?('/absolute/path') // => false
```

## Path Components

### `dir(path)`

Returns the directory portion of a path (everything except the last component).

```oak
{ dir: dir } := import('path')

dir('/usr/local/bin') // => '/usr/local'
dir('/home/user/file.txt') // => '/home/user'
dir('path/to/file.txt') // => 'path/to'
dir('/file.txt') // => ''
dir('file.txt') // => ''

// Trailing slashes are removed before processing
dir('/usr/local/') // => '/usr'
```

### `base(path)`

Returns the last component of a path (typically the filename).

```oak
{ base: base } := import('path')

base('/usr/local/bin') // => 'bin'
base('/home/user/file.txt') // => 'file.txt'
base('path/to/file.txt') // => 'file.txt'
base('/file.txt') // => 'file.txt'
base('file.txt') // => 'file.txt'

// Trailing slashes are removed
base('/usr/local/') // => 'local'
```

### `cut(path)`

Splits a path into `[directory, basename]` pair.

```oak
{ cut: cut } := import('path')

cut('/usr/local/bin')
// => ['/usr/local', 'bin']

cut('/home/user/file.txt')
// => ['/home/user', 'file.txt']

cut('path/to/file.txt')
// => ['path/to', 'file.txt']

cut('file.txt')
// => ['', 'file.txt']
```

## Path Manipulation

### `clean(path)`

Normalizes a path by applying the following transformations:
1. Remove consecutive slashes (except at the beginning)
2. Remove `.` (current directory)
3. Remove `..` and the preceding directory component (if possible)

```oak
{ clean: clean } := import('path')

clean('/usr//local///bin') // => '/usr/local/bin'
clean('/usr/./local/./bin') // => '/usr/local/bin'
clean('/usr/local/../bin') // => '/usr/bin'
clean('./file.txt') // => 'file.txt'
clean('path/./to/../file') // => 'path/file'

// Multiple parent references
clean('/a/b/c/../../d') // => '/a/d'

// Parent refs that can't be resolved stay
clean('../parent') // => '../parent'
clean('../../up') // => '../../up'

// Empty path
clean('') // => ''
```

### `join(parts...)`

Joins multiple path components into a single normalized path.

```oak
{ join: join } := import('path')

join('/usr', 'local', 'bin') // => '/usr/local/bin'
join('/home', 'user', 'documents') // => '/home/user/documents'
join('path', 'to', 'file.txt') // => 'path/to/file.txt'

// Automatically cleans the result
join('/usr', './local', '../lib') // => '/usr/lib'
join('a', 'b/../c') // => 'a/c'

// Works with absolute components
join('/base', '/absolute') // => '/base/absolute'

// Handles empty strings
join('path', '', 'file') // => 'path/file'
join('', 'path') // => 'path'
```

### `split(path)`

Splits a path into its individual components, ignoring trailing slashes. Returns empty list for empty paths.

```oak
{ split: split } := import('path')

split('/usr/local/bin') // => ['usr', 'local', 'bin']
split('path/to/file.txt') // => ['path', 'to', 'file.txt']
split('/home/user/') // => ['home', 'user']

// Empty strings are filtered out
split('path//to/file') // => ['path', 'to', 'file']

// Empty or root path
split('') // => []
split('/') // => []

// Relative paths
split('./path/to/file') // => ['.', 'path', 'to', 'file']
```

### `resolve(path, base?)`

Returns an absolute path by resolving a relative path against a base directory. Uses current working directory if no base given.

```oak
{ resolve: resolve } := import('path')

// Resolve against current directory (env().PWD)
resolve('file.txt')
// => '/current/working/directory/file.txt'

resolve('./path/to/file')
// => '/current/working/directory/path/to/file'

// Resolve against specific base
resolve('file.txt', '/home/user')
// => '/home/user/file.txt'

resolve('../parent/file', '/home/user/documents')
// => '/home/user/parent/file'

// Absolute paths are cleaned and returned as-is
resolve('/absolute/path', '/base')
// => '/absolute/path'

resolve('/usr/../lib', '/base')
// => '/lib'
```

## Examples

### Building File Paths

```oak
{ join: join } := import('path')

homeDir := '/home/user'
documentsPath := join(homeDir, 'Documents', 'report.txt')
// => '/home/user/Documents/report.txt'

projectPath := join('/projects', 'myapp', 'src', 'main.oak')
// => '/projects/myapp/src/main.oak'
```

### Extracting File Information

```oak
{ dir: dir, base: base } := import('path')

filePath := '/home/user/documents/report.pdf'

directory := dir(filePath)
// => '/home/user/documents'

filename := base(filePath)
// => 'report.pdf'

println('File: ' + filename)
println('Location: ' + directory)
```

### Path Navigation

```oak
{ join: join, dir: dir } := import('path')

currentFile := '/projects/myapp/src/components/Header.oak'

// Go to parent directory
parentDir := dir(currentFile)
// => '/projects/myapp/src/components'

// Go to sibling directory
siblingPath := join(dir(currentFile), '../styles/main.css')
// Will need to be cleaned

// Go to project root
projectRoot := join(currentFile, '../../../..')
```

### URL/URI Path Manipulation

```oak
{ join: join, clean: clean } := import('path')

baseURL := '/api/v1'
resource := 'users'
id := '123'

endpoint := join(baseURL, resource, id)
// => '/api/v1/users/123'

// Clean messy paths
cleaned := clean('/api//v1/./users/../posts//' + string(id))
// => '/api/v1/posts/123'
```

### Resolving Relative Imports

```oak
{ resolve: resolve, join: join, dir: dir } := import('path')

fn resolveImport(currentFile, importPath) {
    if abs?(importPath) {
        true -> importPath
        _ -> {
            currentDir := dir(currentFile)
            resolve(importPath, currentDir)
        }
    }
}

currentFile := '/project/src/main.oak'
import1 := resolveImport(currentFile, './utils/helper.oak')
// => '/project/src/utils/helper.oak'

import2 := resolveImport(currentFile, '../lib/core.oak')
// => '/project/lib/core.oak'
```

### Path Validation

```oak
{ clean: clean } := import('path')

fn isValidPath?(path) {
    cleaned := clean(path)
    // Check if path tries to escape root
    !startsWith?(cleaned, '../')
}

isValidPath?('safe/path/file.txt') // => true
isValidPath?('../../../etc/passwd') // => false
isValidPath?('./local/file') // => true
```

### Building Paths Safely

```oak
{ join: join, clean: clean } := import('path')

fn buildSafePath(root, userInput) {
    // Ensure user input doesn't escape root
    fullPath := join(root, userInput)
    cleaned := clean(fullPath)
    
    // Verify the result starts with root
    if startsWith?(cleaned, root) {
        true -> cleaned
        _ -> ? // Reject paths that escape root
    }
}

safePath := buildSafePath('/var/www', 'uploads/image.png')
// => '/var/www/uploads/image.png'

unsafe := buildSafePath('/var/www', '../../../etc/passwd')
// => ? (escaped root directory)
```

### File Extension Extraction

```oak
{ base: base } := import('path')
{ rindexOf: rindexOf, slice: slice } := import('str')

fn getExtension(path) {
    filename := base(path)
    dotIndex := rindexOf(filename, '.')
    
    if dotIndex > 0 {
        true -> slice(filename, dotIndex + 1)
        _ -> '' // No extension
    }
}

getExtension('/path/to/file.txt') // => 'txt'
getExtension('/path/to/archive.tar.gz') // => 'gz'
getExtension('/path/to/README') // => ''
```

### Path Comparison

```oak
{ clean: clean } := import('path')

fn pathsEqual?(path1, path2) {
    clean(path1) = clean(path2)
}

pathsEqual?('/usr/local', '/usr/./local') // => true
pathsEqual?('/usr/local/bin', '/usr/lib/../local/bin') // => true
pathsEqual?('/path/to/file', '/other/path') // => false
```

## Implementation Notes

- All functions work with UNIX-style paths (forward slashes)
- Trailing slashes are removed before processing in `dir()` and `base()`
- Consecutive slashes are preserved at the beginning for absolute paths
- `resolve()` uses `env().PWD` for current working directory
- No validation that paths actually exist on the file system
- No Windows path support (backslashes are not treated specially)
- Empty path components are filtered out in `split()` and `join()`

## Edge Cases

```oak
{ dir: dir, base: base, clean: clean, join: join } := import('path')

// Root directory
dir('/') // => ''
base('/') // => ''

// Single component
dir('file.txt') // => ''
base('file.txt') // => 'file.txt'

// Dotfiles
base('/path/to/.hidden') // => '.hidden'
dir('/path/to/.hidden') // => '/path/to'

// Multiple slashes
clean('///path///to///file') // => '/path/to/file'

// Only dots
clean('.') // => ''
clean('./././.') // => ''

// Parent references at root
clean('/../../..') // => '/' // Can't escape root
clean('a/../../b') // => '../b' // Relative escape
```

## Limitations

- No glob pattern matching
- No ~ (tilde) expansion for home directories
- No environment variable expansion
- No symbolic link resolution
- No path existence checking
- No Windows path support
- No URI scheme handling (http://, file://, etc.)
- Does not handle query strings or fragments in URIs

## See Also

- `fs` library - For file system operations
- `str` library - For string manipulation
- Oak built-in `env()` - For current working directory
