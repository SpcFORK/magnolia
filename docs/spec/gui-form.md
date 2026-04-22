# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `gui-input`

- `VK_BACK` · `8`
- `VK_TAB` · `9`
- `VK_CLEAR` · `12`
- `VK_RETURN` · `13`
- `VK_SHIFT` · `16`
- `VK_CONTROL` · `17`
- `VK_ALT` · `18`
- `VK_PAUSE` · `19`
- `VK_CAPSLOCK` · `20`
- `VK_ESCAPE` · `27`
- `VK_SPACE` · `32`
- `VK_PAGEUP` · `33`
- `VK_PAGEDOWN` · `34`
- `VK_END` · `35`
- `VK_HOME` · `36`
- `VK_LEFT` · `37`
- `VK_UP` · `38`
- `VK_RIGHT` · `39`
- `VK_DOWN` · `40`
- `VK_INSERT` · `45`
- `VK_DELETE` · `46`
- `VK_0` · `48`
- `VK_1` · `49`
- `VK_2` · `50`
- `VK_3` · `51`
- `VK_4` · `52`
- `VK_5` · `53`
- `VK_6` · `54`
- `VK_7` · `55`
- `VK_8` · `56`
- `VK_9` · `57`
- `VK_A` · `65`
- `VK_B` · `66`
- `VK_C` · `67`
- `VK_D` · `68`
- `VK_E` · `69`
- `VK_F` · `70`
- `VK_G` · `71`
- `VK_H` · `72`
- `VK_I` · `73`
- `VK_J` · `74`
- `VK_K` · `75`
- `VK_L` · `76`
- `VK_M` · `77`
- `VK_N` · `78`
- `VK_O` · `79`
- `VK_P` · `80`
- `VK_Q` · `81`
- `VK_R` · `82`
- `VK_S` · `83`
- `VK_T` · `84`
- `VK_U` · `85`
- `VK_V` · `86`
- `VK_W` · `87`
- `VK_X` · `88`
- `VK_Y` · `89`
- `VK_Z` · `90`
- `VK_NUMPAD0` · `96`
- `VK_NUMPAD1` · `97`
- `VK_NUMPAD2` · `98`
- `VK_NUMPAD3` · `99`
- `VK_NUMPAD4` · `100`
- `VK_NUMPAD5` · `101`
- `VK_NUMPAD6` · `102`
- `VK_NUMPAD7` · `103`
- `VK_NUMPAD8` · `104`
- `VK_NUMPAD9` · `105`
- `VK_MULTIPLY` · `106`
- `VK_ADD` · `107`
- `VK_SEPARATOR` · `108`
- `VK_SUBTRACT` · `109`
- `VK_DECIMAL` · `110`
- `VK_DIVIDE` · `111`
- `VK_F1` · `112`
- `VK_F2` · `113`
- `VK_F3` · `114`
- `VK_F4` · `115`
- `VK_F5` · `116`
- `VK_F6` · `117`
- `VK_F7` · `118`
- `VK_F8` · `119`
- `VK_F9` · `120`
- `VK_F10` · `121`
- `VK_F11` · `122`
- `VK_F12` · `123`
- `VK_NUMLOCK` · `144`
- `VK_SCROLLLOCK` · `145`
- `VK_OEM_SEMICOLON` · `186`
- `VK_OEM_PLUS` · `187`
- `VK_OEM_COMMA` · `188`
- `VK_OEM_MINUS` · `189`
- `VK_OEM_PERIOD` · `190`
- `VK_OEM_SLASH` · `191`
- `VK_OEM_TILDE` · `192`
- `VK_OEM_LBRACKET` · `219`
- `VK_OEM_BACKSLASH` · `220`
- `VK_OEM_RBRACKET` · `221`
- `VK_OEM_QUOTE` · `222`
- `VK_LSHIFT` · `160`
- `VK_RSHIFT` · `161`
- `VK_LCONTROL` · `162`
- `VK_RCONTROL` · `163`
- `VK_LALT` · `164`
- `VK_RALT` · `165`
### `isLetterKey?(vk)`

> returns `:bool`

### `isDigitKey?(vk)`

> returns `:bool`

### `isNumpadKey?(vk)`

> returns `:bool`

### `isFunctionKey?(vk)`

> returns `:bool`

### `isArrowKey?(vk)`

> returns `:bool`

### `isModifierKey?(vk)`

> returns `:bool`

## Module: `gui-thread`

- `threadLib` · `import(...)`
### `CommandQueue()`

> returns `:object`

### `FrameFence(workerCount)`

> returns `:object`

### `WorkerPool(numWorkers)`

