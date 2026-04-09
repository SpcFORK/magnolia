# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-events.oak`

- `windows` · `import(...)`
- `eventBusLib` · `import(...)`
- `guiInput` · `import(...)`
- `guiThread` · `import(...)`
### `emitThreadSafe(window, event, payload)`

### `emitFromWorker(window, event, payload)`

- `_mapResult` · `{2 entries}`
### `_mapMouseXY(window, physX, physY)`

### `_ensureEventBus(window)`

### `eventBus(window)`

### `on(window, event, handler)`

> returns `?`

### `once(window, event, handler)`

> returns `?`

### `off(window, event, tokenOrHandler)`

> returns `:int`

### `emit(window, event, payload, onDone)`

> returns `:int`

### `listenerCount(window, event)`

> returns `:int`

### `clearListeners(window, event)`

> returns `:bool`

### `publish(window, event, payload)`

> returns `:int`

### `onDispatch(window, handler)`

### `onceDispatch(window, handler)`

### `onRunStart(window, handler)`

### `onceRunStart(window, handler)`

### `onIdle(window, handler)`

### `onceIdle(window, handler)`

### `onFrame(window, handler)`

### `onceFrame(window, handler)`

### `onClosing(window, handler)`

### `onceClosing(window, handler)`

### `onClosed(window, handler)`

### `onceClosed(window, handler)`

- `MSG_OFF_TYPE` · `8`
- `MSG_OFF_WPARAM` · `16`
- `MSG_OFF_LPARAM` · `24`
- `FORM_WM_MOUSEMOVE` · `512`
- `FORM_WM_LBUTTONDOWN` · `513`
- `FORM_WM_LBUTTONUP` · `514`
- `FORM_WM_RBUTTONDOWN` · `516`
- `FORM_WM_RBUTTONUP` · `517`
- `FORM_WM_MBUTTONDOWN` · `519`
- `FORM_WM_MBUTTONUP` · `520`
- `FORM_WM_MOUSEWHEEL` · `522`
- `FORM_WM_LBUTTONDBLCLK` · `515`
- `FORM_WM_KEYDOWN` · `256`
- `FORM_WM_KEYUP` · `257`
- `FORM_WM_CHAR` · `258`
- `GUI_WM_SIZE` · `5`
- `GUI_WM_PAINT` · `15`
- `GUI_WM_ERASEBKGND` · `20`
- `GUI_WM_SIZING` · `532`
- `GUI_WM_WINDOWPOSCHANGED` · `71`
- `GUI_WM_ENTERSIZEMOVE` · `561`
- `GUI_WM_EXITSIZEMOVE` · `562`
- `GUI_WM_MOUSEHOVER` · `673`
- `GUI_WM_MOUSELEAVE` · `675`
- `GUI_WM_DPICHANGED` · `736`
- `WM_IME_STARTCOMPOSITION` · `269`
- `WM_IME_ENDCOMPOSITION` · `270`
- `WM_IME_COMPOSITION` · `271`
- `WM_IME_SETCONTEXT` · `641`
- `WM_IME_NOTIFY` · `642`
- `WM_IME_CHAR` · `646`
- `GCS_COMPSTR` · `8`
- `GCS_RESULTSTR` · `2048`
- `GCS_COMPATTR` · `16`
- `GCS_CURSORPOS` · `128`
- `TME_HOVER` · `1`
- `TME_LEAVE` · `2`
- `TME_SIZEOF` · `24`
- `FORM_VK_BACK` — constant
- `FORM_VK_TAB` — constant
- `FORM_VK_RETURN` — constant
- `FORM_VK_SHIFT` — constant
- `FORM_VK_CONTROL` — constant
- `FORM_VK_ALT` — constant
- `FORM_VK_ESCAPE` — constant
- `MK_LBUTTON` · `1`
- `MK_RBUTTON` · `2`
- `MK_SHIFT` · `4`
- `MK_CONTROL` · `8`
- `MK_MBUTTON` · `16`
### `modShift?(wp)`

> returns `:bool`

### `modCtrl?(wp)`

> returns `:bool`

### `modAlt?(window)`

> returns `:bool`

### `keyShiftDown?(window)`

> returns `:bool`

### `keyCtrlDown?(window)`

> returns `:bool`

### `keyAltDown?(window)`

### `formMsgType(window)`

### `formMsgWParam(window)`

### `formMsgLParam(window)`

### `formLoWord(v)`

> returns `:bool`

### `formHiWord(v)`

> returns `:bool`

### `_cacheDispatchContext(window)`

### `_clearDispatchContext(window)`

> returns `?`

### `formEventContext(window)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

> returns `:bool`

### `_vkKeyName(vk)`

- `_noKeyMatch` · `{1 entries}`
### `_extractKeyEvent(window, evt, expectedMsgType, expectedEvtType)`

> returns `:object`

### `onKeyDownEvent(window, handler)`

### `onceKeyDownEvent(window, handler)`

### `onKeyUpEvent(window, handler)`

### `onceKeyUpEvent(window, handler)`

### `onMouseMove(window, handler)`

### `onLButtonDown(window, handler)`

### `onLButtonUp(window, handler)`

### `onRButtonDown(window, handler)`

### `onRButtonUp(window, handler)`

### `onMButtonDown(window, handler)`

### `onMButtonUp(window, handler)`

### `onMouseWheel(window, handler)`

### `onLButtonDblClk(window, handler)`

### `_trackMouseEvent(window, flags)`

### `enableMouseTracking(window)`

### `onMouseHover(window, handler)`

### `onMouseLeave(window, handler)`

