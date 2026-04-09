# gui-accessibility — MSAA / UI Automation Support

`import('gui-accessibility')` provides Microsoft Active Accessibility (MSAA) and UI Automation support for the Magnolia GUI, enabling screen readers to read widget names, roles, and state.

## Quick Start

```oak
gui := import('GUI')
acc := import('gui-accessibility')

window := gui.createWindow('Accessible App', 800, 600, {})

// Enable accessibility on the window
acc.enableAccessibility(window)

// Register widgets
acc.accRegisterButton(window, 1, 'Save', 10, 10, 80, 30)
acc.accRegisterTextField(window, 2, 'Name', 10, 50, 200, 24, '')

// Update state dynamically
acc.accSetState(window, 1, acc.STATE_SYSTEM_FOCUSED)
acc.accFocus(window, 1)

// Live announcements
acc.accAnnounce(window, 3, 'File saved successfully')
```

## Accessibility Tree

### `accNode(id, name, role, bounds, state, children)`

Creates an accessible node descriptor.

### `accTree(window)`

Gets or creates the accessibility tree for a window.

### `accRegister(window, id, name, role, bounds, state)`

Registers a widget in the accessibility tree.

### `accUnregister(window, id)`

Removes a widget from the accessibility tree.

### `accSetName(window, id, name)`

Updates the accessible name of a widget.

### `accSetState(window, id, state)`

Updates the state flags of a widget.

### `accSetValue(window, id, value)`

Updates the value of a widget.

### `accSetBounds(window, id, bounds)`

Updates the bounding rectangle of a widget.

### `accSetDescription(window, id, desc)`

Sets the accessible description.

### `accSetDefaultAction(window, id, action)`

Sets the default action label.

### `accFocus(window, id)`

Notifies that a widget has gained focus.

### `accSelection(window, id)`

Notifies that a widget has been selected.

## Event Notification

### `notifyAccEvent(window, event, childId)`

Fires an accessibility event via `NotifyWinEvent`.

### `enableAccessibility(window)`

Sets up `WM_GETOBJECT` handling for the window so screen readers can query the accessibility tree.

## Widget Registration Helpers

Convenience functions that register common widget types with appropriate MSAA roles and states:

| Function | Role |
|----------|------|
| `accRegisterButton(window, id, label, x, y, w, h)` | Push button |
| `accRegisterCheckbox(window, id, label, x, y, w, h, checked?)` | Check box |
| `accRegisterRadio(window, id, label, x, y, w, h, selected?)` | Radio button |
| `accRegisterTextField(window, id, label, x, y, w, h, value)` | Text field |
| `accRegisterSlider(window, id, label, x, y, w, h, value)` | Slider |
| `accRegisterProgressBar(window, id, label, x, y, w, h, value)` | Progress bar |
| `accRegisterTab(window, id, label, x, y, w, h, selected?)` | Page tab |
| `accRegisterListItem(window, id, label, x, y, w, h, selected?)` | List item |
| `accRegisterLink(window, id, label, x, y, w, h)` | Hyperlink |
| `accRegisterStaticText(window, id, label, x, y, w, h)` | Static text |
| `accRegisterGroup(window, id, label, x, y, w, h)` | Groupbox |
| `accRegisterTreeItem(window, id, label, x, y, w, h, expanded?)` | Tree item |
| `accRegisterTable(window, id, label, x, y, w, h)` | Table |
| `accRegisterTableCell(window, id, label, x, y, w, h)` | Table cell |
| `accRegisterDropdown(window, id, label, x, y, w, h, expanded?)` | Combo box |

## Live Regions & Debugging

### `accAnnounce(window, id, text)`

Updates a live-region node's name and fires a change event so screen readers announce it immediately.

### `accDump(window)`

Dumps the accessible tree as a list of node entries for debugging.

### `accVerify(window)`

Verifies all registered widgets have names and roles. Returns a list of issues.

### `accRoleName(role)`

Returns a human-readable name for an MSAA role constant.

### `accStateName(state)`

Returns a human-readable description of MSAA state flags.

## Constants

### Roles

| Constant | Value |
|----------|-------|
| `ROLE_SYSTEM_PUSHBUTTON` | 43 |
| `ROLE_SYSTEM_CHECKBUTTON` | 44 |
| `ROLE_SYSTEM_RADIOBUTTON` | 45 |
| `ROLE_SYSTEM_TEXT` | 42 |
| `ROLE_SYSTEM_SLIDER` | 51 |
| `ROLE_SYSTEM_PROGRESSBAR` | 48 |
| `ROLE_SYSTEM_PAGETAB` | 37 |
| `ROLE_SYSTEM_LISTITEM` | 34 |
| `ROLE_SYSTEM_COMBOBOX` | 46 |
| `ROLE_SYSTEM_TABLE` | 24 |
| `ROLE_SYSTEM_CELL` | 29 |

### States

| Constant | Value |
|----------|-------|
| `STATE_SYSTEM_FOCUSED` | 4 |
| `STATE_SYSTEM_SELECTED` | 2 |
| `STATE_SYSTEM_CHECKED` | 16 |
| `STATE_SYSTEM_PRESSED` | 8 |
| `STATE_SYSTEM_EXPANDED` | 512 |
| `STATE_SYSTEM_COLLAPSED` | 1024 |
| `STATE_SYSTEM_READONLY` | 64 |
| `STATE_SYSTEM_FOCUSABLE` | 1048576 |
| `STATE_SYSTEM_SELECTABLE` | 2097152 |
| `STATE_SYSTEM_HASPOPUP` | 1073741824 |

### Events

| Constant | Value |
|----------|-------|
| `EVENT_OBJECT_NAMECHANGE` | 32780 |
| `EVENT_OBJECT_STATECHANGE` | 32778 |
| `EVENT_OBJECT_LOCATIONCHANGE` | 32779 |
| `EVENT_OBJECT_VALUECHANGE` | 32782 |
| `EVENT_OBJECT_FOCUS` | 32773 |
| `EVENT_OBJECT_SELECTION` | 32774 |
| `EVENT_OBJECT_LIVEREGIONCHANGED` | 32793 |
