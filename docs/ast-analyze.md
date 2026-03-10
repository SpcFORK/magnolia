# AST Analyze Library (ast-analyze)

## Overview

`libast-analyze` performs semantic analysis on Abstract Syntax Trees (ASTs), annotating nodes with information needed for code generation and optimization.

## Import

```oak
analyze := import('ast-analyze')
{ analyzeNode: analyzeNode } := import('ast-analyze')
```

## Features

- **Semantic analysis**: Analyzes AST for variable declarations, scopes, and references
- **Tail call optimization**: Detects and marks tail-recursive functions
- **Scope tracking**: Tracks variable declarations and closures
- **Context annotation**: Adds metadata for code generation

## Functions

### `analyzeNode(node, webTarget?, enclosingContext?)`

Performs semantic analysis on an AST node and returns a transformed tree with annotations.

**Parameters:**
- `node` - Root AST node to analyze
- `webTarget?` - Boolean, enables web-specific optimizations (tail call trampolining)
- `enclosingContext?` - Optional parent context for nested analysis

**Returns:** Transformed AST with semantic annotations

```oak
{ analyzeNode: analyzeNode } := import('ast-analyze')

ast := parseOakCode(sourceCode)
analyzedAST := analyzeNode(ast, true, ?)  // Analyze for web target

// AST now has:
// - .decls arrays on blocks/functions
// - .recurred? flag on recursive functions
// - Trampoline wrappers for tail recursion
```

## Analysis Features

### Variable Declaration Tracking

Tracks which variables are declared in each scope:

```oak
// Input code:
fn example {
    x := 10
    y := 20
    x + y
}

// Analyzed node.decls: ['x', 'y']
```

### Scope Analysis

Differentiates between local declarations and arguments:

```oak
// Input:
fn add(a, b) {
    result := a + b
    result
}

// Analysis:
// - args: ['a', 'b']
// - decls: ['result']  (excludes arguments)
```

### Tail Call Optimization (Web Target)

For recursive functions on web target, wraps with trampoline:

```oak
// Original:
fn factorial(n, acc) if n {
    0 -> acc
    _ -> factorial(n - 1, acc * n)  // Tail call
}

// Analyzed (web?: true):
fn factorial(n, acc) {
    __oak_trampolined_factorial := fn(n, acc) if n {
        0 -> acc
        _ -> factorial(n - 1, acc * n)
    }
    __oak_resolve_trampoline(__oak_trampolined_factorial, n, acc)
}
```

### Closure Detection

Identifies variables captured in closures:

```oak
// Input:
fn outer(x) {
    fn inner(y) {
        x + y  // x is captured from outer scope
    }
    inner
}

// Analysis tracks x as closure variable
```

## Analyzed Node Structure

### Block with Declarations

```oak
{
    type: :block
    exprs: [...]
    decls: ['var1', 'var2']  // Variables declared in this block
}
```

### Function with Scope Info

```oak
{
    type: :function
    name: 'myFunc'
    args: ['a', 'b']
    decls: ['local1', 'local2']  // Local vars (excluding args)
    body: analyzedBlock
    recurred?: true  // Set if function is recursive
}
```

## Usage Examples

### Basic Analysis

```oak
{ analyzeNode: analyzeNode } := import('ast-analyze')
syntax := import('syntax')

code := 'fn add(x, y) { result := x + y; result }'
ast := syntax.parse(code)

analyzed := analyzeNode(ast, false, ?)
// analyzed contains semantic information
```

### Web Target Analysis

```oak
{ analyzeNode: analyzeNode } := import('ast-analyze')

ast := parseCode(sourceCode)

// Enable web optimizations (tail call trampolining)
webAST := analyzeNode(ast, true, ?)

// Recursive functions now have trampoline wrappers
```

### Build Pipeline Integration

```oak
{ analyzeNode: analyzeNode } := import('ast-analyze')

fn compileModule(sourceCode, targetWeb?) {
    ast := parseSource(sourceCode)
    analyzed := analyzeNode(ast, targetWeb?, ?)
    codegen(analyzed)  // Generate code from analyzed AST
}
```

## Recursive Function Handling

### Without Trampolining (Native Target)

```oak
analyzeNode(ast, false, ?)
// Recursive calls remain as-is
// Stack-based recursion
```

### With Trampolining (Web Target)

```oak
analyzeNode(ast, true, ?)
// Recursive tail calls wrapped in trampoline
// Prevents stack overflow in JavaScript
```

## Context Tracking

Analysis tracks enclosing context for nested scopes:

```oak
{
    enclosingFn: functionNode  // Nearest named function
    enclosingFnLit: functionNode  // Nearest function literal
    decls: {}  // Declarations in current scope
    args: {}  // Arguments in current scope
}
```

## Implementation Notes

- Does not mutate original AST (returns new tree)
- Recursively analyzes all sub-expressions
- Tracks declarations separately from arguments
- Tail call optimization only for web target
- Marks recursive functions with `recurred?` flag

## Optimization Passes

1. **Scope analysis**: Track variable declarations
2. **Recursion detection**: Identify recursive calls
3. **Tail call detection**: Find tail-recursive patterns
4. **Trampoline insertion**: Wrap tail calls (web only)
5. **Closure analysis**: Identify captured variables

## Limitations

- Tail call optimization only for directly recursive functions
- No mutual recursion optimization
- Trampoline overhead on web targets
- No dead code detection
- No constant propagation
- Closure analysis is basic (doesn't optimize captures)

## Use Cases

- **Build systems**: Prepare AST for code generation
- **Optimizers**: Enable tail call optimization
- **Transpilers**: Add semantic information to AST
- **Linters**: Check for undefined variables
- **Documentation**: Extract function signatures

## Trampoline Pattern

Web targets use trampoline to avoid stack overflow:

```javascript
// Generated JavaScript
function factorial(n, acc) {
    let __oak_trampolined_factorial = function(n, acc) {
        if (n === 0) return acc;
        return factorial(n - 1, acc * n);
    };
    return __oak_resolve_trampoline(__oak_trampolined_factorial, n, acc);
}
```

## See Also

- [ast-transform.md](ast-transform.md) - AST transformation utilities
- [transpile.md](transpile.md) - AST transformation middleware
- [build.md](build.md) - Build system integration
- [syntax.md](syntax.md) - Oak parser
