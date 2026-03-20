# Audio Complex Helpers (audio-complex)

## Overview

`audio-complex` factors out complex-number and unit-conversion math used by `libaudio`. It provides magnitude and phase helpers for complex values, MIDI frequency conversion, and linear-to-dB gain conversion.

This module is typically used by other audio modules (`audio-fft`, `audio`) rather than imported directly.

## Import

```oak
{ complexMag: complexMag, phaseAt: phaseAt, freqToMidi: freqToMidi, linearToDb: linearToDb } := import('audio-complex')
```

## Functions

### `complexMag(x, y)`

Returns the magnitude (modulus) of the complex number `x + yi`.

```oak
complexMag(3, 4)  // => 5.0
complexMag(1, 0)  // => 1.0
```

### `phaseAt(x, y)`

Returns the phase angle in radians for the complex number `x + yi`. Returns `0` for the zero vector.

```oak
phaseAt(1, 0)  // => 0.0
phaseAt(0, 1)  // => ~1.5708  (π/2)
phaseAt(0, 0)  // => 0
```

### `freqToMidi(freq)`

Converts a frequency in Hz to the nearest MIDI note number. A4 (440 Hz) maps to note 69.

```oak
freqToMidi(440)   // => 69  (A4)
freqToMidi(261.63) // => 60  (C4, middle C)
freqToMidi(880)   // => 81  (A5)
```

### `linearToDb(gain)`

Converts a linear gain factor to decibels (dB). Returns `?` for zero or negative gain values.

```oak
linearToDb(1.0)  // => 0.0   (unity gain)
linearToDb(2.0)  // => ~6.02 (double amplitude)
linearToDb(0.5)  // => ~-6.02
linearToDb(0)    // => ?
```
