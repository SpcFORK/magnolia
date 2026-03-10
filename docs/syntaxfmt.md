# Syntax Format Library (syntaxfmt)

## Overview

`libsyntaxfmt` provides middleware for formatting Oak source files, isolating core formatting logic from CLI concerns. It formats Oak code to consistent style.

## Import

```oak
syntaxfmt := import('syntaxfmt')
{
    formatContent: formatContent
    formatFile: formatFile
    formatFileInPlace: formatFileInPlace
    formatFileWithDiff: formatFileWithDiff
    getChangedOakFiles: getChangedOakFiles
} := import('syntaxfmt')
```

## Functions

### `formatContent(content)`

Formats Oak source code string.

**Parameters:**
- `content` - Oak source code string

**Returns:** Formatted source code string

```oak
{ formatContent: formatContent } := import('syntaxfmt')

code := 'fn add(a,b){a+b}'
formatted := formatContent(code)
println(formatted)
// => fn add(a, b) {
//        a + b
//    }
```

### `formatFile(path)`

Reads and formats a file without modifying it.

**Parameters:**
- `path` - File path to format

**Returns:** Result object:
- `{ ok: true, content: <formatted> }` on success
- `{ ok: false, error: <message> }` on error

```oak
{ formatFile: formatFile } := import('syntaxfmt')

result := formatFile('lib/utils.oak')
if result.ok {
    true -> println(result.content)
    _ -> println('Error: ' + result.error)
}
```

### `formatFileInPlace(path)`

Formats a file and writes the result back to the same file.

**Parameters:**
- `path` - File path to format in-place

**Returns:** Result object:
- `{ ok: true }` on success
- `{ ok: false, error: <message> }` on error

```oak
{ formatFileInPlace: formatFileInPlace } := import('syntaxfmt')

result := formatFileInPlace('src/main.oak')
if result.ok {
    true -> println('File formatted successfully')
    _ -> println('Error: ' + result.error)
}
```

### `formatFileWithDiff(path)`

Formats a file and returns a diff showing changes.

**Parameters:**
- `path` - File path to format

**Returns:** Result object:
- `{ ok: true, diff: <diff output> }` on success
- `{ ok: false, error: <message> }` on error

**Note:** Requires `diff` command available on the system.

```oak
{ formatFileWithDiff: formatFileWithDiff } := import('syntaxfmt')

result := formatFileWithDiff('lib/utils.oak')
if result.ok {
    true -> if result.diff != '' {
        true -> println(result.diff)
        _ -> println('No formatting changes needed')
    }
    _ -> println('Error: ' + result.error)
}
```

### `getChangedOakFiles()`

Gets a list of changed `.oak` files from git.

**Returns:** Result object:
- `{ ok: true, files: [path1, path2, ...] }` on success
- `{ ok: false, error: <message> }` on error

**Note:** Requires git repository and `git` command.

```oak
{ getChangedOakFiles: getChangedOakFiles } := import('syntaxfmt')

result := getChangedOakFiles()
if result.ok {
    true -> {
        println('Changed Oak files:')
        each(result.files, println)
    }
    _ -> println('Error: ' + result.error)
}
```

## Usage Examples

### Format Single File

```oak
{ formatFileInPlace: formatFileInPlace } := import('syntaxfmt')
{ printf: printf } := import('fmt')

path := 'src/app.oak'
result := formatFileInPlace(path)

if result.ok {
    true -> printf('✓ Formatted {{0}}\n', path)
    _ -> printf('✗ Failed to format {{0}}: {{1}}\n', path, result.error)
}
```

### Format Multiple Files

```oak
{ formatFileInPlace: formatFileInPlace } := import('syntaxfmt')

files := ['lib/utils.oak', 'lib/helpers.oak', 'src/main.oak']

each(files, fn(file) {
    result := formatFileInPlace(file)
    if result.ok {
        true -> println('Formatted: ' + file)
        _ -> println('Failed: ' + file + ' - ' + result.error)
    }
})
```

### Check Formatting (Dry Run)

```oak
{ formatFile: formatFile } := import('syntaxfmt')

fn checkFormatting(path) {
    original := readFile(path)
    result := formatFile(path)
    
    if result.ok {
        true -> if result.content = original {
            true -> { formatted: true, path: path }
            _ -> { formatted: false, path: path }
        }
        _ -> { error: result.error, path: path }
    }
}

files := ['src/main.oak', 'lib/utils.oak']
results := files |> map(checkFormatting)

each(results, fn(r) {
    if r.formatted = false {
        true -> println('Needs formatting: ' + r.path)
    } |> r.error {
        ? -> ?
        _ -> println('Error: ' + r.path)
    }
})
```

