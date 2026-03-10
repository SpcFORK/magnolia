# JSON Library (json)

## Overview

`libjson` implements a JSON parser and serializer for converting between Oak values and JSON strings. It handles all standard JSON types and provides robust error handling for invalid JSON.

## Import

```oak
json := import('json')
// or destructure specific functions
{ serialize: serialize, parse: parse } := import('json')
```

## Functions

### `serialize(value)`

Converts an Oak value to its JSON string representation.

**Supported Types:**

- `null`, `empty`, `function` → `"null"`
- `string` → JSON string (with escaping)
- `atom` → JSON string (converts to string)
- `int`, `float`, `bool` → JSON number/boolean
- `list` → JSON array `[...]`
- `object` → JSON object `{...}`

**JSON Escaping:**

The following characters are escaped in strings:
- `\t` → `\\t` (tab)
- `\n` → `\\n` (newline)
- `\r` → `\\r` (carriage return)
- `\f` → `\\f` (form feed)
- `"` → `\\"` (quote)
- `\\` → `\\\\` (backslash)

```oak
{ serialize: serialize } := import('json')

// Simple values
serialize(42) // => "42"
serialize(3.14) // => "3.14"
serialize(true) // => "true"
serialize(?) // => "null"
serialize('hello') // => '"hello"'

// Composite values
serialize([1, 2, 3]) // => "[1,2,3]"
serialize({name: 'Alice', age: 30})
// => '{"name":"Alice","age":30}'

// Nested structures
serialize({
    users: ['Alice', 'Bob']
    count: 2
    active: true
})
// => '{"users":["Alice","Bob"],"count":2,"active":true}'

// Atoms convert to strings
serialize(:success) // => '"success"'

// Functions become null
serialize(fn(x) x + 1) // => "null"

// Special characters are escaped
serialize('Hello\nWorld') // => '"Hello\\nWorld"'
serialize('Say "Hi"') // => '"Say \\"Hi\\""'
```

### `parse(jsonString)`

Parses a JSON string and returns its Oak representation. Returns `:error` if the JSON is invalid.

**JSON to Oak Type Mapping:**

- JSON `null` → Oak `?`
- JSON string → Oak string
- JSON number → Oak `int` or `float`
- JSON boolean → Oak `true` or `false`
- JSON array → Oak list `[]`
- JSON object → Oak object `{}`

**Error Handling:**

Returns `:error` for:
- Invalid JSON syntax
- Unterminated strings
- Malformed numbers
- Missing commas or brackets
- Unexpected characters

```oak
{ parse: parse } := import('json')

// Simple values
parse('42') // => 42
parse('3.14') // => 3.14
parse('true') // => true
parse('null') // => ?
parse('"hello"') // => 'hello'

// Lists
parse('[1, 2, 3]') // => [1, 2, 3]
parse('["a", "b", "c"]') // => ['a', 'b', 'c']

// Objects
parse('{"name": "Alice", "age": 30}')
// => {name: 'Alice', age: 30}

// Nested structures
parse('{
    "users": ["Alice", "Bob"],
    "count": 2,
    "active": true
}')
// => {users: ['Alice', 'Bob'], count: 2, active: true}

// Whitespace is ignored
parse('  {  "a"  :  1  }  ') // => {a: 1}

// Escaped characters
parse('"Hello\\nWorld"') // => 'Hello\nWorld'
parse('"Say \\"Hi\\""') // => 'Say "Hi"'

// Error cases
parse('invalid') // => :error
parse('{broken') // => :error
parse('[1, 2,]') // => :error (trailing comma)
parse('{"key": }') // => :error (missing value)
```

## Examples

### API Data Handling

```oak
json := import('json')

// Serialize request payload
requestData := {
    username: 'alice'
    email: 'alice@example.com'
    preferences: {
        theme: 'dark'
        notifications: true
    }
}

payload := serialize(requestData)
// => '{"username":"alice","email":"alice@example.com",...}'

// Parse API response
response := '{"status":"success","userId":123,"token":"abc123"}'
data := parse(response)

if data != :error -> {
    println('User ID: ' + string(data.userId))
    println('Token: ' + data.token)
}
```

### Configuration Files

```oak
{ serialize: serialize, parse: parse } := import('json')
{ writeFile: writeFile, readFile: readFile } := import('fs')

// Save configuration
config := {
    server: {
        port: 8080
        host: 'localhost'
    }
    database: {
        url: 'postgres://localhost/mydb'
        maxConnections: 10
    }
}

writeFile('config.json', serialize(config))

// Load configuration
fn loadConfig {
    content := readFile('config.json')
    if content != ? -> parse(content)
}

loadedConfig := loadConfig()
if loadedConfig != :error -> {
    println('Server port: ' + string(loadedConfig.server.port))
}
```

### Data Validation

```oak
{ parse: parse } := import('json')

fn validateJSON(jsonStr) if parsed := parse(jsonStr) {
    :error -> {
        println('Invalid JSON!')
        false
    }
    _ -> {
        println('Valid JSON')
        true
    }
}

validateJSON('{"valid": true}') // => Valid JSON, returns true
validateJSON('{invalid}') // => Invalid JSON!, returns false
```

### Round-Trip Conversion

```oak
{ serialize: serialize, parse: parse } := import('json')

original := {
    name: 'Product'
    price: 29.99
    tags: ['new', 'featured']
    inStock: true
}

// Serialize and parse back
jsonStr := serialize(original)
restored := parse(jsonStr)

// Values are preserved
restored.name // => 'Product'
restored.price // => 29.99
restored.tags // => ['new', 'featured']
restored.inStock // => true
```

### Working with Arrays

```oak
{ serialize: serialize, parse: parse } := import('json')

// Serialize list of objects
users := [
    {id: 1, name: 'Alice'}
    {id: 2, name: 'Bob'}
    {id: 3, name: 'Charlie'}
]

jsonArray := serialize(users)
// => '[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"},...]'

// Parse and process
parsedUsers := parse(jsonArray)
with map(parsedUsers) fn(user) user.name
// => ['Alice', 'Bob', 'Charlie']
```

## Implementation Notes

- Serialization produces compact JSON (no extra whitespace)
- Parsing is whitespace-tolerant
- Numbers are automatically typed as `int` or `float` based on content
- Object key order is not guaranteed to be preserved
- Atoms are serialized as strings (no way to distinguish on parse)
- Functions, `empty`, and `null` all serialize to `"null"`
- The parser uses a stateful reader for efficient character-by-character processing
- Error recovery is not attempted—first error causes `:error` return

## Limitations

- No pretty-printing option (always compact)
- No streaming parser (entire string must be in memory)
- No custom serialization hooks
- Cannot distinguish between atoms and strings after round-trip
- No support for JSON5 extensions (comments, trailing commas, etc.)
- Large numbers may lose precision due to floating-point representation

## Related Functions

Uses these imports from other Oak libraries:
- `std.default()`, `std.slice()`, `std.map()` - Utilities
- `str.space?()`, `str.join()` - String operations

## Error Handling

Always check for `:error` when parsing untrusted JSON:

```oak
{ parse: parse } := import('json')

fn safeParseJSON(str) if result := parse(str) {
    :error -> {
        println('JSON parse error')
        ? // Return null or default value
    }
    _ -> result
}
```
