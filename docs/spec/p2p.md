# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\p2p.oak`

- `thread` · `import(...)`
- `json` · `import(...)`
- `ws` · `import(...)`
### `_noop(_)`

> returns `?`

### `_packet(kind, body)`

### `_sendPacket(socket, kind, body)`

### `_terminalPeerError?(evt)`

> returns `:bool`

### `_decodePacket(evt)`

> returns `:object`

### `_findPeer(peers, peerId)`

### `_peerSummaries(peers, excludeId)`

### `_peerSockets(peers, excludeId)`

### `_fanout(state, mutex, excludeId, kind, body)`

### `_removePeer(state, mutex, peerId)`

### `_sendDirect(state, mutex, toPeerId, fromPeerId, channel, payload)`

### `_peerEvent(packetType, body)`

> returns `:object`

### `_hostError(onEvent, peerId, error)`

### `_hostLoop(state, mutex, peerId, socket, onEvent)`

### `_acceptPeer(state, mutex, host, path, socket, req, onEvent)`

### `Host(host, path, onEvent)`

> returns `:object`

### `join(url, peerId, onEvent, meta)`

