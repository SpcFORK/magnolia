# Build Import Resolver (build-imports)

## Overview

`build-imports` handles static import-graph traversal for `oak build`. It walks the AST of each source file, identifies `import('...')` calls with string-literal arguments, and queues their dependencies for parsing. The result is a complete transitive-closure of all modules required by the entry point.

This module is part of `oak build`'s internal pipeline and is not intended for direct use in application code.

## Import

```oak
buildImports := import('build-imports')
{ cachedParse: cachedParse, addImportsFromSource: addImportsFromSource, addImportsFromFile: addImportsFromFile } := import('build-imports')
```

## Functions

### `cachedParse(path, text, moduleNodes, log, abort, syntax, format)`

Parses `text` (the source for `path`) using `syntax.parse()`, caches the result in `moduleNodes.(path)`, and returns the AST node list. Returns the cached result immediately when `moduleNodes.(path)` is already populated.

Calls `abort` with a formatted error message when the parse produces a `:error` node.

```oak
nodes := cachedParse(
    'main.oak'
    src
    moduleNodes
    printf, abortFn, syntax, format
)
```

### `addImportsFromSource(path, file, moduleNodes, importAssignmentNode, web?, runtimeLibQ, runtimeLib, resolve, dir, addImportsFromFile, cachedParseFn)`

Parses `file` as the source for `path` and recursively discovers all statically-known `import(...)` references. For each dependency:

- If it is a runtime library (identified by `runtimeLibQ`) and `web?` is `true`, its source is fetched via `runtimeLib(name)` and recursed into.
- Otherwise the resolved file path is passed to `addImportsFromFile` if not already in `moduleNodes`.

### `addImportsFromFile(path, readFile, printf, addImportsFromSource)`

Reads `path` from disk via `readFile`, warns with `printf` when the file does not exist, and passes the content to `addImportsFromSource`.

```oak
addImportsFromFile(
    'lib/util.oak'
    readFile, printf
    fn(p, src) addImportsFromSource(p, src, ...)
)
```
