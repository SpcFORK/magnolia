# Email Facade Library (email)

## Overview

`libemail` provides a single import point for all email protocol libraries:

- **email-smtp** — SMTP client/server
- **email-pop** — POP3 client/server
- **email-imap** — IMAP client/server

## Import

```oak
email := import('email')
```

This gives you:

```oak
email.smtp   // the email-smtp module
email.pop    // the email-pop module
email.imap   // the email-imap module
```

You can also import each protocol individually:

```oak
smtp := import('email-smtp')
pop := import('email-pop')
imap := import('email-imap')
```

The legacy import names `import('smtp')`, `import('pop')`, and `import('imap')` remain available for backward compatibility.

## Example

```oak
email := import('email')
{ wait: wait } := import('std')

// SMTP
closeSmtp := email.smtp.listen('127.0.0.1:2525', {
    onMessage: fn(msg) { ok?: true, message: 'queued' }
})
wait(0.05)
result := email.smtp.connect('127.0.0.1:2525')
client := result.client
client.ehlo('localhost')
client.send({
    from: 'sender@example.com'
    to: ['rcpt@example.com']
    raw: 'Subject: Test\r\n\r\nhello'
})
client.quit()
closeSmtp()

// POP
closePop := email.pop.listen('127.0.0.1:8110', {
    auth: fn(u, p) true
    messages: fn(_) [{ uid: 'm1', data: 'hello pop' }]
})
wait(0.05)
popResult := email.pop.connect('127.0.0.1:8110')
popResult.client.login('u', 'p')
popResult.client.retr(1)
popResult.client.quit()
closePop()

// IMAP
closeImap := email.imap.listen('127.0.0.1:8143', {
    auth: fn(u, p) true
    mailboxes: fn(_) [{ name: 'INBOX', messages: [{ body: 'hello imap' }] }]
})
wait(0.05)
imapResult := email.imap.connect('127.0.0.1:8143')
imapResult.client.login('u', 'p')
imapResult.client.select('INBOX')
imapResult.client.fetch(1)
imapResult.client.logout()
closeImap()
```

## See Also

- [email-smtp](email-smtp.md)
- [email-pop](email-pop.md)
- [email-imap](email-imap.md)
