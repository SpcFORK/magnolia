# gui-menus — Win32 Menu Bars & Context Menus

`import('gui-menus')` provides Win32 menu bar and context (popup) menu support with keyboard accelerators, system menu customization, and declarative menu building.

## Quick Start

```oak
menus := import('gui-menus')
gui := import('GUI')

window := gui.createWindow('Menu Demo', 800, 600, {})

// Declarative menu bar
hMenu := menus.buildMenu([
    { label: '&File', items: [
        { label: '&New', id: 1 }
        { label: '&Open...', id: 2 }
        { label: :separator }
        { label: 'E&xit', id: 99 }
    ]}
    { label: '&Edit', items: [
        { label: '&Undo', id: 10 }
        { label: '&Redo', id: 11 }
    ]}
])
menus.setMenuBar(window, hMenu)

// Handle menu commands
menus.onMenuCommand(window, fn(cmdId) if cmdId {
    99 -> println('Exit!')
    _ -> println('Command: ' + string(cmdId))
})

// Keyboard accelerators
menus.installAccelerators(window, [
    { key: 78, cmd: 1, ctrl: true }   // Ctrl+N → New
    { key: 79, cmd: 2, ctrl: true }   // Ctrl+O → Open
])
```

## Menu Construction

### `createMenu()`

Creates a menu bar handle.

### `createPopupMenu()`

Creates a popup (context) menu handle.

### `appendItem(hMenu, id, label)`

Adds a text item with a command id.

### `appendSeparator(hMenu)`

Adds a separator line.

### `appendSubmenu(hMenu, hSubMenu, label)`

Adds a submenu with a text label.

### `buildMenu(spec)`

Builds a complete menu from a declarative spec. Each top-level entry is `{ label, items: [...] }` where items are `{ label, id }` or `{ label: :separator }`.

### `destroyMenu(hMenu)`

Frees a menu handle.

## Menu Bar & Popup

### `setMenuBar(window, hMenu)`

Attaches a menu bar to a window.

### `removeMenuBar(window)`

Detaches the menu bar from a window.

### `showPopupMenu(window, hMenu, x, y)`

Displays a context menu at screen coordinates `(x, y)`. Returns the selected command id, or 0 if cancelled.

### `onMenuCommand(window, handler)`

The handler receives `(commandId)` when a menu item is selected.

## Keyboard Accelerators

### `createAcceleratorTable(entries)`

Creates an accelerator table from entries: `{key, cmd, ctrl?, shift?, alt?}`.

### `destroyAcceleratorTable(hAccel)`

Frees an accelerator table handle.

### `translateAccelerator(window, hAccel)`

Translates an accelerator keystroke. Returns `true` if handled.

### `installAccelerators(window, entries)`

Creates an accelerator table and hooks it into the poll loop.

## System Menu

### `getSystemMenu(window, reset)`

Gets the system (title bar) menu handle.

### `appendSystemMenuItem(window, id, label)`

Adds an item to the system menu.

### `appendSystemMenuSeparator(window)`

Adds a separator to the system menu.

### `resetSystemMenu(window)`

Restores the system menu to its default state.

### `onSysCommand(window, handler)`

The handler receives `(commandId)` for custom system menu items.

## Constants

### Menu Item Flags

| Constant | Value |
|----------|-------|
| `MF_STRING` | 0 |
| `MF_SEPARATOR` | 2048 |
| `MF_POPUP` | 16 |
| `MF_CHECKED` | 8 |
| `MF_GRAYED` | 1 |
| `MF_BYCOMMAND` | 0 |
| `MF_BYPOSITION` | 1024 |

### Accelerator Flags

| Constant | Value |
|----------|-------|
| `FVIRTKEY` | 1 |
| `FSHIFT` | 4 |
| `FCONTROL` | 8 |
| `FALT` | 16 |
