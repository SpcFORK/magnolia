# Socket Library (socket)

## Overview

`libsocket` exposes raw TCP/TLS stream helpers used by higher-level protocol libraries such as SMTP, POP3, and IMAP.

This library currently targets Magnolia's native runtime. In the JavaScript runtime, socket built-ins return structured `:error` results because synchronous stream sockets are not exposed there yet.

## Import

```oak
socket := import('socket')
```

## Client API

### `connect(address, options?)`

Connects to `host:port`.

Options:

- `tls: true` for implicit TLS
- `serverName` for TLS SNI / certificate validation
- `insecureSkipVerify: true` to disable certificate verification

Returns either:

```oak
{ type: :ok, socket: {type: :socket, id: 0}, remote: '...', local: '...', tls: false }
```

or an error object.

### `send(socket, data)`

Writes a string to the stream.

### `sendLine(socket, line)`

Writes `line + "\r\n"`.

### `recv(socket, size)`

Reads up to `size` bytes and returns `{type: :data, data: '...'}` or `{type: :closed}`.

### `recvExact(socket, size)`

Reads exactly `size` bytes or returns an error.

### `recvLine(socket)`

Reads one CRLF/LF-terminated line and strips the trailing newline.

### `startTLS(socket, options?)`

Upgrades an existing stream to TLS.

For server-side upgrades, pass:

```oak
{ server: true, certFile: './cert.pem', keyFile: './key.pem' }
```

### `close(socket)`

Closes the stream.

## Server API

### `listen(address, onConnect, options?)`

Starts a TCP listener and calls `onConnect(evt)` for each accepted connection.

Event shape:

```oak
{
    type: :connect
    socket: {type: :socket, id: 1}
    remote: '127.0.0.1:50000'
    local: '127.0.0.1:2525'
    tls: false
}
```

Returns a close function.

If `options.tls = true`, the listener uses implicit TLS from accept time and requires `certFile` and `keyFile`.

## Wrapper

### `Socket(socket)`

Returns an object-oriented wrapper around a raw socket value.

## Example

```oak
socket := import('socket')
thread := import('thread')

closeServer := socket.listen('127.0.0.1:9099', fn(evt) if evt.type {
    :connect -> thread.spawn(fn() {
        socket.sendLine(evt.socket, 'hello')
        socket.close(evt.socket)
    })
    _ -> ?
})

client := socket.connect('127.0.0.1:9099')
line := socket.recvLine(client.socket)
println(line.data)
closeServer()
```