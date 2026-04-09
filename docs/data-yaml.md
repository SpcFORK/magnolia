# data-yaml — YAML Parser & Serializer

`import('data-yaml')` provides YAML parsing and serialization supporting scalars, mappings, sequences, multi-line strings (literal `|` and folded `>`), and flow syntax.

## Quick Start

```oak
yaml := import('data-yaml')

// Parse YAML
config := yaml.parse('
server:
  host: "0.0.0.0"
  port: 8080
  debug: false
tags:
  - web
  - api
description: |
  This is a
  multi-line string
')
// config.server.port => 8080
// config.tags.0 => 'web'

// Serialize to YAML
text := yaml.serialize({
    name: 'Alice'
    age: 30
    hobbies: ['reading', 'coding']
})
```

## API Reference

### `parse(text)`

Parses YAML text into an Oak value.

**Supported features:**
- Scalars: null (`null`, `~`), booleans (`true`, `false`), integers, floats, strings
- Quoted strings (`"..."`, `'...'`) with escape sequences
- Block strings: literal (`|` preserves newlines), folded (`>` folds to spaces)
- Sequences: block (`- item`) and flow (`[a, b, c]`)
- Mappings: block (`key: value`) and flow (`{key: value}`)
- Comments (`#`)
- Nested structures via indentation (2 spaces per level)

### `serialize(val)`

Converts an Oak value to YAML text.

## Notes

- Values that look like keywords (`null`, `true`, `false`, `~`) are automatically quoted during serialization when used as strings.
- Indentation is 2 spaces per nesting level.
- Anchors/aliases (`&`, `*`) and tags (`!!type`) are not supported.
- Multi-line strings use `|` for literal blocks and `>` for folded blocks.
