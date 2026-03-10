# Bitwise and Pointer Helpers (bitwise)

## Overview

`libbitwise` wraps low-level integer bit operations and pointer/memory primitives into reusable helpers.

This library is useful when you need:

- predictable bit mask operations
- explicit pointer arithmetic
- controlled memory reads/writes over byte buffers

Import path: `bitwise`

## Import

```oak
bitwise := import('bitwise')

{
    and: and
    or: or
    xor: xor
    not: not
    shl: shl
    shr: shr
    hasAll?: hasAll?
    hasAny?: hasAny?
    set: set
    clear: clear
    toggle: toggle
    update: update
    ptr: ptr
    null: null
    null?: null?
    addrOf: addrOf
    add: add
    sub: sub
    diff: diff
    read: read
    write: write
} := import('bitwise')
```

## Integer Bit Operations

### `and(a, b)`

Bitwise AND.

### `or(a, b)`

Bitwise OR.

### `xor(a, b)`

Bitwise XOR.

### `not(n)`

Bitwise NOT (implemented as `n ^ -1`).

### `shl(n, by)`

Shift left by `by` bits.

### `shr(n, by)`

Shift right by `by` bits.

```oak
{
    and: and
    or: or
    xor: xor
    not: not
    shl: shl
    shr: shr
} := import('bitwise')

println(and(13, 11)) // 9
println(or(13, 11))  // 15
println(xor(13, 11)) // 6
println(not(10))     // -11
println(shl(3, 4))   // 48
println(shr(48, 4))  // 3
```

## Mask Helpers

### `hasAll?(value, mask)`

Returns true when all mask bits are set in value.

### `hasAny?(value, mask)`

Returns true when any mask bit is set in value.

### `set(value, mask)`

Sets mask bits in value.

### `clear(value, mask)`

Clears mask bits in value.

### `toggle(value, mask)`

Flips mask bits in value.

### `update(value, mask, enabled)`

Calls `set` when `enabled` is true, else `clear`.

```oak
{
    hasAll?: hasAll?
    hasAny?: hasAny?
    set: set
    clear: clear
    toggle: toggle
    update: update
} := import('bitwise')

value := 10 // 1010
mask := 6   // 0110

println(hasAll?(value, mask))   // false
println(hasAny?(value, mask))   // true
println(set(value, mask))       // 14
println(clear(value, mask))     // 8
println(toggle(value, mask))    // 12
println(update(value, mask, true))  // 14
println(update(value, mask, false)) // 8
```

## Pointer and Memory Helpers

### `null()` and `null?(p)`

Creates/checks the null pointer.

### `addrOf(data)`

Returns a pointer to the first byte in a byte string.

### `add(p, offset)` and `sub(p, offset)`

Pointer arithmetic in bytes.

### `diff(a, b)`

Byte distance between pointers.

### `read(p, length)`

Reads bytes from memory and returns a byte string.

### `write(p, data)`

Writes bytes to memory. Accepts either byte string or list of bytes.

```oak
{
    bits: bits
    string: string
} := import('std')

{
    addrOf: addrOf
    add: add
    diff: diff
    read: read
    write: write
} := import('bitwise')

buf := bits([65, 66, 67, 68]) // ABCD
base := addrOf(buf)

println(diff(add(base, 3), base)) // 3
println(string(read(base, 4)))    // ABCD

write(add(base, 1), [90, 89])
println(string(read(base, 4)))    // AZYD
```

## Notes

- `shl` and `shr` are provided as reliable wrappers when direct shift operators are inconvenient.
- Pointer and memory APIs are unsafe by design: invalid addresses can crash the process.
- Keep source byte buffers alive while operating on their addresses.

## Related

- [Go runtime and host interop](./go.md)
- [Language spec](./spec.md)
- [Sample: pointers + bits](../samples/pointers-bits.oak)
