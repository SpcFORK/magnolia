# gui-common

Shared utility helpers used across all Magnolia GUI submodules.

Centralises `_default`, `_err`, and `_clamp` so each GUI submodule can import
them from one place instead of redefining them locally.

## Functions

### `_default(value, fallback)`

Null-coalescing helper. Returns `fallback` when `value` is `?`,
otherwise returns `value`.

### `_err(message, detail)`

Constructs a standard error object:

```oak
{
  type: :error
  error: message
  detail: detail
}
```

### `_clamp(v, minV, maxV)`

Clamps `v` to the range `[minV, maxV]`.
