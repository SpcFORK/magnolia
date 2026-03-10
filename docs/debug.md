# Debug Library (debug)

## Overview

`libdebug` provides utilities for debugging and inspecting runtime values in Oak programs, including pretty-printing data structures and visualizing numeric distributions with histograms.

## Import

```oak
debug := import('debug')
// or destructure specific functions
{ inspect: inspect, println: println, histo: histo, bar: bar } := import('debug')
```

## Functions

### `inspect(value, options?)`

Pretty-prints Oak data structures with customizable indentation and formatting. More readable than `string()` for complex data.

**Options:**
- `indent` - Indentation string (default: `'  '`)
- `depth` - Maximum nesting depth (default: `-1` = unlimited)
- `maxLine` - Maximum line length before multi-line format (default: `80`)
- `maxList` - Max list items before multi-line format (default: `16`)
- `maxObject` - Max object entries before multi-line format (default: `3`)

```oak
{ inspect: inspect } := import('debug')

data := {
    name: 'Alice'
    age: 30
    skills: ['Oak', 'JavaScript', 'Python']
    address: {
        city: 'Seattle'
        state: 'WA'
    }
}

println(inspect(data))
// Output:
// {
//   address: {
//     city: 'Seattle'
//     state: 'WA'
//   }
//   age: 30
//   name: 'Alice'
//   skills: ['Oak', 'JavaScript', 'Python']
// }

// With custom options
println(inspect(data, {
    indent: '    '
    depth: 1
    maxLine: 40
}))
```

### `println(value, options?)`

Shorthand for `println(inspect(value, options))`. Prints the inspected value.

**Note:** Not a drop-in replacement for `std.println` (not variadic).

```oak
{ println: println } := import('debug')

user := {
    id: 123
    name: 'Bob'
    active: true
}

println(user)
// Prints pretty-formatted user object
```

### `bar(n)`

Draws a single histogram bar as a Unicode string, where 1 character ≈ 1 unit of value.

Uses Unicode block characters for sub-character precision:
- Full blocks: `█`
- Partial blocks: `▏ ▎ ▍ ▌ ▋ ▊ ▉`

```oak
{ bar: bar } := import('debug')

bar(5) // => '█████'
bar(2.5) // => '██▌'
bar(0.5) // => '▌'
bar(0.1) // => '▏' (minimum non-zero)
bar(0) // => ''
```

### `histo(values, options?)`

Draws a histogram from a list of numbers using Unicode block characters.

**Options:**
- `min` - Minimum value (default: computed from data)
- `max` - Maximum value (default: computed from data)
- `bars` - Number of bars/buckets (default: `10`)
- `label` - Label position: `:start`, `:end`, or `?` (default: `?` = none)
- `cols` - Maximum column width in characters (default: `80`)

```oak
{ histo: histo } := import('debug')

scores := [65, 72, 88, 92, 95, 78, 81, 85, 90, 93]

println(histo(scores))
// Visual histogram output

// With labels at start
println(histo(scores, { label: :start }))
// Output (example):
// 1 █▌
// 2 ███
// 3 ████▌
// ...

// With custom options
println(histo(scores, {
    bars: 5
    label: :end
    cols: 40
}))
```

## Examples

### Debugging Complex Data Structures

```oak
{ inspect: inspect } := import('debug')

response := {
    status: 200
    headers: {
        'Content-Type': 'application/json'
        'Cache-Control': 'no-cache'
    }
    body: {
        users: [
            { id: 1, name: 'Alice', role: 'admin' }
            { id: 2, name: 'Bob', role: 'user' }
        ]
        total: 2
    }
}

println('API Response:')
println(inspect(response, { depth: 3 }))
```

### Inspecting with Depth Limits

```oak
{ inspect: inspect } := import('debug')

deeplyNested := {
    level1: {
        level2: {
            level3: {
                level4: 'too deep!'
            }
        }
    }
}

// Limit depth to 2 levels
println(inspect(deeplyNested, { depth: 2 }))
// Output:
// {
//   level1: {
//     level2: { 3 entries... }  // Truncated
//   }
// }
```

### Pretty-Printing Lists

```oak
{ inspect: inspect } := import('debug')

shortList := [1, 2, 3]
println(inspect(shortList))
// => [1, 2, 3]

longList := range(20)
println(inspect(longList, { maxList: 5 }))
// Multi-line format for long lists:
// [
//   0
//   1
//   2
//   ...
// ]
```

### String Escaping in Inspect

```oak
{ inspect: inspect } := import('debug')

special := 'Hello\nWorld\t!\x00'
println(inspect(special))
// => 'Hello\nWorld\t!\x00' (properly escaped)

// vs standard string()
println(string(special))
// => Hello
//    World	!  (literal newlines/tabs)
```

