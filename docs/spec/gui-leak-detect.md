# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-leak-detect.oak`

- `guiThread` · `import(...)`
- `windows` · `import(...)`
- `GR_GDIOBJECTS` · `0`
- `GR_USEROBJECTS` · `1`
- `GR_GDIOBJECTS_PEAK` · `2`
- `GR_USEROBJECTS_PEAK` · `4`
### `_getCurrentProcess()`

### `getGDIObjectCount()`

### `getUserObjectCount()`

### `getGDIObjectPeak()`

### `getUserObjectPeak()`

### `getHandleCount()`

- `_PMC_SIZE` · `72`
### `getWorkingSetSize()`

### `getPeakWorkingSetSize()`

### `leakDetectorState()`

> returns `:object`

### `leakSnapshot(state)`

### `leakSnapshotParallel(state)`

### `leakCheck(state)`

### `leakReport(state)`

### `leakSetThresholds(state, thresholds)`

### `leakResetBaseline(state)`

### `leakTrend(state)`

