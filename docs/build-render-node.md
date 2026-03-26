# Build Render Node (build-render-node)

## Overview

`build-render-node` provides a standalone Oak AST node renderer extracted from `build-render`. It converts AST nodes (as produced by `syntax.Parser` or `build-ast`) into Oak source text. Unlike `build-render`, it can be used outside the full build pipeline — for example, to render individual AST nodes in save files, code generators, or REPL tools.

## Import

```oak
buildRenderNode := import('build-render-node')
```

## Functions

### `createRenderer(formatIdent, abort)`

Returns a `renderNode(node)` function bound to the given callbacks.

**Parameters**

- `formatIdent` — `fn(name)` that transforms identifier names (e.g. for minification). Pass `fn(name) name` for no transformation.
- `abort` — `fn(msg)` called on unexpected AST node types. Can raise an error or return an error string.

**Returns** a `renderNode` function.

```oak
buildIdent := import('build-ident')
formatIdent := buildIdent.formatIdent(false)

renderNode := buildRenderNode.createRenderer(formatIdent, fn(msg) msg)
source := renderNode(astNode)
```

### `makeRenderOakNode()`

Factory that returns a `renderNode` function with identity formatting (no renaming) and a no-op abort (returns the error message as a string). Equivalent to `createRenderer(fn(name) name, fn(msg) msg)`.

Use this when you need the renderer function itself — for example, to pass it as a callback or call it multiple times without re-creating it.

```oak
buildRenderNode := import('build-render-node')

renderNode := buildRenderNode.makeRenderOakNode()
renderNode(astNode1)
renderNode(astNode2)
```

### `renderOakNode(node)`

Convenience wrapper that renders an AST node with identity formatting (no renaming) and a no-op abort (returns the error message as a string). Calls `makeRenderOakNode()` internally.

```oak
syntax := import('syntax')
buildRenderNode := import('build-render-node')

tokens := syntax.Tokenizer('x := 1 + 2').tokenize()
ast := syntax.Parser(tokens).parse()
source := buildRenderNode.renderOakNode(ast)
// => 'x:=1+2'
```

### `renderOakVal(value)`

Renders a runtime Oak value (not an AST node) to Oak source text. Handles all primitive types, lists, and objects with proper string escaping.

```oak
buildRenderNode := import('build-render-node')

buildRenderNode.renderOakVal('hello')
// => '\'hello\''

buildRenderNode.renderOakVal([1, 2, 3])
// => '[1, 2, 3]'

buildRenderNode.renderOakVal({ name: 'Alice', score: 42 })
// => '{name: \'Alice\', score: 42}'
```

This is useful for generating importable Oak source files from runtime data (e.g. save files, config generators).

## Relationship to build-render

`build-render` now imports `build-render-node` internally for its Oak bundle renderer (`renderOakBundle`). The JS bundle renderer remains self-contained in `build-render` because it has JS-specific logic (clone, renderAssignTarget, etc.).
