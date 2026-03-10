# Syntax Library (syntax)

## Overview

`libsyntax` provides an Oak tokenizer and parser for analyzing Oak source code, enabling syntax highlighting, code analysis, and AST-level metaprogramming with macros.

## Import

```oak
syntax := import('syntax')
{ Tokenizer: Tokenizer } := import('syntax')
{ parseWithMacros: parseWithMacros, Macro: Macro } := import('syntax')
```

## Metaprogramming and Macros

`syntax` now includes AST macro expansion helpers.

### `Macro(expander)`

Wraps a macro expander function.

**Parameters:**
- `expander(args, callNode, macros)` - Function returning a replacement AST node

**Returns:** Macro descriptor object

### `macro?(value)`

Checks whether a value is a macro descriptor created by `Macro`.

### `expandMacros(ast, macros)`

Recursively walks AST node(s) and expands macro calls.

**Parameters:**
- `ast` - One AST node or a list of AST nodes
- `macros` - Object mapping identifier names to `Macro(...)` values

**Returns:** Expanded AST node(s)

### `parseWithMacros(text, macros)`

Parses source and applies macro expansion in one step.

```oak
syntax := import('syntax')

expanded := syntax.parseWithMacros('inc(2)', {
    inc: syntax.Macro(fn(args) {
        type: :binary
        tok: args.0.tok
        op: :plus
        left: args.0
        right: { type: :int, tok: args.0.tok, val: 1 }
    })
})

// expanded.0 is now a :binary node representing 2 + 1
```

## Components

### `Tokenizer(source)`

Creates a tokenizer for Oak source code.

**Parameters:**
- `source` - Oak source code string

**Returns:** Tokenizer object

```oak
{ Tokenizer: Tokenizer } := import('syntax')

code := '
fn hello(name) {
    println(\'Hello, \' + name)
}
'

tokenizer := Tokenizer(code)
tokens := tokenizer.tokenize()

each(tokens, fn(tok) {
    println(tok.type + ': ' + tok.value)
})
```

## Token Types

Oak recognizes these token types:

### Keywords
- `fn`, `if`, `with`, `each`
- `true`, `false`, `?` (null)

### Identifiers
- Variable names: `myVar`, `count`, `firstName`
- Atom names: `:atom`, `:type`, `:success`

### Literals
- **Numbers**: `42`, `3.14`, `0xFF`, `2e10`
- **Strings**: `'hello'`, `"world"`
- **Atoms**: `:atom`, `:value`

### Operators
- Arithmetic: `+`, `-`, `*`, `/`, `%`
- Comparison: `<`, `>`, `=`, `!=`
- Logical: `&`, `|`, `!`
- Assignment: `<-`
- Pipeline: `|>`
- Property access: `.`

### Delimiters
- `(`, `)` - Function calls, grouping
- `{`, `}` - Blocks, objects
- `[`, `]` - Lists
- `,` - Separators
- `:` - Key-value pairs in objects

### Comments
- Single-line: `// comment`
- Block comments: Not standard in Oak

## Tokenizer API

### `tokenize()`

Returns a list of all tokens in the source.

**Returns:** List of token objects

```oak
tokens := tokenizer.tokenize()

// Each token has:
// {
//   type: 'keyword' | 'ident' | 'number' | 'string' | 'operator' | ...
//   value: 'actual text'
//   line: 1
//   col: 0
// }
```

## Usage Examples

### Syntax Highlighting

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn highlightOak(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    html := ''
    each(tokens, fn(tok) {
        cssClass := if tok.type {
            :keyword -> 'keyword'
            :number -> 'number'
            :string -> 'string'
            :atom -> 'atom'
            :comment -> 'comment'
            :ident -> 'ident'
            _ -> ''
        }
        
        html <- html + '<span class="' + cssClass + '">'
        html <- html + escapeHTML(tok.value)
        html <- html + '</span>'
    })
    
    html
}

highlighted := highlightOak('fn add(a, b) { a + b }')
```

### Count Function Definitions

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn countFunctions(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    count := 0
    each(tokens, fn(tok) {
        if tok.type = :keyword & tok.value = 'fn' -> {
            count <- count + 1
        }
    })
    
    count
}

functions := countFunctions(sourceCode)
println('Functions defined: ' + string(functions))
```

### Extract String Literals

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn extractStrings(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    strings := []
    each(tokens, fn(tok) {
        if tok.type = :string -> {
            strings <- append(strings, tok.value)
        }
    })
    
    strings
}

strings := extractStrings(code)
println('String literals: ' + string(len(strings)))
```

### Find Variable Names

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn findIdentifiers(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    idents := {}
    each(tokens, fn(tok) {
        if tok.type = :ident -> {
            idents.(tok.value) := true
        }
    })
    
    keys(idents) // Unique identifiers
}

variables := findIdentifiers(sourceCode)
```

