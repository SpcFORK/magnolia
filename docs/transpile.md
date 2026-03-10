# Transpile Middleware

The transpile middleware provides a plugin architecture for AST transformations during the build process.

## Overview

The transpile middleware is automatically integrated into both `build.oak` and `pack.oak`. It sits in the compilation pipeline between AST parsing and code generation, allowing you to transform the AST before it's compiled.

## Architecture

```
Source Code → Parse → Wrap Block → [Transpile] → Analyze → Wrap Module → Bundle → Codegen
```

The transpile step applies registered transformers to the AST in sequence.

## Usage

### Using Built-in Transpilers

```oak
build := import('build')

build.run({
    entry: 'src/main.oak'
    output: 'build/bundle.oak'
    transpilers: [
        build.transpile.optimizeConstants
        build.transpile.removeDebugCalls
    ]
})
```

### Creating Custom Transpilers

A transpiler is simply a function that takes an AST node and returns a transformed AST node:

```oak
transpile := import('transpile')

// Example: Convert all "foo" identifiers to "bar"
fn renameFoo(node) if node.type {
    :identifier -> if node.val {
        'foo' -> {
            type: :identifier
            tok: node.tok
            val: 'bar'
        }
        _ -> node
    }
    _ -> node
}

// Register the transpiler
transpile.registerTranspiler(renameFoo)
```

### Using the walkNode Helper

For more complex transformations that need to traverse the entire AST:

```oak
transpile := import('transpile')

// Example: Count all function calls
fnCallCount := { count: 0 }

myTranspiler := transpile.createTranspiler(fn(node) if node.type {
    :fnCall -> {
        fnCallCount.count := fnCallCount.count + 1
        node
    }
    _ -> node
})

transpile.registerTranspiler(myTranspiler)
```

## Configuration

### Enable/Disable Transpilation

```oak
build := import('build')

build.run({
    entry: 'main.oak'
    output: 'bundle.oak'
    transpile?: false  // Disable transpilation
})
```

### Verbose Logging

```oak
build := import('build')

build.run({
    entry: 'main.oak'
    output: 'bundle.oak'
    transpileVerbose?: true  // Enable detailed transpile logs
})
```

## API Reference

### transpile.configure(config)

Configure the transpile middleware.

**Parameters:**
- `config.enabled?` - Enable/disable transpilation (default: true)
- `config.verbose?` - Enable verbose logging (default: false)
- `config.log` - Custom logging function (default: printf)

### transpile.registerTranspiler(fn)

Register a transpiler function to be applied during compilation.

**Parameters:**
- `fn` - A function that takes an AST node and returns a transformed node

### transpile.clearTranspilers()

Remove all registered transpilers.

### transpile.transpileNode(node)

Apply all registered transpilers to an AST node. This is called automatically by the build process.

**Parameters:**
- `node` - An AST node

**Returns:**
- Transformed AST node

### transpile.walkNode(node, visitor)

Recursively walk an AST and apply a visitor function to each node.

**Parameters:**
- `node` - The AST node to walk
- `visitor` - Function called for each node

**Returns:**
- Transformed AST node

### transpile.createTranspiler(visitor)

Create a transpiler from a visitor function.

**Parameters:**
- `visitor` - Function called for each node in the AST

**Returns:**
- A transpiler function

## Built-in Transpilers

### optimizeConstants

Performs constant folding for simple arithmetic expressions.

```oak
// Before: 2 + 3
// After: 5
```

### removeDebugCalls

Removes all calls to the `debug()` function.

```oak
// Before: debug('checking value')
// After: (removed)
```

## Examples

### Dead Code Elimination

```oak
transpile := import('transpile')

fn eliminateDeadCode(node) if node.type {
    :ifExpr -> if node.cond.type {
        :atom -> if node.cond.val {
            :true -> if len(node.branches) > 0 {
                true -> node.branches.(0).body
                _ -> node
            }
            :false -> if len(node.branches) > 1 {
                true -> node.branches.(1).body
                _ -> {
                    type: :atom
                    tok: node.tok
                    val: :null
                }
            }
            _ -> node
        }
        _ -> node
    }
    _ -> node
}

transpile.registerTranspiler(eliminateDeadCode)
```

### String Interpolation Macro

```oak
transpile := import('transpile')
{
    replace: replace
} := import('str')

// Transform template strings: `"Hello ${name}"` → `"Hello " + name`
fn expandTemplates(node) if node.type {
    :string -> if node.val |> contains?('${') {
        true -> {
            // Parse and transform template string
            // (simplified example)
            type: :binary
            tok: node.tok
            op: :plus
            left: { type: :string, val: 'Hello ', tok: node.tok }
            right: { type: :identifier, val: 'name', tok: node.tok }
        }
        _ -> node
    }
    _ -> node
}

transpile.registerTranspiler(expandTemplates)
```

## Integration with Build Tools

The transpile middleware is automatically available in:

- **build.oak** - Direct AST transformation during bundling
- **pack.oak** - Inherits transpilation through build.oak
- **cmd/build.oak** - CLI builds with transpile support
- **cmd/pack.oak** - CLI packaging with transpile support

## Best Practices

1. **Keep transpilers pure** - Avoid side effects; always return a new or modified node
2. **Preserve token information** - Maintain `tok` fields for accurate error reporting
3. **Test thoroughly** - Transpilers can introduce subtle bugs
4. **Order matters** - Transpilers are applied sequentially in registration order
5. **Use walkNode for deep transforms** - When you need to modify nested structures
6. **Clear state** - `clearTranspilers()` is called automatically between builds

## Performance Considerations

- Transpilers are applied to every module in the build
- Complex transformations may slow down build times
- Use `transpile?: false` in config to disable when not needed
- Verbose logging (`transpileVerbose?: true`) adds overhead

## Troubleshooting

### Transpiler not being applied

1. Check that transpilers are registered before `build.run()` or `pack.run()`
2. Verify `transpile?: false` is not set in config
3. Enable verbose logging to see which transpilers are applied

### Build errors after adding transpiler

1. Ensure your transpiler returns valid AST nodes
2. Check that token information is preserved
3. Use `transpile.walkNode` to recursively transform nested nodes
4. Test your transpiler on simple cases first
