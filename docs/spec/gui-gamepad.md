# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-gamepad.oak`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_xinputProcCache` · `{}`
- `_xinputDll` · `?`
### `_resolveXInputDll()`

### `xinput(symbol, args...)`

- `XUSER_MAX_COUNT` · `4`
- `ERROR_SUCCESS` · `0`
- `ERROR_DEVICE_NOT_CONNECTED` · `1167`
- `XINPUT_GAMEPAD_DPAD_UP` · `1`
- `XINPUT_GAMEPAD_DPAD_DOWN` · `2`
- `XINPUT_GAMEPAD_DPAD_LEFT` · `4`
- `XINPUT_GAMEPAD_DPAD_RIGHT` · `8`
- `XINPUT_GAMEPAD_START` · `16`
- `XINPUT_GAMEPAD_BACK` · `32`
- `XINPUT_GAMEPAD_LEFT_THUMB` · `64`
- `XINPUT_GAMEPAD_RIGHT_THUMB` · `128`
- `XINPUT_GAMEPAD_LEFT_SHOULDER` · `256`
- `XINPUT_GAMEPAD_RIGHT_SHOULDER` · `512`
- `XINPUT_GAMEPAD_A` · `4096`
- `XINPUT_GAMEPAD_B` · `8192`
- `XINPUT_GAMEPAD_X` · `16384`
- `XINPUT_GAMEPAD_Y` · `32768`
- `XINPUT_GAMEPAD_LEFT_THUMB_DEADZONE` · `7849`
- `XINPUT_GAMEPAD_RIGHT_THUMB_DEADZONE` · `8689`
- `XINPUT_GAMEPAD_TRIGGER_THRESHOLD` · `30`
- `XINPUT_VIBRATION_MAX` · `65535`
- `_XINPUT_STATE_SIZE` · `16`
- `_XINPUT_VIBRATION_SIZE` · `4`
### `_readI16(base, off)`

### `getGamepadState(playerIndex)`

> returns `:object`

### `gamepadButtonDown?(state, button)`

> returns `:bool`

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

### `gamepadLeftThumb?(state)`

### `gamepadRightThumb?(state)`

### `applyDeadzone(value, deadzone)`

> returns `:int`

### `applyThumbDeadzone(state)`

> returns `:object`

### `setGamepadVibration(playerIndex, leftMotor, rightMotor)`

### `stopGamepadVibration(playerIndex)`

### `gamepadPollState()`

> returns `:object`

### `gamepadPoll(ps)`

### `gamepadButtonPressed?(ps, idx, button)`

> returns `:bool`

### `gamepadButtonReleased?(ps, idx, button)`

> returns `:bool`

### `gamepadConnected?(ps, idx)`

### `gamepadJustConnected?(ps, idx)`

> returns `:bool`

### `gamepadJustDisconnected?(ps, idx)`

> returns `:bool`