### Validate Syntax

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn validateBalanced(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    stack := []
    pairs := {
        '(': ')'
        '[': ']'
        '{': '}'
    }
    
    valid := true
    each(tokens, fn(tok) {
        if tok.value {
            '(' | '[' | '{' -> {
                stack <- append(stack, tok.value)
            }
            ')' | ']' | '}' -> {
                if len(stack) = 0 -> {
                    valid <- false
                } else {
                    opening := stack.(len(stack) - 1)
                    if pairs.(opening) = tok.value {
                        true -> stack <- slice(stack, 0, len(stack) - 1)
                        _ -> valid <- false
                    }
                }
            }
        }
    })
    
    valid & len(stack) = 0
}

if validateBalanced(code) {
    true -> println('Brackets balanced ✓')
    _ -> println('Unbalanced brackets ✗')
}
```

### Token Statistics

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn analyzeCode(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    stats := {
        keywords: 0
        identifiers: 0
        numbers: 0
        strings: 0
        operators: 0
        total: len(tokens)
    }
    
    each(tokens, fn(tok) {
        if tok.type {
            :keyword -> stats.keywords <- stats.keywords + 1
            :ident -> stats.identifiers <- stats.identifiers + 1
            :number -> stats.numbers <- stats.numbers + 1
            :string -> stats.strings <- stats.strings + 1
            :operator -> stats.operators <- stats.operators + 1
        }
    })
    
    stats
}

stats := analyzeCode(sourceCode)
println('Code statistics:')
println('  Keywords: ' + string(stats.keywords))
println('  Identifiers: ' + string(stats.identifiers))
println('  Numbers: ' + string(stats.numbers))
println('  Strings: ' + string(stats.strings))
```

### Simple Formatter

```oak
{ Tokenizer: Tokenizer } := import('syntax')

fn formatOak(code) {
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    formatted := ''
    indent := 0
    
    each(tokens, fn(tok) {
        if tok.value = '{' -> {
            formatted <- formatted + ' {\n'
            indent <- indent + 1
            formatted <- formatted + repeat('  ', indent)
        } |> tok.value = '}' -> {
            indent <- indent - 1
            formatted <- formatted + '\n' + repeat('  ', indent) + '}'
        } |> tok.type = :newline -> {
            formatted <- formatted + '\n' + repeat('  ', indent)
        } else {
            formatted <- formatted + tok.value + ' '
        }
    })
    
    formatted
}
```

## Oak Syntax Elements

### Function Declarations

```oak
fn name(param1, param2) { body }
```

### Conditionals

```oak
if condition { true -> action }
if value { case1 -> result1, case2 -> result2, _ -> default }
```

### Loops

```oak
with std.loop() fn(again) { again() }
each(list, fn(item) { /* ... */ })
```

### Objects

```oak
obj := {
    key: 'value'
    method: fn { /* ... */ }
}
```

### Lists

```oak
list := [1, 2, 3, 4]
```

### Pipeline

```oak
value |> fn1() |> fn2() |> fn3()
```

## Token Properties

Each token object contains:

```oak
{
    type: :keyword      // Token type (atom)
    value: 'fn'         // Literal text
    line: 1             // Line number (1-indexed)
    col: 0              // Column number (0-indexed)
}
```

## Use Cases

- **Syntax highlighting** in editors
- **Code analysis** and metrics
- **Linting** and style checking
- **Code transformation** tools
- **Documentation generation**
- **IDE features** (autocomplete, etc.)
- **Formatting** and beautification

## Limitations

- Tokenization only (no full AST parsing)
- No semantic analysis
- No type information
- No error recovery
- Position information may be approximate
- Does not validate syntax correctness beyond tokenization
- Comments may not preserve all whitespace

## Related Patterns

### Combined with File I/O

```oak
{ Tokenizer: Tokenizer } := import('syntax')
{ readFile: readFile } := import('fs')

code := readFile('script.oak')
tokens := Tokenizer(code).tokenize()
```

### Process Multiple Files

```oak
{ Tokenizer: Tokenizer } := import('syntax')

each(files, fn(file) {
    code := readFile(file)
    tokenizer := Tokenizer(code)
    tokens := tokenizer.tokenize()
    
    analyzeSyntax(file, tokens)
})
```

## Performance Notes

- Tokenization is streaming (efficient for large files)
- Entire source must be in memory as string
- Token list stores all tokens (memory overhead for large files)

## See Also

- `str` library - String manipulation for token processing
- `fs` library - Reading source files
- `fmt` library - Formatting output  
- Oak language specification
