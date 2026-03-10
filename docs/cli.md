# CLI Library (cli)

## Overview

`libcli` parses command-line arguments and flags, providing a structured interface for handling program invocation with verbs, options, and positional arguments.

## Import

```oak
cli := import('cli')
{ parse: parse, parseArgv: parseArgv } := import('cli')
```

## Functions

### `parse()`

Parses command-line arguments from `args()` builtin and returns a structured object.

**Returns:**
```oak
{
    exe: 'program-name'      // Executable name (args.0)
    main: 'main.oak'         // Main file (args.1)
    verb: 'command'          // Subcommand/verb (args.2 or ?)
    opts: { ... }            // Options/flags object
    args: [ ... ]            // Positional arguments
}
```

```oak
{ parse: parse } := import('cli')

// Command: ./oak main.oak serve --port 8080 --debug file1.txt file2.txt

cli := parse()
println(cli.verb)        // => 'serve'
println(cli.opts.port)   // => '8080'
println(cli.opts.debug)  // => true
println(cli.args)        // => ['file1.txt', 'file2.txt']
```

### `parseArgv(argv)`

Parses a custom argument array (instead of using `args()`).

```oak
{ parseArgv: parseArgv } := import('cli')

argv := ['program', 'main.oak', 'build', '--output', 'dist', 'src/']
cli := parseArgv(argv)

println(cli.verb)         // => 'build'
println(cli.opts.output)  // => 'dist'
println(cli.args)         // => ['src/']
```

## Argument Parsing Rules

### Flags (Boolean Options)

```oak
// -flag or --flag (implied true)
./program --verbose
// => opts.verbose = true

./program -v
// => opts.v = true
```

### Options (Key-Value)

```oak
// -opt value or --opt value
./program --port 8080
// => opts.port = '8080'

./program -p 3000
// => opts.p = '3000'
```

### Verb/Subcommand

First non-flag argument after main file:

```oak
./program main.oak build src/
//                  ^^^^^ verb
//                       ^^^^ positional arg
```

### Positional Arguments

Arguments that aren't flags or options:

```oak
./program main.oak file1.txt file2.txt
// => args = ['file1.txt', 'file2.txt']
```

### Double-Dash Convention

`--` marks the start of purely positional arguments:

```oak
./program --flag -- --not-a-flag file.txt
// => opts.flag = true
// => args = ['--not-a-flag', 'file.txt']
```

## Complete Examples

### Simple CLI Tool

```oak
{ parse: parse } := import('cli')

cli := parse()

if cli.verb {
    'help' -> println('Usage: program [options] <verb> [args]')
    'version' -> println('v1.0.0')
    'build' -> {
        output := cli.opts.output |> default('dist')
        verbose := cli.opts.verbose |> default(false)
        
        if verbose -> println('Building to: ' + output)
        
        each(cli.args, fn(file) {
            println('Processing: ' + file)
        })
    }
    _ -> println('Unknown verb: ' + cli.verb)
}
```

### Web Server CLI

```oak
{ parse: parse } := import('cli')
http := import('http')

cli := parse()

port := int(cli.opts.port |> default('8080'))
debug := cli.opts.debug != ?

if debug -> println('Debug mode enabled')

server := http.Server()
// ... configure routes ...

println('Starting server on port ' + string(port))
server.start(port)
```

### File Processor with Options

```oak
{ parse: parse } := import('cli')

cli := parse()

options := {
    recursive: cli.opts.recursive != ? | cli.opts.r != ?
    verbose: cli.opts.verbose != ? | cli.opts.v != ?
    output: cli.opts.output |> cli.opts.o |> default('out/')
}

if options.verbose -> {
    println('Recursive: ' + string(options.recursive))
    println('Output: ' + options.output)
}

each(cli.args, fn(file) {
    processFile(file, options)
})
```

### Git-like Subcommands

```oak
{ parse: parse } := import('cli')

cli := parse()

if cli.verb {
    'init' -> initRepository()
    'clone' -> {
        url := cli.args.0
        if url = ? -> {
            println('Error: clone requires URL')
            exit(1)
        }
        cloneRepository(url)
    }
    'commit' -> {
        message := cli.opts.message | cli.opts.m
        if message = ? -> {
            println('Error: -m or --message required')
            exit(1)
        }
        commitChanges(message)
    }
    ? -> println('No command specified. Use --help for usage.')
    _ -> println('Unknown command: ' + cli.verb)
}
```

### Help System

```oak
{ parse: parse } := import('cli')

cli := parse()

if cli.opts.help | cli.opts.h -> {
    println('Usage: myapp [options] <command> [args]')
    println('')
    println('Commands:')
    println('  build     Build the project')
    println('  test      Run tests')
    println('  serve     Start development server')
    println('')
    println('Options:')
    println('  --port NUM     Port number (default: 8080)')
    println('  --verbose      Enable verbose output')
    println('  --help, -h     Show this help')
    exit(0)
}
```

## Flag Name Conventions

```oak
// Short flags (single dash, single letter)
-v
-p 8080

// Long flags (double dash, full word)
--verbose
--port 8080

// Both can be used together
./program -v --port 8080
```

## Implementation Details

- Automatically detects if first argument is a flag (not a verb)
- If no verb provided, sets `verb` to `?`
- Flags without values default to `true`
- All option values are strings (must convert manually)
- Last occurrence of a flag wins if specified multiple times

## Edge Cases

```oak
// No verb (flag in verb position)
./program --flag arg
// => verb = ?, opts.flag = true, args = ['arg']

// Empty strings preserved
./program '' arg
// => verb = '', args = ['arg']

// Equals sign not supported
./program --key=value
// => opts.key = '=value' (not parsed as key=value pair)
```

## Limitations

- No automatic type conversion (all values are strings)
- No short flag bundling (`-abc` is not `-a -b -c`)
- No `--key=value` syntax support (use `--key value`)
- No automatic help generation
- No argument validation or type checking
- No default value handling (must use `default()` manually)
- No support for required arguments
- No nested subcommands

## See Also

- Oak built-in `args()` - Get command-line arguments
- `std.default()` - For default values
- Oak built-in `exit()` - Exit with code
