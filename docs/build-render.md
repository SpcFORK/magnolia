# Build Renderers (build-render)

## Overview

`build-render` converts a bundle AST produced by `build-ast` into a self-contained source string for either a native Oak bundle or a JavaScript bundle. It handles all AST node types and embeds the appropriate runtime preamble and Virtual FS initializer.

This module is part of `oak build`'s internal code-generation pipeline and is not intended for direct use in application code.

## Import

```oak
buildRender := import('build-render')
{ renderOakBundle: renderOakBundle, renderJSBundle: renderJSBundle } := import('build-render')
```

## Functions

### `renderOakBundle(bundleNode, vfsFiles, encodeJSON, oakNativeRuntime, formatIdent, abort)`

Serializes a bundle AST node to Oak source code.

**Parameters**

- `bundleNode` — the `:block` bundle AST from `wrapBundle`.
- `vfsFiles` — a map of virtual-FS paths to file contents; an empty map omits the VFS initializer.
- `encodeJSON` — function to encode a value as a JSON string (typically `json.encode`).
- `oakNativeRuntime` — the native runtime preamble string (from `runtime-native`).
- `formatIdent` — identifier formatter function (from `build-ident`).
- `abort` — called with an error message on unexpected AST node types.

Returns a string containing the complete native Oak bundle source.

```oak
source := renderOakBundle(bundleAst, {}, json.encode, OakNativeRuntime, formatIdent, abortFn)
writeFile('out.oak', source)
```

### `renderJSBundle(bundleNode, vfsFiles, encodeJSON, oakJSRuntime, formatIdent, clone)`

Serializes a bundle AST node to JavaScript source code.

**Parameters**

- `bundleNode` — the `:block` bundle AST from `wrapBundle`.
- `vfsFiles` — a map of virtual-FS paths to file contents.
- `encodeJSON` — function to encode a value as a JSON string.
- `oakJSRuntime` — the JS runtime preamble string (from `runtime-js`).
- `formatIdent` — identifier formatter with `web? = true`.
- `clone` — the `std.clone` function, used internally during rendering.

Returns a string containing the complete JavaScript bundle source.

```oak
source := renderJSBundle(bundleAst, {}, json.encode, OakJSRuntime, formatIdentWeb, std.clone)
writeFile('out.js', source)
```
