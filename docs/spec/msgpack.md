# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\msgpack.oak`

### `bytes(parts)`

### `packet(data)`

### `fromPacket(parts)`

### `Binary(data)`

> returns `:object`

### `binary?(value)`

> returns `:bool`

### `_u8(n)`

> returns `:list`

### `_u16be(n)`

> returns `:list`

### `_u32be(n)`

> returns `:list`

### `_encodeString(s)`

> returns `:atom`

### `_encodeBinary(data)`

> returns `:atom`

### `_encodeInt(n)`

> returns `:list`

### `_encodeList(xs)`

> returns `:atom`

### `_encodeObject(obj)`

> returns `:atom`

### `_encodeToPacket(value)`

### `_f64ToPacket(f)`

> returns `:list`

### `_f32ToPacket(f)`

> returns `:list`

### `_encodeFloat64(f)`

### `encodeFloat64(f)`

### `encodeFloat32(f)`

### `_readU8(pkt, i)`

### `_readU16(pkt, i)`

> returns `:bool`

### `_readU32(pkt, i)`

> returns `:bool`

### `_readI8(pkt, i)`

> returns `:atom`

### `_readI16(pkt, i)`

> returns `:atom`

### `_readI32(pkt, i)`

> returns `:atom`

### `_readBits64(pkt, i)`

### `_ieee754f64(bs)`

> returns `:float`

### `_ieee754f32(bs)`

> returns `:float`

### `_decode(pkt, i)`

> returns `:atom`

### `_decodeStr(pkt, i, n)`

> returns `:atom`

### `_decodeBin(pkt, i, n)`

> returns `:atom`

### `_decodeArray(pkt, i, n)`

### `_decodeMap(pkt, i, n)`

### `_slicePkt(pkt, a, b)`

### `parsePacket(pkt)`

> returns `:atom`

### `parse(data)`

### `serializeSafe(value)`

> returns `:atom`

### `serialize(value)`

### `serializePacket(value)`

> returns `:atom`

### `pbatchSerialize(values)`

### `pbatchParse(packets)`

