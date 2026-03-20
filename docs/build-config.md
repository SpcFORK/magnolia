# Build Configuration (build-config)

## Overview

`build-config` validates and applies the user-supplied build configuration object to the shared `state` record used throughout the `oak build` pipeline. It also registers transpilers and wires up logging/abort hooks.

This module is part of `oak build`'s internal pipeline and is not intended for direct use in application code.

## Import

```oak
buildConfig := import('build-config')
{ configure: configure, requireOpt: requireOpt } := import('build-config')
```

## Functions

### `configure(config, state, parseIncludes, statFile, resolve, collectVFSFiles, transpile, each, printf, format)`

Applies a `config` object to `state`, validates required fields, and configures the `transpile` subsystem. Calls `state.Abort` (which defaults to printing an error and calling `exit(1)`) when a required option is missing or invalid.

**Config fields**

| Field             | Type      | Description                                              |
|-------------------|-----------|----------------------------------------------------------|
| `entry`           | string    | Path to the entry-point Oak file. **Required.**          |
| `output`          | string    | Output file path. **Required.**                          |
| `web?`            | bool      | Target JavaScript (web) output.                          |
| `wasm?`           | bool      | Target WebAssembly output. Mutually exclusive with `web?`. |
| `includes`        | string/list | Additional files to include in the bundle.             |
| `includeVFS`      | string    | Path glob for Virtual FS embedding.                      |
| `transpile?`      | bool      | Enable transpiler (default: `true`).                     |
| `transpileVerbose?` | bool   | Verbose transpiler logging (default: `false`).           |
| `transpilers`     | list      | Extra transpiler objects to register.                    |
| `log`             | function  | Custom log function (default: `printf`).                 |
| `abort`           | function  | Custom abort function (default: log + `exit(1)`).        |

```oak
configure(
    { entry: 'main.oak', output: 'out.oak' }
    state
    parseIncludes, statFile, resolve, collectVFSFiles, transpile, each, printf, format
)
```

### `requireOpt(flag, value, abort, format)`

Calls `abort` with a formatted message when `value` is `?` or `''`. Used internally by `configure` to validate required flags.

```oak
requireOpt('entry', state.Entry, state.Abort, format)
```
