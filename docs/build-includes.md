# Build Include Parser (build-includes)

## Overview

`build-includes` parses `--include` specifications for `oak build`. An include spec maps a virtual module name to a file path, allowing additional Oak files to be embedded in a bundle under a chosen import name.

This module is part of `oak build`'s internal pipeline and is not intended for direct use in application code.

## Import

```oak
buildIncludes := import('build-includes')
{ parseInclude: parseInclude, parseIncludes: parseIncludes } := import('build-includes')
```

## Include Spec Format

A single spec is a colon-separated `name:path` string. When no colon is present, both the name and path are derived from the bare path. The path is resolved relative to the current working directory, and `.oak` is appended automatically when no recognised extension is present (`.oak`, `.ok`, `.mag`, `.mg`).

```
myLib:lib/myLib.oak   → { name: 'myLib',    path: '/abs/lib/myLib.oak' }
lib/extra             → { name: 'lib/extra', path: '/abs/lib/extra.oak' }
```

## Functions

### `parseInclude(spec)`

Parses a single include specification string and returns a `{ name, path }` object.

```oak
parseInclude('util:lib/util.oak')
// => { name: 'util', path: '/project/lib/util.oak' }

parseInclude('lib/helper')
// => { name: 'lib/helper', path: '/project/lib/helper.oak' }
```

### `parseIncludes(specs)`

Accepts either:
- A comma-separated string of specs (as passed from the CLI `--include` flag).
- A list of strings or pre-parsed `{ name, path }` objects.

Returns a list of `{ name, path }` objects, filtering empty entries.

```oak
parseIncludes('util:lib/util.oak,extra')
// => [
//   { name: 'util',  path: '/project/lib/util.oak' }
//   { name: 'extra', path: '/project/extra.oak' }
// ]

parseIncludes(['util:lib/util.oak', { name: 'x', path: '/abs/x.oak' }])
// => [
//   { name: 'util', path: '/project/lib/util.oak' }
//   { name: 'x',    path: '/abs/x.oak' }
// ]
```
