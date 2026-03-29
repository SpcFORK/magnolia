# Documentation Output Target (`--doc`)

The `--doc` build target generates Markdown API documentation from the bundled AST. It extracts module names, function signatures, and constant declarations.

## Usage

```sh
oak build --entry lib/std.oak --output std-api.md --doc
```

## Output Format

The generated Markdown includes:

- **Module headers**: Each bundled module gets an `## Module:` section
- **Function signatures**: With parameter names and rest-arg indicators
- **Constant declarations**: Listed as bullet points

## Limitations

- Comments are not currently preserved in the AST. Function descriptions rely on naming conventions.
- Type information is not inferred (Oak is dynamically typed).
- Only top-level declarations within modules are documented.

## Example Output

```markdown
# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `std.oak`

### `map(list, fn)`

### `filter(list, fn)`

- `PI` — constant
```
