# gui-leak-detect — Resource Leak Detector

`import('gui-leak-detect')` provides a memory and handle leak detector for GUI applications, tracking GDI objects, USER objects, handles, and working set memory with configurable alert thresholds.

## Quick Start

```oak
ld := import('gui-leak-detect')

// Create detector with baseline snapshot
state := ld.leakDetectorState()

// ... run application logic ...

// Take a snapshot and check for leaks
ld.leakSnapshot(state)
alerts := ld.leakCheck(state)
each(alerts, fn(a) println('LEAK: ' + a))

// Generate a human-readable report
report := ld.leakReport(state)
println(report)
```

## Resource Counters

### `getGDIObjectCount()`

Current count of GDI objects for this process.

### `getUserObjectCount()`

Current count of USER objects for this process.

### `getGDIObjectPeak()`

Peak GDI object count since process start.

### `getUserObjectPeak()`

Peak USER object count since process start.

### `getHandleCount()`

Number of open handles in this process.

### `getWorkingSetSize()`

Current working set (physical memory) in bytes.

### `getPeakWorkingSetSize()`

Peak working set in bytes.

## Leak Detector

### `leakDetectorState()`

Creates a new leak detector and takes a baseline snapshot of all resource counters.

### `leakSnapshot(state)`

Takes a snapshot of current resource usage.

### `leakSnapshotParallel(state)`

Same as `leakSnapshot` but collects all counters concurrently.

### `leakCheck(state)`

Compares the latest snapshot against the baseline. Returns a list of alert strings for any resource that exceeded its threshold.

### `leakReport(state)`

Generates a text report of resource usage and any alerts.

### `leakSetThresholds(state, thresholds)`

Updates alert thresholds. Fields: `gdiDelta`, `userDelta`, `handleDelta`, `memoryDeltaMB`.

### `leakResetBaseline(state)`

Resets the baseline to the latest snapshot.

### `leakTrend(state)`

Returns an array of deltas between consecutive snapshots for trend analysis.

## Default Thresholds

| Resource | Default Delta |
|----------|---------------|
| GDI objects | 50 |
| USER objects | 50 |
| Handles | 100 |
| Memory | 50 MB |

## Constants

| Constant | Value | Description |
|----------|-------|-------------|
| `GR_GDIOBJECTS` | 0 | GetGuiResources type |
| `GR_USEROBJECTS` | 1 | GetGuiResources type |
| `GR_GDIOBJECTS_PEAK` | 2 | GetGuiResources type |
| `GR_USEROBJECTS_PEAK` | 4 | GetGuiResources type |
