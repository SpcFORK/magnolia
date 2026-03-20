# Build AST Helpers (build-ast)

## Overview

`build-ast` provides AST wrapping and path-normalization helpers used by the Oak bundler. It turns individual module ASTs into closure-wrapped module nodes and combines them into a single bundle AST.

This module is part of `oak build`'s internal pipeline and is not intended for direct use in application code.

## Import

```oak
buildAst := import('build-ast')
{ wrapModule: wrapModule, wrapBundle: wrapBundle } := import('build-ast')
```

## Functions

### `wrapModule(block)`

Wraps the body of a parsed module (`block` is a `:block` AST node) in a closure that returns an object containing all top-level declarations. This is the module isolation mechanism used in Oak bundles.

```oak
block := syntax.parse(src)
moduleNode := wrapModule(block)
```

The returned node is a `:function` node whose body ends with an object literal exporting all declared names from `block.decls`.

### `wrapBundle(modules, entryModuleName, includes, importCallNode)`

Combines a list of `[path, moduleNode]` pairs into a single `:block` bundle AST, ready for rendering.

**Parameters**

- `modules` — list of `[modulePath, wrappedModuleNode]` pairs.
- `entryModuleName` — the path of the entry-point module; its module function is called last to boot the program.
- `includes` — list of include spec objects `{ name, path }` for files included via `--include`.
- `importCallNode` — the AST node pattern used to identify `import(...)` calls during path rewriting.

```oak
bundleAst := wrapBundle(modules, entryPath, includes, importNode)
```

All `import('...')` calls inside each module have their argument rewritten to the normalized bundle-relative path before the bundle is assembled.

### `bundleCommonPrefix(paths)`

Returns the longest common path prefix shared by all `paths` in the bundle, used to strip redundant directory segments from module keys. Returns `''` when fewer than two paths are provided.

```oak
prefix := bundleCommonPrefix(['/src/a.oak', '/src/b.oak'])
// => '/src/'
```
