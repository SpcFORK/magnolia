# Writes Library (writes)

## Overview

`writes` provides shared little-endian integer read/write helpers over raw memory.

It is intended for:

- writing portable typed memory helpers once
- reusing the same integer encoding behavior across platform interop modules
- reducing copy-pasted byte packing logic

Import path: `writes`

## Import

```oak
writes := import('writes')
```

## Functions

### `readU32(address)`

Reads 4 bytes at `address` and returns an unsigned 32-bit little-endian integer.

### `writeU32(address, value)`

Writes `value` as 4 little-endian bytes at `address`.

### `readU64(address)`

Reads 8 bytes at `address` and returns an unsigned 64-bit little-endian integer.

### `writeU64(address, value)`

Writes `value` as 8 little-endian bytes at `address`.

## Example

```oak
writes := import('writes')

buf := bits([0, 0, 0, 0, 0, 0, 0, 0])
ptr := addr(buf)

writes.writeU32(ptr, 305419896) // 0x12345678
println(writes.readU32(ptr))    // 305419896

writes.writeU64(ptr, 72623859790382856) // 0x0102030405060708
println(writes.readU64(ptr))             // 72623859790382856
```

## Byte Order

`writes` uses little-endian encoding:

- `writeU32(..., 0x12345678)` writes bytes `[0x78, 0x56, 0x34, 0x12]`
- `writeU64(..., 0x0102030405060708)` writes bytes `[0x08, 0x07, 0x06, 0x05, 0x04, 0x03, 0x02, 0x01]`

## Notes

- These helpers rely on `memread` and `memwrite`; invalid pointers can crash the process.
- Keep backing buffers alive while pointers are in use.
- Values are treated as non-negative integers for byte packing.

## Related

- [Go Runtime and System Interop](./go.md)
- [Windows Interop Library](./windows.md)
- [Linux Interop Library](./linux.md)
