# Virtual Shell

## Overview

The Oak Virtual Shell (`libshell`) provides a lightweight shell environment that can run in any environment, including those without a native shell (like WASM, embedded systems, or sandboxed environments). It features full integration with the Virtual File System (VFS), making it perfect for portable applications and packed binaries.

## Features

- **Cross-Platform**: Works anywhere Oak runs, including WASM
- **VFS Integration**: Full support for virtual filesystem operations
- **Batch Mode**: Execute shell scripts programmatically
- **Interactive REPL**: Command-line shell interface
- **Environment Variables**: Export and manage environment state
- **Path Navigation**: Comprehensive path handling with `..`, `./`, absolute, and relative paths

## API

### Creating a Shell

```js
shell := import('shell')

// Create shell with system filesystem
sh := shell.createShell()

// Create shell with virtual filesystem
Virtual := import('Virtual')
vfs := Virtual.createVirtualFS({ ... })
sh := shell.createShell(vfs)
```

### Shell Methods

#### `exec(cmdline)`

Execute a single command and return its exit code.

```js
exitCode := sh.exec('ls /home')
exitCode := sh.exec('cat file.txt')

if sh.exec('cd /nonexistent') {
    0 -> println('Success')
    _ -> println('Failed')
}
```

Returns:
- `0` on success
- Non-zero error code on failure
- `:exit` if the exit command was executed

#### `repl()`

Start an interactive read-eval-print loop (REPL).

```js
sh.repl()
```

This will show an interactive prompt:
```
Oak Shell v0.2 (type "help" for commands, "exit" to quit)
/> ls
/> cd /home
/home> pwd
/home
/home> exit
```

#### `batch(script)`

Execute multiple commands from a script string.

```js
script := '
# Setup project
cd /app
mkdir src
touch src/main.oak
ls src
'

exitCode := sh.batch(script)
```

Returns the exit code of the last command, or `0` if the script completes successfully.

#### `getCurrentDir()`

Get the current working directory.

```js
cwd := sh.getCurrentDir()
println(cwd)  // "/home/user"
```

#### `getEnv()`

Get a copy of the environment variables.

```js
env := sh.getEnv()
println(env.PATH)
println(env.HOME)
```

#### `setEnv(key, value)`

Set an environment variable programmatically.

```js
sh.setEnv('DEBUG', 'true')
sh.setEnv('APP_NAME', 'MyApp')
```

## Built-in Commands

### File Navigation

#### `pwd`
Print working directory.

```sh
pwd
# /home/user
```

#### `cd [directory]`
Change directory. Supports absolute paths, relative paths, `./`, and `../`.

```sh
cd /home
cd user
cd ./documents
cd ../..
cd  # goes to $HOME
```

#### `ls [directory]`
List files in the current or specified directory.

```sh
ls
ls /home
ls ..
```

### File Operations

#### `cat <file>...`
Print contents of one or more files.

```sh
cat file.txt
cat file1.txt file2.txt file3.txt
```

#### `touch <file>...`
Create empty files.

```sh
touch newfile.txt
touch file1.txt file2.txt
```

#### `cp <source> <destination>`
Copy a file.

```sh
cp original.txt copy.txt
cp /home/file.txt /backup/file.txt
```

#### `mv <source> <destination>`
Move or rename a file.

```sh
mv old.txt new.txt
mv /tmp/file.txt /home/file.txt
```

#### `rm [-r] <file>...`
Remove files. Use `-r` flag for recursive deletion (future support).

```sh
rm file.txt
rm file1.txt file2.txt
rm -r directory
```

#### `mkdir <directory>...`
Create directories.

```sh
mkdir newdir
mkdir dir1 dir2 dir3
mkdir /home/user/projects
```

### Utilities

#### `echo <text>...`
Print text to stdout.

```sh
echo Hello World
echo File created successfully
```

#### `env`
Display all environment variables.

```sh
env
# PATH=/bin:/usr/bin
# HOME=/home
# PWD=/
```

#### `export <VAR>=<value>`
Set environment variables.

```sh
export DEBUG=true
export APP_NAME=MyApplication
export PATH=/usr/local/bin:/usr/bin
```

#### `help`
Display help information about available commands.

```sh
help
```

