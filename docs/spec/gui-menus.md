# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-menus.oak`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `MF_STRING` · `0`
- `MF_SEPARATOR` · `2048`
- `MF_POPUP` · `16`
- `MF_CHECKED` · `8`
- `MF_UNCHECKED` · `0`
- `MF_GRAYED` · `1`
- `MF_ENABLED` · `0`
- `MF_BYCOMMAND` · `0`
- `MF_BYPOSITION` · `1024`
- `TPM_LEFTBUTTON` · `0`
- `TPM_RIGHTBUTTON` · `2`
- `TPM_RETURNCMD` · `256`
- `WM_COMMAND` · `273`
### `createMenu()`

### `createPopupMenu()`

### `appendItem(hMenu, id, label)`

### `appendSeparator(hMenu)`

### `appendSubmenu(hMenu, hSubMenu, label)`

### `setMenuBar(window, hMenu)`

> returns `:bool`

### `removeMenuBar(window)`

> returns `:bool`

### `destroyMenu(hMenu)`

> returns `:bool`

### `showPopupMenu(window, hMenu, x, y)`

### `onMenuCommand(window, handler)`

### `buildMenu(spec)`

- `FVIRTKEY` · `1`
- `FSHIFT` · `4`
- `FCONTROL` · `8`
- `FALT` · `16`
### `createAcceleratorTable(entries)`

### `destroyAcceleratorTable(hAccel)`

> returns `:bool`

### `translateAccelerator(window, hAccel)`

### `installAccelerators(window, entries)`

> returns `:object`

- `WM_SYSCOMMAND` · `274`
### `getSystemMenu(window, reset)`

### `appendSystemMenuItem(window, id, label)`

### `appendSystemMenuSeparator(window)`

### `resetSystemMenu(window)`

### `onSysCommand(window, handler)`

