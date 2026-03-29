# Engine Switching — `bytecode()` and `interpreter()`

## Overview

Magnolia has two execution engines:

- **Tree-walking interpreter** — the default engine (`evalNodes`/`evalExpr`). Walks the AST directly. Highest throughput for callback-heavy code (map, filter, reduce) because builtins can call Oak functions with zero bridging overhead.
- **Bytecode VM** — activated with `--bytecode`. Compiles the AST to stack-based bytecode and runs it on a register-free VM. Fastest for tight loops, arithmetic, recursion, and object/string operations.

The `bytecode()` and `interpreter()` builtins let Oak programs switch engines mid-execution, routing each hot function to whichever engine is fastest for that workload.

```oak
// Run a tight arithmetic loop on the bytecode VM (even in interpreter mode)
result := with bytecode([200000]) fn(n) {
    fn sub(i, acc) if i {
        n -> acc
        _ -> sub(i + 1, acc + i * 3 - i / 2 + 7)
    }
    sub(0, 0)
}

// Run a callback-heavy map on the tree-walker (even in --bytecode mode)
mapped := with interpreter([myList]) fn(xs) {
    map(xs, fn(x) x * 2)
}
```

## Usage

Both builtins share the same calling convention:

```
bytecode(argList, function)
interpreter(argList, function)
```

The idiomatic form uses Oak's `with` sugar, which rewrites the trailing function as the last argument:

```oak
with bytecode([arg1, arg2]) fn(a, b) {
    // body executes on the bytecode VM
}

with interpreter([arg1, arg2]) fn(a, b) {
    // body executes on the tree-walking interpreter
}
```

### `bytecode(argList, fn)`

Compiles `fn` to bytecode and executes it on a pooled VM.

