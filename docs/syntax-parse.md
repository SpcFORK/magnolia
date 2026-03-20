# Syntax Parser (syntax-parse)

## Overview

`syntax-parse` is the second stage of Oak's parsing pipeline. It consumes the filtered token stream produced by `syntax-tokenize` and builds a typed AST node tree.

## Import

```oak
parseLib := import('syntax-parse')
{ Parser: Parser } := import('syntax-parse')
```

## Functions

### `Parser(tokens)`

Creates a parser object over `tokens` (a list of token objects as returned by `Tokenizer.tokenize()`, with `:newline` and `:comment` tokens included — the parser filters them internally).

Returns an object with a single method:

#### `.parse()`

Parses the full token stream and returns a list of top-level AST nodes.

On success each element is a valid AST node object (see **AST Node Types** below).

On parse error, a single `:error` node is returned:

```oak
{ type: :error, error: 'Unexpected token ...', pos: [index, line, col] }
```

```oak
{
    tokenize: tokenize
} := import('syntax-tokenize').Tokenizer(src)
{
    parse: parse
} := import('syntax-parse').Parser(tokenize())

nodes := parse()
if nodes.(0).type = :error -> printf('parse error: {{0}}', nodes.(0).error)
```

## AST Node Types

All nodes share a `type` atom and a `tok` field containing the triggering token.

### Literals

| `type`     | Extra fields            |
|------------|-------------------------|
| `:null`    | —                       |
| `:empty`   | —                       |
| `:int`     | `val` (number)          |
| `:float`   | `val` (number)          |
| `:string`  | `val` (string)          |
| `:bool`    | `val` (bool)            |
| `:atom`    | `val` (string)          |

### Composite

| `type`     | Extra fields                                    |
|------------|-------------------------------------------------|
| `:list`    | `elems: [node, ...]`                            |
| `:object`  | `entries: [{ key: node, val: node }, ...]`      |

### Identifiers

| `type`       | Extra fields        |
|--------------|---------------------|
| `:identifier`| `val` (string name) |

### Expressions

| `type`            | Extra fields                                                    |
|-------------------|-----------------------------------------------------------------|
| `:unary`          | `op` (atom), `right` (node)                                     |
| `:binary`         | `op` (atom), `left` (node), `right` (node)                      |
| `:assignment`     | `local?` (bool), `left` (node), `right` (node)                  |
| `:propertyAccess` | `left` (node), `right` (node)                                   |
| `:fnCall`         | `function` (node), `args: [node]`, `restArg` (node or `?`)     |
| `:ifExpr`         | `cond` (node), `branches: [ifBranch, ...]`                      |
| `:ifBranch`       | `target` (node), `body` (node)                                  |
| `:block`          | `exprs: [node, ...]`                                            |
| `:function`       | `name` (string), `args: [string]`, `restArg` (string), `body` (node) |

### Binary Operators

| Atom        | Symbol | Atom      | Symbol |
|-------------|--------|-----------|--------|
| `:plus`     | `+`    | `:minus`  | `-`    |
| `:times`    | `*`    | `:divide` | `/`    |
| `:modulus`  | `%`    | `:and`    | `&`    |
| `:or`       | `\|`   | `:xor`    | `^`    |
| `:rshift`   | `>>`   | `:pushArrow` | `<<` |
| `:eq`       | `=`    | `:neq`    | `!=`   |
| `:greater`  | `>`    | `:less`   | `<`    |
| `:geq`      | `>=`   | `:leq`    | `<=`   |
