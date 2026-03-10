# WebSocket Library (websocket)

## Overview

`libwebsocket` provides WebSocket client and server helpers for Magnolia.

It is a thin wrapper around native runtime built-ins:

- `ws_dial`
- `ws_send`
- `ws_recv`
- `ws_close`
- `ws_listen`

## Import

```oak
ws := import('websocket')
```

## Client API

### `connect(url, headers?)`

Connects to a WebSocket endpoint.

Success result:

```oak
{
    type: :ok
    socket: {type: :websocket, id: 0}
    status: 101
    headers: {...}
}
```

Error result:

```oak
{
    type: :error
    error: '<message>'
    status: 0
    headers: {}
}
```

### `send(socket, data, opcode?)`

Sends a message.

- `data` must be a string.
- `opcode` defaults to text (`1`).

Returns `{type: :sent}` on success, or `{type: :error, error: ...}`.

### `recv(socket)`

Reads one message.

Message result:

```oak
{
    type: :message
    opcode: 1
    data: 'hello'
}
```

Closed result:

```oak
{
    type: :closed
    code: 1000
    reason: ''
}
```

Error result:

```oak
{
    type: :error
    error: '<message>'
}
```

### `close(socket, code?, reason?)`

Closes the socket. Defaults to close code `1000`.

## Server API

### `listen(host, path, onConnect)`

Starts a WebSocket server on `host` and upgrades requests on `path`.

`onConnect` receives one event per successful upgrade:

```oak
{
    type: :connect
    socket: {type: :websocket, id: 1}
    req: {
        method: 'GET'
        url: '/ws'
        headers: {...}
    }
}
```

On startup/upgrade errors it receives:

```oak
{type: :error, error: '<message>'}
```

Returns a close function that gracefully shuts down the server.

## Constants

### `Opcode`

- `Opcode.text = 1`
- `Opcode.binary = 2`
- `Opcode.close = 8`
- `Opcode.ping = 9`
- `Opcode.pong = 10`

### `CloseCode`

- `CloseCode.normal = 1000`
- `CloseCode.goingAway = 1001`
- `CloseCode.protocolError = 1002`
- `CloseCode.unsupportedData = 1003`
- `CloseCode.invalidFramePayloadData = 1007`
- `CloseCode.policyViolation = 1008`
- `CloseCode.messageTooBig = 1009`
- `CloseCode.internalServerError = 1011`

## Example

```oak
ws := import('websocket')

result := ws.connect('ws://echo.websocket.org')
if result.type = :ok -> {
    sock := result.socket
    ws.send(sock, 'hello from magnolia')
    msg := ws.recv(sock)
    println(msg)
    ws.close(sock)
} else {
    println(result.error)
}
```
