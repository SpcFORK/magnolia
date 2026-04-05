# Windows GUI — Features Yet To Be Added

Tracked features not yet implemented in the Magnolia Windows GUI subsystem.

---

## Rendering Backends

- [x] **Vulkan swapchain presentation** — supports resolution scaling via intermediate VkImage + vkCmdBlitImage; falls back to direct buffer copy when no scaling is active
- [x] **OpenGL SwapBuffers presentation** — resolution scaling via FBO + glBlitFramebuffer (GL 3.0+); graceful fallback when FBO extensions unavailable
- [x] **Direct3D 11 backend** — D3D11CreateDeviceAndSwapChain with BGRA swap chain; DXGI handles resolution scaling; WARP software driver fallback
- [x] **Software rasterizer improvements** — Sutherland-Hodgman viewport triangle clipping eliminates wasted scanline iteration for off-screen geometry; fast path bypasses clipping for fully-visible triangles

---

## Input Handling

- [x] **Mouse wheel scrolling** — `onMouseWheel(window, handler)` in gui-events.oak
- [x] **Middle mouse button** — `onMButtonDown`/`onMButtonUp` in gui-events.oak
- [x] **Double-click detection** — `onLButtonDblClk` + CS_DBLCLKS class style
- [x] **Mouse hover/leave tracking** — `onMouseHover`/`onMouseLeave` via TrackMouseEvent
- [x] **Touch / pen / stylus input** — `registerTouchWindow`/`onTouch`/`enableTouchInput`/`onPointerDown`/`onPointerUp`/`onPointerUpdate`/`enablePointerInput` in gui-events.oak
- [x] **Modifier key state queries** — `modShift?`, `modCtrl?`, `modAlt?`, `keyShiftDown?`, `keyCtrlDown?`, `keyAltDown?`
- [x] **Gamepad / joystick input** — `getGamepadState`/`gamepadPoll`/`setGamepadVibration` + 14 button helpers in gui-gamepad.oak
- [x] **Raw input (HID)** — `registerRawInputDevice`/`onRawInput`/`enableRawMouse`/`enableRawKeyboard`/`enableRawGamepad` in gui-events.oak

---

## Clipboard & Data Transfer

- [x] **Clipboard read / write** — `clipboardGetText`/`clipboardSetText`/`clipboardHasText` in gui-clipboard.oak
- [x] **Clipboard integration in form fields** — `formCopyByFocus`/`formPasteByFocus`/`formCutByFocus` in gui-form.oak
- [x] **OLE drag-and-drop** — `enableOleDrop`/`disableOleDrop`/`onOleDrop`/`onDragOver` in gui-filedrop.oak
- [x] **File drop events** — `enableFileDrop`/`onFileDrop` in gui-filedrop.oak

---

## Native Dialogs

- [x] **File Open / Save dialog** — `openFileDialog`/`saveFileDialog` in gui-dialogs.oak
- [x] **Folder picker** — `pickFolder` in gui-dialogs.oak (SHBrowseForFolderW)
- [x] **Color chooser dialog** — `chooseColor` in gui-dialogs.oak (ChooseColorW)
- [x] **Font chooser dialog** — `chooseFont` in gui-dialogs.oak (ChooseFontW)
- [x] **Message box** — `messageBox` in windows-windowing.oak
- [x] **Print dialog** — `showPrintDialog` (PRINTDLGW) via comdlg32 in gui-print.oak

- [x] **Menu bar** — `createMenu`/`setMenuBar`/`buildMenu` in gui-menus.oak
- [x] **Context (popup) menu** — `createPopupMenu`/`showPopupMenu` in gui-menus.oak
- [x] **Keyboard accelerators** — `createAcceleratorTable`/`installAccelerators` in gui-menus.oak + poll loop integration
- [x] **System menu customization** — `getSystemMenu`/`appendSystemMenuItem`/`onSysCommand` in gui-menus.oak

---

## Window Management

