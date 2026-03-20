# Audio DSP Helpers (audio-dsp)

## Overview

`audio-dsp` provides windowing functions, FIR (finite impulse response) convolution and filter kernel generation, and ADSR envelope shaping. It is used by `libaudio` but can also be imported directly for signal-processing pipelines.

All sample and coefficient values are floating-point. Window functions return normalised coefficients (typically in [0, 1]).

## Import

```oak
dsp := import('audio-dsp')
// or destructure
{ windowHann: windowHann, convolve: convolve, adsr: adsr } := import('audio-dsp')
```

## Constants

### `TwoPi`

`2 * π`, commonly used in window and angle calculations.

```oak
TwoPi // => ~6.2832
```

## Windowing Functions

Window functions shape a block of samples to reduce spectral leakage before an FFT.

### `windowRect(n)`

Returns a rectangular (boxcar) window of length `n` — all coefficients are `1`.

```oak
windowRect(4) // => [1, 1, 1, 1]
```

### `windowHann(n)`

Returns a Hann window of length `n`.

```oak
windowHann(4) // => [0, 0.75, 0.75, 0]
```

### `windowHamming(n)`

Returns a Hamming window of length `n`.

```oak
windowHamming(4) // => [0.08, 0.77, 0.77, 0.08]
```

### `windowBlackman(n)`

Returns a Blackman window of length `n`.

```oak
windowBlackman(4) // => [0, 0.63, 0.63, 0]
```

### `applyWindow(samples, window)`

Multiplies `samples` element-wise by `window`. Both lists must have the same length.

```oak
applyWindow([1, 1, 1, 1], windowHann(4)) // => [0, 0.75, 0.75, 0]
```

## Convolution

### `convolve(signal, kernel)`

Computes the full linear convolution of `signal` and `kernel`. The output length is `len(signal) + len(kernel) - 1`.

```oak
convolve([1, 2, 3], [1, 1]) // => [1, 3, 5, 3]
```

## FIR Filter Kernels

### `lowPassFIR(cutoff, sampleRate, size?)`

Generates a Hamming-windowed sinc low-pass FIR kernel. `cutoff` and `sampleRate` are in Hz. `size` defaults to 63 (must be odd for symmetric kernels).

```oak
kernel := lowPassFIR(1000, 44100, 63)
filtered := convolve(samples, kernel)
```

### `highPassFIR(cutoff, sampleRate, size?)`

Generates a high-pass FIR kernel via spectral inversion of a low-pass kernel with the same parameters.

```oak
kernel := highPassFIR(2000, 44100)
```

## Envelopes

### `adsr(samples, sampleRate, attack?, decay?, sustain?, release?)`

Applies an ADSR volume envelope to a mono sample list. All time parameters are in seconds; `sustain` is a linear gain factor in [0, 1].

| Parameter  | Default |
|------------|---------|
| `attack`   | 0.01 s  |
| `decay`    | 0.1 s   |
| `sustain`  | 0.7     |
| `release`  | 0.2 s   |

```oak
shaped := adsr(samples, 44100, 0.005, 0.05, 0.8, 0.3)
```
