# data-csv — CSV Parser & Serializer

`import('data-csv')` provides RFC 4180 CSV parsing and serialization with support for quoted fields, embedded newlines, and configurable delimiters.

## Quick Start

```oak
csv := import('data-csv')

// Parse CSV text to rows
rows := csv.parse('name,age\nAlice,30\nBob,25')
// => [['name', 'age'], ['Alice', '30'], ['Bob', '25']]

// Parse with headers → list of objects
records := csv.parseWithHeader('name,age\nAlice,30\nBob,25')
// => [{name: 'Alice', age: '30'}, {name: 'Bob', age: '25'}]

// Serialize rows to CSV
text := csv.serialize([['a', 'b'], [1, 2]])

// Serialize objects with headers
text := csv.serializeWithHeader(
    [{name: 'Alice', age: 30}]
    ['name', 'age']
)
```

## API Reference

### `parse(text, opts?)`

Parses a CSV string into a list of rows (each row is a list of string fields).

**Options:**
- `delimiter` — field separator (default: `','`)

### `parseWithHeader(text, opts?)`

Parses CSV using the first row as column headers. Returns a list of objects keyed by header names.

### `serialize(rows, opts?)`

Converts a list of rows (lists of values) to a CSV string.

**Options:**
- `delimiter` — field separator (default: `','`)
- `newline` — line ending (default: `'\r\n'`)

### `serializeWithHeader(records, headers, opts?)`

Serializes a list of objects using the specified header row.

## Notes

- Fields containing the delimiter, quotes, or newlines are automatically quoted.
- Embedded double quotes are escaped as `""` per RFC 4180.
