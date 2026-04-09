# data-toml — TOML Parser & Serializer

`import('data-toml')` provides TOML v1.0 parsing and serialization with support for basic types, tables, arrays of tables, inline tables, and dotted keys.

## Quick Start

```oak
toml := import('data-toml')

// Parse TOML
config := toml.parse('
[server]
host = "0.0.0.0"
port = 8080
debug = false

[database]
name = "mydb"
pool_size = 10

[[routes]]
path = "/api"
handler = "apiHandler"

[[routes]]
path = "/web"
handler = "webHandler"
')

// config.server.port => 8080
// config.routes.0.path => '/api'

// Serialize to TOML
text := toml.serialize({
    server: { host: '0.0.0.0', port: 8080 }
    database: { name: 'mydb' }
})
```

## API Reference

### `parse(text)`

Parses TOML text into an Oak object.

**Supported features:**
- Bare keys (`key = value`) and quoted keys (`"key" = value`)
- Dotted keys (`a.b.c = value`)
- Tables (`[table]`) and arrays of tables (`[[array]]`)
- Inline tables (`{ key = value, ... }`)
- Arrays (`[value1, value2]`)
- Strings (basic `"..."` with escapes, literal `'...'`)
- Integers, floats, booleans
- Comments (`#`)

### `serialize(obj)`

Converts an Oak object to TOML format.

## Notes

- Conforms to TOML v1.0 syntax for common use cases.
- Dotted keys are expanded into nested tables on parse.
- String escape sequences (`\n`, `\t`, `\\`, `\"`) are handled in quoted strings.
