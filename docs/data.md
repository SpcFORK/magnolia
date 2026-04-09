# data — Unified Data Format Facade

`import('data')` provides a single import point for all data format libraries.

## Quick Start

```oak
Data := import('data')

// CSV
rows := Data.CSV.parse('a,b\n1,2')
text := Data.CSV.serialize([[1, 2], [3, 4]])

// JSON
obj := Data.JSON.parse('{"key": "value"}')

// YAML
obj := Data.YAML.parse('name: Alice\nage: 30')

// TOML
obj := Data.TOML.parse('[server]\nport = 8080')

// XML
nodes := Data.XML.parse('<root><item>hi</item></root>')

// INI
obj := Data.INI.parse('[db]\nhost = localhost')

// Markdown
html := Data.Markdown.parse('# Hello')

// MessagePack
raw := Data.MessagePack.serialize({ key: 'value' })
```

## Available Modules

| Property | Underlying Module | Formats |
|----------|-------------------|---------|
| `CSV` | `data-csv` | RFC 4180 CSV |
| `JSON` | `json` | JSON |
| `TOML` | `data-toml` | TOML v1.0 |
| `XML` | `data-xml` | XML |
| `INI` | `data-ini` | INI |
| `YAML` | `data-yaml` | YAML |
| `Markdown` | `md` | Markdown |
| `MessagePack` | `msgpack` | MessagePack |

## See Also

- [data-csv](data-csv.md)
- [data-ini](data-ini.md)
- [data-toml](data-toml.md)
- [data-xml](data-xml.md)
- [data-yaml](data-yaml.md)
- [json](json.md)
- [msgpack](msgpack.md)
- [md](md.md)
