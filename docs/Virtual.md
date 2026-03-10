# Virtual Library (Virtual)

## Overview

`libVirtual` is a self-hosting Oak interpreter written in Oak, capable of parsing and executing Oak source code at runtime. It provides a virtual machine for dynamic code evaluation and metaprogramming.

## Import

```oak
Virtual := import('Virtual')
{ createVM: createVM } := import('Virtual')
```

## Features

- **Self-hosting**: Oak interpreter in Oak
- **Dynamic evaluation**: Execute Oak code at runtime
- **AST interpretation**: Evaluates parsed syntax trees
- **Scope management**: Maintains execution scopes
- **Standard library access**: Integrates with Oak stdlib
- **Macro support**: Extensible with macros

## Functions

### `createVM()`

Creates a new virtual machine instance for code execution.

**Returns:** VM object with evaluation functions

```oak
{ createVM: createVM } := import('Virtual')

vm := createVM()

// VM provides:
// - vm.eval(code) - Evaluate Oak source code
// - vm.evalExpr(ast, scope) - Evaluate AST node
// - vm.globalScope - Global variable scope
```

## VM Object

### `evalExpr(node, scope)`

Evaluates an AST node in a given scope.

**Parameters:**
- `node` - AST node to evaluate
- `scope` - Variable scope object

**Returns:** Evaluation result

```oak
vm := createVM()

// Create AST node
ast := {
    type: :binary
    op: :plus
    left: { type: :int, val: 10 }
    right: { type: :int, val: 32 }
}

result := vm.evalExpr(ast, vm.globalScope)
println(result)  // => 42
```

### `globalScope`

The VM's global variable scope.

```oak
vm := createVM()

// Set global variable
vm.globalScope.myVar := 100

// Evaluate code using globals
result := vm.eval('myVar + 50')
println(result)  // => 150
```

## Supported Operations

### Literals

```oak
vm := createVM()

vm.evalExpr({ type: :null }, {})       // => ?
vm.evalExpr({ type: :empty }, {})      // => _
vm.evalExpr({ type: :int, val: 42 }, {})      // => 42
vm.evalExpr({ type: :float, val: 3.14 }, {})  // => 3.14
vm.evalExpr({ type: :string, val: 'hi' }, {}) // => 'hi'
vm.evalExpr({ type: :bool, val: true }, {})   // => true
vm.evalExpr({ type: :atom, val: 'ok' }, {})   // => :ok
```

### Binary Operations

Supported operators:
- Arithmetic: `+`, `-`, `*`, `/`, `%`, `**`
- Logical: `&`, `|`, `^`
- Comparison: `>`, `<`, `=`, `>=`, `<=`, `!=`

```oak
vm := createVM()

ast := {
    type: :binary
    op: :plus
    left: { type: :int, val: 5 }
    right: { type: :int, val: 3 }
}

vm.evalExpr(ast, {})  // => 8
```

### Unary Operations

Supported: `-` (negation), `!` (logical not), `~` (bitwise not)

```oak
vm := createVM()

ast := {
    type: :unary
    op: :minus
    right: { type: :int, val: 42 }
}

vm.evalExpr(ast, {})  // => -42
```

### Variables

```oak
vm := createVM()

scope := { x: 10, y: 20 }

ast := { type: :identifier, val: 'x' }
result := vm.evalExpr(ast, scope)
println(result)  // => 10
```

### Lists

```oak
vm := createVM()

ast := {
    type: :list
    elems: [
        { type: :int, val: 1 }
        { type: :int, val: 2 }
        { type: :int, val: 3 }
    ]
}

vm.evalExpr(ast, {})  // => [1, 2, 3]
```

### Objects

```oak
vm := createVM()

ast := {
    type: :object
    entries: [
        {
            key: { type: :identifier, val: 'name' }
            val: { type: :string, val: 'Oak' }
        }
        {
            key: { type: :string, val: 'version' }
            val: { type: :int, val: 1 }
        }
    ]
}

vm.evalExpr(ast, {})  // => { name: 'Oak', version: 1 }
```

### Function Calls

```oak
vm := createVM()

// Call built-in function
ast := {
    type: :fnCall
    function: { type: :identifier, val: 'len' }
    args: [{
        type: :list
        elems: [
            { type: :int, val: 1 }
            { type: :int, val: 2 }
        ]
    }]
}

vm.evalExpr(ast, vm.globalScope)  // => 2
```

## Usage Examples

### Simple Evaluation

```oak
{ createVM: createVM } := import('Virtual')
syntax := import('syntax')

vm := createVM()

code := '2 + 2 * 10'
ast := syntax.Tokenizer(code).parse()

result := vm.evalExpr(ast, {})
println(result)  // => 22
```

### Dynamic Code Execution

