# Bundle AST Library (bundle-ast)

## Overview

`libbundle-ast` combines multiple Oak modules into a single bundled AST with proper module system initialization for standalone execution.

## Import

```oak
bundleAST := import('bundle-ast')
{ wrapBundle: wrapBundle } := import('bundle-ast')
```

## Functions

### `wrapBundle(modules, entryModuleName)`

Creates a bundled AST from multiple parsed modules.

**Parameters:**
- `modules` - List of `[modulePath, moduleAST]` pairs
- `entryModuleName` - Name of the entry point module to execute

**Returns:** Block AST node containing all modules and initialization code

```oak
{ wrapBundle: wrapBundle } := import('bundle-ast')

modules := [
    ['main.oak', mainAST]
    ['lib/utils.oak', utilsAST]
    ['lib/helpers.oak', helpersAST]
]

bundled := wrapBundle(modules, 'main.oak')

// bundled is a block containing:
// 1. All module wrapper functions
// 2. Module initialization table
// 3. Entry point invocation
```

## Bundle Structure

The generated bundle has this structure:

```oak
{
    type: :block
    exprs: [
        module1WrapperFunction  // fn { ... }
        module2WrapperFunction  // fn { ... }
        ...
        moduleNWrapperFunction  // fn { ... }
        entryPointCall         // __oak_module_import('main.oak')
    ]
    decls: []
}
```

## Module Wrapping

Each module is wrapped in a function via `__oak_modularize`:

```javascript
// Generated structure
__oak_modularize('module.oak', function() {
    // Original module code
    return { exports }
})
```

## Usage Examples

### Simple Bundle

```oak
{ wrapBundle: wrapBundle } := import('bundle-ast')
syntax := import('syntax')

// Parse modules
main := syntax.parse(readFile('main.oak'))
utils := syntax.parse(readFile('utils.oak'))

modules := [
    ['main.oak', main]
    ['utils.oak', utils]
]

bundled := wrapBundle(modules, 'main.oak')
```

### Build Pipeline Integration

```oak
{ wrapBundle: wrapBundle } := import('bundle-ast')

fn compileProject(entryPoint) {
    // 1. Resolve all imports
    modules := resolveModules(entryPoint)
    
    // 2. Parse each module
    parsedModules := modules |> map(fn([path, code]) {
        [path, parseOakCode(code)]
    })
    
    // 3. Analyze each module
    analyzedModules := parsedModules |> map(fn([path, ast]) {
        [path, analyzeNode(ast, true, ?)]
    })
    
    // 4. Bundle into single AST
    bundled := wrapBundle(analyzedModules, entryPoint)
    
    // 5. Generate code
    codegen(bundled)
}
```

### Custom Module Loader

```oak
{ wrapBundle: wrapBundle } := import('bundle-ast')

fn bundleWithVirtualFS(modules, entry, vfsFiles) {
    // Create module with embedded VFS
    vfsModule := createVFSModule(vfsFiles)
    
    // Add VFS to modules
    allModules := [['__vfs', vfsModule]] |> append(modules)
    
    // Bundle with VFS included
    wrapBundle(allModules, entry)
}
```

## Module Initialization Table

The bundle creates a table mapping module names to their wrapper functions:

```javascript
{
    'main.oak': __oak_modularize('main.oak', mainFunc)
    'lib/utils.oak': __oak_modularize('lib/utils.oak', utilsFunc)
    'lib/helpers.oak': __oak_modularize('lib/helpers.oak', helpersFunc)
}
```

## Entry Point Execution

The bundle ends with a call to load the entry module:

```javascript
__oak_module_import('main.oak')
```

This triggers:
1. Module table lookup
2. Module function execution
3. Caching of module exports
4. Return of cached exports

## Implementation Details

### Module Name Resolution

Module paths are used as-is:

```oak
'main.oak'          // Entry point
'lib/utils.oak'     // Relative to entry
'/abs/path.oak'     // Absolute path
```

### Modularization Function

Each module wrapped with `__oak_modularize`:

```javascript
__oak_modularize(moduleName, moduleFunction)
```

Returns a callable that:
- Executes module function once
- Caches module exports
- Returns cached exports on subsequent calls

### Token Generation

Uses synthetic tokens for generated nodes:

```oak
{ pos: [0, 1, 0], type: :stringLiteral, val: moduleName }
{ pos: [0, 1, 0], type: :identifier, val: '__oak_modularize' }
```

## Bundle Execution Flow

1. **Load bundle**: Execute bundled code
2. **Register modules**: Store module functions in table
3. **Import entry**: Call `__oak_module_import(entryModule)`
4. **Resolve imports**: Load dependencies on demand
5. **Cache exports**: Store module return values
6. **Return result**: Entry module's exports

## Integration Points

### Used by Build System

```oak
// build.oak uses bundle-ast
modules := collectModules(entry)
bundled := wrapBundle(modules, entry)
output := codegen(bundled)
writeFile(outputPath, output)
```

### Used by Pack Tool

```oak
// pack.oak uses bundle-ast via build
bundled := buildAndBundle(entry, includes)
executable := createExecutable(bundled, vfsFiles)
```

## Module Format

Each module AST should be a transformed module wrapper:

```oak
{
    type: :function
    name: ''  // Anonymous
    args: []
    body: {
        type: :block
        exprs: [/* module code */]
        decls: [/* module declarations */]
    }
}
```

## Limitations

- No dynamic imports (all modules known at bundle time)
- Module paths must be string literals
- Circular dependencies not specially handled
- No code splitting
- No lazy loading
- Bundle size includes all modules (no tree shaking here)

## Optimization Opportunities

Bundle-ast doesn't optimize, but bundled AST can be:
- Dead code eliminated
- Minified
- Constant folded
- Inlined

## Use Cases

- **Application bundling**: Combine app modules
- **Library distribution**: Single-file libraries
- **Code deployment**: Deploy as one unit
- **Static analysis**: Analyze entire program
- **Tree shaking**: Enable cross-module optimization

## See Also

- [build.md](build.md) - Build system
- [ast-transform.md](ast-transform.md) - Module wrapping
- [ast-analyze.md](ast-analyze.md) - Semantic analysis
- [pack.md](pack.md) - Create executables
