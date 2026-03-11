# Audio Library (audio)

## Overview

`libaudio` provides raw audio processing utilities for Oak programs. It covers PCM sample-type conversions, WAV file format encoding and parsing, oscillator synthesis, basic digital signal processing (DSP) transforms, windowing functions, FIR filter kernel generation, ADSR envelope shaping, a radix-2 Cooley-Tukey FFT/IFFT implementation, and common unit-conversion helpers.

Samples are represented as normalised floating-point values in the range [-1, 1] unless otherwise documented. For stereo audio, frames may be represented either as interleaved lists `[L0, R0, L1, R1, ...]` or as split channel objects `{ left: [...], right: [...] }`.

## Import

```oak
audio := import('audio')
// or destructure specific functions
{ sine: sine, wav: wav, fft: fft } := import('audio')
```

## Constants

### `TwoPi`

The constant 2π (two times pi), commonly needed for oscillator calculations.

```oak
TwoPi // => 6.28318530717958647692528676655900576839433879875021
```

### Sample Rates

Common standard audio sample rates in Hz.

```oak
audio.SampleRateCD  // => 44100 (CD quality)
audio.SampleRateDVD // => 48000 (DVD / broadcast)
audio.SampleRateHD  // => 96000 (high-definition)
```

### Channel Counts

```oak
audio.Mono   // => 1
audio.Stereo // => 2
```

### Bit Depths

Common PCM bit-width constants.

```oak
audio.BitDepth8  // => 8
audio.BitDepth16 // => 16
audio.BitDepth32 // => 32
```

## Sample-Type Conversions

Audio is typically stored in integer formats but processed as normalised floats. These functions convert between the two.

### `f32ToI16(sample)`

Converts a normalised float in [-1, 1] to a signed 16-bit integer.

```oak
f32ToI16(1.0)    // => 32767
f32ToI16(-1.0)   // => -32767
f32ToI16(0.0)    // => 0
f32ToI16(0.5)    // => 16383
```

### `i16ToF32(sample)`

Converts a signed 16-bit integer to a normalised float in [-1, 1].

```oak
i16ToF32(32767)  // => ~0.9999
i16ToF32(-32768) // => -1.0
i16ToF32(0)      // => 0.0
```

### `f32ToI32(sample)` / `i32ToF32(sample)`

Converts between normalised floats and signed 32-bit integers.

```oak
f32ToI32(1.0)    // => 2147483647
i32ToF32(2147483647) // => ~1.0
```

### `f32ToU8(sample)` / `u8ToF32(sample)`

Converts between normalised floats and unsigned 8-bit integers (used in 8-bit WAV files, where 128 = silence).

```oak
f32ToU8(0.0)     // => 128 (silence)
f32ToU8(1.0)     // => 255
f32ToU8(-1.0)    // => 0

u8ToF32(128)     // => 0.0
u8ToF32(255)     // => ~1.0
```

### `clipF32(sample)`

Hard-clips a single sample to the [-1, 1] range.

```oak
clipF32(0.5)     // => 0.5
clipF32(1.5)     // => 1.0
clipF32(-2.0)    // => -1.0
```

## WAV File Format

### `wav(samples, sampleRate, channels, bitDepth)`

Encodes a list of normalised float samples into a binary WAV file. For stereo audio, `samples` must already be interleaved as `[L0, R0, L1, R1, ...]`.

- **sampleRate** (default: 44100) — samples per second
- **channels** (default: 1) — mono=1, stereo=2
- **bitDepth** (default: 16) — 8, 16, or 32-bit PCM

Returns a binary string suitable for writing to disk.

```oak
// Generate 1 second of 440 Hz sine wave at CD quality
samples := sine(440, 44100, 44100)
wavFile := wav(samples, 44100, 1, 16)

// Write to file (with appropriate IO)
// file := open('tone.wav', :write)
// write(file.fd, 0, wavFile)
```

### `parseWav(data)`

Parses a binary WAV file and returns an object with properties:
- **sampleRate** — samples per second
- **channels** — 1 for mono, 2 for stereo, etc.
- **bitDepth** — 8, 16, or 32
- **samples** — list of normalised float samples

Returns `?` if the file is not a valid PCM WAV.

```oak
// Read WAV file and parse it
fileBytes := readFile('audio.wav')
parsed := parseWav(fileBytes)

samples := parsed.samples     // normalised float list
rate := parsed.sampleRate    // e.g., 44100
```

## Oscillators

Oscillator functions generate periodic waveforms of a specified frequency and length. All oscillators accept an optional `phase` parameter (in radians) to control the starting phase.

### `sine(freq, sampleRate, length, phase)`

Generates a sine wave.

```oak
// 1 second of 440 Hz at 44100 Hz
s := sine(440, 44100, 44100)

// With phase offset
s := sine(440, 44100, 44100, math.Pi / 2)
```

