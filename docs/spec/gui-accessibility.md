# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-accessibility.oak`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `ROLE_SYSTEM_TITLEBAR` · `1`
- `ROLE_SYSTEM_MENUBAR` · `2`
- `ROLE_SYSTEM_SCROLLBAR` · `3`
- `ROLE_SYSTEM_GRIP` · `4`
- `ROLE_SYSTEM_SOUND` · `5`
- `ROLE_SYSTEM_CURSOR` · `6`
- `ROLE_SYSTEM_CARET` · `7`
- `ROLE_SYSTEM_ALERT` · `8`
- `ROLE_SYSTEM_WINDOW` · `9`
- `ROLE_SYSTEM_CLIENT` · `10`
- `ROLE_SYSTEM_MENUPOPUP` · `11`
- `ROLE_SYSTEM_MENUITEM` · `12`
- `ROLE_SYSTEM_TOOLTIP` · `13`
- `ROLE_SYSTEM_APPLICATION` · `14`
- `ROLE_SYSTEM_DOCUMENT` · `15`
- `ROLE_SYSTEM_PANE` · `16`
- `ROLE_SYSTEM_CHART` · `17`
- `ROLE_SYSTEM_DIALOG` · `18`
- `ROLE_SYSTEM_BORDER` · `19`
- `ROLE_SYSTEM_GROUPING` · `20`
- `ROLE_SYSTEM_SEPARATOR` · `21`
- `ROLE_SYSTEM_TOOLBAR` · `22`
- `ROLE_SYSTEM_STATUSBAR` · `23`
- `ROLE_SYSTEM_TABLE` · `24`
- `ROLE_SYSTEM_COLUMNHEADER` · `25`
- `ROLE_SYSTEM_ROWHEADER` · `26`
- `ROLE_SYSTEM_COLUMN` · `27`
- `ROLE_SYSTEM_ROW` · `28`
- `ROLE_SYSTEM_CELL` · `29`
- `ROLE_SYSTEM_LINK` · `30`
- `ROLE_SYSTEM_HELPBALLOON` · `31`
- `ROLE_SYSTEM_CHARACTER` · `32`
- `ROLE_SYSTEM_LIST` · `33`
- `ROLE_SYSTEM_LISTITEM` · `34`
- `ROLE_SYSTEM_OUTLINE` · `35`
- `ROLE_SYSTEM_OUTLINEITEM` · `36`
- `ROLE_SYSTEM_PAGETAB` · `37`
- `ROLE_SYSTEM_PROPERTYPAGE` · `38`
- `ROLE_SYSTEM_INDICATOR` · `39`
- `ROLE_SYSTEM_GRAPHIC` · `40`
- `ROLE_SYSTEM_STATICTEXT` · `41`
- `ROLE_SYSTEM_TEXT` · `42`
- `ROLE_SYSTEM_PUSHBUTTON` · `43`
- `ROLE_SYSTEM_CHECKBUTTON` · `44`
- `ROLE_SYSTEM_RADIOBUTTON` · `45`
- `ROLE_SYSTEM_COMBOBOX` · `46`
- `ROLE_SYSTEM_DROPLIST` · `47`
- `ROLE_SYSTEM_PROGRESSBAR` · `48`
- `ROLE_SYSTEM_DIAL` · `49`
- `ROLE_SYSTEM_HOTKEYFIELD` · `50`
- `ROLE_SYSTEM_SLIDER` · `51`
- `ROLE_SYSTEM_SPINBUTTON` · `52`
- `ROLE_SYSTEM_DIAGRAM` · `53`
- `ROLE_SYSTEM_ANIMATION` · `54`
- `ROLE_SYSTEM_EQUATION` · `55`
- `ROLE_SYSTEM_BUTTONDROPDOWN` · `56`
- `ROLE_SYSTEM_BUTTONMENU` · `57`
- `ROLE_SYSTEM_BUTTONDROPDOWNGRID` · `58`
- `ROLE_SYSTEM_WHITESPACE` · `59`
- `ROLE_SYSTEM_PAGETABLIST` · `60`
- `ROLE_SYSTEM_CLOCK` · `61`
- `ROLE_SYSTEM_SPLITBUTTON` · `62`
- `STATE_SYSTEM_NORMAL` · `0`
- `STATE_SYSTEM_UNAVAILABLE` · `1`
- `STATE_SYSTEM_SELECTED` · `2`
- `STATE_SYSTEM_FOCUSED` · `4`
- `STATE_SYSTEM_PRESSED` · `8`
- `STATE_SYSTEM_CHECKED` · `16`
- `STATE_SYSTEM_MIXED` · `32`
- `STATE_SYSTEM_READONLY` · `64`
- `STATE_SYSTEM_HOTTRACKED` · `128`
- `STATE_SYSTEM_DEFAULT` · `256`
- `STATE_SYSTEM_EXPANDED` · `512`
- `STATE_SYSTEM_COLLAPSED` · `1024`
- `STATE_SYSTEM_BUSY` · `2048`
- `STATE_SYSTEM_INVISIBLE` · `32768`
- `STATE_SYSTEM_OFFSCREEN` · `65536`
- `STATE_SYSTEM_SIZEABLE` · `131072`
- `STATE_SYSTEM_MOVEABLE` · `262144`
- `STATE_SYSTEM_FOCUSABLE` · `1048576`
- `STATE_SYSTEM_SELECTABLE` · `2097152`
- `STATE_SYSTEM_LINKED` · `4194304`
- `STATE_SYSTEM_TRAVERSED` · `8388608`
- `STATE_SYSTEM_MULTISELECTABLE` · `16777216`
- `STATE_SYSTEM_HASPOPUP` · `1073741824`
- `EVENT_SYSTEM_SOUND` · `1`
- `EVENT_SYSTEM_ALERT` · `2`
- `EVENT_SYSTEM_FOREGROUND` · `3`
- `EVENT_SYSTEM_MENUSTART` · `4`
- `EVENT_SYSTEM_MENUEND` · `5`
- `EVENT_SYSTEM_MENUPOPUPSTART` · `6`
- `EVENT_SYSTEM_MENUPOPUPEND` · `7`
- `EVENT_SYSTEM_CAPTURESTART` · `8`
- `EVENT_SYSTEM_CAPTUREEND` · `9`
- `EVENT_SYSTEM_MOVESIZESTART` · `10`
- `EVENT_SYSTEM_MOVESIZEEND` · `11`
- `EVENT_SYSTEM_CONTEXTHELPSTART` · `12`
- `EVENT_SYSTEM_CONTEXTHELPEND` · `13`
- `EVENT_SYSTEM_DRAGDROPSTART` · `14`
- `EVENT_SYSTEM_DRAGDROPEND` · `15`
- `EVENT_SYSTEM_DIALOGSTART` · `16`
- `EVENT_SYSTEM_DIALOGEND` · `17`
- `EVENT_SYSTEM_SCROLLINGSTART` · `18`
- `EVENT_SYSTEM_SCROLLINGEND` · `19`
- `EVENT_SYSTEM_SWITCHSTART` · `20`
- `EVENT_SYSTEM_SWITCHEND` · `21`
- `EVENT_OBJECT_CREATE` · `32768`
- `EVENT_OBJECT_DESTROY` · `32769`
- `EVENT_OBJECT_SHOW` · `32770`
- `EVENT_OBJECT_HIDE` · `32771`
- `EVENT_OBJECT_REORDER` · `32772`
- `EVENT_OBJECT_FOCUS` · `32773`
- `EVENT_OBJECT_SELECTION` · `32774`
- `EVENT_OBJECT_SELECTIONADD` · `32775`
- `EVENT_OBJECT_SELECTIONREMOVE` · `32776`
- `EVENT_OBJECT_SELECTIONWITHIN` · `32777`
- `EVENT_OBJECT_STATECHANGE` · `32778`
- `EVENT_OBJECT_LOCATIONCHANGE` · `32779`
- `EVENT_OBJECT_NAMECHANGE` · `32780`
- `EVENT_OBJECT_DESCRIPTIONCHANGE` · `32781`
- `EVENT_OBJECT_VALUECHANGE` · `32782`
- `EVENT_OBJECT_PARENTCHANGE` · `32783`
- `EVENT_OBJECT_HELPCHANGE` · `32784`
- `EVENT_OBJECT_DEFACTIONCHANGE` · `32785`
- `EVENT_OBJECT_ACCELERATORCHANGE` · `32786`
- `EVENT_OBJECT_INVOKED` · `32787`
- `EVENT_OBJECT_TEXTSELECTIONCHANGED` · `32788`
- `EVENT_OBJECT_CONTENTSCROLLED` · `32789`
- `EVENT_OBJECT_LIVEREGIONCHANGED` · `32793`
- `WM_GETOBJECT` · `61`
- `OBJID_SELF` · `0`
- `OBJID_WINDOW` · `0`
- `OBJID_CLIENT` — constant
- `OBJID_NATIVEOM` — constant
- `_UIA_UIAROOT_OBJECTID` — constant
- `_oleaccProcCache` · `{}`
- `Oleacc` · `'oleacc.dll'`
### `oleacc(symbol, args...)`

### `accNode(id, name, role, bounds, state, children)`

> returns `:object`

### `accTree(window)`

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

### `enableAccessibility(window)`

> returns `:bool`

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

> returns `:object`

### `accRoleName(role)`

> returns `:string`

### `accStateName(state)`

> returns `:string`

