# gui-audio — Spatial Audio & Media Transport

`import('gui-audio')` provides Windows spatial audio with viewport-relative stereo panning and distance attenuation, system volume control, and media transport controls for music playback.

## Quick Start

```oak
audio := import('gui-audio')

// Create a spatial audio source
src := audio.spatialAudioSource(:laser, 200, 100, 0.8)

// Update with listener position to compute pan and gain
audio.spatialAudioUpdate(src, 400, 300, 800, 600, 500)

// System volume
vol := audio.getSystemVolume()
audio.setSystemVolume(50)

// Media transport controls
audio.enableMediaTransportControls({
    title: 'My Song'
    artist: 'Artist Name'
})
audio.setMediaPlaybackStatus('playing')
```

## Spatial Audio

### `spatialAudioSource(id, x, y, volume)`

Creates a spatial audio source descriptor with a position and base volume.

### `spatialAudioUpdate(source, listenerX, listenerY, viewW, viewH, maxDist)`

Updates a source with computed stereo pan `[-1..1]` and gain `[0..1]` based on distance from the listener.

### `spatialApplyToSamples(samples, pan, gain)`

Applies pan and gain to interleaved stereo samples.

### `spatialMixSources(sources, listenerX, listenerY, viewW, viewH, maxDist, bufLen)`

Mixes multiple spatial sources into a single stereo buffer.

## System Volume

### `setAppVolumeName(displayName)`

Sets the display name for this app in the Windows volume mixer.

### `getSystemVolume()`

Returns the current system master volume (0–100).

### `setSystemVolume(level)`

Sets the system master volume (0–100).

## Media Transport Controls

### `enableMediaTransportControls(options)`

Registers media transport control handlers with title and artist, integrating with the Windows system media overlay.

### `setMediaPlaybackStatus(status)`

Updates the playback status. Valid values: `'playing'`, `'paused'`, `'stopped'`, `'changing'`.

### `updateMediaInfo(title, artist, albumTitle)`

Updates the currently displayed media info on system media controls.
