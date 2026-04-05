# gui-input — Virtual Key Code Constants

`gui-input` provides named constants for keyboard virtual key codes used by
`onKeyDown`, `onKeyUp`, and related GUI event handlers. It is the single
source of truth for key codes across all GUI modules (`gui-events`,
`gui-form`, `gui-test`) and samples.

## Usage

Key constants are re-exported via the `GUI` module, so you can use them
directly with the `gui.` prefix:

```oak
gui := import('GUI')

gui.onKeyDown(window, fn(vk) {
    if vk = gui.VK_ESCAPE -> gui.close(window)
    if vk = gui.VK_SPACE  -> togglePause()
    if vk = gui.VK_W | vk = gui.VK_UP -> moveForward()
})
```

Or import `gui-input` directly for standalone use:

```oak
input := import('gui-input')

if vk = input.VK_RETURN -> submit()
```

## Constants

### Modifier & Control Keys

| Constant      | Value | Key         |
|---------------|-------|-------------|
| `VK_BACK`     | 8     | Backspace   |
| `VK_TAB`      | 9     | Tab         |
| `VK_CLEAR`    | 12    | Clear       |
| `VK_RETURN`   | 13    | Enter       |
| `VK_SHIFT`    | 16    | Shift       |
| `VK_CONTROL`  | 17    | Ctrl        |
| `VK_ALT`      | 18    | Alt         |
| `VK_PAUSE`    | 19    | Pause       |
| `VK_CAPSLOCK` | 20    | Caps Lock   |
| `VK_ESCAPE`   | 27    | Escape      |
| `VK_SPACE`    | 32    | Space       |

### Navigation Keys

| Constant      | Value | Key        |
|---------------|-------|------------|
| `VK_PAGEUP`   | 33    | Page Up    |
| `VK_PAGEDOWN` | 34    | Page Down  |
| `VK_END`      | 35    | End        |
| `VK_HOME`     | 36    | Home       |
| `VK_LEFT`     | 37    | Left Arrow |
| `VK_UP`       | 38    | Up Arrow   |
| `VK_RIGHT`    | 39    | Right Arrow|
| `VK_DOWN`     | 40    | Down Arrow |

### Editing Keys

| Constant    | Value | Key    |
|-------------|-------|--------|
| `VK_INSERT` | 45    | Insert |
| `VK_DELETE` | 46    | Delete |

### Letter Keys (VK_A – VK_Z)

`VK_A` = 65 through `VK_Z` = 90. These match ASCII uppercase values.

### Digit Keys (VK_0 – VK_9)

`VK_0` = 48 through `VK_9` = 57. These match ASCII digit values.

### Numpad Keys

| Constant       | Value | Key       |
|----------------|-------|-----------|
| `VK_NUMPAD0`   | 96    | Numpad 0  |
| `VK_NUMPAD1`–`VK_NUMPAD9` | 97–105 | Numpad 1–9 |
| `VK_MULTIPLY`  | 106   | Numpad *  |
| `VK_ADD`       | 107   | Numpad +  |
| `VK_SEPARATOR` | 108   | Separator |
| `VK_SUBTRACT`  | 109   | Numpad -  |
| `VK_DECIMAL`   | 110   | Numpad .  |
| `VK_DIVIDE`    | 111   | Numpad /  |

### Function Keys

`VK_F1` = 112 through `VK_F12` = 123.

### Lock Keys

| Constant        | Value | Key         |
|-----------------|-------|-------------|
| `VK_NUMLOCK`    | 144   | Num Lock    |
| `VK_SCROLLLOCK` | 145   | Scroll Lock |

### OEM Keys (US Standard Keyboard)

| Constant           | Value | Key   |
|--------------------|-------|-------|
| `VK_OEM_SEMICOLON` | 186   | ; :   |
| `VK_OEM_PLUS`      | 187   | = +   |
| `VK_OEM_COMMA`     | 188   | , <   |
| `VK_OEM_MINUS`     | 189   | - _   |
| `VK_OEM_PERIOD`    | 190   | . >   |
| `VK_OEM_SLASH`     | 191   | / ?   |
| `VK_OEM_TILDE`     | 192   | ` ~   |
| `VK_OEM_LBRACKET`  | 219   | [ {   |
| `VK_OEM_BACKSLASH` | 220   | \ \|  |
| `VK_OEM_RBRACKET`  | 221   | ] }   |
| `VK_OEM_QUOTE`     | 222   | ' "   |

### Left/Right Modifier Keys

| Constant     | Value | Key         |
|--------------|-------|-------------|
| `VK_LSHIFT`  | 160   | Left Shift  |
| `VK_RSHIFT`  | 161   | Right Shift |
| `VK_LCONTROL`| 162   | Left Ctrl   |
| `VK_RCONTROL`| 163   | Right Ctrl  |
| `VK_LALT`    | 164   | Left Alt    |
| `VK_RALT`    | 165   | Right Alt   |

## Helper Functions

| Function           | Description                              |
|--------------------|------------------------------------------|
| `isLetterKey?(vk)` | True if vk is A–Z (65–90)               |
| `isDigitKey?(vk)`  | True if vk is 0–9 top row (48–57)       |
| `isNumpadKey?(vk)` | True if vk is numpad 0–9 (96–105)       |
| `isFunctionKey?(vk)`| True if vk is F1–F12 (112–123)          |
| `isArrowKey?(vk)`  | True if vk is an arrow key (37–40)      |
| `isModifierKey?(vk)`| True if vk is Shift, Ctrl, or Alt       |

## Cross-Platform Notes

These values correspond to Windows Virtual-Key codes, which Magnolia uses as
the canonical key representation on all platforms. On Linux (X11), the
`onKeyDown`/`onKeyUp` event handlers translate X11 keysyms to these same
values, so application code uses the same `VK_*` constants regardless of
platform.

## Migration from FORM_VK_*

The older `FORM_VK_*` constants in `gui-events.oak` (and re-exported via
`GUI.oak`) remain available for backward compatibility. They now delegate
to `gui-input` internally. New code should prefer the `VK_*` names:

| Old                 | New           |
|---------------------|---------------|
| `gui.FORM_VK_BACK`  | `gui.VK_BACK` |
| `gui.FORM_VK_TAB`   | `gui.VK_TAB`  |
| `gui.FORM_VK_RETURN` | `gui.VK_RETURN` |
| `gui.FORM_VK_SHIFT` | `gui.VK_SHIFT` |
| `gui.FORM_VK_CONTROL` | `gui.VK_CONTROL` |
| `gui.FORM_VK_ALT`   | `gui.VK_ALT`  |
| `gui.FORM_VK_ESCAPE` | `gui.VK_ESCAPE` |