#### `exit`
Exit the shell.

```sh
exit
```

## Command Line Usage

The shell can be invoked from the command line using the `oak shell` command:

### Interactive Mode

Start an interactive shell session:

```sh
oak shell
```

### Execute Single Command

Execute a single command with `-c`:

```sh
oak shell -c "ls /home"
oak shell -c "cat config.json"
```

### Run Script File

Execute commands from a file:

```sh
oak shell setup.sh
```

Example script file (`setup.sh`):
```sh
#!/usr/bin/env oak shell

# Project setup script
echo Setting up project...
cd /app
mkdir src
mkdir config
touch src/main.oak
touch config/settings.json
echo Done!
```

## Examples

### Example 1: VFS Shell Operations

```js
Virtual := import('Virtual')
shell := import('shell')

vfs := Virtual.createVirtualFS({
    '/home/README.md': '# Project\nWelcome!'
    '/home/config.json': '{"version": "1.0"}'
})

sh := shell.createShell(vfs)

sh.exec('cd /home')
sh.exec('ls')
sh.exec('cat README.md')
sh.exec('cp config.json backup.json')
sh.exec('ls')
```

### Example 2: Automated Setup Script

```js
shell := import('shell')
sh := shell.createShell()

setupScript := '
# Initialize project structure
echo Creating project structure...
mkdir /app
cd /app
mkdir src lib test
touch src/main.oak
touch README.md
echo export PROJECT_NAME=MyApp >> .env
ls
echo Project initialized!
'

result := sh.batch(setupScript)
if result = 0 {
    println('Setup completed successfully')
}
```

### Example 3: File Processing Pipeline

```js
shell := import('shell')
Virtual := import('Virtual')

vfs := Virtual.createVirtualFS({
    '/data/input.txt': 'raw data'
})

sh := shell.createShell(vfs)

sh.exec('cd /data')
sh.exec('cat input.txt')
sh.exec('cp input.txt processed.txt')
sh.exec('cat processed.txt')
```

### Example 4: Environment Configuration

```js
shell := import('shell')
sh := shell.createShell()

// Configure environment
sh.exec('export APP_ENV=production')
sh.exec('export LOG_LEVEL=info')
sh.exec('export MAX_WORKERS=4')

// Verify configuration
sh.exec('env')

// Access from Oak code
env := sh.getEnv()
println('Running in: ' << env.APP_ENV)
```

## Path Resolution

The shell supports comprehensive path resolution:

- **Absolute paths**: `/home/user/file.txt`
- **Relative paths**: `documents/file.txt`
- **Current directory**: `./file.txt`
- **Parent directory**: `../file.txt`
- **Multiple parent levels**: `../../other/file.txt`
- **Home directory**: `~` (via `cd` with no args)

Examples:
```sh
# Starting from /home/user
cd documents        # /home/user/documents
cd ./api           # /home/user/documents/api
cd ..              # /home/user/documents
cd ../../tmp       # /tmp
cd /               # /
```

## Integration with Packed Binaries

When using `oak pack` with embedded VFS, the shell automatically uses the virtual filesystem:

```sh
# Create packed binary with VFS
oak pack --entry main.oak --output myapp --include "config:./config,data:./data"
```

In your `main.oak`:
```js
shell := import('shell')

// Automatically uses packed VFS if available
sh := shell.createShell(___packed_vfs())

sh.repl()  // Interactive shell with access to embedded files
```

## Error Handling

Commands return exit codes following Unix conventions:

- `0`: Success
- `1`: General error
- `127`: Command not found
- `:exit`: Special value indicating shell exit requested

```js
result := sh.exec('cat nonexistent.txt')
if result {
    0 -> println('Success')
    1 -> println('File not found or read error')
    127 -> println('Unknown command')
    :exit -> println('Exit requested')
}
```

## Limitations and Future Work

Current limitations:

- **Limited real directory support**: On real filesystems, directory creation may be limited
- **No job control**: Background jobs with `&` are not supported
- **No command substitution**: `$(command)` or backticks are not supported

Implemented features:

