# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\GUI.oak`

- `_guiImportT0` · `nanotime(...)`
- `std` · `import(...)`
- `windows` · `import(...)`
- `linux` · `import(...)`
- `_guiImportT1` · `nanotime(...)`
- `guiMesh` · `import(...)`
- `guiRender` · `import(...)`
- `gui3dmath` · `import(...)`
- `gui2d` · `import(...)`
- `guiRaster` · `import(...)`
- `guiLighting` · `import(...)`
- `_guiImportT2` · `nanotime(...)`
- `guiWeb` · `import(...)`
- `guiNativeWin` · `import(...)`
- `guiNativeLinux` · `import(...)`
- `_guiImportT3` · `nanotime(...)`
- `guiDraw` · `import(...)`
- `guiColor` · `import(...)`
- `guiEvents` · `import(...)`
- `guiInput` · `import(...)`
- `guiGraph` · `import(...)`
- `guiForm` · `import(...)`
- `guiLoop` · `import(...)`
- `guiShader` · `import(...)`
- `guiFonts` · `import(...)`
- `guiVideo` · `import(...)`
- `guiResolution` · `import(...)`
- `guiCanvas` · `import(...)`
- `guiAcc` · `import(...)`
- `guiClipboard` · `import(...)`
- `guiFiledrop` · `import(...)`
- `guiAudio` · `import(...)`
- `guiGamepad` · `import(...)`
- `guiAA` · `import(...)`
- `guiDrawOps` · `import(...)`
- `guiGpuInfo` · `import(...)`
- `guiLeakDetect` · `import(...)`
- `guiTest` · `import(...)`
- `guiDialogs` · `import(...)`
- `guiMenus` · `import(...)`
- `guiPrint` · `import(...)`
- `guiTheme` · `import(...)`
- `guiSystray` · `import(...)`
- `guiThread` · `import(...)`
- `_guiImportT4` · `nanotime(...)`
- `_importTimings` · `{5 entries}`
### `getImportTimings()`

### `backend()`

> returns `:atom`

### `isWindows?()`

### `isLinux?()`

### `isWeb?()`

### `_ensureEventBus(window)`

### `eventBus(window)`

### `on(window, event, handler)`

### `once(window, event, handler)`

### `off(window, event, tokenOrHandler)`

### `emit(window, event, payload, onDone)`

### `listenerCount(window, event)`

### `clearListeners(window, event)`

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

### `onKeyDownEvent(window, handler)`

### `onceKeyDownEvent(window, handler)`

### `onKeyUpEvent(window, handler)`

### `onceKeyUpEvent(window, handler)`

### `_publish(window, event, payload)`

### `_clampByte(value)`

### `_clampOpacity(value)`

### `rgb(r, g, b)`

### `rgba(r, g, b, a, background)`

### `colorR(color)`

### `colorG(color)`

### `colorB(color)`

### `opacity(color, amount, background)`

- `GL_COLOR_BUFFER_BIT` · `16384`
- `GL_DEPTH_BUFFER_BIT` · `256`
- `GL_TRIANGLES` · `4`
### `createWindow(title, width, height, options)`

### `createWindowAsync(title, width, height, options)`

### `awaitWindow(future)`

### `show(window)`

### `hide(window)`

### `move(window, x, y)`

### `resize(window, width, height)`

### `scale(window, scaleX, scaleY)`

### `fullscreen(window, enabled)`

### `lockResize(window, locked)`

### `setAlwaysOnTop(window, enabled)`

### `setWindowOpacity(window, alpha)`

### `setWindowColorKey(window, colorKey)`

### `removeLayeredStyle(window)`

### `registerWindow(window)`

### `unregisterWindow(window)`

### `allWindows()`

### `pollAllWindows()`

### `closeAllWindows()`

### `anyWindowOpen?()`

### `saveWindowState(window)`

### `restoreWindowState(window, state)`

### `acquireSingleInstance(name)`

### `releaseSingleInstance(inst)`

### `setTaskbarProgress(window, completed, total)`

### `setTaskbarProgressState(window, flags)`

- `TBPF_NOPROGRESS` — constant
- `TBPF_INDETERMINATE` — constant
- `TBPF_NORMAL` — constant
- `TBPF_ERROR` — constant
- `TBPF_PAUSED` — constant
### `createOwnedWindow(parent, title, width, height, options)`

### `showModalDialog(parent, title, width, height, options, setupFn)`

### `closeModalDialog(dialog)`

### `setWindowOwner(child, parent)`

### `getWindowMonitor(window)`

### `centerOnMonitor(window)`

### `moveToMonitor(window, hMonitor)`

### `getWindowDpi(window)`

### `getMonitorInfo(hMonitor)`

### `monitorFromPoint(x, y, flags)`

### `monitorFromRect(left, top, right, bottom, flags)`

### `extendFrameIntoClientArea(window, margins)`

### `enableGlassSheet(window)`

### `onNcHitTest(window, handler)`

### `customChromeHitTest(mx, my, width, height, borderSize, captionHeight)`

### `setWindowRgn(window, hRgn, redraw)`

### `createRoundRectRgn(left, top, right, bottom, rx, ry)`

- `HTCLIENT` — constant
- `HTCAPTION` — constant
- `HTMINBUTTON` — constant
- `HTMAXBUTTON` — constant
- `HTCLOSE` — constant
### `showToastNotification(title, message, options)`

### `showToastWithFallback(title, message, options)`

### `enableMica(window)`

### `enableAcrylic(window)`

### `enableTabbedMica(window)`

### `disableBackdrop(window)`

### `setDwmDarkMode(window, dark)`

### `setDwmAttribute(window, attribute, value)`

### `beginDrag(window)`

### `updateDrag(window)`

### `endDrag(window)`

### `setDesignResolution(window, logicalWidth, logicalHeight, options)`

### `clearDesignResolution(window)`

### `hasDesignResolution?(window)`

### `designWidth(window)`

### `designHeight(window)`

### `physicalWidth(window)`

### `physicalHeight(window)`

### `resolutionScaleX(window)`

### `resolutionScaleY(window)`

### `resolutionOffsetX(window)`

### `resolutionOffsetY(window)`

### `physicalToLogical(window, px, py)`

### `logicalToPhysical(window, lx, ly)`

### `createCanvas(window, id, options)`

### `initWebGL(window, contextName, attrs)`

### `webglCreateShader(window, shaderType, source)`

### `webglCreateProgram(window, vertexShader, fragmentShader)`

### `webglUseProgram(window, program)`

### `webglClearColor(window, r, g, b, a)`

### `webglViewport(window, x, y, width, height)`

### `webglClear(window, mask)`

### `webglDrawArrays(window, mode, first, count)`

### `webglFlush(window)`

### `setTitle(window, title)`

### `setIcon(window, iconSpec)`

### `beginFrame(window)`

### `endFrame(window)`

### `drawText(window, x, y, text, color)`

### `textWidth(text, window)`

