# AST Transform Library (ast-transform)

## Overview

`libast-transform` provides utilities for transforming and wrapping Abstract Syntax Tree (AST) nodes during Oak compilation and bundling.

## Import

```oak
transform := import('ast-transform')
{
    wrapBlock: wrapBlock
    wrapModule: wrapModule
    formatIdentForWeb: formatIdentForWeb
    formatIdentForOak: formatIdentForOak
} := import('ast-transform')
```

## Functions

### `wrapBlock(nodes)`

Creates a block AST node containing multiple expressions.

**Parameters:**
- `nodes` - List of AST nodes

**Returns:** Block AST node

```oak
{ wrapBlock: wrapBlock } := import('ast-transform')

// Combine multiple AST nodes into one block
block := wrapBlock([node1, node2, node3])

// Equivalent to: { node1; node2; node3 }
```

### `wrapModule(block)`

Transforms an AST block into a module wrapper function for bundling.

**Parameters:**
- `block` - Block AST node

**Returns:** Function AST node that exports module contents

```oak
{ wrapModule: wrapModule } := import('ast-transform')

moduleFunc := wrapModule(blockNode)

// Creates: fn { /* block contents */ }
```

### `formatIdentForWeb(name, key?)`

Formats an identifier for JavaScript output, avoiding reserved words.

**Parameters:**
- `name` - Original identifier name
- `key` - Optional unique suffix

**Returns:** Safe JavaScript identifier

```oak
{ formatIdentForWeb: formatIdentForWeb } := import('ast-transform')

formatIdentForWeb('_')          // => '__oak_empty_ident'
formatIdentForWeb('class')      // => '__oak_js_class'
formatIdentForWeb('function')   // => '__oak_js_function'
formatIdentForWeb('myVar')      // => 'myVar'
formatIdentForWeb('_', '1')     // => '__oak_empty_ident1'
```

### `formatIdentForOak(name)`

Formats an identifier for Oak output, handling Oak-specific reserved words.

**Parameters:**
- `name` - Original identifier name

**Returns:** Safe Oak identifier

```oak
{ formatIdentForOak: formatIdentForOak } := import('ast-transform')

formatIdentForOak('import')     // => '__oak_module_import'
formatIdentForOak('myVar')      // => 'myVar'
```

## Reserved Word Handling

### JavaScript Reserved Words

Protected by `formatIdentForWeb`:

```
await, break, case, catch, class, const, continue
debugger, default, delete, do, else, enum, export
extends, false, finally, for, function, if, in
instanceof, new, null, return, super, switch, this
throw, true, try, typeof, var, void, while, with
yield
```

### Special Cases

```oak
// Underscore (_) gets unique suffix to avoid shadowing
formatIdentForWeb('_')       // => '__oak_empty_ident'
formatIdentForWeb('_', '0')  // => '__oak_empty_ident0'

// JS keywords get prefixed
formatIdentForWeb('class')   // => '__oak_js_class'
formatIdentForWeb('return')  // => '__oak_js_return'

// Normal names unchanged
formatIdentForWeb('myVar')   // => 'myVar'
```

## AST Node Structure

### Block Node

```oak
{
    type: :block
    tok: { pos: [0, 1, 1], type: :leftBrace, val: ? }
    exprs: [expr1, expr2, expr3]  // List of expressions
    decls: ['var1', 'var2']       // Declared variables
}
```

### Function/Module Node

```oak
{
    type: :function
    tok: { pos: [0, 1, 1], type: :fnKeyword, val: ? }
    name: ''                      // Empty for anonymous
    args: []                      // Parameters
    restArg: ?                    // Rest parameter or ?
    body: blockNode               // Function body
    decls: ['localVar']           // Local declarations
}
```

## Usage Examples

### Combining Multiple Nodes

```oak
{ wrapBlock: wrapBlock } := import('ast-transform')

statements := [
    { type: :assignment, /* ... */ }
    { type: :fnCall, /* ... */ }
    { type: :identifier, /* ... */ }
]

combinedBlock := wrapBlock(statements)
```

### Creating Module Wrapper

