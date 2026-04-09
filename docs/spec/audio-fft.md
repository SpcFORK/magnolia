# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-fft.oak`

- `TwoPi` — constant
### `_nextPow2(n)`

### `_log2Floor(n)`

### `_bitrev(x, bits)`

### `_fftSwap(re, im, i, j)`

### `_fftTwiddle(sign, j, step)`

> returns `:object`

### `_fftButterflyApply(re, im, u, v, wr, wi)`

### `_fftButterflyRange(re, im, sign, step, k, half, j)`

> returns `?`

### `_fftGroupPass(re, im, sign, step, n, k, half)`

> returns `?`

### `_fftCore(re, im, sign)`

### `fft(re, im)`

### `ifft(re, im)`

> returns `:object`

### `magnitude(re, im)`

### `phase(re, im)`

### `powerSpectrum(samples)`

### `pbatchFFT(signals)`

### `pbatchPowerSpectrum(signals)`

