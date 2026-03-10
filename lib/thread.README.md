# Oak Threading Library (libthread)

A comprehensive concurrency and parallelism library for Oak, providing high-level abstractions for concurrent programming.

## Overview

`libthread` builds on Oak's built-in primitives (`go`, `make_chan`, `chan_send`, `chan_recv`) to provide ergonomic abstractions for common threading patterns including mutexes, semaphores, wait groups, thread pools, and parallel operations.

## Installation

Import the library in your Oak program:

```oak
thread := import('thread')
```

## Core Features

### 1. Goroutine Management

#### spawn(fn, args...)
Executes a function in a new goroutine with optional arguments.

```oak
thread.spawn(fn() println('Hello from thread!'))
thread.spawn(fn(x, y) println(x + y), 10, 20)
```

### 2. Channel Operations

#### makeChannel(size?)
Creates a channel with optional buffer size. Defaults to unbuffered (size = 0).

```oak
ch := thread.makeChannel()       // unbuffered
ch := thread.makeChannel(10)     // buffered with capacity 10
```

#### send(ch, value, callback?)
Sends a value to a channel (sync or async).

```oak
thread.send(ch, 'hello')                          // synchronous
thread.send(ch, 'hello', fn() println('sent!'))  // asynchronous
```

#### recv(ch, callback?)
Receives a value from a channel (sync or async).

```oak
data := thread.recv(ch)                              // synchronous
thread.recv(ch, fn(data) println('got: ', data))    // asynchronous
```

### 3. Synchronization Primitives

#### Mutex()
Provides mutual exclusion for protecting critical sections.

```oak
m := thread.Mutex()
m.lock()
// critical section
m.unlock()
```

#### Semaphore(n)
Counting semaphore for controlling access to a pool of resources.

```oak
sem := thread.Semaphore(3)  // allow 3 concurrent accesses
sem.acquire()
// use resource
sem.release()
```

#### WaitGroup()
Synchronizes multiple goroutines, waiting for all to complete.

```oak
wg := thread.WaitGroup()
wg.add(2)

thread.spawn(fn() {
    doWork()
    wg.done()
})

thread.spawn(fn() {
    doWork()
    wg.done()
})

wg.wait()  // blocks until both goroutines call done()
```

### 4. Async Patterns

#### Future(fn)
Represents a value that will be available in the future.

```oak
f := thread.Future(fn() expensiveComputation())
// do other work...
result := f.get()  // blocks until result is ready

// Chaining
f.then(fn(result) println('Got: ' << string(result)))
```

#### Pool(numWorkers)
Worker pool for managing concurrent tasks with a fixed number of workers.

```oak
pool := thread.Pool(4)  // 4 worker threads

pool.submit(fn() doWork1())
pool.submit(fn() doWork2())

pool.close()  // no more tasks
pool.wait()   // wait for all tasks to complete
```

### 5. Parallel Operations

#### parallel(fns)
Executes multiple functions concurrently and returns results in order.

```oak
results := thread.parallel([
    fn() computation1()
    fn() computation2()
    fn() computation3()
])
// results = [result1, result2, result3]
```

#### pmap(list, fn)
Parallel map - applies a function to each element in parallel.

```oak
squares := thread.pmap([1, 2, 3, 4], fn(x) x * x)
// squares = [1, 4, 9, 16]
```

#### pmapConcurrent(list, fn, maxConcurrent)
Like pmap but limits concurrency to n goroutines.

```oak
results := thread.pmapConcurrent([1, 2, 3, 4], fn(x) x * x, 2)
// Uses only 2 goroutines at a time
```

#### race(fns)
Returns the result of the first function to complete.

```oak
fastest := thread.race([
    fn() slowComputation()
    fn() fastComputation()
])
```

### 6. Advanced Patterns

#### pipeline(input, stages...)
Creates a concurrent processing pipeline.

```oak
result := thread.pipeline(
    10,
    fn(x) x * 2,
    fn(x) x + 1,
    fn(x) x * x
)
// result = ((10 * 2) + 1)^2 = 441
```

