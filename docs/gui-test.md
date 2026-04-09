# gui-test — Automated GUI Test Harness

`import('gui-test')` provides automated GUI testing via synthetic input injection (mouse, keyboard, touch) using Win32 `SendInput`, enabling regression testing without manual interaction.

## Quick Start

```oak
gt := import('gui-test')

// Create a test suite
suite := gt.testSuite('Button Tests')

gt.testCase(suite, 'click save button', fn(result) {
    hwnd := gt.testFindWindow(?, 'My App')
    gt.testAssert(result, hwnd != ?, 'window found')

    rect := gt.testGetWindowRect(hwnd)
    gt.testMouseClick(rect.x + 50, rect.y + 100)

    title := gt.testGetWindowText(hwnd)
    gt.testAssertEqual(result, title, 'Saved', 'title updated')
})

gt.testRun(suite)
println(gt.testReport(suite))
```

## Synthetic Input

### Mouse

| Function | Description |
|----------|-------------|
| `testMouseMove(x, y)` | Move cursor to absolute screen position |
| `testMouseClick(x, y)` | Left click at (x, y) |
| `testMouseRightClick(x, y)` | Right click at (x, y) |
| `testMouseDoubleClick(x, y)` | Double left click at (x, y) |
| `testMouseDrag(x1, y1, x2, y2)` | Drag from (x1,y1) to (x2,y2) |
| `testMouseWheel(delta)` | Scroll wheel (delta > 0 scrolls up) |

### Keyboard

| Function | Description |
|----------|-------------|
| `testKeyDown(vk)` | Press a key down |
| `testKeyUp(vk)` | Release a key |
| `testKeyPress(vk)` | Press and release a key |
| `testTypeText(text)` | Type a string character by character (Unicode) |
| `testKeyCombo(keys)` | Press a key combination (e.g. `[VK_CONTROL, 67]` for Ctrl+C) |

## Test Framework

### `testSuite(name)`

Creates a named test suite for organizing test cases.

### `testCase(suite, name, fn)`

Adds a test case to the suite. The callback receives a result object for assertions.

### `testRun(suite)`

Runs all tests in the suite. Returns results.

### `testAssert(result, condition, message)`

Asserts that `condition` is true.

### `testAssertEqual(result, actual, expected, message)`

Asserts that `actual` equals `expected`.

### `testReport(suite)`

Generates a text report of test results.

## Window Inspection

### `testGetWindowRect(hwnd)`

Returns `{x, y, w, h}` for a window.

### `testGetWindowText(hwnd)`

Returns the window title text.

### `testIsWindowVisible?(hwnd)`

Returns `true` if a window is visible.

### `testGetForegroundWindow()`

Returns the currently focused window handle.

### `testFindWindow(className, windowName)`

Finds a window by class name or title.

### `testScreenshot(hwnd)`

Captures a bitmap of a window for visual regression testing.

### `testFreeScreenshot(ss)`

Frees a captured screenshot.

## Constants

### Input Types

| Constant | Value |
|----------|-------|
| `INPUT_MOUSE` | 0 |
| `INPUT_KEYBOARD` | 1 |

### Mouse Event Flags

| Constant | Value |
|----------|-------|
| `MOUSEEVENTF_MOVE` | 1 |
| `MOUSEEVENTF_LEFTDOWN` | 2 |
| `MOUSEEVENTF_LEFTUP` | 4 |
| `MOUSEEVENTF_RIGHTDOWN` | 8 |
| `MOUSEEVENTF_RIGHTUP` | 16 |
| `MOUSEEVENTF_WHEEL` | 2048 |
| `MOUSEEVENTF_ABSOLUTE` | 32768 |

### Keyboard Event Flags

| Constant | Value |
|----------|-------|
| `KEYEVENTF_KEYDOWN` | 0 |
| `KEYEVENTF_KEYUP` | 2 |
| `KEYEVENTF_UNICODE` | 4 |