### `setFont(window, fontSpec)`

### `clearFont(window)`

### `fillRect(window, x, y, width, height, color, borderColor)`

### `pushMask(window, x, y, w, h)`

### `popMask(window)`

### `degToRad(deg)`

### `Vec3(x, y, z)`

### `_transformPoint(v, transform)`

### `_projectPoint(window, p, camera)`

### `_transformVertices(vertices, transform, i, out)`

### `drawLine(window, x1, y1, x2, y2, color)`

### `drawLinesBatch(window, segs, color)`

- `_graphCtxCached` · `{5 entries}`
### `graphRange(values, options)`

### `graphMapX(index, count, x, width)`

### `graphMapY(value, y, height, range)`

### `drawGraphAxes(window, x, y, width, height, options)`

### `drawLineGraph(window, x, y, width, height, values, options)`

### `drawBarGraph(window, x, y, width, height, values, options)`

### `drawSparkline(window, x, y, width, height, values, options)`

- `MSG_OFF_TYPE` — constant
- `MSG_OFF_WPARAM` — constant
- `MSG_OFF_LPARAM` — constant
- `FORM_WM_MOUSEMOVE` — constant
- `FORM_WM_LBUTTONDOWN` — constant
- `FORM_WM_LBUTTONUP` — constant
- `FORM_WM_RBUTTONDOWN` — constant
- `FORM_WM_RBUTTONUP` — constant
- `FORM_WM_MBUTTONDOWN` — constant
- `FORM_WM_MBUTTONUP` — constant
- `FORM_WM_MOUSEWHEEL` — constant
- `FORM_WM_LBUTTONDBLCLK` — constant
- `FORM_WM_KEYDOWN` — constant
- `FORM_WM_KEYUP` — constant
- `FORM_WM_CHAR` — constant
- `GUI_WM_SIZE` — constant
- `GUI_WM_PAINT` — constant
- `GUI_WM_ERASEBKGND` — constant
- `GUI_WM_SIZING` — constant
- `GUI_WM_WINDOWPOSCHANGED` — constant
- `GUI_WM_ENTERSIZEMOVE` — constant
- `GUI_WM_EXITSIZEMOVE` — constant
- `FORM_VK_BACK` — constant
- `FORM_VK_TAB` — constant
- `FORM_VK_RETURN` — constant
- `FORM_VK_SHIFT` — constant
- `FORM_VK_CONTROL` — constant
- `FORM_VK_ALT` — constant
- `FORM_VK_ESCAPE` — constant
- `MK_LBUTTON` — constant
- `VK_BACK` — constant
- `VK_TAB` — constant
- `VK_CLEAR` — constant
- `VK_RETURN` — constant
- `VK_SHIFT` — constant
- `VK_CONTROL` — constant
- `VK_ALT` — constant
- `VK_PAUSE` — constant
- `VK_CAPSLOCK` — constant
- `VK_ESCAPE` — constant
- `VK_SPACE` — constant
- `VK_PAGEUP` — constant
- `VK_PAGEDOWN` — constant
- `VK_END` — constant
- `VK_HOME` — constant
- `VK_LEFT` — constant
- `VK_UP` — constant
- `VK_RIGHT` — constant
- `VK_DOWN` — constant
- `VK_INSERT` — constant
- `VK_DELETE` — constant
- `VK_0` — constant
- `VK_1` — constant
- `VK_2` — constant
- `VK_3` — constant
- `VK_4` — constant
- `VK_5` — constant
- `VK_6` — constant
- `VK_7` — constant
- `VK_8` — constant
- `VK_9` — constant
- `VK_A` — constant
- `VK_B` — constant
- `VK_C` — constant
- `VK_D` — constant
- `VK_E` — constant
- `VK_F` — constant
- `VK_G` — constant
- `VK_H` — constant
- `VK_I` — constant
- `VK_J` — constant
- `VK_K` — constant
- `VK_L` — constant
- `VK_M` — constant
- `VK_N` — constant
- `VK_O` — constant
- `VK_P` — constant
- `VK_Q` — constant
- `VK_R` — constant
- `VK_S` — constant
- `VK_T` — constant
- `VK_U` — constant
- `VK_V` — constant
- `VK_W` — constant
- `VK_X` — constant
- `VK_Y` — constant
- `VK_Z` — constant
- `VK_NUMPAD0` — constant
- `VK_NUMPAD1` — constant
- `VK_NUMPAD2` — constant
- `VK_NUMPAD3` — constant
- `VK_NUMPAD4` — constant
- `VK_NUMPAD5` — constant
- `VK_NUMPAD6` — constant
- `VK_NUMPAD7` — constant
- `VK_NUMPAD8` — constant
- `VK_NUMPAD9` — constant
- `VK_MULTIPLY` — constant
- `VK_ADD` — constant
- `VK_SEPARATOR` — constant
- `VK_SUBTRACT` — constant
- `VK_DECIMAL` — constant
- `VK_DIVIDE` — constant
- `VK_F1` — constant
- `VK_F2` — constant
- `VK_F3` — constant
- `VK_F4` — constant
- `VK_F5` — constant
- `VK_F6` — constant
- `VK_F7` — constant
- `VK_F8` — constant
- `VK_F9` — constant
- `VK_F10` — constant
- `VK_F11` — constant
- `VK_F12` — constant
- `VK_NUMLOCK` — constant
- `VK_SCROLLLOCK` — constant
- `VK_OEM_SEMICOLON` — constant
- `VK_OEM_PLUS` — constant
- `VK_OEM_COMMA` — constant
- `VK_OEM_MINUS` — constant
- `VK_OEM_PERIOD` — constant
- `VK_OEM_SLASH` — constant
- `VK_OEM_TILDE` — constant
- `VK_OEM_LBRACKET` — constant
- `VK_OEM_BACKSLASH` — constant
- `VK_OEM_RBRACKET` — constant
- `VK_OEM_QUOTE` — constant
- `VK_LSHIFT` — constant
- `VK_RSHIFT` — constant
- `VK_LCONTROL` — constant
- `VK_RCONTROL` — constant
- `VK_LALT` — constant
- `VK_RALT` — constant
### `isLetterKey?(vk)`

### `isDigitKey?(vk)`

### `isNumpadKey?(vk)`

### `isFunctionKey?(vk)`

### `isArrowKey?(vk)`

### `isModifierKey?(vk)`

- `MK_RBUTTON` — constant
- `MK_SHIFT` — constant
- `MK_CONTROL` — constant
- `MK_MBUTTON` — constant
### `modShift?(wp)`

### `modCtrl?(wp)`

### `modAlt?(window)`

### `keyShiftDown?(window)`

### `keyCtrlDown?(window)`

### `keyAltDown?(window)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

### `formMsgType(window)`

### `formMsgWParam(window)`

### `formMsgLParam(window)`

### `formLoWord(v)`

### `formHiWord(v)`

### `formEventContext(window)`

### `onMouseMove(window, handler)`

