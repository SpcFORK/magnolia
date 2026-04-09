# Email SMTP Library (email-smtp)

## Overview

`libemail-smtp` provides a practical SMTP client/server subset built on `socket`.

This is the `email-`prefixed version of `smtp`. The legacy `import('smtp')` still works.

This library currently works on Magnolia's native runtime. In the JavaScript runtime it will fail with structured socket `:error` results until stream socket support is added there.

Supported commands in this implementation:

- Client: `EHLO`, `HELO`, `STARTTLS`, `MAIL FROM`, `RCPT TO`, `DATA`, `QUIT`
- Server: `EHLO`, `HELO`, `NOOP`, `RSET`, `STARTTLS`, `MAIL FROM`, `RCPT TO`, `DATA`, `QUIT`

## Import

```oak
smtp := import('email-smtp')
```

Or via the facade:

```oak
email := import('email')
smtp := email.smtp
```

## Client API

### `connect(address, options?)`

Returns:

```oak
{ type: :ok, greeting: {...}, client: {...} }
```

Client methods:

- `command(line)`
- `ehlo(name?)`
- `helo(name?)`
- `startTLS(options?)`
- `mail(from)`
- `rcpt(to)`
- `data(body)`
- `send({from, to, raw? , headers?, body?})`
- `quit()`
- `close()`

## Server API

### `listen(address, handlers?, options?)`

Starts an SMTP server and returns a close function.

Handler fields:

- `hostname` server banner name
- `onMessage(message)` called after `DATA`

Message shape:

```oak
{
    from: 'sender@example.com'
    to: ['rcpt@example.com']
    body: 'raw message body including headers'
    tls?: false
}
```

`onMessage` should return an object like:

```oak
{ ok?: true, code: 250, message: 'queued' }
```

### STARTTLS

Enable STARTTLS with listener options:

```oak
{
    startTLS?: true
    certFile: './cert.pem'
    keyFile: './key.pem'
}
```

## Example

```oak
smtp := import('email-smtp')
{ wait: wait } := import('std')

closeServer := smtp.listen('127.0.0.1:2525', {
    onMessage: fn(message) {
        println(message.body)
        { ok?: true, message: 'queued' }
    }
})

wait(0.05)
clientResult := smtp.connect('127.0.0.1:2525')
client := clientResult.client
client.ehlo('localhost')
client.send({
    from: 'sender@example.com'
    to: ['rcpt@example.com']
    raw: 'Subject: Demo\r\n\r\nhello smtp'
})
client.quit()
closeServer()
```

## See Also

- [email](email.md) — unified facade
- [email-pop](email-pop.md)
- [email-imap](email-imap.md)
