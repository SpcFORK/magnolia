# POP Library (pop)

## Overview

`libpop` provides a POP3 client/server subset.

This library currently works on Magnolia's native runtime. In the JavaScript runtime it will fail with structured socket `:error` results until stream socket support is added there.

Supported commands in this implementation:

- Client: `USER`, `PASS`, `CAPA`, `STAT`, `LIST`, `UIDL`, `RETR`, `TOP`, `DELE`, `RSET`, `NOOP`, `STLS`, `QUIT`
- Server: same subset

## Import

```oak
pop := import('pop')
```

## Client API

### `connect(address, options?)`

Returns:

```oak
{ type: :ok, banner: {...}, client: {...} }
```

Client methods:

- `user(name)`
- `pass(password)`
- `login(name, password)`
- `capa()`
- `stat()`
- `list(index?)`
- `uidl(index?)`
- `retr(index)`
- `top(index, lines)`
- `dele(index)`
- `rset()`
- `noop()`
- `startTLS(options?)`
- `quit()`

Multi-line responses expose `lines` terminated by protocol dot-handling already removed.

## Server API

### `listen(address, handlers?, options?)`

Handler fields:

- `auth(user, pass)` returns false for rejection or any truthy auth/session object for success
- `messages(authState)` returns a list of mailbox messages

Message input shape:

```oak
{ uid: 'm1', data: 'Subject: Demo\r\n\r\nhello pop' }
```

Listener `options.startTLS? = true` enables `STLS` and requires `certFile` / `keyFile`.

## Example

```oak
pop := import('pop')
{ wait: wait } := import('std')

closeServer := pop.listen('127.0.0.1:8110', {
    auth: fn(user, pass) user = 'demo' & pass = 'secret'
    messages: fn(_) [{ uid: 'm1', data: 'Subject: Demo\r\n\r\nhello pop' }]
})

wait(0.05)
result := pop.connect('127.0.0.1:8110')
client := result.client
client.login('demo', 'secret')
println(client.retr(1).lines)
client.quit()
closeServer()
```