### `onLButtonDown(window, handler)`

### `onLButtonUp(window, handler)`

### `onRButtonDown(window, handler)`

### `onRButtonUp(window, handler)`

### `onMButtonDown(window, handler)`

### `onMButtonUp(window, handler)`

### `onMouseWheel(window, handler)`

### `onLButtonDblClk(window, handler)`

### `enableMouseTracking(window)`

### `onMouseHover(window, handler)`

### `onMouseLeave(window, handler)`

### `onKeyDown(window, handler)`

### `onKeyUp(window, handler)`

### `onChar(window, handler)`

### `onResize(window, handler)`

### `onDpiChanged(window, handler)`

### `onImeStartComposition(window, handler)`

### `onImeEndComposition(window, handler)`

### `onImeComposition(window, handler)`

### `setImePosition(window, x, y)`

### `registerTouchWindow(window, flags)`

### `unregisterTouchWindow(window)`

### `onTouch(window, handler)`

### `enableTouchInput(window, options)`

- `TOUCHEVENTF_MOVE` — constant
- `TOUCHEVENTF_DOWN` — constant
- `TOUCHEVENTF_UP` — constant
- `TOUCHEVENTF_PRIMARY` — constant
- `TOUCHEVENTF_PEN` — constant
- `TOUCHEVENTF_PALM` — constant
- `TWF_FINETOUCH` — constant
- `TWF_WANTPALM` — constant
### `onPointerDown(window, handler)`

### `onPointerUp(window, handler)`

### `onPointerUpdate(window, handler)`

### `onPointerEnter(window, handler)`

### `onPointerLeave(window, handler)`

### `enablePointerInput(window)`

- `PT_POINTER` — constant
- `PT_TOUCH` — constant
- `PT_PEN` — constant
- `PT_MOUSE` — constant
- `PT_TOUCHPAD` — constant
- `POINTER_FLAG_INCONTACT` — constant
- `POINTER_FLAG_PRIMARY` — constant
- `POINTER_FLAG_DOWN` — constant
- `POINTER_FLAG_UPDATE` — constant
- `POINTER_FLAG_UP` — constant
- `PEN_FLAG_BARREL` — constant
- `PEN_FLAG_ERASER` — constant
- `PEN_FLAG_INVERTED` — constant
### `getGamepadState(playerIndex)`

### `setGamepadVibration(playerIndex, left, right)`

### `stopGamepadVibration(playerIndex)`

### `applyThumbDeadzone(state)`

### `applyDeadzone(value, deadzone)`

### `gamepadPollState()`

### `gamepadPoll(ps)`

### `gamepadButtonDown?(state, button)`

### `gamepadButtonPressed?(ps, idx, button)`

### `gamepadButtonReleased?(ps, idx, button)`

### `gamepadConnected?(ps, idx)`

### `gamepadJustConnected?(ps, idx)`

### `gamepadJustDisconnected?(ps, idx)`

### `gamepadDpadUp?(state)`

### `gamepadDpadDown?(state)`

### `gamepadDpadLeft?(state)`

### `gamepadDpadRight?(state)`

### `gamepadStart?(state)`

### `gamepadBack?(state)`

### `gamepadA?(state)`

### `gamepadB?(state)`

### `gamepadX?(state)`

### `gamepadY?(state)`

### `gamepadLeftShoulder?(state)`

### `gamepadRightShoulder?(state)`

- `XUSER_MAX_COUNT` — constant
- `XINPUT_GAMEPAD_DPAD_UP` — constant
- `XINPUT_GAMEPAD_DPAD_DOWN` — constant
- `XINPUT_GAMEPAD_DPAD_LEFT` — constant
- `XINPUT_GAMEPAD_DPAD_RIGHT` — constant
- `XINPUT_GAMEPAD_START` — constant
- `XINPUT_GAMEPAD_BACK` — constant
- `XINPUT_GAMEPAD_LEFT_THUMB` — constant
- `XINPUT_GAMEPAD_RIGHT_THUMB` — constant
- `XINPUT_GAMEPAD_LEFT_SHOULDER` — constant
- `XINPUT_GAMEPAD_RIGHT_SHOULDER` — constant
- `XINPUT_GAMEPAD_A` — constant
- `XINPUT_GAMEPAD_B` — constant
- `XINPUT_GAMEPAD_X` — constant
- `XINPUT_GAMEPAD_Y` — constant
### `registerRawInputDevice(window, usagePage, usage, flags)`

### `unregisterRawInputDevice(usagePage, usage)`

### `onRawInput(window, handler)`

### `enableRawMouse(window)`

### `enableRawKeyboard(window)`

### `enableRawGamepad(window)`

### `enableRawJoystick(window)`

- `HID_USAGE_PAGE_GENERIC` — constant
- `HID_USAGE_GENERIC_POINTER` — constant
- `HID_USAGE_GENERIC_MOUSE` — constant
- `HID_USAGE_GENERIC_JOYSTICK` — constant
- `HID_USAGE_GENERIC_GAMEPAD` — constant
- `HID_USAGE_GENERIC_KEYBOARD` — constant
- `RIDEV_INPUTSINK` — constant
- `RIDEV_NOLEGACY` — constant
- `RIDEV_DEVNOTIFY` — constant
- `RIM_TYPEMOUSE` — constant
- `RIM_TYPEKEYBOARD` — constant
- `RIM_TYPEHID` — constant
### `setProcessDpiAwarenessContext(context)`

### `getDpiForWindow(hwnd)`

### `getDpiForSystem()`

### `getDpiForMonitor(hMonitor, dpiType)`

### `monitorFromWindow(hwnd, flags)`

### `adjustWindowRectExForDpi(x, y, w, h, style, exStyle, dpi)`

### `enableNonClientDpiScaling(hwnd)`

### `dpiScale(value, dpi)`

### `dpiUnscale(value, dpi)`

- `DPI_AWARENESS_CONTEXT_UNAWARE` — constant
- `DPI_AWARENESS_CONTEXT_SYSTEM_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2` — constant
### `clipboardGetText()`

### `clipboardSetText(text)`

### `clipboardHasText()`

### `enableFileDrop(window)`

### `disableFileDrop(window)`

### `onFileDrop(window, handler)`

### `enableOleDrop(window)`

### `disableOleDrop(window)`

### `onOleDrop(window, handler)`

### `dragDropState()`

### `onDragOver(window, state, handler)`

- `DROPEFFECT_NONE` — constant
- `DROPEFFECT_COPY` — constant
- `DROPEFFECT_MOVE` — constant
- `DROPEFFECT_LINK` — constant
### `openFileDialog(options)`

### `saveFileDialog(options)`

### `chooseColor(options)`

### `chooseFont(options)`

### `pickFolder(options)`

### `showPrintDialog(options)`

### `startDoc(hDC, docName, outputFile)`

### `startPage(hDC)`

### `endPage(hDC)`

### `endDoc(hDC)`

### `abortDoc(hDC)`