- [x] **Multi-window support** — `registerWindow`/`unregisterWindow`/`pollAllWindows`/`allWindows` in gui-native-win.oak
- [x] **Child / owned windows** — `createOwnedWindow`/`showModalDialog`/`closeModalDialog`/`setWindowOwner` in gui-native-win.oak
- [x] **Multi-monitor awareness** — `getWindowMonitor`/`centerOnMonitor`/`moveToMonitor`/`getWindowDpi` in gui-native-win.oak
- [x] **Window snap / restore state** — `saveWindowState`/`restoreWindowState` in gui-native-win.oak
- [x] **Minimize-to-tray** — `minimizeToTray`/`restoreFromTray`/`addTrayIcon`/`onTrayEvent` in gui-systray.oak
- [x] **Custom window chrome** — `extendFrameIntoClientArea`/`enableGlassSheet`/`onNcHitTest`/`customChromeHitTest` in gui-native-win.oak
- [x] **Layered / transparent windows** — `setWindowOpacity`/`setWindowColorKey`/`removeLayeredStyle` in gui-native-win.oak
- [x] **Always-on-top toggle** — `setAlwaysOnTop(window, enabled)` via SetWindowPos + HWND_TOPMOST
- [x] **Taskbar progress indicator** — `setTaskbarProgress`/`setTaskbarProgressState` via ITaskbarList3 in gui-native-win.oak

---

## DPI & Display Scaling

- [x] **Per-monitor DPI awareness** — `setProcessDpiAwarenessContext`/`getDpiForWindow`/`getDpiForSystem`/`getDpiForMonitor` in windows-windowing.oak
- [x] **WM_DPICHANGED handling** — `onDpiChanged` handler with automatic SetWindowPos resize in gui-events.oak
- [x] **Logical-to-physical coordinate mapping** — `dpiScale`/`dpiUnscale` helpers in windows-windowing.oak
- [x] **Scaled font rendering** — `getSystemMetricsForDpi` + DPI-aware coordinate mapping in windows-windowing.oak

---

## Text & IME

- [x] **IME support** — `onImeStartComposition`/`onImeEndComposition`/`onImeComposition`/`setImePosition` in gui-events.oak
- [x] **Text composition events** — WM_IME_COMPOSITION fires `:composing`/`:result` events with text and cursor position
- [x] **Text selection in fields** — `formSelectionState`/`formSelMoveCursor`/`formDrawFieldWithSel` in gui-form.oak
- [x] **Undo / redo in text fields** — `formUndoState`/`formUndoPush`/`formUndo`/`formRedo` in gui-form.oak
- [x] **Rich text / multi-style spans** — `formRichTextState`/`formDrawRichText`/`richSpan`/`richBold`/`richItalic`/`richUnderline`/`richColored`/`richStyled` in gui-form.oak

---

## UI Widgets (Missing or Incomplete)

- [x] **Native scrollbar widgets** — `formDrawScrollbar`/`formScrollbarHit` (vertical+horizontal) in gui-form.oak
- [x] **Tooltips / hover hints** — `formTooltipState`/`formTooltipUpdate`/`formDrawTooltip` in gui-form.oak
- [x] **Tab strip / tab control** — `formDrawTabStrip`/`formHitTab` in gui-form.oak
- [x] **Tree view** — `formDrawTreeView`/`formTreeHitRow`/`formTreeToggle` in gui-form.oak
- [x] **Dropdown / combo box** — `formDrawDropdown`/`formDrawDropdownList`/`formHitDropdownItem` in gui-form.oak
- [x] **Progress bar** — `formDrawProgressBar` (determinate + indeterminate) in gui-form.oak
- [x] **Spinner / number input** — `formDrawSpinner`/`formSpinnerHit`/`formSpinnerAdjust` in gui-form.oak
- [x] **Date / time picker** — `formDatePickerState`/`formDrawDateField`/`formDrawDateCalendar`/`formDateCalendarHit` in gui-form.oak
- [x] **Table / data grid** — `formDrawTable`/`formTableHitRow`/`formTableHitColumn` in gui-form.oak
- [x] **Focus management** — `formFocusState`/`formFocusNext`/`formFocusPrev`/`formDrawFocusRing`/`formHandleTabKey` in gui-form.oak

---

## Accessibility

