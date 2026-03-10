# magnolia 🌸

**magnolia** is an expressive, dynamically typed programming language based on [oak](https://oaklang.org/). it extends oak with powerful new features including a transpile middleware system, virtual file system, advanced threading utilities, and gpu computing support, while maintaining the simplicity and elegance of the original language.

here's an example magnolia program.

```js
std := import('std')

fn fizzbuzz(n) if [n % 3, n % 5] {
    [0, 0] -> 'fizzbuzz'
    [0, _] -> 'fizz'
    [_, 0] -> 'buzz'
    _ -> string(n)
}

std.range(1, 101) |> std.each(fn(n) {
    std.println(fizzbuzz(n))
})
```

magnolia has good support for asynchronous i/o. here's how you read a file and print it.

```js
std := import('std')
fs := import('fs')

with fs.readfile('./file.txt') fn(file) if file {
    ? -> std.println('could not read file!')
    _ -> print(file)
}
```

magnolia also has a pragmatic standard library that comes built into the interpreter executable. for example, there's a built-in http server and router in the `http` library.

```js
std := import('std')
fmt := import('fmt')
http := import('http')

server := http.server()
with server.route('/hello/:name') fn(params) {
    fn(req, end) if req.method {
        'get' -> end({
            status: 200
            body: fmt.format('hello, {{ 0 }}!'
                std.default(params.name, 'world'))
        })
        _ -> end(http.methodnotallowed)
    }
}
server.start(9999)
```

## install

magnolia is currently installed from source.

on unix-like systems, build with make:

```sh
make install
```

on windows, use the provided build script:

```bat
build.bat
```

or build directly with go on any platform:

```sh
go build .
```

you can also run without installing:

```sh
go run . <file-or-command>
```

## what's new in magnolia

magnolia extends oak with powerful new features for modern development:

### 🎨 enhanced error display

beautiful, color-coded error messages with source code context to help you quickly identify and fix issues:

```
╭─ runtime error ───────────────────────────────────────────
│
│ file: test.oak
│ position: [4:8]
│
│ division by zero
│
│ context:
│    2 │ x := 10
│    3 │ y := 20
│    4 │ z := x / 0
│      │        ^
│    5 │ 
│    6 │ println(z)
╰───────────────────────────────────────────────────────────
```

see [error-display.md](docs/error-display.md) for more details.

### 🔧 transpile middleware

a plugin architecture for ast transformations during the build process. write custom transpilers to transform your code at compile-time:

```js
build := import('build')
transpile := build.transpile

// create custom transpiler
mytranspiler := transpile.createtranspiler(fn(node) {
    // transform ast nodes
    node
})

build.run({
    entry: 'main.oak'
    transpilers: [mytranspiler]
})
```

### 📁 virtual file system

an in-memory file system that can be embedded in packed binaries, enabling true cross-platform deployment:

```js
virtual := import('virtual')

vfs := virtual.createvirtualfs({
    'config.json': '{"version": "1.0"}'
    'data/test.txt': 'test data'
})

content := vfs.readfile('config.json')
vfs.writefile('output.txt', 'hello world')
```

### 🧵 thread library

high-level utilities for concurrent and parallel programming, including mutexes, semaphores, wait groups, and thread pools:

```js
thread := import('thread')

// parallel map
results := thread.pmap([1, 2, 3, 4], fn(x) x * x)

// mutex for safe shared state
mutex := thread.mutex()
mutex.lock()
// critical section
mutex.unlock()

// thread pool
pool := thread.pool(4)
pool.submit(fn() {
    // work to be done
})
```

### 🎮 gpu computing

low-level