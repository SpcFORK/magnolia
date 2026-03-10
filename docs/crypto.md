# Crypto Library (crypto)

## Overview

`libcrypto` provides utilities for working with cryptographic primitives and cryptographically secure sources of randomness. Currently focuses on UUID generation using secure random bytes.

## Import

```oak
crypto := import('crypto')
// or destructure specific functions
{ uuid: uuid } := import('crypto')
```

## Functions

### `uuid()`

Generates a random UUID v4 (Universally Unique Identifier) using cryptographically secure randomness.

UUID v4 format: `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx`

- All `x` positions are random hexadecimal digits (0-9, a-f)
- The `4` indicates version 4
- The `y` is one of 8, 9, a, or b (variant bits)

```oak
{ uuid: uuid } := import('crypto')

id1 := uuid()
// => e.g., "a3bb189e-8bf9-3b44-9e63-6319c3b5b9a7"

id2 := uuid()
// => e.g., "f47ac10b-58cc-4372-a567-0e02b2c3d479"

// Each call generates a unique ID
id1 != id2 // => true
```

## Implementation Details

The `uuid()` function:

1. Generates 16 secure random bytes using Oak's `srand()` builtin
2. Sets the version bits (byte 6) to indicate UUID v4
3. Sets the variant bits (byte 8) according to RFC 4122
4. Formats the bytes as a hyphenated hexadecimal string

### Version and Variant Bits

```oak
// Pseudocode of the algorithm:
bytes := srand(16)  // 16 random bytes

// UUID v4 version bits
bytes[6] := (bytes[6] & 0x0F) | 0x40  // Version 4

// RFC 4122 variant bits
bytes[8] := (bytes[8] & 0x3F) | 0x80  // Variant 1

// Format as: xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx
```

## Examples

### Generating Unique Identifiers

```oak
{ uuid: uuid } := import('crypto')

// Generate a session ID
sessionId := uuid()
// => "550e8400-e29b-41d4-a716-446655440000"

// Create multiple unique IDs
ids := [uuid(), uuid(), uuid()]
// => ["...", "...", "..."]
```

### Database Key Generation

```oak
{ uuid: uuid } := import('crypto')

fn createUser(name, email) {
    {
        id: uuid()
        name: name
        email: email
        createdAt: time()
    }
}

user := createUser('Alice', 'alice@example.com')
// => {
//   id: "7c9e6679-7425-40de-944b-e07fc1f90ae7"
//   name: "Alice"
//   email: "alice@example.com"
//   createdAt: 1234567890
// }
```

### API Request Tracking

```oak
{ uuid: uuid } := import('crypto')

fn makeRequest(endpoint, data) {
    requestId := uuid()
    println('Request ' + requestId + ' to ' + endpoint)
    
    // ... make HTTP request ...
    
    {
        requestId: requestId
        response: responseData
    }
}
```

### Collision Probability

UUID v4 uses 122 random bits (after accounting for version and variant bits). The probability of collision is extremely low:

- 1 in 2^122 for any two UUIDs
- Need ~2.71 quintillion UUIDs for a 50% collision chance
- Practically collision-free for most applications

## Security Considerations

- Uses Oak's `srand()` builtin for cryptographically secure random bytes
- Suitable for security-sensitive applications like session tokens
- Do not use for cryptographic keys—use dedicated key derivation functions
- UUIDs are not secret—they're unique identifiers, not passwords

## Related Functions

The library uses these imports from other Oak libraries:

- `std.toHex()` - Convert bytes to hexadecimal
- `std.map()` - Process byte arrays
- `str.split()` - Split random bytes into characters

For non-cryptographic random values, see the `random` library.

## Notes

- UUID v4 is the most common UUID variant
- The format is defined in RFC 4122
- UUIDs are 128 bits (16 bytes) total
- represented as 36-character strings (32 hex digits + 4 hyphens)
- Case-insensitive (lowercase is canonical)
- Guaranteed to be unique across space and time (with astronomical probability)
