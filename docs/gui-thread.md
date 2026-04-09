# gui-thread — GUI Threading Primitives

`import('gui-thread')` provides thread-safety primitives and parallel helpers for the GUI rendering pipeline. Worker threads perform computation (mesh transforms, layout, etc.) and post results back to the main thread via command queues, preventing direct Win32/GDI calls from worker threads.

## Quick Start

```oak
gt := import('gui-thread')
gui := import('GUI')

window := gui.createWindow('Threaded App', 800, 600, {})

// Attach threading infrastructure to a window
gt.initWindowThreading(window, { numWorkers: 4 })

// In your render loop:
sched := gt.scheduler(window)

sched.beginFrame()
sched.dispatch(fn() heavyComputation1())
sched.dispatch(fn() heavyComputation2())
results := sched.endFrame()
sched.flush()

// Transform vertices in parallel
transformed := gt.parallelTransformVertices(vertices, fn(v) {
    // apply matrix transform
}, 4)

// Cleanup on close
gt.destroyWindowThreading(window)
```

## Window Integration

### `initWindowThreading(window, options)`

Attaches threading infrastructure (command queue, worker pool, scheduler, state guard) to a window.

**Options:**
- `numWorkers` — number of worker threads (default: 4)
- `commandQueue` — enable command queue (default: `true`)
- `workerPool` — enable worker pool (default: `true`)

### `threadingEnabled?(window)`

Returns `true` if threading is enabled on the window.

### `commandQueue(window)` / `workerPool(window)` / `scheduler(window)` / `stateGuard(window)`

Accessors for the threading objects attached to a window.

### `flushThreadedCommands(window)`

Flushes all pending threaded commands (call from the frame loop).

### `destroyWindowThreading(window)`

Cleans up and releases all threading resources.

## CommandQueue

Created via `CommandQueue()`. Ensures worker thread results are consumed on the main thread.

| Method | Description |
|--------|-------------|
| `post(cmdFn)` | Enqueue a zero-arg function to run on the main thread |
| `postValue(tag, value)` | Enqueue a tagged value for main-thread collection |
| `drain()` | Collect all pending commands without blocking |
| `flush()` | Execute all pending command functions (call once per frame) |
| `size()` | Approximate number of pending commands |

## WorkerPool

Created via `WorkerPool(numWorkers)`. Offloads computation to background goroutines.

| Method | Description |
|--------|-------------|
| `submit(taskFn)` | Send a computation task to the pool |
| `submitWithResult(taskFn)` | Send a task; returns a Future for the result |
| `pmap(list, fn)` | Apply function to each element in parallel |
| `pmapN(list, fn, maxN)` | Apply function with limited concurrency |
| `batch(list, chunkSize, processFn)` | Split list into chunks and process in parallel |
| `close()` | Shut down the pool |
| `size()` | Number of workers |

## FrameFence

Created via `FrameFence(workerCount)`. Synchronizes worker completion per frame.

| Method | Description |
|--------|-------------|
| `arrive()` | Signal one worker has completed |
| `awaitWorkers()` | Block until all workers have arrived |
| `count()` | Number of expected workers |

## StateGuard

Created via `StateGuard()`. Mutex-based wrapper for shared state.

| Method | Description |
|--------|-------------|
| `read(readerFn)` | Lock, call reader, unlock, return result |
| `write(writerFn)` | Lock, call writer, unlock |
| `lock()` / `unlock()` | Manual lock/unlock |

## AsyncLoader

Created via `AsyncLoader(cmdQueue)`. Loads resources in background goroutines.

| Method | Description |
|--------|-------------|
| `load(loaderFn, onComplete)` | Start background load; calls onComplete with result |
| `busy?()` | Returns `true` if loads are in flight |
| `count()` | Number of active loads |

## FrameScheduler

Created via `FrameScheduler(pool, cmdQueue)`. Dispatches per-frame tasks to workers.

| Method | Description |
|--------|-------------|
| `beginFrame()` | Reset per-frame task tracking |
| `dispatch(taskFn)` | Submit a task for this frame |
| `endFrame()` | Block until all tasks complete; return results |
| `flush()` | Drain the command queue |

## Vertex Transform Helper

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

Transforms mesh vertices in parallel batches. Falls back to inline transform if vertex count is too small.

```oak
transformed := gt.parallelTransformVertices(verts, fn(v) {
    // return transformed vertex
}, 4)
```

## Notes

- Worker threads must **not** call Win32/GDI/GPU functions directly. Use the command queue to post results back to the main thread.
- The FrameScheduler pattern is: `beginFrame()` → `dispatch(...)` → `endFrame()` → `flush()`.