> returns `:object`

### `StateGuard()`

> returns `:object`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

### `AsyncLoader(cmdQueue)`

> returns `:object`

### `FrameScheduler(pool, cmdQueue)`

> returns `:object`

### `initWindowThreading(window, options)`

### `threadingEnabled?(window)`

### `commandQueue(window)`

### `workerPool(window)`

### `scheduler(window)`

### `stateGuard(window)`

### `flushThreadedCommands(window)`

### `destroyWindowThreading(window)`

> returns `?`

## Module: `lib\gui-form.oak`

- `std` · `import(...)`
- `guiInput` · `import(...)`
- `guiThread` · `import(...)`
### `formSetStatus(state, message, ok)`

### `formPopChar(s)`

### `formClamp(v, minVal, maxVal)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

> returns `:bool`

### `formHitListIndex(mx, my, x, y, itemW, itemH, count)`

### `formHitRectKey(mx, my, rects)`

### `formToggleIfHit(value, mx, my, x, y, w, h)`

> returns `:bool`

### `formSelectByHit(current, mx, my, rects)`

### `formSetByAssignments(state, assignments)`

### `formResetFlags(state, keys, value)`

### `formSetHoverFromRects(state, mx, my, rects)`

### `formToggleKeysByHit(state, mx, my, rects)`

### `formSetKeyByHit(state, targetKey, mx, my, rects)`

### `formTruncateText(s, maxChars)`

### `formIsPrintableChar?(code)`

> returns `:bool`

- `_charWidth` · `7`
### `formSelectionState()`

> returns `:object`

### `formSelSetCursor(sel, pos, shifting)`

### `formSelMoveCursor(sel, text, dir, shifting)`

### `formSelHome(sel, shifting)`

### `formSelEnd(sel, text, shifting)`

### `formSelAll(sel, text)`

### `formSelRange(sel)`

> returns `:object`

### `formSelHasSelection?(sel)`

### `formSelSelectedText(sel, text)`

### `formSelDeleteSelection(sel, text)`

### `formSelInsertAtCursor(sel, text, insert)`

### `formSelBackspace(sel, text)`

### `formSelClickPos(mx, fieldX, fieldPadding, text)`

### `formDrawFieldWithSel(ctx, window, x, y, w, h, text, placeholder, focused, sel, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor, selColor)`

### `formUndoState(maxHistory)`

> returns `:object`

### `formUndoPush(hist, text, sel)`

### `formUndo(hist, text, sel)`

### `formRedo(hist, text, sel)`

### `formAppendByFocus(state, focus, c, fieldSpecs, notesSpec)`

### `formBackspaceByFocus(state, focus, fieldKeys, notesSpec)`

- `_clipboard` · `?`
### `_cb()`

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

> returns `:object`

### `formFocusNext(fs)`

### `formFocusPrev(fs)`

### `formFocusSet(fs, key)`

### `formFocusIs?(fs, key)`

### `formDrawFocusRing(ctx, window, x, y, w, h, color, options)`

### `formHandleTabKey(fs, shiftDown)`

- `_VK_RETURN` — constant
- `_VK_SPACE` — constant
- `_VK_ESCAPE` — constant
- `_VK_UP` — constant
- `_VK_DOWN` — constant
- `_VK_LEFT` — constant
- `_VK_RIGHT` — constant
- `_VK_HOME` — constant
- `_VK_END` — constant
- `_VK_PGUP` — constant
- `_VK_PGDN` — constant
### `formIsActivateKey?(vk)`

> returns `:bool`

### `formHandleKeyNav(fs, vk, shiftDown, widgetHandlers, options)`

> returns `:atom`

### `formCheckboxKeyToggle(checked, vk)`

> returns `:bool`

### `formRadioKeySelect(current, vk, choices)`

### `formSliderKeyAdjust(value, vk, minVal, maxVal, step)`

### `formDropdownKeyNav(open?, selectedIdx, vk, itemCount)`

> returns `:object`

### `formTreeKeyNav(flatList, selectedIdx, vk)`

> returns `:object`

### `formTableKeyNav(selectedRow, vk, rowCount, pageSize)`

### `formMaskText(s)`

### `formDrawBorder(ctx, window, x, y, w, h, color)`

### `formDrawField(ctx, window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawPasswordField(ctx, window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawCheckbox(ctx, window, x, y, checked, label, fieldColor, borderColor, checkColor)`

### `formDrawRadio(ctx, window, x, y, selected, label, fieldColor, borderColor, accentColor)`

### `formDrawPrimaryButton(ctx, window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, bottomLineColor)`