- [x] Add pipe support (`|`) with full stdout capture between stages
- [x] Add redirection support (`>`, `>>`, `<`) with captured output
- [x] Add wildcard/glob expansion (`*`, `?`)
- [x] Add quoted argument parsing (single and double quotes)
- [x] Add command history (`history` command, `getHistory()` API)
- [x] Add alias support (`alias`, `unalias` commands, `addAlias()`/`removeAlias()` API)
- [x] Add more commands (`grep`, `find`, `head`, `tail`, `wc`, `sort`, `uniq`, `tee`)
- [x] Add tab completion engine (`tab` command, `complete()` API)
- [x] Full stdout capture for pipe stages and redirections

Future enhancements:

- [ ] Add readline-like line editing (arrow keys, Ctrl-A/E)
- [ ] Add job control (`&`)
- [ ] Add command substitution (`$(command)`)

## New Commands Reference

### `grep <pattern> [file...]`
Search for a text pattern in files or piped stdin.

```sh
grep TODO *.oak
cat file.txt | grep error
grep main src/app.oak src/lib.oak
```

### `head [-n N] [file...]`
Print the first N lines (default 10) of a file or stdin.

```sh
head file.txt
head -5 file.txt
cat log.txt | head -20
```

### `tail [-n N] [file...]`
Print the last N lines (default 10) of a file or stdin.

```sh
tail file.txt
tail -5 file.txt
cat log.txt | tail -20
```

### `wc [file...]`
Count lines, words, and characters.

```sh
wc file.txt
cat file.txt | wc
```

### `find [dir] [-name pattern]`
Recursively find files, optionally matching a glob pattern.

```sh
find .
find /home -name *.oak
find src -name test?.txt
```

### `sort [-r] [file...]`
Sort lines alphabetically. Use `-r` for reverse order.

```sh
sort names.txt
cat data.txt | sort -r
```

### `uniq [file]`
Remove adjacent duplicate lines.

```sh
sort names.txt | uniq
uniq sorted.txt
```

### `tee [-a] <file>`
Copy stdin to a file and stdout. Use `-a` to append.

```sh
cat data.txt | tee backup.txt
echo hello | tee -a log.txt
```

### `alias [name=value]`
Define or show command aliases.

```sh
alias ll='ls -l'
alias              # show all aliases
```

### `unalias <name>`
Remove a command alias.

```sh
unalias ll
```

### `history`
Show command history for the current session.

```sh
history
```

### `tab [partial...]`
Show completions for a partial command or path.

```sh
tab gr          # shows: grep
tab so          # shows: sort
tab src/        # shows files in src/ directory
```

## Tab Completion

The shell provides a completion engine accessible in two ways:

### REPL Command

Use the `tab` command interactively:

```sh
/> tab ec
echo
/> tab /ho
/home
```

### Programmatic API

Call `complete(partial)` from Oak code for GUI or custom shell integrations:

```js
shell := import('shell')
sh := shell.Shell()
sh.exec('mkdir src')
sh.exec('touch src/main.oak')
sh.exec('touch src/lib.oak')

matches := sh.complete('cat src/')
// => ['src/main.oak', 'src/lib.oak']

cmdMatches := sh.complete('gr')
// => ['grep']
```

Completion behavior:
- **At command position** (first word): completes command names and aliases
- **At argument position** (after first word): completes file/directory paths
- **After trailing space**: lists all files in current directory

## Pipes

Commands can be chained with `|` to pass data between stages:

```sh
cat file.txt | grep error | head -5
find . -name *.oak | sort
```

## Redirections

Output can be redirected to files with `>` (overwrite) or `>>` (append).
Input can be read from files with `<`.

```sh
ls > files.txt
echo hello >> log.txt
grep error < app.log
```

## Quoting

Arguments with spaces can be quoted using single or double quotes:

```sh
echo "hello world"
echo 'single quoted'
echo "escaped \"quote\""
```

## Glob Expansion

Wildcard patterns are expanded against the filesystem:

```sh
ls *.oak          # all .oak files
cat test?.txt     # test1.txt, testA.txt, etc.
```

## Testing

Run the shell tests:

```sh
oak test/shell.test.oak
```

Or run all tests including shell tests:

```sh
oak test/main.oak
```

## See Also

- [Virtual File System Documentation](virtual-fs.md)
- [Pack Command Documentation](pack.md)
- [Oak Standard Library](https://oaklang.org/lib/)