### `deletePrintDC(hDC)`

### `printTextOut(hDC, x, y, text)`

### `printMoveTo(hDC, x, y)`

### `printLineTo(hDC, x, y)`

### `printRectangle(hDC, l, t, r, b)`

### `printEllipse(hDC, l, t, r, b)`

### `printSetFont(hDC, height, weight, italic, fontName)`

### `printSetTextColor(hDC, r, g, b)`

### `printSetBkMode(hDC, mode)`

### `printSetPen(hDC, style, width, r, g, b)`

### `printDeleteObject(hObj)`

### `getDeviceCaps(hDC, index)`

### `getPrinterPageSize(hDC)`

### `createPreviewDC(width, height)`

### `destroyPreviewDC(preview)`

### `printToFile(outputPath, docName, renderFn)`

### `printDocument(options, renderPageFn)`

- `PD_ALLPAGES` — constant
- `PD_SELECTION` — constant
- `PD_PAGENUMS` — constant
- `PD_COLLATE` — constant
- `PD_PRINTTOFILE` — constant
- `PD_CURRENTPAGE` — constant
- `DEVCAP_HORZRES` — constant
- `DEVCAP_VERTRES` — constant
- `DEVCAP_LOGPIXELSX` — constant
- `DEVCAP_LOGPIXELSY` — constant
### `createMenu()`

### `createPopupMenu()`

### `menuAppendItem(hMenu, id, label)`

### `menuAppendSeparator(hMenu)`

### `menuAppendSubmenu(hMenu, hSub, label)`

### `setMenuBar(window, hMenu)`

### `removeMenuBar(window)`

### `destroyMenu(hMenu)`

### `showPopupMenu(window, hMenu, x, y)`

### `onMenuCommand(window, handler)`

### `buildMenu(spec)`

### `createAcceleratorTable(entries)`

### `destroyAcceleratorTable(hAccel)`

### `installAccelerators(window, entries)`

### `getSystemMenu(window, reset)`

### `appendSystemMenuItem(window, id, label)`

### `appendSystemMenuSeparator(window)`

### `resetSystemMenu(window)`

### `onSysCommand(window, handler)`

### `isDarkMode?()`

### `isHighContrast?()`

### `accentColor()`

### `addTrayIcon(window, id, tooltip, hIcon)`

### `removeTrayIcon(window, id)`

### `updateTrayTooltip(window, id, tooltip)`

### `showBalloonTip(window, id, title, message, iconFlag)`

### `onTrayEvent(window, handler)`

### `minimizeToTray(window)`

### `restoreFromTray(window)`

### `formSetStatus(state, message, ok)`

### `formPopChar(s)`

### `formClamp(v, minVal, maxVal)`

### `formHitListIndex(mx, my, x, y, itemW, itemH, count)`

### `formHitRectKey(mx, my, rects)`

### `formToggleIfHit(value, mx, my, x, y, w, h)`

### `formSelectByHit(current, mx, my, rects)`

### `formSetByAssignments(state, assignments)`

### `formResetFlags(state, keys, value)`

### `formSetHoverFromRects(state, mx, my, rects)`

### `formToggleKeysByHit(state, mx, my, rects)`

### `formSetKeyByHit(state, targetKey, mx, my, rects)`

### `formTruncateText(s, maxChars)`

### `formIsPrintableChar?(code)`

### `formSelectionState()`

### `formSelSetCursor(sel, pos, shifting)`

### `formSelMoveCursor(sel, text, dir, shifting)`

### `formSelHome(sel, shifting)`

### `formSelEnd(sel, text, shifting)`

### `formSelAll(sel, text)`

### `formSelRange(sel)`

### `formSelHasSelection?(sel)`

### `formSelSelectedText(sel, text)`

### `formSelDeleteSelection(sel, text)`

### `formSelInsertAtCursor(sel, text, insert)`

### `formSelBackspace(sel, text)`

### `formSelClickPos(mx, fieldX, fieldPadding, text)`

### `formDrawFieldWithSel(window, x, y, w, h, text, placeholder, focused, sel, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor, selColor)`

### `formUndoState(maxHistory)`

### `formUndoPush(hist, text, sel)`

### `formUndo(hist, text, sel)`

### `formRedo(hist, text, sel)`

### `formAppendByFocus(state, focus, c, fieldSpecs, notesSpec)`

### `formBackspaceByFocus(state, focus, fieldKeys, notesSpec)`

### `formCopyByFocus(state, focus, fieldKeys)`

### `formPasteByFocus(state, focus, fieldSpecs)`

### `formCutByFocus(state, focus, fieldKeys)`

### `formNotesAppendChar(lines, c)`

### `formNotesBackspace(lines)`

### `formNotesNewLine(lines, maxLines)`

### `formSliderValue(mx, sliderX, sliderW, handleW, minVal, maxVal)`

### `formApplySliderDrag(state, dragKey, mx, sliderX, sliderW, handleW, minVal, maxVal, bindings)`

### `formNextField(current, order)`

### `formPrevField(current, order)`

### `formFocusState(tabOrder)`

### `formFocusNext(fs)`

### `formFocusPrev(fs)`

### `formFocusSet(fs, key)`

### `formFocusIs?(fs, key)`

### `formHandleTabKey(fs, shiftDown)`

### `formIsActivateKey?(vk)`

### `formHandleKeyNav(fs, vk, shiftDown, widgetHandlers, options)`

### `formCheckboxKeyToggle(checked, vk)`

### `formRadioKeySelect(current, vk, choices)`

### `formSliderKeyAdjust(value, vk, minVal, maxVal, step)`

### `formDropdownKeyNav(isOpen, selectedIdx, vk, itemCount)`

### `formTreeKeyNav(flatList, selectedIdx, vk)`

### `formTableKeyNav(selectedRow, vk, rowCount, pageSize)`

### `formMaskText(s)`

- `_formCtxCached` · `{4 entries}`
### `formDrawBorder(window, x, y, w, h, color)`

### `formDrawFocusRing(window, x, y, w, h, color, options)`

### `formDrawField(window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawPasswordField(window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawCheckbox(window, x, y, checked, label, fieldColor, borderColor, checkColor)`

### `formDrawRadio(window, x, y, selected, label, fieldColor, borderColor, accentColor)`

### `formDrawPrimaryButton(window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, bottomLineColor)`

### `formDrawSecondaryButton(window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, borderColor)`

### `formDrawSlider(window, x, y, w, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, options)`

### `formDrawLabeledPercentSlider(window, x, labelY, sliderY, w, label, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, percentRectColor)`

### `formDrawNotes(window, x, y, w, h, lines, focused, placeholder, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawStatusBanner(window, x, y, w, h, message, ok, fieldColor, borderColor, successColor, errorColor)`

### `formDrawProgressBar(window, x, y, w, value, maxVal, trackColor, fillColor, borderColor, options)`

