# Syntax Printer (syntax-print)

## Overview

`syntax-print` pretty-prints a stream of Oak tokens back into formatted source code. It preserves all newlines (including those inside comments) and adds appropriate spacing around operators and keywords. This is the backend used by `oak fmt`.

## Import

```oak
printLib := import('syntax-print')
{ Printer: Printer } := import('syntax-print')
```

## Functions

### `Printer(tokens)`

Creates a printer object over `tokens` (a list of token objects as returned by `Tokenizer.tokenize()`). Returns an object with a single method:

#### `.print()`

Formats the token stream and returns the pretty-printed source string.

```oak
{
    tokenize: tokenize
} := import('syntax-tokenize').Tokenizer(src)
{
    print: print
} := import('syntax-print').Printer(tokenize())

formatted := print()
```

## Formatting Rules

- **Indentation** uses tabs. The printer tracks brace/bracket/paren depth and indents accordingly.
- **Newlines** from the token stream are preserved as-is; the printer does not add or remove newlines except to ensure proper indentation after them.
- **Operators and keywords** (`:=`, `<-`, `|>`, `->`, `<<`, arithmetic/comparison ops, `if`, `fn`, `with`, `cs`) are surrounded by a single space.
- **Commas** are followed by a space when on the same line.
- **Parentheses, brackets, and braces** have no inner padding.
- **Comments** (`//`) are reproduced verbatim including their content.

## Example

```oak
src := 'fn add(a,b)a+b'
{ tokenize: tokenize } := import('syntax-tokenize').Tokenizer(src)
{ print: print } := import('syntax-print').Printer(tokenize())
print()
// => 'fn add(a, b) a + b'
```
