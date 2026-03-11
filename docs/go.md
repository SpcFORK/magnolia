# Go Runtime and System Interop

## Overview

Magnolia exposes a small set of low-level built-ins that let Oak code interact with the host Go runtime and operating system primitives.

This includes:

- Goroutines and channels (`go`, `make_chan`, `chan_send`, `chan_recv`)
- Runtime metadata (`___runtime_go_version`, `___runtime_sys_info`)
- Foreign function calls (`sysproc`, `syscall`)
- Thread affinity helpers (`lock_thread`, `unlock_thread`)
- Native Win32 loop helper (`win_msg_loop`, Windows only)
- Raw byte and memory helpers (`bits`, `addr`, `pointer`, `memread`, `memwrite`)

These APIs are powerful and intentionally minimal. They are best used inside thin wrapper libraries for safety and readability.

For a safer wrapper around `sysproc` and `syscall`, see `lib/sys.oak`.

## Concurrency Primitives

### `go(fn, args...)`

Runs a function on a new goroutine.

```oak
ch := make_chan()

go(fn {
    chan_send(ch, :ready)
})

evt := chan_recv(ch)
println(evt.data) // :ready
```

### `make_chan(cap?)`

Creates a channel value. If `cap` is omitted, the channel is unbuffered.

```oak
unbuffered := make_chan()
buffered := make_chan(8)
```

### `chan_send(ch, value, callback?)`

Sends `value` to `ch`. If a callback is provided, the operation is coordinated through callback flow.

```oak
ch := make_chan(1)
chan_send(ch, 42)
```

### `chan_recv(ch, callback?)`

Receives from `ch`, returning an event object.

Event object shape:

- `data`: received value
- `ok`: boolean receive status

```oak
ch := make_chan(1)
chan_send(ch, 'hello')

msg := chan_recv(ch)
println(msg.data) // 'hello'
println(msg.ok)   // true
```

## Runtime Metadata

### `___runtime_go_version()`

Returns the host Go version as a string.

```oak
println(___runtime_go_version())
```

### `___runtime_sys_info()`

Returns an object with host runtime information.

Common fields include:

- `os`
- `arch`
- `cpus`

```oak
info := ___runtime_sys_info()
println(info.os)
println(info.arch)
println(string(info.cpus))
```

## Foreign Function Interop

### `sysproc(library, name)`

Resolves a procedure by library and symbol name.

Returns either a proc object or an error object.

```oak
proc := sysproc('kernel32.dll', 'GetCurrentProcessId')
if proc.type = :proc -> {
    result := syscall(proc)
    println(result.type) // :ok or :error
}
```

### `syscall(procOrAddress, args...)`

Calls a procedure target returned by `sysproc` (or a numeric address).

Result object shape:

- `type`: `:ok` or `:error`
- On success: `r1`, `r2`
- On error: `error`

```oak
pidProc := sysproc('kernel32.dll', 'GetCurrentProcessId')
res := syscall(pidProc)

if res.type = :ok -> {
    println('PID: ' + string(res.r1))
} else {
    println('syscall failed: ' + res.error)
}
```

### `lock_thread()` / `unlock_thread()`

Pins the current Oak execution to the current OS thread, then releases it.

This is especially important for Win32 UI code, where window creation and
message pumping must happen on the same OS thread.

```oak
if ___runtime_sys_info().os = 'windows' {
    true -> {
        lock_thread()
        // Create window + run message loop here.
        unlock_thread()
    }
}
```

### `int(x)` with pointers

`int(...)` accepts pointer values and converts them to integer addresses.

Useful when writing pointer-sized values into native struct buffers.

```oak
buf := bits([0, 0, 0, 0, 0, 0, 0, 0])
ptr := addr(buf)
n := int(ptr)
println(string(n))
```

### `win_msg_loop(hwnd)`

Windows-only helper that runs a native `GetMessageW` / `TranslateMessage` /
`DispatchMessageW` loop for a window handle until exit.

Returns `0` on normal loop exit, or an error object.

```oak
if ___runtime_sys_info().os = 'windows' {
    true -> {
        lock_thread()
        // hwnd should be a valid window handle (int/pointer value)
        result := win_msg_loop(hwnd)
        unlock_thread()
        println(string(result))
    }
}
```

## Raw Bytes and Memory

### `bits(x)`

Converts between byte list and byte string.

```oak
raw := bits([65, 66, 67])
nums := bits(raw) // [65, 66, 67]
```

### `addr(byteString)` and `pointer(x)`

- `addr` returns an address for byte-backed data.
- `pointer` converts integers/pointers/byte strings into pointer values.

```oak
buf := bits([65, 66, 67])
ptr := pointer(addr(buf))
```

### `memread(addressOrPointer, length)`

Reads `length` bytes from memory and returns a byte string.

### `memwrite(addressOrPointer, data)`

Writes byte data to memory.

```oak
buf := bits([65, 66, 67])
ptr := addr(buf)

memwrite(ptr + 1, bits([90]))
println(bits(memread(ptr, 3))) // [65, 90, 67]
```

## Safety Notes

- Avoid arbitrary memory addresses. Invalid reads/writes can crash the process.
- Keep buffers alive while passing pointers to native calls.
- Prefer passing explicit byte strings for pointer-backed arguments.
- Treat `syscall` boundaries as unsafe code and validate all arguments.
- Consider wrapping low-level calls in library APIs (for example, `gpu`) instead of direct use in app code.

## Sample Program

See `samples/go-interop.oak` for a runnable example that combines:

- goroutine + channel coordination
- runtime metadata
- safe `sysproc`/`syscall` invocation via `lib/sys.oak`
- direct memory read/write on a byte buffer

## Related Docs

- [Language spec](./spec.md)
- [GPU interop library](./gpu.md)
- [Thread library](./thread.md)
- [Bitwise and pointer helpers](./bitwise.md)
- [Safer syscall wrappers](../lib/sys.oak)