**Parameters:**
- `argList` — a list of values to pass as arguments (bound to the function's parameters in order)
- `fn` — the function to execute

**Returns:** The result of the function body.

**Behavior:**
- The function AST is compiled to bytecode on first call and **cached** per function definition. Repeated calls with the same function skip recompilation entirely.
- VM structs are **pooled** via `sync.Pool` to minimize allocation overhead.
- Outer-scope variables (imports, builtins, closures from enclosing scopes) are bridged into the VM's scope chain and resolved at runtime via `LOAD_UPVAL`.

```oak
std := import('std')
{ map: map } := std

// map is visible inside the bytecode-compiled function
with bytecode([1000]) fn(n) {
    xs := std.range(n)
    map(xs, fn(x) x * 2)
}
```

### `interpreter(argList, fn)`

Evaluates `fn` using the tree-walking interpreter with tail-call optimization.

**Parameters:**
- `argList` — a list of values to pass as arguments
- `fn` — the function to execute

**Returns:** The result of the function body.

**Behavior:**
- When called from `--bytecode` mode, bytecode closures with preserved AST are converted to `FnValue`s so the body is genuinely tree-walked instead of dispatched back to the VM.
- Tail-call optimization (TCO) is enabled — tail-recursive functions produce thunks that are iteratively unwound, avoiding stack overflow for deep recursion.
- The enclosing bytecode scope is bridged into a tree-walker scope so outer variables remain accessible.

```oak
{ sort: sort } := import('sort')

// sort's callback overhead is lower on the tree-walker
with interpreter([unsorted]) fn(xs) {
    sort(xs)
}
```

### Rest Arguments

Both builtins support functions with rest parameters:

```oak
with bytecode([1, 2, 3, 4, 5]) fn(first, rest...) {
    // first = 1, rest = [2, 3, 4, 5]
    first + reduce(rest, 0, fn(a, b) a + b)
}
```

## When to Use Which

Based on A/B benchmarking across both execution modes:

| Workload | Fastest Engine | Why |
|---|---|---|
| Tight arithmetic loops | `bytecode()` | VM dispatch is faster than AST walking for pure computation |
| Naive recursion (fib) | `bytecode()` | Less overhead per recursive call |
| Tail recursion | `bytecode()` | Negligible difference, but avoids bridge overhead |
| List construction (push) | `bytecode()` | Tight append loop benefits from VM speed |
| `map`, `filter`, `reduce` | `interpreter()` | Callback functions are called directly by Go builtins with zero bridging cost |
| String concat/slice | `bytecode()` | Iteration overhead dominates |
| Object build/read | `bytecode()` | Property access is faster under the VM |
| Closures & compose | `bytecode()` | Closure creation + invocation is cheaper on the VM |
| Sorting | `bytecode()` | Comparison-heavy workload benefits from VM dispatch |
| FizzBuzz (pattern matching) | `bytecode()` | Branch-heavy `if` chains are faster on the VM |
| Sieve of Eratosthenes | `interpreter()` | Heavy use of list mutation + callbacks favors the tree-walker |

**Rule of thumb:** Use `bytecode()` for self-contained computation (arithmetic, recursion, object manipulation). Use `interpreter()` when the function body calls higher-order builtins that take callbacks (`map`, `filter`, `reduce`, `sort` with complex comparators, `each`).

## Architecture

### Compilation Cache

Every `bytecode()` call compiles the function's AST to a `bytecodeChunk`. The chunk is cached in `engine.bytecodeCache` keyed by the `*fnNode` pointer, so the same function definition is only compiled once across all invocations:

```
bytecodeCache: map[*fnNode]*bytecodeChunk
```

The cache is protected by a `sync.RWMutex` for safe concurrent access.

### VM Pool

VM structs are recycled through a `sync.Pool`. Each VM has a pre-allocated 1024-slot stack and 64-capacity call frame array. `acquireVM` initializes a pooled VM for a given chunk; `releaseVM` clears references and returns it to the pool.

### Scope Bridging

The two engines use different scope representations:

- **Tree-walker:** `scope` — a linked list of `map[string]Value` frames
- **Bytecode VM:** `vmScope` — parallel `names[]`/`values[]` arrays with a parent pointer

Two bridge functions convert between them:

- `bridgeToVmScope(*scope) *vmScope` — flattens a tree-walker scope chain into a single `vmScope`. Used by `bytecode()` to make imports and outer variables available to the VM.
- `bridgeVmScope(*vmScope) scope` — copies a `vmScope` chain into a tree-walker scope. Used by `interpreter()` to make bytecode-captured variables available to the tree-walker.

### Upvalue Resolution

When `bytecode()` compiles a function, the compiler emits `LOAD_UPVAL` instructions for any identifier not declared as a local parameter. At runtime, `LOAD_UPVAL` walks the `vmScope` parent chain (seeded from the bridge) and falls back to the Context's top-level scope (builtins + imports).

## Examples

### Routing Hot Paths in a Benchmark

```oak
std := import('std')
{ range: range, map: map, reduce: reduce } := std
{ sort: sort } := import('sort')

// Arithmetic: bytecode is 3-5x faster
result := with bytecode([200000]) fn(n) {
    fn sub(i, acc) if i {
        n -> acc
        _ -> sub(i + 1, acc + i * 3 + 7)
    }
    sub(0, 0)
}

// Map over a list: interpreter is 4-5x faster
mapped := with interpreter([std.range(10000)]) fn(xs) {
    map(xs, fn(x) x * 2 + 1)
}

// Sort: bytecode wins in both modes
sorted := with bytecode([randomList]) fn(xs) {
    sort(xs)
}
```

### Nesting Engines

Engine calls can be nested arbitrarily:

```oak
with bytecode([data]) fn(xs) {
    // outer loop on bytecode VM
    fn process(items) {
        // switch to interpreter for the map
        with interpreter([items]) fn(ys) {
            map(ys, fn(y) y * 2)
        }
    }
    process(xs)
}
```

## Related Docs

- [Virtual-Bytecode.md](Virtual-Bytecode.md) — self-hosted bytecode VM in Oak
- [go.md](go.md) — Go runtime builtins (`go`, channels, syscalls)
- [cli.md](cli.md) — `--bytecode` flag documentation
- [build.md](build.md) — build system and compilation targets
