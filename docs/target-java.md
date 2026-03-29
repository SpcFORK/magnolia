# Java Target (`--java`)

Transpiles Oak/Magnolia source to Java 17+.

## Usage

```sh
magnolia build --entry src/main.oak --output OakBundle.java --java
```

## Output

A single `OakBundle.java` file containing:

- **Runtime helpers** — `__oak_eq`, `__oak_push`, `__oak_acc`, `__as_oak_string`, module system (`__oak_modularize`, `__oak_module_import`)
- **Transpiled AST** — Oak expressions mapped to Java equivalents inside `public static void main`

## Design Decisions

- **All values are `Object`**: Oak is dynamically typed; Java uses `Object` with runtime casts.
- **Functions as `Function<Object[], Object>`**: Oak functions become Java lambdas accepting an `Object[]` and returning `Object`.
- **Maps for objects**: Oak objects become `HashMap<String, Object>`.
- **Lists**: Oak lists become `ArrayList<Object>`.
- **Pattern matching via if-chains**: Oak `if` expressions with pattern branches translate to `if`/`else if` chains comparing with `__oak_eq`.
- **Identifier mangling**: Java reserved words are prefixed with `__oak_java_` (e.g., `class` → `__oak_java_class`). `?` → `_qm`, `!` → `_exclam`.

## Limitations

- No type inference — everything is `Object` with casts.
- Arithmetic requires `Number` casts at runtime.
- Closures work via Java lambda capture but mutable state requires workarounds.
- String concatenation uses `__as_oak_string` helper for type coercion.
- Large Oak programs may hit Java's method size limit; future work may split into multiple methods.

## Requirements

- Java 17+ (uses `Map.of`, enhanced switch, text blocks-compatible style)
- No external dependencies
