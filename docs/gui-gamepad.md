# gui-gamepad — XInput Gamepad Support

`import('gui-gamepad')` provides XInput gamepad/joystick input support for up to 4 controllers with button detection, analog stick reading, vibration, and deadzone handling.

## Quick Start

```oak
gp := import('gui-gamepad')

// Poll state for change detection
ps := gp.gamepadPollState()

// In game loop:
gp.gamepadPoll(ps)

if gp.gamepadConnected?(ps, 0) -> {
    if gp.gamepadButtonPressed?(ps, 0, gp.XINPUT_GAMEPAD_A) -> {
        println('Player 1 pressed A!')
    }

    state := gp.applyThumbDeadzone(ps.current.0)
    println('Left stick X: ' + string(state.thumbLX))
}

// Rumble feedback
gp.setGamepadVibration(0, 32000, 32000)
```

## State Polling

### `getGamepadState(playerIndex)`

Polls the instantaneous state of a gamepad (0–3). Returns button state, triggers, and thumbstick values.

### `gamepadPollState()`

Creates a poll state tracker that stores current and previous frame states for all 4 controllers.

### `gamepadPoll(ps)`

Polls all 4 controllers, shifting current state to previous and reading new state.

### `gamepadConnected?(ps, idx)`

Returns `true` if the controller at index `idx` is currently connected.

### `gamepadJustConnected?(ps, idx)`

Returns `true` if the controller just connected this frame.

### `gamepadJustDisconnected?(ps, idx)`

Returns `true` if the controller just disconnected this frame.

## Button Queries

### `gamepadButtonDown?(state, button)`

Returns `true` if a button is currently held down.

### `gamepadButtonPressed?(ps, idx, button)`

Returns `true` if a button was just pressed this frame (edge-triggered).

### `gamepadButtonReleased?(ps, idx, button)`

Returns `true` if a button was just released this frame.

### D-Pad Helpers

`gamepadDpadUp?(state)`, `gamepadDpadDown?(state)`, `gamepadDpadLeft?(state)`, `gamepadDpadRight?(state)`

### Face Buttons

`gamepadA?(state)`, `gamepadB?(state)`, `gamepadX?(state)`, `gamepadY?(state)`

### Menu Buttons

`gamepadStart?(state)`, `gamepadBack?(state)`

### Shoulder & Thumb Buttons

`gamepadLeftShoulder?(state)`, `gamepadRightShoulder?(state)`, `gamepadLeftThumb?(state)`, `gamepadRightThumb?(state)`

## Analog Input

### `applyDeadzone(value, deadzone)`

Returns 0 if the value is within the deadzone, otherwise returns the normalized value.

### `applyThumbDeadzone(state)`

Returns a new state with deadzones applied to both thumbsticks.

## Vibration

### `setGamepadVibration(playerIndex, leftMotor, rightMotor)`

Sets rumble intensity (0–65535) for both motors.

### `stopGamepadVibration(playerIndex)`

Stops all vibration on a controller.

## Constants

### Button Masks

| Constant | Value |
|----------|-------|
| `XINPUT_GAMEPAD_DPAD_UP` | 1 |
| `XINPUT_GAMEPAD_DPAD_DOWN` | 2 |
| `XINPUT_GAMEPAD_DPAD_LEFT` | 4 |
| `XINPUT_GAMEPAD_DPAD_RIGHT` | 8 |
| `XINPUT_GAMEPAD_START` | 16 |
| `XINPUT_GAMEPAD_BACK` | 32 |
| `XINPUT_GAMEPAD_LEFT_THUMB` | 64 |
| `XINPUT_GAMEPAD_RIGHT_THUMB` | 128 |
| `XINPUT_GAMEPAD_LEFT_SHOULDER` | 256 |
| `XINPUT_GAMEPAD_RIGHT_SHOULDER` | 512 |
| `XINPUT_GAMEPAD_A` | 4096 |
| `XINPUT_GAMEPAD_B` | 8192 |
| `XINPUT_GAMEPAD_X` | 16384 |
| `XINPUT_GAMEPAD_Y` | 32768 |

### Deadzones

| Constant | Value |
|----------|-------|
| `XINPUT_GAMEPAD_LEFT_THUMB_DEADZONE` | 7849 |
| `XINPUT_GAMEPAD_RIGHT_THUMB_DEADZONE` | 8689 |
| `XINPUT_GAMEPAD_TRIGGER_THRESHOLD` | 30 |
| `XINPUT_VIBRATION_MAX` | 65535 |
