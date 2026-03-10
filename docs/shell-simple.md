# Shell Simple Library (shell-simple)

## Overview

`libshell-simple` is a minimal debugging version of the shell library, providing a simplified API for testing and development.

## Import

```oak
shellSimple := import('shell-simple')
{ test: test } := import('shell-simple')
```

## Functions

### `test()`

Simple test function that prints "Test" to stdout.

**Returns:** `0` (success code)

```oak
{ test: test } := import('shell-simple')

test()
// Prints: Test
// Returns: 0
```

## Usage

This library is primarily used for:
- **Debugging**: Verify import system works
- **Testing**: Minimal library for test cases
- **Development**: Placeholder during development

## Example

```oak
{ test: test } := import('shell-simple')

result := test()
if result = 0 -> {
    println('Test passed')
}
```

## Implementation Notes

- Minimal implementation (no shell functionality)
- Used for debugging the library system
- Does not include full shell capabilities
- See `shell` library for complete shell functionality

## Limitations

- Single test function only
- No actual shell operations
- Debugging/testing only

## See Also

- [shell.md](shell.md) - Full shell library
- [fmt.md](fmt.md) - Formatting utilities