### `onKeyDown(window, handler)`

### `onKeyUp(window, handler)`

### `onChar(window, handler)`

### `onResize(window, handler)`

### `onDpiChanged(window, handler)`

### `_imeGetString(hwnd, gcsFlag)`

### `_imeGetCursorPos(hwnd)`

### `onImeStartComposition(window, handler)`

### `onImeEndComposition(window, handler)`

### `onImeComposition(window, handler)`

### `setImePosition(window, x, y)`

- `WM_TOUCH` · `576`
- `_TOUCHINPUT_SIZE` · `40`
- `TOUCHEVENTF_MOVE` · `1`
- `TOUCHEVENTF_DOWN` · `2`
- `TOUCHEVENTF_UP` · `4`
- `TOUCHEVENTF_INRANGE` · `8`
- `TOUCHEVENTF_PRIMARY` · `16`
- `TOUCHEVENTF_NOCOALESCE` · `32`
- `TOUCHEVENTF_PEN` · `64`
- `TOUCHEVENTF_PALM` · `128`
- `TOUCHINPUTMASKF_CONTACTAREA` · `4`
- `TOUCHINPUTMASKF_EXTRAINFO` · `2`
- `TOUCHINPUTMASKF_TIMEFROMSYSTEM` · `1`
- `TWF_FINETOUCH` · `1`
- `TWF_WANTPALM` · `2`
- `WM_POINTERDOWN` · `582`
- `WM_POINTERUP` · `583`
- `WM_POINTERUPDATE` · `581`
- `WM_POINTERENTER` · `585`
- `WM_POINTERLEAVE` · `586`
- `WM_POINTERCAPTURECHANGED` · `588`
- `WM_POINTERWHEEL` · `590`
- `WM_POINTERHWHEEL` · `591`
- `PT_POINTER` · `1`
- `PT_TOUCH` · `2`
- `PT_PEN` · `3`
- `PT_MOUSE` · `4`
- `PT_TOUCHPAD` · `5`
- `POINTER_FLAG_NONE` · `0`
- `POINTER_FLAG_NEW` · `1`
- `POINTER_FLAG_INRANGE` · `2`
- `POINTER_FLAG_INCONTACT` · `4`
- `POINTER_FLAG_FIRSTBUTTON` · `16`
- `POINTER_FLAG_SECONDBUTTON` · `32`
- `POINTER_FLAG_PRIMARY` · `8192`
- `POINTER_FLAG_DOWN` · `65536`
- `POINTER_FLAG_UPDATE` · `131072`
- `POINTER_FLAG_UP` · `262144`
- `PEN_FLAG_BARREL` · `1`
- `PEN_FLAG_INVERTED` · `2`
- `PEN_FLAG_ERASER` · `4`
- `PEN_MASK_PRESSURE` · `1`
- `PEN_MASK_ROTATION` · `2`
- `PEN_MASK_TILT_X` · `4`
- `PEN_MASK_TILT_Y` · `8`
### `registerTouchWindow(window, flags)`

### `unregisterTouchWindow(window)`

### `_parseTouchInputs(window, count)`

### `onTouch(window, handler)`

### `enableTouchInput(window, options)`

### `_pointerIdFromWParam(wp)`

> returns `:bool`

### `_getPointerType(pointerId)`

- `_POINTER_INFO_SIZE` · `96`
### `_getPointerInfo(pointerId)`

> returns `:object`

- `_POINTER_PEN_INFO_SIZE` · `120`
### `_getPointerPenInfo(pointerId)`

> returns `:object`

### `_makePointerEvent(window, pointerId)`

### `onPointerDown(window, handler)`

### `onPointerUp(window, handler)`

### `onPointerUpdate(window, handler)`

### `onPointerEnter(window, handler)`

### `onPointerLeave(window, handler)`

### `enablePointerInput(window)`

- `WM_INPUT` · `255`
- `HID_USAGE_PAGE_GENERIC` · `1`
- `HID_USAGE_PAGE_GAME` · `5`
- `HID_USAGE_PAGE_LED` · `8`
- `HID_USAGE_PAGE_BUTTON` · `9`
- `HID_USAGE_GENERIC_POINTER` · `1`
- `HID_USAGE_GENERIC_MOUSE` · `2`
- `HID_USAGE_GENERIC_JOYSTICK` · `4`
- `HID_USAGE_GENERIC_GAMEPAD` · `5`
- `HID_USAGE_GENERIC_KEYBOARD` · `6`
- `HID_USAGE_GENERIC_KEYPAD` · `7`
- `HID_USAGE_GENERIC_MULTI_AXIS` · `8`
- `RIDEV_REMOVE` · `1`
- `RIDEV_INPUTSINK` · `256`
- `RIDEV_NOLEGACY` · `48`
- `RIDEV_DEVNOTIFY` · `8192`
- `RID_INPUT` · `268435459`
- `RID_HEADER` · `268435461`
- `RIM_TYPEMOUSE` · `0`
- `RIM_TYPEKEYBOARD` · `1`
- `RIM_TYPEHID` · `2`
- `_RAWINPUTDEVICE_SIZE` — constant
### `registerRawInputDevice(window, usagePage, usage, flags)`

### `unregisterRawInputDevice(usagePage, usage)`

- `_RAWINPUTHEADER_SIZE` — constant
### `_getRawInputData(lParam)`

> returns `:object`

### `onRawInput(window, handler)`

### `enableRawMouse(window)`

### `enableRawKeyboard(window)`

### `enableRawGamepad(window)`

### `enableRawJoystick(window)`

### `isResizeDispatch?(window, step)`

> returns `:bool`