### `square(freq, sampleRate, length, phase)`

Generates a band-unlimited square wave (±1). Useful for subtractive synthesis.

```oak
s := square(440, 44100, 44100)
```

### `sawtooth(freq, sampleRate, length, phase)`

Generates a sawtooth wave that ramps from -1 to 1 each period. Rich harmonic content.

```oak
s := sawtooth(220, 44100, 44100)
```

### `triangle(freq, sampleRate, length, phase)`

Generates a triangle wave, smoother than sawtooth but still rich.

```oak
s := triangle(440, 44100, 44100)
```

### `noise(length)`

Generates white noise with amplitude uniformly distributed in [-1, 1].

```oak
// 0.5 seconds of white noise at 44100 Hz
s := noise(22050)
```

### `silence(length)`

Generates zero-amplitude silence of the given length.

```oak
// 100 ms of silence
s := silence(4410)  // at 44100 Hz
```

## Basic DSP

### `gain(samples, factor)`

Scales all samples by a constant multiplicative factor.

```oak
s := sine(440, 44100, 44100)
quiet := gain(s, 0.5)      // 50% volume
loud := gain(s, 2.0)       // 200% volume
```

### `clip(samples)`

Hard-clips all samples to the [-1, 1] range, useful for preventing digital distortion.

```oak
s := gain(sine(440, 44100, 44100), 10)  // very loud
safe := clip(s)  // clamped to [-1, 1]
```

### `normalize(samples)`

Scales the entire signal so that the peak absolute value is 1.0. Returns the list unchanged if all samples are zero.

```oak
s := [0.1, 0.2, -0.15]
normalized := normalize(s)  // peak will be ±1.0
```

### `add(a, b)`

Adds two equal-length sample lists element-wise. Useful for combining signals (mixing before normalization to avoid clipping).

```oak
s1 := sine(440, 44100, 44100)
s2 := sine(880, 44100, 44100)
mixed := add(s1, s2) |> gain(_, 0.5)  // combine and reduce to avoid clipping
```

### `mix(a, b, ratio)`

Crossfades between two equal-length sample lists. `ratio` = 0 returns all of `a`; `ratio` = 1 returns all of `b`; `ratio` = 0.5 is 50/50.

```oak
fade := mix(s1, s2, 0.5)  // blend two sounds equally
blend := mix(s1, s2, 0.3) // more of s1
```

### `dc(samples, offset)`

Adds a constant DC (direct current) offset to all samples. Use to bias or shift a signal.

```oak
s := sine(440, 44100, 44100)
biased := dc(s, 0.2)  // shifts signal up by 0.2
```

### `reverseSignal(samples)`

Reverses the time order of all samples (plays them backwards).

```oak
s := sine(440, 44100, 44100)
rev := reverseSignal(s)
```

### `panStereo(samples, position)`

Applies equal-power stereo panning to a mono sample list. Uses cosine/sine gains to maintain constant perceived loudness.

- **position** in [-1, 1]: -1 is hard-left, 0 is centre, +1 is hard-right

Returns `{ left: [...], right: [...] }` with each channel as a sample list.

```oak
mono := sine(440, 44100, 44100)
stereo := panStereo(mono, 0.3)  // slightly right

stereo.left      // left channel
stereo.right     // right channel
```

## Stereo Utilities

### `interleave(left, right)`

Combines two separate channel lists into a single interleaved list `[L0, R0, L1, R1, ...]`.

```oak
{ left: L, right: R } := panStereo(mono, 0.5)
interleaved := interleave(L, R)
```

### `deinterleave(samples, nChannels)`

Splits an interleaved multi-channel list into a list of per-channel sample lists.

```oak
channels := deinterleave(interleaved, 2)
left := channels.(0)
right := channels.(1)
```

## Windowing Functions

Window functions are used with spectral analysis and FIR filter design to reduce spectral leakage.

### `windowRect(n)`

Rectangular (boxcar) window — all ones. No windowing.

```oak
w := windowRect(256)  // [1, 1, 1, ..., 1]
```

### `windowHann(n)`

Hann (Hanning) window — smooth, gradual taper at ends.

```oak
w := windowHann(256)
```

### `windowHamming(n)`

Hamming window — similar to Hann but never quite reaches zero at edges.

```oak
w := windowHamming(256)
```

### `windowBlackman(n)`

Blackman window — wider main lobe, excellent sidelobe rejection.

```oak
w := windowBlackman(256)
```

### `applyWindow(samples, window)`

Multiplies a sample list element-wise by a window function. Both lists must have equal length.

```oak
frame := sine(440, 44100, 1024)
windowed := applyWindow(frame, windowHann(1024))
```

## FIR Filters

### `convolve(signal, kernel)`

Computes the full linear convolution of a signal with a filter kernel, returning a list of length `len(signal) + len(kernel) - 1`. To get the same length as the input, slice out the centre portion.

