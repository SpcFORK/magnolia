# Native Oak Bundle Runtime (runtime-native)

## Overview

`runtime-native` contains the Oak runtime preamble that is prepended to every native Oak bundle produced by `oak build` (non-web, non-wasm). It implements the module system used in bundled programs: each module is registered as a closure and lazily evaluated on first import.

This module is part of the build toolchain and is not intended for use in application code.

## Import

```oak
native := import('runtime-native')
{ OakNativeRuntime: OakNativeRuntime } := import('runtime-native')
```

## Exported Values

### `OakNativeRuntime`

A string containing the Oak module-system runtime source. It is prepended verbatim by `build-render.renderOakBundle` before the bundled module code.

**Provided definitions:**

| Identifier              | Description                                                        |
|-------------------------|--------------------------------------------------------------------|
| `__Oak_Modules`         | Global map from module path → registered module closure or object. |
| `__Oak_Import_Aliases`  | Optional alias map for alternative module paths.                   |
| `__oak_modularize(name, module)` | Registers a module closure under `name`.              |
| `__oak_module_import(name)`      | Imports a module by name, evaluating its closure once and caching the result. Falls back to the built-in `import` for runtime stdlib paths. |
| `___packed_vfs()`       | Returns `__Oak_VFS` when a Virtual FS is embedded, otherwise `?`.  |

```oak
// Typical usage inside build-render:
output := OakNativeRuntime << renderedBundleSource
writeFile('out.oak', output)
```

The module-import function handles circular imports by checking whether the cached value is a function (unevaluated) before calling it.
