# JSON AST Output Target (`--ast`)

The `--ast` build target serializes the fully bundled and transformed AST as JSON. This is the raw data structure after all parsing, transpilation, analysis, module wrapping, and bundling passes have completed.

## Usage

```sh
oak build --entry main.oak --output bundle.json --ast
```

## Output Format

The output is a JSON object representing the root AST node (a `:block` node containing all bundled module expressions). Each node has at minimum a `type` field indicating its kind.

### Node Types

| Type | Fields |
|------|--------|
| `null` | `type`, `tok` |
| `empty` | `type`, `tok` |
| `string` | `type`, `val`, `tok` |
| `int` | `type`, `val`, `tok` |
| `float` | `type`, `val`, `tok` |
| `bool` | `type`, `val`, `tok` |
| `identifier` | `type`, `val`, `tok` |
| `atom` | `type`, `val`, `tok` |
| `list` | `type`, `elems`, `tok` |
| `object` | `type`, `entries`, `tok` |
| `unary` | `type`, `op`, `right`, `tok` |
| `binary` | `type`, `op`, `left`, `right`, `tok` |
| `assignment` | `type`, `local?`, `left`, `right`, `tok` |
| `propertyAccess` | `type`, `left`, `right`, `tok` |
| `ifExpr` | `type`, `cond`, `branches`, `tok` |
| `block` | `type`, `exprs`, `decls`, `tok` |
| `function` | `type`, `name`, `args`, `restArg`, `body`, `decls`, `tok` |
| `fnCall` | `type`, `function`, `args`, `restArg`, `tok` |
| `class` | `type`, `name`, `args`, `restArg`, `body`, `parents`, `staticExprs`, `decls`, `tok` |

## Use Cases

- External linting and formatting tools
- Language server protocol backends
- Custom codegen pipelines in any language
- Visual AST explorers and debuggers
- Building transpilers to targets not supported natively (C, Python, Rust, etc.)

The JSON AST is the **multiplier target**: once AST is serializable, anyone can write `oak-ast-to-X` tools in their language of choice without modifying the Magnolia core.
