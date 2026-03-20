# Windows Registry Library (windows-registry)

## Overview

`windows-registry` contains Win32 registry wrappers used by `windows`.

## Import

```oak
wreg := import('windows-registry')
```

## Low-level wrappers

- `regCloseKey`
- `regOpenKeyEx`
- `regCreateKeyEx`
- `regQueryValueEx`
- `regSetValueEx`
- `regDeleteValue`

## High-level helpers

- `regReadDword`
- `regWriteDword`
- `regReadString`
- `regWriteString`
