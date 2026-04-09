# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\smtp.oak`

- `thread` · `import(...)`
- `sock` · `import(...)`
### `_okLine(code, text)`

### `_moreLine(code, text)`

### `_resp(code, lines)`

> returns `:object`

### `_readResponse(socket)`

### `_textLines(response)`

### `_dotStuff(body)`

### `_renderMessage(message)`

### `_smtpClient(socket, greeting, options)`

> returns `:object`

### `connect(address, options)`

> returns `:object`

### `_pathValue(line, prefix)`

### `_sessionState(socket, handlers, options)`

> returns `:object`

### `_smtpReply(socket, code, text)`

### `_smtpMultiline(socket, code, lines)`

### `_handleData(socket)`

### `_serveClient(socket, handlers, options)`

### `listen(address, handlers, options)`

