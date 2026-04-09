# Email IMAP Library (email-imap)

## Overview

`libemail-imap` provides a compact IMAP client/server subset built on `socket`.

This is the `email-`prefixed version of `imap`. The legacy `import('imap')` still works.

This library currently works on Magnolia's native runtime. In the JavaScript runtime it will fail with structured socket `:error` results until stream socket support is added there.

Supported commands in this implementation:

- Client: `CAPABILITY`, `LOGIN`, `LIST`, `SELECT`, `FETCH`, `STARTTLS`, `NOOP`, `LOGOUT`
- Server: same subset

This implementation is intentionally lightweight and currently focuses on mailbox listing, selecting, and fetching message bodies or headers.

## Import

```oak
imap := import('email-imap')
```

Or via the facade:

```oak
email := import('email')
imap := email.imap
```

## Client API

### `connect(address, options?)`

Returns:

```oak
{ type: :ok, greeting: {...}, client: {...} }
```

Client methods:

- `command(text)` sends a tagged IMAP command
- `capability()`
- `login(user, pass)`
- `list(reference?, mailbox?)`
- `select(mailbox)`
- `fetch(seqSet, section?)`
- `startTLS(options?)`
- `noop()`
- `logout()`

`fetch()` returns a response object whose untagged entries may include `literal` payloads.

## Server API

### `listen(address, handlers?, options?)`

Handler fields:

- `auth(user, pass)` returns false or a truthy auth/session object
- `mailboxes(authState)` returns mailboxes

Mailbox shape:

```oak
{
    name: 'INBOX'
    flags: '\\HasNoChildren'
    messages: [{
        header: 'Subject: Demo\r\n\r\n'
        body: 'hello imap'
    }]
}
```

Listener `options.startTLS? = true` enables `STARTTLS` and requires `certFile` / `keyFile`.

## Example

```oak
imap := import('email-imap')
{ wait: wait } := import('std')

closeServer := imap.listen('127.0.0.1:8143', {
    auth: fn(user, pass) user = 'demo' & pass = 'secret'
    mailboxes: fn(_) [{
        name: 'INBOX'
        messages: [{
            header: 'Subject: Demo\r\n\r\n'
            body: 'hello imap'
        }]
    }]
})

wait(0.05)
result := imap.connect('127.0.0.1:8143')
client := result.client
client.login('demo', 'secret')
client.select('INBOX')
println(client.fetch(1, 'BODY[]').entries)
client.logout()
closeServer()
```

## See Also

- [email](email.md) — unified facade
- [email-smtp](email-smtp.md)
- [email-pop](email-pop.md)
