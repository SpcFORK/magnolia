# HTTP Library (http)

## Overview

`libhttp` provides utilities for building HTTP server applications in Oak, including routing, static file serving, query string handling, and URL encoding.

## Import

```oak
http := import('http')
// or destructure specific functions
{
    Server: Server
    Router: Router
    queryEncode: queryEncode
    queryDecode: queryDecode
} := import('http')
```

## Server API

### `Server()`

Creates an HTTP server application with routing capabilities.

**Methods:**

- `route(pattern, handler)` - Add a route handler for a path pattern
- `start(port)` - Start the server on the specified port
- `startThreaded(port)` - Start the server on the specified port and handle each request in a separate Oak thread

```oak
{ Server: Server } := import('http')

server := Server()

// Add routes
server.route('/', fn(params) fn(req, end) end({
    status: 200
    body: 'Hello, World!'
}))

server.route('/users/:id', fn(params) fn(req, end) end({
    status: 200
    body: 'User ID: ' + params.id
}))

// Start server on port 8080
server.start(8080)

// Or start with threaded request handling
// (each request is processed in a separate Oak thread)
server.startThreaded(8080)
```

## Router API

### `Router()`

Creates a router for mapping URL paths to handlers. Used internally by `Server()` but can be used standalone.

**Methods:**

- `add(pattern, handler)` - Add a route handler
- `catch(handler)` - Add a catch-all handler (empty pattern)
- `match(path)` - Match a path and return the appropriate handler

**Path Patterns:**

- Static segments: `/users/profile`
- Named parameters: `/users/:id` (captures value in `params.id`)
- Wildcard parameters: `/static/*path` (captures remaining path in `params.path`)
- Query parameters: Automatically parsed from `?key=value&...`

```oak
{ Router: Router } := import('http')

router := Router()

// Static route
router.add('/about', fn(params) fn(req, end) end({
    status: 200
    body: 'About page'
}))

// Named parameter
router.add('/users/:userId', fn(params) fn(req, end) end({
    status: 200
    body: 'User: ' + params.userId
}))

// Wildcard parameter
router.add('/files/*filePath', fn(params) fn(req, end) end({
    status: 200
    body: 'File: ' + params.filePath
}))

// Catch-all (404 handler)
router.catch(fn(params) fn(req, end) end({
    status: 404
    body: 'Not found'
}))

// Match a path
handler := router.match('/users/123')
// Returns the handler function that captures userId: '123'
```

### Route Pattern Matching

```oak
// URL: /users/alice
// Pattern: /users/:name
// Result: params.name = 'alice'

// URL: /static/css/main.css
// Pattern: /static/*path
// Result: params.path = 'css/main.css'

// URL: /search?q=oak&limit=10
// Pattern: /search
// Result: params.q = 'oak', params.limit = '10'

// URL: /api/v1/users/123/posts/456
// Pattern: /api/:version/users/:userId/posts/:postId
// Result: params.version = 'v1', params.userId = '123', params.postId = '456'
```

## Static File Serving

### `handleStatic(path)`

Returns a route handler for serving static files. Only responds to GET requests.

```oak
{ Server: Server, handleStatic: handleStatic } := import('http')

server := Server()

// Serve a single file
server.route('/', handleStatic('./public/index.html'))

// Serve directory with wildcard
with server.route('/static/*path') fn(params) {
    handleStatic('./public/' + params.path)
}

// Serve favicon
server.route('/favicon.ico', handleStatic('./public/favicon.ico'))

server.start(8080)
```

### `mimeForPath(path)`

Returns the appropriate MIME type for a file path based on its extension.

**Supported MIME Types:**

| Extension | MIME Type |
|-----------|-----------|
| `.html` | `text/html; charset=utf-8` |
| `.txt`, `.md` | `text/plain; charset=utf-8` |
| `.css` | `text/css; charset=utf-8` |
| `.js` | `application/javascript; charset=utf-8` |
| `.json` | `application/json; charset=utf-8` |
| `.jpg`, `.jpeg` | `image/jpeg` |
| `.png` | `image/png` |
| `.gif` | `image/gif` |
| `.svg` | `image/svg+xml` |
| `.webp` | `image/webp` |
| `.pdf` | `application/pdf` |
| `.zip` | `application/zip` |
| default | `application/octet-stream` |

```oak
{ mimeForPath: mimeForPath } := import('http')

mimeForPath('index.html') // => 'text/html; charset=utf-8'
mimeForPath('style.css') // => 'text/css; charset=utf-8'
mimeForPath('photo.jpg') // => 'image/jpeg'
mimeForPath('data.bin') // => 'application/octet-stream'
```

## URL Encoding

### `percentEncode(s)`

Encodes a string using percent-encoding (URI component encoding). Similar to JavaScript's `encodeURIComponent()`.

```oak
{ percentEncode: percentEncode } := import('http')

percentEncode('hello world') // => 'hello%20world'
percentEncode('a&b=c') // => 'a%26b%3Dc'
percentEncode('café') // => 'caf%C3%A9'

// Safe characters are not encoded
percentEncode('abc123-_.!~*\'()') // => 'abc123-_.!~*\'()'
```

### `percentEncodeURI(s)`

Encodes a string using percent-encoding but preserves URI-reserved characters. Similar to JavaScript's `encodeURI()`.

```oak
{ percentEncodeURI: percentEncodeURI } := import('http')

percentEncodeURI('https://example.com/path?q=hello world')
// => 'https://example.com/path?q=hello%20world'

// Preserves URI structure characters: ;,/?:@&=+$#
percentEncodeURI('path/to/resource') // => 'path/to/resource'
```

### `percentDecode(s)`

