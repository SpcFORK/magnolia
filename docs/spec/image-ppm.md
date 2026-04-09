# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\image-ppm.oak`

### `clampByte(n)`

> returns `:int`

### `_mkHeader(magic, width, height, maxVal)`

### `_mkBuf(header)`

> returns `:object`

### `ppm(width, height, pixels)`

### `ppmPlain(width, height, pixels)`

### `pgm(width, height, pixels)`

### `pgmPlain(width, height, pixels)`

### `pbm(width, height, pixels)`

### `pbmPlain(width, height, pixels)`

### `_skipWS(data, idx)`

### `_readToken(data, idx)`

### `_readHeader(data)`

> returns `:object`

### `decodePPM(data)`

> returns `:object`

### `decodePGM(data)`

> returns `:object`

