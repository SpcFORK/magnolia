# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `audio-aiff`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI8(s)`

> returns `:int`

### `_i8ToF32(s)`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_byteStr(blist)`

### `_u16BE(n)`

> returns `:list`

### `_i16BE(n)`

### `_u32BE(n)`

> returns `:list`

### `_i32BE(n)`

### `_readU8(s, off)`

### `_readI8(s, off)`

### `_readU16BE(s, off)`

### `_readI16BE(s, off)`

### `_readU32BE(s, off)`

### `_readI32BE(s, off)`

### `_encodeExtended80(val)`

### `_decodeExtended80(s, off)`

### `_aiffSampleBytes(sample, bitDepth)`

### `_aiffDataChunk(samples, bitDepth)`

### `aiff(samples, sampleRate, channels, bitDepth)`

> returns `:string`

### `_aiffReadSample(data, off, bitDepth)`

### `_aiffReadSamples(data, dataOff, totalSamples, bytesPerSample, bitDepth)`

### `_aiffFindChunk(data, off, endOff, tag)`

> returns `?`

### `parseAiff(data)`

> returns `?`

### `pbatchAiff(specs)`

### `pbatchParseAiff(dataList)`

## Module: `audio-au`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI8(s)`

> returns `:int`

### `_i8ToF32(s)`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_byteStr(blist)`

### `_u32BE(n)`

> returns `:list`

### `_readU8(s, off)`

### `_readI8(s, off)`

### `_readU32BE(s, off)`

### `_readI16BE(s, off)`

### `_readI32BE(s, off)`

- `_AU_FMT_PCM8` · `2`
- `_AU_FMT_PCM16` · `3`
- `_AU_FMT_PCM32` · `5`
### `_auSampleBytes(sample, bitDepth)`

### `_auDataChunk(samples, bitDepth)`

### `au(samples, sampleRate, channels, bitDepth)`

### `_auReadSample(data, off, encoding)`

### `_auBytesPerSample(encoding)`

> returns `:int`

### `_auEncodingBitDepth(encoding)`

> returns `:int`

### `_auReadSamples(data, dataOff, totalSamples, bytesPerSample, encoding)`

### `parseAu(data)`

> returns `?`

### `pbatchAu(specs)`

### `pbatchParseAu(dataList)`

## Module: `audio-complex`

### `complexMag(x, y)`

### `phaseAt(x, y)`

> returns `:int`

### `freqToMidi(freq)`

### `linearToDb(g)`

> returns `:int`

## Module: `audio-dsp`

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

## Module: `audio-fft`

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

## Module: `audio-ogg`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_byteStr(blist)`

### `_u16LE(n)`

> returns `:list`

### `_u32LE(n)`

> returns `:list`

### `_i16LE(n)`

### `_readU8(s, off)`

### `_readU16LE(s, off)`

### `_readU32LE(s, off)`

### `_readI16LE(s, off)`

### `_pow2(n)`

### `_xor32(a, b)`

### `_crc32OggTable()`

- `_crcTable` · `_crc32OggTable(...)`
### `_crc32Ogg(data)`

### `_granuleBytes(n)`

### `_segmentTable(payloadLen)`

### `_oggPage(headerType, granule, serial, pageSeq, payload)`

### `_oggPcmIdHeader(sampleRate, channels, bitDepth)`

> returns `:string`

- `_maxPageSamples` · `32000`
### `_encodeSamples(samples, bitDepth)`

### `ogg(samples, sampleRate, channels, bitDepth)`

### `_readOggPageHeader(data, off)`

> returns `:object`

### `_readPcmSamples(data, off, nBytes)`

### `parseOgg(data)`

> returns `:object`

### `pbatchOgg(specs)`

### `pbatchParseOgg(dataList)`

## Module: `audio-raw`

### `_clipF32(s)`

> returns `:int`

### `_f32ToU8(s)`

### `_u8ToF32(s)`

### `_f32ToI8(s)`

> returns `:int`

### `_i8ToF32(s)`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_byteStr(blist)`

### `_i16LE(n)`

> returns `:list`

### `_i32LE(n)`

> returns `:list`

### `_i16BE(n)`

> returns `:list`

### `_i32BE(n)`

> returns `:list`

