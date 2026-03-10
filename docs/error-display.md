# Error Display System

## Overview

Magnolia features an enhanced error display system that provides clear, color-coded error messages with source code context to help developers quickly identify and fix issues.

## Features

### 1. **Color-Coded Output**
Errors are displayed with ANSI color codes for better visibility:
- **Red**: Error messages and error lines
- **Cyan**: Labels and file information
- **Gray**: Context lines and stack traces
- **Yellow**: Stack trace headers

### 2. **Source Code Context**
When an error occurs, the system displays:
- The error line with a visual pointer (`^`) showing the exact column
- Surrounding context lines (configurable, default: 2 lines before and after)
- Line numbers for easy navigation

### 3. **Enhanced Position Information**
Error messages include:
- File name
- Line and column number
- Stack traces (for runtime errors)

### 4. **Error Types**

#### Parse Errors
Errors that occur during parsing (syntax errors):
```
╭─ Parse Error ─────────────────────────────────────────────
│
│ File: test.oak
│ Position: [3:10]
│
│ Unexpected token ']'
│
│ Context:
│    1 │ x := [1, 2, 3
│    2 │ 
│    3 │ y := x]
│      │          ^
╰───────────────────────────────────────────────────────────
```

#### Runtime Errors
Errors that occur during execution:
```
╭─ Runtime Error ───────────────────────────────────────────
│
│ File: test.oak
│ Position: [4:8]
│
│ Division by zero
│
│ Context:
│    2 │ x := 10
│    3 │ y := 20
│    4 │ z := x / 0
│      │        ^
│    5 │ 
│    6 │ println(z)
│
│ Stack Trace:
│   in fn calculate [3:5]
│   in fn main [10:2]
╰───────────────────────────────────────────────────────────
```

## Configuration

The error display system can be configured using `ErrorDisplayConfig`:

```go
type ErrorDisplayConfig struct {
    UseColor       bool  // Enable/disable ANSI colors
    ShowContext    bool  // Show source code context
    ContextLines   int   // Number of context lines before/after error
    ShowStackTrace bool  // Show stack trace for runtime errors
}
```

### Default Configuration
```go
config := DefaultErrorConfig()
// UseColor: true
// ShowContext: true
// ContextLines: 2
// ShowStackTrace: true
```

### Custom Configuration Examples

#### Minimal Error Display (for CI/CD)
```go
config := ErrorDisplayConfig{
    UseColor:       false,  // No colors for log files
    ShowContext:    false,  // No source context
    ContextLines:   0,
    ShowStackTrace: true,   // Keep stack traces
}
DisplayError(err, config)
```

#### Extended Context (for debugging)
```go
config := ErrorDisplayConfig{
    UseColor:       true,
    ShowContext:    true,
    ContextLines:   5,     // Show 5 lines before/after
    ShowStackTrace: true,
}
DisplayError(err, config)
```

## API Usage

### Displaying Errors
```go
// Use default configuration
DisplayError(err, DefaultErrorConfig())

// Use custom configuration
config := ErrorDisplayConfig{
    UseColor:     true,
    ShowContext:  true,
    ContextLines: 3,
}
DisplayError(err, config)
```

### Formatting Errors (without printing)
```go
// Get formatted error string
formatted := FormatError(err)
// Use formatted string elsewhere (logging, etc.)
```

## Implementation Details

### File Tracking
The Context now tracks the current file being executed:
```go
ctx := NewContext("/path/to/dir")
ctx.currentFile = "example.oak"
```

This ensures that:
- Error messages include the correct file name
- Source context is loaded from the right file
- Multi-file projects show accurate error locations

### Context-Specific Display
Different execution contexts use different configurations:

- **File Execution**: Full error display with context
- **REPL**: Error display without source context (since code is entered line-by-line)
- **Eval Mode**: Error display without source context
- **Library Loading**: Errors include library name in file field

## Examples

### Division by Zero
```oak
x := 10
y := 20
z := x / 0   // Error!
```

Output:
```
╭─ Runtime Error ───────────────────────────────────────────
│
│ File: test.oak
│ Position: [3:8]
│
│ Division by zero
│
│ Context:
│    1 │ x := 10
│    2 │ y := 20
│    3 │ z := x / 0   // Error!
│      │        ^
╰───────────────────────────────────────────────────────────
```

### Type Error
```oak
x := 'hello'
y := x / 2   // Cannot divide string by number
```

Output:
```
╭─ Runtime Error ───────────────────────────────────────────
│
│ File: test.oak
│ Position: [2:8]
│
│ Cannot / incompatible values 'hello', 2
│
│ Context:
│    1 │ x := 'hello'
│    2 │ y := x / 2
│      │        ^
╰───────────────────────────────────────────────────────────
```

## Benefits

1. **Faster Debugging**: Visual context helps locate errors quickly
2. **Better Developer Experience**: Clear, formatted output is easier to read
3. **Production-Ready**: Configurable for different environments
4. **Accessible**: Works in all terminal environments (colors degrade gracefully)
5. **Educational**: Helpful for learning the language and understanding errors

## Notes

- ANSI colors work in most modern terminals (Windows 10+, macOS, Linux)
- Colors can be disabled for environments that don't support them
- Source context requires the file to be accessible on disk
- REPL errors don't show source context since input is ephemeral
