# Linux Constants Module (linux-constants)

## Overview

`linux-constants` exports constant values and candidate shared-library names used by
Magnolia's Linux interop modules.

Import this module when you need stable numeric flags or library candidate lists
without bringing in call wrappers.

## Import

```oak
constants := import('linux-constants')
```

## Exported Groups

### Library candidate lists

- `LibC`
- `LibDL`
- `LibX11`

### Memory protection (`mmap`/`mprotect`)

- `PROT_NONE`
- `PROT_READ`
- `PROT_WRITE`
- `PROT_EXEC`

### Mapping flags

- `MAP_SHARED`
- `MAP_PRIVATE`
- `MAP_FIXED`
- `MAP_ANONYMOUS`

### `open(2)` flags

- `O_RDONLY`
- `O_WRONLY`
- `O_RDWR`
- `O_CREAT`
- `O_TRUNC`
- `O_APPEND`

### `lseek(2)` constants

- `SEEK_SET`
- `SEEK_CUR`
- `SEEK_END`

### `access(2)` mode constants

- `F_OK`
- `R_OK`
- `W_OK`
- `X_OK`

### `dlopen(3)` flags

- `RTLD_LAZY`
- `RTLD_NOW`
- `RTLD_GLOBAL`
- `RTLD_LOCAL`

### `sysconf(3)` constants

- `_SC_PAGESIZE`

### X11 masks and event types

- `KeyPressMask`
- `ButtonPressMask`
- `ExposureMask`
- `StructureNotifyMask`
- `KeyPress`
- `Expose`
- `DestroyNotify`
- `ClientMessage`

## Example

```oak
k := import('linux-constants')

println('libc candidate count: ' + string(len(k.LibC)))
println('PROT_READ|PROT_WRITE = ' + string(k.PROT_READ | k.PROT_WRITE))
```
