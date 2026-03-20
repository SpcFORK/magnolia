# Build Identifier Formatter (build-ident)

## Overview

`build-ident` maps Oak identifier names to their safe rendered forms for native Oak bundles and JavaScript bundles. When targeting JavaScript it escapes `?` and `!` characters and avoids all ECMAScript reserved words.

This module is used internally by the code-generator in `oak build` and is not intended for direct use in application code.

## Import

```oak
buildIdent := import('build-ident')
{ formatIdent: formatIdent } := import('build-ident')
```

## Functions

### `formatIdent(web?, name, key?)`

Returns a safe identifier string for `name` in the target context.

**Parameters**

- `web?` — `true` when targeting JavaScript output.
- `name` — the Oak identifier name.
- `key?` — an optional disambiguating key appended to `__oak_empty_ident` when translating the special `_` identifier.

**Native target (`web? = false`)**

- `'import'` → `'__oak_module_import'` (avoids shadowing the built-in).
- All other names are returned unchanged.

**JavaScript target (`web? = true`)**

- `'_'` → `'__oak_empty_ident'` (optionally suffixed with `key`).
- ECMAScript reserved words (e.g. `class`, `return`, `let`, `yield`, …) → `'__oak_js_' << name`.
- `'import'` → `'__oak_module_import'`.
- `'?'` in name → replaced with `'__oak_qm'`.
- `'!'` in name → replaced with `'__oak_exclam'`.
- All other names are returned as-is.

```oak
formatIdent(false, 'myVar')   // => 'myVar'
formatIdent(true, 'class')    // => '__oak_js_class'
formatIdent(true, 'done?')    // => 'done__oak_qm'
formatIdent(true, 'import')   // => '__oak_module_import'
formatIdent(true, '_', 0)     // => '__oak_empty_ident0'
```
