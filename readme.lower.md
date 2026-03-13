<h1>
    <img width="24" height="24" alt="🌳" src="https://github.com/user-attachments/assets/22fadd8f-707e-4279-8ea0-63bd3da6fdba" />
    <em><b>⠀magnolia</b></em>⠀🌸
</h1>

is an expressive, dynamically typed programming language based on [oak](https://oaklang.org/). it extends oak with powerful new features including a transpile middleware system, virtual file system, advanced threading utilities, and gpu computing support, while maintaining the simplicity and elegance of the original language.

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

### command name note

depending on how you build/install magnolia, the executable may be named either `magnolia` or `oak`:

- `go build .` in this repository typically produces `magnolia`/`magnolia.exe`
- `make install` installs the binary as `oak`

in examples below, use whichever name matches your local install.

### quick start

```sh
# start repl
magnolia repl

# run a file
magnolia samples/hello.oak

# evaluate a one-liner
magnolia eval "1 + 2 * 3"

# show cli help
magnolia help
```

## what's new in magnolia

magnolia extends oak with powerful new features for modern development:

### latest platform updates (march 2026)

- windows gui 2d includes vulkan support, while default `auto` prioritizes stable presenters `opengl -> ddraw -> gdi` (`vulkanauto` enables vulkan in auto mode).
- vulkan bootstrap on windows now validates core instance extensions (`vk_khr_surface`, `vk_khr_win32_surface`), creates a win32 surface, and selects a queue family that supports both graphics and present.
- window state now exposes vulkan runtime handles/selection details via `window.vulkansurface`, `window.vulkanphysicaldevice`, and `window.vulkanqueuefamily`.
- win32 class registration uses a null background brush (`hbrbackground = 0`) to reduce background flash between presents.
- gui frame scheduling now supports `maxframedtms` clamping and urgent redraw triggering for resize-related windows messages.

see [docs/gui-native-win.md](docs/gui-native-win.md) and [docs/gui.md](docs/gui.md) for api details.

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
│    6