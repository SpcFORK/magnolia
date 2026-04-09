# API Documentation

_Auto-generated from Magnolia source._

---

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

