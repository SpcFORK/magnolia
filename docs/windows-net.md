# Windows Net Library (windows-net)

## Overview

`windows-net` contains Winsock and WinINet wrappers used by `windows`.

## Import

```oak
wnet := import('windows-net')
```

## Winsock exports

- `wsLastError`
- `wsaStartup`, `wsaCleanup`
- `socket`, `bindSocket`, `connectSocket`, `listenSocket`, `acceptSocket`
- `sendSocket`, `recvSocket`, `shutdownSocket`, `closeSocket`
- `htons`, `htonl`, `inetAddr`
- `sockaddrIn`

## WinINet exports

- `internetOpen`
- `internetConnect`
- `internetOpenUrl`
- `internetReadFile`
- `internetCloseHandle`
- `internetSimpleGet`

## Utility exports

- `callOk?`, `callValueOrZero`