### `formDrawTabStrip(window, x, y, tabs, activeIdx, bgColor, activeBgColor, borderColor, activeTextColor, inactiveTextColor, options)`

### `formHitTab(mx, my, x, y, tabs, options)`

### `formDrawDropdown(window, x, y, w, h, selectedLabel, open, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDropdownList(window, x, y, w, items, hoverIdx, itemHeight, listBgColor, hoverBgColor, borderColor)`

### `formHitDropdownItem(mx, my, x, y, w, items, itemHeight)`

### `formTooltipState()`

### `formTooltipUpdate(state, mx, my, regionKey, text, delay)`

### `formTooltipHide(state)`

### `formDrawTooltip(window, state, bgColor, borderColor, textColor, options)`

### `formDrawSpinner(window, x, y, w, h, value, focused, fieldColor, fieldFocusColor, borderColor, arrowColor, options)`

### `formSpinnerHit(mx, my, x, y, w, h, options)`

### `formSpinnerAdjust(value, direction, options)`

### `formDrawScrollbar(window, x, y, w, h, scrollPos, contentSize, viewSize, trackColor, thumbColor, borderColor, options)`

### `formScrollbarHit(mx, my, x, y, w, h, scrollPos, contentSize, viewSize, options)`

### `formDrawTreeView(window, x, y, w, h, nodes, selectedIdx, scrollPos, bgColor, selectedBgColor, textColor, selectedTextColor, borderColor, options)`

### `formTreeHitRow(mx, my, x, y, w, flat, scrollPos, options)`

### `formTreeToggle(flat, idx)`

### `formTreeContentHeight(flat, options)`

### `formDatePickerState(year, month, day)`

### `formDatePickerPrevMonth(state)`

### `formDatePickerNextMonth(state)`

### `formDateLabel(state)`

### `formDrawDateField(window, x, y, w, h, state, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDateCalendar(window, x, y, state, bgColor, headerBgColor, selectedBgColor, todayBorderColor, textColor, selectedTextColor, headerTextColor, borderColor, options)`

### `formDateCalendarHit(mx, my, x, y, state, options)`

### `formDrawTable(window, x, y, w, h, columns, rows, scrollPos, selectedRow, headerBgColor, rowBgColor, altRowBgColor, selectedBgColor, headerTextColor, textColor, selectedTextColor, borderColor, options)`

### `formTableHitRow(mx, my, x, y, w, columns, rows, scrollPos, options)`

### `formTableContentHeight(rows, options)`

### `formTableHitColumn(mx, x, columns)`

### `formFrameTimerState(maxSamples)`

### `formFrameTimerTick(state)`

### `formDrawFrameTimingOverlay(window, state, x, y, w, h, options)`

### `formRichTextState(lines)`

### `formRichTextSetLines(state, lines)`

### `formRichTextAppendLine(state, spans)`

### `formRichTextInsertSpan(state, lineIdx, spanIdx, span)`

### `formDrawRichText(window, x, y, w, h, state, bgColor, borderColor)`

### `formRichTextScroll(state, delta)`

### `formRichTextTotalHeight(state)`

### `richSpan(text)`

### `richBold(text)`

### `richItalic(text)`

### `richUnderline(text)`

### `richColored(text, r, g, b)`

### `richStyled(text, options)`

### `enableAccessibility(window)`

### `accTree(window)`

### `accNode(id, name, role, bounds, state, children)`

### `accRegister(window, id, name, role, bounds, state)`

### `accUnregister(window, id)`

### `accSetName(window, id, name)`

### `accSetState(window, id, state)`

### `accSetValue(window, id, value)`

### `accSetBounds(window, id, bounds)`

### `accSetDescription(window, id, desc)`

### `accSetDefaultAction(window, id, action)`

### `accFocus(window, id)`

### `accSelection(window, id)`

### `notifyAccEvent(window, event, childId)`

### `accRegisterButton(window, id, label, x, y, w, h)`

### `accRegisterCheckbox(window, id, label, x, y, w, h, checked?)`

### `accRegisterRadio(window, id, label, x, y, w, h, selected?)`

### `accRegisterTextField(window, id, label, x, y, w, h, value)`

### `accRegisterSlider(window, id, label, x, y, w, h, value)`

### `accRegisterProgressBar(window, id, label, x, y, w, h, value)`

### `accRegisterTab(window, id, label, x, y, w, h, selected?)`

### `accRegisterListItem(window, id, label, x, y, w, h, selected?)`

### `accRegisterLink(window, id, label, x, y, w, h)`

### `accRegisterStaticText(window, id, label, x, y, w, h)`

### `accRegisterGroup(window, id, label, x, y, w, h)`

### `accRegisterTreeItem(window, id, label, x, y, w, h, expanded?)`

### `accRegisterTable(window, id, label, x, y, w, h)`

### `accRegisterTableCell(window, id, label, x, y, w, h)`

### `accRegisterDropdown(window, id, label, x, y, w, h, expanded?)`

### `accAnnounce(window, id, text)`

### `accDump(window)`

### `accVerify(window)`

### `accRoleName(role)`

### `accStateName(state)`

- `ROLE_SYSTEM_PUSHBUTTON` — constant
- `ROLE_SYSTEM_CHECKBUTTON` — constant
- `ROLE_SYSTEM_RADIOBUTTON` — constant
- `ROLE_SYSTEM_TEXT` — constant
- `ROLE_SYSTEM_STATICTEXT` — constant
- `ROLE_SYSTEM_SLIDER` — constant
- `ROLE_SYSTEM_PROGRESSBAR` — constant
- `ROLE_SYSTEM_PAGETAB` — constant
- `ROLE_SYSTEM_LISTITEM` — constant
- `ROLE_SYSTEM_LINK` — constant
- `ROLE_SYSTEM_OUTLINEITEM` — constant
- `ROLE_SYSTEM_TABLE` — constant
- `ROLE_SYSTEM_CELL` — constant
- `ROLE_SYSTEM_COMBOBOX` — constant
- `ROLE_SYSTEM_APPLICATION` — constant
- `ROLE_SYSTEM_CLIENT` — constant
- `ROLE_SYSTEM_GROUPING` — constant
- `STATE_SYSTEM_NORMAL` — constant
- `STATE_SYSTEM_FOCUSED` — constant
- `STATE_SYSTEM_SELECTED` — constant
- `STATE_SYSTEM_CHECKED` — constant
- `STATE_SYSTEM_PRESSED` — constant
- `STATE_SYSTEM_EXPANDED` — constant
- `STATE_SYSTEM_COLLAPSED` — constant
- `STATE_SYSTEM_READONLY` — constant
- `STATE_SYSTEM_FOCUSABLE` — constant
- `STATE_SYSTEM_UNAVAILABLE` — constant
- `EVENT_OBJECT_FOCUS` — constant
- `EVENT_OBJECT_STATECHANGE` — constant
- `EVENT_OBJECT_VALUECHANGE` — constant
- `EVENT_OBJECT_NAMECHANGE` — constant
- `EVENT_OBJECT_LIVEREGIONCHANGED` — constant
### `Vec2(x, y)`

