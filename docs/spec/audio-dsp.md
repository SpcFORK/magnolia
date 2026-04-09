# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio-dsp.oak`

- `TwoPi` — constant
### `_windowPhase(i, n)`

### `windowRect(n)`

### `windowHann(n)`

### `windowHamming(n)`

### `windowBlackman(n)`

### `applyWindow(samples, window)`

### `_convolveBounds(n, sLen, kLen)`

> returns `:object`

### `_convolveAt(signal, kernel, n, sLen, kLen)`

### `convolve(signal, kernel)`

### `pconvolve(signal, kernel, numWorkers)`

### `_sincNorm(x)`

> returns `:int`

### `lowPassFIR(cutoff, sampleRate, size)`

### `highPassFIR(cutoff, sampleRate, size)`

### `_adsrEnv(i, aS, dS, sustain, rStart, rS)`

### `adsr(samples, sampleRate, attack, decay, sustain, release)`