### Visualizing Distributions

```oak
{ histo: histo } := import('debug')
{ normal: normal } := import('random')

// Generate random data
samples := range(1000) |> map(fn normal())

// Visualize distribution
println('Normal Distribution:')
println(histo(samples, {
    bars: 20
    label: :start
    cols: 60
}))
```

### Monitoring Value Ranges

```oak
{ histo: histo } := import('debug')

temperatures := [18, 22, 25, 23, 19, 27, 24, 21, 20, 26]

println('Temperature Distribution:')
println(histo(temperatures, {
    min: 15
    max: 30
    bars: 15
    label: :end
}))
```

### Performance Visualization

```oak
{ histo: histo } := import('debug')

// Measure request times
requestTimes := measureRequests()

println('Request Duration Distribution:')
println(histo(requestTimes, {
    label: :end
    cols: 80
}))
```

### Comparing Data Sets

```oak
{ histo: histo } := import('debug')

before := [10, 12, 15, 11, 13, 14, 12]
after := [15, 17, 20, 16, 18, 19, 17]

println('Before optimization:')
println(histo(before, { label: :start }))

println('\nAfter optimization:')
println(histo(after, { label: :start }))
```

### Custom Bar Visualization

```oak
{ bar: bar } := import('debug')
{ scale: scale } := import('math')

values := [25, 50, 75, 100]
maxWidth := 40

each(values, fn(v) {
    scaled := scale(v, 0, 100, 0, maxWidth)
    println(string(v) + '%: ' + bar(scaled))
})
// Output:
// 25%: ██████████
// 50%: ████████████████████
// 75%: ██████████████████████████████
// 100%: ████████████████████████████████████████
```

### Debugging Function Results

```oak
{ inspect: inspect } := import('debug')

fn processData(input) {
    result := {
        input: input
        processed: transform(input)
        metadata: { timestamp: time(), version: '1.0' }
    }
    
    println('DEBUG: Result =', inspect(result, { depth: 2 }))
    result
}
```

### Object Key Formatting

```oak
{ inspect: inspect } := import('debug')

obj := {
    validIdentifier: 1
    'invalid key': 2
    123: 3
}

println(inspect(obj))
// Output:
// {
//   '123': 3
//   'invalid key': 2
//   validIdentifier: 1
// }
```

## Formatting Rules

### Primitive Types

```oak
{ inspect: inspect } := import('debug')

inspect(?) // => '?'
inspect(true) // => 'true'
inspect(42) // => '42'
inspect(3.14) // => '3.14'
inspect('hello') // => "'hello'"
inspect(:atom) // => ':atom'
inspect(fn(x) x) // => 'fn { ... }'
```

### String Escaping

Special characters in strings are escaped:
- `\n` → `\\n`
- `\t` → `\\t`
- `\r` → `\\r`
- `\f` → `\\f`
- `\'` → `\\'`
- `\\` → `\\\\`
- Non-printable bytes → `\\xHH` (hex)

### Object Key Formatting

- Valid identifiers: `key: value`
- Numbers: `123: value`
- Invalid identifiers: `'key name': value`

### Multi-Line Threshold

Switches to multi-line when:
- Line length exceeds `maxLine`
- List has more than `maxList` items
- Object has more than `maxObject` entries
- Contains nested non-primitive values

## Histogram Details

### Bucket Calculation

```oak
// Values are distributed into buckets
unit := (max - min) / numberOfBars

// Each value x goes into bucket:
bucket := int((x - min) / unit)
```

### Unicode Block Characters

- `█` - 8/8 (full block)
- `▉` - 7/8
- `▊` - 6/8
- `▋` - 5/8
- `▌` - 4/8
- `▍` - 3/8
- `▎` - 2/8
- `▏` - 1/8

### Label Options

```oak
// No labels
histo(data)
// =>
// ████
// ██████
// ███

// Labels at start
histo(data, { label: :start })
// =>
// 10 ████
// 15 ██████
//  8 ███

// Labels at end
histo(data, { label: :end })
// =>
// ████ 10
// ██████ 15
// ███ 8
```

## Performance Notes

- `inspect()` recursively traverses data structures
- Deep nesting consumes stack space
- Use `depth` limit for very deep structures
- `histo()` iterates through data twice (once for buckets, once for max)
- Large data sets may take time to format

## Limitations

- `println()` not variadic (unlike `std.println`)
- No custom formatters
- No color output configuration
- Histogram limited to numeric data
- No interactive debugging features
- No stack traces or breakpoints
- Object keys always sorted alphabetically

## See Also

- Oak built-in `string()` - Basic string conversion
- Oak built-in `type()` - Type checking
- Oak built-in `println()` - Standard output
- `fmt` library - For formatted strings
- `math` library - For value scaling
