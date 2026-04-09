# MessagePack Library (msgpack)

## Overview

`libmsgpack` encodes Oak values into MessagePack wire bytes.

MessagePack is a compact binary serialization format that is similar to JSON in
structure but smaller and faster to parse on many workloads. Instead of text,
it uses byte tags and packed payloads.

This module is focused on serialization output, either as:

- a byte string (for files/sockets)
- a packet list (`[0..255]` integers) for inspection or transport adapters

## Import

```oak
msgpack := import('msgpack')
{ serialize: serialize, serializePacket: serializePacket } := import('msgpack')
```

## Functions

### `serializeSafe(value)`

Encodes `value` and returns a MessagePack byte string.

Returns `:error` when the value cannot be encoded by this implementation.

### `serialize(value)`

Encodes `value` and returns a MessagePack byte string.

If the value is unsupported, this falls back to MessagePack `nil` (`0xC0`).

### `serializePacket(value)`

Encodes `value` and returns a list of bytes (`[int]`) instead of a string.

Returns `:error` for unsupported values.

### `packet(byteString)`

Converts a byte string to a byte list.

### `fromPacket(parts)`

Converts a byte list to a byte string.

### `Binary(data)`

Wraps a byte string so it is encoded using MessagePack `bin` family (`bin8`,
`bin16`, `bin32`) instead of UTF-8 string tags.

### `binary?(value)`

Returns true when `value` is a `Binary(...)` wrapper object.

## Supported Value Types

- `?` -> `nil`
- `bool` -> `true` / `false`
- `int` -> fixint/int8/int16/int32/uint8/uint16/uint32
- `string` -> fixstr/str8/str16/str32
- `atom` -> encoded as string
- `list` -> fixarray/array16/array32
- `object` -> fixmap/map16/map32 (keys encoded as strings)
- `Binary(data)` -> bin8/bin16/bin32

## Current Limitations

- `float` values are not encoded yet (`serializeSafe` returns `:error`)
- integer support is currently bounded to signed 32-bit / unsigned 32-bit range
- map/object key order is not guaranteed
- no MessagePack deserializer yet

## Examples

### Basic serialization

```oak
msgpack := import('msgpack')

raw := msgpack.serialize({ kind: 'ping', id: 7, ok: true })
packet := msgpack.packet(raw)

// packet starts with a map header and can be sent over sockets/files
```

### Binary payload packet

```oak
msgpack := import('msgpack')

payload := '\x01\x02\x03\x04'
wrapped := msgpack.Binary(payload)

bytes := msgpack.serializePacket({ type: 'blob', data: wrapped })
```

### Safe mode vs fallback mode

```oak
msgpack := import('msgpack')

msgpack.serializeSafe(3.14) // => :error
msgpack.serialize(3.14)     // => byte string [0xC0]
```

## Parallel Batch Operations

### `pbatchSerialize(values)`

Serializes a list of values into MessagePack byte strings in parallel.

```oak
msgpack.pbatchSerialize([{a: 1}, {b: 2}, 'hello'])
```

### `pbatchParse(packets)`

Parses multiple MessagePack byte strings in parallel.

```oak
msgpack.pbatchParse([packed1, packed2])
// => [value1, value2]
```
