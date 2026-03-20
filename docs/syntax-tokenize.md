# Syntax Tokenizer (syntax-tokenize)

## Overview

`syntax-tokenize` is the first stage of Oak's parsing pipeline. It converts raw Oak source text into a stream of typed token objects, preserving shebangs, newlines, and comments alongside the semantic tokens. The token stream is later filtered and consumed by `syntax-parse`.

## Import

```oak
tok := import('syntax-tokenize')
{ Tokenizer: Tokenizer, renderToken: renderToken, renderPos: renderPos, shebang?: shebang? } := import('syntax-tokenize')
```

## Functions

### `shebang?(text)`

Returns `true` when `text` begins with a shebang (`#!`) line.

```oak
shebang?('#!/usr/bin/env oak\nfn main ...')  // => true
shebang?('fn main ...')                      // => false
```

### `renderPos(pos)`

Formats a position triple `[index, line, col]` as a `[line:col]` string.

```oak
renderPos([0, 3, 12])  // => '[3:12]'
```

### `renderToken(token)`

Returns a human-readable string describing a token, including its type, optional value, and position. Useful for error messages.

```oak
renderToken({ type: :identifier, val: 'foo', pos: [0, 1, 1] })
// => ':identifier(foo) [1:1]'
```

### `Tokenizer(source)`

Creates a tokenizer object for `source`. Call `.tokenize()` on it to produce the full token list.

```oak
{ tokenize: tokenize } := Tokenizer(sourceCode)
tokens := tokenize()
```

The returned token list includes all token types, including `:newline` and `:comment`. Filter these out before passing to the parser when building an AST.

## Token Objects

Each token is an object with the following fields:

| Field  | Type   | Description                                             |
|--------|--------|---------------------------------------------------------|
| `type` | atom   | Token type (see table below).                           |
| `val`  | string | Token value, or `?` for punctuation with no content.   |
| `pos`  | list   | `[byteIndex, line, col]` — 1-based line and column.    |

## Token Types

### Literals

| Type              | Example         |
|-------------------|-----------------|
| `:stringLiteral`  | `'hello'`       |
| `:numberLiteral`  | `42`, `3.14`    |
| `:trueLiteral`    | `true`          |
| `:falseLiteral`   | `false`         |

### Identifiers and Keywords

| Type           | Example        |
|----------------|----------------|
| `:identifier`  | `foo`, `bar?`  |
| `:atom`        | `:ok`, `:error`|
| `:ifKeyword`   | `if`           |
| `:fnKeyword`   | `fn`           |
| `:withKeyword` | `with`         |
| `:csKeyword`   | `cs`           |
| `:underscore`  | `_`            |

### Operators

| Type              | Symbol  | Type          | Symbol |
|-------------------|---------|---------------|--------|
| `:assign`         | `:=`    | `:nonlocalAssign` | `<-` |
| `:pipeArrow`      | `\|>`   | `:branchArrow`  | `->`  |
| `:pushArrow`      | `<<`    | `:rshift`       | `>>`  |
| `:plus`           | `+`     | `:minus`        | `-`   |
| `:times`          | `*`     | `:divide`       | `/`   |
| `:modulus`        | `%`     | `:xor`          | `^`   |
| `:and`            | `&`     | `:or`           | `\|`  |
| `:eq`             | `=`     | `:neq`          | `!=`  |
| `:greater`        | `>`     | `:less`         | `<`   |
| `:geq`            | `>=`    | `:leq`          | `<=`  |
| `:exclam`         | `!`     | `:colon`        | `:`   |
| `:ellipsis`       | `...`   | `:qmark`        | `?`   |

### Delimiters

| Type             | Symbol |
|------------------|--------|
| `:leftParen`     | `(`    |
| `:rightParen`    | `)`    |
| `:leftBracket`   | `[`    |
| `:rightBracket`  | `]`    |
| `:leftBrace`     | `{`    |
| `:rightBrace`    | `}`    |
| `:dot`           | `.`    |
| `:comma`         | `,`    |

### Non-Semantic

| Type       | Notes                        |
|------------|------------------------------|
| `:newline` | Significant for semicolon insertion; usually filtered. |
| `:comment` | `// ...` line comment content.                         |