### Format Changed Files in Git

```oak
{
    getChangedOakFiles: getChangedOakFiles
    formatFileInPlace: formatFileInPlace
} := import('syntaxfmt')

result := getChangedOakFiles()
if result.ok {
    true -> {
        println('Formatting ' + string(len(result.files)) + ' changed file(s)...')
        
        each(result.files, fn(file) {
            formatResult := formatFileInPlace(file)
            if formatResult.ok {
                true -> println('✓ ' + file)
                _ -> println('✗ ' + file + ': ' + formatResult.error)
            }
        })
    }
    _ -> println('Error: ' + result.error)
}
```

### Format with Diff Preview

```oak
{ formatFileWithDiff: formatFileWithDiff } := import('syntaxfmt')

fn previewFormatting(path) {
    result := formatFileWithDiff(path)
    
    if result.ok {
        true -> if result.diff {
            '' -> {
                println(path + ': Already formatted')
                { needsFormat: false }
            }
            _ -> {
                println(path + ' changes:')
                println(result.diff)
                { needsFormat: true, diff: result.diff }
            }
        }
        _ -> {
            println(path + ': Error - ' + result.error)
            { error: result.error }
        }
    }
}

previewFormatting('lib/utils.oak')
```

### CI/CD Formatting Check

```oak
{
    getChangedOakFiles: getChangedOakFiles
    formatFile: formatFile
} := import('syntaxfmt')

fn checkFormattingCI() {
    result := getChangedOakFiles()
    if !result.ok -> {
        println('Error: ' + result.error)
        exit(1)
    }
    
    unformatted := []
    
    each(result.files, fn(file) {
        original := readFile(file)
        formatResult := formatFile(file)
        
        if formatResult.ok & formatResult.content != original -> {
            unformatted <- append(unformatted, file)
        }
    })
    
    if len(unformatted) > 0 {
        true -> {
            println('The following files need formatting:')
            each(unformatted, fn(f) { println('  - ' + f) })
            exit(1)
        }
        _ -> {
            println('All files are properly formatted ✓')
            exit(0)
        }
    }
}

checkFormattingCI()
```

### Batch Format Directory

```oak
{ formatFileInPlace: formatFileInPlace } := import('syntaxfmt')
{ listFiles: listFiles } := import('fs')
{ endsWith?: endsWith? } := import('str')

fn formatDirectory(dir) {
    with listFiles(dir) fn(entries) if entries {
        ? -> println('Error reading directory: ' + dir)
        _ -> {
            oakFiles := entries |> filter(fn(e) {
                !e.dir? & e.name |> endsWith?('.oak')
            })
            
            each(oakFiles, fn(file) {
                path := dir + '/' + file.name
                result := formatFileInPlace(path)
                
                if result.ok {
                    true -> println('✓ ' + path)
                    _ -> println('✗ ' + path)
                }
            })
        }
    }
}

formatDirectory('lib')
formatDirectory('src')
```

## Implementation Notes

- Uses `syntax.print()` for formatting
- File operations use `fs` library
- Diff generation uses system `diff` command
- Git integration uses `exec()` to run `git diff`
- Returns structured result objects (not exceptions)

## Formatting Style

The formatter applies consistent Oak style:
- Standard indentation
- Consistent spacing around operators
- Proper bracket/brace alignment
- Normalized whitespace

## Error Handling

All functions return result objects:

```oak
// Success
{ ok: true, content: '...' }  // or { ok: true }

// Error
{ ok: false, error: 'Could not read file' }
```

## Limitations

- Requires `diff` command for diff generation
- Requires `git` command for changed files
- No configuration options (fixed style)
- Comments may be reformatted
- Original file replaced (no automatic backup)

## Use Cases

- **Code formatting**: Enforce consistent style
- **CI/CD**: Check formatting in pipelines
- **Pre-commit hooks**: Format before commit
- **Batch processing**: Format entire projects
- **Editor integration**: Format on save

## See Also

- [syntax.md](syntax.md) - Oak tokenizer/parser
- [fs.md](fs.md) - File operations
- [fmt.md](fmt.md) - String formatting
- `syntax.print()` - Core formatting function
