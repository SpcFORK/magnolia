# String Library (str)

## Overview

`libstr` is the core string library for Oak, providing utilities for string manipulation, searching, splitting, case conversion, and character classification.

## Import

```oak
str := import('str')
// or destructure specific functions
{ split: split, join: join, trim: trim } := import('str')
```

## Character Classification

All character classification functions work with single ASCII characters.

### `upper?(c)`

Returns true if `c` is an uppercase ASCII letter (A-Z).

```oak
upper?('A') // => true
upper?('a') // => false
```

### `lower?(c)`

Returns true if `c` is a lowercase ASCII letter (a-z).

```oak
lower?('a') // => true
lower?('A') // => false
```

### `digit?(c)`

Returns true if `c` is an ASCII digit (0-9).

```oak
digit?('5') // => true
digit?('a') // => false
```

### `space?(c)`

Returns true if `c` is ASCII whitespace (space, tab, newline, carriage return, or form feed).

```oak
space?(' ') // => true
space?('\t') // => true
space?('a') // => false
```

### `letter?(c)`

Returns true if `c` is an ASCII letter (uppercase or lowercase).

```oak
letter?('a') // => true
letter?('Z') // => true
letter?('5') // => false
```

### `word?(c)`

Returns true if `c` is a letter or digit.

```oak
word?('a') // => true
word?('5') // => true
word?('-') // => false
```

### `checkRange(lo, hi)`

Returns a function that checks if a character's codepoint is within `[lo, hi]` (inclusive).

```oak
isDigit := checkRange('0', '9')
isDigit('5') // => true
isDigit('a') // => false
```

## String Joining

### `join(strings, joiner)`

Concatenates a list of strings with a separator. Default separator is empty string.

```oak
join(['a', 'b', 'c'], ', ') // => 'a, b, c'
join(['hello', 'world'], ' ') // => 'hello world'
join(['a', 'b', 'c']) // => 'abc'
join([]) // => ''
```

## String Testing

### `startsWith?(s, prefix)`

Returns true if string `s` starts with `prefix`.

```oak
startsWith?('hello world', 'hello') // => true
startsWith?('hello world', 'world') // => false
```

### `endsWith?(s, suffix)`

Returns true if string `s` ends with `suffix`.

```oak
endsWith?('hello world', 'world') // => true
endsWith?('hello world', 'hello') // => false
```

### `contains?(s, substr)`

Returns true if `s` contains the substring `substr`.

```oak
contains?('hello world', 'lo wo') // => true
contains?('hello world', 'xyz') // => false
```

## String Searching

### `indexOf(s, substr)`

Returns the first index where `substr` appears in `s`, or `-1` if not found.

```oak
indexOf('hello world', 'world') // => 6
indexOf('hello world', 'xyz') // => -1
indexOf('aaabbb', 'a') // => 0
```

### `rindexOf(s, substr)`

Returns the last index where `substr` appears in `s`, or `-1` if not found.

```oak
rindexOf('hello world', 'o') // => 7
rindexOf('aaabbb', 'a') // => 2
rindexOf('hello', 'xyz') // => -1
```

## String Splitting

### `cut(s, sep)`

Splits string `s` at the first occurrence of separator `sep`. Returns `[before, after]`. If separator not found, returns `[s, '']`.

```oak
cut('key=value', '=') // => ['key', 'value']
cut('a:b:c', ':') // => ['a', 'b:c']
cut('hello', '=') // => ['hello', '']
```

### `split(s, sep)`

Splits string `s` by all occurrences of `sep`. If `sep` is not specified or `?`, returns a list of individual characters.

```oak
split('a,b,c', ',') // => ['a', 'b', 'c']
split('hello') // => ['h', 'e', 'l', 'l', 'o']
split('hello', '') // => ['h', 'e', 'l', 'l', 'o']
split('a::b::c', '::') // => ['a', 'b', 'c']
```

## Case Conversion

### `lower(s)`

