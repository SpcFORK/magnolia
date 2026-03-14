# Oak Programming Language Documentation

This is a work-in-progress rough draft of things that will end up in a rough informal language specification.

## Syntax

Oak, like [Ink](https://dotink.co), has automatic comma insertion at end of lines. This means if a comma can be inserted at the end of a line, it will automatically be inserted.

```go
program := expr*

expr := literal | identifier |
    assignment |
    propertyAccess |
    unaryExpr | binaryExpr |
    prefixCall | infixCall |
    ifExpr | withExpr |
    block

literal := nullLiteral |
    numberLiteral | stringLiteral | atomLiteral | boolLiteral |
    listLiteral | objectLiteral |
    fnLiteral | csLiteral

nullLiteral := '?'
numberLiteral := \d+ | \d* '.' \d+
stringLiteral := // single quoted string with standard escape sequences + \x00 syntax
atomLiteral := ':' + identifier
boolLiteral := 'true' | 'false'
listLiteral := '[' ( expr ',' )* ']' // last comma optional
objectLiteral := '{' ( expr ':' expr ',' )* '}' // last comma optional
fnLiteral := 'fn' '(' ( identifier ',' )* (identifier '...')? ')' expr
csLiteral := 'cs' identifier ('(' ( identifier ',' )* (identifier '...')? ')')? expr

identifier := \w_ (\w\d_?!)* | _

assignment := (
    identifier [':=' '<-'] expr |
    listLiteral [':=' '<-'] expr |
    objectLiteral [':=' '<-'] expr
)

propertyAccess := identifier ('.' identifier)+

unaryExpr := ('!' | '-') expr
binaryExpr := expr (+ - * / % ^ & | > < = >= <= <<) binaryExpr

prefixCall := expr '(' (expr ',')* ')'
infixCall := expr '|>' prefixCall

ifExpr := 'if' expr? '{' ifClause* '}'
ifClause := expr '->' expr ','

withExpr := 'with' prefixCall fnLiteral

block := '{' expr+ '}' | '(' expr* ')'
```

### AST node types

```c
nullLiteral
stringLiteral
numberLiteral
boolLiteral
atomLiteral
listLiteral
objectLiteral
fnLiteral
csLiteral
identifier
assignment
propertyAccess
unaryExpr
binaryExpr
fnCall
ifExpr
block
```

## Language Functions

- `import(path)`: Imports a module located at the specified `path`.
- `string(x)`: Converts the argument `x` to a string.
- `int(x)`: Converts the argument `x` to an integer.
- `float(x)`: Converts the argument `x` to a floating-point number.
- `atom(c)`: Creates an atom with the specified character `c`.
- `codepoint(c)`: Returns the Unicode code point of the character `c`.
- `char(n)`: Converts the Unicode code point `n` to a character.
- `type(x)`: Returns the type of the argument `x`.
- `name(x)`: Returns an atom name for `x` (for example class and function names).
- `csof(a, b)`: Returns `true` when `a` and `b` refer to the same class, or when one is a class and the other is an atom matching that class name.
- `len(x)`: Returns the length of the argument `x`.
- `keys(x)`: Returns an array of keys of the argument `x`.

## OS Functions

- `args()`: Returns command-line arguments as an array of strings.
- `env()`: Returns the environment variables as an object.
- `time()`: Returns the current time as a float.
- `nanotime()`: Returns the current time in nanoseconds as an integer.
- `exit(code)`: Exits the program with the specified exit code.
- `rand()`: Generates a random floating-point number between 0 and 1.
- `srand(length)`: Seeds the random number generator with the specified length.
- `wait(duration)`: Pauses the program execution for the specified duration.
- `exec(path, args, stdin)`: Executes a command specified by `path` with the given `args` and optional standard input `stdin`. Returns stdout, stderr, and end events.
- `sysproc(library, name)`: Gets a system procedure from a library.
- `syscall(proc, args...)`: Calls a system procedure with the given arguments.
- `utf16(string)`: Converts a string to UTF-16 encoding.
- `go(fn, args...)`: Spawns a goroutine with the specified function and arguments.
- `make_chan(cap?)`: Creates a channel with an optional capacity.
- `chan_send(ch, value, callback?)`: Sends a value to a channel with an optional callback.
- `chan_recv(ch, callback?)`: Receives a value from a channel with an optional callback.
- `bits(x)`: Converts between list of bytes and byte string.
- `addr(bits)`: Gets the address from bits.
- `pointer(x)`: Converts an integer, atom name, or byte string to a pointer type.
- `memread(address, length)`: Reads memory at the specified address.
- `memwrite(address, bits)`: Writes bits to memory at the specified address.

### System Interop Result Shapes

- `sysproc(library, name)` result:
  - success: `{type: :proc, library: string, name: string, addr: pointer}`
  - error: `{type: :error, error: string, library: string, name: string}`
- `syscall(procOrAddress, args...)` result:
  - success: `{type: :ok, r1: int, r2: int}`
  - error: `{type: :error, error: string}`
- `chan_recv(ch, callback?)` synchronous return:
  - event: `{data: any, ok: bool}`

### System Interop Notes

- `sysproc`/`syscall` behavior depends on host OS and available native libraries.
- Memory APIs (`memread`, `memwrite`) are unsafe by design and can crash the process on invalid addresses.
- Prefer wrapping interop calls in higher-level modules such as `gpu` or `sys` for safer application code.

## I/O Interfaces

- `input()`: Reads input from the standard input.
- `print()`: Writes output to the standard output.
- `ls(path)`: Lists files and directories in the specified path.
- `mkdir(path)`: Creates a directory at the specified path.
- `rm(path)`: Removes the file or directory at the specified path.
- `stat(path)`: Retrieves file or directory information at the specified path.
- `open(path, flags, perm)`: Opens a file at the specified path with the given flags and permissions.
- `close(fd)`: Closes the file descriptor `fd`.
- `read(fd, offset, length)`: Reads data from the file descriptor `fd` starting at the specified `offset` and reading `length` bytes.
- `write(fd, offset, data)`: Writes data to the file descriptor `fd` starting at the specified `offset`.
- `close := listen(host, handler)`: Listens for incoming connections on the specified `host` and handles them with the provided `handler` function.
- `req(data)`: Sends an HTTP request with the provided data.
  
  ```go
  // Req syntax:
  // ---
  
  req({
    url: ''
    method: 'GET'
    headers: {}
    body: _
  })
  ```

## Math Functions

- Trigonometric functions
  - `sin(n)`: Calculates the sine of the angle `n`.
  - `cos(n)`: Calculates the cosine of the angle `n`.
  - `tan(n)`: Calculates the tangent of the angle `n`.
- Inverse Trigonometric functions
  - `asin(n)`: Calculates the arcsine of the value `n`.
  - `acos(n)`: Calculates the arccosine of the value `n`.
  - `atan(n)`: Calculates the arctangent of the value `n`.
- Power and logarithmic functions
  - `pow(b, n)`: Raises the base `b` to the power of `n`.
  - `log(b, n)`: Calculates the logarithm of `n` with base `b`.

## Code samples

```js
// hello world
std.println('Hello, World!')

// some math
sq := fn(n) n * n
fn sq(n) n * n
fn sq(n) { n * n } // equivalent

// side-effecting functions
fn say() { std.println('Hi!') }
// if no arguments, () is optiona
fn { std.println('Hi!') }

// class constructor sugar
cs Pair(left, right) {
	{ left: left, right: right }
}
Pair(1, 2).left

// factorial
fn factorial(n) if n <= 1 {
	true -> 1
	_ -> n * factorial(n - 1)
}
```

```js
// methods are emulated by pipe notation
scores |> sum()
names |> join(', ')
fn sum(xs...) xs |> reduce(0, fn(a, b) a + b)
oakFiles := fileNames |> filter(fn(name) name |> endsWith?('.oak'))
```

```js
// "with" keyword just makes the last fn a callback as last arg
with loop(10) fn(n) std.println(n)
with wait(1) fn {
	std.println('Done!')
}
with fetch('example.com') fn(resp) {
	with resp.json() fn(json) {
		std.println(json)
	}
}
```

```js
// raw file read
with open('name.txt') fn(evt) if evt.type {
	:error -> std.println(evt.message)
	_ -> with read(fd := evt.fd, 0, -1) fn(evt) {
		if evt.type {
			:error -> std.println(evt.message)
			_ -> fmt.printf('file data: {{0}}', evt.data)
		}
		close(fd)
	}
}

// with stdlib
std := import('std')
fs := import('fs')
with fs.readFile('names.txt') fn(file) if file {
	? -> std.println('[error] could not read file')
	_ -> std.println(file)
}
```