### `Vec4(x, y, w, h)`

### `Rect2(x, y, width, height)`

### `vec2Add(a, b)`

### `vec2Sub(a, b)`

### `vec2Scale(v, s)`

### `vec2Dot(a, b)`

### `vec2Len(v)`

### `vec2Normalize(v)`

### `rectTranslate(rect, dx, dy)`

### `rectContains(rect, point)`

### `rectIntersects(a, b)`

### `Transform2D(options)`

### `applyTransform2D(point, transform)`

### `Camera2D(options)`

### `worldToScreen2D(point, camera, window)`

### `screenToWorld2D(point, camera, window)`

- `_depsLine` · `{1 entries}`
- `_depsLineBatch` · `{2 entries}`
- `_depsLineRect` · `{2 entries}`
- `_depsLineEllipse` · `{2 entries}`
- `_depsLinePoly` · `{2 entries}`
- `_depsLineBatchPoly` · `{3 entries}`
- `_depsLineRectRounded` · `{3 entries}`
- `_depsLineRectEllipse` · `{3 entries}`
### `drawRect2D(window, x, y, width, height, color, filled, borderColor)`

### `drawCircle2D(window, cx, cy, radius, color, filled, borderColor)`

### `drawPolyline2D(window, points, color, closed)`

### `drawPolygon2D(window, points, color, filled, borderColor)`

### `drawGrid2D(window, spacing, color, originX, originY)`

### `drawEllipse2D(window, cx, cy, rx, ry, color, filled, borderColor)`

### `drawArc2D(window, cx, cy, radius, startAngle, endAngle, color)`

### `drawTriangle2D(window, x1, y1, x2, y2, x3, y3, color, filled, borderColor)`

### `drawRoundedRect2D(window, x, y, width, height, radius, color, filled, borderColor)`

### `drawStar2D(window, cx, cy, outerR, innerR, points, color, filled, borderColor)`

### `drawBezier2D(window, points, color, steps)`

### `drawRing2D(window, cx, cy, outerR, innerR, color)`

### `drawCross2D(window, cx, cy, size, thickness, color, filled)`

### `drawDiamond2D(window, cx, cy, width, height, color, filled)`

### `drawArrow2D(window, x1, y1, x2, y2, color, headSize)`

### `drawCapsule2D(window, cx, cy, width, height, color, filled)`

### `drawSector2D(window, cx, cy, radius, startAngle, endAngle, color, filled)`

### `drawRegularPolygon2D(window, cx, cy, radius, sides, color, filled)`

### `drawSpiral2D(window, cx, cy, startRadius, growth, turns, color)`

### `drawThickLine2D(window, x1, y1, x2, y2, thickness, color)`

### `drawDashedLine2D(window, x1, y1, x2, y2, color, dashLen, gapLen)`

### `Element(bounds)`

### `drawLineAA(window, x1, y1, x2, y2, color, bgColor)`

### `drawCircleFilledAA(window, cx, cy, radius, color, bgColor)`

### `drawCircleOutlineAA(window, cx, cy, radius, color, bgColor)`

### `drawEllipseFilledAA(window, cx, cy, rx, ry, color, bgColor)`

### `drawRoundedRectFilledAA(window, x, y, w, h, radius, color, bgColor)`

### `drawTriangleFilledAA(window, p0, p1, p2, color, bgColor)`

- `_drawOpsDeps` · `?`
### `_initDrawOpsDeps()`

> returns `:object`

### `draw(window, op)`

### `drawBatch(window, ops)`

### `drawBatchParallel(window, ops, numWorkers)`

### `drawThreaded(window, op)`

### `drawBatchThreaded(window, ops)`

### `drawBatchParallelThreaded(window, ops, numWorkers)`

### `CubeMesh(size)`

> returns `:object`

### `drawTriangleFilled(window, p0, p1, p2, color)`

- `_depsMesh` · `{5 entries}`
- `_depsMeshWire` · `{3 entries}`
- `_depsRgb` · `{1 entries}`
### `drawMeshSolid(window, mesh, transform, camera, color, light, borderColor)`

### `drawMeshLit(window, mesh, transform, camera, color, scene, material, borderColor)`

### `DirectionalLight(options)`

### `PointLight(options)`

### `SpotLight(options)`

### `AmbientLight(options)`

### `Material(options)`

### `LightScene(options)`

### `addLight(scene, light)`

### `removeLight(scene, index)`

### `clearLights(scene)`

### `lightCount(scene)`

### `faceNormal(pa, pb, pc)`

### `faceCenter(pa, pb, pc)`

### `shadeFaceColor(baseColor, scene, material, pa3, pb3, pc3, camPos)`

### `shadeFaceIntensity(scene, material, pa3, pb3, pc3, camPos)`

### `prepareScene(scene, material, camPos)`

### `shadeFaceColorFast(baseColor, prep, pa3, pb3, pc3)`

### `Mesh(vertices, edges)`

### `GridMesh(size, step)`

### `AxesMesh(length)`

### `VoxelMesh(voxels, voxelSize)`

### `VoxelGrid(options)`

### `SphereMesh(radius, segments, rings)`

### `PyramidMesh(base, height)`

### `CylinderMesh(radius, height, segments)`

### `ConeMesh(radius, height, segments)`

### `TorusMesh(majorRadius, minorRadius, majorSegments, minorSegments)`

### `PlaneMesh(width, depth, subdivisionsW, subdivisionsD)`

### `HemisphereMesh(radius, segments, rings)`

### `WedgeMesh(width, height, depth)`

### `TubeMesh(outerRadius, innerRadius, height, segments)`

### `ArrowMesh(shaftRadius, shaftHeight, headRadius, headHeight, segments)`

### `PrismMesh(radius, height, sides)`

### `StairsMesh(steps, width, stepHeight, stepDepth)`

### `IcosphereMesh(radius)`

### `drawMeshWireframe(window, mesh, transform, camera, color)`

### `Renderer3D(window, options)`

### `poll(window)`

### `_frameIntervalNt(window)`

### `_frameMaxDtNt(window)`

### `_markUrgentFrameIfResizeDispatch(window, step)`

### `_maybeRunFrame(window, onFrame, force)`

### `_sleepUntilNextFrame(window)`

### `run(window, onEvent, onFrame)`

### `close(window)`

> returns `:int`

### `initThreading(window, options)`

### `destroyThreading(window)`

### `threadingEnabled?(window)`

### `threadCommandQueue(window)`

### `threadWorkerPool(window)`

### `threadScheduler(window)`

### `threadStateGuard(window)`

### `flushThreadedCommands(window)`

### `emitThreadSafe(window, event, payload)`

### `emitFromWorker(window, event, payload)`

### `DrawQueue()`

### `parallelFillRects(window, coords, computeFn, numWorkers)`

### `CommandQueue()`

### `WorkerPool(numWorkers)`

