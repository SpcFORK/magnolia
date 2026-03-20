# Build Analyzer (build-analyze)

## Overview

`build-analyze` performs static semantic analysis on Oak ASTs produced by the parser. It annotates the tree with declaration sets, marks tail-recursive calls for trampoline rewriting (when targeting JavaScript), and sets up scoping context for each block and function node.

This module is used internally by `oak build` and is not intended for direct use in application code.

## Import

```oak
analyze := import('build-analyze')
{ analyzeNode: analyzeNode } := import('build-analyze')
```

## Functions

### `analyzeNode(node, web?)`

Performs a full semantic pass over an AST `node` (typically a `:block` returned by the parser). When `web?` is `true`, tail-recursive function calls are rewritten into trampoline form for safe JavaScript tail-call elimination.

Returns the annotated AST node. The top-level node receives a `.decls` list of all top-level identifiers declared within it.

```oak
syntax := import('syntax')
ast := syntax.parse(srcText)
annotated := analyzeNode(ast, false)  // native target
```

```oak
// For JS bundling:
annotated := analyzeNode(ast, true)
```

**Annotation details**

- Every `:block` and `:function` node gains a `.decls` list of identifier names declared within that scope.
- For `:function` nodes with the same name as the enclosing function, recursive tail calls are detected and `node.recurred?` is set to `true`.
- When `web?` is `true` and `recurred?` is set, the function body is rewritten to use `__oak_resolve_trampoline` for stack-safe recursion in JS.
