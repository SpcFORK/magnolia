# Filesystem Library (fs)

## Overview

`libfs` provides ergonomic filesystem APIs for Oak programs, wrapping built-in filesystem functions for safer and more efficient file operations. All functions support both synchronous and asynchronous variants.

## Import

```oak
fs := import('fs')
// or destructure specific functions
{ readFile: readFile, writeFile: writeFile, listFiles: listFiles } := import('fs')
```

## Async vs Sync Usage

Most functions work in two modes:

- **Synchronous** (no callback): Blocks and returns the value immediately
- **Asynchronous** (with callback): Non-blocking, calls the callback with the result

```oak
// Synchronous
content := readFile('file.txt')

// Asynchronous
with readFile('file.txt') fn(content) {
    println(content)
}
```

## Configuration

### `ReadBufSize`

The buffer size for streaming file reads. Default: `4096` bytes. Can be modified globally:

```oak
fs := import('fs')
fs.ReadBufSize := 8192  // Increase buffer size
```

## Functions

### `readFile(path, withFile?)`

Reads the entire contents of a file and returns it as a string. Returns `?` on error.

**Parameters:**
- `path` - File path to read
- `withFile` - Optional callback for async mode

```oak
{ readFile: readFile } := import('fs')

// Sync
content := readFile('hello.txt')
if content != ? -> println(content)

// Async
with readFile('large-file.txt') fn(content) if content != ? -> {
    println('File size: ' + string(len(content)))
}
```

### `writeFile(path, file, withEnd?)`

Writes data to a file, creating it if it doesn't exist or truncating it if it does. Returns `true` on success, `?` on error.

**Parameters:**
- `path` - File path to write
- `file` - String content to write
- `withEnd` - Optional callback for async mode

```oak
{ writeFile: writeFile } := import('fs')

// Sync
success := writeFile('output.txt', 'Hello, World!')
if success -> println('File written successfully')

// Async
with writeFile('output.txt', 'Hello, Async World!') fn(success) if success -> {
    println('File written successfully')
}
```

### `appendFile(path, file, withEnd?)`

Appends data to the end of a file, creating it if it doesn't exist. Returns `true` on success, `?` on error.

**Parameters:**
- `path` - File path to append to
- `file` - String content to append
- `withEnd` - Optional callback for async mode

```oak
{ appendFile: appendFile } := import('fs')

// Sync
appendFile('log.txt', 'New log entry\n')

// Async
with appendFile('log.txt', 'Async log entry\n') fn(success) {
    println('Log appended: ' + string(success))
}
```

### `statFile(path, withStat?)`

Returns file metadata (from Oak's built-in `stat()`). Returns `?` on error.

**Parameters:**
- `path` - File path to stat
- `withStat` - Optional callback for async mode

```oak
{ statFile: statFile } := import('fs')

// Sync
info := statFile('file.txt')
if info != ? -> {
    println('Size: ' + string(info.size))
    println('Modified: ' + string(info.mod))
    println('Is directory: ' + string(info.dir))
}

// Async
with statFile('file.txt') fn(info) if info != ? -> {
    println('File info:', info)
}
```

### `listFiles(path, withFiles?)`

Returns a list of files and directories in a directory. Returns `?` on error or if not a directory.

**Parameters:**
- `path` - Directory path to list
- `withFiles` - Optional callback for async mode

```oak
{ listFiles: listFiles } := import('fs')

// Sync
files := listFiles('.')
if files != ? -> {
    each(files, fn(file) println(file))
}

// Async
with listFiles('.') fn(files) if files != ? -> {
    println('Found ' + string(len(files)) + ' files')
}
```

## Examples

### Reading Configuration Files

```oak
{ readFile: readFile } := import('fs')
json := import('json')

fn loadConfig(configPath) {
    content := readFile(configPath)
    if content != ? -> json.parse(content)
}

config := loadConfig('config.json')
if config != :error -> {
    println('Loaded config for: ' + config.appName)
}
```

### Writing Log Files

```oak
{ appendFile: appendFile } := import('fs')
{ format: format } := import('datetime')

fn log(message) {
    timestamp := format(time())
    entry := timestamp + ': ' + message + '\n'
    appendFile('app.log', entry)
}

log('Application started')
log('Processing request')
log('Application stopped')
```

### Recursive Directory Traversal

```oak
{ listFiles: listFiles, statFile: statFile } := import('fs')
{ join: join } := import('path')

fn walkDir(dir) {
    files := listFiles(dir)
    if files != ? -> {
        each(files, fn(filename) {
            fullPath := join(dir, filename)
            info := statFile(fullPath)
            
            if info != ? -> if info.dir {
                true -> {
                    println('Directory: ' + fullPath)
                    walkDir(fullPath)  // Recurse into subdirectory
                }
                _ -> println('File: ' + fullPath)
            }
        })
    }
}

walkDir('.')
```

### Async File Processing Pipeline

```oak
{ readFile: readFile, writeFile: writeFile } := import('fs')
{ upper: upper } := import('str')

// Async pipeline: read → transform → write
with readFile('input.txt') fn(content) if content != ? -> {
    transformed := upper(content)
    
    with writeFile('output.txt', transformed) fn(success) {
        println('Processing complete: ' + string(success))
    }
}
```

### Checking File Existence

```oak
{ statFile: statFile } := import('fs')

fn fileExists?(path) statFile(path) != ?

if fileExists?('config.json') {
    true -> println('Config file found')
    _ -> println('Config file missing')
}
```

### Batch File Operations

```oak
{ readFile: readFile, writeFile: writeFile } := import('fs')

fn copyFile(src, dest) {
    content := readFile(src)
    if content != ? -> writeFile(dest, content)
}

fn backupFiles(files) {
    each(files, fn(file) {
        if copyFile(file, file + '.bak') -> {
            println('Backed up: ' + file)
        }
    })
}

backupFiles(['data.json', 'config.json', 'state.json'])
```

## Error Handling

All functions return `?` on error. Always check return values:

```oak
{ readFile: readFile, writeFile: writeFile } := import('fs')

content := readFile('nonexistent.txt')
if content {
    ? -> println('Error: Could not read file')
    _ -> {
        // Process content
        result := processData(content)
        
        success := writeFile('output.txt', result)
        if !success -> println('Error: Could not write file')
    }
}
```

## Performance Notes

- **Synchronous operations** block the event loop—use async for I/O-heavy tasks
- **Async operations** allow concurrent file operations and better responsiveness
- **ReadBufSize** affects memory usage vs. read performance trade-off
- Large files are read in chunks of `ReadBufSize` bytes
- `writeFile()` truncates existing files—use `appendFile()` to preserve content

## Limitations

- No streaming API—entire file must fit in memory
- No directory creation (`mkdir`) function—use built-in `make()` syscall
- No file deletion—use built-in `remove()` syscall
- No file copying—must read and write manually
- No file permissions/mode support beyond basic stat info
- Character encoding is assumed to be UTF-8

## See Also

- `path` library - For path manipulation
- Oak built-in functions: `open()`, `close()`, `read()`, `write()`, `stat()`, `ls()`