### `FrameFence(workerCount)`

### `FrameScheduler(pool, cmdQueue)`

### `StateGuard()`

### `AsyncLoader(cmdQueue)`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

- `shaderPI` — constant
- `shaderTAU` — constant
- `shaderHALF_PI` — constant
- `shaderE` — constant
- `shaderDEG2RAD` — constant
- `shaderRAD2DEG` — constant
- `shaderSQRT2` — constant
### `shaderFract(x)`

### `shaderMod(x, y)`

### `shaderSign(x)`

### `shaderAbs2(x)`

### `shaderClamp(x, lo, hi)`

### `shaderSaturate(x)`

### `shaderLerpFloat(a, b, t)`

### `shaderInverseLerp(a, b, x)`

### `shaderRemap(x, inLo, inHi, outLo, outHi)`

### `shaderStep(edge, x)`

### `shaderSmoothstep(edge0, edge1, x)`

### `shaderSmootherstep(edge0, edge1, x)`

### `shaderMin2(a, b)`

### `shaderMax2(a, b)`

### `shaderSqr(x)`

### `shaderSqrt(x)`

### `shaderLerp(a, b, t)`

### `shaderAtan2(y, x)`

### `shaderPingpong(t, length)`

### `shaderDegToRad(d)`

### `shaderRadToDeg(r)`

### `shaderEaseInQuad(t)`

### `shaderEaseOutQuad(t)`

### `shaderEaseInOutQuad(t)`

### `shaderEaseInCubic(t)`

### `shaderEaseOutCubic(t)`

### `shaderEaseInOutCubic(t)`

### `shaderEaseInSine(t)`

### `shaderEaseOutSine(t)`

### `shaderEaseInOutSine(t)`

### `shaderEaseInExpo(t)`

### `shaderEaseOutExpo(t)`

### `shaderEaseOutElastic(t)`

### `shaderEaseOutBounce(t)`

### `shaderVec2(x, y)`

### `shaderDot2(a, b)`

### `shaderLength2(v)`

### `shaderDistance2(a, b)`

### `shaderNormalize2(v)`

### `shaderRotate2(v, angle)`

### `shaderScale2(v, s)`

### `shaderAdd2(a, b)`

### `shaderSub2(a, b)`

### `shaderLerp2(a, b, t)`

### `shaderNegate2(v)`

### `shaderAbs2v(v)`

### `shaderMin2v(a, b)`

### `shaderMax2v(a, b)`

### `shaderFloor2(v)`

### `shaderFract2(v)`

### `shaderReflect2(v, n)`

### `shaderToPolar(v)`

### `shaderFromPolar(r, theta)`

### `shaderVec3(x, y, z)`

### `shaderAdd3(a, b)`

### `shaderSub3(a, b)`

### `shaderScale3(v, s)`

### `shaderDot3(a, b)`

### `shaderLength3(v)`

### `shaderDistance3(a, b)`

### `shaderNormalize3(v)`

### `shaderCross3(a, b)`

### `shaderLerp3(a, b, t)`

### `shaderNegate3(v)`

### `shaderReflect3(v, n)`

### `shaderPackRGB(r, g, b)`

### `shaderUnpackRGB(c)`

### `shaderColorR(c)`

### `shaderColorG(c)`

### `shaderColorB(c)`

### `shaderMix(a, b, t)`

### `shaderMix3(a, b, c, t)`

### `shaderBrighten(c, amount)`

### `shaderDarken(c, amount)`

### `shaderInvert(c)`

### `shaderGrayscale(c)`

### `shaderOverlay(fg, bg, alpha)`

### `shaderHsl2rgb(h, s, l)`

### `shaderRgb2hsl(c)`

### `shaderHsv2rgb(h, s, v)`

### `shaderRgb2hsv(c)`

### `shaderFloatStr(c)`

### `shaderCosinePalette(t, a, b, c, d)`

### `shaderContrast(c, amount)`

### `shaderSepia(c)`

### `shaderBlendMultiply(a, b)`

### `shaderBlendScreen(a, b)`

### `shaderBlendAdd(a, b)`

### `shaderHash(seed)`

### `shaderHash2(a, b)`

### `shaderHash3(a, b, c)`

### `shaderNoise2D(x, y)`

### `shaderFbm(x, y, octaves?)`

### `shaderNoiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `shaderFbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

### `shaderSdCircle(px, py, cx, cy, r)`

### `shaderSdBox(px, py, cx, cy, hw, hh)`

### `shaderSdLine(px, py, ax, ay, bx, by)`

### `shaderSdRoundedBox(px, py, cx, cy, hw, hh, r)`

### `shaderSdfFill(d, color)`

### `shaderSdfSmoothFill(d, color, bg, edge)`

### `shaderSdfStroke(d, thickness, color)`

### `shaderSdfGlow(d, color, intensity, radius)`

### `shaderSdUnion(d1, d2)`

### `shaderSdIntersect(d1, d2)`

### `shaderSdSubtract(d1, d2)`

### `shaderSdSmoothUnion(d1, d2, k)`

### `shaderSdSmoothIntersect(d1, d2, k)`

### `shaderSdSmoothSubtract(d1, d2, k)`

### `shaderSdAnnular(d, r)`

### `shaderSdRepeat2(px, py, cx, cy)`

### `shaderCheckerboard(x, y, size)`

### `shaderStripes(x, y, angle, width)`

### `shaderGrid(x, y, size, thickness)`

### `shaderDots(x, y, spacing, radius)`

### `shaderVoronoi(x, y, scale_)`

### `shaderGlslVersion(ver?)`

### `shaderGlslPrecision(prec?, type?)`

### `shaderGlslStdUniforms()`

### `shaderGlslUniform(type, name)`

### `shaderGlslUniforms(uniforms)`

### `shaderGlslIn(type, name)`

### `shaderGlslOut(type, name)`

### `shaderGlslQuadVertex()`

### `shaderGlslQuadVertexCompat()`

### `shaderGlslFragmentWrap(body, version?)`

### `shaderGlslMathLib()`

### `shaderHlslStdCBuffer()`

### `shaderHlslCBuffer(name, uniforms)`

### `shaderHlslQuadVertex()`

### `shaderHlslFragmentWrap(body)`

### `shaderHlslMathLib()`

### `shaderSubmitWebGL(window, fragSource, vertSource?)`

### `shaderDrawWebGL(window, clearR?, clearG?, clearB?)`

### `shaderRenderWebGL(window, fragSource)`

### `shaderCompileGLSL(source, stage?, outputPath?)`

### `shaderCompileHLSL(source, profile?, entry?, outputPath?)`

### `shaderCompileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `shaderAssembleGLSL(opts)`

### `shaderAssembleHLSL(opts)`

### `Shader(fragment?, opts?)`

### `shaderElapsed(shader)`

### `shaderPause(shader)`

### `shaderResume(shader)`

### `shaderReset(shader)`

### `shaderSetUniform(shader, key, value)`

