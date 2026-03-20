# JavaScript Bundle Runtime (runtime-js)

## Overview

`runtime-js` contains the JavaScript runtime preamble that is prepended to every Oak-to-JavaScript bundle produced by `oak build --web`. It implements Oak's module system, language primitives (equality, string, list, object operations), and the standard library shim layer in pure JavaScript.

This module is part of the build toolchain and is not intended for use in application code.

## Import

```oak
js := import('runtime-js')
{ OakJSRuntime: OakJSRuntime } := import('runtime-js')
```

## Exported Values

### `OakJSRuntime`

A string containing the full JavaScript runtime source, approximately 600–800 lines. It is prepended verbatim by `build-render.renderJSBundle` before the bundled module code.

**Key features provided by the runtime:**

| Feature                        | JS identifier                  |
|--------------------------------|-------------------------------|
| Module registration            | `__oak_modularize(name, fn)`  |
| Module import                  | `__oak_module_import(name)`   |
| Structural equality            | `__oak_eq(a, b)`              |
| Property access with `?` default | `__oak_acc(tgt, prop)`      |
| Object key normalization       | `__oak_obj_key(x)`            |
| Push (`<<`)                    | `__oak_push(a, b)`            |
| Bitwise AND on strings/bools   | `__oak_and(a, b)`             |
| Bitwise OR on strings/bools    | `__oak_or(a, b)`              |
| Bitwise XOR on strings/bools   | `__oak_xor(a, b)`             |
| Oak string wrapper             | `__OakString`, `__as_oak_string`, `__is_oak_string` |
| Oak Empty sentinel             | `__Oak_Empty`                 |
| Trampoline for tail calls      | `__oak_resolve_trampoline`    |
| Standard library built-ins     | `len`, `print`, `string`, `int`, `keys`, etc. |

```oak
// Typical usage inside build-render:
output := OakJSRuntime << renderedBundleSource
writeFile('out.js', output)
```