Decodes a percent-encoded string. Handles `+` as space and hex-encoded characters.

```oak
{ percentDecode: percentDecode } := import('http')

percentDecode('hello%20world') // => 'hello world'
percentDecode('hello+world') // => 'hello world'
percentDecode('a%26b%3Dc') // => 'a&b=c'
percentDecode('caf%C3%A9') // => 'café'
```

## Query String Handling

### `queryEncode(params)`

Converts an object to a URL query string. Composite values (lists, objects) are JSON-serialized. Functions are omitted.

```oak
{ queryEncode: queryEncode } := import('http')

queryEncode({ name: 'Alice', age: 30 })
// => 'age=30&name=Alice' (alphabetically sorted)

queryEncode({ q: 'hello world', limit: 10 })
// => 'limit=10&q=hello%20world'

queryEncode({ tags: ['oak', 'lang'], active: true })
// => 'active=true&tags=["oak","lang"]'

queryEncode({ key: ?, value: 'test' })
// => 'value=test' (null values omitted)
```

### `queryDecode(queryString)`

Parses a query string into an object. All values are strings (query strings are untyped).

```oak
{ queryDecode: queryDecode } := import('http')

queryDecode('name=Alice&age=30')
// => { name: 'Alice', age: '30' }

queryDecode('q=hello+world&limit=10')
// => { q: 'hello world', limit: '10' }

queryDecode('a=1&b=2&c=3')
// => { a: '1', b: '2', c: '3' }

queryDecode('') // => {}
queryDecode('key=') // => { key: '' }
```

## Response Objects

### Standard Response Format

```oak
{
    status: 200                          // HTTP status code
    headers: {                           // Optional headers
        'Content-Type': 'application/json'
        'Cache-Control': 'no-cache'
    }
    body: 'Response body'                // String body
}
```

### Pre-defined Responses

```oak
{ NotFound: NotFound, MethodNotAllowed: MethodNotAllowed } := import('http')

// 404 Not Found
NotFound
// => { status: 404, body: 'file not found' }

// 405 Method Not Allowed
MethodNotAllowed
// => { status: 405, body: 'method not allowed' }
```

## Complete Examples

### Basic Web Server

```oak
{ Server: Server } := import('http')

server := Server()

server.route('/', fn(params) fn(req, end) end({
    status: 200
    headers: { 'Content-Type': 'text/html; charset=utf-8' }
    body: '<h1>Welcome to Oak!</h1>'
}))

server.route('/hello/:name', fn(params) fn(req, end) end({
    status: 200
    body: 'Hello, ' + params.name + '!'
}))

println('Server starting on port 8080...')
server.start(8080)
```

### JSON API Server

```oak
{
    Server: Server
    queryDecode: queryDecode
} := import('http')
json := import('json')

server := Server()

server.route('/api/users/:id', fn(params) fn(req, end) {
    userId := params.id
    
    user := {
        id: userId
        name: 'User ' + userId
        email: 'user' + userId + '@example.com'
    }
    
    end({
        status: 200
        headers: { 'Content-Type': 'application/json' }
        body: json.serialize(user)
    })
})

server.route('/api/search', fn(params) fn(req, end) {
    // params already includes parsed query string
    query := params.q |> default('*')
    limit := int(params.limit) |> default(10)
    
    results := {
        query: query
        limit: limit
        results: []  // Would fetch from database
    }
    
    end({
        status: 200
        headers: { 'Content-Type': 'application/json' }
        body: json.serialize(results)
    })
})

server.start(3000)
```

### Static File Server with API

```oak
{
    Server: Server
    handleStatic: handleStatic
} := import('http')

server := Server()

// API routes
server.route('/api/status', fn(params) fn(req, end) end({
    status: 200
    headers: { 'Content-Type': 'application/json' }
    body: '{"status":"ok","uptime":' + string(time()) + '}'
}))

// Static file routes
server.route('/', handleStatic('./public/index.html'))
with server.route('/static/*path') fn(params) {
    handleStatic('./public/' + params.path)
}

// Catch-all 404
server.route('', fn(params) fn(req, end) end({
    status: 404
    headers: { 'Content-Type': 'text/html' }
    body: '<h1>404 - Page Not Found</h1>'
}))

server.start(8080)
```

## Request Object

The `req` parameter passed to handlers contains:

```oak
{
    method: 'GET'          // HTTP method (GET, POST, etc.)
    url: '/path?query=1'   // Full URL path with query string
    headers: { ... }       // Request headers
    body: '...'            // Request body (for POST, PUT, etc.)
}
```

## Notes

- The server listens on `0.0.0.0` (all interfaces) by default
- Automatic `X-Served-By: oak/libhttp` header is added to all responses
- Route patterns are matched in order—first match wins
- Query parameters are automatically decoded and merged into `params`
- Percent-encoding handles `+` as space for form data compatibility
- Static file handler only responds to GET requests
- MIME types are case-sensitive based on file extension
- All handler functions follow the pattern: `fn(params) fn(req, end) { ... }`

## Related Libraries

- `json` - For JSON serialization/parsing
- `fs` - For reading static files
- `fmt` - For formatted logging
- `str` - For string manipulation

## Parallel Batch Operations

### `pbatchQueryEncode(paramSets)`

Encodes multiple parameter objects into query strings in parallel.

```oak
http.pbatchQueryEncode([{a: 1}, {b: 2}])  // => ['a=1', 'b=2']
```

### `pbatchQueryDecode(queryStrings)`

Decodes multiple query strings into parameter objects in parallel.

```oak
http.pbatchQueryDecode(['a=1', 'b=2'])  // => [{a: '1'}, {b: '2'}]
```
