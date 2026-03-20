# Audio FFT Helpers (audio-fft)

## Overview

`audio-fft` provides a radix-2 Cooley-Tukey FFT/IFFT implementation and frequency-domain utility functions for Oak audio programs. It is used internally by `libaudio` and can also be imported directly for custom DSP pipelines.

All inputs and outputs use lists of floating-point numbers. Real-only signals may omit the imaginary component, which defaults to all-zeros.

## Import

```oak
fftLib := import('audio-fft')
// or destructure
{ fft: fft, ifft: ifft, magnitude: magnitude, phase: phase } := import('audio-fft')
```

## Constants

### `TwoPi`

`2 * π`, used internally for twiddle-factor calculations.

```oak
TwoPi // => ~6.2832
```

## FFT / IFFT

### `fft(re, im?)`

Computes the forward discrete Fourier transform of a signal. `re` is a list of real samples; `im` (imaginary part) is optional and defaults to a zero list of the same length.

The input length must be a power of 2. Returns an object `{ re: [...], im: [...] }`.

```oak
result := fft([1, 0, -1, 0])
result.re  // real components
result.im  // imaginary components
```

### `ifft(re, im)`

Computes the inverse DFT. Output is scaled by `1/N` so that `ifft(fft(x).re, fft(x).im).re ≈ x`.

```oak
signal := [1, 0, -1, 0]
spectrum := fft(signal)
roundtrip := ifft(spectrum.re, spectrum.im)
roundtrip.re  // => ~[1, 0, -1, 0]
```

## Frequency-Domain Utilities

### `magnitude(re, im)`

Returns a list of complex magnitudes, one per FFT bin.

```oak
mags := magnitude(spectrum.re, spectrum.im)
// mags.(i) = sqrt(re(i)^2 + im(i)^2)
```

### `phase(re, im)`

Returns a list of phase angles in radians, one per FFT bin.

```oak
phases := phase(spectrum.re, spectrum.im)
// phases.(i) = atan2(im(i), re(i))
```

## Notes

- Input length must be a power of 2; pad or truncate before calling.
- Only the first `N/2 + 1` bins are unique for real signals (Nyquist symmetry).
- For windowed analysis, apply a window (e.g. `windowHann` from `audio-dsp`) before calling `fft`.
