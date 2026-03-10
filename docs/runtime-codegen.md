# Runtime Codegen Library (runtime-codegen)

## Overview

`libruntime-codegen` provides runtime code templates for bundled Oak modules, including the module system implementation for both Oak native and JavaScript targets.

## Import

```oak
runtimeCodegen := import('runtime-codegen')
{ OakNativeRuntime: OakNativeRuntime, OakJSRuntime: OakJSRuntime } := import('runtime-codegen')
```

## Constants

### `OakNativeRuntime`

Oak source code for the module system runtime (native Oak target).

**Provides:**
- `__Oak_Modules` - Module registry
- `__Oak_Import_Aliases` - Import alias mapping
- `__oak_modularize(name, module)` - Register module
- `__oak_module_import(name)` - Load module

```oak
{ OakNativeRuntime: OakNativeRuntime } := import('runtime-codegen')

// Used in bundled Oak code
bundleSource := OakNativeRuntime + '\n' + moduleCode
```

### `OakJSRuntime`

JavaScript code for the module system runtime (web target).

**Provides:**
- Module system (same functions as Oak runtime)
- Language primitives:
  - `__oak_eq(a, b)` - Deep equality comparison
  - `__oak_resolve_trampoline(fn, ...args)` - Tail call optimization
  - `__oak_trampoline(fn, ...args)` - Trampoline wrapper
- Helper functions for Oak semantics in JavaScript

```oak
{ OakJSRuntime: OakJSRuntime } := import('runtime-codegen')

// Used in web builds
jsBundleSource := OakJSRuntime + '\n' + transpiled JSCode
```

## Oak Native Runtime

### Module Registry

```oak
__Oak_Modules := {}              // Stores all modules
__Oak_Import_Aliases := ?        // Import path aliases
```

### Module Registration

```oak
fn __oak_modularize(name, module) {
    __Oak_Modules.(name) := module
}

// Usage in bundle:
__oak_modularize('utils.oak', fn { /* module code */ })
```

### Module Import

```oak
fn __oak_module_import(name) if ___runtime_lib?(name) {
    true -> import(name)  // Built-in library
    _ -> if type(module := __Oak_Modules.(name)) {
        :null -> if module := __Oak_Modules.(__Oak_Import_Aliases.(name)) {
            ? -> import(name)
            _ -> {
                mod := module()
                __Oak_Modules.(name) := mod  // Cache
                mod
            }
        }
        :function -> {
            m := module()
            __Oak_Modules.(name) := m  // Cache
            m
        }
        _ -> module  // Already cached
    }
}
```

## JavaScript Runtime

### Module System

```javascript
const __Oak_Modules = {};
let __Oak_Import_Aliases;

function __oak_modularize(name, fn) {
    __Oak_Modules[name] = fn;
}

function __oak_module_import(name) {
    if (typeof __Oak_Modules[name] === 'object') {
        return __Oak_Modules[name];  // Cached
    }
    
    const module = __Oak_Modules[name] || 
                   __Oak_Modules[__Oak_Import_Aliases[name]];
    
    if (module) {
        __Oak_Modules[name] = {};  // Break circular imports
        return __Oak_Modules[name] = module();
    } else {
        throw new Error(`Could not import Oak module "${name}" at runtime`);
    }
}
```

### Deep Equality

```javascript
function __oak_eq(a, b) {
    // Handle Oak's empty value (_)
    if (a === __Oak_Empty || b === __Oak_Empty) return true;
    
    // Null handling
    if (a == null && b == null) return true;
    if (a == null || b == null) return false;
    
    // Primitives
    if (typeof a === 'boolean' || typeof a === 'number' ||
        typeof a === 'symbol' || typeof a === 'function') {
        return a === b;
    }
    
    // Strings (Oak strings need special handling)
    a = __as_oak_string(a);
    b = __as_oak_string(b);
    if (typeof a !== typeof b) return false;
    if (__is_oak_string(a) && __is_oak_string(b)) {
        return a.valueOf() === b.valueOf();
    }
    
    // Collections (recursive comparison)
    if (len(a) !== len(b)) return false;
    for (const key of keys(a)) {
        if (!__oak_eq(a[key], b[key])) return false;
    }
    return true;
}
```

