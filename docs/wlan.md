# WLAN Library (WLAN)

## Overview

`libWLAN` provides local-network game discovery and announcement over TCP. A game host runs a **Beacon** that advertises its session on a known port. Clients **sweep** or continuously **scan** the local subnet for active Beacons, enabling LAN multiplayer lobby discovery without a central server.

Because Oak's networking primitives are TCP-based, WLAN uses a lightweight TCP probe protocol rather than UDP multicast. Beacons respond instantly to probe connections with a two-line payload (magic header + JSON metadata) and close the connection.

## Import

```oak
wlan := import('WLAN')

// or destructure specific functions
{
    Beacon: Beacon
    probe: probe
    sweep: sweep
    Scanner: Scanner
    localPrefix: localPrefix
} := import('WLAN')
```

## Constants

### `MAGIC`

The protocol magic string (`'MAGNOLIA_WLAN'`). Sent as the first line of every beacon response to distinguish WLAN beacons from other TCP services.

### `DEFAULT_PORT`

Default beacon port (`19800`).

## Functions

### `localPrefix()`

Detects the local network subnet prefix (e.g., `"192.168.1"`). Runs `ipconfig` on Windows or `ip addr`/`ifconfig` on Unix. Falls back to `"192.168.1"` if detection fails.

```oak
prefix := wlan.localPrefix()
// => "192.168.1"
```

### `Beacon(opts)`

Starts a TCP beacon that advertises a game session. Any incoming connection receives the game info and is closed immediately.

**Options:**

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `port` | int | `19800` | TCP port to listen on |
| `name` | string | `'Unknown Game'` | Display name for the session |
| `meta` | object | `{}` | Arbitrary metadata (map, players, maxPlayers, gamePort, etc.) |

**Returns** a beacon controller, or an error object if the listener fails to start.

```oak
beacon := wlan.Beacon({
    port: 19800
    name: 'Arena Match'
    meta: {
        map: 'castle'
        players: 1
        maxPlayers: 4
        gamePort: 9411
    }
})
```

#### Beacon controller

| Method | Description |
|--------|-------------|
| `info()` | Returns `{ name, meta }` snapshot |
| `update(patch)` | Merges `patch.name` and/or `patch.meta` into the beacon info |
| `close()` | Stops the beacon listener |

```oak
beacon.update({ meta: { players: 3 } })
beacon.close()
```

### `probe(host, port?)`

Checks a single `host:port` for an active WLAN beacon. Returns game info on success, `?` (null) on failure.

```oak
game := wlan.probe('192.168.1.5', 19800)
if game != ? -> println(game.name + ' (' + string(game.meta.players) + '/' + string(game.meta.maxPlayers) + ')')
```

**Returns** (on success):

```oak
{
    host: '192.168.1.5'
    port: 19800
    name: 'Arena Match'
    meta: { map: 'castle', players: 1, maxPlayers: 4, gamePort: 9411 }
}
```

### `sweep(opts?)`

Performs a one-shot parallel scan of a subnet range. Returns a list of discovered games.

**Options:**

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `port` | int | `19800` | Beacon port to probe |
| `prefix` | string | auto-detect | Subnet prefix, e.g., `"192.168.1"` |
| `start` | int | `1` | First host octet |
| `end` | int | `254` | Last host octet |
| `workers` | int | `32` | Number of parallel probe goroutines |

```oak
games := wlan.sweep({ port: 19800 })
games |> each(fn(g) println(g.name + ' at ' + g.host))
```

### `Scanner(opts?)`

Starts continuous background scanning. Reports game sessions as they appear or disappear through callbacks.

**Options** (same as `sweep`, plus):

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| `interval` | number | `3` | Seconds between scan cycles |
| `onFound` | fn(game) | no-op | Called when a new game is discovered |
| `onLost` | fn(game) | no-op | Called when a game disappears |
| `onSweep` | fn(games) | no-op | Called after each sweep completes |

```oak
scanner := wlan.Scanner({
    port: 19800
    onFound: fn(game) println('+ ' + game.name + ' at ' + game.host)
    onLost: fn(game) println('- ' + game.name)
})

// Check current games at any time
scanner.games()

// Stop scanning
scanner.close()
```

#### Scanner controller

| Method | Description |
|--------|-------------|
| `games()` | Returns a snapshot list of currently known games |
| `close()` | Stops the background scan loop |

## Protocol

The beacon protocol is intentionally minimal:

1. Client opens a TCP connection to the beacon port
2. Beacon sends: `MAGNOLIA_WLAN\n` (magic line)
3. Beacon sends: `<JSON>\n` (game info as a single JSON line)
4. Beacon closes the connection

The magic line lets probes reject non-WLAN services quickly.

## Integration with P2P

WLAN handles **discovery** — finding which games exist on the LAN. For the actual game communication, pair it with the `p2p` library:

```oak
wlan := import('WLAN')
p2p := import('p2p')

// --- Host side ---
// Start a P2P relay host for game traffic
host := p2p.Host('0.0.0.0:9411', '/game')

// Advertise the game so others can find it
beacon := wlan.Beacon({
    name: 'My Game Room'
    meta: {
        gamePort: 9411
        gamePath: '/game'
        map: 'forest'
        players: 1
        maxPlayers: 8
    }
})

// --- Client side ---
// Discover games on the LAN
games := wlan.sweep()
if len(games) > 0 -> {
    game := games.0
    url := 'ws://' + game.host + ':' + string(game.meta.gamePort) + game.meta.gamePath
    peer := p2p.join(url, 'player-2', fn(evt) println(evt))
}
```

## Full Example

```oak
{
    each: each
    default: default
} := import('std')

wlan := import('WLAN')

// Host: advertise a game
beacon := wlan.Beacon({
    port: 19800
    name: 'Dungeon Crawl'
    meta: {
        map: 'catacombs'
        players: 1
        maxPlayers: 4
        gamePort: 9411
    }
})

println('Beacon active on port 19800')

// Client: find games
wait(0.5)
games := wlan.sweep({ port: 19800, prefix: '127.0.0' })
games |> each(fn(g) {
    println('Found: ' + g.name + ' at ' + g.host + ':' + string(g.port))
    println('  Map: ' + default(g.meta.map, '?'))
    println('  Players: ' + string(default(g.meta.players, 0)) +
            '/' + string(default(g.meta.maxPlayers, 0)))
})

beacon.close()
```

## Notes

- Scanning speed depends on the local network. On a typical LAN, unreachable hosts fail within 1–3 seconds (ARP timeout). The parallel worker pool (`workers` option) keeps total scan time reasonable.
- The default port `19800` is arbitrary. Choose a port that doesn't conflict with other services. Using the same port across all instances of your game ensures discovery works.
- `localPrefix()` prefers private network addresses (192.168.x, 10.x, 172.16–31.x). On machines with multiple interfaces, it may pick a different subnet than intended — pass `prefix` explicitly if needed.
- The beacon listens on `0.0.0.0` (all interfaces) by default, so it is reachable from any local network the host is connected to.