```oak
{ createVM: createVM } := import('Virtual')

vm := createVM()

// Set up environment
vm.globalScope.data := [1, 2, 3, 4, 5]
vm.globalScope.operation := 'sum'

// Execute user code
userCode := '
    if operation {
        "sum" -> reduce(data, 0, fn(acc, x) { acc + x })
        "product" -> reduce(data, 1, fn(acc, x) { acc * x })
        _ -> ?
    }
'

ast := parseCode(userCode)
result := vm.evalExpr(ast, vm.globalScope)
println(result)  // => 15
```

### REPL Implementation

```oak
{ createVM: createVM } := import('Virtual')
syntax := import('syntax')

vm := createVM()

fn repl {
    with std.loop() fn(again) {
        print('> ')
        input := readLine()
        
        if input = 'exit' -> {
            println('Goodbye!')
        } else {
            ast := syntax.Tokenizer(input).parse()
            result := vm.evalExpr(ast, vm.globalScope)
            println(result)
            again()
        }
    }
}

repl()
```

### Sandboxed Execution

```oak
{ createVM: createVM } := import('Virtual')

fn evaluateSafely(code, timeout) {
    vm := createVM()
    
    // Limited scope (no file I/O, etc.)
    safeScope := {
        // Math operations only
        add: fn(a, b) { a + b }
        mul: fn(a, b) { a * b }
        max: std.max
        min: std.min
    }
    
    ast := parseCode(code)
    
    // Evaluate with limited scope
    result := vm.evalExpr(ast, safeScope)
    result
}

result := evaluateSafely('add(10, mul(5, 3))', 1000)
println(result)  // => 25
```

### Macro System

```oak
{ createVM: createVM } := import('Virtual')

vm := createVM()

// Define macro
vm.macros.unless := fn(cond, body) {
    // Transform: unless(x) { y } → if !x { y }
    {
        type: :ifExpr
        cond: { type: :unary, op: :exclam, right: cond }
        branches: [{ target: { type: :bool, val: true }, body: body }]
    }
}

// Macros expand before evaluation
```

### Custom Operators

```oak
{ createVM: createVM } := import('Virtual')

vm := createVM()

// Override binary operator
originalPlus := vm.evalBinaryOp
vm.evalBinaryOp := fn(op, left, right) if op {
    :plus -> if type(left) = :string & type(right) = :string {
        true -> left + ' ' + right  // Add space between strings
        _ -> originalPlus(op, left, right)
    }
    _ -> originalPlus(op, left, right)
}

// Now string + string adds spaces
ast := parseCode('"Hello" + "World"')
result := vm.evalExpr(ast, {})
println(result)  // => 'Hello World'
```

## VM State

### Call Stack

The VM maintains a call stack for tracking function execution:

```oak
vm.callStack := []  // Stack of calling contexts
```

### Standard Libraries

Access to Oak's built-in libraries:

```oak
vm.stdlibs := ___stdlibs()  // Get all standard libraries
```

## Implementation Details

### Binary Operators

```oak
:plus    // + (numbers, strings, lists)
:minus   // - (numbers)
:times   // * (numbers)
:divide  // / (numbers)
:modulus // % (numbers)
:power   // ** (numbers)
:and     // & (booleans, bitwise)
:or      // | (booleans, bitwise)
:xor     // ^ (bitwise)
:greater // > (comparison)
:less    // < (comparison)
:eq      // = (equality)
:geq     // >= (comparison)
:leq     // <= (comparison)
:neq     // != (inequality)
```

### Type Handling

- Integers: Exact arithmetic
- Floats: IEEE 754 floating point
- Strings: Concatenation with `+`
- Lists: Append with `+`
- Objects: Property access

## Limitations

- Performance overhead (interpreted, not compiled)
- No TCO (tail call optimization) support
- Limited error messages
- No debugger integration
- Stack overflow risk on deep recursion
- No module hot-reloading
- Import only works for built-ins

## Use Cases

- **REPLs**: Interactive Oak shells
- **Scripting**: Dynamic code execution
- **Plugins**: User-provided extensions
- **Configuration**: Executable config files
- **Testing**: Evaluate test code dynamically
- **Metaprogramming**: Code generation at runtime

## Performance Considerations

- Interpretation is slower than native execution
- AST traversal has overhead
- Scope lookups are dictionary accesses
- No JIT compilation
- Best for small, dynamic code snippets

## Security Warnings

⚠️ **Do not evaluate untrusted code**:
- No sandboxing enforced
- Full access to VM scope
- Can call any available functions
- No resource limits
- Potential for infinite loops

## See Also

- [syntax.md](syntax.md) - Oak tokenizer/parser
- [ast-analyze.md](ast-analyze.md) - AST analysis
- [codegen-common.md](codegen-common.md) - AST rendering
- Oak language specification
