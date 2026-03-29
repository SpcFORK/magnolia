# Lua Output Target (`--lua`)

The `--lua` build target transpiles Oak programs to Lua 5.4. This enables embedding Magnolia logic in Lua-hosted environments like game engines, Redis, Nginx/OpenResty, and Neovim.

## Usage

```sh
oak build --entry main.oak --output bundle.lua --lua
```

## Output Format

The output is a self-contained Lua 5.4 script that includes:

1. **Runtime preamble**: Oak compatibility functions (`__oak_eq`, `__oak_push`, `__oak_acc`, etc.)
2. **Module system**: `__oak_modularize` / `__oak_module_import` registry
3. **Transpiled code**: All module code converted to Lua syntax

## Semantic Mapping

| Oak Feature | Lua Equivalent |
|-------------|----------------|
| Functions | `function(...) end` |
| Closures | Lua closures (upvalues are mutable by default) |
| Objects | Lua tables with string keys |
| Lists | Lua tables (note: 0-indexed in Oak, 1-indexed in Lua) |
| Atoms | String constants |
| Pattern matching | `if/elseif/else` chains |
| Pipe `\|>` | Nested function calls |
| `null` | `nil` |
| `true`/`false` | `true`/`false` |

## Runtime Helpers

- `__oak_eq(a, b)` — Deep structural equality comparison
- `__oak_push(tbl, val)` — Append to table
- `__oak_acc(obj, key)` — Safe property access
- `__as_oak_string(v)` — String coercion for `+` operator
- `__oak_and(a, b)` / `__oak_or(a, b)` — Short-circuit logical operators
- `__oak_xor(a, b)` — XOR for booleans and integers

## Identifier Mangling

Lua reserved words are prefixed with `__oak_lua_`:
`and`, `break`, `do`, `else`, `elseif`, `end`, `false`, `for`, `function`,
`goto`, `if`, `in`, `local`, `nil`, `not`, `or`, `repeat`, `return`, `then`,
`true`, `until`, `while`

## Known Limitations

- **Strings**: Lua strings are immutable. Oak's mutable string indexing is not fully supported.
- **Array indexing**: Lua arrays are 1-indexed. External Lua code interfacing with Oak lists should account for this.
- **Async**: Oak's `go()`, channels, and `wait()` are not supported in the Lua target.
- **Operator polymorphism**: Bitwise operations on strings are not supported.

## Use Cases

- Game engines (LÖVE, Defold, Corona)
- Redis stored procedures
- Nginx/OpenResty scripting
- Neovim plugin development
- IoT/embedded scripting environments
