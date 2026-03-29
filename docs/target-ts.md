# TypeScript Output Target (`--ts`)

The `--ts` build target compiles Oak programs to TypeScript. It uses the JavaScript renderer as its base and adds type stub declarations.

## Usage

```sh
oak build --entry main.oak --output bundle.ts --ts
```

## Output Format

The output is a valid TypeScript file that includes:

1. **Type declarations**: Basic Oak type stubs (`OakVal`, `OakFn`, `OakList`, `OakObject`)
2. **JS runtime**: The full Oak JavaScript runtime preamble
3. **Bundled code**: All module code transpiled to JS (same as `--web`)

The file includes `// @ts-nocheck` to allow gradual typing. The type stubs provide a starting point for consuming Oak modules from TypeScript projects.

## Type Strategy

The current approach is **shallow typing**:

- All runtime values are typed as `OakVal` (alias for `any`)
- Module exports are typed as `OakObject`
- Function parameters default to `OakVal`

This provides IDE autocompletion for module shapes while avoiding unsound type inference.

## Use Cases

- Publishing Oak libraries as typed npm packages
- Consuming Oak-generated code from TypeScript projects
- IDE autocompletion and hover documentation
- Type-safe module boundaries between Oak and TypeScript code

## Identifier Mangling

TypeScript output uses the same identifier mangling rules as JavaScript:
- ECMAScript reserved words → `__oak_js_<word>`
- `?` in names → `__oak_qm`
- `!` in names → `__oak_exclam`
- `import` → `__oak_module_import`
