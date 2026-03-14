# P2P Library (p2p)

## Overview

`libp2p` provides a small relay-style peer messaging layer on top of Magnolia's WebSocket helpers.

This is an application-level mesh coordinator, not raw socket hole punching or NAT traversal. One host accepts peer connections, tracks membership, and relays peer messages.

## Import

```oak
p2p := import('p2p')
```

## Host API

### `Host(host, path, onEvent?)`

Starts a relay host on `host` and `path`.

Example:

```oak
p2p := import('p2p')

host := p2p.Host('127.0.0.1:9411', '/mesh', fn(evt) {
    println(evt)
})
```

Returns an object with:

- `type = :host`
- `host`, `path`
- `peers()` → current peer summaries
- `send(peerId, payload, channel?)`
- `broadcast(payload, channel?)`
- `close()`

Host events include:

- `{type: :peer_joined, peerId: 'alice', meta: {...}}`
- `{type: :peer_left, peerId: 'alice', reason: '...'}`
- `{type: :message, mode: :broadcast|:direct, from: 'alice', ...}`
- `{type: :error, scope: :host, error: '...'}`

## Peer API

### `join(url, peerId, onEvent?, meta?)`

Connects a peer to a host.

Success returns a controller object:

```oak
{
    type: :peer
    peerId: 'alice'
    url: 'ws://127.0.0.1:9411/mesh'
    connected?: fn() ...
    ready?: fn() ...
    knownPeers: fn() ...
    send: fn(toPeerId, payload, channel?) ...
    broadcast: fn(payload, channel?) ...
    close: fn(reason?) ...
}
```

Failure returns the WebSocket error object:

```oak
{
    type: :error
    error: '<message>'
    status: 0
    headers: {}
}
```

Peer events include:

- `{type: :ready, peerId: 'alice', peers: [...]}`
- `{type: :peer_joined, peerId: 'bob', meta: {...}}`
- `{type: :peer_left, peerId: 'bob', reason: '...'}`
- `{type: :message, from: 'bob', channel: 'chat', payload: {...}}`
- `{type: :error, scope: :peer|:remote, error: '...'}`
- `{type: :closed, peerId: 'alice', code: 1000, reason: '...'}`

## Protocol

Peers exchange JSON packets over WebSockets:

- `hello` for registration
- `broadcast` for fan-out messages
- `direct` for one-peer messages
- host-emitted `welcome`, `peer-joined`, `peer-left`, `peer-message`, and `error`

## Example

```oak
{
    wait: wait
} := import('std')

p2p := import('p2p')

host := p2p.Host('127.0.0.1:9411', '/mesh')
alice := p2p.join('ws://127.0.0.1:9411/mesh', 'alice')
bob := p2p.join('ws://127.0.0.1:9411/mesh', 'bob')

wait(0.2)
alice.broadcast({ text: 'hello everyone' }, 'chat')
bob.send('alice', { text: 'private hello' }, 'dm')

wait(0.2)
host.close()
```

## Multi-Process CLI Sample

Use [samples/p2p-cli.oak](samples/p2p-cli.oak) to run a host and peers in separate terminals.

Host terminal:

```sh
magnolia samples/p2p-cli.oak host --port 9411 --path /mesh
```

Peer terminal (broadcast):

```sh
magnolia samples/p2p-cli.oak peer --id alice --url ws://127.0.0.1:9411/mesh --broadcast --text "hello room"
```

Peer terminal (direct):

```sh
magnolia samples/p2p-cli.oak peer --id bob --url ws://127.0.0.1:9411/mesh --to alice --text "hello alice"
```