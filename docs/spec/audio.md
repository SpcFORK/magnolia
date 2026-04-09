# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\audio.oak`

- `TwoPi` — constant
- `SampleRateCD` · `44100`
- `SampleRateDVD` · `48000`
- `SampleRateHD` · `96000`
- `Mono` · `1`
- `Stereo` · `2`
- `BitDepth8` · `8`
- `BitDepth16` · `16`
- `BitDepth32` · `32`
### `f32ToI16(s)`

### `i16ToF32(s)`

### `f32ToI32(s)`

### `i32ToF32(s)`

### `f32ToU8(s)`

### `u8ToF32(s)`

### `clipF32(s)`

> returns `:int`

### `wav(samples, sampleRate, channels, bitDepth)`

### `parseWav(data)`

### `aiff(samples, sampleRate, channels, bitDepth)`

### `parseAiff(data)`

### `au(samples, sampleRate, channels, bitDepth)`

### `parseAu(data)`

### `rawEncode(samples, opts)`

### `rawDecode(data, opts)`

### `sine(freq, sampleRate, length, phase)`

### `_oscPhaseTurn(freq, sampleRate, i, phase)`

### `square(freq, sampleRate, length, phase)`

### `sawtooth(freq, sampleRate, length, phase)`

### `triangle(freq, sampleRate, length, phase)`

### `noise(length)`

### `silence(length)`

### `gain(samples, factor)`

### `clip(samples)`

### `normalize(samples)`

### `add(a, b)`

### `mix(a, b, ratio)`

### `dc(samples, offset)`

### `reverseSignal(samples)`

### `panStereo(samples, position)`

> returns `:object`

### `interleave(left, right)`

### `deinterleave(samples, nChannels)`

### `windowRect(n)`

### `windowHann(n)`

### `windowHamming(n)`

### `windowBlackman(n)`

### `applyWindow(samples, window)`

### `convolve(signal, kernel)`

### `lowPassFIR(cutoff, sampleRate, size)`

### `highPassFIR(cutoff, sampleRate, size)`

### `adsr(samples, sampleRate, attack, decay, sustain, release)`

### `fft(re, im)`

### `ifft(re, im)`

### `magnitude(re, im)`

### `phase(re, im)`

### `powerSpectrum(samples)`

### `durationSamples(seconds, sampleRate)`

### `samplesDuration(n, sampleRate)`

### `freqBin(f, sampleRate, fftSize)`

### `binFreq(bin, sampleRate, fftSize)`

### `midiToFreq(note)`

> returns `:int`

### `freqToMidi(freq)`

### `dbToLinear(db)`

### `linearToDb(g)`

### `pmixTracks(tracks)`

### `psynthesize(sampleRate, length, specs)`

