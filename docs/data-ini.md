# data-ini — INI Parser & Serializer

`import('data-ini')` provides INI file parsing and serialization supporting sections, key-value pairs, comments (`#` and `;`), and quoted values.

## Quick Start

```oak
ini := import('data-ini')

// Parse INI text
config := ini.parse('
[database]
host = localhost
port = 5432
debug = true

[app]
name = "My App"
')
// config.database.host => 'localhost'
// config.database.port => 5432
// config.app.name => 'My App'

// Serialize to INI
text := ini.serialize({
    database: { host: 'localhost', port: 5432 }
    app: { name: 'My App' }
})
```

## API Reference

### `parse(text)`

Parses INI text into an Oak object. Sections become nested objects. Supports:
- `[section]` headers
- `key = value` pairs
- Comments: lines starting with `#` or `;`
- Quoted values with `"` or `'`
- Automatic type coercion: booleans, integers, floats, atoms, null

### `serialize(obj)`

Converts an Oak object to INI format. Global (non-object) properties are written first, followed by `[section]` blocks.

## Notes

- Values are automatically typed on parse: `true`/`false` → booleans, numeric strings → numbers.
- Empty values and `null` are supported.
- Nested sections (sub-sections) are not supported — only one level of `[section]` grouping.
