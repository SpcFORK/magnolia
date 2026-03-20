# Syntax Macros (syntax-macros)

## Overview

`syntax-macros` provides a hygienic AST macro expansion pass for Oak. A macro is a function that receives raw AST nodes as arguments and returns a transformed AST node. Macros are expanded recursively before evaluation or code generation.

## Import

```oak
macros := import('syntax-macros')
{ Macro: Macro, macro?: macro?, expandMacros: expandMacros } := import('syntax-macros')
```

## Types

### `Macro(expand)`

Wraps `expand` into a macro descriptor object. `expand` must be a function with the signature `fn(args, callNode, macros)` that returns an AST node.

- `args` — list of AST nodes passed to the macro call site.
- `callNode` — the full `:fnCall` AST node at the call site.
- `macros` — the current macro map (allows macros to call other macros).

```oak
// A macro that swaps two arguments
SwapArgs := Macro(fn(args, _, _) {
    type: :fnCall
    tok: args.(0).tok
    function: args.(0).function
    args: [args.(1), args.(0)]
    restArg: ?
})
```

## Functions

### `macro?(x)`

Returns `true` when `x` is a descriptor created by `Macro(...)`.

```oak
macro?(Macro(fn(args, _, _) args.(0)))  // => true
macro?('not a macro')                   // => false
```

### `expandMacros(ast, macros?)`

Walks `ast` (a list of top-level AST nodes or a single node) and replaces every `:fnCall` whose function identifier matches a key in the `macros` map with the node returned by the macro's `expand` function. Expansion is applied recursively so macros may produce nodes that contain further macro calls.

`macros` defaults to `{}` when omitted (no-op pass).

Returns the fully expanded AST (a list when given a list, a node when given a node).

```oak
macroMap := {
    'double': Macro(fn(args, call, _) {
        // replace double(x) with x + x
        {
            type: :binary
            tok: call.tok
            op: :plus
            left: args.(0)
            right: args.(0)
        }
    })
}

syntax := import('syntax')
ast := syntax.parse('double(5)')
expanded := expandMacros(ast, macroMap)
// expanded represents `5 + 5`
```

**Non-call nodes** (`:list`, `:object`, `:function`, `:assignment`, `:propertyAccess`, `:unary`, `:binary`, `:ifExpr`, `:ifBranch`) are recursed into so macros at any nesting depth are expanded.
