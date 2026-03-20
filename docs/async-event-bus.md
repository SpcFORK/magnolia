# Async Event Bus (async-event-bus)

## Overview

`async-event-bus` provides a lightweight in-process publish/subscribe event bus. Listeners can be registered for named events and notified either synchronously or asynchronously. One-time listeners are automatically removed after their first invocation.

## Import

```oak
bus := import('async-event-bus')
// or destructure
{ create: create } := import('async-event-bus')
```

## Factory

### `create()`

Creates and returns a new `EventBus` instance.

```oak
bus := create()
```

## EventBus Methods

### `on(event, handler)`

Registers a persistent listener for `event`. Returns a numeric token that can be used to remove this specific listener later.

```oak
token := bus.on('data', fn(payload, event) {
    printf('got {{0}}: {{1}}', event, payload)
})
```

### `once(event, handler)`

Registers a one-time listener for `event`. The handler is automatically removed after its first invocation.

```oak
bus.once('ready', fn(payload, _) {
    printf('ready: {{0}}', payload)
})
```

### `off(event, tokenOrHandler)`

Removes a listener. Accepts either the token returned by `on`/`once`, or the handler function itself. Returns the number of listeners removed.

```oak
bus.off('data', token)
// or by reference
bus.off('data', myHandler)
```

### `emit(event, payload)`

Emits `event` synchronously when called with two arguments. All matching listeners are called in registration order; one-time listeners are removed after invocation. Returns the number of listeners dispatched.

```oak
bus.emit('data', { x: 1 })
```

### `emit(event, payload, onDone)`

Emits `event` asynchronously when `onDone` is provided. Each listener is scheduled with `wait(0)` before being called. `onDone` receives the listener count when all listeners have been called.

```oak
bus.emit('data', { x: 1 }, fn(n) printf('dispatched to {{0}} listeners', n))
```

### `emitSync(event, payload)`

Always emits synchronously regardless of argument count. Returns the number of listeners dispatched.

```oak
count := bus.emitSync('tick', ?)
```

### `emitAsync(event, payload, onDone)`

Always emits asynchronously. `onDone` is optional; pass `?` to omit the completion callback.

```oak
bus.emitAsync('tick', ?, fn(n) printf('done, {{0}} handlers', n))
```

### `clear(event?)`

Removes all listeners for the given `event`, or all listeners on the bus when called without an argument.

```oak
bus.clear('data')   // remove all 'data' listeners
bus.clear()         // remove every listener
```

### `listeners(event)`

Returns a list of handler functions registered for `event`.

```oak
handlers := bus.listeners('data')
```

### `listenerCount(event)`

Returns the number of listeners registered for `event`.

```oak
n := bus.listenerCount('data') // => 2
```