### `shaderGetUniform(shader, key)`

### `shaderSetResolution(shader, res)`

### `shaderBeginFrame(shader)`

### `shaderEndFrame(shader)`

### `shaderDt(shader)`

### `shaderIsRunning(shader)`

### `shaderFrameCount(shader)`

### `shaderRegisteredCount()`

### `shaderUnregisterShader(shader)`

### `shaderClearAll()`

### `shaderDestroyAll()`

### `shaderRender(window, shader, x, y, w, h)`

### `shaderRenderShader(window, shader, x, y, w, h)`

### `shaderRenderShaderLines(window, shader, x, y, w, h)`

### `shaderRenderGradientBands(window, gradientFn, x, y, w, h, time, bands?)`

### `shaderRenderGradient(window, gradientFn, x, y, w, h, time)`

### `shaderRenderHorizontalBands(window, gradientFn, x, y, w, h, time, bands?)`

### `shaderRenderColumns(window, columns, ox, oy, h)`

### `shaderUpdateColumns(columns, h, t, rate?)`

### `ShaderPass(shader, x, y, w, h)`

### `shaderComposePasses(window, passes)`

### `shaderCreateBuffer(w, h)`

### `shaderClearBuffer(buf, color?)`

### `shaderSetPixel(buf, x, y, color)`

### `shaderGetPixel(buf, x, y)`

### `shaderRenderBuffer(window, buf, ox, oy)`

### `shaderRenderShaderToBuffer(buf, shader)`

- `fontFW_THIN` — constant
- `fontFW_EXTRALIGHT` — constant
- `fontFW_LIGHT` — constant
- `fontFW_NORMAL` — constant
- `fontFW_MEDIUM` — constant
- `fontFW_SEMIBOLD` — constant
- `fontFW_BOLD` — constant
- `fontFW_EXTRABOLD` — constant
- `fontFW_HEAVY` — constant
### `defaultFontSpec()`

### `createFontFromSpec(spec)`

### `deleteFontFromSpec(fontResult)`

### `fontKey(spec)`

### `cachedFont(windowOrDisplay, spec)`

### `releaseCachedFonts(windowOrDisplay)`

### `measureTextEx(windowOrDisplay, spec, text)`

### `selectFontEx(args)`

### `fontGetTextMetrics(hdc)`

### `fontLineHeight(hdcOrFontStruct)`

### `buildXLFD(spec)`

### `videoFrameToBuffer(frame)`

### `videoBufferToFrame(buf)`

### `videoRenderFrame(window, frame, ox, oy)`

### `videoRenderFrameScaled(window, frame, ox, oy, scale)`

### `videoBufferGrayscale(buf)`

### `videoBufferInvert(buf)`

### `videoBufferThreshold(buf, t)`

### `videoBufferBlend(bufA, bufB, alpha)`

### `videoBufferDiff(bufA, bufB)`

### `videoCaptureBuffer(buf)`

### `videoFrameToBmpPixels(frame)`

### `createCanvas(window, options)`

### `beginCanvas(canvas)`

### `endCanvas(canvas)`

### `moveCanvas(canvas, x, y)`

### `resizeCanvas(canvas, w, h)`

### `setCanvasVisible(canvas, vis)`

### `setCanvasZIndex(canvas, z)`

### `setCanvasOpacity(canvas, alpha)`

### `setCanvasTransparentColor(canvas, color)`

### `destroyCanvas(canvas)`

### `destroyAllCanvases(window)`

### `isCanvas?(obj)`

### `canvases(window)`

### `canvasCount(window)`

### `canvasAt(window, px, py)`

### `canvasHitTest?(canvas, px, py)`

### `canvasToLocal(canvas, px, py)`

### `addJumpListTask(title, path, arguments, iconPath, iconIndex, description)`

### `clearJumpList()`

### `addJumpListRecentFile(filePath)`

### `registerFileAssociation(ext, progId, desc, cmd, icon)`

### `unregisterFileAssociation(ext, progId)`

### `refreshShellAssociations()`

### `addSearchFolder(folderPath, scope)`

### `searchFiles(query, maxResults)`

### `searchFilesWithProperty(query, property, maxResults)`

### `spatialAudioSource(id, x, y, volume)`

### `spatialAudioUpdate(source, lx, ly, vw, vh, maxDist)`

### `spatialApplyToSamples(samples, pan, gain)`

### `spatialMixSources(sources, lx, ly, vw, vh, maxDist, bufLen)`

### `setAppVolumeName(displayName)`

### `getSystemVolume()`

### `setSystemVolume(level)`

### `enableMediaTransportControls(options)`

### `setMediaPlaybackStatus(status)`

### `updateMediaInfo(title, artist, albumTitle)`

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

### `testCase(suite, name, testFn)`

### `testRun(suite)`

### `testAssert(result, condition, message)`

### `testAssertEqual(result, actual, expected, message)`

### `testReport(suite)`

### `testGetWindowRect(hwnd)`

### `testGetWindowText(hwnd)`

### `testIsWindowVisible?(hwnd)`

### `testGetForegroundWindow()`

### `testFindWindow(cls, name)`

### `testScreenshot(hwnd)`

### `testFreeScreenshot(ss)`

- `TEST_VK_RETURN` — constant
- `TEST_VK_ESCAPE` — constant
- `TEST_VK_TAB` — constant
- `TEST_VK_BACK` — constant
- `TEST_VK_SPACE` — constant
- `TEST_VK_LEFT` — constant
- `TEST_VK_UP` — constant
- `TEST_VK_RIGHT` — constant
- `TEST_VK_DOWN` — constant
- `TEST_VK_SHIFT` — constant
- `TEST_VK_CONTROL` — constant
- `TEST_VK_MENU` — constant
- `TEST_VK_DELETE` — constant
- `TEST_VK_HOME` — constant
- `TEST_VK_END` — constant
### `getGPUAdapters()`

### `getGPUAdaptersParsed()`

### `getDXGIAdapters()`

### `getD3DFeatureLevel()`

### `getDisplayModes()`

### `getMonitorInfo()`

### `gpuCapabilityDump()`

### `gpuCapabilityDumpParallel()`

### `getGDIObjectCount()`

### `getUserObjectCount()`

### `getGDIObjectPeak()`

### `getUserObjectPeak()`

### `getHandleCount()`

### `getWorkingSetSize()`

### `getPeakWorkingSetSize()`

### `leakDetectorState()`

### `leakSnapshot(state)`

### `leakSnapshotParallel(state)`

### `leakCheck(state)`

### `leakReport(state)`

### `leakSetThresholds(state, thresholds)`

### `leakResetBaseline(state)`

### `leakTrend(state)`

- `LEAK_GR_GDIOBJECTS` — constant
- `LEAK_GR_USEROBJECTS` — constant
- `LEAK_GR_GDIOBJECTS_PEAK` — constant
- `LEAK_GR_USEROBJECTS_PEAK` — constant