```oak
filtered := convolve(signal, kernel)
same_len := filtered |> slice(len(kernel) / 2 - 1, len(signal))
```

### `lowPassFIR(cutoff, sampleRate, size)`

Generates a Hamming-windowed sinc low-pass FIR filter kernel.

- **cutoff** (Hz) — the -6 dB cut frequency
- **sampleRate** — in Hz
- **size** (default: 63) — kernel length, must be odd

```oak
kernel := lowPassFIR(2000, 44100, 63)
filtered := convolve(signal, kernel) |> slice(31, len(signal) + 31)
```

### `highPassFIR(cutoff, sampleRate, size)`

Generates a high-pass FIR kernel via spectral inversion of a low-pass kernel.

```oak
kernel := highPassFIR(2000, 44100, 63)
filtered := convolve(signal, kernel)
```

## ADSR Envelope

### `adsr(samples, sampleRate, attack, decay, sustain, release)`

Applies an ADSR (Attack, Decay, Sustain, Release) volume envelope to a sample list.

- **attack** (seconds, default: 0.01) — fade in time
- **decay** (seconds, default: 0.1) — fade to sustain level
- **sustain** (gain [0, 1], default: 0.7) — held level between key events
- **release** (seconds, default: 0.2) — fade out time

```oak
note := sine(440, 44100, 44100)
shaped := adsr(note, 44100, 0.05, 0.1, 0.6, 0.2)
```

## Frequency Domain (FFT)

### `fft(re, im)` / `ifft(re, im)`

Compute the forward and inverse Discrete Fourier Transform using a radix-2 Cooley-Tukey algorithm. Input lists must have a power-of-2 length; use `powerSpectrum()` for automatic padding.

- `fft` returns `{ re, im }` complex frequency bins
- `ifft` scales the result by 1/N automatically

```oak
// Forward transform
samples := sine(440, 44100, 1024)
result := fft(samples, ?)

// Inverse transform (note: imaginary part ~0 for real signals)
back := ifft(result.re, result.im)
reconstructed := back.re
```

### `magnitude(re, im)` / `phase(re, im)`

Compute the magnitude and phase of each FFT bin.

```oak
result := fft(samples, ?)
mags := magnitude(result.re, result.im)
phases := phase(result.re, result.im)
```

### `powerSpectrum(samples)`

Computes the power-spectrum magnitude from a real signal. Automatically zero-pads to the next power of 2.

```oak
s := sine(440, 44100, 100)
ps := powerSpectrum(s)
```

## Unit Conversions

### Time ↔ Samples

### `durationSamples(seconds, sampleRate)`

Converts a duration in seconds to a sample count.

```oak
n := durationSamples(2.5, 44100)  // => 110250
```

### `samplesDuration(n, sampleRate)`

Converts a sample count to a duration in seconds.

```oak
t := samplesDuration(44100, 44100)  // => 1.0
```

### Frequency ↔ FFT Bins

### `freqBin(f, sampleRate, fftSize)`

Returns the FFT bin index for a given frequency.

```oak
bin := freqBin(1000, 44100, 1024)  // => 23
```

### `binFreq(bin, sampleRate, fftSize)`

Returns the centre frequency (Hz) of a given FFT bin.

```oak
f := binFreq(23, 44100, 1024)  // => ~991 Hz
```

### MIDI ↔ Frequency

### `midiToFreq(note)`

Converts a MIDI note number to a frequency in Hz. MIDI note 69 = A4 = 440 Hz.

```oak
midiToFreq(69)   // => 440.0 (A4)
midiToFreq(60)   // => 261.63 (middle C)
midiToFreq(72)   // => 523.25 (C5)
```

### `freqToMidi(freq)`

Converts a frequency in Hz to the nearest MIDI note number.

```oak
freqToMidi(440)     // => 69
freqToMidi(261.63)  // => 60
```

### Gain (dB ↔ Linear)

### `dbToLinear(db)`

Converts a dB value to a linear amplitude gain factor.

```oak
dbToLinear(-6)   // => ~0.501 (half amplitude)
dbToLinear(0)    // => 1.0 (no change)
dbToLinear(6)    // => ~1.995 (double amplitude)
```

### `linearToDb(g)`

Converts a linear gain factor to dB. Returns `?` for zero or negative values.

```oak
linearToDb(0.5)   // => ~-6.02
linearToDb(1.0)   // => 0.0
linearToDb(2.0)   // => ~6.02
```

## Example: Synthesizing a Simple Tone

```oak
audio := import('audio')

// Generate a 1-second 440 Hz tone with ADSR
samples := audio.sine(440, 44100, 44100)
shaped := audio.adsr(samples, 44100, 0.05, 0.1, 0.7, 0.2)
normalized := audio.normalize(shaped)

// Encode to WAV
wavData := audio.wav(normalized, 44100, 1, 16)

// (save wavData to file...)
```