### Tail Call Optimization (Trampolining)

```javascript
// Resolve trampoline to final value
function __oak_resolve_trampoline(fn, ...args) {
    let result = fn(...args);
    while (typeof result === 'object' && 
           typeof result[Symbol.iterator] === 'function') {
        result = result[Symbol.iterator]().next().value(...args);
    }
    return result;
}

// Create trampoline for tail recursion
function __oak_trampoline(fn, ...args) {
    return function*() {
        yield function() {
            return fn(...args);
        };
    };
}
```

## Usage Examples

### Oak Native Bundle

```oak
{ OakNativeRuntime: OakNativeRuntime } := import('runtime-codegen')

bundleCode := OakNativeRuntime + '

__oak_modularize("main.oak", fn {
    println("Hello from main!")
})

__oak_modularize("utils.oak", fn {
    { add: fn(a, b) { a + b } }
})

__oak_module_import("main.oak")
'

writeFile('bundle.oak', bundleCode)
```

### JavaScript Web Bundle

```oak
{ OakJSRuntime: OakJSRuntime } := import('runtime-codegen')

jsBundleCode := OakJSRuntime + '

__oak_modularize("app", function() {
    console.log("App loaded");
    return { version: "1.0" };
});

__oak_module_import("app");
'

writeFile('bundle.js', jsBundleCode)
```

### Build System Integration

```oak
{ OakNativeRuntime: OakNativeRuntime, OakJSRuntime: OakJSRuntime } := import('runtime-codegen')

fn generateBundle(modules, target) {
    runtime := if target {
        :web -> OakJSRuntime
        :native -> OakNativeRuntime
        _ -> OakNativeRuntime
    }
    
    moduleCode := modules |> map(fn([name, code]) {
        '__oak_modularize("' + name + '", function() {\n' +
        code + '\n});'
    }) |> join('\n\n')
    
    entryCall := '__oak_module_import("' + modules.0.0 + '");'
    
    runtime + '\n\n' + moduleCode + '\n\n' + entryCall
}
```

### Circular Import Handling

```javascript
// JavaScript runtime breaks circular imports
__Oak_Modules[name] = {};  // Placeholder
return __Oak_Modules[name] = module();  // Replace with actual
```

Example:

```javascript
// a.js imports b.js, b.js imports a.js

__oak_modularize("a", function() {
    const b = __oak_module_import("b");  // Gets placeholder {}
    return { name: "A", b: b };
});

__oak_modularize("b", function() {
    const a = __oak_module_import("a");  // Placeholder prevents infinite loop
    return { name: "B", a: a };
});
```

### Module Caching

Both runtimes cache module results:

```oak
// First call: executes module function
utils := __oak_module_import('utils')

// Second call: returns cached result
utils2 := __oak_module_import('utils')

// utils === utils2 (same object)
```

## Runtime Functions

### `___runtime_lib?(name)`

Built-in function (Oak native only) that checks if a name is a runtime library.

```oak
___runtime_lib?('std')   // => true
___runtime_lib?('http')  // => true
___runtime_lib?('myapp') // => false
```

## Implementation Notes

- Module functions are called only once
- Results are cached in `__Oak_Modules`
- Circular imports return placeholder object
- JavaScript runtime includes polyfills for Oak semantics
- Trampoline optimization prevents stack overflow
- Deep equality handles Oak's value semantics

## Limitations

- Circular imports may have incomplete exports
- Dynamic imports not supported
- Module names must be string literals
- No lazy loading
- No module hot-reloading
- Trampoline has overhead compared to native tail calls

## Use Cases

- **Bundling**: Embed in bundled code
- **Code generation**: Template for module system
- **Web deployment**: JavaScript runtime for browsers
- **Testing**: Standalone module execution

## Security Considerations

- No sandboxing of module code
- All modules share global scope
- No module permission system
- Errors in one module can crash all

## See Also

- [build.md](build.md) - Build system
- [bundle-ast.md](bundle-ast.md) - AST bundling
- [pack.md](pack.md) - Create executables
- Oak module system documentation