- [x] **MSAA / UI Automation provider** — `accNode`/`accTree`/`accRegister`/`notifyAccEvent`/`enableAccessibility` + 16 convenience registrations in gui-accessibility.oak
- [x] **High-contrast mode detection** — `isHighContrast?()` via SystemParametersInfoW in gui-theme.oak
- [x] **Keyboard-only navigation** — `formHandleKeyNav`/`formCheckboxKeyToggle`/`formRadioKeySelect`/`formSliderKeyAdjust`/`formDropdownKeyNav`/`formTreeKeyNav`/`formTableKeyNav` in gui-form.oak
- [x] **Accessible names and roles** — `accSetName`/`accSetState`/`accSetValue`/`accSetBounds`/`accSetDescription`/`accSetDefaultAction`/`accFocus`/`accSelection` in gui-accessibility.oak
- [x] **Narrator / NVDA compatibility testing** — `accDump`/`accVerify`/`accRoleName`/`accStateName`/`accAnnounce` in gui-accessibility.oak

---

## System Integration

- [x] **System tray / notification area icon** — `addTrayIcon`/`removeTrayIcon`/`showBalloonTip`/`onTrayEvent` in gui-systray.oak
- [x] **System notifications (toast)** — `showToastNotification`/`showToastWithFallback` via PowerShell WinRT in gui-native-win.oak
- [x] **Dark mode / system theme detection** — `isDarkMode?()`/`accentColor()` via registry in gui-theme.oak
- [x] **Windows 11 Mica / Acrylic backdrop** — `enableMica`/`enableAcrylic`/`enableTabbedMica`/`disableBackdrop`/`setDwmDarkMode` in gui-native-win.oak
- [x] **Single-instance enforcement** — `acquireSingleInstance`/`releaseSingleInstance` via named mutex in gui-native-win.oak
- [x] **Jump list customization** — `addJumpListTask`/`clearJumpList`/`addJumpListRecentFile` via PowerShell COM in gui-native-win.oak
- [x] **File type association handler** — `registerFileAssociation`/`unregisterFileAssociation`/`refreshShellAssociations` via HKCU registry in gui-native-win.oak
- [x] **Windows Search integration** — `addSearchFolder`/`searchFiles`/`searchFilesWithProperty` via Windows Search COM in gui-native-win.oak

---

## Audio Integration (GUI-specific)

- [x] **Spatial audio tied to GUI viewport** — `spatialAudioSource`/`spatialAudioUpdate`/`spatialApplyToSamples`/`spatialMixSources` in gui-audio.oak
- [x] **System volume mixer integration** — `setAppVolumeName`/`getSystemVolume`/`setSystemVolume` via PowerShell IAudioEndpointVolume COM in gui-audio.oak
- [x] **Media transport controls** — `enableMediaTransportControls`/`setMediaPlaybackStatus`/`updateMediaInfo` via WinRT SystemMediaTransportControls in gui-audio.oak

---

## Printing

- [x] **Print dialog** — `showPrintDialog` (PRINTDLGW) via comdlg32 in gui-print.oak
- [x] **GDI print rendering** — `startDoc`/`startPage`/`endPage`/`endDoc`/`printTextOut`/`printMoveTo`/`printLineTo`/`printRectangle`/`printEllipse`/`printSetFont`/`printSetTextColor`/`printSetBkMode`/`printSetPen` in gui-print.oak
- [x] **Print preview** — `createPreviewDC`/`destroyPreviewDC` (off-screen compatible DC + bitmap) in gui-print.oak
- [x] **PDF export** — `printToFile`/`printDocument` (multi-page with copies via print-to-file) in gui-print.oak

---

## Testing & Diagnostics

- [x] **Automated GUI test harness** — `testMouseMove`/`testMouseClick`/`testKeyPress`/`testTypeText`/`testKeyCombo`/`testSuite`/`testCase`/`testRun`/`testReport` + window introspection in gui-test.oak
- [x] **Frame timing overlay** — `formFrameTimerState`/`formFrameTimerTick`/`formDrawFrameTimingOverlay` in gui-form.oak
- [x] **GPU capability dump** — `getGPUAdapters`/`getGPUAdaptersParsed`/`getDXGIAdapters`/`getD3DFeatureLevel`/`getDisplayModes`/`getMonitorInfo`/`gpuCapabilityDump` in gui-gpu-info.oak
- [x] **Memory / handle leak detector** — `leakDetectorState`/`leakSnapshot`/`leakCheck`/`leakReport`/`leakTrend` + GDI/USER/handle/memory tracking in gui-leak-detect.oak

---

*Last updated: 2026-04-01*
