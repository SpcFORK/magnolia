# Format Library (fmt)

## Overview

`libfmt` is the string formatting library for Oak, providing template-based string formatting with placeholder substitution.

## Import

```oak
fmt := import('fmt')
// or destructure specific functions
{ format: format, printf: printf } := import('fmt')
```

## Functions

### `format(raw, values...)`

Returns a formatted string where placeholders `{{N}}` are replaced by the Nth value from the arguments. Values can be referenced zero or more times.

**Placeholder Syntax:**
- `{{0}}` - First value (`values.0`)
- `{{1}}` - Second value (`values.1`)
- `{{N}}` - Nth value (`values.N`)
- `{{key}}` - Dictionary key (when first argument is an object)

**Features:**
- Whitespace inside `{{ }}` is ignored
- Values are converted to strings using `string()`
- Placeholders can be used multiple times
- Non-integer keys index into the first argument if it's an object

```oak
{ format: format } := import('fmt')

// Positional arguments
format('Hello, {{0}}!', 'World')
// => 'Hello, World!'

format('{{0}} + {{1}} = {{2}}', 5, 3, 8)
// => '5 + 3 = 8'

// Reusing placeholders
format('{{0}} {{0}} {{0}}!', 'Go')
// => 'Go Go Go!'

// Whitespace in placeholders is ignored
format('Result: {{ 0 }}', 42)
// => 'Result: 42'

// Object/dictionary mode
format('Name: {{name}}, Age: {{age}}', {
    name: 'Alice'
    age: 30
})
// => 'Name: Alice, Age: 30'

// Multiple values
format('{{0}} is {{1}} years old', 'Bob', 25)
// => 'Bob is 25 years old'

// Empty placeholder defaults to empty string on missing values
format('Value: {{5}}', 1, 2, 3)
// => 'Value: '
```

### `printf(raw, values...)`

Prints the result of `format(raw, values...)` to standard output  (using `println()`).

```oak
{ printf: printf } := import('fmt')

printf('Hello, {{0}}!', 'World')
// Prints: Hello, World!

printf('{{0}} + {{1}} = {{2}}', 5, 3, 8)
// Prints: 5 + 3 = 8

printf('User {{name}} logged in', { name: 'Alice' })
// Prints: User Alice logged in
```

## Examples

### Basic String Formatting

```oak
{ format: format } := import('fmt')

name := 'Alice'
age := 30

greeting := format('Hello, {{0}}! You are {{1}} years old.', name, age)
// => 'Hello, Alice! You are 30 years old.'
```

### Template-Based Messages

```oak
{ format: format } := import('fmt')

fn successMessage(action, item) {
    format('✓ Successfully {{0}} {{1}}', action, item)
}

fn errorMessage(action, item, reason) {
    format('✗ Failed to {{0}} {{1}}: {{2}}', action, item, reason)
}

println(successMessage('created', 'user account'))
// => ✓ Successfully created user account

println(errorMessage('delete', 'file.txt', 'Permission denied'))
// => ✗ Failed to delete file.txt: Permission denied
```

### Logging with Timestamps

```oak
{ printf: printf } := import('fmt')
datetime := import('datetime')

fn log(level, message) {
    timestamp := datetime.format(time())
    printf('[{{0}}] {{1}}: {{2}}', timestamp, level, message)
}

log('INFO', 'Application started')
log('WARN', 'Low disk space')
log('ERROR', 'Connection failed')
```

### Table Formatting

```oak
{ printf: printf } := import('fmt')
{ padEnd: padEnd, padStart: padStart } := import('str')

fn printRow(id, name, score) {
    printf(
        '| {{0}} | {{1}} | {{2}} |'
        padStart(string(id), 4, ' ')
        padEnd(name, 20, ' ')
        padStart(string(score), 6, ' ')
    )
}

println('| ID   | Name                 | Score  |')
println('|------|----------------------|--------|')
printRow(1, 'Alice', 95)
printRow(2, 'Bob', 87)
printRow(3, 'Charlie', 92)
```

### URL Building

```oak
{ format: format } := import('fmt')

fn buildURL(base, path, params) {
    url := format('{{0}}/{{1}}', base, path)
    if params != {} -> {
        queryString := // encode params
        url <- url + '?' + queryString
    }
    url
}

buildURL('https://api.example.com', 'users/123', {})
// => 'https://api.example.com/users/123'
```

