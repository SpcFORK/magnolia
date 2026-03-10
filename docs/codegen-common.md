# Codegen Common Library (codegen-common)

## Overview

`libcodegen-common` provides shared utilities for code generation, including AST rendering and dependency tracing used across different compilation targets.

## Import

```oak
codegen := import('codegen-common')
{ renderNode: renderNode, traceDependencies: traceDependencies } := import('codegen-common')
```

## Functions

### `renderNode(node)`

Converts an AST node back into Oak source code text.

**Parameters:**
- `node` - AST node to render

**Returns:** Oak source code string

```oak
{ renderNode: renderNode } := import('codegen-common')

// Render various node types
renderNode({ type: :null })
// => '?'

renderNode({ type: :number, val: 42 })
// => '42'

renderNode({ type: :string, val: 'hello' })
// => '\'hello\''

renderNode({ type: :atom, val: 'success' })
// => ':success'

renderNode({ type: :bool, val: true })
// => 'true'
```

### `traceDependencies(nodes)`

Analyzes an AST to find all module dependencies (imports).

**Parameters:**
- `nodes` - AST node or list of nodes

**Returns:** List of module names (strings) that are imported

```oak
{ traceDependencies: traceDependencies } := import('codegen-common')

ast := parseCode('
    http := import("http")
    json := import("json")
    utils := import("lib/utils")
')

deps := traceDependencies(ast)
// => ['http', 'json', 'lib/utils']
```

## Supported Node Types

### Literals

```oak
renderNode({ type: :null })          // => '?'
renderNode({ type: :empty })         // => '_'
renderNode({ type: :number, val: 3.14 })  // => '3.14'
renderNode({ type: :string, val: 'hi' })  // => '\'hi\''
renderNode({ type: :atom, val: 'ok' })    // => ':ok'
renderNode({ type: :bool, val: false })   // => 'false'
```

### Identifiers

```oak
renderNode({ type: :identifier, val: 'myVar' })
// => 'myVar'
```

### Lists

```oak
renderNode({
    type: :list
    elems: [
        { type: :number, val: 1 }
        { type: :number, val: 2 }
    ]
})
// => '[1, 2]'
```

### Objects

```oak
renderNode({
    type: :object
    entries: [
        { key: { type: :identifier, val: 'x' }, val: { type: :number, val: 10 } }
        { key: { type: :string, val: 'y' }, val: { type: :number, val: 20 } }
    ]
})
// => '{x: 10, \'y\': 20}'
```

### Function Calls

```oak
renderNode({
    type: :fnCall
    function: { type: :identifier, val: 'println' }
    args: [{ type: :string, val: 'hello' }]
    restArg: ?
})
// => 'println(\'hello\')'
```

### Functions

```oak
renderNode({
    type: :function
    name: 'add'
    args: ['a', 'b']
    restArg: ?
    body: { type: :identifier, val: 'result' }
})
// => 'fn add(a, b) result'

// Anonymous function
renderNode({
    type: :function
    name: ''
    args: ['x']
    body: { /* ... */ }
})
// => 'fn(x) { ... }'
```

## Usage Examples

### AST to Source Code

```oak
{ renderNode: renderNode } := import('codegen-common')

ast := parseOakCode('fn add(a, b) { a + b }')
source := renderNode(ast)
println(source)
// => 'fn add(a, b) { a + b }'
```

### Dependency Analysis

```oak
{ traceDependencies: traceDependencies } := import('codegen-common')

code := '
    std := import("std")
    http := import("http")
    
    fn main {
        json := import("json")
        utils := import("lib/utils")
    }
'

ast := parseCode(code)
dependencies := traceDependencies(ast)

println('Dependencies found:')
each(dependencies, println)
// => 'std'
// => 'http'
// => 'json'
// => 'lib/utils'
```

### Code Transformation Pipeline

```oak
{ renderNode: renderNode, traceDependencies: traceDependencies } := import('codegen-common')

fn processModule(sourceCode) {
    // 1. Parse to AST
    ast := parseCode(sourceCode)
    
    // 2. Find dependencies
    deps := traceDependencies(ast)
    println('Found ' + string(len(deps)) + ' dependencies')
    
    // 3. Transform AST
    transformedAST := transformNode(ast)
    
    // 4. Render back to source
    output := renderNode(transformedAST)
    
    output
}
```

### Pretty Printing AST

```oak
{ renderNode: renderNode } := import('codegen-common')

fn prettyPrint(node, indent) {
    if node.type = :block -> {
        result := '{\n'
        each(node.exprs, fn(expr) {
            result <- result + repeat('  ', indent + 1)
            result <- result + renderNode(expr) + '\n'
        })
        result <- result + repeat('  ', indent) + '}'
        result
    } else {
        renderNode(node)
    }
}
```

### Dependency Graph Builder

```oak
{ traceDependencies: traceDependencies } := import('codegen-common')

fn buildDependencyGraph(entryPoint) {
    graph := {}
    visited := {}
    
    fn visit(modulePath) {
        if visited.(modulePath) -> ?
        visited.(modulePath) := true
        
        code := readFile(modulePath)
        ast := parseCode(code)
        deps := traceDependencies(ast)
        
        graph.(modulePath) := deps
        
        each(deps, fn(dep) {
            depPath := resolveImport(dep, modulePath)
            visit(depPath)
        })
    }
    
    visit(entryPoint)
    graph
}
```

### Source Map Generation

```oak
{ renderNode: renderNode } := import('codegen-common')

fn generateWithSourceMap(ast, originalSource) {
    output := renderNode(ast)
    
    // Build mapping between original and generated
    mapping := {
        generated: output
        original: originalSource
        ast: ast
    }
    
    mapping
}
```

## String Escaping

The `renderNode` function automatically escapes strings:

```oak
{ renderNode: renderNode } := import('codegen-common')

renderNode({ type: :string, val: 'it\'s' })
// => '\'it\\\'s\''

renderNode({ type: :string, val: 'line1\nline2' })
// => '\'line1\nline2\''
```

## Implementation Notes

- Recursively traverses AST to find `import()` calls
- Only detects static imports (not dynamic strings)
- String rendering escapes single quotes
- Object keys can be identifiers or strings
- Rest arguments rendered with `...` suffix
- Unknown node types render as `'(unknown node)'`

## Limitations

- No source formatting/indentation
- No comment preservation
- No whitespace control
- Dynamic imports not detected (e.g., `import(variable)`)
- Rest arguments use simplified syntax
- No validation of generated code
- String escaping is basic (only handles quotes)

## Use Cases

- **Code generation**: Convert AST back to source
- **Bundling**: Extract dependencies for module resolution
- **Documentation**: Extract imports for docs
- **Analysis**: Build dependency graphs
- **Debugging**: Print AST as readable code
- **Transpilation**: Transform and re-emit code

## Dependency Detection

Finds imports in:
- Top-level statements
- Function bodies
- Conditional branches
- Nested blocks
- Assignment expressions

Does not find:
- Dynamic imports: `import(varName)`
- String concatenation: `import('li' + 'b')`
- Computed imports

## See Also

- [ast-transform.md](ast-transform.md) - AST transformation
- [ast-analyze.md](ast-analyze.md) - Semantic analysis
- [build.md](build.md) - Build system
- [syntax.md](syntax.md) - Oak parser
