# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-test.oak`

- `threadLib` · `import(...)`
- `windows` · `import(...)`
- `guiInput` · `import(...)`
- `INPUT_MOUSE` · `0`
- `INPUT_KEYBOARD` · `1`
- `INPUT_HARDWARE` · `2`
- `MOUSEEVENTF_MOVE` · `1`
- `MOUSEEVENTF_LEFTDOWN` · `2`
- `MOUSEEVENTF_LEFTUP` · `4`
- `MOUSEEVENTF_RIGHTDOWN` · `8`
- `MOUSEEVENTF_RIGHTUP` · `16`
- `MOUSEEVENTF_MIDDLEDOWN` · `32`
- `MOUSEEVENTF_MIDDLEUP` · `64`
- `MOUSEEVENTF_WHEEL` · `2048`
- `MOUSEEVENTF_ABSOLUTE` · `32768`
- `MOUSEEVENTF_VIRTUALDESK` · `16384`
- `KEYEVENTF_KEYDOWN` · `0`
- `KEYEVENTF_KEYUP` · `2`
- `KEYEVENTF_UNICODE` · `4`
- `KEYEVENTF_SCANCODE` · `8`
- `VK_RETURN` — constant
- `VK_ESCAPE` — constant
- `VK_TAB` — constant
- `VK_BACK` — constant
- `VK_SPACE` — constant
- `VK_LEFT` — constant
- `VK_UP` — constant
- `VK_RIGHT` — constant
- `VK_DOWN` — constant
- `VK_SHIFT` — constant
- `VK_CONTROL` — constant
- `VK_MENU` — constant
- `VK_DELETE` — constant
- `VK_HOME` — constant
- `VK_END` — constant
- `_INPUT_SIZE` · `40`
### `_buildMouseInput(dx, dy, flags, data)`

### `_buildKeyboardInput(vk, scan, flags)`

### `_sendInputs(inputs)`

### `testMouseMove(x, y)`

### `testMouseClick(x, y)`

### `testMouseRightClick(x, y)`

### `testMouseDoubleClick(x, y)`

### `testMouseDrag(x1, y1, x2, y2)`

### `testMouseWheel(delta)`

### `testKeyDown(vk)`

### `testKeyUp(vk)`

### `testKeyPress(vk)`

### `testTypeText(text)`

### `testKeyCombo(keys)`

### `testSuite(name)`

> returns `:object`

### `testCase(suite, name, testFn)`

### `testRun(suite)`

### `testAssert(result, condition, message)`

### `testAssertEqual(result, actual, expected, message)`

### `testReport(suite)`

### `testGetWindowRect(hwnd)`

> returns `:object`

### `testGetWindowText(hwnd)`

### `testIsWindowVisible?(hwnd)`

### `testGetForegroundWindow()`

### `testFindWindow(className, windowName)`

### `testScreenshot(hwnd)`

### `testFreeScreenshot(ss)`

