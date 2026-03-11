# Crypto Library (crypto)

## Overview

`libcrypto` provides utilities for working with cryptographic primitives and cryptographically secure sources of randomness. The library focuses on practical cryptographic functions for session tokens, unique identifiers, secure random values, and secret comparison.

## Import

```oak
crypto := import('crypto')
// or destructure specific functions
{ uuid: uuid, randomBytes: randomBytes, randomInt: randomInt, sha256: sha256 } := import('crypto')
```

## Functions

### `uuid()`

Generates a random UUID v4 (Universally Unique Identifier) using cryptographically secure randomness, compliant with RFC 4122 §4.4.

UUID v4 format: `xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx`

- All `x` positions are random hexadecimal digits (0-9, a-f)
- The `4` indicates version 4 (random)
- The `y` is one of 8, 9, a, or b (variant bits per RFC 4122)

```oak
{ uuid: uuid } := import('crypto')

id1 := uuid()
// => "ebe94230-d33a-4d38-bfb2-8b31226d2f4f" (example - value varies)

id2 := uuid()
// => "a04e77d8-1b9c-487b-b5b8-ac7c867ad129" (example - value varies)

// Each call generates a unique ID
id1 != id2 // => true
```

### `randomBytes(n)`

Returns n cryptographically secure random bytes encoded as a lowercase hexadecimal string (2n characters). Suitable for tokens, nonces, salts, and any use case requiring verifiable entropy.

Parameters:
- `n` (int): Number of random bytes to generate

Returns: Lowercase hex string of 2n characters

```oak
{ randomBytes: randomBytes } := import('crypto')

token := randomBytes(16)
// => "96d90e9995801c315b66d1474ebcd3da" (example - 32 hex chars)

nonce := randomBytes(8)
// => "4563f6b02e2f2f79" (example - 16 hex chars)
```

### `randomInt(min, max)`

Returns a cryptographically secure integer in the half-open interval [min, max). Uses 4 bytes of OS entropy to generate the value.

Parameters:
- `min` (int): Lower bound (inclusive)
- `max` (int): Upper bound (exclusive)

Returns: Random integer where min ≤ result < max

**Note:** Slight modulo bias exists for ranges that do not divide 2^32 evenly, but is negligible in practice for most applications.

```oak
{ randomInt: randomInt } := import('crypto')

// Fair six-sided die roll (1-6)
roll := randomInt(1, 7)
// => 4 (example)

// Random percentage
percent := randomInt(0, 101)
// => 73 (example)
```

### `randomFloat()`

Returns a cryptographically secure float uniformly distributed in [0.0, 1.0). Uses 53 bits of entropy to fill an IEEE 754 double-precision mantissa.

Returns: Float in range [0.0, 1.0)

```oak
{ randomFloat: randomFloat } := import('crypto')

value := randomFloat()
// => 0.6610978990802057 (example)

// Scale to desired range
scaled := randomFloat() * 100.0
// => 22.567810370679853 (example - 0-100 range)
```

### `timingSafeEqual(a, b)`

Compares two strings byte-by-byte without short-circuiting, reducing timing side-channel leakage when comparing secret values such as MAC tags or derived keys.

Parameters:
- `a` (string): First value
- `b` (string): Second value

Returns: `true` if a equals b, `false` otherwise

**Note:** If the lengths of a and b differ, the function returns `false` without inspecting the bytes, leaking the length difference. This is acceptable for fixed-length outputs (e.g., hex-encoded hashes or MAC tags from symmetric algorithms).

```oak
{ timingSafeEqual: timingSafeEqual } := import('crypto')

correctSecret := 'mySecretValue'
userInput := 'mySecretValue'

if timingSafeEqual(correctSecret, userInput) {
	true -> println('Secret is correct')
	_ -> println('Secret is incorrect')
}

// Constant-time comparison prevents timing attacks on secrets
wrongInput := 'wrongSecret'
timingSafeEqual(correctSecret, wrongInput) // => false (same time as correct match)
```

### `sha256(data)`

Computes SHA-256 over a byte string and returns a 64-character lowercase hexadecimal digest.

Parameters:
- `data` (string): Input byte string to hash

Returns: 64-character lowercase hex digest

```oak
{ sha256: sha256 } := import('crypto')

sha256('')
// => "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

sha256('hello')
// => "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824"
```

## Implementation Details

### Entropy Source

All functions use Oak's built-in `srand()` function, which provides cryptographically secure random bytes using the operating system's entropy source (e.g., `/dev/urandom` on Unix, `CryptGenRandom` on Windows).

### UUID v4 Generation

The `uuid()` function:
1. Generates 16 secure random bytes using a single `srand(16)` call
2. Sets the version bits (byte 6, high nibble) to 0x4, indicating UUID v4
3. Sets the variant bits (byte 8, top two bits) to 0b10, per RFC 4122
4. Formats the bytes as a hyphenated hexadecimal string

### Hex Encoding

Byte-to-hex conversion uses a lookup table for efficiency with two operations per byte (high and low nibbles).

### SHA-256

`sha256` implements the full SHA-256 compression pipeline in Oak:
1. Message padding and 64-bit length encoding (big-endian)
2. 64-word schedule expansion
3. 64 compression rounds with RFC constants
4. Final state serialization to lowercase hexadecimal

## Use Cases

### Session Tokens

```oak
{ uuid: uuid } := import('crypto')

sessionId := uuid()  // Globally unique, unpredictable session ID
```

### CSRF Token Generation

```oak
{ randomBytes: randomBytes } := import('crypto')

csrfToken := randomBytes(32)  // 64-character hex string for CSRF protection
```

### Cryptographic Nonces

```oak
{ randomBytes: randomBytes } := import('crypto')

nonce := randomBytes(12)  // For use in authenticated encryption schemes
```

### Secure Random Selection

```oak
{ randomInt: randomInt } := import('crypto')

items := ['apple', 'banana', 'cherry', 'date']
chosen := items.(randomInt(0, len(items)))
```

### Monte Carlo Simulations

```oak
{ randomFloat: randomFloat } := import('crypto')

samples := 100000
insideCircle := 0
each(range(samples), fn(_) {
	x := randomFloat()
	y := randomFloat()
	if (x * x + y * y) < 1.0 -> insideCircle <- insideCircle + 1
})
```

### HMAC Verification

```oak
{ timingSafeEqual: timingSafeEqual } := import('crypto')

// Compare computed HMAC with received HMAC
if timingSafeEqual(computedMac, receivedMac) {
	true -> println('Signature is valid')
	_ -> println('Signature verification failed')
}
```

## Security Considerations

- **All functions use cryptographically secure randomness** from `srand()`, making them suitable for security-critical applications.
- **`timingSafeEqual` prevents timing attacks** on fixed-length secrets by avoiding early exit on mismatch.
- **Variable-length secrets** should have their length protected separately (e.g., constant-length encoding or length-oblivious comparison).
- **Output size matters**: Larger `n` in `randomBytes(n)` provides proportionally more entropy. Typical secure tokens use 16-32 bytes (128-256 bits).

## Related Functions

For non-cryptographic random values, see the `random` library.

## Notes

- UUID v4 is the most common UUID variant, defined in RFC 4122
- UUIDs are 128 bits (16 bytes) total, represented as 36-character strings (32 hex digits + 4 hyphens)
- Guaranteed to be unique across space and time (with astronomical probability)
- All hex output is lowercase for consistency