```oak
{ wrapBlock: wrapBlock, wrapModule: wrapModule } := import('ast-transform')

// Wrap module exports in function
exports := wrapBlock([
    exportNode1
    exportNode2
])

moduleFunc := wrapModule(exports)

// Result is callable module initializer
```

### Safe Identifier Generation

```oak
{ formatIdentForWeb: formatIdentForWeb } := import('ast-transform')

// Generate safe variable names for JS output
varNames := ['class', 'return', 'myVar', '_']

safeNames := varNames |> map(fn(name, i) {
    formatIdentForWeb(name, string(i))
})

// => ['__oak_js_class', '__oak_js_return', 'myVar', '__oak_empty_ident0']
```

### Build Pipeline Integration

```oak
{ wrapBlock: wrapBlock, formatIdentForWeb: formatIdentForWeb } := import('ast-transform')

fn processModule(moduleAST, moduleName) {
    // Rename identifiers to avoid conflicts
    processedAST := walkNode(moduleAST, fn(node) if node.type = :identifier {
        true -> {
            node.val := formatIdentForWeb(node.val, moduleName)
            node
        }
        _ -> node
    })
    
    // Wrap in module function
    wrapModule(processedAST)
}
```

### Bundler Implementation

```oak
{
    wrapBlock: wrapBlock
    wrapModule: wrapModule
    formatIdentForWeb: formatIdentForWeb
} := import('ast-transform')

fn bundleModules(modules) {
    // Transform each module
    moduleWrappers := modules |> map(fn(mod) {
        // Wrap module contents
        moduleFunc := wrapModule(mod.ast)
        
        // Safe module name
        safeName := formatIdentForWeb('mod_' + mod.name)
        
        {
            name: safeName
            func: moduleFunc
        }
    })
    
    // Combine all modules into one block
    wrapBlock(moduleWrappers |> map(fn(m) m.func))
}
```

## Token Structure

Token objects in AST nodes:

```oak
{
    pos: [offset, line, column]  // Source position
    type: :tokenType             // Token type atom
    val: value                   // Token value or ?
}
```

## Helper Patterns

### Creating Empty Tokens

```oak
emptyToken := { pos: [0, 1, 1], type: :leftBrace, val: ? }
```

### Creating Synthetic Nodes

```oak
{ wrapBlock: wrapBlock } := import('ast-transform')

// Create node from scratch
syntheticNode := {
    type: :identifier
    tok: { pos: [0, 1, 1], type: :identifier, val: 'x' }
    val: 'x'
}

block := wrapBlock([syntheticNode])
```

## Integration with Other Libraries

### Used By Build System

```oak
// build.oak uses ast-transform
buildAST := parseSource(code)
wrappedModule := wrapModule(buildAST)
jsCode := codegen(wrappedModule)
```

### Used By Bundler

```oak
// bundle-ast.oak uses ast-transform
modules := collectModules()
bundled := wrapBlock(modules |> map(wrapModule))
```

### Used By Transpiler

```oak
// transpile.oak uses formatIdent functions
transpileNode := fn(node) {
    if node.type = :identifier -> {
        node.val := formatIdentForWeb(node.val)
    }
    node
}
```

## Implementation Notes

- All functions return new AST nodes (immutable)
- Token positions use synthetic values ([0, 1, 1])
- Block wrapping preserves declaration information
- Module wrapping creates anonymous functions
- Identifier formatting is deterministic

## Limitations

- Synthetic tokens have no real source position
- No validation of reserved words beyond predefined list
- Module wrapping doesn't handle circular dependencies
- No automatic import/export transformation
- formatIdent functions are target-specific

## Use Cases

- **Bundling**: Combine multiple modules
- **Code generation**: Create synthetic AST nodes
- **Transpilation**: Safe identifier renaming
- **Optimization**: Restructure code blocks
- **Module systems**: Wrap exports in functions

## See Also

- [ast-analyze.md](ast-analyze.md) - Semantic AST analysis
- [bundle-ast.md](bundle-ast.md) - AST bundling
- [transpile.md](transpile.md) - AST transformation middleware
- [build.md](build.md) - Build system