### `formDrawSecondaryButton(ctx, window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, borderColor)`

### `formDrawSlider(ctx, window, x, y, w, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, options)`

### `formDrawLabeledPercentSlider(ctx, window, x, labelY, sliderY, w, label, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, percentRectColor)`

### `formDrawNotes(ctx, window, x, y, w, h, lines, focused, placeholder, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawStatusBanner(ctx, window, x, y, w, h, message, ok, fieldColor, borderColor, successColor, errorColor)`

### `formDrawProgressBar(ctx, window, x, y, w, value, maxVal, trackColor, fillColor, borderColor, options)`

### `formDrawTabStrip(ctx, window, x, y, tabs, activeIdx, bgColor, activeBgColor, borderColor, activeTextColor, inactiveTextColor, options)`

### `formHitTab(mx, my, x, y, tabs, options)`

### `formDrawDropdown(ctx, window, x, y, w, h, selectedLabel, open, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDropdownList(ctx, window, x, y, w, items, hoverIdx, itemHeight, listBgColor, hoverBgColor, borderColor)`

### `formHitDropdownItem(mx, my, x, y, w, items, itemHeight)`

### `formTooltipState()`

> returns `:object`

### `formTooltipUpdate(state, mx, my, regionKey, text, delay)`

### `formTooltipHide(state)`

### `formDrawTooltip(ctx, window, state, bgColor, borderColor, textColor, options)`

### `formDrawSpinner(ctx, window, x, y, w, h, value, focused, fieldColor, fieldFocusColor, borderColor, arrowColor, options)`

### `formSpinnerHit(mx, my, x, y, w, h, options)`

> returns `:atom`

### `formSpinnerAdjust(value, direction, options)`

### `formDrawScrollbar(ctx, window, x, y, w, h, scrollPos, contentSize, viewSize, trackColor, thumbColor, borderColor, options)`

### `formScrollbarHit(mx, my, x, y, w, h, scrollPos, contentSize, viewSize, options)`

### `_flattenTree(nodes, depth, out)`

### `formDrawTreeView(ctx, window, x, y, w, h, nodes, selectedIdx, scrollPos, bgColor, selectedBgColor, textColor, selectedTextColor, borderColor, options)`

### `formTreeHitRow(mx, my, x, y, w, flat, scrollPos, options)`

### `formTreeToggle(flat, idx)`

> returns `:bool`

### `formTreeContentHeight(flat, options)`

### `_isLeapYear?(y)`

> returns `:bool`

### `_daysInMonth(year, month)`

> returns `:int`

### `_dayOfWeek(y, m, d)`

### `formDatePickerState(year, month, day)`

> returns `:object`

### `formDatePickerPrevMonth(state)`

### `formDatePickerNextMonth(state)`

### `_padZero(n)`

> returns `:string`

- `_monthNames` · `[12 items]`
### `formDateLabel(state)`

### `formDrawDateField(ctx, window, x, y, w, h, state, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDateCalendar(ctx, window, x, y, state, bgColor, headerBgColor, selectedBgColor, todayBorderColor, textColor, selectedTextColor, headerTextColor, borderColor, options)`

### `formDateCalendarHit(mx, my, x, y, state, options)`

> returns `:object`

### `formDrawTable(ctx, window, x, y, w, h, columns, rows, scrollPos, selectedRow, headerBgColor, rowBgColor, altRowBgColor, selectedBgColor, headerTextColor, textColor, selectedTextColor, borderColor, options)`

### `formTableHitRow(mx, my, x, y, w, columns, rows, scrollPos, options)`

### `formTableContentHeight(rows, options)`

### `formTableHitColumn(mx, x, columns)`

### `formFrameTimerState(maxSamples)`

> returns `:object`

### `formFrameTimerTick(state)`

### `formDrawFrameTimingOverlay(ctx, window, state, x, y, w, h, options)`

### `formRichTextState(lines)`

> returns `:object`

### `formRichTextSetLines(state, lines)`

### `formRichTextAppendLine(state, spans)`

### `formRichTextInsertSpan(state, lineIdx, spanIdx, span)`

### `_spanWidth(text, fontSize)`

### `formDrawRichText(ctx, window, x, y, w, h, state, bgColor, borderColor)`

### `formRichTextScroll(state, delta)`

### `formRichTextTotalHeight(state)`

### `richSpan(text)`

> returns `:object`

### `richBold(text)`

> returns `:object`

### `richItalic(text)`

> returns `:object`

### `richUnderline(text)`

> returns `:object`

### `richColored(text, r, g, b)`

> returns `:object`

### `richStyled(text, options)`

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