#### retry(fn, maxAttempts)
Retries a function up to maxAttempts times until it succeeds.

```oak
result := thread.retry(fn() unstableNetworkCall(), 3)
```

#### debounce(fn, waitTime)
Delays execution until waitTime seconds after the last call.

```oak
debouncedSave := thread.debounce(fn() saveData(), 0.5)
// Calling multiple times within 0.5s will only execute once
```

#### throttle(fn, waitTime)
Executes at most once every waitTime seconds.

```oak
throttledLog := thread.throttle(fn() logMessage(), 1)
// Will execute at most once per second
```

## Examples

### Example 1: Concurrent Web Scraping

```oak
thread := import('thread')

urls := [
    'https://example.com/page1'
    'https://example.com/page2'
    'https://example.com/page3'
]

// Scrape all URLs concurrently with max 2 concurrent requests
results := thread.pmapConcurrent(urls, fn(url) {
    // fetch and parse URL
    fetchAndParse(url)
}, 2)
```

### Example 2: Producer-Consumer Pattern

```oak
thread := import('thread')

ch := thread.makeChannel(10)
wg := thread.WaitGroup()

// Producer
wg.add(1)
thread.spawn(fn() {
    range(100) |> each(fn(i) thread.send(ch, i))
    wg.done()
})

// Consumer
wg.add(1)
thread.spawn(fn() {
    range(100) |> each(fn(_) {
        item := thread.recv(ch)
        process(item.data)
    })
    wg.done()
})

wg.wait()
```

### Example 3: Worker Pool

```oak
thread := import('thread')

pool := thread.Pool(5)

tasks := [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
tasks |> each(fn(task) {
    pool.submit(fn() processTask(task))
})

pool.close()
pool.wait()
```

### Example 4: Parallel Computation with Wait Group

```oak
thread := import('thread')

results := []
wg := thread.WaitGroup()
m := thread.Mutex()

range(10) |> each(fn(i) {
    wg.add(1)
    thread.spawn(fn() {
        result := compute(i)
        
        m.lock()
        results << result
        m.unlock()
        
        wg.done()
    })
})

wg.wait()
println('All computations complete')
```

## Performance Considerations

1. **Buffered vs Unbuffered Channels**: Use buffered channels when you know the capacity in advance to reduce blocking.

2. **Pool Size**: Choose pool size based on workload:
   - CPU-bound tasks: number of CPU cores
   - I/O-bound tasks: higher number (10-100+)

3. **Goroutine Overhead**: Each goroutine has memory overhead. For millions of tasks, use worker pools instead of spawning individual goroutines.

4. **Memory**: Channels and goroutines consume memory. Close/cleanup when done.

## Thread Safety

- All channel operations are thread-safe
- Mutex and Semaphore provide synchronization
- Shared data should be protected with mutexes or communicated via channels
- Follow the principle: "Don't communicate by sharing memory; share memory by communicating"

## API Reference

| Function | Description |
|----------|-------------|
| `spawn(fn, args...)` | Execute function in new goroutine |
| `makeChannel(size?)` | Create channel |
| `send(ch, value, cb?)` | Send to channel |
| `recv(ch, cb?)` | Receive from channel |
| `Mutex()` | Create mutex |
| `Semaphore(n)` | Create semaphore with n tokens |
| `WaitGroup()` | Create wait group |
| `Future(fn)` | Create future |
| `Pool(n)` | Create worker pool with n workers |
| `parallel(fns)` | Run functions in parallel |
| `pmap(list, fn)` | Parallel map |
| `pmapConcurrent(list, fn, n)` | Parallel map with concurrency limit |
| `race(fns)` | Return result of fastest function |
| `pipeline(input, stages...)` | Create processing pipeline |
| `retry(fn, n)` | Retry function n times |
| `debounce(fn, t)` | Debounce function by t seconds |
| `throttle(fn, t)` | Throttle function to once per t seconds |

## License

Same as Oak language (MIT).
