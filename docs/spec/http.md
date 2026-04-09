# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\http.oak`

- `sort` · `import(...)`
- `json` · `import(...)`
### `queryEncode(params)`

### `queryDecode(params)`

### `_encodeChar(uri?)`

> **thunk** returns `:function`

### `percentEncode(s)`

### `percentEncodeURI(s)`

### `_hex?(c)`

> returns `:bool`

### `percentDecode(s)`

### `cs Router()`

> returns `:object`

- `MimeTypes` · `{17 entries}`
### `mimeForPath(path)`

- `NotFound` · `{2 entries}`
- `MethodNotAllowed` · `{2 entries}`
### `_hdr(attrs)`

### `cs Server()`

> returns `:object`

### `handleStaticUnsafe(path)`

> **thunk** returns `:function`

### `handleStatic(path)`

### `pbatchQueryEncode(paramSets)`

### `pbatchQueryDecode(queryStrings)`