Converts all uppercase letters in `s` to lowercase.

```oak
lower('HELLO World') // => 'hello world'
lower('ABC123') // => 'abc123'
```

### `upper(s)`

Converts all lowercase letters in `s` to uppercase.

```oak
upper('hello WORLD') // => 'HELLO WORLD'
upper('abc123') // => 'ABC123'
```

## String Replacement

### `replace(s, old, new)`

Replaces all occurrences of substring `old` with `new` in string `s`. Does nothing for empty strings.

```oak
replace('hello world', 'world', 'Oak') // => 'hello Oak'
replace('aaa', 'a', 'b') // => 'bbb'
replace('test', 'x', 'y') // => 'test'
replace('hello', '', 'x') // => 'hello' (no-op for empty old)
```

## String Padding

### `padStart(s, n, pad)`

Prepends repetitions of `pad` to string `s` until it reaches at least `n` characters. Returns `s` if already long enough.

```oak
padStart('5', 3, '0') // => '005'
padStart('hello', 10, '-') // => '-----hello'
padStart('longstring', 5, 'x') // => 'longstring'
```

### `padEnd(s, n, pad)`

Appends repetitions of `pad` to string `s` until it reaches at least `n` characters. Returns `s` if already long enough.

```oak
padEnd('5', 3, '0') // => '500'
padEnd('hello', 10, '-') // => 'hello-----'
padEnd('longstring', 5, 'x') // => 'longstring'
```

## String Trimming

### `trimStart(s, prefix)`

Removes all repeated occurrences of `prefix` from the start of `s`. If `prefix` is `?` (not specified), removes whitespace.

```oak
trimStart('   hello', ?) // => 'hello'
trimStart('!!!hello', '!') // => 'hello'
trimStart('aaabbb', 'a') // => 'bbb'
trimStart('xxhello', 'xx') // => 'hello'
```

### `trimEnd(s, suffix)`

Removes all repeated occurrences of `suffix` from the end of `s`. If `suffix` is `?` (not specified), removes whitespace.

```oak
trimEnd('hello   ', ?) // => 'hello'
trimEnd('hello!!!', '!') // => 'hello'
trimEnd('aaabbb', 'b') // => 'aaa'
trimEnd('helloxx', 'xx') // => 'hello'
```

### `trim(s, part)`

Removes repeated occurrences of `part` from both ends of `s`. If `part` is `?` (not specified), removes whitespace.

```oak
trim('  hello  ', ?) // => 'hello'
trim('!!!hello!!!', '!') // => 'hello'
trim('  \t hello \n  ') // => 'hello'
trim('xxhelloxx', 'xx') // => 'hello'
```

## Examples

### URL Parsing

```oak
{
    split: split
    cut: cut
    contains?: contains?
} := import('str')

url := 'https://example.com/path/to/page?key=value'

// Extract protocol
[protocol, rest] := cut(url, '://')
// protocol = 'https', rest = 'example.com/path/to/page?key=value'

// Check if URL uses HTTPS
if contains?(url, 'https') {
    println('Secure connection')
}
```

### CSV Parsing

```oak
{ split: split, trim: trim } := import('str')

csvLine := 'John Doe, 30, Engineer'
fields := csvLine |> split(',') |> map(fn(f) trim(f, ?))
// fields = ['John Doe', '30', 'Engineer']
```

### String Formatting

```oak
{
    padStart: padStart
    padEnd: padEnd
    upper: upper
} := import('str')

fn formatId(n) padStart(string(n), 5, '0')
formatId(42) // => '00042'

fn formatName(s) upper(s.0) << lower(slice(s, 1))
formatName('JOHN') // => 'John'
```

## Notes

- All functions create new strings rather than modifying the original
- Empty string `''` is used as the default joiner in `join()`
- Character classification functions only work reliably with ASCII characters
- `split()` without a separator returns individual characters
- Trimming functions can remove multi-character patterns, not just single characters