### Error Messages with Context

```oak
{ format: format } := import('fmt')

fn parseError(line, col, message) {
    format('Parse error at line {{0}}, column {{1}}: {{2}}', line, col, message)
}

fn runtimeError(file, line, message) {
    format('Runtime error in {{0}}:{{1}}: {{2}}', file, line, message)
}

println(parseError(42, 15, 'Unexpected token'))
// => Parse error at line 42, column 15: Unexpected token

println(runtimeError('main.oak', 100, 'Division by zero'))
// => Runtime error in main.oak:100: Division by zero
```

### Object-Based Formatting

```oak
{ format: format } := import('fmt')

user := {
    name: 'Alice'
    email: 'alice@example.com'
    role: 'admin'
}

message := format(
    'User: {{name}} ({{email}})\nRole: {{role}}'
    user
)
println(message)
// User: Alice (alice@example.com)
// Role: admin
```

### Progress Indicators

```oak
{ printf: printf } := import('fmt')

fn showProgress(current, total) {
    percent := int(current * 100 / total)
    printf('Progress: {{0}}/{{1}} ({{2}}%)', current, total, percent)
}

showProgress(45, 100)
// Prints: Progress: 45/100 (45%)

showProgress(732, 1000)
// Prints: Progress: 732/1000 (73%)
```

### Reusable Templates

```oak
{ format: format } := import('fmt')

HTTPResponseTemplate := 'HTTP/1.1 {{0}} {{1}}\nContent-Type: {{2}}\n\n{{3}}'

fn httpResponse(status, statusText, contentType, body) {
    format(HTTPResponseTemplate, status, statusText, contentType, body)
}

response := httpResponse(200, 'OK', 'application/json', '{"success":true}')
println(response)
// HTTP/1.1 200 OK
// Content-Type: application/json
//
// {"success":true}
```

### Multiple Value References

```oak
{ format: format } := import('fmt')

// Using same value multiple times
repeated := format('{{0}}, {{0}}, and {{0}} again!', 'Hello')
// => 'Hello, Hello, and Hello again!'

// Complex expression
equation := format(
    '{{0}} × {{1}} = {{2}}, therefore {{2}} ÷ {{0}} = {{1}}'
    4, 5, 20
)
// => '4 × 5 = 20, therefore 20 ÷ 4 = 5'
```

## Implementation Notes

- Escaping is not supported—`{{` is always parsed as a placeholder start
- Invalid placeholder numbers access undefined values (result: empty string)
- Placeholders are zero-indexed (`{{0}}` is first argument)
- String conversion is automatic using Oak's `string()` builtin
- Whitespace trimming in placeholders: `{{ 0 }}` is equivalent to `{{0}}`
- For object mode, non-integer keys are looked up in `values.0`

## Edge Cases

```oak
{ format: format } := import('fmt')

// Empty placeholder
format('{{}}', 'test')
// => '' (empty placeholder = empty string)

// Out of range
format('{{5}}', 'a', 'b', 'c')
// => '' (no 6th argument)

// Non-existent object key
format('{{missing}}', { name: 'Alice' })
// => '' (key doesn't exist)

// Literal braces (not escaped)
format('Not a {{0}}placeholder', 'real')
// => 'Not a realplaceholder' (no space after placeholder)

// Single brace
format('Single { brace', 'test')
// => 'Single { brace' (not a placeholder)
```

## Performance Notes

- Simple string concatenation may be faster for very simple cases
- `format()` builds the result character-by-character with state management
- Useful when placeholders are reused or order varies
- `printf()` is just `println(format(...))` convenience wrapper

## Comparison with Alternatives

```oak
// Without format library (string concatenation)
'Hello, ' + name + '! You are ' + string(age) + ' years old.'

// With format library
format('Hello, {{0}}! You are {{1}} years old.', name, age)

// When you reuse values:
name + ' ' + name + ' ' + name        // Repetitive
format('{{0}} {{0}} {{0}}', name)     // Cleaner
```

## See Also

- `std.string()` - For type conversion
- `str` library - For string manipulation
- Oak's built-in `println()` - For basic output