### `_readU8(s, off)`

### `_readI8(s, off)`

### `_readI16LE(s, off)`

### `_readI32LE(s, off)`

### `_readI16BE(s, off)`

### `_readI32BE(s, off)`

### `_rawSampleBytes(sample, bitDepth, signed, endian)`

### `rawEncode(samples, opts)`

### `_rawReadSample(data, off, bitDepth, signed, endian)`

### `rawDecode(data, opts)`

> returns `:object`

### `pbatchRawEncode(specs)`

### `pbatchRawDecode(specs)`

## Module: `audio-wav`

### `_clipF32(s)`

> returns `:int`

### `_f32ToI16(s)`

### `_i16ToF32(s)`

### `_f32ToI32(s)`

### `_i32ToF32(s)`

### `_f32ToU8(s)`

### `_u8ToF32(s)`

### `_byteStr(blist)`

### `_u16LE(n)`

> returns `:list`

### `_i16LE(n)`

### `_u32LE(n)`

> returns `:list`

### `_readU8(s, off)`

### `_readU16LE(s, off)`

### `_readI16LE(s, off)`

### `_readU32LE(s, off)`

### `_readI32LE(s, off)`

### `_wavSampleBytes(sample, bitDepth)`

### `_wavDataChunk(samples, bitDepth)`

### `_wavHeader(fileLen, sampleRate, channels, bitDepth, byteRate, blockAlign, dataLen)`

> returns `:string`

### `wav(samples, sampleRate, channels, bitDepth)`

### `_wavHeaderValid(data)`

> returns `:bool`

### `_wavFindDataChunk(data, off)`

> returns `?`

### `_wavReadSample(data, off, bitDepth)`

### `_wavReadSamples(data, dataOff, totalSamples, bytesPerSample, bitDepth)`

### `pbatchWav(specs)`

### `pbatchParseWav(dataList)`

### `parseWav(data)`

> returns `?`

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

### `ogg(samples, sampleRate, channels, bitDepth)`

### `parseOgg(data)`

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

## Module: `math`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

### `orient(x0, y0, x1, y1)`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

### `median(xs)`

### `stddev(xs)`

### `round(n, decimals)`

## Module: `math-base`

- `Pi` · `3.141592653589793`
- `E` · `2.718281828459045`
### `sign(n)`

> returns `:int`

### `abs(n)`

### `sqrt(n)`

## Module: `math-geo`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

> returns `:list`

### `orient(x0, y0, x1, y1)`

> returns `:int`

## Module: `math-stats`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

> returns `?`

### `median(xs)`

> returns `?`

### `stddev(xs)`

### `pbatchMean(datasets)`

### `pbatchStddev(datasets)`

### `round(n, decimals)`

## Module: `sort`

### `sort!(xs, pred)`

### `sort(xs, pred)`

### `_mergeSorted(a, b, pred)`

### `psort(xs, pred)`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `str`

### `checkRange(lo, hi)`

> **thunk** returns `:function`

### `upper?(c)`

> returns `:bool`

### `lower?(c)`

> returns `:bool`

### `digit?(c)`

> returns `:bool`

### `space?(c)`

> returns `:bool`

### `letter?(c)`

> returns `:bool`

### `word?(c)`

> returns `:bool`

### `join(strings, joiner)`

> returns `:string`

### `startsWith?(s, prefix)`

### `endsWith?(s, suffix)`

### `_matchesAt?(s, substr, idx)`

> returns `:bool`

### `indexOf(s, substr)`

### `rindexOf(s, substr)`

### `contains?(s, substr)`

### `cut(s, sep)`

> returns `:list`

### `lower(s)`

### `upper(s)`

### `_replaceNonEmpty(s, old, new)`

### `replace(s, old, new)`

### `_splitNonEmpty(s, sep)`

### `split(s, sep)`

### `_extend(pad, n)`

### `padStart(s, n, pad)`

### `padEnd(s, n, pad)`

### `_trimStartSpace(s)`

### `_trimStartNonEmpty(s, prefix)`

### `trimStart(s, prefix)`

### `_trimEndSpace(s)`

### `_trimEndNonEmpty(s, suffix)`

### `trimEnd(s, suffix)`

### `trim(s, part)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

